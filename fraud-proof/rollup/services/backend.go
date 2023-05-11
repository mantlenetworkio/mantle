package services

import (
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/state"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
)

// Required interface for interacting with Ethereum instance
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	StateAtBlock(block *types.Block, reexec uint64, base *state.StateDB, checkLive bool, preferDisk bool) (statedb *state.StateDB, err error)
}
