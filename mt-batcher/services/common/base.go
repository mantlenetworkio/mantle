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

type BaseResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func BaseResource(status bool, code int, data interface{}, msg string) (baseRep *BaseResponse) {
	baseRep = &BaseResponse{Status: status, Code: code, Data: data, Msg: msg}
	return
}
