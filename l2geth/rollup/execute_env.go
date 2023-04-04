package rollup

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus"
	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/state"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/params"
	"math/big"
	"sync"
	"time"
)

type environment struct {
	signer types.Signer

	state     *state.StateDB // apply state changes here
	ancestors mapset.Set     // ancestor set (used for checking uncle parent validity)
	family    mapset.Set     // family set (used for checking uncle invalidity)
	uncles    mapset.Set     // uncle set
	tcount    int            // tx count in cycle
	gasPool   *core.GasPool  // available gas used to pack transactions

	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt
}

type executor struct {
	mu          sync.RWMutex
	chain       *core.BlockChain
	chainConfig *params.ChainConfig
	extra       []byte
	GasFloor    uint64
	engine      consensus.Engine
	current     *environment
	coinbase    common.Address
}

func newExecutor(gasFloor uint64, chainConfig *params.ChainConfig, engine consensus.Engine, chain *core.BlockChain) *executor {
	return &executor{
		chain:       chain,
		GasFloor:    gasFloor,
		chainConfig: chainConfig,
		engine:      engine,
	}
}

// commitNewTx is an BVM addition that mines a block with a single tx in it.
// It needs to return an error in the case there is an error to prevent waiting
// on reading from a channel that is written to when a new block is added to the
// chain.
func (e *executor) applyTx(tx *types.Transaction) (error, *types.Block) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	tstart := time.Now()

	parent := e.chain.CurrentBlock()
	num := parent.Number()

	// Preserve liveliness as best as possible. Must panic on L1 to L2
	// transactions as the timestamp cannot be malleated
	if parent.Time() > tx.L1Timestamp() {
		log.Error("Monotonicity violation", "index", num, "parent", parent.Time(), "tx", tx.L1Timestamp())
	}

	// Fill in the index field in the tx meta if it is `nil`.
	// This should only ever happen in the case of the sequencer
	// receiving a queue origin sequencer transaction. The verifier
	// should always receive transactions with an index as they
	// have already been confirmed in the canonical transaction chain.
	// Use the parent's block number because the CTC is 0 indexed.
	if meta := tx.GetMeta(); meta.Index == nil {
		index := num.Uint64()
		meta.Index = &index
		tx.SetTransactionMeta(meta)
	}
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     new(big.Int).Add(num, common.Big1),
		GasLimit:   e.GasFloor,
		Extra:      e.extra,
		Time:       tx.L1Timestamp(),
	}
	if err := e.engine.Prepare(e.chain, header); err != nil {
		return fmt.Errorf("Failed to prepare header for mining: %w", err), nil
	}
	// Could potentially happen if starting to mine in an odd state.
	err := e.makeCurrent(parent, header)
	if err != nil {
		return fmt.Errorf("Failed to create mining context: %w", err), nil
	}
	transactions := make(map[common.Address]types.Transactions)
	acc, _ := types.Sender(e.current.signer, tx)
	transactions[acc] = types.Transactions{tx}
	txs := types.NewTransactionsByPriceAndNonce(e.current.signer, transactions)
	if err := e.commitTransactionsWithError(txs, e.coinbase); err != nil {
		return err, nil
	}

	return e.commit(tstart)
}

// makeCurrent creates a new environment for the current cycle.
func (e *executor) makeCurrent(parent *types.Block, header *types.Header) error {
	state, err := e.chain.StateAt(parent.Root())
	if err != nil {
		return err
	}
	env := &environment{
		signer:    types.NewEIP155Signer(e.chainConfig.ChainID),
		state:     state,
		ancestors: mapset.NewSet(),
		family:    mapset.NewSet(),
		uncles:    mapset.NewSet(),
		header:    header,
	}

	// when 08 is processed ancestors contain 07 (quick block)
	for _, ancestor := range e.chain.GetBlocksFromHash(parent.Hash(), 7) {
		for _, uncle := range ancestor.Uncles() {
			env.family.Add(uncle.Hash())
		}
		env.family.Add(ancestor.Hash())
		env.ancestors.Add(ancestor.Hash())
	}

	// Keep track of transactions which return errors so they can be removed
	env.tcount = 0
	e.current = env
	return nil
}

func (e *executor) commitTransactionsWithError(txs *types.TransactionsByPriceAndNonce, coinbase common.Address) error {

	if e.current.gasPool == nil {
		e.current.gasPool = new(core.GasPool).AddGas(e.current.header.GasLimit)
	}

	for {

		// If we don't have enough gas for any further transactions then we're done
		if e.current.gasPool.Gas() < params.TxGas {
			log.Trace("Not enough gas for further transactions", "have", e.current.gasPool, "want", params.TxGas)
			break
		}
		// Retrieve the next transaction and abort if all done
		tx := txs.Peek()
		if tx == nil {
			break
		}

		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		//
		// We use the eip155 signer regardless of the current hf.
		from, _ := types.Sender(e.current.signer, tx)
		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !e.chainConfig.IsEIP155(e.current.header.Number) {
			log.Trace("Ignoring reply protected transaction", "hash", tx.Hash(), "eip155", e.chainConfig.EIP155Block)

			txs.Pop()
			continue
		}
		// Start executing the transaction
		e.current.state.Prepare(tx.Hash(), common.Hash{}, e.current.tcount)

		_, err := e.commitTransaction(tx, coinbase)
		switch err {
		case core.ErrGasLimitReached:
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Trace("Gas limit exceeded for current block", "sender", from)
			txs.Pop()

		case core.ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
			txs.Shift()

		case core.ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Trace("Skipping account with high nonce", "sender", from, "nonce", tx.Nonce())
			txs.Pop()

		case nil:
			// Everything ok, collect the logs and shift in the next transaction from the same account
			e.current.tcount++
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}

		// UsingBVM
		// Return specific execution errors directly to the user to
		// avoid returning the generic ErrCannotCommitTxnErr. It is safe
		// to return the error directly since l2geth only processes at
		// most one transaction per block.
		if err != nil {
			return err
		}
	}

	if e.current.tcount == 0 {
		return errors.New("can not commit transaction in executor")
	}
	return nil
}

func (e *executor) commitTransaction(tx *types.Transaction, coinbase common.Address) ([]*types.Log, error) {
	// Make sure there's only one tx per block
	if e.current != nil && len(e.current.txs) > 0 {
		return nil, core.ErrGasLimitReached
	}
	snap := e.current.state.Snapshot()
	receipt, err := core.ApplyTransaction(e.chainConfig, e.chain, &coinbase, e.current.gasPool, e.current.state, e.current.header, tx, &e.current.header.GasUsed, *e.chain.GetVMConfig())
	if err != nil {
		e.current.state.RevertToSnapshot(snap)
		return nil, err
	}
	e.current.txs = append(e.current.txs, tx)
	e.current.receipts = append(e.current.receipts, receipt)
	return receipt.Logs, nil
}

// commit runs any post-transaction state modifications, assembles the final block
// and commits new work if consensus engine is running.
func (e *executor) commit(start time.Time) (error, *types.Block) {
	// Deep copy receipts here to avoid interaction between different tasks.
	receipts := make([]*types.Receipt, len(e.current.receipts))
	for i, l := range e.current.receipts {
		receipts[i] = new(types.Receipt)
		*receipts[i] = *l
	}
	s := e.current.state.Copy()
	block, err := e.engine.FinalizeAndAssemble(e.chain, e.current.header, s, e.current.txs, nil, e.current.receipts)
	if err != nil {
		return err, nil
	}

	// As a sanity check, ensure all new blocks have exactly one
	// transaction. This check is done here just in case any of our
	// higher-evel checks failed to catch empty blocks passed to commit.
	txs := block.Transactions()
	if len(txs) != 1 {
		return fmt.Errorf("Block created with %d transactions rather than 1 at %d", len(txs), block.NumberU64()), nil
	}
	tend := time.Now()
	log.Info("apply transaction ", "start ", start.String(), "end", tend.String(), "total cost time", tend.Sub(start).String())

	return nil, block
}

// setEtherbase sets the etherbase used to initialize the block coinbase field.
func (e *executor) setEtherbase(addr common.Address) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.coinbase = addr
}

// setExtra sets the content used to initialize the block extra field.
func (e *executor) setExtra(extra []byte) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.extra = extra
}
