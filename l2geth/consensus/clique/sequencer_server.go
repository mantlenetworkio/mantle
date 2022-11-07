package clique

import (
	"bytes"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/sequencer"
	"github.com/mantlenetworkio/mantle/l2geth/event"
	"github.com/mantlenetworkio/mantle/l2geth/log"
)

// NewSequencerBlockEvent is posted when sequencer set has been imported.
type ProducersUpdateEvent struct{ Producers *Producers }

type SequencerServer struct {
	wg  sync.WaitGroup
	mux *event.TypeMux

	engine *Clique
	ticker *time.Ticker
}

func NewSequencerServer(epoch time.Duration, clique *Clique, mux *event.TypeMux) *SequencerServer {
	return &SequencerServer{
		ticker: time.NewTicker(epoch),
		engine: clique,
		mux:    mux,
	}
}

func (seqS *SequencerServer) GetScheduler() ([]byte, error) {
	return sequencer.GetScheduler()
}

func (seqS *SequencerServer) Start() error {
	seqS.wg.Add(1)
	go seqS.readLoop()
	return nil
}

func (seqS *SequencerServer) readLoop() {
	defer seqS.wg.Done()
	for {
		select {
		case <-seqS.ticker.C:
			seqSet, err := sequencer.GetSequencerSet()
			if err != nil {
				log.Error("Get sequencer set failed, err : ", err)
				continue
			}
			var request GetProducers
			pros := seqS.engine.GetProducers(request)
			// get changes
			changes := CompareSequencerSet(pros.SequencerSet.Sequencers, seqSet)
			// update sequencer set
			pros.SequencerSet.UpdateWithChangeSet(changes)
			pros.increment()
			// Broadcast the producer and announce event
			seqS.mux.Post(ProducersUpdateEvent{Producers: &pros})
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
			power := v.Amount.Div(v.Amount, big.NewInt(scale)).Int64()
			if bytes.Equal(seq.Address.Bytes(), v.MintAddress.Bytes()) && power == seq.Power {
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
