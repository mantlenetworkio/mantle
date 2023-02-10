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

import "./OneStepProof.sol";
import "./VerificationContext.sol";
import "./EVMTypesLib.sol";
import "./BloomLib.sol";
import "./Params.sol";

library VerifierHelper {
    using VerificationContext for VerificationContext.Context;
    using BloomLib for BloomLib.Bloom;
    using OneStepProof for OneStepProof.LogProof;

    function decodeAndVerifyStackProof(
        OneStepProof.StateProof memory stateProof,
        uint64 popNum,
        uint64 offset,
        bytes calldata encoded
    ) internal pure returns (uint64, OneStepProof.StackProof memory stackProof) {
        // Decode StackProof
        (offset, stackProof) = OneStepProof.decodeStackProof(encoded, offset, popNum);
        if (popNum != 0) {
            // Reconstruct the stack hash before pops
            bytes32 h = stackProof.stackHashAfterPops;
            for (uint8 i = uint8(stackProof.pops.length); i > 0; i--) {
                h = keccak256(abi.encodePacked(h, stackProof.pops[i - 1]));
            }
            // Ensure the stack hash reconstructed is the same as the stack hash in stateProof
            require(h == stateProof.stackHash, "Bad StackProof");
        } else {
            // If no pop, the stackHashAfterPops is just the current stack hash
            stackProof.stackHashAfterPops = stateProof.stackHash;
        }
        return (offset, stackProof);
    }

    function decodeAndVerifyBlockHashProof(
        uint64 offset,
        bytes calldata encoded,
        uint64 blockNumber,
        bytes32 blockHashRoot
    ) internal pure returns (uint64, bytes32) {
        OneStepProof.BlockHashProof memory blockHashProof;
        (offset, blockHashProof) = OneStepProof.decodeBlockHashProof(encoded, offset);
        OneStepProof.BlockHashMerkleProof memory blockHashMerkleProof;
        (offset, blockHashMerkleProof) = OneStepProof.decodeBlockHashMerkleProof(encoded, offset);
        // Verify blockhash
        bytes32 h = blockHashProof.blockHash;
        uint64 path = blockHashMerkleProof.path;
        uint64 num = 0;
        for (uint8 i = 0; i < blockHashMerkleProof.proof.length; i++) {
            if (path & 0x1 == 1) {
                h = keccak256(abi.encodePacked(h, blockHashMerkleProof.proof[i]));
                num |= uint64(1 << i);
            } else {
                h = keccak256(abi.encodePacked(blockHashMerkleProof.proof[i], h));
            }
            path = path >> 1;
        }
        require(h == blockHashRoot, "Bad BlockHashProof");
        require(num == blockNumber, "Bad BlockHashProof");

        return (offset, blockHashProof.blockHash);
    }

    function decodeAndInsertBlockHashProof(
        uint64 offset,
        bytes calldata encoded,
        uint64 blockNumber,
        bytes32 blockHashRoot,
        bytes32 newBlockHash
    ) internal pure returns (uint64, bytes32) {
        OneStepProof.BlockHashProof memory blockHashProof;
        (offset, blockHashProof) = OneStepProof.decodeBlockHashProof(encoded, offset);
        OneStepProof.BlockHashMerkleProof memory blockHashMerkleProof;
        (offset, blockHashMerkleProof) = OneStepProof.decodeBlockHashMerkleProof(encoded, offset);
        // Verify blockhash
        bytes32 h = blockHashProof.blockHash;
        uint64 path = blockHashMerkleProof.path;
        uint64 num = 0;
        for (uint8 i = 0; i < blockHashMerkleProof.proof.length; i++) {
            if (path & 0x1 == 1) {
                h = keccak256(abi.encodePacked(h, blockHashMerkleProof.proof[i]));
                num |= uint64(1 << i);
            } else {
                h = keccak256(abi.encodePacked(blockHashMerkleProof.proof[i], h));
            }
            path = path >> 1;
        }
        require(h == blockHashRoot, "Bad BlockHashProof");
        require(num == blockNumber, "Bad BlockHashProof");

        h = newBlockHash;
        path = blockHashMerkleProof.path;
        for (uint8 i = 0; i < blockHashMerkleProof.proof.length; i++) {
            if (path & 0x1 == 1) {
                h = keccak256(abi.encodePacked(h, blockHashMerkleProof.proof[i]));
            } else {
                h = keccak256(abi.encodePacked(blockHashMerkleProof.proof[i], h));
            }
            path = path >> 1;
        }
        // h is the new root now
        return (offset, h);
    }

    function updateAndHashLogProof(
        OneStepProof.LogProof memory logProof,
        address addr,
        uint256[] memory topics,
        bytes memory data
    ) internal pure returns (bytes32) {
        bytes32 logHash = EVMTypesLib.hashLogEntry(addr, topics, data);
        logProof.accumulateHash = keccak256(abi.encodePacked(logProof.accumulateHash, logHash));
        logProof.bloom.add(addr);
        for (uint8 i = 0; i < topics.length; i++) {
            logProof.bloom.add(bytes32(topics[i]));
        }
        return logProof.hashLogProof();
    }

    function verifyRevertByError(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
        returns (OneStepProof.StateProof memory)
    {
        bytes32 lastDepthHash = stateProof.lastDepthHash;
        if (stateProof.depth == 1) {}
        // require(offset == encoded.length, "Proof Overflow");
        return stateProof;
    }

    function verifyTransfer(bytes32 worldState, uint256 value, uint64 offset, bytes calldata encoded)
        internal
        pure
        returns (bytes32)
    {}

    function verifyTransactionInitiation(
        VerificationContext.Context memory ctx,
        OneStepProof.InterStateProof memory stateProof,
        EVMTypesLib.Account memory senderAccountProof,
        uint64 offset,
        bytes calldata encoded
    ) internal pure returns (bytes32) {}

    function hashBlockHeader(
        VerificationContext.Context memory ctx,
        bytes32 parentBlockHash,
        OneStepProof.InterStateProof memory stateProof
    ) internal pure returns (bytes32) {
        EVMTypesLib.BlockHeader memory blockHeader;
        blockHeader.parentHash = parentBlockHash;
        blockHeader.ommerHash = Params.EMPTY_UNCLE_HASH;
        blockHeader.beneficiary = ctx.getCoinbase();
        blockHeader.stateRoot = stateProof.globalStateRoot;
        blockHeader.transactionRoot = stateProof.transactionTrieRoot;
        blockHeader.receiptsRoot = stateProof.receiptTrieRoot;
        blockHeader.difficulty = ctx.getDifficulty();
        blockHeader.number = stateProof.blockNumber;
        blockHeader.gasLimit = ctx.getGasLimit();
        blockHeader.gasUsed = uint64(stateProof.blockGasUsed);
        blockHeader.timestamp = uint64(ctx.getTimestamp());
        blockHeader.logsBloom = stateProof.logsBloom;
        return EVMTypesLib.hashBlockHeader(blockHeader);
    }

    function BuildCreateState(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildCreate2State(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildCallState(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildCallCodeState(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildDelegateCallState(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildStaticCallState(
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StackProof memory stackProof,
        EVMTypesLib.Account memory accountProof,
        OneStepProof.MemoryReadProof memory memoryReadProof
    ) internal pure {}

    function BuildIntraReturnState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof,
        OneStepProof.MemoryReadProof memory memoryReadProof,
        OneStepProof.MemoryWriteProof memory memoryWriteProof
    ) internal pure {}

    function BuildInterReturnState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.InterStateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof
    ) internal pure {}

    function BuildIntraRevertState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof,
        OneStepProof.MemoryReadProof memory memoryReadProof,
        OneStepProof.MemoryWriteProof memory memoryWriteProof
    ) internal pure {}

    function BuildInterRevertState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.InterStateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof
    ) internal pure {}

    function BuildIntraSuicideState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.StateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof,
        OneStepProof.MemoryReadProof memory memoryReadProof,
        OneStepProof.MemoryWriteProof memory memoryWriteProof
    ) internal pure {}

    function BuildInterSuicideState(
        VerificationContext.Context memory ctx,
        OneStepProof.StateProof memory stateProof,
        OneStepProof.InterStateProof memory lastStateProof,
        OneStepProof.StackProof memory stackProof
    ) internal pure {}
}
