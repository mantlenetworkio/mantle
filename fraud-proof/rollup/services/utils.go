package services

import (
	"bytes"
	"context"
	"math/big"

	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

func SubmitOneStepProof(
	challengeSession *bindings.IChallengeSession,
	proofBackend proof.Backend,
	ctx context.Context,
	state *proof.ExecutionState,
	challengedStepIndex *big.Int,
//prevBisection [][32]byte,
	prevChallengedSegmentStart *big.Int,
	prevChallengedSegmentLength *big.Int,
) error {
	osp, err := proof.GenerateProof(proofBackend, ctx, state, nil)
	if err != nil {
		log.Crit("UNHANDELED: osp generation failed", "err", err)
	}
	_, err = challengeSession.VerifyOneStepProof(
		osp.Encode(),
		challengedStepIndex,
		//prevBisection,
		prevChallengedSegmentStart,
		prevChallengedSegmentLength,
	)
	log.Info("OSP submitted")
	if err != nil {
		log.Error("OSP verification failed")
	}
	return err
}

// Responder -> startStateHash, endStateHash
func RespondBisection(
	b *BaseService,
	challengeSession *bindings.IChallengeSession,
	ev *bindings.IChallengeBisected,
	states []*proof.ExecutionState,
) error {
	//var prevBisection = [2][32]byte{ev.StartState, ev.EndState}
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
			state = states[segStart+segLen/2]
		} else if !bytes.Equal(endState[:], ev.EndState[:]) {
			state = states[segStart+segLen]
		} else {
			log.Crit("RespondBisection can't find state difference")
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
