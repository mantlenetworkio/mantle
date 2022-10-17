package gasfee

import (
	_ "embed"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/abi"
	"strings"
)

const jsondata = `[{"inputs":[],"outputs":[],"name":"withdraw","stateMutability":"nonpayable","type":"function"}]`

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
func PacketData() ([]byte, error) {
	sequencerFeeVault = Abi()
	data, err := sequencerFeeVault.Pack("withdraw")
	if err != nil {
		return nil, err
	}
	return data, nil
}
