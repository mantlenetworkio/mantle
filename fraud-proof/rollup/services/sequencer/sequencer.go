package sequencer

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/p2p"
)

func RegisterService(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) {
	sequencer, err := New(eth, proofBackend, cfg, auth)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	sequencer.Start()
	log.Info("Sequencer registered")
}

type ChallengeCtx struct {
	ChallengeAddr common.Address
	Assertion     *rollupTypes.Assertion
	Parent        *rollupTypes.Assertion
}

// Sequencer run confirming loop and respond challenge, assumes no Berlin+London fork on L2
type Sequencer struct {
	*services.BaseService

	confirmedIDCh         chan *big.Int
	challengeCh           chan *ChallengeCtx
	challengeResolutionCh chan struct{}

	confirmations uint64
}

func New(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) (*Sequencer, error) {
	base, err := services.NewBaseService(eth, proofBackend, cfg, auth)
	if err != nil {
		return nil, err
	}
	s := &Sequencer{
		BaseService:           base,
		confirmedIDCh:         make(chan *big.Int, 4096),
		challengeCh:           make(chan *ChallengeCtx),
		challengeResolutionCh: make(chan struct{}),
		confirmations:         cfg.L1Confirmations,
	}
	return s, nil
}

// This goroutine tries to confirm created assertions
func (s *Sequencer) confirmationLoop() {
	defer s.Wg.Done()

	// Watch AssertionCreated event
	var createdCh = make(chan *bindings.RollupAssertionCreated, 4096)
	createdSub, err := s.Rollup.Contract.WatchAssertionCreated(&bind.WatchOpts{Context: s.Ctx}, createdCh)
	if err != nil {
		log.Error("Failed to watch rollup event", "err", err)
	}
	defer createdSub.Unsubscribe()

	// Watch AssertionConfirmed event
	var confirmedCh = make(chan *bindings.RollupAssertionConfirmed, 4096)
	confirmedSub, err := s.Rollup.Contract.WatchAssertionConfirmed(&bind.WatchOpts{Context: s.Ctx}, confirmedCh)
	if err != nil {
		log.Error("Failed to watch rollup event", "err", err)
	}
	defer confirmedSub.Unsubscribe()

	// Watch L1 blockchain for confirmation period
	var headCh = make(chan *types.Header, 4096)
	headSub, err := s.L1.SubscribeNewHead(s.Ctx, headCh)
	if err != nil {
		log.Error("Failed to watch l1 chain head", "err", err)
	}
	defer headSub.Unsubscribe()

	var challengedCh = make(chan *bindings.RollupAssertionChallenged, 4096)
	challengedSub, err := s.Rollup.Contract.WatchAssertionChallenged(&bind.WatchOpts{Context: s.Ctx}, challengedCh)
	if err != nil {
		log.Error("Failed to watch rollup event", "err", err)
	}
	defer challengedSub.Unsubscribe()

	challengeAssertions := make(map[uint64]bool)

	challengeContext, _ := s.Rollup.ChallengeCtx()
	isInChallenge := challengeContext.DefenderAssertionID.Uint64() != 0 && !challengeContext.Completed
	// restart with challengeCtx
	if isInChallenge {
		defenderAssertion, _ := s.AssertionMap.Assertions(challengeContext.DefenderAssertionID)
		parentAssertionID, _ := s.AssertionMap.GetParentID(challengeContext.DefenderAssertionID)
		parentAssertion, _ := s.AssertionMap.Assertions(parentAssertionID)

		challengeCtx := ChallengeCtx{
			ChallengeAddr: common.Address(challengeContext.ChallengeAddress),
			Assertion: &rollupTypes.Assertion{
				ID:           challengeContext.DefenderAssertionID,
				VmHash:       defenderAssertion.StateHash,
				InboxSize:    defenderAssertion.InboxSize,
				Parent:       defenderAssertion.Parent,
				Deadline:     defenderAssertion.Deadline,
				ProposalTime: defenderAssertion.ProposalTime,
			},
			Parent: &rollupTypes.Assertion{
				ID:           parentAssertionID,
				VmHash:       parentAssertion.StateHash,
				InboxSize:    parentAssertion.InboxSize,
				Parent:       parentAssertion.Parent,
				Deadline:     parentAssertion.Deadline,
				ProposalTime: parentAssertion.ProposalTime,
			},
		}
		s.challengeCh <- &challengeCtx
	}

	for {
		if isInChallenge {
			// Waif for the challenge resolved
			select {
			case <-s.challengeResolutionCh:
				log.Info("Sequencer finished challenge, reset isInChallenge status")
				isInChallenge = false
			case <-s.Ctx.Done():
				log.Error("Scheduler confirmationLoop ctx done")
				return
			}
		} else {
			select {
			case ev := <-createdCh:
				// New assertion created on L1 Rollup
				log.Info("Get New Assertion...", "AssertionID", ev.AssertionID,
					"AsserterAddr", ev.AsserterAddr, "VmHash", ev.VmHash, "InboxSize", ev.InboxSize)
				if !bytes.Equal(ev.AsserterAddr.Bytes(), s.Config.StakeAddr.Bytes()) {
					log.Info("Get Assertion for challenge,store it")
					challengeAssertions[ev.AssertionID.Uint64()] = true
				}
			case header := <-headCh:
				// todo : optimization the check with block height
				// Get confirm block header
				if s.confirmations != 0 {
					num := new(big.Int)
					num.SetUint64(s.confirmations)
					num.Sub(header.Number, num)
					header, err = s.L1.HeaderByNumber(s.Ctx, num)
					if err != nil {
						log.Error("Failed to get confirmed header", "err", err)
						continue
					}
				}
				// Get first unresolved confirm assertion and check deadline
				lastRSAID, err := s.Rollup.LastResolvedAssertionID()
				if err != nil {
					log.Error("Failed to get last resolved assertion ID", "err", err)
					continue
				}
				lastCAID, err := s.Rollup.LastCreatedAssertionID()
				if err != nil {
					log.Error("Failed to get last created assertion ID", "err", err)
					continue
				}
				if lastCAID.Uint64() <= lastRSAID.Uint64() {
					continue
				}
				// New block mined on L1
				log.Info("Sequencer sync new layer1 block", "height", header.Number)
				firstUnresolvedID := lastRSAID.Uint64() + 1
				firstUnresolvedAssertion, err := s.AssertionMap.Assertions(new(big.Int).SetUint64(firstUnresolvedID))
				if err != nil {
					log.Error("Failed to get first unresolved Assertion", "err", err, "ID", firstUnresolvedID)
					continue
				}
				if header.Time >= firstUnresolvedAssertion.Deadline.Uint64() {
					log.Info("Current assertion", "id", firstUnresolvedID)
					if !challengeAssertions[firstUnresolvedID] {
						// Confirmation period has past, confirm it
						log.Info("Sequencer call ConfirmFirstUnresolvedAssertion...")
						_, err := s.Rollup.ConfirmFirstUnresolvedAssertion()
						if err != nil {
							if err.Error() == "execution reverted: InvalidParent" {
								challengeAssertions[firstUnresolvedID] = true
							}
							log.Error("Failed to confirm DA", "err", err)
						}
						continue
					}
					log.Info("Sequencer call RejectFirstUnresolvedAssertion...")
					// reject challenge assertion
					_, err := s.Rollup.RejectFirstUnresolvedAssertion()
					if err != nil {
						log.Error("Failed to reject DA", "err", err)
						continue
					}
					_, err = s.Rollup.RemoveOldZombies()
					if err != nil {
						log.Error("Failed to remove zombies", "err", err)
						continue
					}
					delete(challengeAssertions, firstUnresolvedID)
				}

			case ev := <-confirmedCh:
				log.Info("New confirmed assertion", "id", ev.AssertionID)
			case ev := <-challengedCh:
				log.Warn("New challenge rise!!!!!!", "ev", ev)
				// todo when interrupt at this moment, check staker`s status
				challengeAssertion := new(rollupTypes.Assertion)
				parent := new(rollupTypes.Assertion)
				if ret, err := s.AssertionMap.Assertions(ev.AssertionID); err != nil {
					log.Crit("Get assertion failed", "id", ev.AssertionID, "err", err)
				} else {
					challengeAssertion.ID = ev.AssertionID
					challengeAssertion.VmHash = ret.StateHash
					challengeAssertion.InboxSize = ret.InboxSize
					challengeAssertion.Parent = ret.Parent
					challengeAssertion.Deadline = ret.Deadline
					challengeAssertion.ProposalTime = ret.ProposalTime
				}
				if ret, err := s.AssertionMap.Assertions(challengeAssertion.Parent); err != nil {
					log.Crit("Get assertion failed", "id", challengeAssertion.Parent, "err", err)
				} else {
					parent.ID = ev.AssertionID
					parent.VmHash = ret.StateHash
					parent.InboxSize = ret.InboxSize
					parent.Parent = ret.Parent
					parent.Deadline = ret.Deadline
					parent.ProposalTime = ret.ProposalTime
				}

				challengeCtx := ChallengeCtx{
					common.Address(ev.ChallengeAddr),
					challengeAssertion,
					parent,
				}

				s.challengeCh <- &challengeCtx
				isInChallenge = true

			case <-s.Ctx.Done():
				return
			}
		}
	}
}

func (s *Sequencer) challengeLoop() {
	defer s.Wg.Done()

	// Watch L1 blockchain for challenge timeout
	var headCh = make(chan *types.Header, 4096)
	headSub, err := s.L1.SubscribeNewHead(s.Ctx, headCh)
	if err != nil {
		log.Error("Failed to watch l1 chain head", "err", err)
	}
	defer headSub.Unsubscribe()

	var challengeSession *bindings.ChallengeSession
	var states []*proof.ExecutionState

	var bisectedCh = make(chan *bindings.ChallengeBisected, 4096)
	var bisectedSub event.Subscription
	var challengeCompletedCh = make(chan *bindings.ChallengeChallengeCompleted, 4096)
	var challengeCompletedSub event.Subscription

	inChallenge := false
	var opponentTimeout uint64

	for {
		if inChallenge {
			select {
			case ev := <-bisectedCh:
				// case get bisection, if is our turn
				//   if in single step, submit proof
				//   if multiple step, track current segment, update
				responder, err := challengeSession.CurrentResponder()
				if err != nil {
					// TODO: error handling
					log.Error("Can not get current responder", "error", err)
					continue
				}
				log.Info("Responder info...", "responder", responder, "staker", s.Config.StakeAddr)
				if common.Address(responder) == s.Config.StakeAddr {
					log.Info("Sequencer start to respond new bisection...")
					// If it's our turn
					//err := services.RespondBisection(s.BaseService, abi, challengeSession, ev, states, common.Hash{}, false)
					err := services.RespondBisection(s.BaseService, challengeSession, ev, states)
					if err != nil {
						// TODO: error handling
						log.Error("Can not respond to bisection", "error", err)
						continue
					}
				} else {
					opponentTimeLeft, err := challengeSession.CurrentResponderTimeLeft()
					if err != nil {
						// TODO: error handling
						log.Error("Can not get current responder left time", "error", err)
						continue
					}
					log.Info("[Sequencer] Opponent time left", "blockTime", ev.BlockTime.Uint64(), "timeLeft", opponentTimeLeft)
					opponentTimeout = ev.BlockTime.Uint64() + opponentTimeLeft.Uint64()
				}
			case header := <-headCh:
				if opponentTimeout == 0 {
					continue
				}
				// TODO: can we use >= here?
				log.Info("New header incoming...", "header.Number", header.Number, "header.Time", header.Time, "opponentTimeout", opponentTimeout)
				if header.Time > opponentTimeout {
					_, err = challengeSession.Timeout()
					if err != nil {
						log.Error("Can not timeout opponent", "error", err)
						continue
						// TODO: wait some time before retry
						// TODO: fix race condition
					}
					log.Info("Timeout challenge...")
				}
			case ev := <-challengeCompletedCh:
				// TODO: handle if we are not winner --> state corrupted
				log.Info("[challenge] Try to challenge completed", "winner", ev.Winner)
				_, err = challengeSession.CompleteChallenge(s.Config.ChallengeVerify)
				if err != nil {
					log.Error("Can not complete challenge", "error", err)
					continue
				}
				bisectedSub.Unsubscribe()
				challengeCompletedSub.Unsubscribe()
				states = []*proof.ExecutionState{}
				inChallenge = false
				challengeSession = nil
				s.challengeResolutionCh <- struct{}{}
				log.Info("[challenge] Challenge completed", "winner", ev.Winner)
			case <-s.Ctx.Done():
				bisectedSub.Unsubscribe()
				challengeCompletedSub.Unsubscribe()
				return
			}
		} else {
			select {
			case ctx := <-s.challengeCh:
				log.Warn("Sequencer receive new challenge!!!", "handle it", ctx)
				challenge, err := bindings.NewChallenge(ethc.Address(ctx.ChallengeAddr), s.L1)
				if err != nil {
					log.Crit("Failed to access ongoing challenge", "address", ctx.ChallengeAddr, "err", err)
				}
				challengeSession = &bindings.ChallengeSession{
					Contract:     challenge,
					CallOpts:     bind.CallOpts{Pending: true, Context: s.Ctx},
					TransactOpts: *s.TransactOpts,
				}
				// use staker status to check challenge status
				// 1. challenge contract not exist
				// 2. challenge exist and already completed
				stakeStatus, _ := s.Rollup.Stakers(s.Rollup.TransactOpts.From)
				if bytes.Equal(stakeStatus.CurrentChallenge.Bytes(), common.BigToAddress(common.Big0).Bytes()) {
					winner, err := challengeSession.Winner()
					if err != nil || bytes.Equal(winner.Bytes(), common.BigToAddress(common.Big0).Bytes()) {
						// challenge not exit or winner not exist
						log.Info("Challenge not exist", "err", err)
					}
					challengeCompletedCh <- &bindings.ChallengeChallengeCompleted{
						Winner: winner,
					}
					continue
				}
				bisectedSub, err = challenge.WatchBisected(&bind.WatchOpts{Context: s.Ctx}, bisectedCh)
				if err != nil {
					log.Crit("Failed to watch challenge event", "err", err)
				}
				challengeCompletedSub, err = challenge.WatchChallengeCompleted(&bind.WatchOpts{Context: s.Ctx}, challengeCompletedCh)
				if err != nil {
					log.Crit("Failed to watch challenge event", "err", err)
				}
				log.Info("Sequencer generate state...")
				states, err = proof.GenerateStates(
					s.ProofBackend,
					s.Ctx,
					ctx.Parent.InboxSize.Uint64(),
					ctx.Assertion.InboxSize.Uint64(),
					nil,
				)
				if err != nil {
					log.Crit("Failed to generate states", "err", err)
				}
				log.Info("Sequencer generate state end...")

				// initialized: get current bisectedCh;
				// not initialized: InitializeChallengeLength;
				if bytes.Equal(stakeStatus.CurrentChallenge.Bytes(), ctx.ChallengeAddr.Bytes()) {
					bisectionHash, _ := challengeSession.BisectionHash()
					if bytes.Equal(bisectionHash[:], common.BigToHash(common.Big0).Bytes()) {
						// when not init
						numSteps := uint64(len(states)) - 1
						log.Info("Print generated states", "states[0]", states[0].Hash().String(), "states[numSteps]", states[numSteps].Hash().String(), "numSteps", numSteps)
						_, err = challengeSession.InitializeChallengeLength(services.MidState(states, 0, numSteps), new(big.Int).SetUint64(numSteps))
						if err != nil {
							log.Crit("Failed to initialize challenge", "err", err)
						}
					} else {
						//bisectedCh chan *bindings.ChallengeBisected
						curr, err := challengeSession.CurrentBisected()
						if err != nil {
							log.Crit("Failed to get current bisected", "err", err)
						}
						bisectedCh <- &bindings.ChallengeBisected{
							StartState:              curr.StartState,
							MidState:                curr.MidState,
							EndState:                curr.EndState,
							BlockNum:                curr.BlockNum,
							BlockTime:               curr.BlockTime,
							ChallengedSegmentStart:  curr.ChallengedSegmentStart,
							ChallengedSegmentLength: curr.ChallengedSegmentLength,
						}
					}
					inChallenge = true
				} else {
					log.Crit("Unrecognized challenge")
				}

			case <-headCh:
				continue // consume channel values
			case <-s.Ctx.Done():
				return
			}
		}
	}
}

func (s *Sequencer) Protocols() []p2p.Protocol {
	// TODO: sequencer APIs
	return []p2p.Protocol{}
}

func (s *Sequencer) APIs() []rpc.API {
	// TODO: sequencer APIs
	return []rpc.API{}
}

func (s *Sequencer) Start() error {
	_ = s.BaseService.Start(true, true)

	s.Wg.Add(2)
	go s.confirmationLoop()
	go s.challengeLoop()
	log.Info("fraud-proof defender started")
	return nil
}

func (s *Sequencer) Stop() error {
	log.Info("fraud-proof defender stopped")
	s.Cancel()
	s.Wg.Wait()
	return nil
}
