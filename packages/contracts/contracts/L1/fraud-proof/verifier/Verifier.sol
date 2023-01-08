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

import "./IVerifier.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Verifier is IVerifier, Initializable {
    function initialize() public initializer {}

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function verifyOneStepProof(IVerificationContext, bytes32, bytes calldata)
        external
        pure
        override
        returns (bytes32 nextStateHash)
    {
        nextStateHash = 0x0;
    }
}
