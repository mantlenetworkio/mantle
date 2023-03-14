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
import "../IRollup.sol";

contract Challenge is IChallenge {
    struct BisectedStore {
        bytes32 startState;
        bytes32 midState;
        bytes32 endState;
        uint256 blockNum;
        uint256 blockTime;
        uint256 challengedSegmentStart;
        uint256 challengedSegmentLength;
    }

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
    uint256 public lastMoveBlockTime;
    uint256 public defenderTimeLeft;
    uint256 public challengerTimeLeft;

    Turn public turn;
    // See `ChallengeLib.computeBisectionHash` for the format of this commitment.
    bytes32 public bisectionHash;
    bytes32[2] public prevBisection;

    // Initial state used to initialize bisectionHash (write-once).
    bytes32 private startStateHash;
    bytes32 private endStateHash;

    address public winner;

    bool public rollback;
    uint256 public startInboxSize;

    BisectedStore public currentBisected;

    /**
     * @notice Pre-condition: `msg.sender` is correct and still has time remaining.
     * Post-condition: `turn` changes and `lastMoveBlock` set to current `block.number`.
     */
    modifier onlyOnTurn() {
        require(msg.sender == currentResponder(), BIS_SENDER);
        require(block.timestamp - lastMoveBlockTime <= currentResponderTimeLeft(), BIS_DEADLINE);

        _;

        if (turn == Turn.Challenger) {
            challengerTimeLeft = challengerTimeLeft - (block.timestamp- lastMoveBlockTime);
            turn = Turn.Defender;
        } else if (turn == Turn.Defender) {
            defenderTimeLeft = defenderTimeLeft - (block.timestamp - lastMoveBlockTime);
            turn = Turn.Challenger;
        }
        lastMoveBlockTime = block.timestamp;
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
        uint256 _startInboxSize,
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
        lastMoveBlockTime = block.timestamp;
        // TODO(ujval): initialize timeout
        defenderTimeLeft = 150;
        challengerTimeLeft = 150;
        prevBisection[0] = _startStateHash;
        prevBisection[1] = _endStateHash;

        startInboxSize = _startInboxSize;
    }

    function initializeChallengeLength(bytes32 checkStateHash, uint256 _numSteps) external override onlyOnTurn {
        require(bisectionHash == 0, CHAL_INIT_STATE);
        require(_numSteps > 0, "INVALID_NUM_STEPS");
        bisectionHash = ChallengeLib.computeBisectionHash(0, _numSteps);
        // TODO: consider emitting a different event?
        currentBisected = BisectedStore(startStateHash, checkStateHash, endStateHash, block.number, block.timestamp, 0, _numSteps);
        emit Bisected(startStateHash, checkStateHash, endStateHash, block.number, block.timestamp, 0, _numSteps);
    }

    function bisectExecution(
        bytes32[3] calldata bisection,
        uint256 challengedSegmentIndex,
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external override onlyOnTurn postInitialization {
        // Verify provided prev bisection.
        bytes32 prevHash = ChallengeLib.computeBisectionHash(prevChallengedSegmentStart, prevChallengedSegmentLength);
        require(prevHash == bisectionHash, BIS_PREV);

        // Require agreed upon start state hash and disagreed upon end state hash.
        require(bisection[0] == prevBisection[0] || bisection[2] == prevBisection[1], "INVALID_START_OR_END");

        // Compute segment start/length.
        require(challengedSegmentLength > 0, "TOO_SHORT");

        // Compute new challenge state.
        prevBisection[0] = bisection[0];
        prevBisection[1] = bisection[2];
        bisectionHash = ChallengeLib.computeBisectionHash(challengedSegmentStart, challengedSegmentLength);
        currentBisected = BisectedStore(bisection[0], bisection[1], bisection[2], block.number, block.timestamp, challengedSegmentStart, challengedSegmentLength);
        emit Bisected(bisection[0], bisection[1], bisection[2], block.number, block.timestamp, challengedSegmentStart, challengedSegmentLength);
    }

    function verifyOneStepProof(
        VerificationContext.Context calldata ctx,
        uint8 verifyType,
        bytes calldata proof,
        uint256 prevChallengedSegmentStart,
        uint256 prevChallengedSegmentLength
    ) external override onlyOnTurn {
         // Verify provided prev bisection.
         bytes32 prevHash =
            ChallengeLib.computeBisectionHash(prevChallengedSegmentStart, prevChallengedSegmentLength);
         require(prevHash == bisectionHash, BIS_PREV);
         // require(challengedStepIndex > 0 && challengedStepIndex < prevBisection.length, "INVALID_INDEX");
         // Require that this is the last round.
         require(prevChallengedSegmentLength / MAX_BISECTION_DEGREE <= 1, "BISECTION_INCOMPLETE");

         // verify OSP
         // IVerificationContext ctx = <get ctx from sequenced txs>;
         bytes32 nextStateHash = verifier.verifyOneStepProof(
             ctx,
             verifyType,
             prevBisection[1],
             proof
         );
         if (nextStateHash == prevBisection[1]) {
             // osp verified, current win
             _currentWin(CompletionReason.OSP_VERIFIED);
         } else {
             _currentLose(CompletionReason.OSP_VERIFIED);
         }
    }

    function setRollback() public {
        if (rollback) {
            revert("ALREADY_SET_ROLLBACK");
        }
        rollback = true;
    }

    function timeout() external override {
        require(block.timestamp - lastMoveBlockTime > currentResponderTimeLeft(), TIMEOUT_DEADLINE);
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
            winner = defender;
            _asserterWin(reason);
        } else {
            winner = challenger;
            _challengerWin(reason);
        }
    }

    function _currentLose(CompletionReason reason) private {
        if (turn == Turn.Defender) {
            winner = challenger;
            _challengerWin(reason);
        } else {
            winner = defender;
            _asserterWin(reason);
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
