package clique

import (
	"encoding/binary"
	"encoding/json"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
)

type Proposers struct {
	Number       uint64       `json:"number"`       // Block number where the snapshot was created
	Index        uint64       `json:"index"`        // Index of proposers update times
	Epoch        uint64       `json:"epoch"`        // Epoch represents the block number for each producer
	SchedulerID  []byte       `json:"schedulerID"`  // SchedulerID represents scheduler's peer.id
	SequencerSet SequencerSet `json:"sequencerSet"` // Set of sequencers
}

// GetProducers represents a proposers query.
type GetProducers struct{}

// newProposers creates a new ProducersData.
func newProposers(number, index, epoch uint64, schedulerID []byte, sequencerSet SequencerSet) *Proposers {
	data := &Proposers{
		Number:       number,
		Index:        index,
		Epoch:        epoch,
		SchedulerID:  schedulerID,
		SequencerSet: sequencerSet,
	}
	return data
}

// loadProducers loads an existing proposers from the database.
func loadProducers(db ethdb.Database, number uint64) (*Proposers, error) {

	blob, err := db.Get(append([]byte("coterie-"), Uint64ToBytes(number)[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Proposers)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// store inserts the producersData into the database.
func (s *Proposers) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("coterie-"), Uint64ToBytes(s.Number)[:]...), blob)
}

func (s *Proposers) GetSequencerSet() SequencerSet {
	return s.SequencerSet
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Proposers) inturn(signer common.Address) bool {
	return s.SequencerSet.GetProducer().Address == signer
}

// increment index
func (s *Proposers) increment() {
	// todo clear index each epoch?
	s.Index++
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}
