// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./IDelegationShare.sol";

/**
 * @title Abstract interface for a contract that helps structure the delegation relationship.
 * @notice The gas budget provided to this contract in calls from contracts is limited.
 */
//TODO: discuss if we can structure the inputs of these functions better
interface IDelegationCallback {
    function payForService(IERC20 token, uint256 amount) external payable;

    function onDelegationReceived(
        address delegator,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory investorShares
    ) external;

    function onDelegationWithdrawn(
        address delegator,
        address operator,
        IDelegationShare[] memory delegationShares,
        uint256[] memory investorShares
    ) external;
}
