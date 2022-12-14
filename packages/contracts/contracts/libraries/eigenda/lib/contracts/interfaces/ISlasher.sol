// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

/**
 * @title Interface for the primary 'slashing' contract for EigenLayr.
 * @author Layr Labs, Inc.
 * @notice See the `Slasher` contract itself for implementation details.
 */
interface ISlasher {
    /**
     * @notice Gives the `contractAddress` permission to slash the funds of the caller.
     * @dev Typically, this function must be called prior to registering for a middleware.
     */
    function allowToSlash(address contractAddress) external;

    /// @notice Called by a contract to revoke its ability to slash `operator`, once `unbondedAfter` is reached.
    function revokeSlashingAbility(address operator, uint32 unbondedAfter) external;

    /**
     * @notice Used for 'slashing' a certain operator.
     * @param toBeFrozen The operator to be frozen.
     * @dev Technically the operator is 'frozen' (hence the name of this function), and then subject to slashing pending a decision by a human-in-the-loop.
     * @dev The operator must have previously given the caller (which should be a contract) the ability to slash them, through a call to `allowToSlash`.
     */
    function freezeOperator(address toBeFrozen) external;

    /**
     * @notice Used to determine whether `staker` is actively 'frozen'. If a staker is frozen, then they are potentially subject to
     * slashing of their funds, and cannot cannot deposit or withdraw from the investmentManager until the slashing process is completed
     * and the staker's status is reset (to 'unfrozen').
     * @return Returns 'true' if `staker` themselves has their status set to frozen, OR if the staker is delegated
     * to an operator who has their status set to frozen. Otherwise returns 'false'.
     */
    function isFrozen(address staker) external view returns (bool);

    /// @notice Returns true if `slashingContract` is currently allowed to slash `toBeSlashed`.
    function canSlash(address toBeSlashed, address slashingContract) external view returns (bool);

    /// @notice Returns the UTC timestamp until which `slashingContract` is allowed to slash the `operator`.
    function bondedUntil(address operator, address slashingContract) external view returns (uint32);

    /**
     * @notice Removes the 'frozen' status from each of the `frozenAddresses`
     * @dev Callable only by the contract owner (i.e. governance).
     */
    function resetFrozenStatus(address[] calldata frozenAddresses) external;

    /**
     * @notice Used to give global slashing permission to `contracts`.
     * @dev Callable only by the contract owner (i.e. governance).
     */
    function addGloballyPermissionedContracts(address[] calldata contracts) external;

    /**
     * @notice Used to revoke global slashing permission from `contracts`.
     * @dev Callable only by the contract owner (i.e. governance).
     */
    function removeGloballyPermissionedContracts(address[] calldata contracts) external;
}
