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
	basic_merkletree "github.com/cbergoon/merkletree"
	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
)

const RECENT_BLOCK_HASHES_LENGTH = 256

type BlockHash common.Hash
type BlockHashes [RECENT_BLOCK_HASHES_LENGTH]BlockHash

func (b BlockHash) CalculateHash() ([]byte, error) {
	return b[:], nil
}

func (b BlockHash) Equals(other basic_merkletree.Content) (bool, error) {
	return b == other.(BlockHash), nil
}

type BlockHashTree struct {
	Hashes BlockHashes
	tree   *basic_merkletree.MerkleTree
}

func NewBlockHashTree(hashes *BlockHashes) (*BlockHashTree, error) {
	contents := make([]basic_merkletree.Content, RECENT_BLOCK_HASHES_LENGTH)
	for i, h := range hashes {
		contents[i] = h
	}
	tree, err := basic_merkletree.NewTreeWithHashStrategy(contents, newHash)
	if err != nil {
		return nil, err
	}
	return &BlockHashTree{
		Hashes: *hashes,
		tree:   tree,
	}, nil
}

func BlockHashTreeFromBlockContext(blockCtx *vm.Context) (*BlockHashTree, error) {
	// Get BlockHashes of block [current-255, current]
	blockHashes := BlockHashes{}
	currentBlockNum := blockCtx.BlockNumber.Uint64()
	start := uint64(0)
	if currentBlockNum >= RECENT_BLOCK_HASHES_LENGTH {
		start = currentBlockNum - RECENT_BLOCK_HASHES_LENGTH + 1
	}
	for i := start; i <= currentBlockNum; i++ {
		blockHashes[(currentBlockNum-i)%RECENT_BLOCK_HASHES_LENGTH] = BlockHash(blockCtx.GetHash(currentBlockNum - i))
	}
	tree, err := NewBlockHashTree(&blockHashes)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

func (b *BlockHashTree) Copy() *BlockHashTree {
	tree, _ := NewBlockHashTree(&b.Hashes)
	return tree
}

func (b *BlockHashTree) Root() common.Hash {
	return common.BytesToHash(b.tree.MerkleRoot())
}

func (b *BlockHashTree) EncodeState() []byte {
	return b.tree.MerkleRoot()
}

func (b *BlockHashTree) SetBlockHash(number *uint256.Int, hash common.Hash) error {
	index := number.Mod(number, uint256.NewInt(RECENT_BLOCK_HASHES_LENGTH)).Uint64()
	b.Hashes[index] = BlockHash(hash)
	contents := make([]basic_merkletree.Content, len(b.Hashes))
	for i, h := range b.Hashes {
		contents[i] = h
	}
	tree, err := basic_merkletree.NewTreeWithHashStrategy(contents, newHash)
	if err != nil {
		return err
	}
	b.tree = tree
	return nil
}

func (b *BlockHashTree) GetBlockHash(number uint64) common.Hash {
	index := number % RECENT_BLOCK_HASHES_LENGTH
	return common.Hash(b.Hashes[index])
}

func (b *BlockHashTree) GetProof(number uint64) ([]common.Hash, uint64, error) {
	index := number % RECENT_BLOCK_HASHES_LENGTH
	proofs, indices, err := b.tree.GetMerklePath(b.Hashes[index])
	if err != nil {
		return nil, 0, err
	}
	hashes := make([]common.Hash, len(proofs))
	for idx, proofItem := range proofs {
		hashes[idx] = common.BytesToHash(proofItem)
	}
	path := uint64(0)
	// Path is reversed
	for i := range indices {
		path *= 2
		path += uint64(indices[len(indices)-i-1])
	}
	return hashes, path, nil
}
