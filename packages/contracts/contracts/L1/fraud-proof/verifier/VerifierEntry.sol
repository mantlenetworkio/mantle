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

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./libraries/VerificationContext.sol";
import "./libraries/Params.sol";
import "./IVerifier.sol";
import "./IVerifierEntry.sol";

contract VerifierEntry is IVerifierEntry, Initializable, OwnableUpgradeable {
    IVerifier blockInitiationVerifier;
    IVerifier blockFinalizationVerifier;
    IVerifier interTxVerifier;
    IVerifier stackOpVerifier;
    IVerifier environmentalOpVerifier;
    IVerifier memoryOpVerifier;
    IVerifier storageOpVerifier;
    IVerifier callOpVerifier;
    IVerifier invalidOpVerifier;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function setVerifier(uint8 verifier, IVerifier impl) external onlyOwner {
        if (verifier == Params.V_BLOCK_INIT) {
            blockInitiationVerifier = impl;
        } else if (verifier == Params.V_BLOCK_FINAL) {
            blockFinalizationVerifier = impl;
        } else if (verifier == Params.V_INTER_TX) {
            interTxVerifier = impl;
        } else if (verifier == Params.V_STACK_OP) {
            stackOpVerifier = impl;
        } else if (verifier == Params.V_ENVIRONMENTAL_OP) {
            environmentalOpVerifier = impl;
        } else if (verifier == Params.V_MEMORY_OP) {
            memoryOpVerifier = impl;
        } else if (verifier == Params.V_STORAGE_OP) {
            storageOpVerifier = impl;
        } else if (verifier == Params.V_CALL_OP) {
            callOpVerifier = impl;
        } else if (verifier == Params.V_INVALID_OP) {
            invalidOpVerifier = impl;
        } else {
            revert("unreachable");
        }
    }

    function verifyOneStepProof(
        VerificationContext.Context memory ctx,
        uint8 verifier,
        bytes32 currStateHash,
        bytes calldata encoded
    ) external view override returns (bytes32) {
        IVerifier impl;
        if (verifier == Params.V_BLOCK_INIT) {
            impl = blockInitiationVerifier;
        } else if (verifier == Params.V_BLOCK_FINAL) {
            impl = blockFinalizationVerifier;
        } else if (verifier == Params.V_INTER_TX) {
            impl = interTxVerifier;
        } else if (verifier == Params.V_STACK_OP) {
            impl = stackOpVerifier;
        } else if (verifier == Params.V_ENVIRONMENTAL_OP) {
            impl = environmentalOpVerifier;
        } else if (verifier == Params.V_MEMORY_OP) {
            impl = memoryOpVerifier;
        } else if (verifier == Params.V_STORAGE_OP) {
            impl = storageOpVerifier;
        } else if (verifier == Params.V_CALL_OP) {
            impl = callOpVerifier;
        } else if (verifier == Params.V_INVALID_OP) {
            impl = invalidOpVerifier;
        } else {
            revert("unreachable");
        }
        return impl.verifyOneStepProof(ctx, currStateHash, encoded);
    }
}
