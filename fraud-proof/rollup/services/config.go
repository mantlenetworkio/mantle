package services

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
)

const (
	NODE_SEQUENCER = "sequencer"
	NODE_VALIDATOR = "validator"
)

// Config is the configuration of rollup services
type Config struct {
	Node          string         // Rollup node type, either sequencer or validator
	Coinbase      common.Address // The account used for L1 and L2 activity
	Passphrase    string         // The passphrase of the coinbase account
	L1Endpoint    string         // L1 API endpoint
	L1ChainID     uint64         // L1 chain ID
	SequencerAddr common.Address // Validator only
	//SequencerInboxAddr common.Address // L1 SequencerInbox contract address
	RollupAddr        common.Address // L1 Rollup contract address
	RollupStakeAmount uint64         // Amount of stake
}
