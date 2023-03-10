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

import "../verifier/IVerifierEntry.sol";

/**
 * @notice Protocol execution:
 * `initialize` (challenger, via Rollup) ->
 * `initializeChallengeLength` (defender) ->
 * `bisectExecution` (challenger, defender -- alternating) ->
 * `verifyOneStepProof`
 */
interface IChallenge {
    enum CompletionReason {
        OSP_VERIFIED, // OSP verified by winner.
        TIMEOUT // Loser timed out before completing their round.
    }

    event ChallengeCompleted(address winner, address loser, CompletionReason reason);

    event Bisected(bytes32 startState, bytes32 midState, bytes32 endState, uint256 blockNum, uint256 blockTime, uint256 challengedSegmentStart, uint256 challengedSegmentLength);

    /**
     * @notice Initializes contract.
     * @param _defender Defending party.
     * @param _challenger Challenging party. Challenger starts.
     * @param _verifier Address of the verifier contract.
     * @param _resultReceiver Address of contract that will receive the outcome (via callback `completeChallenge`).
     * @param _startStateHash Bisection root being challenged.
     * @param _endStateHash Bisection root being challenged.
     */
    function initialize(
        address _defender,
        address _challenger,
        IVerifierEntry _verifier,
        address _resultReceiver,
        uint256 _startInboxSize,
        bytes32 _startStateHash,
        bytes32 _endStateHash
    ) external;

    /**
     * @notice Initializes the length of the challenge. Must be called by defender before bisection rounds begin.
     * @param _numSteps Number of steps executed from the start of the assertion to its end.
     * If this parameter is incorrect, the defender will be slashed (assuming successful execution of the protocol by the challenger).
     */
    function initializeChallengeLength(bytes32 checkStateHash, uint256 _numSteps) external;

    /**
     * @notice Bisects a segment. The challenged segment is defined by: {`challengedSegmentStart`, `challengedSegmentLength`, `bisection[0]`, `oldEndHash`}
     * @param bisection Bisection of challenged segment. Each element is a state hash (see `ChallengeLib.stateHash`).
     * The first element is the last agreed upon state hash. Must be of length MAX_BISECTION_LENGTH for all rounds except the last.
     * In the last round, the bisection segments must be single steps.
     * @param challengedSegmentIndex Index into `prevBisection`. Must be greater than 0 (since the first is agreed upon).
     * @param challengedSegmentStart Offset of the segment challenged in the preceding round (in steps).
     * @param challengedSegmentLength Length of the segment challenged in the preceding round (in steps).
     * @param prevChallengedSegmentStart Offset of the segment challenged in the preceding round (in steps).
     * Note: this is relative to the assertion being challenged (i.e. always between 0 and the initial `numSteps`).
     * @param prevChallengedSegmentLength Length of the segment challenged in the preceding round (in steps).
     */
    function bisectExecution(
        bytes32[3] calldata bisection,
        uint256 challengedSegmentIndex,
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external;

    /**
     * @notice Verifies one step proof and completes challenge protocol.
     * @param ctx execution context.
     * @param verifyType Index into `prevBisection`. Must be greater than 0 (since the first is agreed upon).
     * @param proof one step proof.
     * @param prevChallengedSegmentStart Offset of the segment challenged in the preceding round (in steps).
     * Note: this is relative to the assertion being challenged (i.e. always between 0 and the initial `numSteps`).
     * @param prevChallengedSegmentLength Length of the segment challenged in the preceding round (in steps).
     */
    function verifyOneStepProof(
        VerificationContext.Context calldata ctx,
        uint8 verifyType,
        bytes calldata proof,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external;

    function setRollback() external;

    /**
     * @notice Triggers completion of challenge protocol if a responder timed out.
     */
    function timeout() external;

    function currentResponder() external view returns (address);

    function currentResponderTimeLeft() external view returns (uint256);
}
