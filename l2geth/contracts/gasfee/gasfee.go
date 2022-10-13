package gasfee

import (
	_ "embed"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/abi"
	"math/big"
	"strings"
)

const jsondata = `[{"inputs":[{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"burn","stateMutability":"nonpayable","type":"function"}]`

var (
	sequencerFeeVault *abi.ABI
)

func Abi() *abi.ABI {
	if sequencerFeeVault != nil {
		return sequencerFeeVault
	}
	sequencerFeeVaultABI, err := abi.JSON(strings.NewReader(jsondata))
	if err != nil {
		panic(err)
	}
	sequencerFeeVault = &sequencerFeeVaultABI
	return sequencerFeeVault
}
func PacketData(amount *big.Int) ([]byte, error) {
	sequencerFeeVault = Abi()
	data, err := sequencerFeeVault.Pack("burn", amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}
