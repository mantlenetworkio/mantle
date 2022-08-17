package scheduler

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"strings"
)

// Volatile state for each Sequencer
// NOTE: The ProducerPriority is not included in Sequencer.Hash();
// make sure to update that method if changes are made here
type Sequencer struct {
	Address     common.Address  `json:"address"`
	PubKey      ecdsa.PublicKey `json:"pub_key"`
	VotingPower int64           `json:"voting_power"`

	ProducerPriority int64 `json:"producer_priority"`
}

// NewSequencer returns a new sequencer with the given pubkey and voting power.
func NewSequencer(addr common.Address, pubKey ecdsa.PublicKey, votingPower int64) *Sequencer {
	return &Sequencer{
		Address:          addr,
		PubKey:           pubKey,
		VotingPower:      votingPower,
		ProducerPriority: 0,
	}
}

// SequencerBasic performs basic validation.
func (s *Sequencer) SequencerBasic() error {
	if s == nil {
		return errors.New("nil sequencer")
	}
	if s.PubKey.Equal(nil) {
		return errors.New("sequencer does not have a public key")
	}

	if s.VotingPower < 0 {
		return errors.New("sequencer has negative voting power")
	}

	if len(s.Address) != common.AddressLength {
		return fmt.Errorf("sequencer address is the wrong size: %v", s.Address)
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
	return fmt.Sprintf("Sequencer{%v %v VP:%v A:%v}",
		s.Address,
		s.PubKey,
		s.VotingPower,
		s.ProducerPriority)
}

// SequencerListString returns a prettified sequencer list for logging purposes.
func SequencerListString(seqs []*Sequencer) string {
	chunks := make([]string, len(seqs))
	for i, seq := range seqs {
		chunks[i] = fmt.Sprintf("%s:%d", seq.Address, seq.VotingPower)
	}

	return strings.Join(chunks, ",")
}
