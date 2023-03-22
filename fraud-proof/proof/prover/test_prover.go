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
	"encoding/json"
	"errors"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"math/big"
	"strconv"
	"strings"
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

func bytesToHex(s []byte) string {
	return "0x" + common.Bytes2Hex(s)
}

func bigToHex(n *big.Int) string {
	if n == nil {
		return "0x0"
	}
	return "0x" + n.Text(16)
}

func uintToHex(n uint64) string {
	return "0x" + strconv.FormatUint(n, 16)
}

func addrToHex(a common.Address) string {
	return strings.ToLower(a.Hex())
}

type OspTestResult struct {
	Ctx   OspTestGeneratedCtx `json:"ctx"`
	Proof OspTestProof        `json:"proof"`
}

type OspTestGeneratedCtx struct {
	TxnHash     string `json:"txnHash"`
	TxNonce     string `json:"txNonce"`
	TxV         string `json:"txV"`
	TxR         string `json:"txR"`
	TxS         string `json:"txS"`
	Coinbase    string `json:"coinbase"`
	Timestamp   string `json:"timestamp"`
	BlockNumber string `json:"blockNumber"`
	Difficulty  string `json:"difficulty"`
	GasLimit    string `json:"gasLimit"`
	ChainID     string `json:"chainID"`
	BaseFee     string `json:"baseFee"`
	Origin      string `json:"origin"`
	Recipient   string `json:"recipient"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Input       string `json:"input"`
	InputSize   string `json:"inputSize"`
}

type OspTestProof struct {
	Opcode    string `json:"opcode"`
	Verifier  uint64 `json:"verifier"`
	CurrHash  string `json:"currHash"`
	NextHash  string `json:"nextHash"`
	ProofSize string `json:"proofSize"`
	CodeSize  string `json:"codeSize"`
	Proof     string `json:"proof"`
	Idx       uint64 `json:"idx"`
}

type TestProver struct {
	// Config
	step   int64
	opcode int64
	// Context (read-only)
	transaction          *types.Transaction
	txctx                *vm.Context
	receipt              *types.Receipt
	rules                params.Rules
	blockNumber          uint64
	transactionIdx       uint64
	committedGlobalState vm.StateDB
	startInterState      *state.InterState
	blockHashTree        *state.BlockHashTree

	// Global
	env             *vm.EVM
	counter         uint64
	vmerr           error // Error from EVM execution
	err             error // Error from the tracer
	done            bool
	ctx             OspTestGeneratedCtx
	proof           OspTestProof
	selfDestructSet *state.SelfDestructSet
	accessListTrie  *state.AccessListTrie

	// Current Call Frame
	callFlag       state.CallFlag
	lastState      *state.IntraState
	lastCost       uint64
	lastCode       []byte
	lastOpcode     byte
	lastDepthState state.OneStepState
	input          *state.Memory
	out            uint64
	outSize        uint64
	selfDestructed bool
}

func NewTestProver(
	step int64,
	opcode int64,
	transaction *types.Transaction,
	txctx *vm.Context,
	receipt *types.Receipt,
	rules params.Rules,
	blockNumber uint64,
	transactionIdx uint64,
	committedGlobalState vm.StateDB,
	interState state.InterState,
	blockHashTree *state.BlockHashTree,
) *TestProver {
	return &TestProver{
		step:                 step,
		opcode:               opcode,
		transaction:          transaction,
		txctx:                txctx,
		receipt:              receipt,
		rules:                rules,
		blockNumber:          blockNumber,
		transactionIdx:       transactionIdx,
		committedGlobalState: committedGlobalState,
		startInterState:      &interState,
		blockHashTree:        blockHashTree,
	}
}

func (l *TestProver) CaptureTxStart(gasLimit uint64) {}

func (l *TestProver) CaptureTxEnd(restGas uint64) {}

func (l *TestProver) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) error {
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
	vmctx := l.env.Context
	recipient := common.Address{}
	if l.transaction.To() != nil {
		recipient = *l.transaction.To()
	}
	v, r, s := l.transaction.RawSignatureValues()
	l.ctx = OspTestGeneratedCtx{
		TxnHash:     bytesToHex(l.transaction.Hash().Bytes()),
		TxNonce:     uintToHex(l.transaction.Nonce()),
		TxV:         bigToHex(v),
		TxR:         bigToHex(r),
		TxS:         bigToHex(s),
		Coinbase:    addrToHex(vmctx.Coinbase),
		Timestamp:   bigToHex(vmctx.Time),
		BlockNumber: bigToHex(vmctx.BlockNumber),
		Difficulty:  bigToHex(vmctx.Difficulty),
		GasLimit:    uintToHex(vmctx.GasLimit),
		ChainID:     bigToHex(l.transaction.ChainId()),
		//BaseFee:     bigToHex(vmctx.BaseFee),
		Origin:    addrToHex(l.txctx.Origin),
		Recipient: addrToHex(recipient),
		Value:     bigToHex(l.transaction.Value()),
		Gas:       uintToHex(l.transaction.Gas()),
		GasPrice:  bigToHex(l.txctx.GasPrice),
		Input:     bytesToHex(input),
		InputSize: uintToHex(l.input.Size()),
	}
	log.Info("Capture Start", "from", from, "to", to)
	return nil
}

func (l *TestProver) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, rData []byte, depth int, vmerr error) error {
	if l.done {
		return vmerr
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
	// log.Info("State", "state", fmt.Sprintf("%+v", s))
	// log.Info("State", "stack", fmt.Sprintf("%+v", s.Stack))
	// log.Info("State", "memory", fmt.Sprintf("%+v", s.Memory))
	// log.Info("State", "input", fmt.Sprintf("%+v", s.InputData))
	// log.Info("State", "output", fmt.Sprintf("%+v", s.ReturnData))

	// The target state is found, generate the one-step proof
	if int64(l.counter-1) == l.step || int64(l.lastOpcode) == l.opcode {
		l.done = true
		if l.lastState == nil {
			l.err = ErrStepIdxAndHashMismatch
			return vmerr
		}
		// l.vmerr is the error of l.lastState, either before/during the opcode execution
		// if l.vmerr is not nil, the current state s must be in the parent call frame of l.lastState
		ctx := proof.NewProofGenContext(l.rules, l.env.Context.Coinbase, l.transaction, l.receipt, l.lastCode)
		osp, err := proof.GetIntraProof(ctx, l.lastState, s, l.vmerr)
		if err != nil {
			l.err = err
		} else {
			encoded := osp.Encode()
			l.proof = OspTestProof{
				Opcode:    l.lastState.OpCode.String(),
				Verifier:  uint64(osp.VerifierType),
				CurrHash:  bytesToHex(l.lastState.Hash().Bytes()),
				NextHash:  bytesToHex(s.Hash().Bytes()),
				ProofSize: uintToHex(uint64(len(encoded))),
				CodeSize:  uintToHex(osp.TotalCodeSize),
				Proof:     bytesToHex(encoded),
				Idx:       l.counter - 1,
			}
			log.Info("SHOW STATE", "length of lastState", len(l.lastState.Encode()))
			log.Info("SHOW STATE", "length of encoded", len(encoded))

			log.Info("SHOW STATE", "hash of lastState", l.lastState.Hash().String())
			log.Info("SHOW STATE", "hash of encoded", crypto.Keccak256Hash(encoded[:len(l.lastState.Encode())]).String())
		}
		return vmerr
	}
	l.lastState = s
	l.lastCode = contract.Code
	l.lastCost = cost
	l.lastOpcode = byte(op)
	// vmerr is not nil means the gas/stack validation failed, the opcode execution will
	// not happen and the current call frame will be immediately reverted. This is the
	// last CaptureState call for this call frame and there won't be any CaptureFault call.
	// Otherwise, vmerr should be cleared.
	l.vmerr = vmerr
	return vmerr
}

func (l *TestProver) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
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

func (l *TestProver) CaptureExit(output []byte, gasUsed uint64, vmerr error) {
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
func (l *TestProver) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, vmerr error) error {
	l.vmerr = vmerr
	// The next CaptureState or CaptureEnd will handle the proof generation if needed
	return vmerr
}

func (l *TestProver) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, vmerr error) error {
	log.Info("Capture End", "output", output)
	if l.done {
		// Something went wrong during tracing, exit early
		return vmerr
	}

	if int64(l.counter-1) == l.step {
		l.done = true
		// If l.vmerr is not nil, the entire transaction execution will be reverted.
		// Otherwise, the execution ended through STOP or RETURN opcode.
		ctx := proof.NewProofGenContext(l.rules, l.env.Context.Coinbase, l.transaction, l.receipt, l.lastCode)
		osp, err := proof.GetIntraProof(ctx, l.lastState, nil, l.vmerr)
		if err != nil {
			l.err = err
		} else {
			encoded := osp.Encode()
			l.proof = OspTestProof{
				Opcode:    l.lastState.OpCode.String(),
				CurrHash:  bytesToHex(l.lastState.Hash().Bytes()),
				NextHash:  bytesToHex(common.Hash{}.Bytes()), // TODO: get the hash of next InterState
				ProofSize: uintToHex(uint64(len(encoded))),
				Proof:     bytesToHex(encoded),
				CodeSize:  uintToHex(osp.TotalCodeSize),
				Idx:       l.counter - 1,
			}
		}
	}
	return vmerr
}

func (l *TestProver) GetResult() (json.RawMessage, error) {
	if l.err != nil {
		return nil, l.err
	}
	result := OspTestResult{
		Ctx:   l.ctx,
		Proof: l.proof,
	}
	res, err := json.Marshal(result)
	if err != nil {
		log.Error("Err", "err", err)
		return nil, err
	}
	return json.RawMessage(res), nil
}
