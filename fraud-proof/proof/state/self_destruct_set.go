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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SelfDestructSet struct {
	Contracts []common.Address
	Hash      common.Hash
}

func NewSelfDestructSet() *SelfDestructSet {
	return &SelfDestructSet{
		Contracts: make([]common.Address, 0),
		Hash:      common.Hash{},
	}
}

func (s *SelfDestructSet) Add(addr common.Address) *SelfDestructSet {
	contracts := make([]common.Address, len(s.Contracts)+1)
	copy(contracts, s.Contracts)
	h := crypto.Keccak256Hash(s.Hash.Bytes(), addr.Bytes())
	return &SelfDestructSet{
		Contracts: append(contracts, addr),
		Hash:      h,
	}
}

func (s *SelfDestructSet) EncodeState() []byte {
	return s.Hash.Bytes()
}
