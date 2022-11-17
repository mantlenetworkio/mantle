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
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/sequencer"
	"github.com/mantlenetworkio/mantle/l2geth/event"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

// ProducersUpdateEvent is posted when sequencer set has been imported.
type ProducersUpdateEvent struct{ Update *ProducerUpdate }

type ProducerUpdate struct {
	Producers *Producers
	Signature []byte
}

func (pro *ProducerUpdate) Serialize() []byte {
	return append(pro.Producers.serialize(), pro.Signature...)
}

func (pro *ProducerUpdate) Deserialize(buf []byte) {
	pro.Producers = deserialize(buf[:len(buf)-65])
	pro.Signature = buf[len(buf)-65:]
}

type SequencerServer struct {
	wg  sync.WaitGroup
	mux *event.TypeMux

	engine *Clique
	ticker *time.Ticker
	check  func() bool

	wallet      accounts.Wallet
	signAccount accounts.Account
}

func NewSequencerServer(epoch time.Duration, clique *Clique, mux *event.TypeMux, check func() bool) *SequencerServer {
	log.Info("Create Sequencer Server")
	sequencer.Initialize()
	return &SequencerServer{
		ticker: time.NewTicker(epoch * time.Second),
		engine: clique,
		mux:    mux,
		check:  check,
	}
}

func (seqS *SequencerServer) SetWallet(wallet accounts.Wallet, acc accounts.Account) {
	seqS.wallet = wallet
	seqS.signAccount = acc
}

func (seqS *SequencerServer) GetScheduler() (common.Address, error) {
	scheduler, err := sequencer.GetScheduler()
	if err != nil {
		return common.BigToAddress(common.Big0), err
	}
	return common.BytesToAddress(scheduler.Bytes()), nil
}

func (seqS *SequencerServer) Start() {
	// check
	if seqS.check == nil {
		panic("Sequencer server need method to check pre-preparation status")
	}
	if seqS.wallet == nil || len(seqS.signAccount.Address.Bytes()) == 0 {
		panic("Sequencer server need wallet to sign msgs")
	}

	// we need pre-preparation is ready first then we can start the server
	for times := 0; !seqS.check(); times++ {
		log.Debug("Sequencer server pre-preparation is not ready, times : ", times)
	}
	seqS.wg.Add(1)
	go seqS.readLoop()
}

func (seqS *SequencerServer) readLoop() {
	defer seqS.wg.Done()
	for {
		// we need pre-preparation is ready first then we can restart the server
		if !seqS.check() {
			log.Debug("Sequencer server pre-preparation is not ready")
			return
		}
		select {
		case <-seqS.ticker.C:
			seqSet, err := sequencer.GetSequencerSet()
			if err != nil {
				log.Error("Get sequencer set failed, err : ", err)
				continue
			}
			var request GetProducers
			proUpdate := seqS.engine.GetProducers(request)
			pros := proUpdate.Producers
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
			signature, err := seqS.engine.signFn(seqS.signAccount, accounts.MimetypeClique, pros.serialize())
			if err != nil {
				log.Error(fmt.Sprintf("Sign data error, err : %v ,Account address : %v ", err, seqS.signAccount.Address.String()))
				continue
			}
			seqS.engine.producers = pros
			seqS.engine.signature = signature
			// Broadcast the producer and announce event by post event
			seqS.mux.Post(
				ProducersUpdateEvent{
					&ProducerUpdate{
						Producers: pros,
						Signature: signature,
					},
				},
			)

		}
	}
}

// CompareSequencerSet will return the update with Driver.seqz
func CompareSequencerSet(old []*Sequencer, newSeq sequencer.SequencerSequencerInfos) []*Sequencer {
	var tmp sequencer.SequencerSequencerInfos
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

func bindToSeq(binds sequencer.SequencerSequencerInfos) []*Sequencer {
	scale := int64(math.Pow10(18))
	var seqs []*Sequencer
	for _, v := range binds {
		// todo: pubkey check
		seq := &Sequencer{
			Address: common.BytesToAddress(v.MintAddress.Bytes()),
			// PubKey:  nil,
			Power: v.Amount.Div(v.Amount, big.NewInt(scale)).Int64(),
		}
		seqs = append(seqs, seq)
	}
	return seqs
}
