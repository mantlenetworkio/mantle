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

import "../../libraries/RLPReader.sol";
import "../../libraries/BytesLib.sol";
import "./BloomLib.sol";
import "./VerificationContext.sol";

library OneStepProof {
    using BytesLib for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;
    using VerificationContext for VerificationContext.Context;

    // [encode rule]
    struct StateProof {
        uint64 blockNumber; // Block number of current transaction [always]
        uint64 transactionIdx; // Transaction index in block [always]
        uint16 depth; // Current call depth [always]
        uint64 gas; // Gas left in the current call [always]
        uint64 refund; // Gas refund accumulated in the current transaction [always]
        bytes32 lastDepthHash; // The state hash of the last depth call frame [always]
        address contractAddress; // Current executing contract address [depth > 1]
        address caller; // Current caller [depth > 1]
        uint256 value; // Current call value [depth > 1]
        uint8 callFlag; // Current call type [depth > 1]
        uint64 out; // Offset of the return data of current call to be copied to the last depth call frame [depth > 1]
        uint64 outSize; // Size of the return data of current call to be copied to the last depth call frame [depth > 1]
        uint64 pc; // Current program counter [always]
        uint8 opCode; // Current opcode to be executed [always]
        bytes32 codeHash; // Current executing contract code hash [always]
        uint64 stackSize; // Size of the stack [always]
        bytes32 stackHash; // Commitment of the stack [always]
        uint64 memSize; // Size of the memory [always]
        bytes32 memRoot; // Commitment of the memory [memSize > 0]
        uint64 inputDataSize; // Size of the call data [depth > 1]
        bytes32 inputDataRoot; // Commitment of the return data [depth > 1 && inputDataSize > 0]
        uint64 returnDataSize; // Size of the return data [always]
        bytes32 returnDataRoot; // Commitment of the return data [returnDataSize > 0]
        bytes32 committedGlobalStateRoot; // Commitment of the global MPT state at the start of transaction [always]
        bytes32 globalStateRoot; // Commitment of the global MPT state [always]
        bytes32 selfDestructAcc; // Commitment of the self destructed contracts in the current transaction [always]
        bytes32 logAcc; // Commitment of the logs emitted in the current transaction [always]
        bytes32 blockHashRoot; // Commitment of the 256 previous blockhash in the current block [always]
        bytes32 accessListRoot; // Commitment of the access list in the current transaction [always]
    }

    function decodeStateProof(VerificationContext.Context memory ctx, bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, StateProof memory proof)
    {
        uint64 remainLength = uint64(encoded.length) - offset;
        uint64 stateProofLen = 323;
        require(remainLength >= stateProofLen, "Proof Underflow (State)");
        proof.blockNumber = encoded.toUint64(offset);
        proof.transactionIdx = encoded.toUint64(offset + 8);
        proof.depth = encoded.toUint16(offset + 16);
        proof.gas = encoded.toUint64(offset + 18);
        proof.refund = encoded.toUint64(offset + 26);
        proof.lastDepthHash = encoded.toBytes32(offset + 34);
        offset = offset + 66;
        if (proof.depth > 1) {
            stateProofLen += 97;
            require(remainLength >= stateProofLen, "Proof Underflow (State)");
            proof.contractAddress = encoded.toAddress(offset);
            proof.caller = encoded.toAddress(offset + 20);
            proof.value = encoded.toUint256(offset + 40);
            proof.callFlag = encoded.toUint8(offset + 72);
            proof.out = encoded.toUint64(offset + 73);
            proof.outSize = encoded.toUint64(offset + 81);
            offset += 89;
        } else {
            proof.contractAddress = ctx.getRecipient();
            proof.caller = ctx.getOrigin();
            proof.value = ctx.getValue();
            if (ctx.getRecipient() == address(0)) {
                proof.callFlag = 4;
            } else {
                proof.callFlag = 0;
            }
        }
        proof.pc = encoded.toUint64(offset);
        proof.opCode = encoded.toUint8(offset + 8);
        proof.codeHash = encoded.toBytes32(offset + 9);
        proof.stackSize = encoded.toUint64(offset + 41);
        offset += 49;
        if (proof.stackSize != 0) {
            stateProofLen += 32;
            require(remainLength >= stateProofLen, "Proof Underflow (State)");
            proof.stackHash = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.memSize = encoded.toUint64(offset);
        offset += 8;
        if (proof.memSize != 0) {
            stateProofLen += 32;
            require(remainLength >= stateProofLen, "Proof Underflow (State)");
            proof.memRoot = encoded.toBytes32(offset);
            offset += 32;
        }
        if (proof.depth > 1) {
            proof.inputDataSize = encoded.toUint64(offset);
            offset += 8;
            if (proof.inputDataSize != 0) {
                stateProofLen += 32;
                require(remainLength >= stateProofLen, "Proof Underflow (State)");
                proof.inputDataRoot = encoded.toBytes32(offset);
                offset += 32;
            }
        } else {
            proof.inputDataSize = ctx.getInputSize();
            proof.inputDataRoot = ctx.getInputRoot();
        }
        proof.returnDataSize = encoded.toUint64(offset);
        offset += 8;
        if (proof.returnDataSize != 0) {
            stateProofLen += 32;
            require(remainLength >= stateProofLen, "Proof Underflow (State)");
            proof.returnDataRoot = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.committedGlobalStateRoot = encoded.toBytes32(offset);
        proof.globalStateRoot = encoded.toBytes32(offset + 32);
        proof.selfDestructAcc = encoded.toBytes32(offset + 64);
        proof.logAcc = encoded.toBytes32(offset + 96);
        proof.blockHashRoot = encoded.toBytes32(offset + 128);
        proof.accessListRoot = encoded.toBytes32(offset + 160);
        return (offset + 192, proof);
    }

    function encodeStateProof(StateProof memory proof) internal pure returns (bytes memory encoded) {
        encoded = encoded.concat(abi.encodePacked(proof.blockNumber));
        encoded = encoded.concat(abi.encodePacked(proof.transactionIdx));
        encoded = encoded.concat(abi.encodePacked(proof.depth));
        encoded = encoded.concat(abi.encodePacked(proof.gas));
        encoded = encoded.concat(abi.encodePacked(proof.refund));
        encoded = encoded.concat(abi.encodePacked(proof.lastDepthHash));
        if (proof.depth > 1) {
            encoded = encoded.concat(abi.encodePacked(proof.contractAddress));
            encoded = encoded.concat(abi.encodePacked(proof.caller));
            encoded = encoded.concat(abi.encodePacked(proof.value));
            encoded = encoded.concat(abi.encodePacked(proof.callFlag));
            encoded = encoded.concat(abi.encodePacked(proof.out));
            encoded = encoded.concat(abi.encodePacked(proof.outSize));
        }
        encoded = encoded.concat(abi.encodePacked(proof.pc));
        encoded = encoded.concat(abi.encodePacked(proof.opCode));
        encoded = encoded.concat(abi.encodePacked(proof.codeHash));
        encoded = encoded.concat(abi.encodePacked(proof.stackSize));
        if (proof.stackSize != 0) {
            encoded = encoded.concat(abi.encodePacked(proof.stackHash));
        }
        encoded = encoded.concat(abi.encodePacked(proof.memSize));
        if (proof.memSize != 0) {
            encoded = encoded.concat(abi.encodePacked(proof.memRoot));
        }
        if (proof.depth > 1) {
            encoded = encoded.concat(abi.encodePacked(proof.inputDataSize));
            if (proof.inputDataSize != 0) {
                encoded = encoded.concat(abi.encodePacked(proof.inputDataRoot));
            }
        }
        encoded = encoded.concat(abi.encodePacked(proof.returnDataSize));
        if (proof.returnDataSize != 0) {
            encoded = encoded.concat(abi.encodePacked(proof.returnDataRoot));
        }
        encoded = encoded.concat(abi.encodePacked(proof.committedGlobalStateRoot));
        encoded = encoded.concat(abi.encodePacked(proof.globalStateRoot));
        encoded = encoded.concat(abi.encodePacked(proof.selfDestructAcc));
        encoded = encoded.concat(abi.encodePacked(proof.logAcc));
        encoded = encoded.concat(abi.encodePacked(proof.blockHashRoot));
        encoded = encoded.concat(abi.encodePacked(proof.accessListRoot));
    }

    function hashStateProof(StateProof memory proof) internal pure returns (bytes32) {
        if (proof.depth == 0) {
            // When returning/reverting from depth 1, we can't directly return an InterStateProof
            // Therefore we reuse some of the fields in the IntraStateProof to store an InterStateProof
            // The field mappings are as follows:
            InterStateProof memory interProof;
            interProof.blockNumber = proof.blockNumber;
            interProof.transactionIdx = proof.transactionIdx;
            interProof.globalStateRoot = proof.globalStateRoot;
            interProof.cumulativeGasUsed = proof.value;
            interProof.blockGasUsed = uint256(proof.lastDepthHash);
            interProof.blockHashRoot = proof.blockHashRoot;
            interProof.transactionTrieRoot = proof.selfDestructAcc;
            interProof.receiptTrieRoot = proof.logAcc;
            return hashInterStateProof(interProof);
        }
        return keccak256(encodeStateProof(proof));
    }

    struct InterStateProof {
        uint64 blockNumber;
        uint64 transactionIdx;
        bytes32 globalStateRoot;
        uint256 cumulativeGasUsed;
        uint256 blockGasUsed;
        bytes32 blockHashRoot;
        bytes32 transactionTrieRoot;
        bytes32 receiptTrieRoot;
        BloomLib.Bloom logsBloom;
    }

    function decodeInterStateProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, InterStateProof memory proof)
    {
        require(encoded.length - offset >= 464, "Proof Underflow (Inter)");
        proof.blockNumber = encoded.toUint64(offset);
        proof.transactionIdx = encoded.toUint64(offset + 8);
        proof.globalStateRoot = encoded.toBytes32(offset + 16);
        proof.cumulativeGasUsed = encoded.toUint64(offset + 48);
        proof.blockGasUsed = encoded.toUint64(offset + 80);
        proof.blockHashRoot = encoded.toBytes32(offset + 112);
        proof.transactionTrieRoot = encoded.toBytes32(offset + 144);
        proof.receiptTrieRoot = encoded.toBytes32(offset + 176);
        proof.logsBloom = BloomLib.decodeBloom(encoded, offset + 208);
        return (offset + 464, proof);
    }

    function encodeInterStateProof(InterStateProof memory proof) internal pure returns (bytes memory encoded) {
        encoded = encoded.concat(abi.encodePacked(proof.blockNumber));
        encoded = encoded.concat(abi.encodePacked(proof.transactionIdx));
        encoded = encoded.concat(abi.encodePacked(proof.globalStateRoot));
        encoded = encoded.concat(abi.encodePacked(proof.cumulativeGasUsed));
        encoded = encoded.concat(abi.encodePacked(proof.blockGasUsed));
        encoded = encoded.concat(abi.encodePacked(proof.blockHashRoot));
        encoded = encoded.concat(abi.encodePacked(proof.transactionTrieRoot));
        encoded = encoded.concat(abi.encodePacked(proof.receiptTrieRoot));
        encoded = encoded.concat(abi.encodePacked(proof.logsBloom.data));
    }

    function hashInterStateProof(InterStateProof memory proof) internal pure returns (bytes32) {
        return keccak256(encodeInterStateProof(proof));
    }

    struct BlockStateProof {
        uint64 blockNumber;
        bytes32 globalStateRoot;
        uint256 cumulativeGasUsed;
        bytes32 blockHashRoot;
    }

    function decodeBlockStateProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, BlockStateProof memory proof)
    {
        require(encoded.length - offset >= 104, "Proof Underflow (Block)");
        proof.blockNumber = encoded.toUint64(offset);
        proof.globalStateRoot = encoded.toBytes32(offset + 8);
        proof.cumulativeGasUsed = encoded.toUint64(offset + 40);
        proof.blockHashRoot = encoded.toBytes32(offset + 72);
        return (offset + 104, proof);
    }

    function encodeBlockStateProof(BlockStateProof memory proof) internal pure returns (bytes memory encoded) {
        encoded = encoded.concat(abi.encodePacked(proof.blockNumber));
        encoded = encoded.concat(abi.encodePacked(proof.globalStateRoot));
        encoded = encoded.concat(abi.encodePacked(proof.cumulativeGasUsed));
        encoded = encoded.concat(abi.encodePacked(proof.blockHashRoot));
    }

    function hashBlockStateProof(BlockStateProof memory proof) internal pure returns (bytes32) {
        return keccak256(encodeBlockStateProof(proof));
    }

    struct CodeProof {
        uint64 ptr;
        uint64 size;
    }

    function decodeCodeProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, CodeProof memory proof)
    {
        require(encoded.length - offset >= 8, "Proof Underflow (Code)");
        // Decode bytecode size in bytes
        uint64 contentSize = encoded.toUint64(offset);
        require(encoded.length - offset >= 8 + contentSize, "Proof Underflow (Code)");
        offset += 8;
        proof.ptr = offset;
        proof.size = contentSize;
        return (offset + contentSize, proof);
    }

    function getOpCodeAt(CodeProof memory proof, bytes calldata encoded, uint64 idx) internal pure returns (uint8) {
        if (idx >= proof.size) {
            return 0;
        }
        return uint8(encoded[proof.ptr + idx]);
    }

    function getCodeSlice(CodeProof memory proof, bytes calldata encoded, uint64 offset, uint64 size)
        internal
        pure
        returns (bytes memory)
    {
        if (offset + size > proof.size) {
            return encoded.slice(proof.ptr + offset, size).concat(new bytes(size - (proof.size - offset)));
        }
        return encoded.slice(proof.ptr + offset, size);
    }

    function hashCodeProof(CodeProof memory proof, bytes calldata encoded) internal pure returns (bytes32) {
        return keccak256(encoded[proof.ptr:proof.ptr + proof.size]);
    }

    struct StackProof {
        // The elements popped in the step
        uint256[] pops;
        // The stack hash after popping above elements
        bytes32 stackHashAfterPops;
    }

    function decodeStackProof(bytes calldata encoded, uint64 offset, uint64 popNum)
        internal
        pure
        returns (uint64, StackProof memory proof)
    {
        if (popNum == 0) {
            // No StackProof needed for popNum == 0
            return (offset, proof);
        }
        require(encoded.length - offset >= 32 * (popNum + 1), "Proof Underflow (Stack)");
        proof.pops = new uint256[](popNum);
        // Decode popped elements
        for (uint64 i = 0; i < popNum; i++) {
            proof.pops[i] = encoded.toUint256(offset);
            offset += 32;
        }
        // Decode stackHashAfterPops
        proof.stackHashAfterPops = encoded.toBytes32(offset);
        offset += 32;
        return (offset, proof);
    }

    function encodeStackProof(StackProof memory proof) internal pure returns (bytes memory encoded) {
        for (uint64 i = 0; i < proof.pops.length; i++) {
            encoded = encoded.concat(abi.encodePacked(proof.pops[i]));
        }
        encoded = encoded.concat(abi.encodePacked(proof.stackHashAfterPops));
    }

    struct MemoryMerkleProof {
        bytes32[] proof;
    }

    function decodeMemoryMerkleProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, MemoryMerkleProof memory proof)
    {
        require(encoded.length - offset >= 8, "Proof Underflow");
        uint64 len = encoded.toUint64(offset);
        offset += 8;
        require(encoded.length - offset >= 32 * len, "Proof Underflow");
        proof.proof = new bytes32[](len);
        for (uint64 i = 0; i < len; i++) {
            proof.proof[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct MemoryReadProof {
        bytes32[] cells;
    }

    function decodeMemoryReadProof(bytes calldata encoded, uint64 offset, uint64 cellNum)
        internal
        pure
        returns (uint64, MemoryReadProof memory proof)
    {
        require(encoded.length - offset >= 32 * cellNum, "Proof Underflow");
        proof.cells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.cells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct MemoryWriteProof {
        bytes32[] cells;
        bytes32[] updatedCells;
    }

    function decodeMemoryWriteProof(bytes calldata encoded, uint64 offset, uint64 cellNum)
        internal
        pure
        returns (uint64, MemoryWriteProof memory proof)
    {
        require(encoded.length - offset >= 64 * cellNum, "Proof Underflow");
        proof.cells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.cells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.updatedCells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.updatedCells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct MemoryAppendProof {
        bytes32[] appendCells;
    }

    function decodeMemoryAppendProof(bytes calldata encoded, uint64 offset, uint64 cellNum)
        internal
        pure
        returns (uint64, MemoryAppendProof memory proof)
    {
        require(encoded.length - offset >= 32 * cellNum, "Proof Underflow");
        proof.appendCells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.appendCells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct MemoryCombinedReadProof {
        bytes32[] cells;
        bytes32[] appendCells;
    }

    function decodeMemoryCombinedReadProof(bytes calldata encoded, uint64 offset, uint64 cellNum, uint64 appendCellNum)
        internal
        pure
        returns (uint64, MemoryCombinedReadProof memory proof)
    {
        require(encoded.length - offset >= 32 * (cellNum + appendCellNum), "Proof Underflow");
        proof.cells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.cells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.appendCells = new bytes32[](appendCellNum);
        for (uint64 i = 0; i < appendCellNum; i++) {
            proof.appendCells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct MemoryCombinedWriteProof {
        bytes32[] cells;
        bytes32[] updatedCells;
        bytes32[] appendCells;
    }

    function decodeMemoryCombinedWriteProof(bytes calldata encoded, uint64 offset, uint64 cellNum, uint64 appendCellNum)
        internal
        pure
        returns (uint64, MemoryCombinedWriteProof memory proof)
    {
        require(encoded.length - offset >= 32 * (2 * cellNum + appendCellNum), "Proof Underflow");
        proof.cells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.cells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.updatedCells = new bytes32[](cellNum);
        for (uint64 i = 0; i < cellNum; i++) {
            proof.updatedCells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        proof.appendCells = new bytes32[](appendCellNum);
        for (uint64 i = 0; i < appendCellNum; i++) {
            proof.appendCells[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    // For MPT proof, receipt proof
    struct RLPProof {
        RLPReader.RLPItem proof;
    }

    function decodeRLPProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, RLPProof memory proof)
    {
        require(encoded.length - offset >= 8, "Proof Underflow");
        uint64 len = encoded.toUint64(offset);
        offset += 8;
        require(encoded.length - offset >= len, "Proof Underflow");
        proof.proof = encoded.slice(offset, len).toRlpItem();
        return (offset + len, proof);
    }

    struct BlockHashProof {
        bytes32 blockHash;
    }

    function decodeBlockHashProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, BlockHashProof memory proof)
    {
        require(encoded.length - offset >= 32, "Proof Underflow");
        proof.blockHash = encoded.toBytes32(offset);
        return (offset + 32, proof);
    }

    struct BlockHashMerkleProof {
        uint64 path;
        bytes32[] proof;
    }

    function decodeBlockHashMerkleProof(bytes memory encoded, uint64 offset)
        internal
        pure
        returns (uint64, BlockHashMerkleProof memory proof)
    {
        require(encoded.length - offset >= 9, "Proof Underflow");
        proof.path = encoded.toUint64(offset);
        uint8 len = encoded.toUint8(offset + 8);
        offset += 9;
        require(encoded.length - offset >= 32 * len, "Proof Underflow");
        proof.proof = new bytes32[](len);
        for (uint64 i = 0; i < len; i++) {
            proof.proof[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return (offset, proof);
    }

    struct LogProof {
        bytes32 accumulateHash;
        BloomLib.Bloom bloom;
    }

    function decodeLogProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, LogProof memory proof)
    {
        require(encoded.length - offset >= 288, "Proof Underflow");
        proof.accumulateHash = encoded.toBytes32(offset);
        proof.bloom = BloomLib.decodeBloom(encoded, offset + 32);
        return (offset + 288, proof);
    }

    function hashLogProof(LogProof memory proof) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(proof.accumulateHash, proof.bloom.data));
    }

    struct SelfDestructSetProof {
        address[] contracts;
    }

    function decodeSelfDestructSetProof(bytes calldata encoded, uint64 offset)
        internal
        pure
        returns (uint64, SelfDestructSetProof memory proof)
    {
        require(encoded.length - offset >= 8, "Proof Underflow");
        uint64 len = encoded.toUint64(offset);
        offset += 8;
        require(encoded.length - offset >= 20 * len, "Proof Underflow");
        proof.contracts = new address[](len);
        for (uint64 i = 0; i < len; i++) {
            proof.contracts[i] = encoded.toAddress(offset);
            offset += 20;
        }
        return (offset, proof);
    }
}
