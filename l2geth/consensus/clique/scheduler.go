package clique

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"

	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	"github.com/mantlenetworkio/mantle/l2geth/event"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
)

type Scheduler struct {
	wg       sync.WaitGroup
	eventMux *event.TypeMux

	sequencerSet    *SequencerSet
	consensusEngine *Clique
	blockchain      *core.BlockChain

	chainHeadSub event.Subscription
	chainHeadCh  chan core.ChainHeadEvent

	ticker *time.Ticker

	wallet      accounts.Wallet
	signAccount accounts.Account

	syncer *synchronizer.Synchronizer
}

func NewScheduler(epoch time.Duration, clique *Clique, blockchain *core.BlockChain, eventMux *event.TypeMux) (*Scheduler, error) {
	log.Info("Create Sequencer Server")

	syncer := synchronizer.NewSynchronizer()
	seqSet, err := syncer.GetSequencerSet()
	if err != nil {
		return nil, err
	}

	var seqz []*Sequencer
	for _, item := range seqSet {
		var addrTemp common.Address
		copy(addrTemp[:], item.MintAddress[:])
		votingPower := big.NewInt(0).Div(item.Amount, big.NewInt(1e16))
		seqz = append(seqz, NewSequencer(addrTemp, votingPower.Int64()))
		log.Info("sequencer: ", "address", item.MintAddress.String())
	}

	if err != nil {
		return nil, fmt.Errorf("get sequencer set failed, err: %v", err)
	}

	return &Scheduler{
		ticker:          time.NewTicker(epoch * time.Second),
		consensusEngine: clique,
		eventMux:        eventMux,
		syncer:          syncer,
		sequencerSet:    NewSequencerSet(seqz),
		blockchain:      blockchain,
		chainHeadCh:     make(chan core.ChainHeadEvent, chainHeadChanSize),
	}, nil
}

func (schedulerInst *Scheduler) SetWallet(wallet accounts.Wallet, acc accounts.Account) {
	schedulerInst.wallet = wallet
	schedulerInst.signAccount = acc
}

func (schedulerInst *Scheduler) GetScheduler() (common.Address, error) {
	scheduler, err := schedulerInst.syncer.GetSchedulerAddr()
	if err != nil {
		return common.BigToAddress(common.Big0), err
	}
	return common.BytesToAddress(scheduler.Bytes()), nil
}

func (schedulerInst *Scheduler) Start() {
	if schedulerInst.wallet == nil || len(schedulerInst.signAccount.Address.Bytes()) == 0 {
		panic("Sequencer server need wallet to sign msgs")
	}
	schedulerInst.chainHeadSub = schedulerInst.blockchain.SubscribeChainHeadEvent(schedulerInst.chainHeadCh)

	schedulerInst.wg.Add(1)
	go schedulerInst.readLoop()
	go schedulerInst.schedulerRoutine()
	go schedulerInst.handleChainHeadEventLoop()
}

func (schedulerInst *Scheduler) Close() {
	schedulerInst.chainHeadSub.Unsubscribe()
	close(schedulerInst.chainHeadCh)
}

func (schedulerInst *Scheduler) schedulerRoutine() {
	sequencerSet := []common.Address{
		//common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
		common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8"),
		common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc"),
	}
	mockSignature := common.Hex2Bytes("2020a0bbf67b08b1df594333c1ead3a771d9742d2f33798e050da744b1255bb67860d672a5055429cc53d17e6c57550989b39cf997e2fb58d1ec6aae198a471501")

	batchSize := uint64(10) // 10 transaction in one batch
	expireTime := int64(15) // 15s
	for {
		randomIdx := rand.Intn(len(sequencerSet))
		sequencerAddr := sequencerSet[randomIdx]

		currentBlock := schedulerInst.blockchain.CurrentBlock()

		msg := types.BatchPeriodStartMsg{
			ReorgIndex:   0,
			BatchIndex:   1,
			StartHeight:  currentBlock.NumberU64() + 1,
			MaxHeight:    currentBlock.NumberU64() + 1 + batchSize,
			ExpireTime:   uint64(time.Now().Unix() + expireTime),
			MinerAddress: sequencerAddr,
			SequencerSet: sequencerSet,
			Signature:    mockSignature,
		}

		// index block and sign L1toL2Tx here?
		// for tx in currentBlock.Transactions() {
		// if tx.QueueOrigin() == types.QueueOriginL1ToL2 {
		// tx.metadata.schedulerSignature = schedulerInst.wallet.SignTx(schedulerInst.signingAccount, tx, schedulerInst.blockchain.Config().ChainID)
		//}

		log.Info("Start Post BatchPeriodStartEvent")
		err := schedulerInst.eventMux.Post(core.BatchPeriodStartEvent{
			Msg:   &msg,
			ErrCh: nil,
		})
		log.Info("End Post BatchPeriodStartEvent")
		if err != nil {
			log.Error("generate BatchPeriodStartEvent error")
			return
		} else {
			log.Info("generate BatchPeriodStartEvent success", "startHeight", msg.StartHeight, "maxHeight", msg.MaxHeight, "expireTime", msg.ExpireTime, "minerAddress", msg.MinerAddress)
		}
		ticker := time.NewTicker(time.Duration(expireTime) * time.Second)
		select {
		case <-schedulerInst.ticker.C:
			log.Info("ticker timeout")
			ticker.Stop()
		}
	}
}

func (schedulerInst *Scheduler) handleChainHeadEventLoop() {
	for {
		select {
		case chainHead := <-schedulerInst.chainHeadCh:
			log.Debug("chainHead", "block number", chainHead.Block.NumberU64(), "extra data", hex.EncodeToString(chainHead.Block.Extra()))
		}
	}
}

func (schedulerInst *Scheduler) readLoop() {
	defer schedulerInst.wg.Done()
	for {
		select {
		case <-schedulerInst.ticker.C:
			seqSet, err := schedulerInst.syncer.GetSequencerSet()
			if err != nil {
				log.Error("Get sequencer set failed, err : ", err)
				continue
			}
			// get changes
			changes := compareSequencerSet(schedulerInst.sequencerSet.Sequencers, seqSet)
			log.Debug(fmt.Sprintf("Get sequencer set success, have changes: %d", len(changes)))

			// todo : should it post every times? or post only have changes
			// update sequencer set and consensus_engine
			err = schedulerInst.sequencerSet.UpdateWithChangeSet(changes)
			if err != nil {
				log.Error(fmt.Sprintf("update sequencer set failed, err :%v", err))
				continue
			}
		}
	}
}

// compareSequencerSet will return the update with Driver.seqz
func compareSequencerSet(old []*Sequencer, newSeq synchronizer.SequencerSequencerInfos) []*Sequencer {
	var tmp synchronizer.SequencerSequencerInfos
	// voting power = deposit / scale (10^18)
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
			// PubKey:  nil,
			Power: v.Amount.Div(v.Amount, big.NewInt(scale)).Int64(),
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
