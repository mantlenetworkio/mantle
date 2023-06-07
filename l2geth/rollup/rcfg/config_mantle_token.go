package rcfg

import (
	"math/big"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/rollup/dump"
)

var (
	// predeploy
	L2MantleTokenAddress = dump.BvmMantleAddress

	// L2WrappedMantleTokenAddress deployed by WMANTLEDeployer, the Wrapped Mantle Token Address is deterministic
	// 0x9523886f149E12D04A9dd1E97214B00A039df199
	L2WrappedMantleTokenAddress = common.HexToAddress("0x9523886f149E12D04A9dd1E97214B00A039df199")

	MantleTokenNameSlot          = common.BigToHash(big.NewInt(3))
	MantleTokenSymbolSlot        = common.BigToHash(big.NewInt(4))
	WrappedMantleTokenNameSlot   = common.BigToHash(big.NewInt(3))
	WrappedMantleTokenSymbolSlot = common.BigToHash(big.NewInt(4))

	MantleTokenName               = "Mantle Token"
	MantleTokenSymbol             = "MNT"
	WrappedMantleTokenName        = "Wrapped MANTLE"
	WrappedMantleTokenSymbol      = "WMANTLE"
	MantleTokenNameValue          = common.BytesToHash(packStringForSlot(MantleTokenName))
	MantleTokenSymbolValue        = common.BytesToHash(packStringForSlot(MantleTokenSymbol))
	WrappedMantleTokenNameValue   = common.BytesToHash(packStringForSlot(WrappedMantleTokenName))
	WrappedMantleTokenSymbolValue = common.BytesToHash(packStringForSlot(WrappedMantleTokenSymbol))
)

// packStringForSlot pack a string value to store in contract directly
// just test when len(string) < 32
func packStringForSlot(value string) []byte {
	valueBytes := common.RightPadBytes([]byte(value), 32)
	valueBytes[31] = byte(len(value) * 2)
	return valueBytes
}
