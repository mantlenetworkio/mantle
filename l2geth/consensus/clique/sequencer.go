package clique

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/mantlenetworkio/mantle/l2geth/common"
)

// Volatile state for each Sequencer
// NOTE: The ProducerPriority is not included in Sequencer.Hash();
// make sure to update that method if changes are made here
type Sequencer struct {
	Address common.Address `json:"address"`
	Power   int64          `json:"power"`
	NodeID  []byte         `json:"node_id"`

	ProducerPriority int64 `json:"producer_priority"`
}

// NewSequencer returns a new sequencer with the given pubkey and power.
func NewSequencer(addr common.Address, power int64, nodeId []byte) *Sequencer {
	return &Sequencer{
		Address:          addr,
		Power:            power,
		NodeID:           nodeId,
		ProducerPriority: 0,
	}
}

// SequencerBasic performs basic validation.
func (s *Sequencer) SequencerBasic() error {
	if s == nil {
		return errors.New("nil sequencer")
	}

	if s.Power < 0 {
		return errors.New("sequencer has negative voting power")
	}

	return nil
}

// Creates a new copy of the sequencer so we can mutate ProducerPriority.
// Panics if the sequencer is nil.
func (s *Sequencer) Copy() *Sequencer {
	vCopy := *s
	return &vCopy
}

// Returns the one with higher ProducerPriority.
func (s *Sequencer) CompareProducerPriority(other *Sequencer) *Sequencer {
	if s == nil {
		return other
	}
	switch {
	case s.ProducerPriority > other.ProducerPriority:
		return s
	case s.ProducerPriority < other.ProducerPriority:
		return other
	default:
		result := bytes.Compare(s.Address.Bytes(), other.Address.Bytes())
		switch {
		case result < 0:
			return s
		case result > 0:
			return other
		default:
			panic("Cannot compare identical sequencers")
		}
	}
}

// String returns a string representation of String.
//
// 1. address
// 2. public key
// 3. voting power
// 4. proposer priority
func (s *Sequencer) String() string {
	if s == nil {
		return "nil-Sequencer"
	}
	return fmt.Sprintf("Sequencer{%v %v A:%v}",
		s.Address,
		s.Power,
		s.ProducerPriority)
}

// SequencerListString returns a prettified sequencer list for logging purposes.
func SequencerListString(seqs []*Sequencer) string {
	chunks := make([]string, len(seqs))
	for i, seq := range seqs {
		chunks[i] = fmt.Sprintf("%s:%d", seq.Address, seq.Power)
	}

	return strings.Join(chunks, ",")
}

//----------------------------------------
// RandSequencer

// RandSequencer returns a randomized validator, useful for testing.
// UNSTABLE
func RandSequencer(randPower bool, minPower int64) *Sequencer {
	power := minPower
	if randPower {
		power += int64(rand.Uint32())
	}
	var seed []byte
	rand.Read(seed)
	var addr common.Address
	copy(addr[:], seed)
	seq := NewSequencer(addr, power, []byte{})
	return seq
}
