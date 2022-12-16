// Copyright 2022, Specular contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proof

import (
	"encoding/binary"
	"fmt"

	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

type StackProof struct {
	Pops               []uint256.Int
	StackHashAfterPops common.Hash
}

func (s *StackProof) Encode() []byte {
	var proofLen int
	if len(s.Pops) == 0 {
		return []byte{}
	}
	proofLen = 32 * (len(s.Pops) + 1)
	encoded := make([]byte, proofLen)
	for idx, val := range s.Pops {
		valBytes := val.Bytes32()
		copy(encoded[32*idx:], valBytes[:])
	}
	copy(encoded[32*len(s.Pops):], s.StackHashAfterPops.Bytes())
	return encoded
}

type MemoryReadProof struct {
	Cells []uint256.Int
	Proof []common.Hash
}

func (m *MemoryReadProof) Encode() []byte {
	encoded := make([]byte, 32*len(m.Cells)+1+32*len(m.Proof))
	for idx, cell := range m.Cells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx:], cellBytes[:])
	}
	encoded[32*len(m.Cells)] = byte(len(m.Proof))
	encodedOffset := 32*len(m.Cells) + 1
	for _, hash := range m.Proof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

type MemoryWriteProof struct {
	Cells       []uint256.Int
	UpdateCells []uint256.Int
	Proof       []common.Hash
}

func (m *MemoryWriteProof) Encode() []byte {
	encoded := make([]byte, 32*len(m.Cells)+32*len(m.UpdateCells)+1+32*len(m.Proof))
	for idx, cell := range m.Cells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx:], cellBytes[:])
	}
	encodedOffset := 32 * len(m.Cells)
	for idx, cell := range m.UpdateCells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx+encodedOffset:], cellBytes[:])
	}
	encoded[32*(len(m.Cells)+len(m.UpdateCells))] = byte(len(m.Proof))
	encodedOffset += 32*len(m.UpdateCells) + 1
	for _, hash := range m.Proof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

type MemoryAppendProof struct {
	AppendCells []uint256.Int
	Proof       []common.Hash
}

func (m *MemoryAppendProof) Encode() []byte {
	encoded := make([]byte, 32*len(m.AppendCells)+1+32*len(m.Proof))
	for idx, cell := range m.AppendCells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx:], cellBytes[:])
	}
	encoded[32*len(m.AppendCells)] = byte(len(m.Proof))
	encodedOffset := 32*len(m.AppendCells) + 1
	for _, hash := range m.Proof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

type MemoryCombinedReadProof struct {
	Cells       []uint256.Int
	AppendCells []uint256.Int
	Proof       []common.Hash
}

func (m *MemoryCombinedReadProof) Encode() []byte {
	encoded := make([]byte, 32*len(m.Cells)+32*len(m.AppendCells)+1+32*len(m.Proof))
	for idx, cell := range m.Cells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx:], cellBytes[:])
	}
	encodedOffset := 32 * len(m.Cells)
	for idx, cell := range m.AppendCells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx+encodedOffset:], cellBytes[:])
	}
	encoded[32*(len(m.Cells)+len(m.AppendCells))] = byte(len(m.Proof))
	encodedOffset += 32*len(m.AppendCells) + 1
	for _, hash := range m.Proof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

type MemoryCombinedWriteProof struct {
	Cells        []uint256.Int
	UpdatedCells []uint256.Int
	AppendCells  []uint256.Int
	Proof        []common.Hash
}

func (m *MemoryCombinedWriteProof) Encode() []byte {
	encoded := make([]byte, 32*len(m.Cells)+32*len(m.UpdatedCells)+32*len(m.AppendCells)+1+32*len(m.Proof))
	for idx, cell := range m.Cells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx:], cellBytes[:])
	}
	encodedOffset := 32 * len(m.Cells)
	for idx, cell := range m.UpdatedCells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx+encodedOffset:], cellBytes[:])
	}
	encodedOffset += 32 * len(m.UpdatedCells)
	for idx, cell := range m.AppendCells {
		cellBytes := cell.Bytes32()
		copy(encoded[32*idx+encodedOffset:], cellBytes[:])
	}
	encoded[32*(len(m.Cells)+len(m.UpdatedCells)+len(m.AppendCells))] = byte(len(m.Proof))
	encodedOffset += 32*len(m.AppendCells) + 1
	for _, hash := range m.Proof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

func generateMemoryReadProof(memory *state.Memory, offset uint64, size uint64) (Proof, error) {
	startCell := offset / 32
	cellNum := calcCellNum(offset, size)
	if memory.CellNum() <= startCell {
		// The start position pasts the end of the memory
		// AppendProof
		pf := memory.GetAppendProof()
		cellNum += startCell - memory.CellNum()
		return &MemoryAppendProof{
			AppendCells: make([]uint256.Int, cellNum), // All empty
			Proof:       pf,
		}, nil
	}
	if memory.CellNum() >= startCell+cellNum {
		// The end position is within the memory
		// ReadProof
		indices := make([]uint64, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			indices[i] = startCell + i
		}
		pf := memory.GetProof(indices)
		cells := make([]uint256.Int, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			cells[i] = *memory.Cell(startCell + i)
		}
		return &MemoryReadProof{
			Cells: cells,
			Proof: pf,
		}, nil
	}
	// The start position is within the memory, but the end position is not
	// CombinedReadProof
	indices := make([]uint64, memory.CellNum()-startCell)
	for i := startCell; i < memory.CellNum(); i++ {
		indices[i-startCell] = i
	}
	pf, err := memory.GetCombinedProof(indices)
	if err != nil {
		return nil, err
	}
	cells := make([]uint256.Int, memory.CellNum()-startCell)
	for i := startCell; i < memory.CellNum(); i++ {
		cells[i-startCell] = *memory.Cell(i)
	}
	appendCells := make([]uint256.Int, cellNum-(memory.CellNum()-startCell))
	return &MemoryCombinedReadProof{
		Cells:       cells,
		AppendCells: appendCells,
		Proof:       pf,
	}, nil
}

func generateMemoryReadProofNoAppend(memory *state.Memory, offset uint64, size uint64) (Proof, error) {
	startCell := offset / 32
	cellNum := calcCellNum(offset, size)
	if memory.CellNum() <= startCell {
		// The start position pasts the end of the memory
		return nil, nil
	}
	if memory.CellNum() >= startCell+cellNum {
		// The end position is within the memory
		indices := make([]uint64, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			indices[i] = startCell + i
		}
		pf := memory.GetProof(indices)
		cells := make([]uint256.Int, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			cells[i] = *memory.Cell(startCell + i)
		}
		return &MemoryReadProof{
			Cells: cells,
			Proof: pf,
		}, nil
	}
	// The start position is within the memory, but the end position is not
	indices := make([]uint64, memory.CellNum()-startCell)
	for i := startCell; i < memory.CellNum(); i++ {
		indices[i-startCell] = i
	}
	pf, err := memory.GetCombinedProof(indices)
	if err != nil {
		return nil, err
	}
	cells := make([]uint256.Int, memory.CellNum()-startCell)
	for i := startCell; i < memory.CellNum(); i++ {
		cells[i-startCell] = *memory.Cell(i)
	}
	return &MemoryReadProof{
		Cells: cells,
		Proof: pf,
	}, nil
}

func generateMemoryWriteProof(memory, memoryAfter *state.Memory, offset uint64, size uint64) (Proof, error) {
	startCell := offset / 32
	cellNum := calcCellNum(offset, size)
	if memory.CellNum() <= startCell {
		// The start position pasts the end of the memory
		// AppendProof
		pf := memory.GetAppendProof()
		cellNum += startCell - memory.CellNum()
		appendCells := make([]uint256.Int, cellNum)
		for i := startCell - memory.CellNum(); i < cellNum; i++ {
			appendCells[i] = *memoryAfter.Cell(i + memory.CellNum())
		}
		return &MemoryAppendProof{
			AppendCells: make([]uint256.Int, cellNum), // All empty
			Proof:       pf,
		}, nil
	}
	if memory.CellNum() >= startCell+cellNum {
		// The end position is within the memory
		// ReadProof
		indices := make([]uint64, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			indices[i] = startCell + i
		}
		pf := memory.GetProof(indices)
		cells := make([]uint256.Int, cellNum)
		updatecCells := make([]uint256.Int, cellNum)
		for i := uint64(0); i < cellNum; i++ {
			cells[i] = *memory.Cell(startCell + i)
			updatecCells[i] = *memoryAfter.Cell(startCell + i)
		}
		return &MemoryWriteProof{
			Cells:       cells,
			UpdateCells: updatecCells,
			Proof:       pf,
		}, nil
	}
	// The start position is within the memory, but the end position is not
	// CombinedReadProof
	indices := make([]uint64, memory.CellNum()-startCell)
	for i := startCell; i < memory.CellNum(); i++ {
		indices[i-startCell] = i
	}
	pf, err := memory.GetCombinedProof(indices)
	if err != nil {
		return nil, err
	}
	cells := make([]uint256.Int, memory.CellNum()-startCell)
	updatecCells := make([]uint256.Int, cellNum)
	for i := startCell; i < memory.CellNum(); i++ {
		cells[i-startCell] = *memory.Cell(i)
		updatecCells[i-startCell] = *memoryAfter.Cell(i)
	}
	appendCells := make([]uint256.Int, cellNum-(memory.CellNum()-startCell))
	for i := memory.CellNum() - startCell; i < cellNum; i++ {
		appendCells[i-(memory.CellNum()-startCell)] = *memoryAfter.Cell(i + startCell)
	}
	return &MemoryCombinedWriteProof{
		Cells:        cells,
		UpdatedCells: updatecCells,
		AppendCells:  appendCells,
		Proof:        pf,
	}, nil
}

type MPTProof struct {
	Proof [][]byte
}

func (m *MPTProof) Encode() []byte {
	data, err := rlp.EncodeToBytes(m.Proof)
	if err != nil {
		log.Error(fmt.Sprint(err))
		panic(err)
	}
	dataLen := uint64(len(data))
	encoded := make([]byte, 8+dataLen)
	binary.BigEndian.PutUint64(encoded, dataLen)
	copy(encoded[8:], data)
	return encoded
}

type CodeProof struct {
	Content []byte
}

func (c *CodeProof) Encode() []byte {
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(len(c.Content)))
	encoded := make([]byte, 8+len(c.Content))
	copy(encoded, length)
	copy(encoded[8:], c.Content)
	return encoded
}

type BlockHashProof struct {
	BlockHash common.Hash
}

func (b *BlockHashProof) Encode() []byte {
	return b.BlockHash.Bytes()
}

type BlockHashMerkleProof struct {
	MerklePath  uint64
	MerkleProof []common.Hash
}

func GetBlockHashMerkleProof(tree *state.BlockHashTree, blockNum uint64) (*BlockHashMerkleProof, error) {
	proof, path, err := tree.GetProof(blockNum)
	if err != nil {
		return nil, err
	}
	return &BlockHashMerkleProof{
		MerklePath:  path,
		MerkleProof: proof,
	}, nil
}

func (p *BlockHashMerkleProof) Encode() []byte {
	encoded := make([]byte, 8+1+32*len(p.MerkleProof))
	path := make([]byte, 8)
	binary.BigEndian.PutUint64(path, p.MerklePath)
	copy(encoded, path)
	encoded[8] = byte(len(p.MerkleProof))
	encodedOffset := 9
	for _, hash := range p.MerkleProof {
		copy(encoded[encodedOffset:], hash.Bytes())
		encodedOffset += 32
	}
	return encoded
}

type LogProof struct {
	AccumulateHash common.Hash
	Bloom          types.Bloom
}

func LogProofFromLogSeries(logSeries *state.LogSeries) *LogProof {
	return &LogProof{
		AccumulateHash: logSeries.AccumulateHash,
		Bloom:          logSeries.Bloom,
	}
}

func (p *LogProof) Encode() []byte {
	encoded := make([]byte, 32+256)
	copy(encoded, p.AccumulateHash.Bytes())
	copy(encoded[32:], p.Bloom.Bytes())
	return encoded
}
