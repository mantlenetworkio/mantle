// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2022, Specular contributors
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

import "../../libraries/BytesLib.sol";
import "../../libraries/MPT.sol";
import "../IVerifier.sol";
import "../libraries/VerificationContext.sol";
import "../libraries/OneStepProof.sol";
import "../libraries/Params.sol";
import "../libraries/GasTable.sol";
import "../libraries/VerifierHelper.sol";
import "../libraries/EVMTypesLib.sol";
import "../libraries/BloomLib.sol";

contract BlockInitiationVerifier is IVerifier {
    function verifyOneStepProof(VerificationContext.Context memory, bytes32 currStateHash, bytes calldata encoded)
        external
        pure
        override
        returns (bytes32)
    {
        return OneStepProof.hashInterStateProof(executeOneStepProof(currStateHash, encoded));
    }

    function executeOneStepProof(bytes32 currStateHash, bytes calldata encoded)
        public
        pure
        returns (OneStepProof.InterStateProof memory endState)
    {
        uint64 offset = 0;
        // Decode block state proof
        OneStepProof.BlockStateProof memory blockStateProof;
        (offset, blockStateProof) = OneStepProof.decodeBlockStateProof(encoded, offset);

        // Calculate the state hash from the submitted proof
        bytes32 stateHashFromProof;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, encoded.offset, offset)
            stateHashFromProof := keccak256(ptr, offset)
        }
        // Ensure the block state proof is valid
        require(stateHashFromProof == currStateHash, "Bad State Proof");

        // Create the first inter-state in the block using block state proof
        endState.blockNumber = blockStateProof.blockNumber + 1;
        endState.transactionIdx = 0;
        endState.globalStateRoot = blockStateProof.globalStateRoot;
        endState.cumulativeGasUsed = blockStateProof.cumulativeGasUsed;
        endState.blockHashRoot = blockStateProof.blockHashRoot;
        endState.transactionTrieRoot = Params.EMPTY_TRIE_ROOT;
        endState.receiptTrieRoot = Params.EMPTY_TRIE_ROOT;
        endState.logsBloom = BloomLib.emptyBloom();
        return endState;
    }
}
