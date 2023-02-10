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

// For EXP opcode gas calculation
// Based on go packages `math/bits` and `holiman/uint256`
// TODO: seems for EXP we only need bytelen
library BitLen {
    uint64 constant UINT16_MINIMUM = 1 << 8;
    uint64 constant UINT32_MINIMUM = 1 << 16;
    uint64 constant UINT64_MINIMUM = 1 << 32;
    uint256 constant UINT128_MINIMUM = 1 << 64;
    uint256 constant UINT192_MINIMUM = 1 << 128;
    uint256 constant UINT256_MINIMUM = 1 << 192;

    function len8(uint8 x) internal pure returns (uint64 n) {
        if (x == 0) {
            return 0;
        } else if (x == 1) {
            return 1;
        } else if (x <= 3) {
            return 2;
        } else if (x <= 7) {
            return 3;
        } else if (x <= 15) {
            return 4;
        } else if (x <= 31) {
            return 5;
        } else if (x <= 63) {
            return 6;
        } else if (x <= 127) {
            return 7;
        }
        return 8;
    }

    function len64(uint64 x) internal pure returns (uint64 n) {
        if (x >= UINT64_MINIMUM) {
            x >>= 32;
            n = 32;
        }
        if (x >= UINT32_MINIMUM) {
            x >>= 16;
            n += 16;
        }
        if (x >= UINT16_MINIMUM) {
            x >>= 8;
            n += 8;
        }
        return n + len8(uint8(x));
    }

    function bitLen(uint256 x) internal pure returns (uint64) {
        if (x >= UINT256_MINIMUM) {
            return 192 + len64(uint64(x >> 192));
        } else if (x >= UINT192_MINIMUM) {
            return 128 + len64(uint64(x >> 128));
        } else if (x >= UINT128_MINIMUM) {
            return 64 + len64(uint64(x >> 64));
        }
        return len64(uint64(x));
    }
}
