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

library ChallengeLib {
    /**
     * @notice Computes the initial bisection hash.
     * @param numSteps Number of steps from the end of `startState` to the end of `endState`.
     */
    function initialBisectionHash(uint256 numSteps)
        internal
        pure
        returns (bytes32)
    {
        return ChallengeLib.computeBisectionHash(0, numSteps);
    }

    /**
     * @notice Computes H(bisection || segmentStart || segmentLength)
     * @param challengedSegmentStart The number of steps preceding `bisection[1]`, relative to the assertion being challenged.
     * @param challengedSegmentLength Length of bisected segment (in steps), from the start of bisection[1] to the end of bisection[-1].
     */
    function computeBisectionHash(
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(challengedSegmentStart, challengedSegmentLength));
    }

    /**
     * @notice Returns length of first segment in a bisection.
     */
    function firstSegmentLength(uint256 length, uint256 bisectionDegree) internal pure returns (uint256) {
        return length / bisectionDegree + length % bisectionDegree;
    }

    /**
     * @notice Returns length of a segment (after first) in a bisection.
     */
    function otherSegmentLength(uint256 length, uint256 bisectionDegree) internal pure returns (uint256) {
        return length / bisectionDegree;
    }
}
