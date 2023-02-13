package common

import (
	"github.com/pkg/errors"
	"math/big"
	"time"
)

var (
	DefaultTimeout                  = 15 * time.Second
	ErrMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type RollupStoreResponse struct {
	DataStoreId uint32 `json:"data_store_id"`
	ConfirmAt   uint32 `json:"confirm_at"`
	Status      uint8  `json:"status"`
}
