package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

// TxBatch represents a transaction batch to be sequenced to L1 sequencer inbox
// It may contain multiple blocks
type TxBatch struct {
	Blocks   types.Blocks
	Contexts []SequenceContext
	Txs      types.Transactions
	GasUsed  *big.Int
}

// SequenceContext is the relavent context of each block sequenced to L1 sequncer inbox
type SequenceContext struct {
	NumTxs      uint64
	BlockNumber uint64
	Timestamp   uint64
}

func NewTxBatch(blocks []*types.Block, maxBatchSize uint64) *TxBatch {
	// TODO: handle maxBatchSize constraint
	var contexts []SequenceContext
	var txs []*types.Transaction
	gasUsed := new(big.Int)
	for _, block := range blocks {
		blockTxs := block.Transactions()
		ctx := SequenceContext{
			NumTxs:      uint64(len(blockTxs)),
			BlockNumber: block.Number().Uint64(), // TODO just use bigint
			Timestamp:   block.Time(),
		}
		contexts = append(contexts, ctx)
		txs = append(txs, blockTxs...)
		gasUsed.Add(gasUsed, new(big.Int).SetUint64(block.GasUsed()))
	}
	return &TxBatch{blocks, contexts, txs, gasUsed}
}

func (b *TxBatch) LastBlockNumber() uint64 {
	return b.Contexts[len(b.Contexts)-1].BlockNumber
}

func (b *TxBatch) LastBlockRoot() common.Hash {
	return b.Blocks[len(b.Blocks)-1].Root()
}

func (b *TxBatch) InboxSize() *big.Int {
	return new(big.Int).SetUint64(uint64(b.Txs.Len()))
}

func (b *TxBatch) SerializeToArgs() ([]*big.Int, []*big.Int, []byte, error) {
	var contexts, txLengths []*big.Int
	for _, ctx := range b.Contexts {
		contexts = append(contexts, big.NewInt(0).SetUint64(ctx.NumTxs))
		contexts = append(contexts, big.NewInt(0).SetUint64(ctx.BlockNumber))
		contexts = append(contexts, big.NewInt(0).SetUint64(ctx.Timestamp))
	}

	buf := new(bytes.Buffer)
	for _, tx := range b.Txs {
		curLen := buf.Len()
		if err := writeTx(buf, tx); err != nil {
			return nil, nil, nil, err
		}
		txLengths = append(txLengths, big.NewInt(int64(buf.Len()-curLen)))
	}
	return contexts, txLengths, buf.Bytes(), nil
}

func (b *TxBatch) ToAssertion(parent *Assertion) *Assertion {
	return &Assertion{
		ID:        new(big.Int).Add(parent.ID, big.NewInt(1)),
		VmHash:    b.LastBlockRoot(),
		InboxSize: big.NewInt(int64(b.LastBlockNumber())), // TODO-FIXME uint64 -> int64 -> big.Int
		GasUsed:   new(big.Int).Add(b.GasUsed, parent.GasUsed),
		Parent:    parent.ID,
	}
}

func writeContext(w *bytes.Buffer, ctx *SequenceContext) error {
	if err := writePrimitive(w, ctx.NumTxs); err != nil {
		return err
	}
	if err := writePrimitive(w, ctx.BlockNumber); err != nil {
		return err
	}
	return writePrimitive(w, ctx.Timestamp)
}

func writeTx(w *bytes.Buffer, tx *types.Transaction) error {
	var txBuf bytes.Buffer
	if err := tx.EncodeRLP(&txBuf); err != nil {
		return err
	}
	txBytes := txBuf.Bytes()
	_, err := w.Write(txBytes)
	return err
}

func writePrimitive(w *bytes.Buffer, data interface{}) error {
	return binary.Write(w, binary.BigEndian, data)
}

var TxBatchParseFailedError = errors.New("Failed to create TxBatch from decoded tx data")

// TxBatchFromDecoded decodes the input of SequencerInbox#appendTxBatch call
// It will only fill Contexts and Txs fields
func TxBatchFromDecoded(decoded []interface{}) (*TxBatch, error) {
	if len(decoded) != 3 {
		return nil, TxBatchParseFailedError
	}
	contexts := decoded[0].([]*big.Int)
	txLengths := decoded[1].([]*big.Int)
	txBatch := decoded[2].([]byte)
	blockNum := len(txLengths)
	if len(contexts) != 3*blockNum {
		return nil, TxBatchParseFailedError
	}
	var txs []*types.Transaction
	var ctxs []SequenceContext
	var txLen uint64 = 0
	for idx, l := range txLengths {
		length := l.Uint64()
		if uint64(len(txBatch)) < txLen+length {
			return nil, TxBatchParseFailedError
		}
		ctx := SequenceContext{
			NumTxs:      contexts[3*idx].Uint64(),
			BlockNumber: contexts[3*idx+1].Uint64(),
			Timestamp:   contexts[3*idx+2].Uint64(),
		}
		raw := txBatch[txLen : txLen+length]
		var tx types.Transaction
		err := rlp.DecodeBytes(raw, &tx)
		if err != nil {
			return nil, err
		}
		ctxs = append(ctxs, ctx)
		txs = append(txs, &tx)
	}
	batch := &TxBatch{
		Contexts: ctxs,
		Txs:      txs,
	}
	return batch, nil
}

// SequenceBlock is a block sequenced to L1 sequencer inbox
// It is used by validators to reconstruct L2 chain from L1 activities
type SequenceBlock struct {
	SequenceContext
	Txs types.Transactions
}
