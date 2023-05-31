package rcfg

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"math/big"
)

var (
	// predeploy
	L2MantleTokenAddress        = common.HexToAddress("0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000")
	L2WrappedMantleTokenAddress = common.HexToAddress("0x9523886f149E12D04A9dd1E97214B00A039df199")

	MantleTokenNameSlot          = common.BigToHash(big.NewInt(3))
	MantleTokenSymbolSlot        = common.BigToHash(big.NewInt(4))
	WrappedMantleTokenNameSlot   = common.BigToHash(big.NewInt(3))
	WrappedMantleTokenSymbolSlot = common.BigToHash(big.NewInt(4))

	MantleTokenName               = "Mantle Tokens"
	MantleTokenSymbol             = "MNTs"
	WrappedMantleTokenName        = "Wrapped MANTLEs"
	WrappenMantleTokenSymbol      = "WMANTLEs"
	MantleTokenNameValue          = common.BytesToHash(packBytesSliceWithLen([]byte(MantleTokenName), len(MantleTokenName)))
	MantleTokenSymbolValue        = common.BytesToHash(packBytesSliceWithLen([]byte(MantleTokenSymbol), len(MantleTokenSymbol)))
	WrappedMantleTokenNameValue   = common.BytesToHash(packBytesSliceWithLen([]byte(WrappedMantleTokenName), len(WrappedMantleTokenName)))
	WrappedMantleTokenSymbolValue = common.BytesToHash(packBytesSliceWithLen([]byte(WrappenMantleTokenSymbol), len(WrappenMantleTokenSymbol)))
)

// packBytesSlice packs the given bytes as [L, V] as the canonical representation
// bytes slice
// string	->	slot
// MNT 		-> 	[77 78 84 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 6]
func packBytesSliceWithLen(bytes []byte, len int) []byte {
	val := common.RightPadBytes(bytes, 32)
	val[31] = byte(len * 2)
	return val
}
