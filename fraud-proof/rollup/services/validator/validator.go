package validator

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	rpc2 "github.com/mantlenetworkio/mantle/l2geth/rpc"
	"math/big"
)

func RegisterService(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) {
	validator, err := New(eth, proofBackend, cfg, auth)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	validator.Start()
	log.Info("Validator registered")
}

type challengeCtx struct {
	opponentAssertion *rollupTypes.Assertion
	ourAssertion      *rollupTypes.Assertion
}

type Validator struct {
	*services.BaseService

	batchCh              chan *rollupTypes.TxBatch
	challengeCh          chan *challengeCtx
	challengeResoutionCh chan struct{}
}

func New(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) (*Validator, error) {
	base, err := services.NewBaseService(eth, proofBackend, cfg, auth)
	if err != nil {
		return nil, err
	}
	v := &Validator{
		BaseService:          base,
		batchCh:              make(chan *rollupTypes.TxBatch, 4096),
		challengeCh:          make(chan *challengeCtx),
		challengeResoutionCh: make(chan struct{}),
	}
	return v, nil
}

// This goroutine validates the assertion posted to L1 Rollup, advances
// stake if validated, or challenges if not
func (v *Validator) validationLoop(genesisRoot common.Hash) {
	defer v.Wg.Done()

	// Listen to AssertionCreated event
	assertionEventCh := make(chan *bindings.IRollupAssertionCreated, 4096)
	assertionEventSub, err := v.Rollup.Contract.WatchAssertionCreated(&bind.WatchOpts{Context: v.Ctx}, assertionEventCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer assertionEventSub.Unsubscribe()

	isInChallenge := false

	for {
		if isInChallenge {
			// Wait for the challenge resolution
			select {
			case <-v.challengeResoutionCh:
				log.Info("Validator finished challenge, reset isInChallenge status")
				isInChallenge = false
			case <-v.Ctx.Done():
				return
			}
		} else {
			select {
			case ev := <-assertionEventCh:
				if common.Address(ev.AsserterAddr) == v.Config.StakeAddr {
					// Create by our own for challenge
					continue
				}
				// New assertion created on Rollup
				log.Info("Validator get new assertion, check it with local block....")
				checkAssertion := &rollupTypes.Assertion{
					ID:        ev.AssertionID,
					VmHash:    ev.VmHash,
					InboxSize: ev.InboxSize,
				}

				block, err := v.BaseService.ProofBackend.BlockByNumber(v.Ctx, rpc2.BlockNumber(checkAssertion.InboxSize.Int64()))
				if err != nil {
					log.Error("Validator get block failed, err: ", err)
				}
				if bytes.Compare(checkAssertion.VmHash.Bytes(), block.Hash().Bytes()) != 0 {
					// Validation failed
					log.Info("Validator check assertion vmHash failed, start challenge assertion....")
					ourAssertion := &rollupTypes.Assertion{
						VmHash:    block.Hash(),
						InboxSize: ev.InboxSize,
					}
					v.challengeCh <- &challengeCtx{checkAssertion, ourAssertion}
					isInChallenge = true
				} else {
					// Validation succeeded, confirm assertion and advance stake
					log.Info("Validator advance stake into assertion, ID: ", ev.AssertionID)
					_, err = v.Rollup.AdvanceStake(ev.AssertionID)
					if err != nil {
						log.Crit("UNHANDELED: Can't advance stake, validator state corrupted", "err", err)
					}
				}
			case <-v.Ctx.Done():
				return
			}
		}
	}
}

func (v *Validator) challengeLoop() {
	defer v.Wg.Done()

	abi, err := bindings.IChallengeMetaData.GetAbi()
	if err != nil {
		log.Crit("Failed to get IChallenge ABI", "err", err)
	}

	// Watch AssertionCreated event
	createdCh := make(chan *bindings.IRollupAssertionCreated, 4096)
	createdSub, err := v.Rollup.Contract.WatchAssertionCreated(&bind.WatchOpts{Context: v.Ctx}, createdCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer createdSub.Unsubscribe()

	challengedCh := make(chan *bindings.IRollupAssertionChallenged, 4096)
	challengedSub, err := v.Rollup.Contract.WatchAssertionChallenged(&bind.WatchOpts{Context: v.Ctx}, challengedCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer challengedSub.Unsubscribe()

	// Watch L1 blockchain for challenge timeout
	headCh := make(chan *ethtypes.Header, 4096)
	headSub, err := v.L1.SubscribeNewHead(v.Ctx, headCh)
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
	var ctx *challengeCtx
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
				// If it's our turn
				if common.Address(responder) == v.Config.StakeAddr {
					err := services.RespondBisection(v.BaseService, abi, challengeSession, ev, states, ctx.opponentAssertion.VmHash, false)
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
				v.challengeResoutionCh <- struct{}{}
			case <-v.Ctx.Done():
				bisectedSub.Unsubscribe()
				challengeCompletedSub.Unsubscribe()
				return
			}
		} else {
			select {
			case ctx = <-v.challengeCh:
				log.Info("Validator get challenge context, create challenge assertion)")
				_, err = v.Rollup.CreateAssertion(
					ctx.ourAssertion.VmHash,
					ctx.ourAssertion.InboxSize,
				)
				if err != nil {
					log.Crit("UNHANDELED: Can't create assertion for challenge, validator state corrupted", "err", err)
				}
			case ev := <-createdCh:
				if common.Address(ev.AsserterAddr) == v.Config.StakeAddr {
					if ev.VmHash == ctx.ourAssertion.VmHash {
						_, err := v.Rollup.ChallengeAssertion(
							[2]ethcommon.Address{
								ethcommon.Address(v.Config.SequencerAddr),
								ethcommon.Address(v.Config.StakeAddr),
							},
							[2]*big.Int{
								ctx.opponentAssertion.ID,
								ev.AssertionID,
							},
						)
						if err != nil {
							log.Crit("UNHANDELED: Can't start challenge, validator state corrupted", "err", err)
						}
					}
				}
			case ev := <-challengedCh:
				if ctx == nil {
					continue
				}
				log.Info("Validator saw new challenge", "assertion id", ev.AssertionID, "expected id", ctx.opponentAssertion.ID, "block", ev.Raw.BlockNumber)
				if ev.AssertionID.Cmp(ctx.opponentAssertion.ID) == 0 {
					challenge, err := bindings.NewIChallenge(ev.ChallengeAddr, v.L1)
					if err != nil {
						log.Crit("Failed to access ongoing challenge", "address", ev.ChallengeAddr, "err", err)
					}
					challengeSession = &bindings.IChallengeSession{
						Contract:     challenge,
						CallOpts:     bind.CallOpts{Pending: true, Context: v.Ctx},
						TransactOpts: *v.TransactOpts,
					}
					bisectedCh = make(chan *bindings.IChallengeBisected, 4096)
					bisectedSub, err = challenge.WatchBisected(&bind.WatchOpts{Context: v.Ctx}, bisectedCh)
					if err != nil {
						log.Crit("Failed to watch challenge event", "err", err)
					}
					challengeCompletedCh = make(chan *bindings.IChallengeChallengeCompleted, 4096)
					challengeCompletedSub, err = challenge.WatchChallengeCompleted(&bind.WatchOpts{Context: v.Ctx}, challengeCompletedCh)
					if err != nil {
						log.Crit("Failed to watch challenge event", "err", err)
					}
					parentAssertion, err := ctx.ourAssertion.GetParentAssertion(v.AssertionMap)
					if err != nil {
						log.Crit("Failed to watch challenge event", "err", err)
					}
					states, err = proof.GenerateStates(
						v.ProofBackend,
						v.Ctx,
						parentAssertion.InboxSize.Uint64(),
						ctx.ourAssertion.InboxSize.Uint64(),
						nil,
					)
					if err != nil {
						log.Crit("Failed to generate states", "err", err)
					}
					inChallenge = true
				}
			case <-headCh:
				continue // consume channel values
			case <-v.Ctx.Done():
				return
			}
		}
	}
}

func (v *Validator) Start() error {
	genesis := v.BaseService.Start(true, true)

	v.Wg.Add(2)
	go v.validationLoop(genesis.Root())
	go v.challengeLoop()
	log.Info("Validator started")
	return nil
}

func (v *Validator) Stop() error {
	log.Info("Validator stopped")
	v.Cancel()
	v.Wg.Wait()
	return nil
}

func (v *Validator) APIs() []rpc.API {
	// TODO: validator APIs
	return []rpc.API{}
}
