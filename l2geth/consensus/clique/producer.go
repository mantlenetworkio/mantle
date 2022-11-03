package clique

import (
	"encoding/binary"
	"encoding/json"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
)

var (
	extraNumberLength    = 8
	extraEpochLength     = 8
	extraSchedulerLength = common.AddressLength
	extraSequencerLength = common.AddressLength + 8 + 8 // each sequencer serialize by address + power + priority
)

type Producers struct {
	Number       uint64       `json:"number"`       // Block number where the snapshot was created
	Epoch        uint64       `json:"epoch"`        // Epoch represents the block number for each producer
	SchedulerID  []byte       `json:"schedulerID"`  // SchedulerID represents scheduler's peer.id
	SequencerSet SequencerSet `json:"sequencerSet"` // Set of sequencers
}

// GetProducers represents a producers query.
type GetProducers struct{}

// newProducers creates a new ProducersData.
func newProducers(number uint64, epoch uint64, schedulerID []byte, sequencerSet SequencerSet) *Producers {
	data := &Producers{
		Number:       number,
		Epoch:        epoch,
		SchedulerID:  schedulerID,
		SequencerSet: sequencerSet,
	}
	return data
}

// loadProducers loads an existing producers from the database.
func loadProducers(db ethdb.Database, number uint64) (*Producers, error) {

	blob, err := db.Get(append([]byte("coterie-"), Uint64ToBytes(number)[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Producers)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// store inserts the producersData into the database.
func (s *Producers) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("coterie-"), Uint64ToBytes(s.Number)[:]...), blob)
}

func (s *Producers) GetSequencerSet() SequencerSet {
	return s.SequencerSet
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Producers) inturn(signer common.Address) bool {
	return s.SequencerSet.GetProducer().Address == signer
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func (s *Producers) serialize() []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, s.Number)

	buf = binary.BigEndian.AppendUint64(buf, s.Epoch)
	buf = append(buf, s.SchedulerID...)

	for _, sequencer := range s.SequencerSet.Sequencers {
		buf = append(buf, sequencer.Address.Bytes()...)
		buf = binary.BigEndian.AppendUint64(buf, uint64(sequencer.Power))
		buf = binary.BigEndian.AppendUint64(buf, uint64(sequencer.ProducerPriority))
	}

	return buf
}

func deserialize(buf []byte) *Producers {
	if len(buf) < extraNumberLength+extraEpochLength+extraSchedulerLength {
		return nil
	}

	if (len(buf)-extraNumberLength-extraEpochLength-extraSchedulerLength)%extraSequencerLength != 0 {
		return nil
	}

	number := binary.BigEndian.Uint64(buf[:extraNumberLength])
	epoch := binary.BigEndian.Uint64(buf[extraNumberLength : extraNumberLength+extraEpochLength])
	schedulerID := buf[extraNumberLength+extraEpochLength : extraNumberLength+extraEpochLength+extraSchedulerLength]

	sequencers := make([]*Sequencer, 0)
	for i := extraNumberLength + extraEpochLength + extraSchedulerLength; i < len(buf); i += extraSequencerLength {
		sequencer := &Sequencer{
			Address:          common.BytesToAddress(buf[i : i+common.AddressLength]),
			Power:            int64(binary.BigEndian.Uint64(buf[i+common.AddressLength : i+common.AddressLength+8])),
			ProducerPriority: int64(binary.BigEndian.Uint64(buf[i+common.AddressLength+8 : i+common.AddressLength+8+8])),
		}
		sequencers = append(sequencers, sequencer)
	}

	sequencerSet := &SequencerSet{
		Sequencers: sequencers,
	}

	sequencerSet.updateTotalPower()
	sequencerSet.Producer = sequencerSet.findProducer()

	return newProducers(number, epoch, schedulerID, *sequencerSet)
}
