// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "./interfaces/IDelegation.sol";
import "./DelegationManagerStorage.sol";
import "./WhiteListBase.sol";
/**
 * @title The primary entry- and exit-point for funds into and out.
 * @author Layr Labs, Inc.
 * @notice This contract is for managing investments in different strategies. The main
 * functionalities are:
 * - adding and removing investment strategies that any delegator can invest into
 * - enabling deposit of assets into specified investment delegation(s)
 * - enabling removal of assets from specified investment delegation(s)
 * - recording deposit of ETH into settlement layer
 * - recording deposit for securing
 * - slashing of assets for permissioned strategies
 */
abstract contract DelegationManager is
    Initializable,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    DelegationManagerStorage,
    WhiteList
{
    using SafeERC20 for IERC20;

    /**
     * @notice Value to which `initTimestamp` and `unlockTimestamp` to is set to indicate a withdrawal is queued/initialized,
     * but has not yet had its waiting period triggered
     */
    uint32 internal constant QUEUED_WITHDRAWAL_INITIALIZED_VALUE = type(uint32).max;

    /**
     * @notice Emitted when a new withdrawal is queued by `depositor`.
     * @param depositor Is the staker who is withdrawing funds.
     * @param withdrawer Is the party specified by `staker` who will be able to complete the queued withdrawal and receive the withdrawn funds.
     * @param delegatedAddress Is the party who the `staker` was delegated to at the time of creating the queued withdrawal
     * @param withdrawalRoot Is a hash of the input data for the withdrawal.
     */
    event WithdrawalQueued(
        address indexed depositor, address indexed withdrawer, address indexed delegatedAddress, bytes32 withdrawalRoot
    );

    /// @notice Emitted when a queued withdrawal is completed
    event WithdrawalCompleted(address indexed depositor, address indexed withdrawer, bytes32 withdrawalRoot);

    modifier onlyNotFrozen(address staker) {
        require(
            !delegationSlasher.isFrozen(staker),
            "DelegationManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"
        );
        _;
    }

    modifier onlyFrozen(address staker) {
        require(delegationSlasher.isFrozen(staker), "DelegationManager.onlyFrozen: staker has not been frozen");
        _;
    }

    /**
     * @param _delegation The delegation contract.
     * @param _delegationSlasher The primary slashing contract.
     */
    constructor(IDelegation _delegation, IDelegationSlasher _delegationSlasher)
        DelegationManagerStorage(_delegation, _delegationSlasher)
    {
        _disableInitializers();
    }

    // EXTERNAL FUNCTIONS

    /**
     * @notice Initializes the investment manager contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
     * and transfers contract ownership to the specified `initialOwner`.
     * @param initialOwner Ownership of this contract is transferred to this address.
     */
    function initialize(address initialOwner)
        external
        initializer
    {
        DOMAIN_SEPARATOR = keccak256(abi.encode(DOMAIN_TYPEHASH, bytes("Mantle"), block.chainid, address(this)));
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Deposits `amount` of `token` into the specified `delegationShare`, with the resultant shares credited to `depositor`
     * @param delegationShare is the specified delegation contract where investment is to be made,
     * @param token is the denomination in which the investment is to be made,
     * @param amount is the amount of token to be invested in the delegation contract by the depositor
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
     */
    function depositInto(IDelegationShare delegationShare, IERC20 token, uint256 amount)
        external
        onlyNotFrozen(msg.sender)
        nonReentrant
        whitelistOnly(address(delegationShare))
        returns (uint256 shares)
    {
        shares = _depositInto(msg.sender, delegationShare, token, amount);
    }

    /**
     * @notice Called by a staker to undelegate entirely. The staker must first withdraw all of their existing deposits
     * (through use of the `queueWithdrawal` function), or else otherwise have never deposited prior to delegating.
     */
    function undelegate() external {
        _undelegate(msg.sender);
    }

    /**
     * @notice Called by a staker to queue a withdraw in the given token and shareAmount from each of the respective given strategies.
     * @dev Stakers will complete their withdrawal by calling the 'completeQueuedWithdrawal' function.
     * User shares are decreased in this function, but the total number of shares in each delegation contract remains the same.
     * The total number of shares is decremented in the 'completeQueuedWithdrawal' function instead, which is where
     * the funds are actually sent to the user through use of the strategies' 'withdrawal' function. This ensures
     * that the value per share reported by each delegation contract will remain consistent, and that the shares will continue
     * to accrue gains during the enforced WITHDRAWAL_WAITING_PERIOD.
     * @param delegationIndexes is a list of the indices in `investorStrats[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @dev strategies are removed from `investorStrats` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `investorStrats`. The simplest way to calculate the correct `delegationIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `investorStrats` to lowest index
     */
    function queueWithdrawal(
        uint256[] calldata delegationIndexes,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata shares,
        WithdrawerAndNonce calldata withdrawerAndNonce,
        bool undelegateIfPossible
    )
        external
        virtual
        whenNotPaused
        onlyNotFrozen(msg.sender)
        nonReentrant
        returns (bytes32)
    {
        require(
            withdrawerAndNonce.nonce == numWithdrawalsQueued[msg.sender],
            "DelegationManager.queueWithdrawal: provided nonce incorrect"
        );
        // increment the numWithdrawalsQueued of the sender
        unchecked {
            ++numWithdrawalsQueued[msg.sender];
        }

        uint256 delegationIndex;

        // modify delegated shares accordingly, if applicable
        delegation.decreaseDelegatedShares(msg.sender, delegationShares, shares);

        for (uint256 i = 0; i < delegationShares.length;) {
            // the internal function will return 'true' in the event the delegation contrat was
            // removed from the depositor's array of strategies -- i.e. investorStrats[depositor]
            if (_removeShares(msg.sender, delegationIndexes[delegationIndex], delegationShares[i], shares[i])) {
                unchecked {
                    ++delegationIndex;
                }
            }

            //increment the loop
            unchecked {
                ++i;
            }
        }

        // fetch the address that the `msg.sender` is delegated to
        address delegatedAddress = delegation.delegatedTo(msg.sender);

        // copy arguments into struct and pull delegation info
        QueuedWithdrawal memory queuedWithdrawal = QueuedWithdrawal({
            delegations: delegationShares,
            tokens: tokens,
            shares: shares,
            depositor: msg.sender,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: delegatedAddress
        });

        // calculate the withdrawal root
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);

        //update storage in mapping of queued withdrawals
        queuedWithdrawals[withdrawalRoot] = WithdrawalStorage({
            /**
             * @dev We add `REASONABLE_STAKES_UPDATE_PERIOD` to the current time here to account for the fact that it may take some time for
             * the operator's stake to be updated on all the middlewares. New tasks created between now at this 'initTimestamp' may still
             * subject the `msg.sender` to slashing!
             */
            initTimestamp: uint32(block.timestamp + REASONABLE_STAKES_UPDATE_PERIOD),
            withdrawer: withdrawerAndNonce.withdrawer,
            unlockTimestamp: QUEUED_WITHDRAWAL_INITIALIZED_VALUE
        });

        // If the `msg.sender` has withdrawn all of their funds in this transaction, then they can choose to also undelegate
        /**
         * Checking that `investorStrats[msg.sender].length == 0` is not strictly necessary here, but prevents reverting very late in logic,
         * in the case that 'undelegate' is set to true but the `msg.sender` still has active deposits.
         */
        if (undelegateIfPossible && investorDelegations[msg.sender].length == 0) {
            _undelegate(msg.sender);
        }

        emit WithdrawalQueued(msg.sender, withdrawerAndNonce.withdrawer, delegatedAddress, withdrawalRoot);

        return withdrawalRoot;
    }

    /*
    *
    * @notice The withdrawal flow is:
    * - Depositor starts a queued withdrawal, setting the receiver of the withdrawn funds as withdrawer
    * - Withdrawer then waits for the queued withdrawal tx to be included in the chain, and then sets the stakeInactiveAfter. This cannot
    *   be set when starting the queued withdrawal, as it is there may be transactions the increase the tasks upon which the stake is active
    *   that get mined before the withdrawal.
    * - The withdrawer completes the queued withdrawal after the stake is inactive or a withdrawal fraud proof period has passed,
    *   whichever is longer. They specify whether they would like the withdrawal in shares or in tokens.
    */
    function startQueuedWithdrawalWaitingPeriod(bytes32 withdrawalRoot, uint32 stakeInactiveAfter) external virtual {
        require(
            queuedWithdrawals[withdrawalRoot].unlockTimestamp == QUEUED_WITHDRAWAL_INITIALIZED_VALUE,
            "DelegationManager.startQueuedWithdrawalWaitingPeriod: Withdrawal stake inactive claim has already been made"
        );
        require(
            queuedWithdrawals[withdrawalRoot].withdrawer == msg.sender,
            "DelegationManager.startQueuedWithdrawalWaitingPeriod: Sender is not the withdrawer"
        );
        require(
            block.timestamp > queuedWithdrawals[withdrawalRoot].initTimestamp,
            "DelegationManager.startQueuedWithdrawalWaitingPeriod: Stake may still be subject to slashing based on new tasks. Wait to set stakeInactiveAfter."
        );
        //they can only unlock after a withdrawal waiting period or after they are claiming their stake is inactive
        queuedWithdrawals[withdrawalRoot].unlockTimestamp = max((uint32(block.timestamp) + WITHDRAWAL_WAITING_PERIOD), stakeInactiveAfter);
    }

    /**
     * @notice Used to complete the specified `queuedWithdrawal`. The function caller must match `queuedWithdrawal.withdrawer`
     * @param queuedWithdrawal The QueuedWithdrawal to complete.
     * @param receiveAsTokens If true, the shares specified in the queued withdrawal will be withdrawn from the specified strategies themselves
     * and sent to the caller, through calls to `queuedWithdrawal.delegations[i].withdraw`. If false, then the shares in the specified strategies
     * will simply be transferred to the caller directly.
     */
    function completeQueuedWithdrawal(QueuedWithdrawal calldata queuedWithdrawal, bool receiveAsTokens)
        external
        whenNotPaused
        // check that the address that the staker *was delegated to* – at the time that they queued the withdrawal – is not frozen
        onlyNotFrozen(queuedWithdrawal.delegatedAddress)
        nonReentrant
    {
        // find the withdrawalRoot
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);
        // copy storage to memory
        WithdrawalStorage memory withdrawalStorageCopy = queuedWithdrawals[withdrawalRoot];

        // verify that the queued withdrawal actually exists
        require(
            withdrawalStorageCopy.unlockTimestamp != 0,
            "DelegationManager.completeQueuedWithdrawal: withdrawal does not exist"
        );

        require(
            uint32(block.timestamp) >= withdrawalStorageCopy.unlockTimestamp
                || (queuedWithdrawal.delegatedAddress == address(0)),
            "DelegationManager.completeQueuedWithdrawal: withdrawal waiting period has not yet passed and depositor was delegated when withdrawal initiated"
        );

        // TODO: add testing coverage for this
        require(
            msg.sender == queuedWithdrawal.withdrawerAndNonce.withdrawer,
            "DelegationManager.completeQueuedWithdrawal: only specified withdrawer can complete a queued withdrawal"
        );

        // reset the storage slot in mapping of queued withdrawals
        delete queuedWithdrawals[withdrawalRoot];

        // store length for gas savings
        uint256 strategiesLength = queuedWithdrawal.delegations.length;
        // if the withdrawer has flagged to receive the funds as tokens, withdraw from strategies
        if (receiveAsTokens) {
            // actually withdraw the funds
            for (uint256 i = 0; i < strategiesLength;) {
                // tell the delegation to send the appropriate amount of funds to the depositor
                queuedWithdrawal.delegations[i].withdraw(
                    withdrawalStorageCopy.withdrawer, queuedWithdrawal.tokens[i], queuedWithdrawal.shares[i]
                );
                unchecked {
                    ++i;
                }
            }
        } else {
            // else increase their shares
            for (uint256 i = 0; i < strategiesLength;) {
                _addShares(withdrawalStorageCopy.withdrawer, queuedWithdrawal.delegations[i], queuedWithdrawal.shares[i]);
                unchecked {
                    ++i;
                }
            }
        }

        emit WithdrawalCompleted(queuedWithdrawal.depositor, withdrawalStorageCopy.withdrawer, withdrawalRoot);
    }

    /**
     * @notice Slashes the shares of a 'frozen' operator (or a staker delegated to one)
     * @param slashedAddress is the frozen address that is having its shares slashed
     * @param delegationIndexes is a list of the indices in `investorStrats[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @param recipient The slashed funds are withdrawn as tokens to this address.
     * @dev delegationShares are removed from `investorStrats` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `investorStrats`. The simplest way to calculate the correct `delegationIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `investorStrats` to lowest index
     */
    function slashShares(
        address slashedAddress,
        address recipient,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata delegationIndexes,
        uint256[] calldata shareAmounts
    )
        external
        virtual
        whenNotPaused
        onlyOwner
        onlyFrozen(slashedAddress)
        nonReentrant
    {
        uint256 delegationIndex;
        uint256 strategiesLength = delegationShares.length;
        for (uint256 i = 0; i < strategiesLength;) {
            // the internal function will return 'true' in the event the delegation contract was
            // removed from the slashedAddress's array of strategies -- i.e. investorStrats[slashedAddress]
            if (_removeShares(slashedAddress, delegationIndexes[delegationIndex], delegationShares[i], shareAmounts[i])) {
                unchecked {
                    ++delegationIndex;
                }
            }

            // withdraw the shares and send funds to the recipient
            delegationShares[i].withdraw(recipient, tokens[i], shareAmounts[i]);

            // increment the loop
            unchecked {
                ++i;
            }
        }

        // modify delegated shares accordingly, if applicable
        delegation.decreaseDelegatedShares(slashedAddress, delegationShares, shareAmounts);
    }

    /**
     * @notice Slashes an existing queued withdrawal that was created by a 'frozen' operator (or a staker delegated to one)
     * @param recipient The funds in the slashed withdrawal are withdrawn as tokens to this address.
     */
    function slashQueuedWithdrawal(address recipient, QueuedWithdrawal calldata queuedWithdrawal)
        external
        whenNotPaused
        onlyOwner
        nonReentrant
    {
        // find the withdrawalRoot
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);

        // verify that the queued withdrawal actually exists
        require(
            queuedWithdrawals[withdrawalRoot].unlockTimestamp != 0,
            "DelegationManager.slashQueuedWithdrawal: withdrawal does not exist"
        );

        // verify that *either* the queued withdrawal has been successfully challenged, *or* the `depositor` has been frozen
        require(
            queuedWithdrawals[withdrawalRoot].withdrawer == address(0) || delegationSlasher.isFrozen(queuedWithdrawal.depositor),
            "DelegationManager.slashQueuedWithdrawal: withdrawal has not been successfully challenged or depositor is not frozen"
        );

        // reset the storage slot in mapping of queued withdrawals
        delete queuedWithdrawals[withdrawalRoot];

        uint256 strategiesLength = queuedWithdrawal.delegations.length;
        for (uint256 i = 0; i < strategiesLength;) {
            // tell the delegation contract to send the appropriate amount of funds to the recipient
            queuedWithdrawal.delegations[i].withdraw(recipient, queuedWithdrawal.tokens[i], queuedWithdrawal.shares[i]);
            unchecked {
                ++i;
            }
        }
    }

    // INTERNAL FUNCTIONS

    /**
     * @notice This function adds `shares` for a given `delegationShare` to the `depositor` and runs through the necessary update logic.
     * @dev In particular, this function calls `delegation.increaseDelegatedShares(depositor, delegationShare, shares)` to ensure that all
     * delegated shares are tracked, increases the stored share amount in `investorStratShares[depositor][delegationShare]`, and adds `delegationShare`
     * to the `depositor`'s list of strategies, if it is not in the list already.
     */
    function _addShares(address depositor, IDelegationShare delegationShare, uint256 shares) internal {
        // sanity check on `shares` input
        require(shares != 0, "DelegationManager._addShares: shares should not be zero!");

        // if they dont have existing shares of this delegation contract, add it to their strats
        if (investorDelegationShares[depositor][delegationShare] == 0) {
            require(
                investorDelegations[depositor].length <= MAX_INVESTOR_DELEGATION_LENGTH,
                "DelegationManager._addShares: deposit would exceed MAX_INVESTOR_DELEGATION_LENGTH"
            );
            investorDelegations[depositor].push(delegationShare);
        }

        // add the returned shares to their existing shares for this delegation contract
        investorDelegationShares[depositor][delegationShare] += shares;

        // if applicable, increase delegated shares accordingly
        delegation.increaseDelegatedShares(depositor, delegationShare, shares);
    }

    /**
     * @notice Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the InvestmentDelegation-type contract
     * `delegationShare`, with the resulting shares credited to `depositor`.
     * @return shares The amount of *new* shares in `delegationShare` that have been credited to the `depositor`.
     */
    function _depositInto(address depositor, IDelegationShare delegationShare, IERC20 token, uint256 amount)
        internal
        returns (uint256 shares)
    {

        // transfer tokens from the sender to the delegation contract
        token.safeTransferFrom(depositor, address(delegationShare), amount);

        // deposit the assets into the specified delegation contract and get the equivalent amount of shares in that delegation contract
        shares = delegationShare.deposit(depositor, token, amount);

        // add the returned shares to the depositor's existing shares for this delegation contract
        _addShares(depositor, delegationShare, shares);

        return shares;
    }

    /**
     * @notice Decreases the shares that `depositor` holds in `delegationShare` by `shareAmount`.
     * @dev If the amount of shares represents all of the depositor`s shares in said delegation contract,
     * then the delegation contract is removed from investorStrats[depositor] and 'true' is returned. Otherwise 'false' is returned.
     */
    function _removeShares(address depositor, uint256 delegationIndex, IDelegationShare delegationShare, uint256 shareAmount)
        internal
        returns (bool)
    {
        // sanity check on `shareAmount` input
        require(shareAmount != 0, "DelegationManager._removeShares: shareAmount should not be zero!");

        //check that the user has sufficient shares
        uint256 userShares = investorDelegationShares[depositor][delegationShare];


        require(shareAmount <= userShares, "DelegationManager._removeShares: shareAmount too high");
        //unchecked arithmetic since we just checked this above
        unchecked {
            userShares = userShares - shareAmount;
        }

        // subtract the shares from the depositor's existing shares for this delegation contract
        investorDelegationShares[depositor][delegationShare] = userShares;
        // if no existing shares, remove is from this investors strats

        if (userShares == 0) {
            // remove the delegation contract from the depositor's dynamic array of strategies
            _removeDelegationFromInvestorDelegations(depositor, delegationIndex, delegationShare);

            // return true in the event that the delegation contract was removed from investorStrats[depositor]
            return true;
        }
        // return false in the event that the delegation contract was *not* removed from investorStrats[depositor]
        return false;
    }

    /**
     * @notice Removes `delegationShare` from `depositor`'s dynamic array of strategies, i.e. from `investorStrats[depositor]`
     * @dev the provided `delegationIndex` input is optimistically used to find the delegation contract quickly in the list. If the specified
     * index is incorrect, then we revert to a brute-force search.
     */
    function _removeDelegationFromInvestorDelegations(address depositor, uint256 delegationIndex, IDelegationShare delegationShare) internal {
        // if the delegation contract matches with the delegation contract index provided
        if (investorDelegations[depositor][delegationIndex] == delegationShare) {
            // replace the delegation contract with the last delegation contract in the list
            investorDelegations[depositor][delegationIndex] =
            investorDelegations[depositor][investorDelegations[depositor].length - 1];
        } else {
            //loop through all of the strategies, find the right one, then replace
            uint256 delegationLength = investorDelegations[depositor].length;

            for (uint256 j = 0; j < delegationLength;) {
                if (investorDelegations[depositor][j] == delegationShare) {
                    //replace the delegation contract with the last delegation contract in the list
                    investorDelegations[depositor][j] = investorDelegations[depositor][investorDelegations[depositor].length - 1];
                    break;
                }
                unchecked {
                    ++j;
                }
            }
        }

        // pop off the last entry in the list of strategies
        investorDelegations[depositor].pop();
    }

    /**
     * @notice If the `depositor` has no existing shares, then they can `undelegate` themselves.
     * This allows people a "hard reset" in their relationship after withdrawing all of their stake.
     */
    function _undelegate(address depositor) internal {
        require(investorDelegations[depositor].length == 0, "InvestmentManager._undelegate: depositor has active deposits");
        delegation.undelegate(depositor);
    }

    function max(uint32 x, uint32 y) internal pure returns (uint32) {
        return x > y ? x : y;
    }

    // VIEW FUNCTIONS

    /**
     * @notice Used to check if a queued withdrawal can be completed. Returns 'true' if the withdrawal can be immediately
     * completed, and 'false' otherwise.
     * @dev This function will revert if the specified `queuedWithdrawal` does not exist
     */
    function canCompleteQueuedWithdrawal(QueuedWithdrawal calldata queuedWithdrawal) external view returns (bool) {
        // find the withdrawalRoot
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);

        // verify that the queued withdrawal actually exists
        require(
            queuedWithdrawals[withdrawalRoot].unlockTimestamp != 0,
            "DelegationManager.canCompleteQueuedWithdrawal: withdrawal does not exist"
        );

        if (delegationSlasher.isFrozen(queuedWithdrawal.delegatedAddress)) {
            return false;
        }

        return (
            uint32(block.timestamp) >= queuedWithdrawals[withdrawalRoot].unlockTimestamp
                || (queuedWithdrawal.delegatedAddress == address(0))
        );
    }

    /**
     * @notice Get all details on the depositor's investments and corresponding shares
     * @return (depositor's strategies, shares in these strategies)
     */
    function getDeposits(address depositor) external view returns (IDelegationShare[] memory, uint256[] memory) {
        uint256 delegationLength = investorDelegations[depositor].length;
        uint256[] memory shares = new uint256[](delegationLength);

        for (uint256 i = 0; i < delegationLength;) {
            shares[i] = investorDelegationShares[depositor][investorDelegations[depositor][i]];
            unchecked {
                ++i;
            }
        }
        return (investorDelegations[depositor], shares);
    }

    /// @notice Simple getter function that returns `investorStrats[staker].length`.
    function investorDelegationLength(address staker) external view returns (uint256) {
        return investorDelegations[staker].length;
    }

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.
    function calculateWithdrawalRoot(QueuedWithdrawal memory queuedWithdrawal) public pure returns (bytes32) {
        return (
            keccak256(
                abi.encode(
                    queuedWithdrawal.delegations,
                    queuedWithdrawal.tokens,
                    queuedWithdrawal.shares,
                    queuedWithdrawal.depositor,
                    queuedWithdrawal.withdrawerAndNonce,
                    queuedWithdrawal.delegatedAddress
                )
            )
        );
    }
}
