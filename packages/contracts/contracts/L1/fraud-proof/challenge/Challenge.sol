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

import "hardhat/console.sol";

import "./IChallenge.sol";
import "./ChallengeLib.sol";
import "../verifier/IVerifierEntry.sol";
import "../IRollup.sol";

contract Challenge is IChallenge {
    enum Turn {
        NoChallenge,
        Challenger,
        Defender
    }

    // Error codes

    // Can only initialize once
    string private constant CHAL_INIT_STATE = "CHAL_INIT_STATE";
    // deadline expired
    string private constant BIS_DEADLINE = "BIS_DEADLINE";
    // Only original asserter can continue bisect
    string private constant BIS_SENDER = "BIS_SENDER";
    // Incorrect previous state
    string private constant BIS_PREV = "BIS_PREV";
    // Can't timeout before deadline
    string private constant TIMEOUT_DEADLINE = "TIMEOUT_DEADLINE";

    bytes32 private constant UNREACHABLE_ASSERTION = bytes32(uint256(0));

    uint256 private constant MAX_BISECTION_DEGREE = 2;

    // Other contracts
    address internal resultReceiver;
    IVerifierEntry internal verifier;

    // Challenge state
    address public defender;
    address public challenger;
    uint256 public lastMoveBlock;
    uint256 public defenderTimeLeft;
    uint256 public challengerTimeLeft;

    Turn public turn;
    // See `ChallengeLib.computeBisectionHash` for the format of this commitment.
    bytes32 public bisectionHash;
    // Initial state used to initialize bisectionHash (write-once).
    bytes32 private startStateHash;
    bytes32 private endStateHash;

    /**
     * @notice Pre-condition: `msg.sender` is correct and still has time remaining.
     * Post-condition: `turn` changes and `lastMoveBlock` set to current `block.number`.
     */
    modifier onlyOnTurn() {
        require(msg.sender == currentResponder(), BIS_SENDER);
        require(block.number - lastMoveBlock <= currentResponderTimeLeft(), BIS_DEADLINE);

        _;

        if (turn == Turn.Challenger) {
            challengerTimeLeft = challengerTimeLeft - (block.number - lastMoveBlock);
            turn = Turn.Defender;
        } else if (turn == Turn.Defender) {
            defenderTimeLeft = defenderTimeLeft - (block.number - lastMoveBlock);
            turn = Turn.Challenger;
        }
        lastMoveBlock = block.number;
    }

    /**
     * @notice Ensures challenge has been initialized.
     */
    modifier postInitialization() {
        require(bisectionHash != 0, "NOT_INITIALIZED");
        _;
    }

    function initialize(
        address _defender,
        address _challenger,
        IVerifierEntry _verifier,
        address _resultReceiver,
        bytes32 _startStateHash,
        bytes32 _endStateHash
    ) external override {
        require(turn == Turn.NoChallenge, CHAL_INIT_STATE);
        require(_defender != address(0) && _challenger != address(0) && _resultReceiver != address(0), "ZERO_ADDRESS");
        defender = _defender;
        challenger = _challenger;
        verifier = _verifier;
        resultReceiver = _resultReceiver;
        startStateHash = _startStateHash;
        endStateHash = _endStateHash;

        turn = Turn.Defender;
        lastMoveBlock = block.number;
        // TODO(ujval): initialize timeout
        defenderTimeLeft = 10;
        challengerTimeLeft = 10;
    }

    function initializeChallengeLength(uint256 _numSteps) external override onlyOnTurn {
        require(bisectionHash == 0, CHAL_INIT_STATE);
        require(_numSteps > 0, "INVALID_NUM_STEPS");
        bisectionHash = ChallengeLib.initialBisectionHash(startStateHash, endStateHash, _numSteps);
        // TODO: consider emitting a different event?
        emit Bisected(bisectionHash, 0, _numSteps);
    }

    function bisectExecution(
        bytes32[] calldata bisection,
        uint256 challengedSegmentIndex,
        bytes32[] calldata prevBisection,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external override onlyOnTurn postInitialization {
        // Verify provided prev bisection.
        bytes32 prevHash =
            ChallengeLib.computeBisectionHash(prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength);
        require(prevHash == bisectionHash, BIS_PREV);
        console.log("log out bisect hash...");
        console.logBytes32(prevHash);
        console.logBytes32(bisectionHash);

        require(challengedSegmentIndex > 0 && challengedSegmentIndex < prevBisection.length, "INVALID_INDEX");
        // Require agreed upon start state hash and disagreed upon end state hash.
        require(bisection[0] == prevBisection[challengedSegmentIndex - 1], "INVALID_START");
        require(bisection[bisection.length - 1] != prevBisection[challengedSegmentIndex], "INVALID_END");

        // Compute segment start/length.
        uint256 challengedSegmentStart = prevChallengedSegmentStart;
        uint256 challengedSegmentLength = prevChallengedSegmentLength;
        if (prevBisection.length > 2) {
            // prevBisection.length == 2 means first round
            uint256 firstSegmentLength =
                ChallengeLib.firstSegmentLength(prevChallengedSegmentLength, MAX_BISECTION_DEGREE);
            uint256 otherSegmentLength =
                ChallengeLib.otherSegmentLength(prevChallengedSegmentLength, MAX_BISECTION_DEGREE);
            challengedSegmentLength = challengedSegmentIndex == 1 ? firstSegmentLength : otherSegmentLength;

            if (challengedSegmentIndex > 1) {
                challengedSegmentStart += firstSegmentLength + (otherSegmentLength * (challengedSegmentIndex - 2));
            }
        }
        require(challengedSegmentLength > 1, "TOO_SHORT");

        // Require that bisection has the correct length. This is only ever less than BISECTION_DEGREE at the last bisection.
        uint256 target = challengedSegmentLength < MAX_BISECTION_DEGREE ? challengedSegmentLength : MAX_BISECTION_DEGREE;
        require(bisection.length == target + 1, "CUT_COUNT");

        // Compute new challenge state.
        bisectionHash = ChallengeLib.computeBisectionHash(bisection, challengedSegmentStart, challengedSegmentLength);
        emit Bisected(bisectionHash, challengedSegmentStart, challengedSegmentLength);
    }

    function verifyOneStepProof(
        bytes calldata proof,
        uint256 challengedStepIndex,
        bytes32[] calldata prevBisection,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external override onlyOnTurn {
        // Verify provided prev bisection.
        bytes32 prevHash =
            ChallengeLib.computeBisectionHash(prevBisection, prevChallengedSegmentStart, prevChallengedSegmentLength);
        require(prevHash == bisectionHash, BIS_PREV);
        require(challengedStepIndex > 0 && challengedStepIndex < prevBisection.length, "INVALID_INDEX");
        // Require that this is the last round.
        require(prevChallengedSegmentLength / MAX_BISECTION_DEGREE <= 1, "BISECTION_INCOMPLETE");

        // TODO: verify OSP
        // IVerificationContext ctx = <get ctx from sequenced txs>;
        // bytes32 nextStateHash = verifier.verifyOneStepProof(
        //     ctx,
        //     prevBisection[challengedStepIndex - 1],
        //     proof
        // );
        // if (nextStateHash == prevBisection[challengedStepIndex]) {
        //     // osp verified, current win
        // } else {
        //     // current lose?
        // }

        _currentWin(CompletionReason.OSP_VERIFIED);
    }

    function timeout() external override {
        require(block.number - lastMoveBlock > currentResponderTimeLeft(), TIMEOUT_DEADLINE);
        if (turn == Turn.Defender) {
            _challengerWin(CompletionReason.TIMEOUT);
        } else {
            _asserterWin(CompletionReason.TIMEOUT);
        }
    }

    function currentResponder() public view override returns (address) {
        if (turn == Turn.Defender) {
            return defender;
        } else if (turn == Turn.Challenger) {
            return challenger;
        } else {
            revert("NO_TURN");
        }
    }

    function currentResponderTimeLeft() public view override returns (uint256) {
        if (turn == Turn.Defender) {
            return defenderTimeLeft;
        } else if (turn == Turn.Challenger) {
            return challengerTimeLeft;
        } else {
            revert("NO_TURN");
        }
    }

    function _currentWin(CompletionReason reason) private {
        if (turn == Turn.Defender) {
            _asserterWin(reason);
        } else {
            _challengerWin(reason);
        }
    }

    function _asserterWin(CompletionReason reason) private {
        emit ChallengeCompleted(defender, challenger, reason);
        IRollup(resultReceiver).completeChallenge(defender, challenger); // safeSelfDestruct(msg.sender);
    }

    function _challengerWin(CompletionReason reason) private {
        emit ChallengeCompleted(challenger, defender, reason);
        IRollup(resultReceiver).completeChallenge(challenger, defender); // safeSelfDestruct(msg.sender);
    }
}
