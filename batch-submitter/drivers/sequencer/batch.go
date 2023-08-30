package sequencer

import (
	"errors"
	"fmt"

	l2types "github.com/mantlenetworkio/mantle/l2geth/core/types"
)

var (
	// ErrBlockWithInvalidContext signals an attempt to generate a
	// BatchContext that specifies a total of zero txs.
	ErrBlockWithInvalidContext = errors.New("attempted to generate batch " +
		"context with 0 queued and 0 sequenced txs")
)

// BatchElement reflects the contents of an atomic update to the L2 state.
// Currently, each BatchElement is constructed from a single block containing
// exactly one tx.
type BatchElement struct {
	// Timestamp is the L1 timestamp of the batch.
	Timestamp uint64

	// BlockNumber is the L1 BlockNumber of the batch.
	BlockNumber uint64

	// Tx is the optional transaction that was applied in this batch.
	//
	// NOTE: This field will only be populated for sequencer txs.
	Tx *CachedTx
}

// IsSequencerTx returns true if this batch contains a tx that needs to be
// posted to the L1 CTC contract.
func (b *BatchElement) IsSequencerTx() bool {
	return b.Tx != nil
}

// BatchElementFromBlock constructs a BatchElement from a single L2 block. This
// method expects that there is exactly ONE tx per block. The returned
// BatchElement will reflect whether or not the lone tx is a sequencer tx or a
// queued tx.
func BatchElementFromBlock(block *l2types.Block) BatchElement {
	txs := block.Transactions()
	if len(txs) != 1 {
		panic(fmt.Sprintf("attempting to create batch element from block %d, "+
			"found %d txs instead of 1", block.Number(), len(txs)))
	}

	tx := txs[0]

	// Extract L2 metadata.
	l1BlockNumber := tx.L1BlockNumber().Uint64()
	isSequencerTx := tx.QueueOrigin() == l2types.QueueOriginSequencer

	// Only include sequencer txs in the returned BatchElement.
	var cachedTx *CachedTx
	if isSequencerTx {
		cachedTx = NewCachedTx(tx)
	}

	return BatchElement{
		Timestamp:   block.Time(),
		BlockNumber: l1BlockNumber,
		Tx:          cachedTx,
	}
}

type groupedBlock struct {
	sequenced []BatchElement
	queued    []BatchElement
}

// GenSequencerBatchParams generates a valid AppendSequencerBatchParams from a
// list of BatchElements. The BatchElements are assumed to be ordered in
// ascending order by L2 block height.
func GenSequencerBatchParams(
	shouldStartAtElement uint64,
	blockOffset uint64,
	batchNumber uint64,
	timestamp uint64,
	blockNumber uint64,

	numSequencedTxs uint64,
	numSubsequentQueueTxs uint64,

) (*AppendSequencerBatchParams, error) {
	var (
		contexts []BatchContext
	)
	contexts = append(contexts, BatchContext{
		NumSequencedTxs:       numSequencedTxs,
		NumSubsequentQueueTxs: numSubsequentQueueTxs,
		Timestamp:             timestamp,
		BlockNumber:           blockNumber,
	})

	return &AppendSequencerBatchParams{
		ShouldStartAtElement:  shouldStartAtElement - blockOffset,
		TotalElementsToAppend: batchNumber,
		Contexts:              contexts,
	}, nil
}
