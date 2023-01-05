package sequencer

import (
	"math/big"
)

type BatchTx struct {
	BlockNumber *big.Int
	rawTx       []byte
}

type BatchTxVector struct {
	txData []BatchTx
}

type SliceBatchTx struct {
	addr uintptr
	len  int
	cap  int
}
