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

import "../../libraries/BitLen.sol";
import "./Params.sol";

library GasTable {
    using BitLen for uint256;

    function costMemory(uint64 mwords) internal pure returns (uint64) {
        return Params.G_MEMORY * mwords + mwords * mwords / 512;
    }

    function gasMemory(uint64 msize, uint64 offset, uint64 length) internal pure returns (uint64) {
        if (length == 0) {
            return 0;
        }
        uint64 newMwords = (offset + length + 31) / 32;
        uint64 mwords = (msize + 31) / 32;
        if (newMwords <= mwords) {
            return 0;
        }
        return costMemory(newMwords) - costMemory(mwords);
    }

    function gasExp(uint256 exponent) internal pure returns (uint64) {
        if (exponent == 0) {
            return Params.G_EXP;
        } else {
            return Params.G_EXP + Params.G_EXPBYTE * ((exponent.bitLen() + 7) / 8);
        }
    }

    function gasKeccak(uint64 msize, uint64 offset, uint64 length) internal pure returns (uint64) {
        return gasMemory(msize, offset, length) + Params.G_KECCAK + Params.G_KECCAKWORD * ((length + 31) / 32);
    }

    function gasCopy(uint64 msize, uint64 offset, uint64 length) internal pure returns (uint64) {
        return gasMemory(msize, offset, length) + Params.G_VERYLOW + Params.G_COPY * ((length + 31) / 32);
    }

    function gasExtCopy(uint64 msize, uint64 offset, uint64 length) internal pure returns (uint64) {
        return gasMemory(msize, offset, length) + Params.G_EXTCODE + Params.G_COPY * ((length + 31) / 32);
    }

    function gasLog(uint64 msize, uint64 offset, uint64 length, uint64 topicNum) internal pure returns (uint64) {
        return gasMemory(msize, offset, length) + Params.G_LOG + Params.G_LOGTOPIC * topicNum
            + Params.G_LOGDATA * ((length + 31) / 32);
    }
}
