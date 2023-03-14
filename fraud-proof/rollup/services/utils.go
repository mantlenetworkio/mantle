package services

import (
	"bytes"
	"context"
	"fmt"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"math/big"

	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

func SubmitOneStepProof(
	challengeSession *bindings.ChallengeSession,
	proofBackend proof.Backend,
	ctx context.Context,
	state *proof.ExecutionState,
	challengedStepIndex *big.Int,
	prevChallengedSegmentStart *big.Int,
	prevChallengedSegmentLength *big.Int,
) error {
	log.Info("OSP GenerateProof...")
	osp, err := proof.GenerateProof(ctx, proofBackend, state, nil)
	if err != nil {
		log.Error("UNHANDELED: osp generation failed", "err", err)
		return err
	}
	log.Info("OSP GenerateProof success")

	log.Info("OSP BuildVerificationContext...")
	verificationContext, err := BuildVerificationContext(ctx, proofBackend, state)
	if err != nil {
		log.Error("UNHANDELED: osp build verification context failed", "err", err)
		return err
	}
	log.Info("OSP BuildVerificationContext success")

	log.Info("OSP VerifyOneStepProof...")
	_, err = challengeSession.VerifyOneStepProof(
		*verificationContext,
		uint8(osp.VerifierType),
		osp.Encode(),
		prevChallengedSegmentStart,
		prevChallengedSegmentLength,
	)
	if err != nil {
		log.Error("OSP verification failed")
		return err
	}
	log.Info("OSP VerifyOneStepProof submitted")
	return nil
}

// Responder -> startStateHash, endStateHash
func RespondBisection(
	b *BaseService,
	challengeSession *bindings.ChallengeSession,
	ev *bindings.ChallengeBisected,
	states []*proof.ExecutionState,
) error {
	var bisection [3][32]byte
	var challengeIdx uint64
	var newStart uint64
	var newLen uint64 // New segment length
	var err error
	// Get bisection info from event
	segStart := ev.ChallengedSegmentStart.Uint64()
	segLen := ev.ChallengedSegmentLength.Uint64()

	if segStart+segLen >= uint64(len(states)) {
		log.Crit("RespondBisection out of range", "segStart+segLen", segStart+segLen, "len(states)", len(states))
	}

	startState := states[segStart].Hash()
	midState := states[segStart+segLen/2+segLen%2].Hash()
	endState := states[segStart+segLen].Hash()

	if segLen >= 3 {
		if !bytes.Equal(midState[:], ev.MidState[:]) {
			newStart = segStart
			newLen = segLen/2 + segLen%2
			bisection[0] = startState
			bisection[1] = states[segStart+newLen/2+newLen%2].Hash()
			bisection[2] = midState
			challengeIdx = 1
		} else {
			newStart = segStart + segLen/2 + segLen%2
			newLen = segLen / 2
			bisection[0] = midState
			bisection[1] = states[segStart+segLen/2+segLen%2+newLen/2+newLen%2].Hash()
			bisection[2] = endState
			challengeIdx = 2
		}
	} else if segLen <= 2 && segLen > 0 {
		var state *proof.ExecutionState
		if !bytes.Equal(startState[:], ev.StartState[:]) {
			state = states[segStart]
		} else if !bytes.Equal(midState[:], ev.MidState[:]) {
			state = states[segStart+segLen/2+segLen%2]
		} else if !bytes.Equal(endState[:], ev.EndState[:]) {
			state = states[segStart+segLen]
		} else {
			log.Error("RespondBisection can't find state difference")
			return nil
		}

		// We've reached one step
		err = SubmitOneStepProof(
			challengeSession,
			b.ProofBackend,
			b.Ctx,
			state,
			common.Big1,
			ev.ChallengedSegmentStart,
			ev.ChallengedSegmentLength,
		)
		if err != nil {
			log.Crit("UNHANDELED: osp failed")
		}
		return err
	} else {
		log.Crit("RespondBisection segLen in event is illegal")
	}
	log.Info("BisectExecution", "bisection", bisection, "cidx", challengeIdx, "psegStart", segStart, "psegLen", segLen)
	_, err = challengeSession.BisectExecution(
		bisection,
		new(big.Int).SetUint64(challengeIdx),
		new(big.Int).SetUint64(newStart),
		new(big.Int).SetUint64(newLen),
		ev.ChallengedSegmentStart,
		ev.ChallengedSegmentLength,
	)
	if err != nil {
		log.Crit("UNHANDELED: bisection excution failed", "err", err)
	}
	return nil
}

// MidState mid-states with floor index
func MidState(states []*proof.ExecutionState, segStart, segLen uint64) common.Hash {
	return states[segStart+segLen/2].Hash()
}

func BuildVerificationContext(ctx context.Context, proofBackend proof.Backend, state *proof.ExecutionState) (*bindings.VerificationContextContext, error) {
	var evmTx bindings.EVMTypesLibTransaction
	var tx *types.Transaction
	var header *types.Header
	var err error
	// get block
	if state != nil && state.Block != nil {
		header = state.Block.Header()
	} else {
		return nil, fmt.Errorf("get nil block from ExecutionState status")
	}
	// get transaction
	if state != nil && state.Block.Transactions() != nil {
		txs := state.Block.Transactions()
		if uint64(len(txs)) < state.TransactionIdx+1 {
			return nil, fmt.Errorf("get transaction index from ExecutionState out of range")
		}
		tx = state.Block.Transactions()[state.TransactionIdx]
	} else {
		return nil, fmt.Errorf("get nil transactions from ExecutionState status")
	}
	// build EVMTypesLibTransaction
	var txOrigin common.Address
	evmTx.Nonce = tx.Nonce()
	evmTx.GasPrice = tx.GasPrice()
	evmTx.To = ethc.Address(*tx.To())
	evmTx.Value = tx.Value()
	evmTx.Data = tx.Data()
	if tx.QueueOrigin() == types.QueueOriginSequencer {
		evmTx.V, evmTx.R, evmTx.S = tx.RawSignatureValues()
		signer := types.NewEIP155Signer(tx.ChainId())
		txOrigin, err = types.Sender(signer, tx)
		if err != nil {
			return nil, err
		}
	} else {

	}
	return &bindings.VerificationContextContext{
		Coinbase:    ethc.Address(header.Coinbase),
		Timestamp:   new(big.Int).SetUint64(tx.L1Timestamp()),
		Number:      header.Number,
		Origin:      ethc.Address(txOrigin),
		Transaction: evmTx,
		InputRoot:   [32]byte{0},
		TxHash:      tx.Hash(),
	}, nil
}
