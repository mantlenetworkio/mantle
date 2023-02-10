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

library BloomLib {
    using BytesLib for bytes;

    struct Bloom {
        bytes32[8] data;
    }

    function emptyBloom() internal pure returns (Bloom memory b) {
        return b;
    }

    function decodeBloom(bytes calldata encoded, uint64 offset) internal pure returns (Bloom memory) {
        Bloom memory bloom;
        for (uint256 i = 0; i < 8; i++) {
            bloom.data[i] = encoded.toBytes32(offset);
            offset += 32;
        }
        return bloom;
    }

    function addHash(Bloom memory bloom, bytes32 h) internal pure {
        uint16 i1 = 255 - (uint16(uint256(h) >> 240) & 0x7ff) >> 3;
        uint8 v1 = uint8(1 << (uint8(h[1]) & 0x7));
        bloom.data[i1 >> 5] = bytes32(uint256(bloom.data[i1 >> 5]) | (uint256(v1) << 8 * (31 - (i1 & 0x1f))));
        uint16 i2 = 255 - (uint16(uint256(h) >> 224) & 0x7ff) >> 3;
        uint8 v2 = uint8(1 << (uint8(h[3]) & 0x7));
        bloom.data[i2 >> 5] = bytes32(uint256(bloom.data[i2 >> 5]) | (uint256(v2) << 8 * (31 - (i2 & 0x1f))));
        uint16 i3 = 255 - (uint16(uint256(h) >> 208) & 0x7ff) >> 3;
        uint8 v3 = uint8(1 << (uint8(h[5]) & 0x7));
        bloom.data[i3 >> 5] = bytes32(uint256(bloom.data[i3 >> 5]) | (uint256(v3) << 8 * (31 - (i3 & 0x1f))));
    }

    function add(Bloom memory bloom, bytes memory data) internal pure {
        bytes32 h;
        assembly {
            h := keccak256(add(data, 0x20), mload(data))
        }
        addHash(bloom, h);
    }

    function add(Bloom memory bloom, address data) internal pure {
        bytes32 h = keccak256(abi.encodePacked(data));
        addHash(bloom, h);
    }

    function add(Bloom memory bloom, bytes32 data) internal pure {
        bytes32 h = keccak256(abi.encodePacked(data));
        addHash(bloom, h);
    }
}
