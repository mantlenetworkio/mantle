package sequencer

import (
	"fmt"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/l2geth/p2p"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	"github.com/mantlenetworkio/mantle/l2geth/common"
)

func RegisterService(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) {
	sequencer, err := New(eth, proofBackend, cfg, auth)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	sequencer.Start()
	log.Info("Sequencer registered")
}

type challengeCtx struct {
	challengeAddr common.Address
	parent        *rollupTypes.Assertion
	assertion     *rollupTypes.Assertion
}

// Current Sequencer assumes no Berlin+London fork on L2
type Sequencer struct {
	*services.BaseService

	pendingAssertionCh   chan *rollupTypes.Assertion
	confirmedIDCh        chan *big.Int
	challengeCh          chan *challengeCtx
	challengeResoutionCh chan struct{}

	confirmations uint64
}

func New(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) (*Sequencer, error) {
	base, err := services.NewBaseService(eth, proofBackend, cfg, auth)
	if err != nil {
		return nil, err
	}
	s := &Sequencer{
		BaseService:          base,
		pendingAssertionCh:   make(chan *rollupTypes.Assertion, 4096),
		confirmedIDCh:        make(chan *big.Int, 4096),
		challengeCh:          make(chan *challengeCtx),
		challengeResoutionCh: make(chan struct{}),
		confirmations:        cfg.L1Confirmations,
	}
	return s, nil
}

// This goroutine tries to confirm created assertions
func (s *Sequencer) confirmationLoop() {
	defer s.Wg.Done()

	// Watch AssertionCreated event
	createdCh := make(chan *bindings.IRollupAssertionCreated, 4096)
	createdSub, err := s.Rollup.Contract.WatchAssertionCreated(&bind.WatchOpts{Context: s.Ctx}, createdCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer createdSub.Unsubscribe()

	// Watch AssertionConfirmed event
	confirmedCh := make(chan *bindings.IRollupAssertionConfirmed, 4096)
	confirmedSub, err := s.Rollup.Contract.WatchAssertionConfirmed(&bind.WatchOpts{Context: s.Ctx}, confirmedCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer confirmedSub.Unsubscribe()

	// Watch L1 blockchain for confirmation period
	headCh := make(chan *types.Header, 4096)
	headSub, err := s.L1.SubscribeNewHead(s.Ctx, headCh)
	if err != nil {
		log.Crit("Failed to watch l1 chain head", "err", err)
	}
	defer headSub.Unsubscribe()

	challengedCh := make(chan *bindings.IRollupAssertionChallenged, 4096)
	challengedSub, err := s.Rollup.Contract.WatchAssertionChallenged(&bind.WatchOpts{Context: s.Ctx}, challengedCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer challengedSub.Unsubscribe()
	isInChallenge := false

	// Current pending assertion from sequencing goroutine
	pendingAssertion := new(rollupTypes.Assertion)
	pendingConfirmationSent := true
	pendingConfirmed := true

	for {
		if isInChallenge {
			// Waif for the challenge resolved
			select {
			case <-s.challengeResoutionCh:
				log.Info("challenge finished")
				isInChallenge = false
			case <-s.Ctx.Done():
				return
			}
		} else {
			select {
			case ev := <-createdCh:
				// New assertion created on L1 Rollup
				log.Info(fmt.Sprintf("Get New Assertion, AssertionID: %s, AsserterAddress: %s",
					ev.AssertionID.String(), ev.AsserterAddr.String()))
				if common.Address(ev.AsserterAddr) == s.Config.StakeAddr {
					log.Info("confirmAssertion.....")
					pendingAssertion.ID = ev.AssertionID
					pendingAssertion.VmHash = ev.VmHash
					pendingAssertion.InboxSize = ev.InboxSize
					pendingAssertion.GasUsed = ev.L2GasUsed
					pendingAssertion.Parent = new(big.Int).Sub(ev.AssertionID, big.NewInt(1))
					pendingAssertion.Deadline, err = s.AssertionMap.GetDeadline(ev.AssertionID)
					if err != nil {
						log.Error("Can not get Assertion deadline", "error", err)
						continue
					}
					pendingAssertion.ProposalTime, err = s.AssertionMap.GetProposalTime(ev.AssertionID)
					if err != nil {
						log.Error("Can not get Assertion proposal time", "error", err)
						continue
					}
					// New assertion created by sequencing goroutine
					if !pendingConfirmed {
						log.Error("Got another DA request before current is confirmed")
						continue
					}
					log.Info("confirmAssertion setup states.....")
					pendingConfirmationSent = false
					pendingConfirmed = false
				}
			case header := <-headCh:
				// Get confirm block header
				if s.confirmations != 0 {
					num := new(big.Int)
					num.SetUint64(s.confirmations)
					num.Sub(header.Number, num)
					header, err = s.L1.HeaderByNumber(s.Ctx, num)
					if err != nil {
						log.Crit("Failed to get confirmed header", "err", err)
						continue
					}
				}
				// New block mined on L1
				log.Info("sequencer sync new layer1 block...")
				if !pendingConfirmationSent && !pendingConfirmed {
					if header.Time >= pendingAssertion.Deadline.Uint64() {
						// Confirmation period has past, confirm it
						log.Info("call ConfirmFirstUnresolvedAssertion...")
						_, err := s.Rollup.ConfirmFirstUnresolvedAssertion()
						if err != nil {
							log.Crit("Failed to confirm DA", "err", err)
							continue
						}
						pendingConfirmationSent = true
					}
				}
			case ev := <-confirmedCh:
				// New confirmed assertion
				if ev.AssertionID.Cmp(pendingAssertion.ID) == 0 {
					// Notify sequencing goroutine
					s.confirmedIDCh <- pendingAssertion.ID
					pendingConfirmed = true
				}
			case newPendingAssertion := <-s.pendingAssertionCh:
				// New assertion created by sequencing goroutine
				if !pendingConfirmed {
					// TODO: support multiple pending assertion
					log.Error("Got another DA request before current is confirmed")
					continue
				}
				log.Info("confirmAssertion setup states.....")
				pendingAssertion = newPendingAssertion.Copy()
				pendingConfirmationSent = false
				pendingConfirmed = false
			case ev := <-challengedCh:
				log.Warn("new challenge rise", ev)
				// New challenge raised
				//if ev.AssertionID.Cmp(pendingAssertion.ID) == 0 {
				//	s.challengeCh <- &challengeCtx{
				//		common.Address(ev.ChallengeAddr),
				//		pendingAssertion,
				//	}
				//	isInChallenge = true
				//}
			case <-s.Ctx.Done():
				return
			}
		}
	}
}

func (s *Sequencer) challengeLoop() {
	defer s.Wg.Done()

	abi, err := bindings.IChallengeMetaData.GetAbi()
	if err != nil {
		log.Crit("Failed to get IChallenge ABI", "err", err)
	}

	// Watch L1 blockchain for challenge timeout
	headCh := make(chan *types.Header, 4096)
	headSub, err := s.L1.SubscribeNewHead(s.Ctx, headCh)
	if err != nil {
		log.Crit("Failed to watch l1 chain head", "err", err)
	}
	defer headSub.Unsubscribe()

	var challengeSession *bindings.IChallengeSession
	var states []*proof.ExecutionState

	var bisectedCh chan *bindings.IChallengeBisected
	var bisectedSub event.Subscription
	var challengeCompletedCh chan *bindings.IChallengeChallengeCompleted
	var challengeCompletedSub event.Subscription

	inChallenge := false
	var opponentTimeoutBlock uint64

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
				if common.Address(responder) == s.Config.StakeAddr {
					// If it's our turn
					err := services.RespondBisection(s.BaseService, abi, challengeSession, ev, states, common.Hash{}, false)
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
					log.Info("[challenge] Opponent time left", "time", opponentTimeLeft)
					opponentTimeoutBlock = ev.Raw.BlockNumber + opponentTimeLeft.Uint64()
				}
			case header := <-headCh:
				if opponentTimeoutBlock == 0 {
					continue
				}
				// TODO: can we use >= here?
				if header.Number.Uint64() > opponentTimeoutBlock {
					_, err := challengeSession.Timeout()
					if err != nil {
						log.Error("Can not timeout opponent", "error", err)
						continue
						// TODO: wait some time before retry
						// TODO: fix race condition
					}
				}
			case ev := <-challengeCompletedCh:
				// TODO: handle if we are not winner --> state corrupted
				log.Info("[challenge] Challenge completed", "winner", ev.Winner)
				bisectedSub.Unsubscribe()
				challengeCompletedSub.Unsubscribe()
				states = []*proof.ExecutionState{}
				inChallenge = false
				s.challengeResoutionCh <- struct{}{}
			case <-s.Ctx.Done():
				bisectedSub.Unsubscribe()
				challengeCompletedSub.Unsubscribe()
				return
			}
		} else {
			select {
			case ctx := <-s.challengeCh:
				log.Warn("new challenge rise, handle it", ctx)
				//challenge, err := bindings.NewIChallenge(ethc.Address(ctx.challengeAddr), s.L1)
				//if err != nil {
				//	log.Crit("Failed to access ongoing challenge", "address", ctx.challengeAddr, "err", err)
				//}
				//challengeSession = &bindings.IChallengeSession{
				//	Contract:     challenge,
				//	CallOpts:     bind.CallOpts{Pending: true, Context: s.Ctx},
				//	TransactOpts: *s.TransactOpts,
				//}
				//bisectedCh = make(chan *bindings.IChallengeBisected, 4096)
				//bisectedSub, err = challenge.WatchBisected(&bind.WatchOpts{Context: s.Ctx}, bisectedCh)
				//if err != nil {
				//	log.Crit("Failed to watch challenge event", "err", err)
				//}
				//challengeCompletedCh = make(chan *bindings.IChallengeChallengeCompleted, 4096)
				//challengeCompletedSub, err = challenge.WatchChallengeCompleted(&bind.WatchOpts{Context: s.Ctx}, challengeCompletedCh)
				//if err != nil {
				//	log.Crit("Failed to watch challenge event", "err", err)
				//}
				//log.Info("to generate state from", "start", ctx.assertion.StartBlock, "to", ctx.assertion.InboxSize)
				//log.Info("backend", "start", ctx.assertion.StartBlock, "to", ctx.assertion.EndBlock)
				//states, err = proof.GenerateStates(
				//	s.ProofBackend,
				//	s.Ctx,
				//	ctx.assertion.PrevCumulativeGasUsed,
				//	ctx.assertion.StartBlock,
				//	ctx.assertion.EndBlock+1,
				//	nil,
				//)
				//if err != nil {
				//	log.Crit("Failed to generate states", "err", err)
				//}
				//_, err = challengeSession.InitializeChallengeLength(new(big.Int).SetUint64(uint64(len(states)) - 1))
				//if err != nil {
				//	log.Crit("Failed to initialize challenge", "err", err)
				//}
				//inChallenge = true
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
	//go s.confirmationLoop()
	//go s.challengeLoop()
	log.Info("fraud-proof defender started")
	return nil
}

func (s *Sequencer) Stop() error {
	log.Info("fraud-proof defender stopped")
	s.Cancel()
	//s.Wg.Wait()
	return nil
}

func (s *Sequencer) CreateAssertion(obj interface{}) error {
	assertion, _ := obj.(rollupTypes.Assertion)
	_, err := s.Rollup.CreateAssertion(
		assertion.VmHash,
		assertion.InboxSize,
		assertion.GasUsed,
	)
	return err
}

func (s *Sequencer) CreateAssertionWithStateBatch(batches [][32]byte, shouldStartAtElement *big.Int, signature []byte, obj interface{}) error {
	assertion, _ := obj.(rollupTypes.Assertion)
	_, err := s.Rollup.CreateAssertionWithStateBatch(
		assertion.VmHash,
		assertion.InboxSize,
		assertion.GasUsed,
		batches,
		shouldStartAtElement,
		signature,
	)
	return err
}

func (s *Sequencer) GetLatestAssertion(staker common.Address) (interface{}, error) {
	var latestAssertion rollupTypes.Assertion
	var assertionID *big.Int
	var err error
	if assertionID, err = s.AssertionMap.GetLatestAssertionID(ethc.Address(staker)); err != nil {
		return nil, err
	}
	if ret, err := s.AssertionMap.Assertions(assertionID); err != nil {
		return nil, err
	} else {
		latestAssertion.ID = assertionID
		latestAssertion.VmHash = ret.StateHash
		latestAssertion.InboxSize = ret.InboxSize
		latestAssertion.GasUsed = ret.GasUsed
		latestAssertion.Parent = ret.Parent
		latestAssertion.Deadline = ret.Deadline
		latestAssertion.ProposalTime = ret.ProposalTime
	}
	return latestAssertion, nil
}

func (s *Sequencer) GenerateState() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sequencer) InChallenge() bool {
	//TODO implement me
	return false
}

func (s *Sequencer) RespondChallenge() error {
	//TODO implement me
	panic("implement me")
}
