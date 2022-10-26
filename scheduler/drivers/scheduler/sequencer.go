package scheduler

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
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
	Address     common.Address `json:"address"`
	NodeID      []byte         `json:"node_id"`
	VotingPower int64          `json:"voting_power"`

	ProducerPriority int64 `json:"producer_priority"`
}

// NewSequencer returns a new sequencer with the given pubkey and voting power.
func NewSequencer(addr common.Address, nodeId []byte, votingPower int64) *Sequencer {
	return &Sequencer{
		Address:          addr,
		NodeID:           nodeId,
		VotingPower:      votingPower,
		ProducerPriority: 0,
	}
}

// SequencerBasic performs basic validation.
func (s *Sequencer) SequencerBasic() error {
	if s == nil {
		return errors.New("nil sequencer")
	}
	if s.NodeID == nil || len(s.NodeID) != 32 {
		return errors.New("sequencer does not have a node id")
	}

	if s.VotingPower < 0 {
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
	return fmt.Sprintf("Sequencer{%v %v VP:%v A:%v}",
		s.Address,
		s.NodeID[:],
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

//----------------------------------------
// RandSequencer

// RandSequencer returns a randomized validator, useful for testing.
// UNSTABLE
func RandSequencer(randPower bool, minPower int64) *Sequencer {
	votePower := minPower
	if randPower {
		votePower += int64(rand.Uint32())
	}
	var seed []byte
	rand.Read(seed)
	var addr common.Address
	copy(addr[:], seed)
	priKey, err := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	if err != nil {
		panic(fmt.Errorf("could not retrieve priKey %w", err))
	}
	seq := NewSequencer(addr, priKey.D.Bytes(), votePower)
	return seq
}
