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

library Params {
    // Verifier
    uint8 constant V_STACK_OP = 0;
    uint8 constant V_ENVIRONMENTAL_OP = 1;
    uint8 constant V_MEMORY_OP = 2;
    uint8 constant V_STORAGE_OP = 3;
    uint8 constant V_CALL_OP = 4;
    uint8 constant V_INVALID_OP = 5;
    uint8 constant V_INTER_TX = 6;
    uint8 constant V_BLOCK_INIT = 7;
    uint8 constant V_BLOCK_FINAL = 8;

    // Gas
    uint64 constant G_ZERO = 0;
    uint64 constant G_JUMPDEST = 1;
    uint64 constant G_BASE = 2;
    uint64 constant G_VERYLOW = 3;
    uint64 constant G_LOW = 5;
    uint64 constant G_MID = 8;
    uint64 constant G_HIGH = 10;
    uint64 constant G_EXTCODE = 700;
    uint64 constant G_BALANCE = 700;
    uint64 constant G_SLOAD = 800;
    uint64 constant G_SSET = 20000;
    uint64 constant G_SRESET = 5000;
    uint64 constant R_SCLEAR = 15000;
    uint64 constant R_SELFDESTRUCT = 24000;
    uint64 constant G_SELFDESTRUCT = 5000;
    uint64 constant G_CREATE = 32000;
    uint64 constant G_CODEDEPOSIT = 200;
    uint64 constant G_CALL = 700;
    uint64 constant G_CALLVALUE = 9000;
    uint64 constant G_CALLSTIPEND = 2300;
    uint64 constant G_NEWACCOUNT = 25000;
    uint64 constant G_EXP = 10;
    uint64 constant G_EXPBYTE = 50;
    uint64 constant G_MEMORY = 3;
    uint64 constant G_TXCREATE = 32000;
    uint64 constant G_TXDATAZERO = 4;
    uint64 constant G_TXDATANONZERO = 16;
    uint64 constant G_TRANSACTION = 21000;
    uint64 constant G_LOG = 375;
    uint64 constant G_LOGDATA = 8;
    uint64 constant G_LOGTOPIC = 375;
    uint64 constant G_KECCAK = 30;
    uint64 constant G_KECCAKWORD = 6;
    uint64 constant G_COPY = 3;
    uint64 constant G_BLOCKHASH = 20;
    uint64 constant G_QUADDIVISOR = 20;

    // Stack
    uint64 constant STACK_LIMIT = 1024;

    bytes32 constant EMPTY_TRIE_ROOT = 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421;
    bytes32 constant EMPTY_CODE_HASH = 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470;
    bytes32 constant EMPTY_UNCLE_HASH = 0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347;
    uint64 constant RECENT_BLOCK_HASHES_LENGTH = 256;
}
