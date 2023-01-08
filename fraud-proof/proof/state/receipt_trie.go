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
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb/memorydb"
	"github.com/mantlenetworkio/mantle/l2geth/trie"
)

type ReceiptTrie struct {
	receipts types.Receipts
	trie     *trie.Trie
}

func NewReceiptTrie(receipts types.Receipts) *ReceiptTrie {
	t, _ := trie.New(common.Hash{}, trie.NewDatabase(memorydb.New()))
	valueBuf := new(bytes.Buffer)
	var indexBuf []byte
	for i := 0; i < receipts.Len(); i++ {
		indexBuf = rlp.AppendUint64(indexBuf[:0], uint64(i))
		valueBuf.Reset()
		valueBuf.Write(receipts.GetRlp(i))
		t.Update(indexBuf, common.CopyBytes(valueBuf.Bytes()))
	}
	return &ReceiptTrie{
		receipts: receipts,
		trie:     t,
	}
}

func (r *ReceiptTrie) Root() common.Hash {
	return r.trie.Hash()
}

func (r *ReceiptTrie) Bloom() types.Bloom {
	return types.CreateBloom(r.receipts)
}

func (r *ReceiptTrie) EncodeState() []byte {
	encoded := r.Root().Bytes()
	encoded = append(encoded, r.Bloom().Bytes()...)
	return encoded
}

func (r *ReceiptTrie) GetReceipt(index int) *types.Receipt {
	return r.receipts[index]
}

func (r *ReceiptTrie) Prove(index int) ([][]byte, error) {
	var proof mptProofList
	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf[:0], uint64(index))
	err := r.trie.Prove(indexBuf, 0, &proof)
	return proof, err
}
