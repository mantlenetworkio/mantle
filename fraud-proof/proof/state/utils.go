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
	"hash"

	"github.com/ethereum/go-ethereum/crypto"
)

type mptProofList [][]byte

func (n *mptProofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

func (n *mptProofList) Delete(key []byte) error {
	panic("not supported")
}

func newHash() hash.Hash {
	return crypto.NewKeccakState()
}
