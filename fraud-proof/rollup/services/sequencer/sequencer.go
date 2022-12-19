package sequencer

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/specularl2/specular/clients/geth/specular/bindings"
	"github.com/specularl2/specular/clients/geth/specular/proof"
	"github.com/specularl2/specular/clients/geth/specular/rollup/services"
	rollupTypes "github.com/specularl2/specular/clients/geth/specular/rollup/types"
)

func RegisterService(stack *node.Node, eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) {
	sequencer, err := New(eth, proofBackend, cfg, auth)
	if err != nil {
		log.Crit("Failed to register the Rollup service", "err", err)
	}
	stack.RegisterLifecycle(sequencer)
	// stack.RegisterAPIs(seq.APIs())
	log.Info("Sequencer registered")
}

type challengeCtx struct {
	challengeAddr common.Address
	assertion     *rollupTypes.Assertion
}

// Current Sequencer assumes no Berlin+London fork on L2
type Sequencer struct {
	*services.BaseService

	batchCh              chan *rollupTypes.TxBatch
	pendingAssertionCh   chan *rollupTypes.Assertion
	confirmedIDCh        chan *big.Int
	challengeCh          chan *challengeCtx
	challengeResoutionCh chan struct{}
}

func New(eth services.Backend, proofBackend proof.Backend, cfg *services.Config, auth *bind.TransactOpts) (*Sequencer, error) {
	base, err := services.NewBaseService(eth, proofBackend, cfg, auth)
	if err != nil {
		return nil, err
	}
	s := &Sequencer{
		BaseService:          base,
		batchCh:              make(chan *rollupTypes.TxBatch, 4096),
		pendingAssertionCh:   make(chan *rollupTypes.Assertion, 4096),
		confirmedIDCh:        make(chan *big.Int, 4096),
		challengeCh:          make(chan *challengeCtx),
		challengeResoutionCh: make(chan struct{}),
	}
	return s, nil
}

// This goroutine fetches txs from txpool and batches them
// TODO: fetching txs through sub is unstable
// TODO: batch txs by timer and directly querying txpool as worker does
func (s *Sequencer) batchingLoop() {
	defer s.Wg.Done()
	defer close(s.batchCh)

	// Watch transactions in TxPool
	txsCh := make(chan core.NewTxsEvent, 4096)
	txsSub := s.Eth.TxPool().SubscribeNewTxsEvent(txsCh)
	defer txsSub.Unsubscribe()

	batcher, err := NewBatcher(s.Config.Coinbase, s.Eth)
	if err != nil {
		log.Crit("Failed to start batcher", "err", err)
	}

	for {
		select {
		case ev := <-txsCh:
			// New transactions from TxPool
			err = batcher.CommitTransactions(ev.Txs)
			if err != nil {
				log.Crit("Failed to commit transactions", "err", err)
			}
			blocks, err := batcher.Batch()
			if err != nil {
				log.Crit("Failed to batch blocks", "err", err)
			}
			batch := rollupTypes.NewTxBatch(blocks, 0) // TODO: add max batch size
			s.batchCh <- batch
		case <-s.Ctx.Done():
			return
		}
	}
}

// This goroutine sequences batches to L1 SequencerInbox and creates DAs
// TODO: create assertions by timer
func (s *Sequencer) sequencingLoop(genesisRoot common.Hash) {
	defer s.Wg.Done()

	// Watch AssertionCreated event
	createdCh := make(chan *bindings.IRollupAssertionCreated, 4096)
	createdSub, err := s.Rollup.Contract.WatchAssertionCreated(&bind.WatchOpts{Context: s.Ctx}, createdCh)
	if err != nil {
		log.Crit("Failed to watch rollup event", "err", err)
	}
	defer createdSub.Unsubscribe()

	// Current confirmed assertion, initalize it to genesis
	// TODO: sync from L1 Rollup
	confirmedAssertion := &rollupTypes.Assertion{
		ID:                    new(big.Int),
		VmHash:                genesisRoot,
		CumulativeGasUsed:     new(big.Int),
		InboxSize:             new(big.Int),
		Deadline:              new(big.Int),
		PrevCumulativeGasUsed: new(big.Int),
	}
	// Assertion created and pending for confirmation
	var pendingAssertion *rollupTypes.Assertion
	// Assertion to be created on L1 Rollup
	queuedAssertion := confirmedAssertion.Copy()
	queuedAssertion.StartBlock = 1

	// Create assertion on L1 Rollup
	commitAssertion := func() {
		pendingAssertion = queuedAssertion.Copy()
		queuedAssertion.StartBlock = queuedAssertion.EndBlock + 1
		queuedAssertion.PrevCumulativeGasUsed = new(big.Int).Set(queuedAssertion.CumulativeGasUsed)
		_, err = s.Rollup.CreateAssertion(
			pendingAssertion.VmHash,
			pendingAssertion.InboxSize,
			pendingAssertion.CumulativeGasUsed,
			confirmedAssertion.VmHash,
			confirmedAssertion.CumulativeGasUsed,
		)
		if err != nil {
			log.Error("Can not create DA", "error", err)
		}
	}

	for {
		select {
		case batch := <-s.batchCh:
			// New batch from Batcher
			log.Info("Get New Batch", "length", len(batch.Txs))
			// Sequence batch to SequencerInbox
			contexts, txLengths, txs, err := batch.SerializeToArgs()
			if err != nil {
				log.Error("Can not serialize batch", "error", err)
				continue
			}
			_, err = s.Inbox.AppendTxBatch(contexts, txLengths, txs)
			if err != nil {
				log.Error("Can not sequence batch", "error", err)
				continue
			}
			// Update queued assertion to latest batch
			queuedAssertion.VmHash = batch.LastBlockRoot()
			queuedAssertion.CumulativeGasUsed.Add(queuedAssertion.CumulativeGasUsed, batch.GasUsed)
			queuedAssertion.InboxSize.Add(queuedAssertion.InboxSize, batch.InboxSize())
			queuedAssertion.EndBlock = batch.LastBlockNumber()
			// If no assertion is pending, commit it
			if pendingAssertion == nil {
				log.Info("commitAssertion.....")
				commitAssertion()
			}
		case ev := <-createdCh:
			// New assertion created on L1 Rollup
			log.Info(fmt.Sprintf("Get New Assertion, AssertionID: %s, AsserterAddress: %s",
				ev.AssertionID.String(), ev.AsserterAddr.String()))
			if common.Address(ev.AsserterAddr) == s.Config.Coinbase {
				if ev.VmHash == pendingAssertion.VmHash {
					log.Info("confirmAssertion.....")
					// If assertion is created by us, get ID and deadline
					pendingAssertion.ID = ev.AssertionID
					pendingAssertion.Deadline, err = s.AssertionMap.GetDeadline(ev.AssertionID)
					if err != nil {
						log.Error("Can not get DA deadline", "error", err)
						continue
					}
					// Send to confirmation goroutine to confirm it
					s.pendingAssertionCh <- pendingAssertion
				}
			}
		case id := <-s.confirmedIDCh:
			// New assertion confirmed
			if pendingAssertion.ID.Cmp(id) == 0 {
				confirmedAssertion = pendingAssertion
				if pendingAssertion.VmHash == queuedAssertion.VmHash {
					// We are done here, waiting for new batches
					pendingAssertion = nil
				} else {
					// Commit queued assertion
					commitAssertion()
				}
			} else {
				// TODO: decentralized sequencer
				// TODO: rewind blockchain, sync from L1, reset states
				log.Error("Confirmed ID is not current pending one", "get", id.String(), "expected", pendingAssertion.ID.String())
			}
		case <-s.Ctx.Done():
			return
		}
	}
}

// This goroutine tries to confirm created assertions
func (s *Sequencer) confirmationLoop() {
	defer s.Wg.Done()

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
	// TODO: watch multiple pending assertions
	var pendingAssertion *rollupTypes.Assertion
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
			case header := <-headCh:
				// New block mined on L1
				log.Info("sequencer sync new layer1 block...")
				if !pendingConfirmationSent && !pendingConfirmed {
					if header.Number.Uint64() >= pendingAssertion.Deadline.Uint64() {
						// Confirmation period has past, confirm it
						log.Info("call ConfirmFirstUnresolvedAssertion...")
						_, err := s.Rollup.ConfirmFirstUnresolvedAssertion()
						if err != nil {
							// log.Error("Failed to confirm DA", "error", err)
							log.Crit("Failed to confirm DA", "err", err)
							// TODO: wait some time before retry
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
				// New challenge raised
				if ev.AssertionID.Cmp(pendingAssertion.ID) == 0 {
					s.challengeCh <- &challengeCtx{
						ev.ChallengeAddr,
						pendingAssertion,
					}
					isInChallenge = true
				}
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
				if responder == common.Address(s.Config.Coinbase) {
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
				challenge, err := bindings.NewIChallenge(ctx.challengeAddr, s.L1)
				if err != nil {
					log.Crit("Failed to access ongoing challenge", "address", ctx.challengeAddr, "err", err)
				}
				challengeSession = &bindings.IChallengeSession{
					Contract:     challenge,
					CallOpts:     bind.CallOpts{Pending: true, Context: s.Ctx},
					TransactOpts: *s.TransactOpts,
				}
				bisectedCh = make(chan *bindings.IChallengeBisected, 4096)
				bisectedSub, err = challenge.WatchBisected(&bind.WatchOpts{Context: s.Ctx}, bisectedCh)
				if err != nil {
					log.Crit("Failed to watch challenge event", "err", err)
				}
				challengeCompletedCh = make(chan *bindings.IChallengeChallengeCompleted, 4096)
				challengeCompletedSub, err = challenge.WatchChallengeCompleted(&bind.WatchOpts{Context: s.Ctx}, challengeCompletedCh)
				if err != nil {
					log.Crit("Failed to watch challenge event", "err", err)
				}
				log.Info("to generate state from", "start", ctx.assertion.StartBlock, "to", ctx.assertion.EndBlock)
				log.Info("backend", "start", ctx.assertion.StartBlock, "to", ctx.assertion.EndBlock)
				states, err = proof.GenerateStates(
					s.ProofBackend,
					s.Ctx,
					ctx.assertion.PrevCumulativeGasUsed,
					ctx.assertion.StartBlock,
					ctx.assertion.EndBlock+1,
					nil,
				)
				if err != nil {
					log.Crit("Failed to generate states", "err", err)
				}
				_, err = challengeSession.InitializeChallengeLength(new(big.Int).SetUint64(uint64(len(states)) - 1))
				if err != nil {
					log.Crit("Failed to initialize challenge", "err", err)
				}
				inChallenge = true
			case <-headCh:
				continue // consume channel values
			case <-s.Ctx.Done():
				return
			}
		}
	}
}

func (s *Sequencer) Start() error {
	genesis := s.BaseService.Start()

	s.Wg.Add(4)
	go s.batchingLoop()
	go s.sequencingLoop(genesis.Root())
	go s.confirmationLoop()
	go s.challengeLoop()
	log.Info("Sequencer started")
	return nil
}

func (s *Sequencer) Stop() error {
	log.Info("Sequencer stopped")
	s.Cancel()
	s.Wg.Wait()
	return nil
}

func (s *Sequencer) APIs() []rpc.API {
	// TODO: sequencer APIs
	return []rpc.API{}
}
