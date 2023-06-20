// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "./DelegationStorage.sol";
import "./DelegationSlasher.sol";
import "./WhiteListBase.sol";
/**
 * @title The primary delegation contract.
 * @notice  This is the contract for delegation. The main functionalities of this contract are
 * - for enabling any staker to register as a delegate and specify the delegation terms it has agreed to
 * - for enabling anyone to register as an operator
 * - for a registered staker to delegate its stake to the operator of its agreed upon delegation terms contract
 * - for a staker to undelegate its assets
 * - for anyone to challenge a staker's claim to have fulfilled all its obligation before undelegation
 */
abstract contract Delegation is Initializable, OwnableUpgradeable, PausableUpgradeable, WhiteList, DelegationStorage {
    /// @notice Simple permission for functions that are only callable by the InvestmentManager contract.
    modifier onlyDelegationManager() {
        require(msg.sender == address(delegationManager), "onlyDelegationManager");
        _;
    }

    // INITIALIZING FUNCTIONS
    constructor(IDelegationManager _delegationManager)
        DelegationStorage(_delegationManager)
    {
        _disableInitializers();
    }

    /// @dev Emitted when a low-level call to `delegationTerms.onDelegationReceived` fails, returning `returnData`
    event OnDelegationReceivedCallFailure(IDelegationCallback indexed delegationTerms, bytes32 returnData);

    /// @dev Emitted when a low-level call to `delegationTerms.onDelegationWithdrawn` fails, returning `returnData`
    event OnDelegationWithdrawnCallFailure(IDelegationCallback indexed delegationTerms, bytes32 returnData);

    event RegisterOperator(address delegationCallback, address register);

    event DelegateTo(address delegatior, address operator);

    event DecreaseDelegatedShares(address delegatedShare, address operator, uint256 share);

    event IncreaseDelegatedShares(address delegatedShare, address operator, uint256 share);

    function initialize(address initialOwner)
        external
        initializer
    {
        DOMAIN_SEPARATOR = keccak256(abi.encode(DOMAIN_TYPEHASH, bytes("Mantle"), block.chainid, address(this)));
        _transferOwnership(initialOwner);
    }

    // PERMISSION FUNCTIONS
    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    // EXTERNAL FUNCTIONS
    /**
     * @notice This will be called by an operator to register itself as an operator that stakers can choose to delegate to.
     * @param dt is the `DelegationTerms` contract that the operator has for those who delegate to them.
     * @dev An operator can set `dt` equal to their own address (or another EOA address), in the event that they want to split payments
     * in a more 'trustful' manner.
     * @dev In the present design, once set, there is no way for an operator to ever modify the address of their DelegationTerms contract.
     */
    function registerAsOperator(IDelegationCallback dt) external whitelistOnly(msg.sender) {
        require(
            address(delegationCallback[msg.sender]) == address(0),
            "Delegation.registerAsOperator: Delegate has already registered"
        );
        // store the address of the delegation contract that the operator is providing.
        delegationCallback[msg.sender] = dt;
        _delegate(msg.sender, msg.sender);
        emit RegisterOperator(address(dt),msg.sender);
    }

    /**
     *  @notice This will be called by a staker to delegate its assets to some operator.
     *  @param operator is the operator to whom staker (msg.sender) is delegating its assets
     */
    function delegateTo(address operator) external whenNotPaused {
        _delegate(msg.sender, operator);
    }

    /**
     * @notice Delegates from `staker` to `operator`.
     * @dev requires that r, vs are a valid ECSDA signature from `staker` indicating their intention for this action
     */
    function delegateToSignature(address staker, address operator, uint256 expiry, bytes32 r, bytes32 vs)
        external
        whenNotPaused
    {
        require(expiry == 0 || expiry >= block.timestamp, "delegation signature expired");
        // calculate struct hash, then increment `staker`'s nonce
        // EIP-712 standard
        bytes32 structHash = keccak256(abi.encode(DELEGATION_TYPEHASH, staker, operator, nonces[staker]++, expiry));
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));
        //check validity of signature

        address recoveredAddress = ECDSA.recover(digestHash, r, vs);

        require(recoveredAddress == staker, "Delegation.delegateToBySignature: sig not from staker");
        _delegate(staker, operator);
    }

    /**
     * @notice Undelegates `staker` from the operator who they are delegated to.
     * @notice Callable only by the InvestmentManager
     * @dev Should only ever be called in the event that the `staker` has no active deposits.
     */
    function undelegate(address staker) external onlyDelegationManager {
        delegationStatus[staker] = DelegationStatus.UNDELEGATED;
        delegatedTo[staker] = address(0);
    }

    /**
     * @notice Increases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker has further deposits
     * @dev Callable only by the InvestmentManager
     */
    function increaseDelegatedShares(address staker, IDelegationShare delegationShare, uint256 shares)
        external
        onlyDelegationManager
    {
        //if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // add strategy shares to delegate's shares
            operatorShares[operator][delegationShare] += shares;

            //Calls into operator's delegationTerms contract to update weights of individual staker
            IDelegationShare[] memory investorDelegations = new IDelegationShare[](1);
            uint256[] memory investorShares = new uint[](1);
            investorDelegations[0] = delegationShare;
            investorShares[0] = shares;

            // call into hook in delegationCallback contract
            IDelegationCallback dt = delegationCallback[operator];
            _delegationReceivedHook(dt, staker, operator, investorDelegations, investorShares);
            emit IncreaseDelegatedShares(address(delegationShare), operator, shares);
        }
    }

    /**
     * @notice Decreases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker withdraws
     * @dev Callable only by the InvestmentManager
     */
    function decreaseDelegatedShares(address staker, IDelegationShare delegationShare, uint256 shares)
        external
        onlyDelegationManager
    {
        //if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegate's shares
            operatorShares[operator][delegationShare] -= shares;

            //Calls into operator's delegationCallback contract to update weights of individual staker
            IDelegationShare[] memory investorDelegationShares = new IDelegationShare[](1);
            uint256[] memory investorShares = new uint[](1);
            investorDelegationShares[0] = delegationShare;
            investorShares[0] = shares;

            // call into hook in delegationCallback contract
            IDelegationCallback dt = delegationCallback[operator];
            _delegationWithdrawnHook(dt, staker, operator, investorDelegationShares, investorShares);
            emit DecreaseDelegatedShares(address(delegationShare), operator, shares);
        }
    }

    /// @notice Version of `decreaseDelegatedShares` that accepts an array of inputs.
    function decreaseDelegatedShares(
        address staker,
        IDelegationShare[] calldata strategies,
        uint256[] calldata shares
    )
        external
        onlyDelegationManager
    {
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegate's shares
            uint256 stratsLength = strategies.length;
            for (uint256 i = 0; i < stratsLength;) {
                operatorShares[operator][strategies[i]] -= shares[i];
                emit DecreaseDelegatedShares(address(strategies[i]), operator, shares[i]);
                unchecked {
                    ++i;
                }
            }

            // call into hook in delegationCallback contract
            IDelegationCallback dt = delegationCallback[operator];
            _delegationWithdrawnHook(dt, staker, operator, strategies, shares);
        }
    }

    // INTERNAL FUNCTIONS

    /**
     * @notice Makes a low-level call to `dt.onDelegationReceived(staker, strategies, shares)`, ignoring reverts and with a gas budget
     * equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).
     * @dev *If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
     * `returnData` is *only the first 32 bytes* returned by the call to `dt`.
     */
    function _delegationReceivedHook(
        IDelegationCallback dt,
        address staker,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory shares
    )
        internal
    {
        /**
         * We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
         * In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.
         */
        // format calldata
        (bool success, bytes memory returnData) = address(dt).call{gas: LOW_LEVEL_GAS_BUDGET}(
            abi.encodeWithSelector(IDelegationCallback.onDelegationReceived.selector, staker, operator, delegationShares, shares)
        );

        // if the call fails, we emit a special event rather than reverting
        if (!success) {
            emit OnDelegationReceivedCallFailure(dt, returnData[0]);
        }
    }

    /**
     * @notice Makes a low-level call to `dt.onDelegationWithdrawn(staker, strategies, shares)`, ignoring reverts and with a gas budget
     * equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).
     * @dev *If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
     * `returnData` is *only the first 32 bytes* returned by the call to `dt`.
     */
    function _delegationWithdrawnHook(
        IDelegationCallback dt,
        address staker,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory shares
    )
        internal
    {
        /**
         * We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
         * In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.
         */

        (bool success, bytes memory returnData) = address(dt).call{gas: LOW_LEVEL_GAS_BUDGET}(
            abi.encodeWithSelector(IDelegationCallback.onDelegationWithdrawn.selector, staker, operator, delegationShares, shares)
        );

        // if the call fails, we emit a special event rather than reverting
        if (!success) {
            emit OnDelegationWithdrawnCallFailure(dt, returnData[0]);
        }
    }

    /**
     * @notice Internal function implementing the delegation *from* `staker` *to* `operator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @dev Ensures that the operator has registered as a delegate (`address(dt) != address(0)`), verifies that `staker` is not already
     * delegated, and records the new delegation.
     */
    function _delegate(address staker, address operator) internal {

        IDelegationCallback dt = delegationCallback[operator];
        require(
            address(dt) != address(0), "Delegation._delegate: operator has not yet registered as a delegate"
        );
        require(isNotDelegated(staker), "Delegation._delegate: staker has existing delegation");

        // checks that operator has not been frozen
        IDelegationSlasher slasher = delegationManager.delegationSlasher();
        require(!slasher.isFrozen(operator), "Delegation._delegate: cannot delegate to a frozen operator");
        // record delegation relation between the staker and operator
        delegatedTo[staker] = operator;

        // record that the staker is delegated
        delegationStatus[staker] = DelegationStatus.DELEGATED;
        // retrieve list of strategies and their shares from investment manager
        (IDelegationShare[] memory delegationShares, uint256[] memory shares) = delegationManager.getDeposits(staker);

        // add strategy shares to delegate's shares
        uint256 delegationLength = delegationShares.length;
        for (uint256 i = 0; i < delegationLength;) {
            // update the share amounts for each of the operator's strategies
            operatorShares[operator][delegationShares[i]] += shares[i];
            unchecked {
                ++i;
            }
        }
        // call into hook in delegationCallback contract
        _delegationReceivedHook(dt, staker, operator, delegationShares, shares);
        emit DelegateTo(staker, operator);
    }

    // VIEW FUNCTIONS

    /// @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
    function isDelegated(address staker) public view returns (bool) {
        return (delegationStatus[staker] == DelegationStatus.DELEGATED);
    }

    /// @notice Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.
    function isNotDelegated(address staker) public view returns (bool) {
        return (delegationStatus[staker] == DelegationStatus.UNDELEGATED);
    }

    /// @notice Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.
    function isOperator(address operator) external view returns (bool) {
        return (address(delegationCallback[operator]) != address(0));
    }
}
