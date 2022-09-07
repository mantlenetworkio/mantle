package coterie

import (
	"encoding/binary"
	"encoding/json"

	"github.com/bitdao-io/mantle/l2geth/common"
	"github.com/bitdao-io/mantle/l2geth/ethdb"
)

type ProducersData struct {
	Number       uint64       `json:"number"`       // Block number where the snapshot was created
	Epoch        uint64       `json:"epoch"`        // Epoch represents the block number for each producer
	SchedulerID  string       `json:"schedulerID"`  // SchedulerID represents scheduler's peer.id
	SequencerSet SequencerSet `json:"sequencerSet"` // Set of sequencers
}

// GetProducersData represents a producers query.
type GetProducersData struct{}

// newProducersData creates a new ProducersData.
func newProducersData(number uint64, epoch uint64, schedulerID string, sequencerSet SequencerSet) *ProducersData {
	data := &ProducersData{
		Number:       number,
		Epoch:        epoch,
		SchedulerID:  schedulerID,
		SequencerSet: sequencerSet,
	}
	return data
}

// loadSnapshot loads an existing producersData from the database.
func loadSnapshot(db ethdb.Database, number uint64) (*ProducersData, error) {

	blob, err := db.Get(append([]byte("coterie-"), Uint64ToBytes(number)[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(ProducersData)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// store inserts the producersData into the database.
func (s *ProducersData) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("coterie-"), Uint64ToBytes(s.Number)[:]...), blob)
}

func (s *ProducersData) sequencerSet() SequencerSet {
	return s.SequencerSet
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *ProducersData) inturn(signer common.Address) bool {
	return s.SequencerSet.GetProducer().Address == signer
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}
