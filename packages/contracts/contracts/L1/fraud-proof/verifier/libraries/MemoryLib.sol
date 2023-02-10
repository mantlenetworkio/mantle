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

import "../../libraries/DeserializationLib.sol";
import "../../libraries/MerkleLib.sol";
import "../../libraries/BytesLib.sol";
import "./OneStepProof.sol";

library MemoryLib {
    using BytesLib for bytes;

    function calcCellNum(uint64 offset, uint64 length) internal pure returns (uint64) {
        return (offset + length + 31) / 32 - offset / 32;
    }

    function getMemoryRoot(bytes memory content) internal pure returns (bytes32) {
        uint64 cellNum = MemoryLib.calcCellNum(0, uint64(content.length));
        bytes32[] memory elements = new bytes32[](cellNum);
        for (uint256 i = 0; i < cellNum - 1; i++) {
            elements[i] = content.toBytes32(i * 32);
        }
        elements[cellNum - 1] = content.toBytes32Pad((cellNum - 1) * 32);
        return MerkleLib.create_from_many(elements);
    }

    function decodeAndVerifyMemoryReadProof(
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded,
        uint64 offset,
        uint64 memoryOffset,
        uint64 memoryReadLength
    ) internal pure returns (uint64, bytes memory) {
        if (stateProof.memSize == 0 || memoryReadLength == 0) {
            return (offset, new bytes(memoryReadLength));
        }
        uint64 startCell = memoryOffset / 32;
        uint64 cellNum = calcCellNum(memoryOffset, memoryReadLength);
        uint64 memoryCell = calcCellNum(0, stateProof.memSize);
        OneStepProof.MemoryMerkleProof memory merkleProof;
        {
            if (memoryCell <= startCell) {
                cellNum += startCell - memoryCell;
                OneStepProof.MemoryAppendProof memory appendProof;
                (offset, appendProof) = OneStepProof.decodeMemoryAppendProof(encoded, offset, cellNum);
                (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
                stateProof.memRoot =
                    MerkleLib.try_append_many(stateProof.memRoot, appendProof.appendCells, merkleProof.proof);
                if (memoryOffset + memoryReadLength > stateProof.memSize) {
                    stateProof.memSize = (memoryOffset + memoryReadLength + 31) / 32 * 32; // Expand by words
                }
                bytes memory readContent = new bytes(memoryReadLength);
                return (offset, readContent);
            }
        }
        {
            if (memoryCell >= startCell + cellNum) {
                OneStepProof.MemoryReadProof memory readProof;
                (offset, readProof) = OneStepProof.decodeMemoryReadProof(encoded, offset, cellNum);
                (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
                if (cellNum == 1) {
                    MerkleLib.element_exists(stateProof.memRoot, startCell, readProof.cells[0], merkleProof.proof);
                    require(
                        MerkleLib.element_exists(stateProof.memRoot, startCell, readProof.cells[0], merkleProof.proof),
                        "IMP"
                    );
                } else {
                    {
                        // avoid stack too deep
                        uint256[] memory indices = MerkleLib.get_indices(readProof.cells, merkleProof.proof);
                        for (uint64 i = 0; i < cellNum; i++) {
                            require(indices[i] == startCell + i, "IMP");
                        }
                    }
                    MerkleLib.elements_exist(stateProof.memRoot, readProof.cells, merkleProof.proof);
                    require(MerkleLib.elements_exist(stateProof.memRoot, readProof.cells, merkleProof.proof), "IMP");
                }
                bytes memory readContent = abi.encodePacked(readProof.cells).slice(memoryOffset % 32, memoryReadLength);
                return (offset, readContent);
            }
        }
        uint64 existCellNum = memoryCell - startCell;
        OneStepProof.MemoryCombinedReadProof memory combinedReadProof;
        (offset, combinedReadProof) =
            OneStepProof.decodeMemoryCombinedReadProof(encoded, offset, existCellNum, cellNum - existCellNum);
        (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
        if (existCellNum == 1) {
            stateProof.memRoot = MerkleLib.try_append_many_using_one(
                stateProof.memRoot,
                startCell,
                combinedReadProof.cells[0],
                combinedReadProof.appendCells,
                merkleProof.proof
            );
        } else {
            {
                // avoid stack too deep
                uint256[] memory indices = MerkleLib.get_indices(combinedReadProof.cells, merkleProof.proof);
                for (uint64 i = 0; i < existCellNum; i++) {
                    require(indices[i] == startCell + i, "IMP");
                }
            }
            stateProof.memRoot = MerkleLib.try_append_many_using_many(
                stateProof.memRoot, combinedReadProof.cells, combinedReadProof.appendCells, merkleProof.proof
            );
        }
        if (memoryOffset + memoryReadLength > stateProof.memSize) {
            stateProof.memSize = (memoryOffset + memoryReadLength + 31) / 32 * 32; // Expand by words
        }
        bytes memory readContent = abi.encodePacked(combinedReadProof.cells, combinedReadProof.appendCells).slice(
            memoryOffset % 32, memoryReadLength
        );
        return (offset, readContent);
    }

    function decodeAndVerifyMemoryLikeReadProofNoAppend(
        bytes32 memoryLikeRoot,
        uint64 memoryLikeSize,
        bytes calldata encoded,
        uint64 offset,
        uint64 memoryLikeOffset,
        uint64 memoryLikeReadLength
    ) internal pure returns (uint64, bytes memory) {
        if (memoryLikeSize == 0 || memoryLikeReadLength == 0) {
            return (offset, new bytes(memoryLikeReadLength));
        }
        uint64 startCell = memoryLikeOffset / 32;
        uint64 cellNum = calcCellNum(memoryLikeOffset, memoryLikeReadLength);
        uint64 memoryCell = calcCellNum(0, memoryLikeSize);
        {
            if (memoryCell <= startCell) {
                bytes memory readContent;
                readContent = new bytes(memoryLikeReadLength);
                return (offset, readContent);
            }
        }
        {
            if (memoryCell >= startCell + cellNum) {
                bytes memory readContent;
                OneStepProof.MemoryReadProof memory readProof;
                OneStepProof.MemoryMerkleProof memory merkleProof;
                (offset, readProof) = OneStepProof.decodeMemoryReadProof(encoded, offset, cellNum);
                (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
                if (cellNum == 1) {
                    MerkleLib.element_exists(memoryLikeRoot, startCell, readProof.cells[0], merkleProof.proof);
                    require(
                        MerkleLib.element_exists(memoryLikeRoot, startCell, readProof.cells[0], merkleProof.proof),
                        "IMP"
                    );
                } else {
                    {
                        uint256[] memory indices = MerkleLib.get_indices(readProof.cells, merkleProof.proof);
                        for (uint64 i = 0; i < cellNum; i++) {
                            require(indices[i] == startCell + i, "IMP2");
                        }
                    }
                    MerkleLib.elements_exist(memoryLikeRoot, readProof.cells, merkleProof.proof);
                    require(MerkleLib.elements_exist(memoryLikeRoot, readProof.cells, merkleProof.proof), "IMP");
                }
                readContent = abi.encodePacked(readProof.cells).slice(memoryLikeOffset % 32, memoryLikeReadLength);
                return (offset, readContent);
            }
        }
        uint64 existCellNum = memoryCell - startCell;
        OneStepProof.MemoryReadProof memory readProof;
        OneStepProof.MemoryMerkleProof memory merkleProof;
        (offset, readProof) = OneStepProof.decodeMemoryReadProof(encoded, offset, existCellNum);
        (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
        if (existCellNum == 1) {
            MerkleLib.element_exists(memoryLikeRoot, startCell, readProof.cells[0], merkleProof.proof);
            require(MerkleLib.element_exists(memoryLikeRoot, startCell, readProof.cells[0], merkleProof.proof), "IMP");
        } else {
            {
                uint256[] memory indices = MerkleLib.get_indices(readProof.cells, merkleProof.proof);
                for (uint64 i = 0; i < cellNum; i++) {
                    require(indices[i] == startCell + i, "IMP");
                }
            }
            MerkleLib.elements_exist(memoryLikeRoot, readProof.cells, merkleProof.proof);
            require(MerkleLib.elements_exist(memoryLikeRoot, readProof.cells, merkleProof.proof), "IMP");
        }
        bytes memory padding = new bytes(32 * (cellNum - existCellNum));
        bytes memory readContent;
        readContent = abi.encodePacked(readProof.cells, padding).slice(memoryLikeOffset % 32, memoryLikeReadLength);
        return (offset, readContent);
    }

    function decodeAndVerifyMemoryWriteProof(
        OneStepProof.StateProof memory stateProof,
        bytes calldata encoded,
        uint64 offset,
        uint64 memoryOffset,
        uint64 memoryWriteLength
    ) internal pure returns (uint64, bytes memory) {
        if (memoryWriteLength == 0) {
            return (offset, new bytes(0));
        }
        if (stateProof.memSize == 0) {
            // Don't call decodeMemoryWriteProof if memory is empty
            // Instead, update memory root and size directly
            revert();
        }
        uint64 startCell = memoryOffset / 32;
        uint64 cellNum = calcCellNum(memoryOffset, memoryWriteLength);
        uint64 memoryCell = calcCellNum(0, stateProof.memSize);
        OneStepProof.MemoryMerkleProof memory merkleProof;

        {
            if (memoryCell <= startCell) {
                cellNum += startCell - memoryCell;
                OneStepProof.MemoryAppendProof memory appendProof;
                (offset, appendProof) = OneStepProof.decodeMemoryAppendProof(encoded, offset, cellNum);
                (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
                if (cellNum == 1) {
                    stateProof.memRoot =
                        MerkleLib.try_append_one(stateProof.memRoot, appendProof.appendCells[0], merkleProof.proof);
                } else {
                    stateProof.memRoot =
                        MerkleLib.try_append_many(stateProof.memRoot, appendProof.appendCells, merkleProof.proof);
                }
                if (memoryOffset + memoryWriteLength > stateProof.memSize) {
                    stateProof.memSize = (memoryOffset + memoryWriteLength + 31) / 32 * 32; // Expand by words
                }
                bytes memory writeContent =
                    abi.encodePacked(appendProof.appendCells).slice(memoryOffset % 32, memoryWriteLength);
                return (offset, writeContent);
            }
        }
        {
            if (memoryCell >= startCell + cellNum) {
                OneStepProof.MemoryWriteProof memory writeProof;
                (offset, writeProof) = OneStepProof.decodeMemoryWriteProof(encoded, offset, cellNum);
                (offset, merkleProof) = OneStepProof.decodeMemoryMerkleProof(encoded, offset);
                if (cellNum == 1) {
                    stateProof.memRoot = MerkleLib.try_update_one(
                        stateProof.memRoot,
                        startCell,
                        writeProof.cells[0],
                        writeProof.updatedCells[0],
                        merkleProof.proof
                    );
                } else {
                    {
                        // Avoid stack too deep
                        uint256[] memory indices = MerkleLib.get_indices(writeProof.cells, merkleProof.proof);
                        for (uint64 i = 0; i < cellNum; i++) {
                            require(indices[i] == startCell + i, "IMP");
                        }
                    }
                    stateProof.memRoot = MerkleLib.try_update_many(
                        stateProof.memRoot, writeProof.cells, writeProof.updatedCells, merkleProof.proof
                    );
                }
                bytes memory writeContent =
                    abi.encodePacked(writeProof.updatedCells).slice(memoryOffset % 32, memoryWriteLength);
                return (offset, writeContent);
            }
        }
        uint64 existCellNum = memoryCell - startCell;
        OneStepProof.MemoryCombinedWriteProof memory combinedWriteProof;
        (offset, combinedWriteProof) =
            OneStepProof.decodeMemoryCombinedWriteProof(encoded, offset, existCellNum, cellNum - existCellNum);
        if (cellNum == 1) {
            stateProof.memRoot = MerkleLib.try_update_one_and_append_many(
                stateProof.memRoot,
                startCell,
                combinedWriteProof.cells[0],
                combinedWriteProof.updatedCells[0],
                combinedWriteProof.appendCells,
                merkleProof.proof
            );
        } else {
            {
                // avoid stack too deep
                uint256[] memory indices = MerkleLib.get_indices(combinedWriteProof.cells, merkleProof.proof);
                for (uint64 i = 0; i < cellNum; i++) {
                    require(indices[i] == startCell + i, "IMP");
                }
            }
            stateProof.memRoot = MerkleLib.try_update_many_and_append_many(
                stateProof.memRoot,
                combinedWriteProof.cells,
                combinedWriteProof.updatedCells,
                combinedWriteProof.appendCells,
                merkleProof.proof
            );
        }
        if (memoryOffset + memoryWriteLength > stateProof.memSize) {
            stateProof.memSize = (memoryOffset + memoryWriteLength + 31) / 32 * 32; // Expand by words
        }
        bytes memory writeContent = abi.encodePacked(combinedWriteProof.updatedCells, combinedWriteProof.appendCells)
            .slice(memoryOffset % 32, memoryWriteLength);
        return (offset, writeContent);
    }
}
