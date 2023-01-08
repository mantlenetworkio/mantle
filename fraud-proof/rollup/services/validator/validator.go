package validator

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	"github.com/mantlenetworkio/mantle/fraud-proof/proof"
	"github.com/mantlenetworkio/mantle/fraud-proof/rollup/services"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
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
	opponentAssertion  *rollupTypes.Assertion
	ourAssertion       *rollupTypes.Assertion
	confirmedAssertion *rollupTypes.Assertion
}

// TODO: abstract the common field/init to a base service
type Validator struct {
	*services.BaseService

	batchCh              chan *rollupTypes.TxBatch
	challengeCh          chan *challengeCtx
	challengeResoutionCh chan struct{}
}

// TODO: this shares a lot of code with sequencer
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

// commitBlocks executes and commits sequenced blocks to local blockchain
// It returns the afterwards state hash and total gas used for these blocks
// TODO: this function shares a lot of codes with Batcher
// TODO: use StateProcessor::Process() instead
func (v *Validator) commitBlocks(blocks []*rollupTypes.SequenceBlock) (common.Hash, *big.Int, error) {
	var targetVmHash common.Hash
	targetGasUsed := new(big.Int)
	chainConfig := v.Chain.Config()
	parent := v.Chain.CurrentBlock()
	if parent == nil {
		return common.Hash{}, nil, fmt.Errorf("missing parent")
	}
	num := parent.Number()
	if num.Uint64() != blocks[0].BlockNumber-1 {
		return common.Hash{}, nil, fmt.Errorf("validator unsynced")
	}
	state, err := v.Chain.StateAt(parent.Root())
	// TODO: in newest version of geth, below codes are removed
	if err != nil {
		// Note since the sealing block can be created upon the arbitrary parent
		// block, but the state of parent block may already be pruned, so the necessary
		// state recovery is needed here in the future.
		//
		// The maximum acceptable reorg depth can be limited by the finalised block
		// somehow. TODO(rjl493456442) fix the hard-coded number here later.
		state, err = v.Eth.StateAtBlock(parent, 1024, nil, false, false)
		log.Warn("Recovered mining state", "root", parent.Root(), "err", err)
	}
	if err != nil {
		return common.Hash{}, nil, err
	}

	for _, sblock := range blocks {
		header := &types.Header{
			ParentHash: parent.Hash(),
			Number:     new(big.Int).SetUint64(sblock.BlockNumber),
			GasLimit:   core.CalcGasLimit(parent, parent.GasLimit(), ethconfig.Defaults.Miner.GasCeil), // TODO: this may cause problem
			Time:       sblock.Timestamp,
			//Coinbase:   v.Config.SequencerAddr, //TODO-FIXME
			Difficulty: common.Big1, // Fake difficulty. Avoid use 0 here because it means the merge happened
		}
		gasPool := new(core.GasPool).AddGas(header.GasLimit)
		var receipts []*types.Receipt
		for idx, tx := range sblock.Txs {
			state.Prepare(tx.Hash(), tx.Hash(), idx) // TODO-FIXME test out if tx hash ==== block hash
			receipt, err := core.ApplyTransaction(chainConfig, v.Chain, &v.Config.SequencerAddr, gasPool, state, header, tx, &header.GasUsed, *v.Chain.GetVMConfig())
			if err != nil {
				return common.Hash{}, nil, err
			}
			receipts = append(receipts, receipt)
		}
		// Finalize header
		header.Root = state.IntermediateRoot(v.Chain.Config().IsEIP158(header.Number))
		header.UncleHash = types.CalcUncleHash(nil)
		// Assemble block
		block := types.NewBlock(header, sblock.Txs, nil, receipts)
		hash := block.Hash()
		// Finalize receipts and logs
		var logs []*types.Log
		for i, receipt := range receipts {
			// Add block location fields
			receipt.BlockHash = hash
			receipt.BlockNumber = block.Number()
			receipt.TransactionIndex = uint(i)

			// Update the block hash in all logs since it is now available and not when the
			// receipt/log of individual transactions were created.
			for _, log := range receipt.Logs {
				log.BlockHash = hash
			}
			logs = append(logs, receipt.Logs...)
		}
		_, err := v.Chain.WriteBlockAndSetHead(block, receipts, logs, state, false)
		if err != nil {
			return common.Hash{}, nil, err
		}
		targetVmHash = header.Root
		targetGasUsed.Add(targetGasUsed, new(big.Int).SetUint64(header.GasUsed))
	}
	return targetVmHash, targetGasUsed, nil
}

// This goroutine synchronizes sequenced transactions from L1 SequencerInbox
func (v *Validator) collectingLoop() {
	defer v.Wg.Done()

	//abi, err := bindings.ISequencerInboxMetaData.GetAbi()
	//if err != nil {
	//	log.Crit("Failed to get ISequencerInbox ABI", "err", err)
	//}

	//// Listen to TxBatchAppendEvent
	//batchEventCh := make(chan *bindings.ISequencerInboxTxBatchAppended, 4096)
	//batchEventSub, err := v.Inbox.Contract.WatchTxBatchAppended(&bind.WatchOpts{Context: v.Ctx}, batchEventCh)
	//if err != nil {
	//	log.Crit("Failed to watch rollup event", "err", err)
	//}
	//defer batchEventSub.Unsubscribe()

	for {
		// if challenge, pause
		select {
		//case ev := <-batchEventCh:
		//	log.Warn("new batch rise", ev)
		//	// New appendTxBatch call
		//	log.Info(fmt.Sprintf("Get New Batch, batchNum: %s, startTxNum: %s, endTxNum: %s",
		//		ev.BatchNumber, ev.StartTxNumber, ev.EndTxNumber))
		//	tx, _, err := v.L1.TransactionByHash(v.Ctx, ev.Raw.TxHash)
		//	if err != nil {
		//		log.Error("Failed to get tx batch data", "error", err)
		//		continue
		//	}
		//	// Decode input from abi
		//	decoded, err := abi.Methods["appendTxBatch"].Inputs.Unpack(tx.Data()[4:])
		//	if err != nil {
		//		log.Error("Failed to decode tx batch data", "error", err)
		//		continue
		//	}
		//	// Construct batch
		//	batch, err := rollupTypes.TxBatchFromDecoded(decoded)
		//	if err != nil {
		//		log.Error("Failed to decode tx batch data", "error", err)
		//		continue
		//	}
		//	v.batchCh <- batch
		case <-v.Ctx.Done():
			return
		}
	}
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

	// Sequenced blocks collected from SequencerInbox
	var pendingBlocks []*rollupTypes.SequenceBlock
	// Current agreed assertion, initalize to genesis assertion
	// TODO: sync from L1 when restart
	confirmedAssertion := &rollupTypes.Assertion{
		ID:        new(big.Int),
		VmHash:    genesisRoot,
		GasUsed:   new(big.Int),
		InboxSize: new(big.Int),
		Deadline:  new(big.Int),
	}

	isInChallenge := false

	for {
		if isInChallenge {
			// Wait for the challenge resolution
			select {
			case <-v.challengeResoutionCh:
				log.Info("challenge finished")
				isInChallenge = false
			case <-v.Ctx.Done():
				return
			}
		} else {
			select {
			case batch := <-v.batchCh:
				// New batch collected from SequencerInbox, split it into blocks
				txNum := 0
				for _, ctx := range batch.Contexts {
					block := &rollupTypes.SequenceBlock{
						SequenceContext: ctx,
						Txs:             batch.Txs[txNum : txNum+int(ctx.NumTxs)],
					}
					pendingBlocks = append(pendingBlocks, block)
					txNum += int(ctx.NumTxs)
				}
				log.Info("preparing pendingBlocks, SequenceBlocks blockNUm: ", len(pendingBlocks))
			case ev := <-assertionEventCh:
				if common.Address(ev.AsserterAddr) == v.Config.Coinbase {
					// Create by our own for challenge
					continue
				}
				// New assertion created on Rollup
				log.Info("Get New Assertion....")
				assertion := &rollupTypes.Assertion{
					ID:        ev.AssertionID,
					VmHash:    ev.VmHash,
					InboxSize: ev.InboxSize,
					GasUsed:   ev.L2GasUsed,
					//StartBlock:            confirmedAssertion.EndBlock + 1,
					//PrevCumulativeGasUsed: new(big.Int).Set(confirmedAssertion.CumulativeGasUsed),
				}
				// Pop correct amount of pending blocks asserted
				inboxSizeDiff := assertion.InboxSize.Uint64() - confirmedAssertion.InboxSize.Uint64()
				var blocksToCommit []*rollupTypes.SequenceBlock
				for _, block := range pendingBlocks {
					if block.NumTxs > inboxSizeDiff {
						log.Crit("UNHANDELED: Assertion created in the middle of block, validator state corrupted!")
					}
					inboxSizeDiff -= block.NumTxs
					blocksToCommit = append(blocksToCommit, block)
					if inboxSizeDiff == 0 {
						break
					}
				}
				if inboxSizeDiff != 0 {
					log.Crit("UNHANDELED: SequencerInbox overflow, validator state corrupted!")
				}
				//assertion.EndBlock = assertion.StartBlock + uint64(len(blocksToCommit)) - 1
				pendingBlocks = pendingBlocks[len(blocksToCommit):]
				// Commit asserted blocks
				log.Info("Commit Blocks....")
				targetVmHash, targetGasUsed, err := v.commitBlocks(blocksToCommit)
				if err != nil {
					// strange error, TODO: rewind
					log.Crit("UNHANDELED: Can't execute sequence blocks, validator state corrupted", "err", err)
				}
				// Check result vm hash and gas
				targetGasUsed.Add(targetGasUsed, confirmedAssertion.GasUsed)
				if targetVmHash != assertion.VmHash || targetGasUsed.Cmp(assertion.GasUsed) != 0 {
					// Validation failed
					log.Info("Challenge Assertion....")
					ourAssertion := &rollupTypes.Assertion{
						VmHash:    targetVmHash,
						InboxSize: ev.InboxSize,
						GasUsed:   targetGasUsed,
					}
					v.challengeCh <- &challengeCtx{assertion, ourAssertion, confirmedAssertion}
					isInChallenge = true
				} else {
					// Validation succeeded, confirm assertion and advance stake
					log.Info("Advance State....")
					_, err = v.Rollup.AdvanceStake(ev.AssertionID)
					if err != nil {
						log.Crit("UNHANDELED: Can't advance stake, validator state corrupted", "err", err)
					}
					confirmedAssertion = assertion
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
				if common.Address(responder) == v.Config.Coinbase {
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
				_, err = v.Rollup.CreateAssertion(
					ctx.ourAssertion.VmHash,
					ctx.ourAssertion.InboxSize,
					ctx.ourAssertion.GasUsed.Add(ctx.ourAssertion.GasUsed, big.NewInt(1)),
				)
				if err != nil {
					log.Crit("UNHANDELED: Can't create assertion for challenge, validator state corrupted", "err", err)
				}
			case ev := <-createdCh:
				if common.Address(ev.AsserterAddr) == v.Config.Coinbase {
					if ev.VmHash == ctx.ourAssertion.VmHash {
						_, err := v.Rollup.ChallengeAssertion(
							[2]ethcommon.Address{
								ethcommon.Address(v.Config.SequencerAddr),
								ethcommon.Address(v.Config.Coinbase),
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
				log.Info("validator saw challenge", "assertion id", ev.AssertionID, "expected id", ctx.opponentAssertion.ID, "block", ev.Raw.BlockNumber)
				if ev.AssertionID.Cmp(ctx.opponentAssertion.ID) == 0 {
					// start := ev.Raw.BlockNumber - 2
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
					states, err = proof.GenerateStates(
						v.ProofBackend,
						v.Ctx,
						big.NewInt(0), //ctx.opponentAssertion.PrevCumulativeGasUsed,
						0,             //ctx.opponentAssertion.StartBlock,
						0,             //ctx.opponentAssertion.EndBlock+1,
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
	genesis := v.BaseService.Start()

	v.Wg.Add(3)
	go v.collectingLoop()
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
