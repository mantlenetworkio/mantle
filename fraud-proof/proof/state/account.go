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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

type Account struct {
	Nonce       uint64
	Balance     uint256.Int
	StorageRoot common.Hash
	CodeHash    common.Hash
}

func AccountFromEVM(acc types.StateAccount) *Account {
	balance, _ := uint256.FromBig(acc.Balance)
	codeHash := common.Hash{}
	if acc.CodeHash != nil && len(acc.CodeHash) >= 0 {
		codeHash = common.BytesToHash(acc.CodeHash)
	}
	return &Account{
		Nonce:       acc.Nonce,
		Balance:     *balance,
		StorageRoot: acc.Root,
		CodeHash:    codeHash,
	}
}

func (acc *Account) Encode() []byte {
	encLen := 8 + 32 + 32 + 32
	encoded := make([]byte, encLen)
	nonce := make([]byte, 8)
	binary.BigEndian.PutUint64(nonce, acc.Nonce)
	copy(encoded, nonce)
	copy(encoded[8:], acc.Balance.Bytes())
	copy(encoded[8+32:], acc.StorageRoot.Bytes())
	copy(encoded[8+32+32:], acc.CodeHash.Bytes())
	return encoded
}
