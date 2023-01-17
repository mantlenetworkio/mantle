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
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10

	defaultBatchSize  = int64(100)
	defaultExpireTime = int64(60)
	defaultBatchEpoch = time.Duration(86400)
)

var (
	scale = big.NewInt(1e18)
)

type Config struct {
	BatchSize  int64
	BatchTime  int64
	BatchEpoch int64
}

type Scheduler struct {
	sequencerSetMtx sync.Mutex // The lock used to protect the sequencerSet field
	batchMtx        sync.Mutex // The lock used to protect the batchEndFlag and batchDone fields

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

	// consensus channel
	batchDone       chan struct{}
	currentStartMsg types.BatchPeriodStartMsg
	currentHeight   uint64
	batchEndFlag    bool

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
		votingPower := big.NewInt(0).Div(item.Amount, scale)
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

func (schedulerInst *Scheduler) SetWallet(wallet accounts.Wallet, acc accounts.Account) {
	schedulerInst.wallet = wallet
	schedulerInst.signAccount = acc
}

func (schedulerInst *Scheduler) Scheduler() common.Address {
	return schedulerInst.schedulerAddr
}

func (schedulerInst *Scheduler) CurrentStartMsg() types.BatchPeriodStartMsg {
	return schedulerInst.currentStartMsg
}

func (schedulerInst *Scheduler) Start() {
	if schedulerInst.wallet == nil || len(schedulerInst.signAccount.Address.Bytes()) == 0 {
		panic("Sequencer server need wallet to sign msgs")
	}
	schedulerInst.exitCh = make(chan struct{}, 1)
	schedulerInst.batchDone = make(chan struct{}, 1)

	schedulerInst.addPeerSub = schedulerInst.eventMux.Subscribe(core.PeerAddEvent{})
	schedulerInst.batchEndSub = schedulerInst.eventMux.Subscribe(core.BatchEndEvent{})
	schedulerInst.chainHeadSub = schedulerInst.blockchain.SubscribeChainHeadEvent(schedulerInst.chainHeadCh)

	go schedulerInst.syncSequencerSetRoutine()
	go schedulerInst.addPeerCheckLoop()
	go schedulerInst.batchEndLoop()
	go schedulerInst.schedulerRoutine()
	go schedulerInst.handleChainHeadEventLoop()

	atomic.StoreInt32(&schedulerInst.running, 1)
}

func (schedulerInst *Scheduler) Stop() {
	atomic.StoreInt32(&schedulerInst.running, 0)
	schedulerInst.Close()
}

// IsRunning returns an indicator whether schedulerInst is running or not.
func (schedulerInst *Scheduler) IsRunning() bool {
	return atomic.LoadInt32(&schedulerInst.running) == 1
}

func (schedulerInst *Scheduler) Close() {
	schedulerInst.chainHeadSub.Unsubscribe()
	schedulerInst.addPeerSub.Unsubscribe()
	schedulerInst.batchEndSub.Unsubscribe()
	close(schedulerInst.exitCh)
}

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

func (schedulerInst *Scheduler) batchEndLoop() {
	// automatically stops if unsubscribe
	for obj := range schedulerInst.batchEndSub.Chan() {
		if _, ok := obj.Data.(core.BatchEndEvent); ok {
			// if batch already exitCh with timeout or height at max height then pass
			func() {
				schedulerInst.batchMtx.Lock()
				defer schedulerInst.batchMtx.Unlock()
				if !schedulerInst.batchEndFlag {
					schedulerInst.batchDone <- struct{}{}
					schedulerInst.batchEndFlag = true
				} else {
					log.Debug("Batch already exitCh with timeout or height at max height")
				}
			}()
		}
	}
}

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
			schedulerInst.currentStartMsg = msg
			rawdb.WriteCurrentBatchPeriodIndex(schedulerInst.db, msg.BatchIndex)

			// clean channel batchDone
			select {
			case <-schedulerInst.batchDone:
			default:
			}

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

		func() {
			schedulerInst.batchMtx.Lock()
			defer schedulerInst.batchMtx.Unlock()
			schedulerInst.batchEndFlag = false
		}()

		ticker := time.NewTicker(time.Duration(expireTime) * time.Second)
		select {
		case <-ticker.C:
			log.Info("schedulerRoutine, ticker timeout")
		case <-schedulerInst.batchDone:
			log.Info("schedulerRoutine, batch done")
		case <-schedulerInst.exitCh:
			log.Info("schedulerRoutine stop")
			return
		}
	}
}

func (schedulerInst *Scheduler) handleChainHeadEventLoop() {
	for {
		select {
		case chainHead := <-schedulerInst.chainHeadCh:
			if chainHead.Block.Transactions().Len() != 0 && chainHead.Block.Transactions()[0].GetMeta() != nil && chainHead.Block.Transactions()[0].QueueOrigin() == types.QueueOriginL1ToL2 {
				log.Debug("chainHead", "block_number", chainHead.Block.NumberU64(), "extra_data", hex.EncodeToString(chainHead.Block.Extra()))
				continue
			}
			log.Info("schedulerInst handle chain head", "current_height", schedulerInst.blockchain.CurrentBlock().NumberU64(), "max_height", schedulerInst.currentStartMsg.MaxHeight)
			if schedulerInst.blockchain.CurrentBlock().NumberU64() == schedulerInst.currentStartMsg.MaxHeight {
				func() {
					schedulerInst.batchMtx.Lock()
					defer schedulerInst.batchMtx.Unlock()
					if !schedulerInst.batchEndFlag {
						log.Info("Batch done with height at max height")
						schedulerInst.batchDone <- struct{}{}
						schedulerInst.batchEndFlag = true
					} else {
						log.Debug("Batch already done with tx apply failed")
					}
				}()
			}
		case <-schedulerInst.exitCh:
			log.Info("scheduler chain head loop stop")
			return
		}
	}
}

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
				log.Error("update sequencer set error", "err_msg", err.Error())
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
	var notDel map[common.Address]bool
	var tmp synchronizer.SequencerSequencerInfos
	for i, v := range newSeq {
		changed := true
		for _, seq := range preSeqs {
			power := big.NewInt(1).Div(v.Amount, scale)
			// found the sequencer and their power equal
			if bytes.Equal(seq.Address.Bytes(), v.MintAddress.Bytes()) {
				// find the sequencer in previous sequencer set
				notDel[seq.Address] = true
				if power.Int64() == seq.Power {
					// voting power have not change
					changed = false
					break
				}
			}
		}
		// sequencer add or update
		if changed {
			tmp = append(tmp, newSeq[i])
		}
	}
	changes := bindToSeq(tmp)
	// select the deleted sequencer
	for _, v := range preSeqs {
		del := true
		for addr, _ := range notDel {
			// find in new sequencer set, this sequencer still in sequencer set
			if bytes.Equal(addr.Bytes(), v.Address.Bytes()) {
				del = false
				break
			}
		}
		// not find in new sequencer set, should set the sequencer power to 0
		if del {
			tmpDel := v
			tmpDel.Power = 0
			changes = append(changes, tmpDel)
		}
	}
	return changes
}

func bindToSeq(binds synchronizer.SequencerSequencerInfos) []*Sequencer {
	var seqs []*Sequencer
	for _, v := range binds {
		seq := &Sequencer{
			Address: common.BytesToAddress(v.MintAddress.Bytes()),
			NodeID:  v.NodeID,
			Power:   big.NewInt(0).Div(v.Amount, scale).Int64(),
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
