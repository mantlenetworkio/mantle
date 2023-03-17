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
	"errors"
	"math/big"

	"github.com/holiman/uint256"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/params"
)

func (osp *OneStepProof) addStateProof(s *state.IntraState) {
	osp.AddProof(StateProofFromState(s))
}

func (osp *OneStepProof) addInterStateProof(s *state.InterState) {
	osp.AddProof(InterStateProofFromInterState(s))
}

func (osp *OneStepProof) addRawCodeProof(code []byte) {
	codeProof := &CodeProof{
		Content: code,
	}
	osp.TotalCodeSize += uint64(len(code))
	osp.AddProof(codeProof)
}

func (osp *OneStepProof) addCodeProof(currState *state.IntraState, address common.Address) {
	osp.addRawCodeProof(currState.GlobalState.GetCode(address))
}

func (osp *OneStepProof) addOpCodeProof(ctx ProofGenContext, currState *state.IntraState) {
	osp.addRawCodeProof(ctx.actualCode)
}

func (osp *OneStepProof) addBlockHashProof(num uint64, currState *state.IntraState) error {
	if num >= currState.BlockNumber || num < currState.BlockNumber-state.RECENT_BLOCK_HASHES_LENGTH {
		return nil
	}
	pf, err := GetBlockHashMerkleProof(currState.BlockHashTree, num)
	if err != nil {
		return err
	}
	osp.AddProof(&BlockHashProof{currState.BlockHashTree.GetBlockHash(num)})
	osp.AddProof(pf)
	return nil
}

func (osp *OneStepProof) addStackProof(popNum int, currState *state.IntraState) {
	if popNum == 0 {
		return
	}
	pops := make([]uint256.Int, popNum)
	for i := 0; i < popNum; i++ {
		pops[i] = *currState.Stack.Back(i)
	}
	stackProof := &StackProof{
		Pops:               pops,
		StackHashAfterPops: currState.Stack.HashAfterPops(popNum),
	}
	osp.AddProof(stackProof)
}

func (osp *OneStepProof) addMemoryLikeReadProof(offset, size uint64, memoryLike *state.Memory) error {
	// We never need Merkle proof is the memory is empty
	if memoryLike.CellNum() == 0 {
		return nil
	}
	pf, err := generateMemoryReadProof(memoryLike, offset, size)
	if err != nil {
		return err
	}
	if pf != nil {
		osp.AddProof(pf)
	}
	return nil
}

func (osp *OneStepProof) addMemoryLikeReadProofNoAppend(offset, size uint64, memoryLike *state.Memory) error {
	// We never need Merkle proof is the memory is empty
	if memoryLike.CellNum() == 0 {
		return nil
	}
	pf, err := generateMemoryReadProofNoAppend(memoryLike, offset, size)
	if err != nil {
		return err
	}
	if pf != nil {
		osp.AddProof(pf)
	}
	return nil
}

func (osp *OneStepProof) addMemoryReadProof(offset, size uint64, currState *state.IntraState) error {
	return osp.addMemoryLikeReadProof(offset, size, currState.Memory)
}

func (osp *OneStepProof) addInputDataReadProof(offset, size uint64, currState *state.IntraState) error {
	return osp.addMemoryLikeReadProofNoAppend(offset, size, currState.InputData)
}

func (osp *OneStepProof) addReturnDataReadProof(offset, size uint64, currState *state.IntraState) error {
	return osp.addMemoryLikeReadProofNoAppend(offset, size, currState.ReturnData)
}

func (osp *OneStepProof) addMemoryWriteProof(offset, size uint64, currState, nextState *state.IntraState) error {
	// We never need Merkle proof is the memory is empty
	if currState.Memory.CellNum() == 0 {
		return nil
	}
	pf, err := generateMemoryWriteProof(currState.Memory, nextState.Memory, offset, size)
	if err != nil {
		return err
	}
	if pf != nil {
		osp.AddProof(pf)
	}
	return nil
}

func (osp *OneStepProof) addCurrAccountProof(currState *state.IntraState) error {
	return osp.addAccountProof(currState, currState.ContractAddress)
}

func (osp *OneStepProof) addAccountProof(currState *state.IntraState, address common.Address) error {
	accPf, err := GetAccountProof(currState.GlobalState, address)
	if err != nil {
		return err
	}
	osp.AddProof(accPf)
	return nil
}

func (osp *OneStepProof) addCommittedAccountProof(currState *state.IntraState, address common.Address) error {
	accPf, err := GetAccountProof(currState.CommittedGlobalState, address)
	if err != nil {
		return err
	}
	osp.AddProof(accPf)
	return nil
}

func (osp *OneStepProof) addStorageLoadProof(key common.Hash, currState *state.IntraState) error {
	accPf, stPf, err := GetStorageProof(currState.GlobalState, currState.ContractAddress, key)
	if err != nil {
		return err
	}
	osp.AddProof(accPf)
	osp.AddProof(stPf)
	return nil
}

func (osp *OneStepProof) addCommittedStorageProof(key common.Hash, currState *state.IntraState) error {
	accPf, stPf, err := GetStorageProof(currState.CommittedGlobalState, currState.ContractAddress, key)
	if err != nil {
		return err
	}
	osp.AddProof(accPf)
	osp.AddProof(stPf)
	return nil
}

func (osp *OneStepProof) addStorageStoreProof(key common.Hash, currState, nextState *state.IntraState) error {
	// We need to add 2 storage proofs here
	osp.addCommittedStorageProof(key, currState)
	osp.addStorageLoadProof(key, currState)
	return nil
}

func (osp *OneStepProof) addLogProof(currState *state.IntraState) error {
	logProof := LogProofFromLogSeries(currState.LogSeries)
	osp.AddProof(logProof)
	return nil
}

// If a call frame is reverted to its parent call frame, we only need to provide the proof
// of the lastDepthState, which is the state it will revert to
// If the revert is called by opRevert, we also need to provide memory proofs for
// return data reading and writing in memory, but it is handled in [opRevertProof]
func (osp *OneStepProof) addCallRevertProof(currState *state.IntraState) error {
	lastDepthState := currState.LastDepthState.(*state.IntraState)
	osp.addStateProof(lastDepthState)
	return nil
}

// If a transaction is reverted, nothing changed in the state except the gas tip to the coinbase.
// We need to send the InterStateProof for the verifier to construct the next InterState.
// We also need to send the receipt to help the verifier update the receipt trie.
func (osp *OneStepProof) addTransactionRevertProof(ctx ProofGenContext, currState *state.IntraState, vmerr error) error {
	lastDepthState := currState.LastDepthState.(*state.InterState)
	osp.addInterStateProof(lastDepthState)
	err := osp.addCommittedAccountProof(currState, ctx.coinbase)
	if err != nil {
		return err
	}
	osp.AddProof(ReceiptProofFromReceipt(ctx.receipt))
	return nil
}

func (osp *OneStepProof) addRevertProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) error {
	if nextState == nil {
		// We're in the topmost call frame, the entire transaction will be reverted
		return osp.addTransactionRevertProof(ctx, currState, vmerr)
	}
	return osp.addCallRevertProof(currState)
}

func makeStackOnlyInstructionProof(popNum int) func(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	return func(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
		osp := EmptyProof()
		if isEnvironmentalOp(currState.OpCode) {
			osp.SetVerifierType(VerifierTypeEnvironmentalOp)
		} else {
			osp.SetVerifierType(VerifierTypeStackOp)
		}
		osp.addStateProof(currState)
		osp.addOpCodeProof(ctx, currState)
		if popNum > 0 {
			// If the stack validation fails, we don't need to provide stack proofs
			if IsStackError(vmerr) {
				err := osp.addRevertProof(ctx, currState, nextState, vmerr)
				if err != nil {
					return nil, err
				}
				return osp, nil
			}
			osp.addStackProof(popNum, currState)
		}
		// If error happens after the stack validation, in most cases we only need the stack proof
		// and revert proof
		if vmerr != nil {
			err := osp.addRevertProof(ctx, currState, nextState, vmerr)
			if err != nil {
				return nil, err
			}
		}
		return osp, nil
	}
}

func opPushProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	// We reuse makeStackOnlyInstructionProof here because the content pushed is covered by the opcode proof
	// Note: if we switch to other proof scheme of code, we may need to provide a separate proof for the pushed content
	return makeStackOnlyInstructionProof(0)(ctx, currState, nextState, vmerr)
}

func makeLogProof(topicNum int) func(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	return func(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
		osp := EmptyProof()
		osp.SetVerifierType(VerifierTypeMemoryOp)
		osp.addStateProof(currState)
		osp.addOpCodeProof(ctx, currState)
		// If the stack validation fails, we don't need to provide stack proofs
		if IsStackError(vmerr) {
			err := osp.addRevertProof(ctx, currState, nextState, vmerr)
			if err != nil {
				return nil, err
			}
			return osp, nil
		}
		osp.addStackProof(topicNum+2, currState)
		// If error happens after the stack validation, in most cases we only need the stack proof
		// and revert proof
		if vmerr != nil {
			err := osp.addRevertProof(ctx, currState, nextState, vmerr)
			if err != nil {
				return nil, err
			}
		}
		offset := currState.Stack.Peek().Uint64()
		size := currState.Stack.Back(1).Uint64()
		err := osp.addMemoryReadProof(offset, size, currState)
		if err != nil {
			return nil, err
		}
		err = osp.addLogProof(currState)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
}

func opKeccak256Proof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	offset := currState.Stack.Peek().Uint64()
	size := currState.Stack.Back(1).Uint64()
	err := osp.addMemoryReadProof(offset, size, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opBalanceProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	addrBytes := currState.Stack.Peek().Bytes32()
	address := common.BytesToAddress(addrBytes[:])
	err := osp.addAccountProof(currState, address)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCallDataLoadProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	offset := currState.Stack.Peek().Uint64()
	err := osp.addInputDataReadProof(offset, 32, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCallDataCopyProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(3, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	memOffset := currState.Stack.Peek().Uint64()
	offset := currState.Stack.Back(1).Uint64()
	size := currState.Stack.Back(2).Uint64()
	err := osp.addInputDataReadProof(offset, size, currState)
	if err != nil {
		return nil, err
	}
	err = osp.addMemoryWriteProof(memOffset, size, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCodeCopyProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(3, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	memOffset := currState.Stack.Peek().Uint64()
	size := currState.Stack.Back(2).Uint64()
	err := osp.addMemoryWriteProof(memOffset, size, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opExtCodeSizeProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	addrBytes := currState.Stack.Peek().Bytes32()
	address := common.BytesToAddress(addrBytes[:])
	err := osp.addAccountProof(currState, address)
	if err != nil {
		return nil, err
	}
	osp.addCodeProof(currState, address)
	return osp, nil
}

func opExtCodeCopyProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(4, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	addrBytes := currState.Stack.Peek().Bytes32()
	address := common.BytesToAddress(addrBytes[:])
	memOffset := currState.Stack.Back(1).Uint64()
	size := currState.Stack.Back(3).Uint64()
	err := osp.addAccountProof(currState, address)
	if err != nil {
		return nil, err
	}
	osp.addCodeProof(currState, address)
	err = osp.addMemoryWriteProof(memOffset, size, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opExtCodeHashProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	addrBytes := currState.Stack.Peek().Bytes32()
	address := common.BytesToAddress(addrBytes[:])
	err := osp.addAccountProof(currState, address)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opReturnDataCopyProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(3, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	memOffset := currState.Stack.Peek().Uint64()
	offset := currState.Stack.Back(1).Uint64()
	size := currState.Stack.Back(2).Uint64()
	err := osp.addReturnDataReadProof(offset, size, currState)
	if err != nil {
		return nil, err
	}
	err = osp.addMemoryWriteProof(memOffset, size, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opSelfBalanceProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	err := osp.addCurrAccountProof(currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opBlockHashProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeEnvironmentalOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	num := currState.Stack.Peek().Uint64()
	err := osp.addBlockHashProof(num, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opMLoadProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	offset := currState.Stack.Peek().Uint64()
	err := osp.addMemoryReadProof(offset, 32, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opMStoreProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	offset := currState.Stack.Peek().Uint64()
	err := osp.addMemoryWriteProof(offset, 32, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opMStore8Proof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeMemoryOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	offset := currState.Stack.Peek().Uint64()
	err := osp.addMemoryWriteProof(offset, 32, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opSLoadProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	keyBytes := currState.Stack.Peek().Bytes32()
	err := osp.addStorageLoadProof(common.BytesToHash(keyBytes[:]), currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opSStoreProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStorageOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	// If error happens after the stack validation, in most cases we only need the stack proof
	// and revert proof
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	keyBytes := currState.Stack.Peek().Bytes32()
	err := osp.addStorageStoreProof(common.BytesToHash(keyBytes[:]), currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opJumpDestProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeStackOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	if vmerr != nil {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	return osp, nil
}

// Call opcode proofs

func (osp *OneStepProof) addCallPostProof(ctx ProofGenContext, currState *state.IntraState) error {
	addrBytes := currState.Stack.Back(1).Bytes32()
	address := common.BytesToAddress(addrBytes[:])
	var in, inSize uint64
	if currState.OpCode == vm.CALL || currState.OpCode == vm.CALLCODE {
		in = currState.Stack.Back(3).Uint64()
		inSize = currState.Stack.Back(4).Uint64()
	} else if currState.OpCode == vm.DELEGATECALL || currState.OpCode == vm.STATICCALL {
		in = currState.Stack.Back(2).Uint64()
		inSize = currState.Stack.Back(3).Uint64()
	} else {
		log.Crit("Unreachable")
	}
	// For input data
	err := osp.addMemoryReadProof(in, inSize, currState)
	if err != nil {
		return err
	}
	// We don't need to prove the recepient account and the opcode for precompiles
	if _, ok := precompile(ctx.rules, address); !ok {
		// For recipient address
		if currState.OpCode == vm.CALL {
			value := currState.Stack.Back(2)
			if value.Sign() != 0 && currState.Caller != address {
				currState.GlobalState.SubBalance(currState.Caller, value.ToBig())
				currState.GlobalState.CommitForProof()
				err = osp.addAccountProof(currState, address)
				if err != nil {
					return err
				}
			}
		}
		// For opcode proof
		osp.addCodeProof(currState, address)
	}
	return nil
}

func (osp *OneStepProof) addCreatePostProof(currState *state.IntraState, nonce uint64, contractAddr common.Address) error {
	value := currState.Stack.Peek()
	in := currState.Stack.Back(1).Uint64()
	inSize := currState.Stack.Back(2).Uint64()
	err := osp.addMemoryReadProof(in, inSize, currState)
	if err != nil {
		return err
	}
	currState.GlobalState.SetNonce(currState.Caller, nonce+1)
	currState.GlobalState.SubBalance(currState.Caller, value.ToBig())
	currState.GlobalState.CommitForProof()
	return osp.addAccountProof(currState, contractAddr)
}

func opCreateProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(3, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCurrAccountProof(currState)
	if err != nil {
		return nil, err
	}
	if vmerr == vm.ErrInsufficientBalance || vmerr == vm.ErrNonceUintOverflow || vmerr == vm.ErrContractAddressCollision {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	nonce := currState.GlobalState.GetNonce(currState.Caller)
	contractAddr := crypto.CreateAddress(currState.Caller, nonce)
	err = osp.addCreatePostProof(currState, nonce, contractAddr)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCallProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(7, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCurrAccountProof(currState)
	if err != nil {
		return nil, err
	}
	if vmerr == vm.ErrInsufficientBalance {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err = osp.addCallPostProof(ctx, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCallCodeProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(7, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCallPostProof(ctx, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

// If a call frame is returned to its parent call frame, we need to provide the proof
// of the lastDepthState. The return data or contract creation is handled in [opReturnProof].
func (osp *OneStepProof) addCallReturnProof(currState *state.IntraState) error {
	lastDepthState := currState.LastDepthState.(*state.IntraState)
	osp.addStateProof(lastDepthState)
	return nil
}

func (osp *OneStepProof) addSelfDestructSetProof(currState *state.IntraState) {
	osp.AddProof(SelfDestructSetProofFromSelfDestructSet(currState.SelfDestructSet))
}

// If a transaction is returned, transaction needs to be finalized.
//  1. refund gas
//  2. tip the coinbase
//  3. delete all suicided contracts
//  4. update the transaction trie and receipt trie
func (osp *OneStepProof) addTransactionReturnProof(ctx ProofGenContext, currState *state.IntraState) error {
	lastDepthState := currState.LastDepthState.(*state.InterState)
	osp.addInterStateProof(lastDepthState)

	err := osp.addAccountProof(currState, currState.Caller)
	if err != nil {
		return err
	}
	refundQuotient := params.RefundQuotient
	if ctx.rules.IsLondon {
		refundQuotient = params.RefundQuotientEIP3529
	}
	initialGas := ctx.transaction.Gas()
	gasUsed := ctx.receipt.GasUsed
	gasPrice := ctx.transaction.GasPrice()
	gasLeft := initialGas - gasUsed
	refund := gasUsed / refundQuotient
	if refund > currState.Refund {
		refund = currState.Refund
	}
	gasLeft += refund
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(gasLeft), gasPrice)
	currState.GlobalState.AddBalance(currState.Caller, remaining)
	currState.GlobalState.CommitForProof()

	err = osp.addAccountProof(currState, ctx.coinbase)
	if err != nil {
		return err
	}
	fee := new(big.Int).Mul(new(big.Int).SetUint64(gasUsed), gasPrice)
	currState.GlobalState.AddBalance(ctx.coinbase, fee)
	currState.GlobalState.CommitForProof()

	osp.addSelfDestructSetProof(currState)
	for _, addr := range currState.SelfDestructSet.Contracts {
		err = osp.addAccountProof(currState, addr)
		if err != nil {
			return err
		}
		currState.GlobalState.DeleteSuicidedAccountForProof(addr)
		currState.GlobalState.CommitForProof()
	}

	osp.AddProof(ReceiptProofFromReceipt(ctx.receipt))
	return nil
}

func (osp *OneStepProof) addReturnProof(ctx ProofGenContext, currState, nextState *state.IntraState) error {
	if nextState == nil {
		// We're in the topmost call frame, the entire transaction will be finalized
		return osp.addTransactionReturnProof(ctx, currState)
	}
	return osp.addCallReturnProof(currState)
}

func opStopProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// TODO: EIP-3541 invalid code check in London
	if vmerr != nil && !IsStopTokenError(vmerr) {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	err := osp.addReturnProof(ctx, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opReturnProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	// TODO: EIP-3541 invalid code check in London
	if vmerr != nil && !IsStopTokenError(vmerr) {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	if currState.Depth > 1 || currState.CallFlag.IsCreate() {
		offset := currState.Stack.Peek().Uint64()
		size := currState.Stack.Back(1).Uint64()
		err := osp.addMemoryReadProof(offset, size, currState)
		if err != nil {
			return nil, err
		}
		if currState.CallFlag.IsCreate() {
			err = osp.addCurrAccountProof(currState)
			if err != nil {
				return nil, err
			}
			// Simulate contract creation
			createdCode := currState.Memory.Data()[offset : offset+size]
			currState.GlobalState.SetCode(currState.ContractAddress, createdCode)
			currState.GlobalState.CommitForProof()
		}
	}
	err := osp.addReturnProof(ctx, currState, nextState)
	if currState.Depth > 1 {
		if !currState.CallFlag.IsCreate() {
			lastDepthState := currState.LastDepthState.(*state.IntraState)
			err = osp.addMemoryWriteProof(currState.Out, currState.OutSize, lastDepthState, nextState)
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opDelegateCallProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(6, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCallPostProof(ctx, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opCreate2Proof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(4, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCurrAccountProof(currState)
	if err != nil {
		return nil, err
	}
	if vmerr == vm.ErrInsufficientBalance || vmerr == vm.ErrNonceUintOverflow || vmerr == vm.ErrContractAddressCollision {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	nonce := currState.GlobalState.GetNonce(currState.Caller)
	// TODO: revise it
	contractAddr := nextState.ContractAddress
	err = osp.addCreatePostProof(currState, nonce, contractAddr)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opStaticCallProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	// If the stack validation fails or write protection or depth, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection || vmerr == vm.ErrDepth {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(6, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addCallPostProof(ctx, currState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opRevertProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(2, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	if currState.Depth > 1 {
		lastDepthState := currState.LastDepthState.(*state.IntraState)
		offset := currState.Stack.Peek().Uint64()
		size := currState.Stack.Back(1).Uint64()
		err := osp.addMemoryReadProof(offset, size, currState)
		if err != nil {
			return nil, err
		}
		err = osp.addMemoryWriteProof(currState.Out, currState.OutSize, lastDepthState, nextState)
		if err != nil {
			return nil, err
		}
	}
	err := osp.addRevertProof(ctx, currState, nextState, errors.New("execution reverted"))
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opInvalidProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeInvalidOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	err := osp.addRevertProof(ctx, currState, nextState, vmerr)
	if err != nil {
		return nil, err
	}
	return osp, nil
}

func opSelfDestructProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeCallOp)
	osp.addStateProof(currState)
	osp.addOpCodeProof(ctx, currState)
	// If the stack validation fails or write protection, we don't need to provide stack proofs
	if IsStackError(vmerr) || vmerr == vm.ErrWriteProtection {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
		return osp, nil
	}
	osp.addStackProof(1, currState)
	if vmerr == vm.ErrOutOfGas || vmerr == vm.ErrGasUintOverflow {
		err := osp.addRevertProof(ctx, currState, nextState, vmerr)
		if err != nil {
			return nil, err
		}
	}
	beneficiary := common.Address(currState.Stack.Back(0).Bytes20())
	contract := currState.ContractAddress
	err := osp.addAccountProof(currState, contract)
	if err != nil {
		return nil, err
	}
	if currState.GlobalState.GetBalance(contract).Sign() > 0 && beneficiary != contract {
		currState.GlobalState.SubBalance(contract, currState.GlobalState.GetBalance(contract))
		currState.GlobalState.CommitForProof()
		err = osp.addAccountProof(currState, beneficiary)
		if err != nil {
			return nil, err
		}
	}
	err = osp.addReturnProof(ctx, currState, nextState)
	if err != nil {
		return nil, err
	}
	return osp, nil
}
