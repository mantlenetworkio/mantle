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

contract CallOpVerifier is IVerifier {
    using BytesLib for bytes;
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
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

        // Simulate EVM one-step execution based on the opcode
        uint8 opCode = stateProof.opCode;
        if (opCode == 0xf0) {
            // CREATE
            verifyOpCREATE(offset, stateProof, encoded);
        } else if (opCode == 0xf1) {
            // CALL
            verifyOpCALL(offset, stateProof, encoded);
        } else if (opCode == 0xf3) {
            // RETURN
            verifyOpRETURN(ctx, offset, stateProof, encoded);
        } else if (opCode == 0xf2) {
            // CALLCODE
            verifyOpCALLCODE(offset, stateProof, encoded);
        } else if (opCode == 0xf4) {
            // DELEGATECALL
            verifyOpDELEGATECALL(offset, stateProof, encoded);
        } else if (opCode == 0xf5) {
            // CREATE2
            verifyOpCREATE2(offset, stateProof, encoded);
        } else if (opCode == 0xfa) {
            // STATICCALL
            verifyOpSTATICCALL(offset, stateProof, encoded);
        } else if (opCode == 0xfd) {
            // REVERT
            verifyOpREVERT(ctx, offset, stateProof, encoded);
        } else if (opCode == 0xff) {
            // SELFDESTRUCT
            verifyOpSELFDESTRUCT(ctx, offset, stateProof, encoded);
        } else {
            revert("Unreachable");
        }

        // Return the state hash after one-step execution
        return stateProof;
    }

    function verifyOpCREATE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpCALL(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpRETURN(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {}

    function verifyOpCALLCODE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpDELEGATECALL(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpCREATE2(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpSTATICCALL(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}

    function verifyOpREVERT(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {}

    function verifyOpSELFDESTRUCT(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {}
}
