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
import "../IVerifier.sol";
import "../libraries/VerificationContext.sol";
import "../libraries/OneStepProof.sol";
import "../libraries/Params.sol";
import "../libraries/GasTable.sol";
import "../libraries/VerifierHelper.sol";

contract InvalidOpVerifier is IVerifier {
    using BytesLib for bytes;
    using VerificationContext for VerificationContext.Context;
    using OneStepProof for OneStepProof.CodeProof;

    function verifyOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        external
        pure
        override
        returns (bytes32)
    {
        return OneStepProof.hashStateProof(executeOneStepProof(ctx, currStateHash, encoded));
    }

    function executeOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        public
        pure
        returns (OneStepProof.StateProof memory endState)
    {
        uint64 offset = 0;
        // Decode state proof
        OneStepProof.StateProof memory stateProof;
        (offset, stateProof) = OneStepProof.decodeStateProof(ctx, encoded, offset);
        // Calculate the state hash from the submitted proof
        bytes32 stateHashFromProof;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, encoded.offset, offset)
            stateHashFromProof := keccak256(ptr, offset)
        }
        // Ensure the state proof is valid
        require(stateHashFromProof == currStateHash, "Bad State Proof");

        // Decode the code proof
        OneStepProof.CodeProof memory codeProof;
        (offset, codeProof) = OneStepProof.decodeCodeProof(encoded, offset);

        // Simulate EVM one-step execution based on the opcode
        uint8 opCode = stateProof.opCode;
        if (!isInvalidOp(opCode)) {
            revert("Unreachable");
        }

        VerifierHelper.verifyRevertByError(offset, stateProof, encoded);

        // Obtain the opcode at new pc
        if (stateProof.depth > 0) {
            if (codeProof.size > uint256(stateProof.pc)) {
                stateProof.opCode = codeProof.getOpCodeAt(encoded, stateProof.pc);
            } else {
                stateProof.opCode = 0x00;
            }
        }
        // Return the state hash after one-step execution
        return stateProof;
    }

    function isInvalidOp(uint8 opCode) internal pure returns (bool) {
        if (opCode == 0xfe) {
            // INVALID
            return true;
        }
        if (opCode >= 0x0c && opCode <= 0x0f) {
            return true;
        }
        if (opCode >= 0x1e && opCode <= 0x1f) {
            return true;
        }
        if (opCode >= 0x21 && opCode <= 0x2f) {
            return true;
        }
        if (opCode >= 0x47 && opCode <= 0x4f) {
            return true; // BASEFEE not supported in Istanbul protocol
        }
        if (opCode >= 0x5c && opCode <= 0x5f) {
            return true;
        }
        if (opCode >= 0xa5 && opCode <= 0xef) {
            return true;
        }
        if (opCode >= 0xf6 && opCode <= 0xf9 || opCode == 0xfb || opCode == 0xfc) {
            return true;
        }
        return false;
    }
}
