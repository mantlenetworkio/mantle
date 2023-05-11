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
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
)

type Stack struct {
	data []uint256.Int
	hash []common.Hash
}

func EmptyStack() *Stack {
	stack := Stack{
		data: make([]uint256.Int, 0),
		hash: []common.Hash{{}},
	}
	return &stack
}

func NewStack(values []uint256.Int) *Stack {
	stack := EmptyStack()
	stack.data = append(stack.data, values...)
	for _, val := range stack.data {
		valueBytes := val.Bytes32()
		h := crypto.Keccak256Hash(stack.Hash().Bytes(), valueBytes[:])
		stack.hash = append(stack.hash, h)
	}
	return stack
}

func StackFromEVMStack(s *vm.Stack) *Stack {
	var values []uint256.Int
	for _, v := range s.Data() {
		value, _ := uint256.FromBig(v)
		values = append(values, *value)
	}
	return NewStack(values)
}

func (st *Stack) Copy() *Stack {
	data := make([]uint256.Int, len(st.data))
	copy(data, st.data)
	hash := make([]common.Hash, len(st.hash))
	copy(hash, st.hash)
	return &Stack{data, hash}
}

func (st *Stack) Data() []uint256.Int {
	return st.data
}

func (st *Stack) Hash() common.Hash {
	return st.hash[len(st.hash)-1]
}

func (st *Stack) PopN(n int) {
	st.data = st.data[:len(st.data)-n]
	st.hash = st.hash[:len(st.hash)-n]
}

func (st *Stack) Len() int {
	return len(st.data)
}

func (st *Stack) Peek() *uint256.Int {
	return &st.data[st.Len()-1]
}

func (st *Stack) Back(n int) *uint256.Int {
	return &st.data[st.Len()-n-1]
}

func (st *Stack) HashAfterPops(n int) common.Hash {
	return st.hash[st.Len()-n]
}

func (st *Stack) EncodeState() []byte {
	encoded := make([]byte, 8)
	binary.BigEndian.PutUint64(encoded, uint64(st.Len()))
	if st.Len() != 0 {
		encoded = append(encoded, st.Hash().Bytes()...)
	}
	return encoded
}
