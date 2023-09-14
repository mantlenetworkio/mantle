package sequencer

import (
	"fmt"

	l2types "github.com/mantlenetworkio/mantle/l2geth/core/types"
)

// BatchElement reflects the contents of an atomic update to the L2 state.
// Currently, each BatchElement is constructed from a single block containing
// exactly one tx.
type BatchElement struct {
	// Timestamp is the L1 timestamp of the batch.
	Timestamp uint64

	// BlockNumber is the L1 BlockNumber of the batch.
	BlockNumber uint64

	// whether it is a Sequencer transaction
	IsSequencerTx bool
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
	return BatchElement{
		Timestamp:     block.Time(),
		BlockNumber:   l1BlockNumber,
		IsSequencerTx: isSequencerTx,
	}
}
