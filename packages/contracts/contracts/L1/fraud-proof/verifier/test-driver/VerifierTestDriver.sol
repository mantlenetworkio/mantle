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

import "hardhat/console.sol";

import "../IVerifier.sol";
import "../libraries/Params.sol";
import "../libraries/EVMTypesLib.sol";
import "../libraries/VerificationContext.sol";

contract VerifierTestDriver {
    IVerifier blockInitiationVerifier;
    IVerifier blockFinalizationVerifier;
    IVerifier interTxVerifier;
    IVerifier stackOpVerifier;
    IVerifier environmentalOpVerifier;
    IVerifier memoryOpVerifier;
    IVerifier storageOpVerifier;
    IVerifier callOpVerifier;
    IVerifier invalidOpVerifier;

    constructor(
        address _blockInitiationVerifier,
        address _blockFinalizationVerifier,
        address _interTxVerifier,
        address _stackOpVerifier,
        address _environmentalOpVerifier,
        address _memoryOpVerifier,
        address _storageOpVerifier,
        address _callOpVerifier,
        address _invalidOpVerifier
    ) {
        blockInitiationVerifier = IVerifier(_blockInitiationVerifier);
        blockFinalizationVerifier = IVerifier(_blockFinalizationVerifier);
        interTxVerifier = IVerifier(_interTxVerifier);
        stackOpVerifier = IVerifier(_stackOpVerifier);
        environmentalOpVerifier = IVerifier(_environmentalOpVerifier);
        memoryOpVerifier = IVerifier(_memoryOpVerifier);
        storageOpVerifier = IVerifier(_storageOpVerifier);
        callOpVerifier = IVerifier(_callOpVerifier);
        invalidOpVerifier = IVerifier(_invalidOpVerifier);
    }

    function getVerifier(uint8 verifier) private view returns (IVerifier) {
        if (verifier == Params.V_BLOCK_INIT) {
            return blockInitiationVerifier;
        } else if (verifier == Params.V_BLOCK_FINAL) {
            return blockFinalizationVerifier;
        } else if (verifier == Params.V_INTER_TX) {
            return interTxVerifier;
        } else if (verifier == Params.V_STACK_OP) {
            return stackOpVerifier;
        } else if (verifier == Params.V_ENVIRONMENTAL_OP) {
            return environmentalOpVerifier;
        } else if (verifier == Params.V_MEMORY_OP) {
            return memoryOpVerifier;
        } else if (verifier == Params.V_STORAGE_OP) {
            return storageOpVerifier;
        } else if (verifier == Params.V_CALL_OP) {
            return callOpVerifier;
        } else if (verifier == Params.V_INVALID_OP) {
            return invalidOpVerifier;
        } else {
            revert("unreachable");
        }
    }

    function verifyProof(
        address sequencerAddress,
        uint256 timestamp,
        uint256 number,
        address origin,
        bytes32 txHash,
        EVMTypesLib.Transaction memory transaction,
        uint8 verifier,
        bytes32 currStateHash,
        bytes calldata proof
    ) external returns (bytes32) {
        VerificationContext.Context memory ctx =
            VerificationContext.Context(sequencerAddress, timestamp, number, origin, transaction, bytes32(0), txHash);
        bytes32 res = getVerifier(verifier).verifyOneStepProof(ctx, currStateHash, proof);
        console.logBytes32(res);
        return res;
    }
}
