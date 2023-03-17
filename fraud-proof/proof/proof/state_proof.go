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

	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

type IntraStateProof struct {
	BlockNumber              uint64
	TransactionIdx           uint64
	Depth                    uint16
	Gas                      uint64
	Refund                   uint64
	LastDepthHash            common.Hash
	ContractAddress          common.Address
	Caller                   common.Address
	Value                    uint256.Int
	CallFlag                 state.CallFlag
	Out                      uint64
	OutSize                  uint64
	Pc                       uint64
	OpCode                   vm.OpCode
	CodeHash                 common.Hash
	StackHash                common.Hash
	StackSize                uint64
	MemorySize               uint64
	MemoryRoot               common.Hash
	InputDataSize            uint64
	InputDataRoot            common.Hash
	ReturnDataSize           uint64
	ReturnDataRoot           common.Hash
	CommittedGlobalStateRoot common.Hash
	GlobalStateRoot          common.Hash
	SelfDestructAcc          common.Hash
	LogAcc                   common.Hash
	BlockHashRoot            common.Hash
	AccesslistRoot           common.Hash
}

func StateProofFromState(s *state.IntraState) *IntraStateProof {
	var lastDepthHash common.Hash
	if s.LastDepthState != nil {
		lastDepthHash = s.LastDepthState.Hash()
	}
	return &IntraStateProof{
		BlockNumber:              s.BlockNumber,
		TransactionIdx:           s.TransactionIdx,
		Depth:                    s.Depth,
		Gas:                      s.Gas,
		Refund:                   s.Refund,
		LastDepthHash:            lastDepthHash,
		ContractAddress:          s.ContractAddress,
		Caller:                   s.Caller,
		Value:                    s.Value,
		CallFlag:                 s.CallFlag,
		Out:                      s.Out,
		OutSize:                  s.OutSize,
		Pc:                       s.Pc,
		OpCode:                   s.OpCode,
		CodeHash:                 s.CodeHash,
		StackSize:                uint64(s.Stack.Len()),
		StackHash:                s.Stack.Hash(),
		MemorySize:               s.Memory.Size(),
		MemoryRoot:               s.Memory.Root(),
		InputDataSize:            s.InputData.Size(),
		InputDataRoot:            s.InputData.Root(),
		ReturnDataSize:           s.ReturnData.Size(),
		ReturnDataRoot:           s.ReturnData.Root(),
		CommittedGlobalStateRoot: s.CommittedGlobalState.GetRootForProof(),
		GlobalStateRoot:          s.GlobalState.GetRootForProof(),
		SelfDestructAcc:          s.SelfDestructSet.Hash,
		LogAcc:                   s.LogSeries.Hash(),
		BlockHashRoot:            s.BlockHashTree.Root(),
		AccesslistRoot:           s.AccessListTrie.Root(),
	}
}

func (s *IntraStateProof) Encode() []byte {
	proofLen := 8 + 8 + 2 + 8 + 8 + 32 + 8 + 1 + 32 + 8 + 8 + 8 + 32 + 32 + 32 + 32 + 32 + 32 // BlockNumber, TransactionIdx, Depth, Gas, Refund, LastDepthHash, Pc, OpCode, CodeMerkle, StackSize, MemorySize, ReturnDataSize, CommittedGlobalStateRoot, GlobalStateRoot, SelfDestructAcc, LogAcc, BlockHashRoot, AccesslistRoot
	if s.Depth != 1 {
		proofLen += 20 + 20 + 32 + 8 + 1 + 8 + 8 // ContractAddress, Caller, Value, CallFlag, Out, OutSize, InputDataSize
		if s.InputDataSize != 0 {
			proofLen += 32 // InputDataRoot
		}
	}
	if s.StackSize != 0 {
		proofLen += 32 // StackHash
	}
	if s.MemorySize != 0 {
		proofLen += 32 // MemoryRoot
	}
	if s.ReturnDataSize != 0 {
		proofLen += 32 // ReturnDataRoot
	}
	encoded := make([]byte, proofLen)
	blockNumber := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumber, s.BlockNumber)
	transactionIdx := make([]byte, 8)
	binary.BigEndian.PutUint64(transactionIdx, s.TransactionIdx)
	depth := make([]byte, 2)
	binary.BigEndian.PutUint16(depth, s.Depth)
	gas := make([]byte, 8)
	binary.BigEndian.PutUint64(gas, s.Gas)
	refund := make([]byte, 8)
	binary.BigEndian.PutUint64(refund, s.Refund)
	pc := make([]byte, 8)
	binary.BigEndian.PutUint64(pc, s.Pc)
	stackSize := make([]byte, 8)
	binary.BigEndian.PutUint64(stackSize, s.StackSize)
	memSize := make([]byte, 8)
	binary.BigEndian.PutUint64(memSize, s.MemorySize)
	var inputDataSize []byte
	if s.Depth != 1 {
		inputDataSize = make([]byte, 8)
		binary.BigEndian.PutUint64(inputDataSize, s.InputDataSize)
	}
	returnDataSize := make([]byte, 8)
	binary.BigEndian.PutUint64(returnDataSize, s.ReturnDataSize)
	copy(encoded, blockNumber)
	copy(encoded[8:], transactionIdx)
	copy(encoded[16:], depth)
	copy(encoded[18:], gas)
	copy(encoded[26:], refund)
	copy(encoded[34:], s.LastDepthHash.Bytes())
	offset := 66
	if s.Depth != 1 {
		copy(encoded[offset:], s.ContractAddress.Bytes())
		copy(encoded[offset+20:], s.Caller.Bytes())
		valueBytes := s.Value.Bytes32()
		copy(encoded[offset+20+20:], valueBytes[:])
		out := make([]byte, 8)
		binary.BigEndian.PutUint64(out, s.Out)
		outSize := make([]byte, 8)
		binary.BigEndian.PutUint64(outSize, s.OutSize)
		encoded[offset+20+20+32] = byte(s.CallFlag)
		copy(encoded[offset+20+20+32+1:], out)
		copy(encoded[offset+20+20+32+1+8:], outSize)
		offset += 20 + 20 + 32 + 1 + 8 + 8
	}
	copy(encoded[offset:], pc)
	encoded[offset+8] = byte(s.OpCode)
	copy(encoded[offset+8+1:], s.CodeHash.Bytes())
	copy(encoded[offset+8+1+32:], stackSize)
	offset += 8 + 1 + 32 + 8
	if s.StackSize != 0 {
		copy(encoded[offset:], s.StackHash.Bytes())
		offset += 32
	}
	copy(encoded[offset:], memSize)
	offset += 8
	if s.MemorySize != 0 {
		copy(encoded[offset:], s.MemoryRoot.Bytes())
		offset += 32
	}
	if s.Depth != 1 {
		copy(encoded[offset:], inputDataSize)
		offset += 8
		if s.InputDataSize != 0 {
			copy(encoded[offset:], s.InputDataRoot.Bytes())
			offset += 32
		}
	}
	copy(encoded[offset:], returnDataSize)
	offset += 8
	if s.ReturnDataSize != 0 {
		copy(encoded[offset:], s.ReturnDataRoot.Bytes())
		offset += 32
	}
	copy(encoded[offset:], s.CommittedGlobalStateRoot.Bytes())
	copy(encoded[offset+32:], s.GlobalStateRoot.Bytes())
	copy(encoded[offset+64:], s.SelfDestructAcc.Bytes())
	copy(encoded[offset+96:], s.LogAcc.Bytes())
	copy(encoded[offset+128:], s.BlockHashRoot.Bytes())
	copy(encoded[offset+160:], s.AccesslistRoot.Bytes())
	return encoded
}

func (s *IntraStateProof) Hash() common.Hash {
	return crypto.Keccak256Hash(s.Encode())
}

type InterStateProof struct {
	BlockNumber         uint64
	TransactionIdx      uint64
	GlobalStateRoot     common.Hash
	BlockGasUsed        uint256.Int
	BlockHashRoot       common.Hash
	TransactionTireRoot common.Hash
	ReceiptTrieRoot     common.Hash
	LogsBloom           types.Bloom
}

func InterStateProofFromInterState(s *state.InterState) *InterStateProof {
	return &InterStateProof{
		BlockNumber:         s.BlockNumber,
		TransactionIdx:      s.TransactionIdx,
		GlobalStateRoot:     s.GlobalState.GetRootForProof(),
		BlockGasUsed:        *s.BlockGasUsed,
		BlockHashRoot:       s.BlockHashTree.Root(),
		TransactionTireRoot: s.TransactionTrie.Root(),
		ReceiptTrieRoot:     s.ReceiptTrie.Root(),
		LogsBloom:           s.ReceiptTrie.Bloom(),
	}
}

func (s *InterStateProof) Encode() []byte {
	encoded := make([]byte, 464)
	binary.BigEndian.PutUint64(encoded, s.BlockNumber)
	binary.BigEndian.PutUint64(encoded[8:], s.TransactionIdx)
	copy(encoded[16:], s.GlobalStateRoot.Bytes())
	blockGasBytes := s.BlockGasUsed.Bytes32()
	copy(encoded[80:], blockGasBytes[:])
	copy(encoded[112:], s.BlockHashRoot.Bytes())
	copy(encoded[144:], s.TransactionTireRoot.Bytes())
	copy(encoded[176:], s.ReceiptTrieRoot.Bytes())
	copy(encoded[208:], s.LogsBloom.Bytes())
	return encoded
}

type BlockStateProof struct {
	BlockNumber       uint64
	GlobalStateRoot   common.Hash
	CumulativeGasUsed uint256.Int
	BlockHashRoot     common.Hash
}

func BlockStateProofFromBlockState(s *state.BlockState) *BlockStateProof {
	return &BlockStateProof{
		BlockNumber:     s.BlockNumber,
		GlobalStateRoot: s.GlobalState.GetRootForProof(),
		BlockHashRoot:   s.BlockHashTree.Root(),
	}
}

func (s *BlockStateProof) Encode() []byte {
	encoded := make([]byte, 104)
	binary.BigEndian.PutUint64(encoded, s.BlockNumber)
	copy(encoded[8:], s.GlobalStateRoot.Bytes())
	gasBytes := s.CumulativeGasUsed.Bytes32()
	copy(encoded[40:], gasBytes[:])
	copy(encoded[72:], s.BlockHashRoot.Bytes())
	return encoded
}

type ReceiptProof struct {
	RLPEncodedReceipt []byte
}

func ReceiptProofFromReceipt(r *types.Receipt) *ReceiptProof {
	encoded, _ := rlp.EncodeToBytes(r)
	return &ReceiptProof{
		RLPEncodedReceipt: encoded,
	}
}

func (s *ReceiptProof) Encode() []byte {
	length := len(s.RLPEncodedReceipt)
	encoded := make([]byte, 8+length)
	binary.BigEndian.PutUint64(encoded, uint64(length))
	copy(encoded[8:], s.RLPEncodedReceipt)
	return encoded
}
