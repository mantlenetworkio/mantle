package sequencer

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/specularl2/specular/clients/geth/specular/rollup/services"
)

type Batch struct {
	Blocks            types.Blocks
	CumulativeGasUsed *big.Int
}

// Batcher assumes exclusive control of underlying blockchain, i.e.
// mining and blockchain insertion can not happen.
// TODO: support Berlin+London fork
type Batcher struct {
	eth         services.Backend
	chain       *core.BlockChain
	chainConfig *params.ChainConfig

	state    *state.StateDB // apply state changes here
	coinbase common.Address

	blocks []*types.Block // batched blocks

	// current batching block
	header   *types.Header
	gasPool  *core.GasPool // available gas used to pack transactions
	tcount   int
	txs      []*types.Transaction
	receipts []*types.Receipt
}

func NewBatcher(coinbase common.Address, eth services.Backend) (*Batcher, error) {
	b := &Batcher{
		coinbase: coinbase,
		eth:      eth,
	}
	b.chain = eth.BlockChain()
	b.chainConfig = b.chain.Config()
	err := b.startNewBlock()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// CommitTransactions will try fill transactions into blocks, and insert
// full blocks into the blockchain
// TODO: recover from failed commitBlock, rewind blockchain
func (b *Batcher) CommitTransactions(txs []*types.Transaction) error {
	// Update timestamp if we do not have any transactions in current block
	if len(b.txs) == 0 {
		b.header.Time = uint64(time.Now().Unix())
	}

	for i := 0; i < len(txs); i++ {
		// If we don't have enough gas for any further transactions, commit and
		// start a new block
		if b.gasPool.Gas() < params.TxGas {
			err := b.commitBlock()
			if err != nil {
				return err
			}
		}
		// Retrieve the next transaction
		tx := txs[i]
		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !b.chainConfig.IsEIP155(b.header.Number) {
			log.Trace("Ignoring reply protected transaction", "hash", tx.Hash(), "eip155", b.chainConfig.EIP155Block)
			continue
		}
		// Start executing the transaction
		b.state.Prepare(tx.Hash(), b.tcount)
		snap := b.state.Snapshot()
		receipt, err := core.ApplyTransaction(b.chainConfig, b.chain, &b.coinbase, b.gasPool, b.state, b.header, tx, &b.header.GasUsed, *b.chain.GetVMConfig())
		if err != nil {
			b.state.RevertToSnapshot(snap)
		}
		switch {
		case errors.Is(err, core.ErrGasLimitReached):
			// Commit block and retry transaction in new block
			err = b.commitBlock()
			if err != nil {
				return err
			}
			i--

		case errors.Is(err, nil):
			// Everything ok, collect the tx and receipt
			b.txs = append(b.txs, tx)
			b.receipts = append(b.receipts, receipt)
			b.tcount++
		}
	}

	return nil
}

// Batch will force the remaining transactions to form a block, insert it into
// the blockchain, and return all blocks created between Batch calls
func (b *Batcher) Batch() (types.Blocks, error) {
	err := b.commitBlock()
	if err != nil {
		return nil, err
	}
	blocks := b.blocks
	b.blocks = nil
	return blocks, nil
}

// commitBlock will assemble the current block, insert it into the blockchain
// and start a new block
func (b *Batcher) commitBlock() error {
	// TODO: return if nothing to be committed

	// Stop state prefetcher (see env.discard in worker.go)
	if b.state != nil {
		b.state.StopPrefetcher()
	}
	// Finalize header
	b.header.Root = b.state.IntermediateRoot(b.chain.Config().IsEIP158(b.header.Number))
	b.header.UncleHash = types.CalcUncleHash(nil)
	// Assemble block
	block := types.NewBlock(b.header, b.txs, nil, b.receipts, trie.NewStackTrie(nil))
	hash := block.Hash()
	// Finalize receipts and logs
	var logs []*types.Log
	for i, receipt := range b.receipts {
		// Add block location fields
		receipt.BlockHash = hash
		receipt.BlockNumber = block.Number()
		receipt.TransactionIndex = uint(i)

		// Update the block hash in all logs since it is now available and not when the
		// receipt/log of individual transactions were created.
		for _, log := range receipt.Logs {
			log.BlockHash = hash
		}
		logs = append(logs, receipt.Logs...)
	}
	// Write block to chain
	// Do not emit headEvent, it will cause worker to try sealing even if it is
	// not started
	_, err := b.chain.WriteBlockAndSetHead(block, b.receipts, logs, b.state, false)
	if err != nil {
		return err
	}
	err = b.startNewBlock()
	if err != nil {
		return err
	}
	b.blocks = append(b.blocks, block)
	return nil
}

// startNewBlock should be called when a block is full and inserted into the
// blockchain. It will reset the batcher except batched blocks
func (b *Batcher) startNewBlock() error {
	parent := b.chain.CurrentBlock()
	if parent == nil {
		return fmt.Errorf("missing parent")
	}
	num := parent.Number()
	b.header = &types.Header{
		ParentHash: parent.Hash(),
		Number:     num.Add(num, common.Big1),
		GasLimit:   core.CalcGasLimit(parent.GasLimit(), ethconfig.Defaults.Miner.GasCeil),
		Time:       uint64(time.Now().Unix()),
		Coinbase:   b.coinbase,
		Difficulty: common.Big1, // Fake difficulty. Avoid use 0 here because it means the merge happened
	}
	state, err := b.chain.StateAt(parent.Root())
	if err != nil {
		// Note since the sealing block can be created upon the arbitrary parent
		// block, but the state of parent block may already be pruned, so the necessary
		// state recovery is needed here in the future.
		//
		// The maximum acceptable reorg depth can be limited by the finalised block
		// somehow. TODO(rjl493456442) fix the hard-coded number here later.
		state, err = b.eth.StateAtBlock(parent, 1024, nil, false, false)
		log.Warn("Recovered mining state", "root", parent.Root(), "err", err)
	}
	if err != nil {
		return err
	}
	state.StartPrefetcher("sequencer")
	b.state = state
	b.gasPool = new(core.GasPool).AddGas(b.header.GasLimit)
	b.tcount = 0
	b.txs = nil
	b.receipts = nil
	return nil
}
