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

import "./BytesLib.sol";

library DeserializationLib {
    function deserializeAddress(bytes memory data, uint256 startOffset) internal pure returns (uint256, address) {
        return (startOffset + 20, BytesLib.toAddress(data, startOffset));
    }

    function deserializeUint256(bytes memory data, uint256 startOffset) internal pure returns (uint256, uint256) {
        require(data.length >= startOffset && data.length - startOffset >= 32, "too short");
        return (startOffset + 32, BytesLib.toUint256(data, startOffset));
    }

    function deserializeBytes32(bytes memory data, uint256 startOffset) internal pure returns (uint256, bytes32) {
        require(data.length >= startOffset && data.length - startOffset >= 32, "too short");
        return (startOffset + 32, BytesLib.toBytes32(data, startOffset));
    }
}
