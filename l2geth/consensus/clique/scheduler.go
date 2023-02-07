package clique

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/rawdb"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/event"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent
	chainHeadChanSize = 10

	// default config of scheduler consensus
	// defaultBatchSize is the size of batch tx num
	defaultBatchSize = int64(100)
	// defaultExpireTime is the default config of a batch timeout
	defaultExpireTime = int64(60)
	// defaultBatchEpoch is the epoch time of the scheduler synchronizing the sequencer collection from l1, default 1 day
	defaultBatchEpoch = time.Duration(86400)
)

var (
	scale = big.NewInt(1e18)
)

// Config sets the parameters required for the scheduler service
type Config struct {
	BatchSize  int64
	BatchTime  int64
	BatchEpoch int64
}

// Scheduler service has the ability to update the sequencer set
// and select a producer from the sequencer set according to the priority algorithm.
// The scheduler then generates blocks and broadcasts them based
// on the transaction set sent by the producer
type Scheduler struct {
	sequencerSetMtx sync.Mutex // The lock used to protect the sequencerSet field

	eventMux *event.TypeMux
	exitCh   chan struct{}
	running  int32

	config *Config

	db ethdb.Database

	schedulerAddr   common.Address
	sequencerSet    *SequencerSet
	consensusEngine *Clique
	blockchain      *core.BlockChain
	txpool          *core.TxPool
	verifiedTxCount func() (uint64, error)

	// consensus channel
	batchDone       chan uint64
	currentStartMsg types.BatchPeriodStartMsg
	currentHeight   uint64

	expectMinTxsCount uint64
	sequencerAssessor *healthAssessor

	chainHeadSub event.Subscription
	chainHeadCh  chan core.ChainHeadEvent
	addPeerSub   *event.TypeMuxSubscription
	batchEndSub  *event.TypeMuxSubscription

	ticker *time.Ticker

	wallet      accounts.Wallet
	signAccount accounts.Account

	syncer *synchronizer.Synchronizer
}

// NewScheduler will create a scheduler server
func NewScheduler(db ethdb.Database, config *Config, schedulerAddress common.Address, clique *Clique, blockchain *core.BlockChain, txpool *core.TxPool, eventMux *event.TypeMux) (*Scheduler, error) {
	log.Info("Create Sequencer Server")

	syncer := synchronizer.NewSynchronizer()
	schedulerAddr, err := syncer.GetSchedulerAddr()
	if err != nil {
		return nil, err
	}
	if schedulerAddr.String() != schedulerAddress.String() {
		return nil, fmt.Errorf("scheduler address mismatch, schedulerAddr from L1 %s,schedulerAddr from config %s", schedulerAddr.String(), schedulerAddress.String())
	}
	seqSet, err := syncer.GetSequencerSet()
	if err != nil {
		return nil, fmt.Errorf("get sequencer set failed, err: %v", err)
	}

	var seqz []*Sequencer
	for _, item := range seqSet {
		var addrTemp common.Address
		copy(addrTemp[:], item.MintAddress[:])
		votingPower := new(big.Int).Div(item.Amount, scale)
		seqz = append(seqz, NewSequencer(addrTemp, votingPower.Int64(), item.NodeID))
		log.Info("sequencer: ", "address", item.MintAddress.String(), "node_ID", hex.EncodeToString(item.NodeID))
	}

	// default epoch 1 day = 86400 second
	batchEpoch := defaultBatchEpoch * time.Second
	if config.BatchEpoch != 0 {
		batchEpoch = time.Duration(config.BatchEpoch) * time.Second
	}

	schedulerInst := &Scheduler{
		config:            config,
		running:           0,
		currentHeight:     0,
		db:                db,
		ticker:            time.NewTicker(batchEpoch),
		consensusEngine:   clique,
		eventMux:          eventMux,
		syncer:            syncer,
		schedulerAddr:     common.BytesToAddress(schedulerAddr[:]),
		sequencerSet:      NewSequencerSet(seqz),
		blockchain:        blockchain,
		txpool:            txpool,
		sequencerAssessor: NewHealthAssessor(),
		chainHeadCh:       make(chan core.ChainHeadEvent, chainHeadChanSize),
	}
	schedulerInst.setSequencerHealthPoints(seqSet)

	return schedulerInst, nil
}

// SetVerifiedTxCount set VerifiedTxCount function
func (schedulerInst *Scheduler) SetVerifiedTxCount(verifiedTxCount func() (uint64, error)) {
	schedulerInst.verifiedTxCount = verifiedTxCount
}

// SetWallet set signer and wallet, so that scheduler servier
func (schedulerInst *Scheduler) SetWallet(wallet accounts.Wallet, acc accounts.Account) {
	schedulerInst.wallet = wallet
	schedulerInst.signAccount = acc
}

// Scheduler return scheduler address
func (schedulerInst *Scheduler) Scheduler() common.Address {
	return schedulerInst.schedulerAddr
}

// CurrentStartMsg get the startMsg of this batch
func (schedulerInst *Scheduler) CurrentStartMsg() types.BatchPeriodStartMsg {
	return schedulerInst.currentStartMsg
}

// Start initializes and starts the sub-thread services of the scheduler
func (schedulerInst *Scheduler) Start() {
	if schedulerInst.wallet == nil || len(schedulerInst.signAccount.Address.Bytes()) == 0 {
		panic("Sequencer server need wallet to sign msgs")
	}

	// channel init
	schedulerInst.exitCh = make(chan struct{}, 1)
	schedulerInst.batchDone = make(chan uint64, 2)

	// Subscribe to events PeerAddEvent
	schedulerInst.addPeerSub = schedulerInst.eventMux.Subscribe(core.PeerAddEvent{})
	// Subscribe to events BatchEndEvent
	schedulerInst.batchEndSub = schedulerInst.eventMux.Subscribe(core.BatchEndEvent(0))
	// Subscribe to events chainHeadSub
	schedulerInst.chainHeadSub = schedulerInst.blockchain.SubscribeChainHeadEvent(schedulerInst.chainHeadCh)

	go schedulerInst.syncSequencerSetRoutine()
	go schedulerInst.addPeerCheckLoop()
	go schedulerInst.batchEndLoop()
	go schedulerInst.schedulerRoutine()
	go schedulerInst.handleChainHeadEventLoop()

	atomic.StoreInt32(&schedulerInst.running, 1)
}

// Stop set schedulerInst the indicator whether schedulerInst is running or not to 1
// and call the method Close to unsubscribe events and close channel exitCh
func (schedulerInst *Scheduler) Stop() {
	atomic.StoreInt32(&schedulerInst.running, 0)
	schedulerInst.Close()
}

// IsRunning returns an indicator whether schedulerInst is running or not
func (schedulerInst *Scheduler) IsRunning() bool {
	return atomic.LoadInt32(&schedulerInst.running) == 1
}

// Close cancel all subscriptions and close channel exitCh
func (schedulerInst *Scheduler) Close() {
	schedulerInst.chainHeadSub.Unsubscribe()
	schedulerInst.addPeerSub.Unsubscribe()
	schedulerInst.batchEndSub.Unsubscribe()
	close(schedulerInst.exitCh)
}

// addPeerCheckLoop will not be called before schedulerInst starts.
// When the miner starts, the schedulerInst service will be started synchronously
// and the scheduler will use addPeerCheckLoop to check the connected p2p nodes.
// The p2p connected after schedulerInst starts will call addPeerCheckLoop to check
// whether it is in the sequencer set
func (schedulerInst *Scheduler) addPeerCheckLoop() {
	// automatically stops if unsubscribe
	for obj := range schedulerInst.addPeerSub.Chan() {
		if ape, ok := obj.Data.(core.PeerAddEvent); ok {
			seqs := schedulerInst.sequencerSet.Sequencers
			find := false
			for _, v := range seqs {
				if bytes.Equal(v.NodeID, ape.PeerId) {
					find = true
					break
				}
			}
			ape.Has <- find
		}
	}
}

// batchEndLoop subscription BatchEndEvent then end this batch,
// this will happen when the batch has tx failed
func (schedulerInst *Scheduler) batchEndLoop() {
	// automatically stops if unsubscribe
	for obj := range schedulerInst.batchEndSub.Chan() {
		if e, ok := obj.Data.(core.BatchEndEvent); ok {
			schedulerInst.batchDone <- uint64(e)
		}
	}
}

// schedulerRoutine the main process of schedulerInst
// uses the batch parameters in the config to rotate each round of batches.
// Each round of batches will conduct a health check based on the results of
// the previous round of batches first, and then perform L1ToL2Tx processing.
// According to the sequencer set The priority selects the producer,
// generates BatchPeriodStartMsg and broadcasts, then updates the priority of the sequencer set.
// The batch will end only when it times out or receives batchDone signal,
// and the process will be interrupted when the exitCh signal is received
func (schedulerInst *Scheduler) schedulerRoutine() {
	batchSize := defaultBatchSize
	expireTime := defaultExpireTime
	if schedulerInst.config.BatchSize != 0 {
		batchSize = schedulerInst.config.BatchSize
	}
	if schedulerInst.config.BatchTime != 0 {
		expireTime = schedulerInst.config.BatchTime
	}
	for {
		// acceptance check last sequencer
		schedulerInst.checkSequencer()

		schedulerCh := make(chan struct{})
		err := schedulerInst.eventMux.Post(core.L1ToL2TxStartEvent{
			ErrCh:       nil,
			SchedulerCh: schedulerCh,
		})
		if err != nil {
			log.Error("generate BatchPeriodStartEvent error")
			return
		}
		select {
		case <-schedulerCh:
			log.Debug("produce block for L1ToL2Tx end", "block_number", schedulerInst.blockchain.CurrentBlock().Number().Uint64())
		}

		err = func() error {
			schedulerInst.sequencerSetMtx.Lock()
			defer schedulerInst.sequencerSetMtx.Unlock()

			if schedulerInst.sequencerSet == nil {
				return fmt.Errorf("empty sequencer set")
			}

			// get producer
			seq := schedulerInst.sequencerSet.getSeqWithMostPriority()
			currentBlock := schedulerInst.blockchain.CurrentBlock()
			currentIndex := rawdb.ReadStartMsgIndex(schedulerInst.db)

			msg := types.BatchPeriodStartMsg{
				RollbackStates: rawdb.ReadRollbackStates(schedulerInst.db),
				BatchIndex:     currentIndex + 1,
				StartHeight:    currentBlock.NumberU64() + 1,
				MaxHeight:      currentBlock.NumberU64() + uint64(batchSize),
				ExpireTime:     uint64(time.Now().Unix() + expireTime),
				Sequencer:      seq.Address,
			}
			signature, err := schedulerInst.wallet.SignData(schedulerInst.signAccount, accounts.MimetypeTypedData, msg.GetSignData())
			if err != nil {
				return fmt.Errorf("sign BatchPeriodStartEvent error %s", err.Error())
			}
			msg.Signature = signature
			expectMinTxsCount, err := schedulerInst.getExpectMinTxsCount(uint64(batchSize))
			if err != nil {
				return fmt.Errorf("get minimum block count failed %s", err.Error())
			}
			schedulerInst.expectMinTxsCount = expectMinTxsCount
			// store msg as currentStartMsg
			schedulerInst.currentStartMsg = msg
			// set BatchIndex to db
			rawdb.WriteCurrentBatchPeriodIndex(schedulerInst.db, msg.BatchIndex)

			// broadcast BatchPeriodStartMsg
			err = schedulerInst.eventMux.Post(core.BatchPeriodStartEvent{
				Msg:   &msg,
				ErrCh: nil,
			})
			if err != nil {
				return fmt.Errorf("generate BatchPeriodStartEvent error %s", err.Error())
			}
			schedulerInst.sequencerSet.IncrementProducerPriority(1)
			log.Info("Generate BatchPeriodStartEvent", "batch_index", msg.BatchIndex, "start_height", msg.StartHeight, "max_height", msg.MaxHeight)
			return nil
		}()

		if err != nil {
			log.Info("Generate BatchPeriodStartEvent failed", "err_msg", err.Error())
			continue
		}

		ticker := time.NewTicker(time.Duration(expireTime) * time.Second)
	loop:
		select {
		case <-ticker.C:
			log.Info("schedulerRoutine, ticker timeout")
		case index := <-schedulerInst.batchDone:
			if schedulerInst.CurrentStartMsg().BatchIndex != index {
				goto loop
			}
			log.Info("schedulerRoutine, batch done")
		case <-schedulerInst.exitCh:
			log.Info("schedulerRoutine stop")
			return
		}
	}
}

// handleChainHeadEventLoop checks whether the current block height reaches the currentStartMsg.maxHeight,
// if it reaches the maxHeight, send batchIndex to channel batchDone.
func (schedulerInst *Scheduler) handleChainHeadEventLoop() {
	for {
		select {
		case chainHead := <-schedulerInst.chainHeadCh:
			// Transactions of type QueueOriginL1ToL2 are not included in the batch
			if chainHead.Block.Transactions().Len() != 0 && chainHead.Block.Transactions()[0].GetMeta() != nil && chainHead.Block.Transactions()[0].QueueOrigin() == types.QueueOriginL1ToL2 {
				log.Debug("chainHead", "block_number", chainHead.Block.NumberU64(), "extra_data", hex.EncodeToString(chainHead.Block.Extra()))
				continue
			}
			log.Debug("schedulerInst handle chain head", "current_height", schedulerInst.blockchain.CurrentBlock().NumberU64(), "max_height", schedulerInst.currentStartMsg.MaxHeight)
			if schedulerInst.blockchain.CurrentBlock().NumberU64() == schedulerInst.currentStartMsg.MaxHeight {
				schedulerInst.batchDone <- schedulerInst.CurrentStartMsg().BatchIndex
			}
		case <-schedulerInst.exitCh:
			log.Info("scheduler chain head loop stop")
			return
		}
	}
}

// syncSequencerSetRoutine uses the batchEpoch parameter set in the config to regularly obtain
// the sequencer set from L1 and update the schedulerInst.sequencerSet
func (schedulerInst *Scheduler) syncSequencerSetRoutine() {
	for {
		select {
		case <-schedulerInst.ticker.C:
			err := func() error {
				log.Info("Sync sequencer set")
				schedulerInst.sequencerSetMtx.Lock()
				defer schedulerInst.sequencerSetMtx.Unlock()
				seqSet, err := schedulerInst.syncer.GetSequencerSet()
				if err != nil {
					return fmt.Errorf("get sequencer set failed %s", err.Error())
				}
				schedulerInst.setSequencerHealthPoints(seqSet)

				// todo do we need get changes then use UpdateWithChangeSet to update sequencer set or just replace schedulerInst.sequencerSet
				// get changes
				changes := compareSequencerSet(schedulerInst.sequencerSet.Sequencers, seqSet)
				log.Info(fmt.Sprintf("Get sequencer set success, have changes: %d", len(changes)))

				// update sequencer set and consensus_engine
				err = schedulerInst.sequencerSet.UpdateWithChangeSet(changes)
				if err != nil {
					return fmt.Errorf("sequencer set update failed %s", err.Error())
				}
				return err
			}()
			if err != nil {
				log.Error("syncSequencerSetRoutine update sequencer set error", "err_msg", err.Error())
				continue
			}
		case <-schedulerInst.exitCh:
			log.Info("Get scheduler stop signal, scheduler syncSequencerSetRoutine stop")
			return
		}
	}
}

// compareSequencerSet will return the update with Driver.seqz
func compareSequencerSet(preSeqs []*Sequencer, newSeq synchronizer.SequencerSequencerInfos) []*Sequencer {
	notDel := make(map[common.Address]bool)
	var tmp synchronizer.SequencerSequencerInfos
	for i, v := range newSeq {
		for _, seq := range preSeqs {
			power := new(big.Int).Div(v.Amount, scale)
			// figure out unchanged sequencer
			if bytes.Equal(seq.Address.Bytes(), v.MintAddress.Bytes()) {
				// find the sequencer in previous sequencer set
				notDel[seq.Address] = true
				if power.Int64() != seq.Power {
					// voting power changed sequencer update
					log.Debug("Compare sequencer set", "update_sequencer", newSeq[i].MintAddress.String())
					tmp = append(tmp, newSeq[i])
				}
				break
			}
		}
		// sequencer add
		if !notDel[common.Address(newSeq[i].MintAddress)] {
			log.Debug("Compare sequencer set", "add_sequencer", newSeq[i].MintAddress.String())
			tmp = append(tmp, newSeq[i])
		}
	}
	changes := bindToSeq(tmp)
	// select the deleted sequencer
	for _, v := range preSeqs {
		if !notDel[v.Address] {
			tmpDel := v
			tmpDel.Power = 0
			changes = append(changes, tmpDel)
			log.Debug("Compare sequencer set", "delete_sequencer", v.Address.String())
		}
	}
	return changes
}

// bindToSeq replaces the sequencer structure from SequencerSequencerInfos to Sequencer
func bindToSeq(binds synchronizer.SequencerSequencerInfos) []*Sequencer {
	var seqs []*Sequencer
	for _, v := range binds {
		seq := &Sequencer{
			Address: common.BytesToAddress(v.MintAddress.Bytes()),
			NodeID:  v.NodeID,
			Power:   new(big.Int).Div(v.Amount, scale).Int64(),
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
