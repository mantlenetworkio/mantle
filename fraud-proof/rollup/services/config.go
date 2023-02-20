package services

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
)

const (
	NODE_SCHEDULER = "scheduler"
	NODE_VERIFIER  = "verifier"
)

// Config is the configuration of rollup services
type Config struct {
	Node            string         // Rollup node type, either sequencer or validator
	Passphrase      string         // The passphrase of the coinbase account
	L1Endpoint      string         // L1 API endpoint
	L1ChainID       uint64         // L1 chain ID
	L1Confirmations uint64         // L1 confirmation block number
	SequencerAddr   common.Address // Validator only
	RollupAddr      common.Address // L1 Rollup contract address
	StakeAddr       common.Address // The account used for rollup assertion stake
	StakeAmount     uint64         // Amount of stake
}
