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

const jsondata = `[{"inputs":[{"internalType":"address","name":"_deadAddress","type":"address"},{"internalType":"address payable","name":"_owner","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"blockStartHeight","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"length","type":"uint256"},{"indexed":false,"internalType":"address[]","name":"tssMembers","type":"address[]"}],"name":"DistributeTssReward","type":"event"},{"inputs":[],"name":"bestBlockID","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_blockStartHeight","type":"uint256"},{"internalType":"uint32","name":"_length","type":"uint32"},{"internalType":"address[]","name":"_tssMembers","type":"address[]"}],"name":"claimReward","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"deadAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"dust","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"ledger","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address payable","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"queryReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_blockID","type":"uint256"},{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"updateReward","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"withdraw","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"withdrawDust","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

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
