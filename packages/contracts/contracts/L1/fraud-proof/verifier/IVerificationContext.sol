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

interface IVerificationContext {
    // Block
    function getBlockHash(uint8 number) external view returns (bytes32);
    function getCoinbase() external view returns (address);
    function getTimestamp() external view returns (uint256);
    function getBlockNumber() external view returns (uint256);
    function getDifficulty() external view returns (uint256);
    function getGasLimit() external view returns (uint64);
    function getChainID() external view returns (uint256);
    function getBaseFee() external view returns (uint256);

    // Transaction
    enum TxnType {
        TRANSFER,
        CREATE,
        CALL
    }

    function getStateRoot() external view returns (bytes32);
    function getEndStateRoot() external view returns (bytes32);
    function getOrigin() external view returns (address);
    function getRecipient() external view returns (address);
    function getTxnType() external view returns (TxnType);
    function getValue() external view returns (uint256);
    function getGas() external view returns (uint256);
    function getGasPrice() external view returns (uint256);
    function getInput() external view returns (bytes memory);
    function getInputSize() external view returns (uint64);
    function getInputRoot() external view returns (bytes32);
    // Only called when txn type is CREATE
    function getCodeMerkleFromInput() external view returns (bytes32);
}
