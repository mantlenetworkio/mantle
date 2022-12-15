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

package state

import (
	"encoding/binary"

	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/specularl2/specular/clients/geth/specular/merkletree"
)

type Memory struct {
	content []byte
	size    uint64
	tree    *merkletree.MerkleTree
}

func NewMemoryFromBytes(values []byte) *Memory {
	if len(values) == 0 {
		return &Memory{tree: merkletree.New([]*uint256.Int{})}
	}
	size := len(values)
	bytesToPad := 32 - size%32
	if bytesToPad != 32 {
		values = append(values, make([]byte, bytesToPad)...)
	}
	elements := make([]*uint256.Int, len(values)/32)
	for idx := range elements {
		elements[idx] = new(uint256.Int).SetBytes(values[idx*32 : (idx+1)*32])
	}
	return &Memory{
		content: values,
		size:    uint64(size),
		tree:    merkletree.New(elements),
	}
}

func MemoryFromEVMMemory(mem *vm.Memory) *Memory {
	return NewMemoryFromBytes(mem.Data())
}

func (m *Memory) Size() uint64 {
	return m.size
}

func (m *Memory) Data() []byte {
	return m.content[:m.size]
}

func (m *Memory) CellNum() uint64 {
	return m.tree.ElementCount()
}

func (m *Memory) Empty() bool {
	return m.size == 0
}

func (m *Memory) Cell(i uint64) *uint256.Int {
	return m.tree.Element(i)
}

func (m *Memory) Root() common.Hash {
	return common.Hash(m.tree.Root())
}

func (m *Memory) GetProof(indices []uint64) []common.Hash {
	var hashes []common.Hash
	for _, hash := range m.tree.GenerateProof(indices) {
		hashes = append(hashes, common.Hash(hash))
	}
	return hashes
}

func (m *Memory) GetAppendProof() []common.Hash {
	var hashes []common.Hash
	for _, hash := range m.tree.GenerateAppendProof() {
		hashes = append(hashes, common.Hash(hash))
	}
	return hashes
}

func (m *Memory) GetCombinedProof(indices []uint64) ([]common.Hash, error) {
	var hashes []common.Hash
	results, err := m.tree.GenerateCombinedProof(indices)
	if err != nil {
		return nil, err
	}
	for _, hash := range results {
		hashes = append(hashes, common.Hash(hash))
	}
	return hashes, nil
}

func (m *Memory) EncodeState() []byte {
	encoded := make([]byte, 8)
	binary.BigEndian.PutUint64(encoded, m.size)
	if m.size != 0 {
		encoded = append(encoded, m.Root().Bytes()...)
	}
	return encoded
}
