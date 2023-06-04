// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IDelegationShare.sol";
import "./IDelegationSlasher.sol";
import "./IDelegation.sol";

/**
 * @title Interface for the primary entrypoint for funds.
 * @author Layr Labs, Inc.
 * @notice See the `DelegationManager` contract itself for implementation details.
 */
interface IDelegationManager {
    // used for storing details of queued withdrawals
    struct WithdrawalStorage {
        uint32 initTimestamp;
        uint32 unlockTimestamp;
        address withdrawer;
    }

    // packed struct for queued withdrawals
    struct WithdrawerAndNonce {
        address withdrawer;
        uint96 nonce;
    }

    /**
     * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
     * In functions that operate on existing queued withdrawals -- e.g. `startQueuedWithdrawalWaitingPeriod` or `completeQueuedWithdrawal`,
     * the data is resubmitted and the hash of the submitted data is computed by `calculateWithdrawalRoot` and checked against the
     * stored hash in order to confirm the integrity of the submitted data.
     */
    struct QueuedWithdrawal {
        IDelegationShare[] delegations;
        IERC20[] tokens;
        uint256[] shares;
        address depositor;
        WithdrawerAndNonce withdrawerAndNonce;
        address delegatedAddress;
    }

    /**
     * @notice Deposits `amount` of `token` into the specified `DelegationShare`, with the resultant shares credited to `depositor`
     * @param delegationShare is the specified shares record where investment is to be made,
     * @param token is the ERC20 token in which the investment is to be made,
     * @param amount is the amount of token to be invested in the delegationShare by the depositor
     */
    function depositInto(IDelegationShare delegationShare, IERC20 token, uint256 amount)
        external
        returns (uint256);

    /// @notice Returns the current shares of `user` in `delegationShare`
    function investorDelegationShares(address user, IDelegationShare delegationShare) external view returns (uint256 shares);

    /**
     * @notice Get all details on the depositor's investments and corresponding shares
     * @return (depositor's delegationShare record, shares in these DelegationShare contract)
     */
    function getDeposits(address depositor) external view returns (IDelegationShare[] memory, uint256[] memory);

    /// @notice Simple getter function that returns `investorDelegations[staker].length`.
    function investorDelegationLength(address staker) external view returns (uint256);

    /**
     * @notice Called by a staker to queue a withdraw in the given token and shareAmount from each of the respective given strategies.
     * @dev Stakers will complete their withdrawal by calling the 'completeQueuedWithdrawal' function.
     * User shares are decreased in this function, but the total number of shares in each delegation strategy remains the same.
     * The total number of shares is decremented in the 'completeQueuedWithdrawal' function instead, which is where
     * the funds are actually sent to the user through use of the delegation strategies' 'withdrawal' function. This ensures
     * that the value per share reported by each strategy will remain consistent, and that the shares will continue
     * to accrue gains during the enforced WITHDRAWAL_WAITING_PERIOD.
     * @param delegationShareIndexes is a list of the indices in `investorDelegationShare[msg.sender]` that correspond to the delegation strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @dev strategies are removed from `delegationShare` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `delegationShares`. The simplest way to calculate the correct `delegationShareIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `delegationShares` to lowest index
     */
    function queueWithdrawal(
        uint256[] calldata delegationShareIndexes,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata shareAmounts,
        WithdrawerAndNonce calldata withdrawerAndNonce,
        bool undelegateIfPossible
    )
        external returns(bytes32);

    function startQueuedWithdrawalWaitingPeriod(
        bytes32 withdrawalRoot,
        uint32 stakeInactiveAfter
    ) external;

    /**
     * @notice Used to complete the specified `queuedWithdrawal`. The function caller must match `queuedWithdrawal.withdrawer`
     * @param queuedWithdrawal The QueuedWithdrawal to complete.
     * @param receiveAsTokens If true, the shares specified in the queued withdrawal will be withdrawn from the specified delegation strategies themselves
     * and sent to the caller, through calls to `queuedWithdrawal.strategies[i].withdraw`. If false, then the shares in the specified delegation strategies
     * will simply be transferred to the caller directly.
     */
    function completeQueuedWithdrawal(
        QueuedWithdrawal calldata queuedWithdrawal,
        bool receiveAsTokens
    )
        external;

    /**
     * @notice Slashes the shares of 'frozen' operator (or a staker delegated to one)
     * @param slashedAddress is the frozen address that is having its shares slashes
     * @param delegationShareIndexes is a list of the indices in `investorStrats[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @param recipient The slashed funds are withdrawn as tokens to this address.
     * @dev strategies are removed from `investorStrats` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `investorStrats`. The simplest way to calculate the correct `strategyIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `investorStrats` to lowest index
     */
    function slashShares(
        address slashedAddress,
        address recipient,
        IDelegationShare[] calldata delegationShares,
        IERC20[] calldata tokens,
        uint256[] calldata delegationShareIndexes,
        uint256[] calldata shareAmounts
    )
        external;

    function slashQueuedWithdrawal(
        address recipient,
        QueuedWithdrawal calldata queuedWithdrawal
    )
        external;

    /**
     * @notice Used to check if a queued withdrawal can be completed. Returns 'true' if the withdrawal can be immediately
     * completed, and 'false' otherwise.
     * @dev This function will revert if the specified `queuedWithdrawal` does not exist
     */
    function canCompleteQueuedWithdrawal(
        QueuedWithdrawal calldata queuedWithdrawal
    )
        external
        returns (bool);

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.
    function calculateWithdrawalRoot(
        QueuedWithdrawal memory queuedWithdrawal
    )
        external
        pure
        returns (bytes32);

    /// @notice Returns the single, central Delegation contract
    function delegation() external view returns (IDelegation);

    /// @notice Returns the single, central DelegationSlasher contract
    function delegationSlasher() external view returns (IDelegationSlasher);
}
