package coterie

import (
	"encoding/json"

	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/ethdb"
)

type Snapshot struct {
	Number       uint64       `json:"number"`       // Block number where the snapshot was created
	Hash         common.Hash  `json:"hash"`         // Block hash where the snapshot was created
	SequencerSet SequencerSet `json:"sequencerSet"` // Set of sequencers
	Epoch        uint64       `json:"epoch"`
}

// newSnapshot creates a new snapshot.
func newSnapshot(number uint64, hash common.Hash, sequencerSet SequencerSet, epoch uint64) *Snapshot {
	snap := &Snapshot{
		Number:       number,
		Hash:         hash,
		SequencerSet: sequencerSet,
		Epoch:        epoch,
	}
	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	blob, err := db.Get(append([]byte("coterie-"), hash[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("coterie-"), s.Hash[:]...), blob)
}

func (s *Snapshot) sequencerSet() SequencerSet {
	return s.SequencerSet
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Snapshot) inturn(signer common.Address) bool {
	return s.SequencerSet.GetProducer().Address == signer
}
