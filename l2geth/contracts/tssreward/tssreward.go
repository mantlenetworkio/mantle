package tssreward

import (
	_ "embed"
	"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi"
	"math/big"
	"strings"
) // embed compiled smart contract

var (
	tssRewardABI *abi.ABI
)

const jsondata = `
[
	{ "type" : "function", "name" : "balance", "constant" : true },
	{ "type" : "function", "name" : "send", "constant" : false, "inputs" : [ { "name" : "amount", "type" : "uint256" } ] }
]`

func Abi() *abi.ABI {
	if tssRewardABI != nil {
		return tssRewardABI
	}
	var err error
	*tssRewardABI, err = abi.JSON(strings.NewReader(jsondata))
	if err != nil {
		panic(err)
	}
	return tssRewardABI
}

func UpdateTssRewardData(blockID *big.Int, amount *big.Int) ([]byte, error) {
	tssRewardABI = Abi()
	data, err := tssRewardABI.Pack("updateTssReward", blockID, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}
