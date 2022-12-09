// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IEigenLayrDelegation.sol";

/**
 * @title Interface for a `ServiceManager`-type contract.
 * @author Layr Labs, Inc.
 */
// TODO: provide more functions for this spec
interface IServiceManager {
    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    /// @notice Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contreact
    function freezeOperator(address operator) external;

    /// @notice Permissioned function that causes the ServiceManager to revoke its ability to slash the operator on EigenLayer, through a call to the Slasher contreact
    function revokeSlashingAbility(address operator, uint32 unbondedAfter) external;

    /// @notice Collateral token used for placing collateral on challenges & payment commits
    function collateralToken() external view returns (IERC20);

    /// @notice The Delegation contract of EigenLayer.
    function eigenLayrDelegation() external view returns (IEigenLayrDelegation);

    /**
     * @notice Verifies that a task for this middleware exists which was created *at or before* `initTimestamp` *AND* that expires *strictly prior to* the
     * specified `unlockTime`.
     * @dev Function reverts if the verification fails.
     */
    function stakeWithdrawalVerification(bytes calldata data, uint256 initTimestamp, uint256 unlockTime)
        external
        view;

    /// @notice Returns the `latestTime` until which operators must serve.
    function latestTime() external view returns (uint32);
}