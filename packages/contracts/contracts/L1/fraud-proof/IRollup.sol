// SPDX-License-Identifier: Apache-2.0

/*
 * Modifications Copyright 2022, Specular contributors
 *
 * This file was changed in accordance to Apache License, Version 2.0.
 *
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.0;

import "./AssertionMap.sol";

interface IRollup {
    event AssertionCreated(
        uint256 assertionID, address asserterAddr, bytes32 vmHash, uint256 inboxSize, uint256 l2GasUsed
    );

    event AssertionChallenged(uint256 assertionID, address challengeAddr);

    event AssertionConfirmed(uint256 assertionID);

    event AssertionRejected(uint256 assertionID);

    event AdvanceStake(uint256 assertionID);

    function assertions() external view returns (AssertionMap);

    /**
     * @param addr User address.
     * @return True if address is staked, else False.
     */
    function isStaked(address addr) external view returns (bool);

    /**
     * @return The current required stake amount.
     */
    function currentRequiredStake() external view returns (uint256);

    /**
     * @return confirmedInboxSize size of inbox confirmed
     */
    function confirmedInboxSize() external view returns (uint256);

    /**
     * @notice Deposits stake on staker's current assertion (or the last confirmed assertion if not currently staked).
     * @notice currently use Ether to stake; stakeAmount Token amount to deposit. Must be > than defined threshold if this is a new stake.
     */
    // function stake(uint256 stakeAmount) external payable;
    function stake() external payable;

    /**
     * @notice Withdraws stakeAmount from staker's stake by if assertion it is staked on is confirmed.
     * @param stakeAmount Token amount to withdraw. Must be <= sender's current stake minus the current required stake.
     */
    function unstake(uint256 stakeAmount) external;

    /**
     * @notice Removes stakerAddress from the set of stakers and withdraws the full stake amount to stakerAddress.
     * This can be called by anyone since it is currently necessary to keep the chain progressing.
     * @param stakerAddress Address of staker for which to unstake.
     */
    function removeStake(address stakerAddress) external;

    /**
     * @notice Advances msg.sender's existing sake to assertionID.
     * @param assertionID ID of assertion to advance stake to. Currently this must be a child of the current assertion.
     * TODO: generalize to arbitrary descendants.
     */
    function advanceStake(uint256 assertionID) external;

    /**
     *
     * @notice create assertion with scc state batch
     *
     * @param vmHash New VM hash.
     * @param inboxSize Size of inbox corresponding to assertion (number of transactions).
     * @param l2GasUsed Total L2 gas used as of the end of this assertion's last transaction.
     * @param _batch Batch of state roots.
     * @param _shouldStartAtElement Index of the element at which this batch should start.
     * @param _signature tss group signature of state batches.
     */
    function createAssertionWithStateBatch(
        bytes32 vmHash,
        uint256 inboxSize,
        uint256 l2GasUsed,
        bytes32[] calldata _batch,
        uint256 _shouldStartAtElement,
        bytes calldata _signature
    ) external;

    /**
     * @notice Creates a new DA representing the rollup state after executing a block of transactions (sequenced in SequencerInbox).
     * Block is represented by all transactions in range [prevInboxSize, inboxSize]. The latest staked DA of the sender
     * is considered to be the predecessor. Moves sender stake onto the new DA.
     *
     * The new DA stores the hash of the parameters: H(l2GasUsed || vmHash)
     *
     * @param vmHash New VM hash.
     * @param inboxSize Size of inbox corresponding to assertion (number of transactions).
     * @param l2GasUsed Total L2 gas used as of the end of this assertion's last transaction.
     */
    function createAssertion(
        bytes32 vmHash,
        uint256 inboxSize,
        uint256 l2GasUsed
    ) external;

    /**
     * @notice Initiates a dispute between a defender and challenger on an unconfirmed DA.
     * @param players Defender (first) and challenger (second) addresses. Must be staked on DAs on different branches.
     * @param assertionIDs Assertion IDs of the players engaged in the challenge. The first ID should be the earlier-created and is the one being challenged.
     * @return Newly created challenge contract address.
     */
    function challengeAssertion(address[2] calldata players, uint256[2] calldata assertionIDs)
        external
        returns (address);

    /**
     * @notice Confirms first unresolved assertion. Assertion is confirmed if and only if:
     * (1) there is at least one staker, and
     * (2) challenge period has passed, and
     * (3) predecessor has been confirmed, and
     * (4) all stakers are staked on the assertion.
     */
    function confirmFirstUnresolvedAssertion() external;

    /**
     * @notice Rejects first unresolved assertion. Assertion is rejected if and only if:
     * (1) all of the following are true:
     * (a) challenge period has passed, and
     * (b) at least one staker exists, and
     * (c) no staker remains staked on the assertion (all have been destroyed).
     * OR
     * (2) predecessor has been rejected
     * @param stakerAddress Address of a staker staked on a different branch to the first unresolved assertion.
     * If the first unresolved assertion's parent is confirmed, this parameter is used to establish that a staker exists
     * on a different branch of the assertion chain. This parameter is ignored when the parent of the first unresolved
     * assertion is not the last confirmed assertion.
     */
    function rejectFirstUnresolvedAssertion(address stakerAddress) external;

    /**
     * @notice Completes ongoing challenge. Callback, called by a challenge contract.
     * @param winner Address of winning staker.
     * @param loser Address of losing staker.
     */
    function completeChallenge(address winner, address loser) external;
}
