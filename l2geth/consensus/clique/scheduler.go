package clique

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
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
)

type Config struct {
	BatchSize int64
	BatchTime int64
}

type Scheduler struct {
	wg       sync.WaitGroup
	l        sync.Mutex
	eventMux *event.TypeMux
	done     chan struct{}
	running  int32

	config *Config

	db ethdb.Database

	schedulerAddr   common.Address
	sequencerSet    *SequencerSet
	consensusEngine *Clique
	blockchain      *core.BlockChain

	// consensus channel
	batchDone       chan struct{}
	currentStartMsg types.BatchPeriodStartMsg
	currentHeight   uint64

	chainHeadSub event.Subscription
	chainHeadCh  chan core.ChainHeadEvent
	addPeerSub   *event.TypeMuxSubscription

	ticker *time.Ticker

	wallet      accounts.Wallet
	signAccount accounts.Account

	syncer *synchronizer.Synchronizer
}

func NewScheduler(db ethdb.Database, config *Config, schedulerAddress common.Address, clique *Clique, blockchain *core.BlockChain, eventMux *event.TypeMux) (*Scheduler, error) {
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
		return nil, err
	}

	var seqz []*Sequencer
	for _, item := range seqSet {
		var addrTemp common.Address
		copy(addrTemp[:], item.MintAddress[:])
		votingPower := big.NewInt(0).Div(item.Amount, big.NewInt(1e16))
		seqz = append(seqz, NewSequencer(addrTemp, votingPower.Int64(), item.NodeID))
		log.Info("sequencer: ", "address", item.MintAddress.String(), "nodeID", hex.EncodeToString(item.NodeID))
	}

	if err != nil {
		return nil, fmt.Errorf("get sequencer set failed, err: %v", err)
	}
	schedulerInst := &Scheduler{
		config:          config,
		running:         0,
		currentHeight:   0,
		db:              db,
		done:            make(chan struct{}, 1),
		batchDone:       make(chan struct{}, 1),
		ticker:          time.NewTicker(10 * time.Second), //TODO
		consensusEngine: clique,
		eventMux:        eventMux,
		syncer:          syncer,
		schedulerAddr:   common.BytesToAddress(schedulerAddr[:]),
		sequencerSet:    NewSequencerSet(seqz),
		blockchain:      blockchain,
		chainHeadCh:     make(chan core.ChainHeadEvent, chainHeadChanSize),
	}

	schedulerInst.addPeerSub = schedulerInst.eventMux.Subscribe(core.PeerAddEvent{})
	go schedulerInst.AddPeerCheck()
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
	schedulerInst.chainHeadSub = schedulerInst.blockchain.SubscribeChainHeadEvent(schedulerInst.chainHeadCh)
	atomic.StoreInt32(&schedulerInst.running, 1)

	schedulerInst.wg.Add(1)
	go schedulerInst.readLoop()
	go schedulerInst.schedulerRoutine()
	go schedulerInst.handleChainHeadEventLoop()
}

func (schedulerInst *Scheduler) Stop() {
	atomic.StoreInt32(&schedulerInst.running, 0)
	schedulerInst.done <- struct{}{}
}

// IsRunning returns an indicator whether schedulerInst is running or not.
func (schedulerInst *Scheduler) IsRunning() bool {
	return atomic.LoadInt32(&schedulerInst.running) == 1
}

func (schedulerInst *Scheduler) Close() {
	schedulerInst.chainHeadSub.Unsubscribe()
	schedulerInst.addPeerSub.Unsubscribe()
	close(schedulerInst.chainHeadCh)
}

func (schedulerInst *Scheduler) AddPeerCheck() {
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
			log.Debug("produce block for L1ToL2Tx end", "blockNumber", schedulerInst.blockchain.CurrentBlock().Number().Uint64())
		}

		schedulerInst.l.Lock()

		if schedulerInst.sequencerSet == nil {
			continue
		}
		seq := schedulerInst.sequencerSet.getSeqWithMostPriority()
		var seqSet []common.Address
		for _, v := range schedulerInst.sequencerSet.Sequencers {
			seqSet = append(seqSet, v.Address)
		}

		currentBlock := schedulerInst.blockchain.CurrentBlock()
		currentIndex := rawdb.ReadStartMsgIndex(schedulerInst.db)
		log.Debug("now index ", "index", currentIndex)
		msg := types.BatchPeriodStartMsg{
			ReorgIndex:  0,
			BatchIndex:  currentIndex + 1,
			StartHeight: currentBlock.NumberU64() + 1,
			MaxHeight:   currentBlock.NumberU64() + 1 + uint64(batchSize),
			ExpireTime:  uint64(time.Now().Unix() + expireTime),
			Sequencer:   seq.Address,
		}
		sign, err := schedulerInst.wallet.SignData(schedulerInst.signAccount, accounts.MimetypeTypedData, msg.GetSignData())
		if err != nil {
			log.Error("sign BatchPeriodStartEvent error")
			return
		}
		msg.Signature = sign
		schedulerInst.currentStartMsg = msg
		rawdb.WriteCurrentBatchPeriodIndex(schedulerInst.db, msg.BatchIndex)

		err = schedulerInst.eventMux.Post(core.BatchPeriodStartEvent{
			Msg:   &msg,
			ErrCh: nil,
		})
		if err != nil {
			log.Error("generate BatchPeriodStartEvent error")
			return
		}
		schedulerInst.sequencerSet.IncrementProducerPriority(1)
		schedulerInst.l.Unlock()
		ticker := time.NewTicker(time.Duration(expireTime) * time.Second)
		select {
		case <-ticker.C:
			log.Debug("ticker timeout")
		case <-schedulerInst.batchDone:
			log.Debug("batch done")
		}
	}
}

func (schedulerInst *Scheduler) handleChainHeadEventLoop() {
	for {
		select {
		case chainHead := <-schedulerInst.chainHeadCh:
			if chainHead.Block.Transactions().Len() != 0 && chainHead.Block.Transactions()[0].GetMeta() != nil && chainHead.Block.Transactions()[0].QueueOrigin() == types.QueueOriginL1ToL2 {
				log.Debug("chainHead", "blockNumber", chainHead.Block.NumberU64(), "extraData", hex.EncodeToString(chainHead.Block.Extra()))
				continue
			}
			if schedulerInst.blockchain.CurrentBlock().NumberU64() == schedulerInst.currentStartMsg.MaxHeight {
				log.Debug("Batch done with height at max height")
				schedulerInst.batchDone <- struct{}{}
			}
			log.Debug("chainHead", "blockNumber", chainHead.Block.NumberU64(), "extraData", hex.EncodeToString(chainHead.Block.Extra()))
		}
	}
}

func (schedulerInst *Scheduler) readLoop() {
	defer schedulerInst.wg.Done()
	for {
		select {
		case <-schedulerInst.ticker.C:
			schedulerInst.l.Lock()
			seqSet, err := schedulerInst.syncer.GetSequencerSet()
			if err != nil {
				log.Error("Get sequencer set failed, err : ", err)
				continue
			}
			// get changes
			changes := compareSequencerSet(schedulerInst.sequencerSet.Sequencers, seqSet)
			log.Debug(fmt.Sprintf("Get sequencer set success, have changes: %d", len(changes)))

			// update sequencer set and consensus_engine
			err = schedulerInst.sequencerSet.UpdateWithChangeSet(changes)
			if err != nil {
				log.Error("sequencer set update failed", "err", err)
				continue
			}
			schedulerInst.l.Unlock()
		case <-schedulerInst.done:
			log.Info("Get scheduler stop signal")
			return
		}
	}
}

// compareSequencerSet will return the update with Driver.seqz
func compareSequencerSet(old []*Sequencer, newSeq synchronizer.SequencerSequencerInfos) []*Sequencer {
	var tmp synchronizer.SequencerSequencerInfos
	scale := int64(math.Pow10(18))
	for i, v := range newSeq {
		changed := true
		for _, seq := range old {
			power := big.NewInt(v.Amount.Int64())
			power = power.Div(power, big.NewInt(scale))
			if bytes.Equal(seq.Address.Bytes(), v.MintAddress.Bytes()) && power.Int64() == seq.Power {
				changed = false
				break
			}
		}
		if changed {
			tmp = append(tmp, newSeq[i])
		}
	}
	changes := bindToSeq(tmp)
	return changes
}

func bindToSeq(binds synchronizer.SequencerSequencerInfos) []*Sequencer {
	scale := int64(math.Pow10(18))
	var seqs []*Sequencer
	for _, v := range binds {
		seq := &Sequencer{
			Address: common.BytesToAddress(v.MintAddress.Bytes()),
			NodeID:  v.NodeID,
			Power:   v.Amount.Div(v.Amount, big.NewInt(scale)).Int64(),
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
