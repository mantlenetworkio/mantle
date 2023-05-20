package types

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"math/big"
)

// Receipt represents the results of a transaction.
type TxStatus struct {
	Status uint `json:"status"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint64      `json:"transactionIndex"`
}

type TxStatusDetail struct {
	Status uint `json:"status"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint64      `json:"transactionIndex"`
}

// StateRootResponse represents the response from the remote server when querying stateroot
type TxStatusResponse struct {
	StateRoot       *StateRoot `json:"stateRoots"`
	Batch           *Batch     `json:"batch"`
	CurrentL1Height int64      `json:"currentL1BlockNumber"`
}

// stateroot represents the return result of the remote server.
// it came from a batch or was replicated from the sequencer.
type StateRoot struct {
	Index      uint64 `json:"index"`
	BatchIndex uint64 `json:"batchIndex"`
	Value      string `json:"value"`
}

// Batch represents the data structure that is submitted with
// a series of transactions to layer one
type Batch struct {
	Index             uint64         `json:"index"`
	Root              common.Hash    `json:"root,omitempty"`
	Size              uint32         `json:"size,omitempty"`
	PrevTotalElements uint32         `json:"prevTotalElements,omitempty"`
	ExtraData         hexutil.Bytes  `json:"extraData,omitempty"`
	BlockNumber       uint64         `json:"blockNumber"`
	Timestamp         uint64         `json:"timestamp"`
	Submitter         common.Address `json:"submitter"`
}
