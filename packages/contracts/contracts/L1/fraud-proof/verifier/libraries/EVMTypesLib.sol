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
import "../../libraries/RLPReader.sol";
import "../../libraries/RLPWriter.sol";
import "./BloomLib.sol";

library EVMTypesLib {
    using BytesLib for bytes;
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    struct BlockHeader {
        bytes32 parentHash;
        bytes32 ommerHash;
        address beneficiary;
        bytes32 stateRoot;
        bytes32 transactionRoot;
        bytes32 receiptsRoot;
        uint256 difficulty;
        uint256 number;
        uint64 gasLimit;
        uint64 gasUsed;
        uint64 timestamp;
        BloomLib.Bloom logsBloom;
    }

    function hashBlockHeader(BlockHeader memory header) internal pure returns (bytes32) {
        bytes[] memory raw = new bytes[](15);
        raw[0] = RLPWriter.writeBytes(abi.encodePacked(header.parentHash));
        raw[1] = RLPWriter.writeBytes(abi.encodePacked(header.ommerHash));
        raw[2] = RLPWriter.writeAddress(header.beneficiary);
        raw[3] = RLPWriter.writeBytes(abi.encodePacked(header.stateRoot));
        raw[4] = RLPWriter.writeBytes(abi.encodePacked(header.transactionRoot));
        raw[5] = RLPWriter.writeBytes(abi.encodePacked(header.receiptsRoot));
        raw[6] = RLPWriter.writeBytes(abi.encodePacked(header.logsBloom.data));
        raw[7] = RLPWriter.writeUint(header.difficulty);
        raw[8] = RLPWriter.writeUint(header.number);
        raw[9] = RLPWriter.writeUint(uint256(header.gasLimit));
        raw[10] = RLPWriter.writeUint(uint256(header.gasUsed));
        raw[11] = RLPWriter.writeUint(uint256(header.timestamp));
        raw[12] = RLPWriter.writeBytes(""); // Extra
        raw[13] = RLPWriter.writeBytes(abi.encodePacked(bytes32(0))); // MixDigest
        raw[14] = RLPWriter.writeBytes(abi.encodePacked(bytes8(0))); // Nonce
        return keccak256(RLPWriter.writeList(raw));
    }

    struct Transaction {
        uint64 nonce;
        uint256 gasPrice;
        uint64 gas;
        address to;
        uint256 value;
        bytes data;
        uint256 v;
        uint256 r;
        uint256 s;
    }

    function decodeTransaction(bytes memory data) internal pure returns (Transaction memory transaction) {
        RLPReader.RLPItem[] memory decoded = data.toRlpItem().toList();
        transaction.nonce = uint64(decoded[0].toUint());
        transaction.gasPrice = decoded[1].toUint();
        transaction.gas = uint64(decoded[2].toUint());
        transaction.to = address(uint160(decoded[3].toUint()));
        transaction.value = decoded[4].toUint();
        transaction.data = decoded[5].toBytes();
        transaction.v = decoded[6].toUint();
        transaction.r = decoded[7].toUint();
        transaction.s = decoded[8].toUint();
    }

    function hashTransaction(Transaction memory txn) internal pure returns (bytes32) {
        bytes[] memory raw = new bytes[](9);
        raw[0] = RLPWriter.writeUint(uint256(txn.nonce));
        raw[1] = RLPWriter.writeUint(txn.gasPrice);
        raw[2] = RLPWriter.writeUint(uint256(txn.gas));
        raw[3] = RLPWriter.writeAddress(txn.to);
        raw[4] = RLPWriter.writeUint(txn.value);
        raw[5] = RLPWriter.writeBytes(txn.data);
        raw[6] = RLPWriter.writeUint(txn.v);
        raw[7] = RLPWriter.writeUint(txn.r);
        raw[8] = RLPWriter.writeUint(txn.s);
        return keccak256(RLPWriter.writeList(raw));
    }

    struct Account {
        uint64 nonce;
        uint256 balance;
        bytes32 storageRoot;
        bytes32 codeHash;
    }

    function decodeAccount(RLPReader.RLPItem memory encoded) internal pure returns (Account memory proof) {
        RLPReader.RLPItem[] memory items = encoded.toList();
        require(items.length == 4, "Invalid Account");
        proof.nonce = uint64(items[0].toUint());
        proof.balance = items[1].toUint();
        proof.storageRoot = bytes32(items[2].toUint());
        proof.codeHash = bytes32(items[3].toUint());
    }

    function encodeRLP(Account memory account) internal pure returns (bytes memory) {
        bytes[] memory raw = new bytes[](4);
        raw[0] = RLPWriter.writeUint(uint256(account.nonce));
        raw[1] = RLPWriter.writeUint(account.balance);
        raw[2] = RLPWriter.writeBytes(abi.encodePacked(account.storageRoot));
        raw[3] = RLPWriter.writeBytes(abi.encodePacked(account.codeHash));
        return RLPWriter.writeList(raw);
    }

    function hashLogEntry(address addr, uint256[] memory topics, bytes memory data) internal pure returns (bytes32) {
        bytes[] memory topicRaw = new bytes[](topics.length);
        for (uint256 i = 0; i < topics.length; i++) {
            topicRaw[i] = RLPWriter.writeBytes(abi.encodePacked(bytes32(topics[i])));
        }
        bytes[] memory raw = new bytes[](3);
        raw[0] = RLPWriter.writeAddress(addr);
        raw[1] = RLPWriter.writeBytes(RLPWriter.writeList(topicRaw));
        raw[2] = RLPWriter.writeBytes(data);
        return keccak256(RLPWriter.writeList(raw));
    }
}
