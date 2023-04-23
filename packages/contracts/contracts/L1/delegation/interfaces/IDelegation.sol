// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./IDelegationCallback.sol";

/**
 * @title Interface for the primary delegation contract.
 * @notice See the `Delegation` contract itself for implementation details.
 */
interface IDelegation {
    enum DelegationStatus {
        UNDELEGATED,
        DELEGATED
    }

    /**
     * @notice This will be called by an operator to register itself as an operator that stakers can choose to delegate to.
     * @param dt is the `DelegationTerms` contract that the operator has for those who delegate to them.
     * @dev An operator can set `dt` equal to their own address (or another EOA address), in the event that they want to split payments
     * in a more 'trustful' manner.
     * @dev In the present design, once set, there is no way for an operator to ever modify the address of their DelegationTerms contract.
     */
    function registerAsOperator(IDelegationCallback dt) external;

    /**
     *  @notice This will be called by a staker to delegate its assets to some operator.
     *  @param operator is the operator to whom staker (msg.sender) is delegating its assets
     */
    function delegateTo(address operator) external;

    /**
     * @notice Delegates from `staker` to `operator`.
     * @dev requires that r, vs are a valid ECSDA signature from `staker` indicating their intention for this action
     */
    function delegateToSignature(address staker, address operator, uint256 expiry, bytes32 r, bytes32 vs) external;

    /**
     * @notice Undelegates `staker` from the operator who they are delegated to.
     * @notice Callable only by the InvestmentManager
     * @dev Should only ever be called in the event that the `staker` has no active deposits.
     */
    function undelegate(address staker) external;

    /// @notice returns the address of the operator that `staker` is delegated to.
    function delegatedTo(address staker) external view returns (address);

    /// @notice returns the delegationCallback of the `operator`, which may mediate their interactions with stakers who delegate to them.
    function delegationCallback(address operator) external view returns (IDelegationCallback);

    /// @notice returns the total number of shares in `DelegationShare` that are delegated to `operator`.
    function operatorShares(address operator, IDelegationShare delegationShare) external view returns (uint256);

    /// @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
    function isDelegated(address staker) external view returns (bool);

    /// @notice Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.
    function isNotDelegated(address staker) external returns (bool);

    /// @notice Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.
    function isOperator(address operator) external view returns (bool);

    /**
     * @notice Increases the `staker`'s delegated shares in `delegationShare` by `shares`, typically called when the staker has further deposits.
     * @dev Callable only by the DelegationManager
     */
    function increaseDelegatedShares(address staker, IDelegationShare delegationShare, uint256 shares) external;

    /**
     * @notice Decreases the `staker`'s delegated shares in `delegationShare` by `shares, typically called when the staker withdraws
     * @dev Callable only by the DelegationManager
     */
    function decreaseDelegatedShares(address staker, IDelegationShare delegationShare, uint256 shares) external;

    /// @notice Version of `decreaseDelegatedShares` that accepts an array of inputs.
    function decreaseDelegatedShares(
        address staker,
        IDelegationShare[] calldata delegationShares,
        uint256[] calldata shares
    ) external;
}
