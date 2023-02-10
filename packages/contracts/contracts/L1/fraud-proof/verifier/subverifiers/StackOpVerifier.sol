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

contract StackOpVerifier is IVerifier {
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
        if (opCode == 0x01) {
            // ADD
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x02) {
            // MUL
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x03) {
            // SUB
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x04) {
            // DIV
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x05) {
            // SDIV
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x06) {
            // MOD
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x07) {
            // SMOD
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x08) {
            // ADDMOD
            verifyTrinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x09) {
            // MULMOD
            verifyTrinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x0a) {
            // EXP
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x0b) {
            // SIGNEXTEND
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x10) {
            // LT
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x11) {
            // GT
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x12) {
            // SLT
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x13) {
            // SGT
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x14) {
            // EQ
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x15) {
            // ISZERO
            verifyUnaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x16) {
            // AND
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x17) {
            // OR
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x18) {
            // XOR
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x19) {
            // NOT
            verifyUnaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x1a) {
            // BYTE
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x1b) {
            // SHL
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x1c) {
            // SHR
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x1d) {
            // SAR
            verifyBinaryOpCode(offset, stateProof, encoded);
        } else if (opCode == 0x50) {
            // POP
            verifyOpPOP(offset, stateProof, encoded);
        } else if (opCode == 0x56) {
            // JUMP
            verifyOpJUMP(offset, stateProof, encoded);
        } else if (opCode == 0x57) {
            // JUMPI
            verifyOpJUMPI(offset, stateProof, encoded);
        } else if (opCode == 0x58) {
            // PC
            verifyOpPC(offset, stateProof, encoded);
        } else if (opCode == 0x59) {
            // MSIZE
            verifyOpMSIZE(offset, stateProof, encoded);
        } else if (opCode == 0x5a) {
            // GAS
            verifyOpJUMPDEST(offset, stateProof, encoded);
        } else if (opCode == 0x5b) {
            // JUMPDEST
            verifyOpJUMPDEST(offset, stateProof, encoded);
        } else if (opCode >= 0x60 && opCode <= 0x7f) {
            // PUSH
            verifyPushOpCode(offset, stateProof, codeProof, encoded);
        } else if (opCode >= 0x80 && opCode <= 0x8f) {
            // DUP
            verifyDupOpCode(offset, stateProof, encoded);
        } else if (opCode >= 0x90 && opCode <= 0x9f) {
            // SWAP
            verifySwapOpCode(offset, stateProof, encoded);
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

    function verifyStackOnlyOpcode(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        uint256 push
    ) internal pure {
        // Simulate pushing the `push` element to the stack
        bytes32 h = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, push));
        stateProof.stackHash = h;
        stateProof.stackSize += 1;
        // Increment pc
        stateProof.pc += 1;
    }

    function verifyUnaryOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        // Decode StackProof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint8 opCode = stateProof.opCode;
        uint64 gasCost;
        // Get the operand from stackProof
        uint256 a = stackProof.pops[0];
        // Simulate the opcode execution, get result and gas cost
        uint256 expected;
        if (opCode == 0x15) {
            assembly {
                expected := iszero(a)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x19) {
            assembly {
                expected := not(a)
            }
            gasCost = Params.G_VERYLOW;
        } else {
            revert("Unreachable");
        }
        // Consume gas cost
        if (stateProof.gas < gasCost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= gasCost;
        // Simulate stack pop
        stateProof.stackSize -= 1;
        // Verify the stack proof and reconstruct the state after one-step execution
        verifyStackOnlyOpcode(stateProof, stackProof, expected);
    }

    function verifyBinaryOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 2) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        // Decode StackProof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2, offset, encoded);

        uint8 opCode = stateProof.opCode;
        uint64 gasCost;
        // Get operands from stackProof
        uint256 a = stackProof.pops[0];
        uint256 b = stackProof.pops[1];
        // Simulate the opcode execution, get result and gas cost
        uint256 expected;
        if (opCode == 0x01) {
            assembly {
                expected := add(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x02) {
            assembly {
                expected := mul(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x03) {
            assembly {
                expected := sub(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x04) {
            assembly {
                expected := div(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x05) {
            assembly {
                expected := sdiv(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x06) {
            assembly {
                expected := mod(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x07) {
            assembly {
                expected := smod(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x0a) {
            assembly {
                expected := exp(a, b)
            }
            gasCost = GasTable.gasExp(b);
        } else if (opCode == 0x0b) {
            assembly {
                expected := signextend(a, b)
            }
            gasCost = Params.G_LOW;
        } else if (opCode == 0x10) {
            assembly {
                expected := lt(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x11) {
            assembly {
                expected := gt(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x12) {
            assembly {
                expected := slt(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x13) {
            assembly {
                expected := sgt(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x14) {
            assembly {
                expected := eq(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x16) {
            assembly {
                expected := and(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x17) {
            assembly {
                expected := or(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x18) {
            assembly {
                expected := xor(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x1a) {
            assembly {
                expected := byte(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x1b) {
            assembly {
                expected := shl(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x1c) {
            assembly {
                expected := shr(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else if (opCode == 0x1d) {
            assembly {
                expected := sar(a, b)
            }
            gasCost = Params.G_VERYLOW;
        } else {
            revert("Unreachable");
        }
        // Consume gas cost
        if (stateProof.gas < gasCost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= gasCost;
        // Simulate stack pop
        stateProof.stackSize -= 2;
        // Verify the stack proof and reconstruct the state after one-step execution
        verifyStackOnlyOpcode(stateProof, stackProof, expected);
    }

    function verifyTrinaryOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 3) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        // Decode StackProof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 3, offset, encoded);

        uint8 opCode = stateProof.opCode;
        uint64 gasCost;
        // Get operands from stackProof
        uint256 a = stackProof.pops[0];
        uint256 b = stackProof.pops[1];
        uint256 c = stackProof.pops[2];
        // Simulate the opcode execution, get result and gas cost
        uint256 expected;
        if (opCode == 0x08) {
            assembly {
                expected := addmod(a, b, c)
            }
            gasCost = Params.G_MID;
        } else if (opCode == 0x09) {
            assembly {
                expected := mulmod(a, b, c)
            }
            gasCost = Params.G_MID;
        } else {
            revert("Unreachable");
        }
        // Consume gas cost
        if (stateProof.gas < gasCost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= gasCost;
        // Simulate stack pop
        stateProof.stackSize -= 3;
        // Verify the stack proof and reconstruct the state after one-step execution
        verifyStackOnlyOpcode(stateProof, stackProof, expected);
    }

    function verifyPushOpCode(
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.CodeProof memory codeProof,
        bytes calldata encoded
    ) internal pure {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        uint8 opCode = stateProof.opCode;
        // Get the number of bytes to push
        uint8 pushBytes = opCode - 0x5f;

        // Slice out the push bytes from the bytecode
        uint256 expected = 0;
        // TODO: optimize this with assembly
        for (uint64 i = 1; i <= pushBytes; i++) {
            expected *= 256;
            uint64 pushOffset = stateProof.pc + i;
            if (pushOffset + 1 <= codeProof.size) {
                expected += codeProof.getOpCodeAt(encoded, pushOffset);
            }
        }
        // Simulate pushing the content to the stack
        bytes32 h = keccak256(abi.encodePacked(stateProof.stackHash, expected));
        stateProof.stackHash = h;
        stateProof.stackSize += 1;

        // Consume the gas
        if (stateProof.gas < Params.G_VERYLOW) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_VERYLOW;
        // Skip the opcode and content, jump to the correct pc
        stateProof.pc += 1 + pushBytes;
    }

    function verifyDupOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        uint8 opCode = stateProof.opCode;
        // Get the stack duplication position
        uint8 dupPos = opCode - 0x80 + 1;
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        if (stateProof.stackSize < dupPos) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Decode the stack proof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, dupPos, offset, encoded);

        // Simulate duplicating the stack content
        // pops[0] is the element to be duplicated
        bytes32 h = keccak256(abi.encodePacked(stateProof.stackHash, stackProof.pops[0]));
        stateProof.stackHash = h;
        stateProof.stackSize += 1;

        // Consume the gas
        if (stateProof.gas < Params.G_VERYLOW) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_VERYLOW;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifySwapOpCode(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        uint8 opCode = stateProof.opCode;
        // Get the stack swap position
        uint8 swapPos = opCode - 0x90 + 2;
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        if (stateProof.stackSize < swapPos) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Decode the stack proof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, swapPos, offset, encoded);

        // Simulate pushing the first swapped element
        bytes32 h = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, stackProof.pops[swapPos - 1]));
        // Simulate pushing the intermediate unchanged elements
        for (uint8 i = swapPos - 1; i > 1; i--) {
            h = keccak256(abi.encodePacked(h, stackProof.pops[i - 1]));
        }
        // Simulate pushing the last swapped element
        h = keccak256(abi.encodePacked(h, stackProof.pops[0]));
        stateProof.stackHash = h;

        // Consume the gas
        if (stateProof.gas < Params.G_VERYLOW) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_VERYLOW;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpPOP(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Decode the stack proof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        // `stackHashAfterPops` is the exact stack hash in the next step state
        stateProof.stackHash = stackProof.stackHashAfterPops;
        stateProof.stackSize -= 1;

        // Consume the gas
        if (stateProof.gas < Params.G_BASE) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_BASE;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpJUMP(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        // Decode the stack proof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        // `stackHashAfterPops` is the exact stack hash in the next step state
        stateProof.stackHash = stackProof.stackHashAfterPops;
        stateProof.stackSize -= 1;

        // Consume the gas
        if (stateProof.gas < Params.G_MID) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_MID;
        // Jump the pc according to the popped element from the stack
        uint64 nextPc = uint64(stackProof.pops[0]);
        stateProof.pc = nextPc;
    }

    function verifyOpJUMPI(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 2) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Decode the stack proof
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 2, offset, encoded);

        // `stackHashAfterPops` is the exact stack hash in the next step state
        stateProof.stackHash = stackProof.stackHashAfterPops;
        stateProof.stackSize -= 2;

        // Calculate the correct next pc
        uint64 nextPc = stateProof.pc + 1;
        if (stackProof.pops[0] != 0) {
            nextPc = uint64(stackProof.pops[1]);
        }

        // Consume the gas
        if (stateProof.gas < Params.G_HIGH) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_HIGH;
        // Jump the pc
        stateProof.pc = nextPc;
    }

    function verifyOpPC(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Consume the gas
        if (stateProof.gas < Params.G_BASE) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_BASE;
        // Simulate pushing `pc` to the stack
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, uint256(stateProof.pc)));
        stateProof.stackSize += 1;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpMSIZE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Consume the gas
        if (stateProof.gas < Params.G_BASE) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_BASE;
        // Simulate pushing `msize` to the stack
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, uint256(stateProof.memSize)));
        stateProof.stackSize += 1;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpGAS(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        // Consume the gas
        if (stateProof.gas < Params.G_BASE) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_BASE;
        // Simulate pushing `gas` to the stack
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, uint256(stateProof.gas)));
        stateProof.stackSize += 1;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpJUMPDEST(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        // Consume the gas
        if (stateProof.gas < Params.G_JUMPDEST) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= Params.G_JUMPDEST;
        // Increment the pc
        stateProof.pc += 1;
    }
}
