package services

import (
	"context"
	"math/big"

	l1abi "github.com/ethereum/go-ethereum/accounts/abi"
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
	prevBisection [][32]byte,
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
		prevBisection,
		prevChallengedSegmentStart,
		prevChallengedSegmentLength,
	)
	log.Info("OSP submitted")
	if err != nil {
		log.Error("OSP verification failed")
	}
	return err
}

func RespondBisection(
	b *BaseService,
	abi *l1abi.ABI,
	challengeSession *bindings.IChallengeSession,
	ev *bindings.IChallengeBisected,
	states []*proof.ExecutionState,
	opponentEndStateHash common.Hash,
	isSequencer bool,
) error {
	// Get bisection info from event
	segStart := ev.ChallengedSegmentStart.Uint64()
	segLen := ev.ChallengedSegmentLength.Uint64()
	// Get previous bisections from call data
	tx, _, err := b.L1.TransactionByHash(b.Ctx, ev.Raw.TxHash)
	if err != nil {
		// TODO: error handling
		log.Error("Failed to get challenge data", "error", err)
		return nil
	}
	decoded, err := abi.Methods["bisectExecution"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		if isSequencer {
			// Sequencer always start first
			log.Error("Failed to decode bisection data", "error", err)
			return nil
		}
		// We are in the first round when the defender calls initializeChallengeLength
		// Get initialized challenge length from event
		steps := segLen
		if steps != uint64(len(states)-1) {
			log.Crit("UNHANDELED: currently not support diverge on steps")
		}
		prevBisection := [][32]byte{
			states[0].Hash(),
			opponentEndStateHash,
		}
		if segLen == 1 {
			// This assertion only has one step
			err = SubmitOneStepProof(
				challengeSession,
				b.ProofBackend,
				b.Ctx,
				states[0],
				common.Big1,
				prevBisection,
				ev.ChallengedSegmentStart,
				ev.ChallengedSegmentLength,
			)
			if err != nil {
				log.Crit("UNHANDELED: osp failed")
			}
		} else {
			// This assertion has multiple steps
			startState := states[0].Hash()
			midState := states[steps/2+steps%2].Hash()
			endState := states[steps].Hash()
			bisection := [][32]byte{
				startState,
				midState,
				endState,
			}
			_, err := challengeSession.BisectExecution(
				bisection,
				common.Big1,
				prevBisection,
				ev.ChallengedSegmentStart,
				ev.ChallengedSegmentLength,
			)
			log.Info("BisectExecution", "bisection", bisection, "cidx", common.Big1, "psegStart", segStart, "psegLen", segLen, "prev", prevBisection)
			if err != nil {
				log.Crit("UNHANDELED: bisection excution failed", "err", err)
			}
		}
		return nil
	}
	prevBisection := decoded[0].([][32]byte)
	startState := states[segStart].Hash()
	midState := states[segStart+segLen/2+segLen%2].Hash()
	endState := states[segStart+segLen].Hash()
	if segLen == 1 {
		// We've reached one step
		err = SubmitOneStepProof(
			challengeSession,
			b.ProofBackend,
			b.Ctx,
			states[segStart],
			common.Big1,
			prevBisection,
			ev.ChallengedSegmentStart,
			ev.ChallengedSegmentLength,
		)
		if err != nil {
			log.Crit("UNHANDELED: osp failed")
		}
	} else {
		challengeIdx := uint64(1)
		if prevBisection[1] == midState {
			challengeIdx = 2
		}
		if segLen == 2 || (segLen == 3 && challengeIdx == 2) {
			// The next challenge segment is a single step
			stateIndex := segStart
			if challengeIdx != 1 {
				stateIndex = segStart + segLen/2
			}
			err = SubmitOneStepProof(
				challengeSession,
				b.ProofBackend,
				b.Ctx,
				states[stateIndex],
				new(big.Int).SetUint64(challengeIdx),
				prevBisection,
				ev.ChallengedSegmentStart,
				ev.ChallengedSegmentLength,
			)
			if err != nil {
				log.Crit("UNHANDELED: osp failed")
			}
		} else {
			var newLen uint64 // New segment length
			var bisection [][32]byte
			if challengeIdx == 1 {
				newLen = segLen/2 + segLen%2
				bisection = [][32]byte{
					startState,
					states[segStart+newLen/2+newLen%2].Hash(),
					midState,
				}
			} else {
				newLen = segLen / 2
				bisection = [][32]byte{
					midState,
					states[segStart+segLen/2+segLen%2+newLen/2+newLen%2].Hash(),
					endState,
				}
			}
			_, err := challengeSession.BisectExecution(
				bisection,
				new(big.Int).SetUint64(challengeIdx),
				prevBisection,
				ev.ChallengedSegmentStart,
				ev.ChallengedSegmentLength,
			)
			log.Info("BisectExecution", "bisection", bisection, "cidx", challengeIdx, "psegStart", segStart, "psegLen", segLen, "prev", prevBisection)
			if err != nil {
				log.Crit("UNHANDELED: bisection excution failed", "err", err)
			}
		}
	}
	return nil
}
