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
	"math/big"

	"github.com/mantlenetworkio/mantle/fraud-proof/proof/state"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/params"
)

type ProofGenContext struct {
	rules       params.Rules
	coinbase    common.Address
	transaction *types.Transaction
	receipt     *types.Receipt
	actualCode  []byte // For opcode proof when in CALLCODE/DELEGATECALL/STATICCALL
}

func NewProofGenContext(rules params.Rules, coinbase common.Address, transaction *types.Transaction, receipt *types.Receipt, actualCode []byte) ProofGenContext {
	return ProofGenContext{
		rules:       rules,
		coinbase:    coinbase,
		receipt:     receipt,
		transaction: transaction,
		actualCode:  actualCode,
	}
}

type Proof interface {
	Encode() []byte
}

type OneStepProof struct {
	VerifierType VerifierType
	Proofs       []Proof

	TotalCodeSize uint64 // for statistics
}

func EmptyProof() *OneStepProof {
	return &OneStepProof{}
}

func (p *OneStepProof) AddProof(proof Proof) {
	p.Proofs = append(p.Proofs, proof)
}

func (p *OneStepProof) SetVerifierType(ty VerifierType) {
	p.VerifierType = ty
}

func (p *OneStepProof) Encode() []byte {
	if len(p.Proofs) == 0 {
		// Empty proof!
		return []byte{}
	}
	encodedLen := 0
	encodedProofs := make([][]byte, len(p.Proofs))
	for idx, proof := range p.Proofs {
		encodedProofs[idx] = proof.Encode()
		encodedLen += len(encodedProofs[idx])
	}
	encoded := make([]byte, encodedLen)
	offset := 0
	for _, encodedProof := range encodedProofs {
		copy(encoded[offset:], encodedProof)
		offset += len(encodedProof)
	}
	return encoded
}

// Type 1 BlockState -> InterState: block initiation
// A BlockStateProof suffices to prove this state transition.
//
// The verifier can generate the first InterState by setting the transaction trie root
// and receipt trie root to the empty trie root, logs bloom to empty, and setting the
// block gas used to 0.
func GetBlockInitiationProof(blockState *state.BlockState) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeBlockInit)
	osp.AddProof(BlockStateProofFromBlockState(blockState))
	return osp, nil
}

// Type 6 InterState -> BlockState: block finalization
// The verifier needs to first pack the block header and calculate the block hash;
// then, the verifier can simulate changing the block hash in the block hash tree
// using the plain Merkle tree proof.
func GetBlockFinalizationProof(interState *state.InterState) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeBlockFinal)
	// This proof reveals the transaction trie root, receipt trie root, logs bloom,
	// and block gas used. Verifier can calculate the block hash from these values.
	osp.AddProof(InterStateProofFromInterState(interState))
	// Prove the parent block hash
	osp.AddProof(&BlockHashProof{interState.BlockHashTree.GetBlockHash(interState.BlockNumber - 1)})
	// This proof provides the block hash tree Merkle proof of the parent block.
	blockHashMerkleProof, err := GetBlockHashMerkleProof(interState.BlockHashTree, interState.BlockNumber-1)
	if err != nil {
		return nil, err
	}
	osp.AddProof(blockHashMerkleProof)
	// Prove the current block hash to be updated
	osp.AddProof(&BlockHashProof{interState.BlockHashTree.GetBlockHash(interState.BlockNumber)})
	// This proof provides the block hash tree Merkle proof at the designated index.
	blockHashMerkleProof, err = GetBlockHashMerkleProof(interState.BlockHashTree, interState.BlockNumber)
	if err != nil {
		return nil, err
	}
	osp.AddProof(blockHashMerkleProof)
	return osp, nil
}

// Type 2 InterState -> IntraState: transaction initiation
// Type 3 InterState -> InterState: EOA transfer transaction
// See go-ethereum/core/state_transition.go for the full state transition logic.
//  1. Provide the inter state proof.
//  2. Provide the MPT proof for the sender account. Verifier can perform nonce
//     check and balance check through the MPT proof.
//  3. Simulate the transaction initation including nonce increment, gas buying,
//     and balance transfer. If error, stop.
//     Note: sadly EVM does not expose methods for these so we have to implement
//     them by ourselves.
//  4. Provide the MPT proof for the recipient account, based on the simulated
//     state. Verifier can construct the state trie root after transfer.
//  5. If EOA transfer, also provde the MPT proof for both transaction trie, receipt
//     trie, and account proof of coinbase so verifier can finalize the transaction.
func GetTransactionInitaitionProof(
	chainConfig *params.ChainConfig,
	vmctx *vm.BlockContext,
	tx *types.Transaction,
	txctx *vm.Context,
	interState *state.InterState,
	statedb vm.StateDB,
) (*OneStepProof, error) {
	osp := EmptyProof()
	osp.SetVerifierType(VerifierTypeInterTx)
	osp.AddProof(InterStateProofFromInterState(interState))

	// MPT proof of the sender account
	senderProof, err := statedb.GetProof(txctx.Origin)
	if err != nil {
		return nil, err
	}
	osp.AddProof(&MPTProof{senderProof})

	// Simulate the transaction initiation
	success := true
	// 1. Nonce check
	stNonce := statedb.GetNonce(txctx.Origin)
	msgNonce := tx.Nonce()
	if stNonce < msgNonce || stNonce > msgNonce || stNonce+1 < stNonce {
		// Nonce check failed
		success = false
	}
	// 2. Sender EOA check
	if success && len(statedb.GetCode(txctx.Origin)) != 0 {
		// Sender is a contract
		success = false
	}
	// 3. Gas buying
	if success {
		requiredBalance := new(big.Int).SetUint64(tx.Gas())
		requiredBalance = requiredBalance.Mul(requiredBalance, tx.GasPrice())
		stBalance := statedb.GetBalance(txctx.Origin)
		if stBalance.Cmp(requiredBalance) < 0 {
			success = false
		}
		// Here we don't check if there is enough gas left in the block for the transaction.
		// Only happens when the sequencer is malicious.
		// TODO: reason it or check
		if success {
			statedb.SubBalance(txctx.Origin, requiredBalance)
		}
	}
	var rules params.Rules
	// 4. Instrinsic gas deduction
	if success {
		rules = chainConfig.Rules(vmctx.BlockNumber)
		// TODO: can we just hard-code the consensus rules here?
		gas, err := core.IntrinsicGas(tx.Data(), tx.To() == nil, rules.IsHomestead, rules.IsIstanbul)
		if err != nil {
			// Gas overflow
			// TODO: reason it
			success = false
		} else if tx.Gas() < gas {
			// ErrIntrinsicGas
			success = false
		}
	}
	// 5. Balance transfer check
	if success {
		if tx.Value().Sign() > 0 && !vmctx.CanTransfer(statedb, txctx.Origin, tx.Value()) {
			// ErrInsufficientFunds
			success = false
		}
	}
	// 6. Simulate the transaction initiation on the sender account
	if success {
		// Increment nonce
		statedb.SetNonce(txctx.Origin, stNonce+1)
		// Deduct sender's balance
		statedb.SubBalance(txctx.Origin, tx.Value())
		// Commit for new trie root after sender account changes
		statedb.CommitForProof()
		// Generate proof for the recipient account
		if tx.To() != nil {
			// Call or EOA transfer
			recipientProof, err := statedb.GetProof(*tx.To())
			if err != nil {
				return nil, err
			}
			osp.AddProof(&MPTProof{recipientProof})
			// Add recipient's balance
			statedb.AddBalance(*tx.To(), tx.Value())
		} else {
			// Contract creation
			contractAddr := crypto.CreateAddress(txctx.Origin, stNonce) // stNonce is the nonce before increment
			recipientProof, err := statedb.GetProof(contractAddr)
			if err != nil {
				return nil, err
			}
			osp.AddProof(&MPTProof{recipientProof})
			// Add contract's balance only if no address collision
			if statedb.GetNonce(contractAddr) != 0 || len(statedb.GetCode(contractAddr)) != 0 {
				success = false
			} else {
				// Set contract nonce
				if rules.IsEIP158 {
					statedb.SetNonce(contractAddr, 1)
				}
				// Add recipient's balance
				statedb.AddBalance(contractAddr, tx.Value())
			}
		}
	}
	// 7. Commit statedb for the recipient account changes
	if success {
		// TODO: EOA transfer can't be contract creation, refactor
		statedb.CommitForProof()
	}

	// Determine if the transaction is an EOA transfer
	// 1. The transaction canot be a contract creation
	// 2. The recipient's code length should be 0
	if tx.To() != nil && len(statedb.GetCode(*tx.To())) == 0 {
		// This is an EOA transfer, provide proofs for finalization here
		// Note: transaction/receipt trie does not contain the current transaction
		//       so it is acutally proofs of exclusions. The verifier can perform
		//       insertions based on the proof.
		txProof, err := interState.TransactionTrie.Prove(int(interState.TransactionIdx))
		if err != nil {
			return nil, err
		}
		osp.AddProof(&MPTProof{txProof})
		receiptProof, err := interState.ReceiptTrie.Prove(int(interState.TransactionIdx))
		if err != nil {
			return nil, err
		}
		osp.AddProof(&MPTProof{receiptProof})
		if success {
			// Only provide the coinbase proof if the transaction is successful started
			coinbaseProof, err := statedb.GetProof(vmctx.Coinbase)
			if err != nil {
				return nil, err
			}
			osp.AddProof(&MPTProof{coinbaseProof})
		}
	}
	return osp, nil
}

// Type 4 IntraState -> IntraState: one-step EVM execution
// Type 5 IntraState -> InterState: transaction finalization
func GetIntraProof(ctx ProofGenContext, currState, nextState *state.IntraState, vmerr error) (*OneStepProof, error) {
	// log.Info("Generating intra proof", "op", currState.OpCode, "gen", proofJumpTable[currState.OpCode])
	return proofJumpTable[currState.OpCode].genProof(ctx, currState, nextState, vmerr)
}
