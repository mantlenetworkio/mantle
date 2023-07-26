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
	weiToEth          = new(big.Float).SetFloat64(1e-18)
)

func WeiToEth64(wei *big.Int) float64 {
	eth := new(big.Float).SetInt(wei)
	eth.Mul(eth, weiToEth)
	eth64, _ := eth.Float64()
	return eth64
}
