// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../delegation/DelegationManager.sol";
import "../ITssGroupManager.sol";
import "../TssStakingSlashing.sol";


/**
 * @title The primary entry- and exit-point for funds into and out.
 * @notice This contract is for managing investments in different strategies. The main
 * functionalities are:
 * - adding and removing investment strategies that any delegator can invest into
 * - enabling deposit of assets into specified investment delegation(s)
 * - enabling removal of assets from specified investment delegation(s)
 * - recording deposit of ETH into settlement layer
 * - recording deposit for securing
 * - slashing of assets for permissioned strategies
 */
contract TssDelegationManager is DelegationManager {


    address public stakingSlash;
    address public tssGroupManager;

    uint256 public minStakeAmount;




    /**
     * @param _delegation The delegation contract.
     * @param _delegationSlasher The primary slashing contract.
     */
    constructor(IDelegation _delegation, IDelegationSlasher _delegationSlasher)
    DelegationManager(_delegation, _delegationSlasher)
    {
        _disableInitializers();
    }

    function initialize(address _stakingSlashing,
        address _tssGroupManager,
        uint256 _minStakeAmount
    ) public initializer {
        stakingSlash = _stakingSlashing;
        tssGroupManager = _tssGroupManager;
        minStakeAmount = _minStakeAmount;
    }


    modifier onlyStakingSlash() {
        require(msg.sender == stakingSlash, "contract call is not staking slashing");
        _;
    }

    function setStakingSlash(address _address) public onlyOwner {
        stakingSlash = _address;
    }

    function setMinStakeAmount(uint256 _amount) public onlyOwner {
        minStakeAmount = _amount;
    }

    function setTssGroupManager(address _addr) public onlyOwner {
        tssGroupManager = _addr;
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
        override
        whenNotPaused
        onlyStakingSlash
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

    function queueWithdrawal(
        uint256[] calldata delegationIndexes,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata shares,
        WithdrawerAndNonce calldata withdrawerAndNonce,
        bool undelegateIfPossible
    )
    external
    override
    whenNotPaused
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (bytes32)
    {
        require(
            withdrawerAndNonce.nonce == numWithdrawalsQueued[msg.sender],
            "InvestmentManager.queueWithdrawal: provided nonce incorrect"
        );
        require(delegationShares.length == 1, "only tss delegation share");
        require(shares.length == 1,"only tss delegation share");
        // increment the numWithdrawalsQueued of the sender
        unchecked {
            ++numWithdrawalsQueued[msg.sender];
        }
        _checkMinStakeAmount(msg.sender,delegationShares[0],shares[0]);

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


    function isCanOperator(address _addr, IDelegationShare delegationShare) external returns (bool)  {
        if (delegation.isOperator(_addr)) {
            uint256 share = delegation.operatorShares(_addr, delegationShare);
            uint256 balance = delegationShare.sharesToUnderlying(share);
            if (balance > minStakeAmount) {
                return true;
            }
        }
        return false;
    }

    function depositInto(IDelegationShare delegationShare, IERC20 token, uint256 amount, address sender)
    external
    onlyNotFrozen(sender)
    nonReentrant
    whitelistOnly(address(delegationShare))
    onlyStakingSlash
    returns (uint256 shares)
    {
        shares = _depositInto(sender, delegationShare, token, amount);
    }

    function queueWithdrawal(
        address sender,
        uint256[] calldata delegationIndexes,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata shares,
        WithdrawerAndNonce calldata withdrawerAndNonce
    )
    external
    whenNotPaused
    onlyNotFrozen(sender)
    onlyStakingSlash
    nonReentrant
    returns (bytes32)
    {
        require(
            withdrawerAndNonce.nonce == numWithdrawalsQueued[sender],
            "InvestmentManager.queueWithdrawal: provided nonce incorrect"
        );
        require(delegationShares.length == 1, "only tss delegation share");
        require(shares.length == 1,"only tss delegation share");
        // increment the numWithdrawalsQueued of the sender
        unchecked {
            ++numWithdrawalsQueued[sender];
        }
        address operator = delegation.delegatedTo(sender);

        _checkMinStakeAmount(sender, delegationShares[0], shares[0]);

        // modify delegated shares accordingly, if applicable
        delegation.decreaseDelegatedShares(sender, delegationShares, shares);

        // the internal function will return 'true' in the event the delegation contrat was
        // removed from the depositor's array of strategies -- i.e. investorStrats[depositor]
        _removeShares(sender, delegationIndexes[0], delegationShares[0], shares[0]);

        // copy arguments into struct and pull delegation info
        QueuedWithdrawal memory queuedWithdrawal = QueuedWithdrawal({
            delegations: delegationShares,
            tokens: tokens,
            shares: shares,
            depositor: sender,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: operator
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

        address staker = sender;
        // If the `msg.sender` has withdrawn all of their funds in this transaction, then they can choose to also undelegate
        /**
         * Checking that `investorStrats[msg.sender].length == 0` is not strictly necessary here, but prevents reverting very late in logic,
         * in the case that 'undelegate' is set to true but the `msg.sender` still has active deposits.
         */
        if (investorDelegations[staker].length == 0) {
            _undelegate(staker);
        }

        emit WithdrawalQueued(staker, withdrawerAndNonce.withdrawer, operator, withdrawalRoot);

        return withdrawalRoot;
    }


    function startQueuedWithdrawalWaitingPeriod(bytes32 withdrawalRoot, address sender, uint32 stakeInactiveAfter) external onlyStakingSlash {
        require(
            queuedWithdrawals[withdrawalRoot].unlockTimestamp == QUEUED_WITHDRAWAL_INITIALIZED_VALUE,
            "InvestmentManager.startQueuedWithdrawalWaitingPeriod: Withdrawal stake inactive claim has already been made"
        );
        require(
            queuedWithdrawals[withdrawalRoot].withdrawer == sender,
            "InvestmentManager.startQueuedWithdrawalWaitingPeriod: Sender is not the withdrawer"
        );
        require(
            block.timestamp > queuedWithdrawals[withdrawalRoot].initTimestamp,
            "InvestmentManager.startQueuedWithdrawalWaitingPeriod: Stake may still be subject to slashing based on new tasks. Wait to set stakeInactiveAfter."
        );
        //they can only unlock after a withdrawal waiting period or after they are claiming their stake is inactive
        queuedWithdrawals[withdrawalRoot].unlockTimestamp = max((uint32(block.timestamp) + WITHDRAWAL_WAITING_PERIOD), stakeInactiveAfter);
    }

    function getWithdrawNonce(address staker) external view onlyStakingSlash returns (uint256) {
        return numWithdrawalsQueued[staker];
    }

    function getDelegationShares(address staker,IDelegationShare delegationShare) external view onlyStakingSlash returns (uint256) {
        return investorDelegationShares[staker][delegationShare];
    }

    function _checkMinStakeAmount(address sender,IDelegationShare delegationShare, uint256 shares) internal {
        address operator = delegation.delegatedTo(sender);
        // check if the operator is still mpc node, if the remaining shares meet the mini requirement
        if (delegation.isDelegated(sender)){
            require(!TssStakingSlashing(stakingSlash).isJailed(operator),"the operator is in jail status");

            uint256 rest= delegation.operatorShares(operator, delegationShare) - shares;
            uint256 balance = delegationShare.sharesToUnderlying(rest);
            if (ITssGroupManager(tssGroupManager).isTssGroupUnJailMembers(operator)) {
                require(balance > minStakeAmount,"unable withdraw due to operator's rest shares smaller than mini requirement");
            }
        }
    }


}

