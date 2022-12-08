package rcfg

import (
	"math/big"
	"os"

	"github.com/mantlenetworkio/mantle/l2geth/common"
)

// UsingBVM is used to enable or disable functionality necessary for the BVM.
var UsingBVM bool

var (
	// L2GasPriceSlot refers to the storage slot that the L2 gas price is stored
	// in in the BVM_GasPriceOracle predeploy
	L2GasPriceSlot = common.BigToHash(big.NewInt(1))
	// L1GasPriceSlot refers to the storage slot that the L1 gas price is stored
	// in in the BVM_GasPriceOracle predeploy
	L1GasPriceSlot = common.BigToHash(big.NewInt(2))
	// L2GasPriceOracleOwnerSlot refers to the storage slot that the owner of
	// the BVM_GasPriceOracle is stored in
	L2GasPriceOracleOwnerSlot = common.BigToHash(big.NewInt(0))
	// L2GasPriceOracleAddress is the address of the BVM_GasPriceOracle
	// predeploy
	L2GasPriceOracleAddress = common.HexToAddress("0x420000000000000000000000000000000000000F")
	// OverheadSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the per transaction overhead. This is added to the L1 cost portion
	// of the fee
	OverheadSlot = common.BigToHash(big.NewInt(3))
	// ScalarSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the transaction fee scalar. This value is scaled upwards by
	// the number of decimals
	ScalarSlot = common.BigToHash(big.NewInt(4))
	// DecimalsSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the number of decimals in the fee scalar
	DecimalsSlot = common.BigToHash(big.NewInt(5))
	// IsBurningSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds switch controls whether GasFee is brun
	IsBurningSlot = common.BigToHash(big.NewInt(6))
	// ChargeSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds switch controls whether charge
	ChargeSlot = common.BigToHash(big.NewInt(7))
)

func init() {
	UsingBVM = os.Getenv("USING_BVM") == "true"
}
