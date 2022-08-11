package types

import "math/big"

type SignStateRequest struct {
	StartBlock          big.Int    `json:"start_block"`
	OffsetStartsAtIndex big.Int    `json:"offset_starts_at_index"`
	StateRoots          [][32]byte `json:"state_roots"`
}
