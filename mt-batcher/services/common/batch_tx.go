package common

import (
	"math/big"
)

type BatchTx struct {
	BlockNumber *big.Int
	RawTx       []byte
}
