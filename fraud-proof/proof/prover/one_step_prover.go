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

package prover

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/mantlenetworkio/mantle/fraud-proof/proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/params"
)

var ErrStepIdxAndHashMismatch = errors.New("step index and hash mismatch")

type OneStepProver struct {
	// Config
	target common.Hash
	step   uint64

	// Context (read-only)
	rules                params.Rules
	blockNumber          uint64
	transactionIdx       uint64
	committedGlobalState vm.StateDB
	startInterState      *state.InterState
	blockHashTree        *state.BlockHashTree
	transaction          *types.Transaction
	receipt              *types.Receipt

	// Global
	env             *vm.EVM
	counter         uint64
	proof           *proof.OneStepProof
	vmerr           error // Error from EVM execution
	err             error // Error from the tracer
	done            bool
	selfDestructSet *state.SelfDestructSet
	accessListTrie  *state.AccessListTrie

	// Current Call Frame
	callFlag       state.CallFlag
	lastState      *state.IntraState
	lastCost       uint64
	lastCode       []byte
	lastDepthState state.OneStepState
	input          *state.Memory
	out            uint64
	outSize        uint64
	selfDestructed bool
}

// [NewProver] creates a new tracer that generates proofs for:
//   - Type 4 IntraState -> IntraState: one-step EVM execution
//   - Type 5 IntraState -> InterState: transaction finalization
//
// target is the hash of the start state that we want to prove
// step is the step number of the start state (step 0 is the InterState before the transaction)
// The target hash and step should be matched, otherwise [ErrStepIdxAndHashMismatch] will be returned
// receipt is the receipt of the *current* transaction traced
func NewProver(
	target common.Hash,
	step uint64,
	rules params.Rules,
	blockNumber uint64,
	transactionIdx uint64,
	committedGlobalState vm.StateDB,
	interState state.InterState,
	blockHashTree *state.BlockHashTree,
	transaction *types.Transaction,
	receipt *types.Receipt,
) *OneStepProver {
	return &OneStepProver{
		target:               target,
		step:                 step,
		rules:                rules,
		blockNumber:          blockNumber,
		transactionIdx:       transactionIdx,
		committedGlobalState: committedGlobalState,
		startInterState:      &interState,
		blockHashTree:        blockHashTree,
		transaction:          transaction,
		receipt:              receipt,
	}
}

func (l *OneStepProver) CaptureTxStart(gasLimit uint64) {}

func (l *OneStepProver) CaptureTxEnd(restGas uint64) {}

func (l *OneStepProver) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) error {
	// We won't handle transaction initiation proof here, it should be handled outside tracing
	l.env = env
	l.counter = 1
	if create {
		l.callFlag = state.CALLFLAG_CREATE
	} else {
		l.callFlag = state.CALLFLAG_CALL
	}
	l.input = state.NewMemoryFromBytes(input)
	l.accessListTrie = state.NewAccessListTrie()
	l.selfDestructSet = state.NewSelfDestructSet()
	l.startInterState.GlobalState = l.env.StateDB.Copy() // This state includes gas-buying and nonce-increment
	l.lastDepthState = l.startInterState
	log.Info("Capture Start", "from", from, "to", to)
	return nil
}

// CaptureState will be called before the opcode execution
// vmerr is for stack validation and gas validation
// the execution error is captured in CaptureFault
func (l *OneStepProver) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, rData []byte, depth int, vmerr error) error {
	if l.done {
		// Something went wrong during tracing, exit early
		return nil
	}

	defer func() {
		l.counter += 1
	}()

	// Construct the IntraState before the opcode execution
	s := state.StateFromCaptured(
		l.blockNumber,
		l.transactionIdx,
		l.committedGlobalState,
		l.selfDestructSet,
		l.blockHashTree,
		l.accessListTrie,
		l.env,
		l.lastDepthState,
		l.callFlag,
		l.input,
		l.out, l.outSize, pc,
		op,
		gas, cost,
		memory,
		stack,
		contract,
		rData,
		depth,
	)

	log.Info("Generated state", "idx", l.counter, "hash", hexutil.Encode(s.Hash().Bytes()), "op", op)
	log.Info("State", "state", fmt.Sprintf("%+v", s))
	log.Info("State", "stack", fmt.Sprintf("%+v", s.Stack))
	log.Info("State", "memory", fmt.Sprintf("%+v", s.Memory))
	log.Info("State", "input", fmt.Sprintf("%+v", s.InputData))
	log.Info("State", "output", fmt.Sprintf("%+v", s.ReturnData))

	// The target state is found, generate the one-step proof
	if l.counter-1 == l.step {
		l.done = true
		if l.lastState == nil || l.lastState.Hash() != l.target {
			l.err = ErrStepIdxAndHashMismatch
			return nil
		}
		// l.vmerr is the error of l.lastState, either before/during the opcode execution
		// if l.vmerr is not nil, the current state s must be in the parent call frame of l.lastState
		ctx := proof.NewProofGenContext(l.rules, l.env.Context.Coinbase, l.transaction, l.receipt, l.lastCode)
		osp, err := proof.GetIntraProof(ctx, l.lastState, s, l.vmerr)
		if err != nil {
			l.err = err
		} else {
			l.proof = osp
		}
		return nil
	}
	l.lastState = s
	l.lastCode = contract.Code
	// vmerr is not nil means the gas/stack validation failed, the opcode execution will
	// not happen and the current call frame will be immediately reverted. This is the
	// last CaptureState call for this call frame and there won't be any CaptureFault call.
	// Otherwise, vmerr should be cleared.
	l.vmerr = vmerr
	return nil
}

func (l *OneStepProver) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	if l.done {
		// Something went wrong during tracing, exit early
		return
	}
	if typ == vm.SELFDESTRUCT {
		// This enter is for the selfdestruct, record the address
		l.selfDestructed = true
		l.selfDestructSet = l.selfDestructSet.Add(from)
		return
	}
	if l.lastDepthState == nil {
		// Strange error, should not happen
		l.err = errors.New("lastDepthState is nil")
		l.done = true
		return
	}
	// Since CaptureState is called before the opcode execution, here l.lastState is exactly
	// the state before call, so update out and outSize by l.lastState
	// Note: we don't want to update out and outSize in CaptureState because the call opcode
	// might fail before entering the call frame
	if typ == vm.CALL || typ == vm.CALLCODE {
		l.out = l.lastState.Stack.Back(5).Uint64()
		l.outSize = l.lastState.Stack.Back(6).Uint64()
	} else if typ == vm.DELEGATECALL || typ == vm.STATICCALL {
		l.out = l.lastState.Stack.Back(4).Uint64()
		l.outSize = l.lastState.Stack.Back(5).Uint64()
	}
	l.callFlag = state.OpCodeToCallFlag(typ)
	l.lastDepthState = l.lastState.StateAsLastDepth(l.callFlag, l.lastCost)
	l.input = state.NewMemoryFromBytes(input)
}

func (l *OneStepProver) CaptureExit(output []byte, gasUsed uint64, vmerr error) {
	if l.done {
		// Something went wrong during tracing, exit early
		return
	}
	if l.selfDestructed {
		// This exit is for selfdestruct
		l.selfDestructed = false
		return
	}
	// TODO: next line seems unnecessary because CaptureEnd will be instantly called
	// if depth of the last state is 1
	if l.lastState.Depth > 1 {
		lastDepthState := l.lastDepthState.(*state.IntraState)
		l.callFlag = lastDepthState.CallFlag
		l.out = lastDepthState.Out
		l.outSize = lastDepthState.OutSize
		l.input = lastDepthState.InputData
		l.lastDepthState = lastDepthState.LastDepthState
		if vmerr != nil {
			// Call reverted, so revert the selfdestructs and access list changes
			l.selfDestructSet = lastDepthState.SelfDestructSet
			l.accessListTrie = lastDepthState.AccessListTrie
		}
	}
}

// CaptureFault will be called when the stack/gas validation is passed but
// the execution failed. The current call will immediately be reverted.
func (l *OneStepProver) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, vmerr error) error {
	l.vmerr = vmerr
	// The next CaptureState or CaptureEnd will handle the proof generation if needed
	return nil
}

func (l *OneStepProver) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	log.Info("Capture End", "output", output)
	if l.done {
		// Something went wrong during tracing, exit early
		return nil
	}

	// If the last state is the target state, generate the transaction finalization proof
	if l.counter-1 == l.step {
		l.done = true
		if l.lastState.Hash() != l.target {
			l.err = ErrStepIdxAndHashMismatch
			return nil
		}
		// If l.vmerr is not nil, the entire transaction execution will be reverted.
		// Otherwise, the execution ended through STOP or RETURN opcode.
		ctx := proof.NewProofGenContext(l.rules, l.env.Context.Coinbase, l.transaction, l.receipt, l.lastCode)
		osp, err := proof.GetIntraProof(ctx, l.lastState, nil, l.vmerr)
		if err != nil {
			l.err = err
		} else {
			l.proof = osp
		}
	}
	return nil
}

func (l *OneStepProver) GetProof() (*proof.OneStepProof, error) {
	if !l.done {
		return nil, errors.New("proof not generated")
	}
	return l.proof, l.err
}
