package eigenda

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"math/big"
)

type BatchTx struct {
	BlockNumber []byte
	TxMeta      []byte
	RawTx       []byte
}

type RollupStoreResponse struct {
	DataStoreId uint32 `json:"data_store_id"`
	ConfirmAt   uint32 `json:"confirm_at"`
	Status      uint8  `json:"status"`
}

type TransactionMeta struct {
	L1BlockNumber   *big.Int        `json:"l1BlockNumber"`
	L1Timestamp     uint64          `json:"l1Timestamp"`
	L1MessageSender *common.Address `json:"l1MessageSender"`
	Index           *uint64         `json:"index"`
	QueueIndex      *uint64         `json:"queueIndex"`
	RawTransaction  []byte          `json:"rawTransaction"`
}
