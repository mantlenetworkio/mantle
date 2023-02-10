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
import "../libraries/VerifierHelper.sol";
import "../libraries/OneStepProof.sol";
import "../libraries/Params.sol";
import "../libraries/GasTable.sol";

contract EnvironmentalOpVerifier is IVerifier {
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
        if (opCode == 0x30) {
            // ADDRESS
            verifyOpADDRESS(offset, stateProof, encoded);
        } else if (opCode == 0x32) {
            // ORIGIN
            verifyOpORIGIN(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x33) {
            // CALLER
            verifyOpCALLER(offset, stateProof, encoded);
        } else if (opCode == 0x34) {
            // CALLVALUE
            verifyOpCALLVALUE(offset, stateProof, encoded);
        } else if (opCode == 0x38) {
            // CODESIZE
            verifyOpCODESIZE(offset, stateProof, codeProof, encoded);
        } else if (opCode == 0x36) {
            // CALLDATASIZE
            verifyOpCALLDATASIZE(offset, stateProof, encoded);
        } else if (opCode == 0x3a) {
            // GASPRICE
            verifyOpGASPRICE(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x3d) {
            // RETURNDATASIZE
            verifyOpRETURNDATASIZE(offset, stateProof, encoded);
        } else if (opCode == 0x40) {
            // BLOCKHASH
            verifyOpBLOCKHASH(offset, stateProof, encoded);
        } else if (opCode == 0x41) {
            // COINBASE
            verifyOpCOINBASE(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x42) {
            // TIMESTAMP
            verifyOpTIMESTAMP(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x43) {
            // NUMBER
            verifyOpNUMBER(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x44) {
            // DIFFICULTY
            verifyOpDIFFICULTY(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x45) {
            // GASLIMIT
            verifyOpGASLIMIT(ctx, offset, stateProof, encoded);
        } else if (opCode == 0x46) {
            // CHAINID
            verifyOpCHAINID(ctx, offset, stateProof, encoded);
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

    function verifyOnePushOpcode(
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        uint64 gas,
        uint256 push,
        bytes calldata encoded
    ) internal pure {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        // Consume the gas
        if (stateProof.gas < gas) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= gas;

        // Simulate pushing `push` to the stack
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, push));
        stateProof.stackSize += 1;
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpADDRESS(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(uint160(stateProof.contractAddress)), encoded);
    }

    function verifyOpORIGIN(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(uint160(ctx.getOrigin())), encoded);
    }

    function verifyOpCALLER(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(uint160(stateProof.caller)), encoded);
    }

    function verifyOpCALLVALUE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(uint160(stateProof.value)), encoded);
    }

    function verifyOpCODESIZE(
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.CodeProof memory codeProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(codeProof.size), encoded);
    }

    function verifyOpCALLDATASIZE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(stateProof.inputDataSize), encoded);
    }

    function verifyOpGASPRICE(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, ctx.getGasPrice(), encoded);
    }

    function verifyOpRETURNDATASIZE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(stateProof.returnDataSize), encoded);
    }

    function verifyOpBLOCKHASH(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        // Consume the gas
        uint64 gas = Params.G_BLOCKHASH;
        if (stateProof.gas < gas) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        stateProof.gas -= gas;

        uint64 num = uint64(stackProof.pops[0]);
        bytes32 bhash;
        if (stateProof.blockNumber >= num || stateProof.blockNumber < num - Params.RECENT_BLOCK_HASHES_LENGTH) {
            bhash = 0x00;
        } else {
            (offset, bhash) =
                VerifierHelper.decodeAndVerifyBlockHashProof(offset, encoded, num, stateProof.blockHashRoot);
        }

        // Simulate pushing `bhash` to the stack
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, bhash));
        // Increment the pc
        stateProof.pc += 1;
    }

    function verifyOpCOINBASE(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(uint160(ctx.getCoinbase())), encoded);
    }

    function verifyOpTIMESTAMP(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, ctx.getTimestamp(), encoded);
    }

    function verifyOpNUMBER(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, ctx.getBlockNumber(), encoded);
    }

    function verifyOpDIFFICULTY(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, ctx.getDifficulty(), encoded);
    }

    function verifyOpGASLIMIT(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(ctx.getGasLimit()), encoded);
    }

    function verifyOpCHAINID(
        VerificationContext.Context memory ctx,
        uint64 offset,
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded
    ) internal pure {
        verifyOnePushOpcode(offset, stateProof, Params.G_BASE, uint256(ctx.getChainID()), encoded);
    }
}
