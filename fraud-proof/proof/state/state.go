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
	"math/big"

	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
)

type OneStepState interface {
	Hash() common.Hash
	IsInter() bool
}

type StateType string

const (
	BlockStateType StateType = "BlockState"
	InterStateType StateType = "InterState"
	IntraStateType StateType = "IntraState"
)

type IntraState struct {
	BlockNumber          uint64
	TransactionIdx       uint64
	Depth                uint16
	Gas                  uint64
	Refund               uint64
	LastDepthState       OneStepState
	ContractAddress      common.Address
	Caller               common.Address
	Value                uint256.Int
	CallFlag             CallFlag
	Out                  uint64
	OutSize              uint64
	Pc                   uint64
	OpCode               vm.OpCode
	CodeHash             common.Hash
	Stack                *Stack
	Memory               *Memory
	InputData            *Memory
	ReturnData           *Memory
	CommittedGlobalState vm.StateDB
	GlobalState          vm.StateDB
	SelfDestructSet      *SelfDestructSet
	LogSeries            *LogSeries
	BlockHashTree        *BlockHashTree
	AccessListTrie       *AccessListTrie
}

func (s *IntraState) Hash() common.Hash {
	items := [][]byte{}
	blockNumBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumBytes, s.BlockNumber)
	items = append(items, blockNumBytes)
	txIdxBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(txIdxBytes, s.TransactionIdx)
	items = append(items, txIdxBytes)
	depth := make([]byte, 2)
	binary.BigEndian.PutUint16(depth, s.Depth)
	items = append(items, depth)
	gas := make([]byte, 8)
	binary.BigEndian.PutUint64(gas, s.Gas)
	items = append(items, gas)
	refund := make([]byte, 8)
	binary.BigEndian.PutUint64(refund, s.Refund)
	items = append(items, refund)
	if s.Depth != 1 {
		items = append(items, s.LastDepthState.Hash().Bytes())
		items = append(items, s.ContractAddress.Bytes())
		items = append(items, s.Caller.Bytes())
		valueBytes := s.Value.Bytes32()
		items = append(items, valueBytes[:])
		out := make([]byte, 8)
		binary.BigEndian.PutUint64(out, s.Out)
		outSize := make([]byte, 8)
		binary.BigEndian.PutUint64(outSize, s.OutSize)
		items = append(items, []byte{byte(s.CallFlag)})
		items = append(items, out)
		items = append(items, outSize)
	}
	pc := make([]byte, 8)
	binary.BigEndian.PutUint64(pc, s.Pc)
	items = append(items, pc)
	items = append(items, []byte{byte(s.OpCode)})
	items = append(items, s.CodeHash.Bytes())
	items = append(items, s.Stack.EncodeState())
	items = append(items, s.Memory.EncodeState())
	if s.Depth != 1 {
		items = append(items, s.InputData.EncodeState())
	}
	items = append(items, s.ReturnData.EncodeState())
	items = append(items, s.CommittedGlobalState.GetRootForProof().Bytes())
	items = append(items, s.GlobalState.GetRootForProof().Bytes())
	items = append(items, s.SelfDestructSet.EncodeState())
	items = append(items, s.LogSeries.EncodeState())
	items = append(items, s.BlockHashTree.EncodeState())
	items = append(items, s.AccessListTrie.EncodeState())
	return crypto.Keccak256Hash(items...)
}

func (s *IntraState) IsInter() bool {
	return false
}

func (s *IntraState) StateAsLastDepth(callFlag CallFlag) *IntraState {
	s_ := *s
	s_.Stack = s.Stack.Copy()
	if callFlag == CALLFLAG_CALL || callFlag == CALLFLAG_CALLCODE {
		s_.Stack.PopN(7)
	} else if callFlag == CALLFLAG_DELEGATECALL || callFlag == CALLFLAG_STATICCALL {
		s_.Stack.PopN(6)
	} else if callFlag == CALLFLAG_CREATE {
		s_.Stack.PopN(3)
	} else {
		s_.Stack.PopN(4)
	}
	return &s_
}

func (s *IntraState) HashAsLastDepth(callFlag CallFlag) common.Hash {
	return s.StateAsLastDepth(callFlag).Hash()
}

func StateFromCaptured(
	blockNumber, transactionIdx uint64,
	committedGlobalState vm.StateDB,
	selfDestructSet *SelfDestructSet,
	blockHashTree *BlockHashTree,
	accessListTrie *AccessListTrie,
	evm *vm.EVM,
	lastDepthState OneStepState,
	callFlag CallFlag,
	inputData *Memory,
	out, outSize, pc uint64,
	op vm.OpCode,
	gas, cost uint64,
	memory *vm.Memory,
	stack *vm.Stack,
	contract *vm.Contract,
	rData []byte,
	depth int,
) *IntraState {
	value, _ := uint256.FromBig(contract.Value())
	contractAddress := contract.Address()
	pstack := StackFromEVMStack(stack)
	pmemory := MemoryFromEVMMemory(memory)
	returnData := NewMemoryFromBytes(rData)
	globalState := evm.StateDB.Copy()
	refund := globalState.GetRefund()
	// All pending changes must be committed before getting the root
	globalState.CommitForProof()
	logSeries := LogSeriesFromLogs(globalState.GetCurrentLogs())
	return &IntraState{
		BlockNumber:          blockNumber,
		TransactionIdx:       transactionIdx,
		Depth:                uint16(depth),
		Gas:                  gas,
		Refund:               refund,
		LastDepthState:       lastDepthState,
		ContractAddress:      contractAddress,
		Caller:               contract.Caller(),
		Value:                *value,
		CallFlag:             callFlag,
		Out:                  out,
		OutSize:              outSize,
		Pc:                   pc,
		OpCode:               op,
		CodeHash:             evm.StateDB.GetCodeHash(contractAddress),
		Stack:                pstack,
		Memory:               pmemory,
		InputData:            inputData,
		ReturnData:           returnData,
		CommittedGlobalState: committedGlobalState,
		GlobalState:          globalState,
		SelfDestructSet:      selfDestructSet,
		LogSeries:            logSeries,
		BlockHashTree:        blockHashTree,
		AccessListTrie:       accessListTrie,
	}
}

type InterState struct {
	BlockNumber       uint64
	TransactionIdx    uint64
	GlobalState       vm.StateDB
	CumulativeGasUsed *uint256.Int
	BlockGasUsed      *uint256.Int
	BlockHashTree     *BlockHashTree
	TransactionTrie   *TransactionTrie
	ReceiptTrie       *ReceiptTrie
}

func (s *InterState) Hash() common.Hash {
	items := [][]byte{}
	blockNumber := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumber, s.BlockNumber)
	items = append(items, blockNumber)
	transactionIdx := make([]byte, 8)
	binary.BigEndian.PutUint64(transactionIdx, s.TransactionIdx)
	items = append(items, transactionIdx)
	items = append(items, s.GlobalState.GetRootForProof().Bytes())
	gasBytes := s.CumulativeGasUsed.Bytes32()
	items = append(items, gasBytes[:])
	blockGasBytes := s.BlockGasUsed.Bytes32()
	items = append(items, blockGasBytes[:])
	items = append(items, s.BlockHashTree.EncodeState())
	items = append(items, s.TransactionTrie.EncodeState())
	items = append(items, s.ReceiptTrie.EncodeState())
	return crypto.Keccak256Hash(items...)
}

func (s *InterState) IsInter() bool {
	return true
}

func InterStateFromCaptured(
	blockNumber, transactionIdx uint64,
	statedb vm.StateDB,
	cumulativeGasUsed, blockGasUsed *big.Int,
	transactions types.Transactions,
	receipts types.Receipts,
	blockHashTree *BlockHashTree,
) *InterState {
	cg, _ := uint256.FromBig(cumulativeGasUsed)
	bg, _ := uint256.FromBig(blockGasUsed)
	transactionTrie := NewTransactionTrie(transactions[:transactionIdx])
	receiptTrie := NewReceiptTrie(receipts[:transactionIdx])
	return &InterState{
		BlockNumber:       blockNumber,
		TransactionIdx:    transactionIdx,
		GlobalState:       statedb.Copy(),
		CumulativeGasUsed: cg,
		BlockGasUsed:      bg,
		TransactionTrie:   transactionTrie,
		ReceiptTrie:       receiptTrie,
		BlockHashTree:     blockHashTree,
	}
}

// Represent the state at the end of a finalized block
type BlockState struct {
	BlockNumber       uint64
	GlobalState       vm.StateDB
	CumulativeGasUsed *uint256.Int
	BlockHashTree     *BlockHashTree
}

func (s *BlockState) Hash() common.Hash {
	items := [][]byte{}
	blockNumBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumBytes, s.BlockNumber)
	items = append(items, blockNumBytes)
	items = append(items, s.GlobalState.GetRootForProof().Bytes())
	gasBytes := s.CumulativeGasUsed.Bytes32()
	items = append(items, gasBytes[:])
	items = append(items, s.BlockHashTree.EncodeState())
	return crypto.Keccak256Hash(items...)
}

func (s *BlockState) IsInter() bool {
	return true
}

func BlockStateFromBlock(blockNumber uint64, stateDB vm.StateDB, cumulativeGasUsed *big.Int, blockHashTree *BlockHashTree) (*BlockState, error) {
	g, _ := uint256.FromBig(cumulativeGasUsed)
	return &BlockState{
		BlockNumber:       blockNumber,
		GlobalState:       stateDB.Copy(),
		CumulativeGasUsed: g,
		BlockHashTree:     blockHashTree,
	}, nil
}
