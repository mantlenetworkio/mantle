// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IPaymentManager.sol";
import "./IDataLayrServiceManager.sol";

/**
 * @title Minimal interface extension to `IPaymentManager`.
 * @author Layr Labs, Inc.
 * @notice Adds a single DataLayr-specific function to the base interface.
 */
interface IDataLayrPaymentManager is IPaymentManager {
    /**
     * @notice Used to perform the final step in a payment challenge, in which the 'trueAmount' is determined and the winner of the challenge is decided.
     * This function is called by a party after the other party has bisected the challenged payments to a difference of one, i.e., further bisection
     * is not possible. Once the payments can no longer be bisected, the function resolves the challenge by determining who is wrong.
     */
    function respondToPaymentChallengeFinal(
        address operator,
        uint256 stakeIndex,
        uint48 nonSignerIndex,
        bytes32[] memory nonSignerPubkeyHashes,
        TotalStakes calldata totalStakes,
        IDataLayrServiceManager.DataStoreSearchData calldata searchData
    ) external;
}
