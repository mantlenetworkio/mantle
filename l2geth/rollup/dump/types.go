package dump

import (
	"github.com/mantlenetworkio/mantle/l2geth/accounts/abi"
	"github.com/mantlenetworkio/mantle/l2geth/common"
)

type BvmDumpAccount struct {
	Address  common.Address         `json:"address"`
	Code     string                 `json:"code"`
	CodeHash string                 `json:"codeHash"`
	Storage  map[common.Hash]string `json:"storage"`
	ABI      abi.ABI                `json:"abi"`
	Nonce    uint64                 `json:"nonce"`
}

type BvmDump struct {
	Accounts map[string]BvmDumpAccount `json:"accounts"`
}
