package clique

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/accounts"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	"github.com/mantlenetworkio/mantle/l2geth/event"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

type BatchPeriodStartEvent struct{ Msg *BatchPeriodStart }

type BatchPeriodStart struct {
	BatchIndex   uint64
	MinerAddress common.Address
	StartHeight  uint64
	MaxHeight    uint64
	ExpireTime   uint64
	Signature    []byte
}

type BatchPeriodEndEvent struct{ Msg *BatchPeriodEnd }

type BatchPeriodEnd struct {
	BatchIndex   uint64
	MinerAddress common.Address
	StartHeight  uint64
	EndHeight    uint64
	Signatures   [][]byte
	Signature    []byte
}

type FraudProofReorgEvent struct{ Msg *FraudProofReorg }

type FraudProofReorg struct {
	Index         uint64
	reorgToHeight uint64
	tssSignature  []byte
}

// ProposersUpdateEvent is posted when sequencer set has been imported.
type ProposersUpdateEvent struct{ Update *ProposerUpdate }

type ProposerUpdate struct {
	Proposers Proposers
	Signature []byte
}

func (pro *ProposerUpdate) Serialize() []byte {
	return append(pro.Proposers.serialize(), pro.Signature...)
}

func (pro *ProposerUpdate) Deserialize(buf []byte) {
	tmp := deserialize(buf[:len(buf)-65])
	if tmp != nil {
		pro.Proposers = *tmp
		pro.Signature = buf[len(buf)-65:]
	} else {
		log.Error("Deserialize producerUpdate err got nil")
	}
}

type Scheduler struct {
	wg  sync.WaitGroup
	mux *event.TypeMux

	engine *Clique
	ticker *time.Ticker
	check  func() bool

	wallet      accounts.Wallet
	signAccount accounts.Account

	syncer *synchronizer.Synchronizer
}

func NewSequencerServer(epoch time.Duration, clique *Clique, mux *event.TypeMux, check func() bool) *Scheduler {
	log.Info("Create Sequencer Server")
	return &Scheduler{
		ticker: time.NewTicker(epoch * time.Second),
		engine: clique,
		mux:    mux,
		check:  check,
		syncer: synchronizer.NewSynchronizer(),
	}
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
	// check
	if schedulerInst.check == nil {
		panic("Sequencer server need method to check pre-preparation status")
	}
	if schedulerInst.wallet == nil || len(schedulerInst.signAccount.Address.Bytes()) == 0 {
		panic("Sequencer server need wallet to sign msgs")
	}

	// we need pre-preparation is ready first then we can start the server
	for times := 0; !schedulerInst.check(); times++ {
		log.Debug("Sequencer server pre-preparation is not ready, times : ", times)
	}
	schedulerInst.wg.Add(1)
	go schedulerInst.readLoop()
}

func (schedulerInst *Scheduler) readLoop() {
	defer schedulerInst.wg.Done()
	for {
		// we need pre-preparation is ready first then we can restart the server
		if !schedulerInst.check() {
			log.Debug("Sequencer server pre-preparation is not ready")
			return
		}
		select {
		case <-schedulerInst.ticker.C:
			seqSet, err := schedulerInst.syncer.GetSequencerSet()
			if err != nil {
				log.Error("Get sequencer set failed, err : ", err)
				continue
			}
			var request GetProducers
			proUpdate := schedulerInst.engine.GetProducers(request)
			pros := proUpdate.Proposers
			// get changes
			changes := CompareSequencerSet(pros.SequencerSet.Sequencers, seqSet)
			log.Debug(fmt.Sprintf("Get sequencer set success, have changes: %d", len(changes)))

			// todo : should it post every times? or post only have changes
			// update sequencer set and engine
			err = pros.SequencerSet.UpdateWithChangeSet(changes)
			if err != nil {
				log.Error(fmt.Sprintf("update sequencer set failed, err :%v ", err))
				continue
			}
			pros.increment()
			signature, err := schedulerInst.engine.signFn(schedulerInst.signAccount, accounts.MimetypeClique, pros.serialize())
			if err != nil {
				log.Error(fmt.Sprintf("Sign data error, err : %v ,Account address : %v ", err, schedulerInst.signAccount.Address.String()))
				continue
			}
			schedulerInst.engine.proposers = pros
			schedulerInst.engine.signature = signature
			// Broadcast the producer and announce event by post event
			schedulerInst.mux.Post(
				ProposersUpdateEvent{
					&ProposerUpdate{
						Proposers: pros,
						Signature: signature,
					},
				},
			)

		}
	}
}

// CompareSequencerSet will return the update with Driver.seqz
func CompareSequencerSet(old []*Sequencer, newSeq synchronizer.SequencerSequencerInfos) []*Sequencer {
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
