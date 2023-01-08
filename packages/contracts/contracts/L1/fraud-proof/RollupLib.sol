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

import "./challenge/ChallengeLib.sol";

// TODO: move into ChallengeLib.
library RollupLib {
    struct ExecutionState {
        uint256 l2GasUsed;
        bytes32 vmHash;
    }

    /**
     * @notice Computes the hash of `execState`.
     */
    function stateHash(ExecutionState memory execState) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(execState.l2GasUsed, execState.vmHash));
    }
}
