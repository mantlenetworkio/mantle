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
import "../../libraries/MPT.sol";
import "../IVerifier.sol";
import "../libraries/VerificationContext.sol";
import "../libraries/OneStepProof.sol";
import "../libraries/Params.sol";
import "../libraries/GasTable.sol";
import "../libraries/VerifierHelper.sol";
import "../libraries/EVMTypesLib.sol";

contract StorageOpVerifier is IVerifier {
    using BytesLib for bytes;
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using VerificationContext for VerificationContext.Context;
    using OneStepProof for OneStepProof.CodeProof;
    using OneStepProof for EVMTypesLib.Account;

    function verifyOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        external
        pure
        override
        returns (bytes32)
    {
        return OneStepProof.hashStateProof(executeOneStepProof(ctx, currStateHash, encoded));
    }

    function executeOneStepProof(VerificationContext.Context memory ctx, bytes32 currStateHash, bytes calldata encoded)
        public
        pure
        returns (OneStepProof.StateProof memory)
    {
        uint64 offset = 0;
        // Decode state proof
        OneStepProof.StateProof memory stateProof;
        (offset, stateProof) = OneStepProof.decodeStateProof(ctx, encoded, offset);
        // Calculate the state hash from the submitted proof
        bytes32 stateHashFromProof;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, encoded.offset, offset)
            stateHashFromProof := keccak256(ptr, offset)
        }
        // Ensure the state proof is valid
        require(stateHashFromProof == currStateHash, "Bad State Proof");

        // Decode the code proof
        OneStepProof.CodeProof memory codeProof;
        (offset, codeProof) = OneStepProof.decodeCodeProof(encoded, offset);

        // Simulate EVM one-step execution based on the opcode
        uint8 opCode = stateProof.opCode;
        if (opCode == 0x31) {
            // BALANCE
            verifyOpBALANCE(offset, stateProof, encoded);
        } else if (opCode == 0x3b) {
            // EXTCODESIZE
            verifyOpEXTCODESIZE(offset, stateProof, encoded);
        } else if (opCode == 0x3c) {
            // EXTCODECOPY
            verifyOpEXTCODECOPY(offset, stateProof, encoded);
        } else if (opCode == 0x3f) {
            // EXTCODEHASH
            verifyOpEXTCODEHASH(offset, stateProof, encoded);
        } else if (opCode == 0x47) {
            // SELFBALANCE
            verifyOpSELFBALANCE(offset, stateProof, encoded);
        } else if (opCode == 0x54) {
            // SLOAD
            verifyOpSLOAD(offset, stateProof, encoded);
        } else if (opCode == 0x55) {
            // SSTORE
            verifyOpSSTORE(offset, stateProof, encoded);
        } else {
            revert("Unreachable");
        }

        // Obtain the opcode at new pc
        if (stateProof.depth > 0) {
            if (codeProof.size > uint256(stateProof.pc)) {
                stateProof.opCode = codeProof.getOpCodeAt(encoded, stateProof.pc);
            } else {
                stateProof.opCode = 0x00;
            }
        }
        // Return the state hash after one-step execution
        return stateProof;
    }

    function verifyOpBALANCE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_BALANCE;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
        }
        stateProof.gas -= cost;

        OneStepProof.RLPProof memory accountProof;
        (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes32 addressHash = keccak256(abi.encodePacked(address(uint160(stackProof.pops[0]))));
        bytes memory encodedAccount = MPT.extractProofValue(
            stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
        );
        EVMTypesLib.Account memory account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());

        stateProof.pc += 1;
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, account.balance));
    }

    function verifyOpEXTCODESIZE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_EXTCODE;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
        }
        stateProof.gas -= cost;

        OneStepProof.RLPProof memory accountProof;
        (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes32 addressHash = keccak256(abi.encodePacked(address(uint160(stackProof.pops[0]))));
        bytes memory encodedAccount = MPT.extractProofValue(
            stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
        );
        EVMTypesLib.Account memory account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());

        OneStepProof.CodeProof memory codeProof;
        (offset, codeProof) = OneStepProof.decodeCodeProof(encoded, offset);

        require(codeProof.hashCodeProof(encoded) == account.codeHash, "Bad Code Proof");

        stateProof.pc += 1;
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, codeProof.size));
    }

    function verifyOpEXTCODECOPY(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 4) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 4, offset, encoded);

        {
            uint64 cost =
                GasTable.gasExtCopy(stateProof.memSize, uint64(stackProof.pops[1]), uint64(stackProof.pops[3]));
            if (stateProof.gas < cost) {
                VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            }
            stateProof.gas -= cost;
        }

        EVMTypesLib.Account memory account;
        {
            OneStepProof.RLPProof memory accountProof;
            (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
            bytes32 addressHash = keccak256(abi.encodePacked(address(uint160(stackProof.pops[0]))));
            bytes memory encodedAccount = MPT.extractProofValue(
                stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
            );
            account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());
        }

        OneStepProof.CodeProof memory codeProof;
        (offset, codeProof) = OneStepProof.decodeCodeProof(encoded, offset);

        require(codeProof.hashCodeProof(encoded) == account.codeHash, "Bad Code Proof");
        {
            bytes memory readContent =
                codeProof.getCodeSlice(encoded, uint64(stackProof.pops[2]), uint64(stackProof.pops[3]));
            bytes memory writeContent;
            (offset, writeContent) = MemoryLib.decodeAndVerifyMemoryWriteProof(
                stateProof, encoded, offset, uint64(stackProof.pops[1]), uint64(stackProof.pops[3])
            );

            require(readContent.equal(writeContent), "Inconsistent Copy");
        }

        stateProof.pc += 1;
        stateProof.stackSize -= 4;
        stateProof.stackHash = stackProof.stackHashAfterPops;
    }

    function verifyOpEXTCODEHASH(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_EXTCODE;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
        }
        stateProof.gas -= cost;

        OneStepProof.RLPProof memory accountProof;
        (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes32 addressHash = keccak256(abi.encodePacked(address(uint160(stackProof.pops[0]))));
        bytes memory encodedAccount = MPT.extractProofValue(
            stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
        );
        EVMTypesLib.Account memory account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());

        stateProof.pc += 1;
        stateProof.stackHash = keccak256(abi.encodePacked(stackProof.stackHashAfterPops, account.codeHash));
    }

    function verifyOpSELFBALANCE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize >= Params.STACK_LIMIT - 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }

        uint64 cost = Params.G_LOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
        }
        stateProof.gas -= cost;

        OneStepProof.RLPProof memory accountProof;
        (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes32 addressHash = keccak256(abi.encodePacked(stateProof.contractAddress));
        bytes memory encodedAccount = MPT.extractProofValue(
            stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
        );
        EVMTypesLib.Account memory account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());

        stateProof.pc += 1;
        stateProof.stackSize += 1;
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, account.codeHash));
    }

    function verifyOpSLOAD(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {
        if (stateProof.stackSize < 1) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
            return;
        }
        OneStepProof.StackProof memory stackProof;
        (offset, stackProof) = VerifierHelper.decodeAndVerifyStackProof(stateProof, 1, offset, encoded);

        uint64 cost = Params.G_LOW;
        if (stateProof.gas < cost) {
            VerifierHelper.verifyRevertByError(offset, stateProof, encoded);
        }
        stateProof.gas -= cost;

        OneStepProof.RLPProof memory accountProof;
        (offset, accountProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes32 addressHash = keccak256(abi.encodePacked(stateProof.contractAddress));
        bytes memory encodedAccount = MPT.extractProofValue(
            stateProof.globalStateRoot, abi.encodePacked(addressHash), accountProof.proof.toList()
        );
        EVMTypesLib.Account memory account = EVMTypesLib.decodeAccount(encodedAccount.toRlpItem());

        OneStepProof.RLPProof memory storageProof;
        (offset, storageProof) = OneStepProof.decodeRLPProof(encoded, offset);
        bytes memory value = MPT.extractProofValue(
            account.storageRoot, abi.encodePacked(bytes32(stackProof.pops[0])), storageProof.proof.toList()
        );

        stateProof.pc += 1;
        stateProof.stackHash = keccak256(abi.encodePacked(stateProof.stackHash, value.toBytes32Pad(0)));
    }

    function verifyOpSSTORE(uint64 offset, OneStepProof.StateProof memory stateProof, bytes calldata encoded)
        internal
        pure
    {}
}
