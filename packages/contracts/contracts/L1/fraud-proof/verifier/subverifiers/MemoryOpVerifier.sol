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
import "../../libraries/MerkleLib.sol";
import "../libraries/VerificationContext.sol";
import "../libraries/OneStepProof.sol";
import "../libraries/Params.sol";
import "../libraries/GasTable.sol";
import "../libraries/VerifierHelper.sol";
import "../IVerifier.sol";

contract MemoryOpVerifier is IVerifier {
    using BytesLib for bytes;
    using VerificationContext for VerificationContext.Context;
    using OneStepProof for OneStepProof.CodeProof;
    using OneStepProof for OneStepProof.LogProof;

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
        returns (OneStepProof.StateProof memory)
    {
        uint64 offset = 0;
        // Decode state proof
        OneStepProof.StateProof memory stateProof;
        (offset, stateProof) = OneStepProof.decodeStateProof(ctx, encoded, offset);
        // Calculate the state hash from the submitted proof
        bytes32 stateHashFromProof;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, encoded.offset, encoded.length)
            stateHashFromProof := keccak256(ptr, offset)
        }
        // Ensure the state proof is valid
        require(stateHashFromProof == currStateHash, "Bad State Proof");

        // Decode the code proof
        OneStepProof.CodeProof memory codeProof;
        (offset, codeProof) = OneStepProof.decodeCodeProof(encoded, offset);

        // Simulate EVM one-step execution based on the opcode
        uint8 opCode = stateProof.opCode;
        if (opCode == 0x20) {
            // KECCAK256
            verifyOpKECCAK256(offset, stateProof, encoded);
        } else if (opCode == 0x35) {
            // CALLDATALOAD
            verifyOpCALLDATALOAD(offset, stateProof, encoded);
        } else if (opCode == 0x37) {
            // CALLDATACOPY
            verifyOpCALLDATACOPY(offset, stateProof, encoded);
        } else if (opCode == 0x39) {
            // CODECOPY
            verifyOpCODECOPY(offset, stateProof, codeProof, encoded);
        } else if (opCode == 0x3e) {
            // RETURNDATACOPY
            verifyOpRETURNDATACOPY(offset, stateProof, encoded);
        } else if (opCode == 0x51) {
            // MLOAD
            verifyOpMLOAD(offset, stateProof, encoded);
        } else if (opCode == 0x52) {
            // MSTORE
            verifyOpMSTORE(offset, stateProof, encoded);
        } else if (opCode == 0x53) {
            // MSTORE8
            verifyOpMSTORE8(offset, stateProof, encoded);
        } else if (opCode >= 0xa0 && opCode <= 0xa4) {
            // LOG
            verifyLogOpCode(offset, stateProof, encoded);
        } else {
            revert("Unreachable");
        }

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

    function verifyOpKECCAK256(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 2) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2, offset, encoded);

        uint64 memOffset = uint64(stackProof.pops[0]);
        uint64 length = uint64(stackProof.pops[1]);
        uint64 cost = GasTable.gasKeccak(stateProof.memSize, memOffset, length);
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        bytes memory readContent;
        (offset, readContent) = MemoryLib.decodeAndVerifyMemoryReadProof(stateProof, encoded, offset, memOffset, length);

        bytes32 result = keccak256(abi.encodePacked(readContent));
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, result));
        stateProof.pc += 1;
    }

    function verifyOpCALLDATALOAD(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_VERYLOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        uint64 memOffset = uint64(stackProof.pops[0]);
        bytes memory readContent;
        (offset, readContent) = MemoryLib.decodeAndVerifyMemoryLikeReadProofNoAppend(
            stateProof.inputDataRoot, stateProof.inputDataSize, encoded, offset, memOffset, 32
        );

        bytes32 result = readContent.toBytes32(0);
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, result));
        stateProof.pc += 1;
    }

    function verifyOpCALLDATACOPY(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 3) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 3, offset, encoded);
        uint64 memOffset = uint64(stackProof.pops[0]);
        uint64 inputOffset = uint64(stackProof.pops[1]);
        uint64 length = uint64(stackProof.pops[2]);

        uint64 cost = GasTable.gasCopy(stateProof.memSize, memOffset, length);
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        {
            bytes memory readContent;
            (offset, readContent) = MemoryLib.decodeAndVerifyMemoryLikeReadProofNoAppend(
                stateProof.inputDataRoot, stateProof.inputDataSize, encoded, offset, inputOffset, length
            );

            bytes memory writeContent;
            (offset, writeContent) =
                MemoryLib.decodeAndVerifyMemoryWriteProof(stateProof, encoded, offset, memOffset, length);

            require(readContent.equal(writeContent), "Inconsistent Copy");
        }

        stateProof.pc += 1;
        stateProof.stackHash = stackProof.stackHashAfterPops;
    }

    function verifyOpCODECOPY(
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.CodeProof memory codeProof,
        bytes calldata encoded
    ) internal pure {
        if (stateProof.stackSize < 3) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 3, offset, encoded);
        uint64 memOffset = uint64(stackProof.pops[0]);
        uint64 codeOffset = uint64(stackProof.pops[1]);
        uint64 length = uint64(stackProof.pops[2]);

        uint64 cost = GasTable.gasCopy(stateProof.memSize, memOffset, length);
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        {
            bytes memory readContent = codeProof.getCodeSlice(encoded, codeOffset, length);

            bytes memory writeContent;
            (offset, writeContent) =
                MemoryLib.decodeAndVerifyMemoryWriteProof(stateProof, encoded, offset, memOffset, length);

            require(readContent.equal(writeContent), "Inconsistent Copy");
        }

        stateProof.pc += 1;
        stateProof.stackHash = stackProof.stackHashAfterPops;
    }

    function verifyOpRETURNDATACOPY(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 3) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 3, offset, encoded);
        uint64 memOffset = uint64(stackProof.pops[0]);
        uint64 returnOffset = uint64(stackProof.pops[1]);
        uint64 length = uint64(stackProof.pops[2]);

        uint64 cost = GasTable.gasCopy(stateProof.memSize, memOffset, length);
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        {
            bytes memory readContent;
            (offset, readContent) = MemoryLib.decodeAndVerifyMemoryLikeReadProofNoAppend(
                stateProof.returnDataRoot, stateProof.returnDataSize, encoded, offset, returnOffset, length
            );

            bytes memory writeContent;
            (offset, writeContent) =
                MemoryLib.decodeAndVerifyMemoryWriteProof(stateProof, encoded, offset, memOffset, length);

            require(readContent.equal(writeContent), "Inconsistent Copy");
        }

        stateProof.pc += 1;
        stateProof.stackHash = stackProof.stackHashAfterPops;
    }

    function verifyOpMLOAD(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_VERYLOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        uint64 memOffset = uint64(stackProof.pops[0]);
        bytes memory readContent;
        (offset, readContent) = MemoryLib.decodeAndVerifyMemoryReadProof(stateProof, encoded, offset, memOffset, 32);

        bytes32 result = readContent.toBytes32(0);
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, result));
        stateProof.pc += 1;
    }

    function verifyOpMSTORE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 2) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2, offset, encoded);

        uint64 cost = Params.G_VERYLOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        uint64 memOffset = uint64(stackProof.pops[0]);
        bytes32 value = bytes32(stackProof.pops[1]);
        bytes memory writeContent;
        if (stateProof.memSize == 0) {
            stateProof.memSize = (memOffset + 32 + 31) / 32 * 32;
            writeContent = abi.encodePacked(new bytes(memOffset), value);
            stateProof.memRoot = MemoryLib.getMemoryRoot(writeContent);
        } else {
            (offset, writeContent) =
                MemoryLib.decodeAndVerifyMemoryWriteProof(stateProof, encoded, offset, memOffset, 32);
            require(writeContent.toBytes32(0) == value, "Bad MemoryProof");
        }
        stateProof.stackHash = stackProof.stackHashAfterPops;
        stateProof.pc += 1;
    }

    function verifyOpMSTORE8(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 2) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2, offset, encoded);

        uint64 cost = Params.G_VERYLOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        uint64 memOffset = uint64(stackProof.pops[0]);
        uint8 value = uint8(stackProof.pops[1]);
        bytes memory writeContent;
        if (stateProof.memSize == 0) {
            stateProof.memSize = (memOffset + 1 + 31) / 32 * 32;
            writeContent = abi.encodePacked(new bytes(memOffset), value);
            stateProof.memRoot = MemoryLib.getMemoryRoot(writeContent);
        } else {
            (offset, writeContent) =
                MemoryLib.decodeAndVerifyMemoryWriteProof(stateProof, encoded, offset, memOffset, 1);
            require(writeContent.toUint8(0) == value, "Bad MemoryProof");
        }
        stateProof.stackHash = stackProof.stackHashAfterPops;
        stateProof.pc += 1;
    }

    function verifyLogOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        uint8 topicNum = stateProof.opCode - 0xa0;
        if (stateProof.stackSize < 2 + topicNum) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2 + topicNum, offset, encoded);

        uint64 memOffset = uint64(stackProof.pops[0]);
        uint64 length = uint64(stackProof.pops[1]);
        uint64 cost = GasTable.gasLog(stateProof.memSize, memOffset, length, topicNum);
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= cost;

        bytes memory readContent;
        (offset, readContent) = MemoryLib.decodeAndVerifyMemoryReadProof(stateProof, encoded, offset, memOffset, length);

        OneStepProof.LogProof memory logProof;
        (offset, logProof) = OneStepProof.decodeLogProof(encoded, offset);
        require(logProof.hashLogProof() == stateProof.logAcc, "Bad LogProof");

        stateProof.logAcc =
            VerifierHelper.updateAndHashLogProof(logProof, stateProof.contractAddress, stackProof.pops, readContent);
        stateProof.pc += 1;
        stateProof.stackHash = stackProof.stackHashAfterPops;
    }
}
