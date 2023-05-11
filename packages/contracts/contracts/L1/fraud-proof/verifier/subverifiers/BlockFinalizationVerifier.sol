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

contract BlockFinalizationVerifier is IVerifier {
    function verifyOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        external
        pure
        override
        returns (bytes32)
    {
        return OneStepProof.hashBlockStateProof(executeOneStepProof(ctx, currStateHash, encoded));
    }

    function executeOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        public
        pure
        returns (OneStepProof.BlockStateProof memory endState)
    {
        uint64 offset = 0;
        // Decode inter-state proof
        OneStepProof.InterStateProof memory stateProof;
        (offset, stateProof) = OneStepProof.decodeInterStateProof(encoded, offset);

        // Calculate the state hash from the submitted proof
        bytes32 stateHashFromProof;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, encoded.offset, encoded.length)
            stateHashFromProof := keccak256(ptr, offset)
        }
        // Ensure the state proof is valid
        require(stateHashFromProof == currStateHash, "Bad State Proof");

        // Obtain the parent block hash from the block hash tree
        bytes32 parentBlockHash;
        (offset, parentBlockHash) = VerifierHelper.decodeAndVerifyBlockHashProof(
            offset, encoded, stateProof.blockNumber - 1, stateProof.blockHashRoot
        );
        // Update the current block hash to the block hash tree
        bytes32 currentBlockHash = VerifierHelper.hashBlockHeader(ctx, parentBlockHash, stateProof);

        // Create the finalized block state using last inter-state proof
        (offset, endState.blockHashRoot) = VerifierHelper.decodeAndInsertBlockHashProof(
            offset, encoded, stateProof.blockNumber, stateProof.blockHashRoot, currentBlockHash
        );
        endState.blockNumber = stateProof.blockNumber;
        endState.globalStateRoot = stateProof.globalStateRoot;
        endState.cumulativeGasUsed = stateProof.cumulativeGasUsed;
        return endState;
    }
}
