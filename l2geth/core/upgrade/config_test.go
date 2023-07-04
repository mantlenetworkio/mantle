package upgrade

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mantlenetworkio/mantle/l2geth/params"
)

func TestConfig_IsTssReward(t *testing.T) {
	upgradeConfig := NewMantleUpgradeConfig(params.MantleMainnetChainID)
	result := upgradeConfig.IsTssReward(big.NewInt(11_000_000))
	if result {
		t.Error("tss reward could not used for mainnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsTssReward(big.NewInt(11_000_000))
	if !result {
		t.Error("tss reward should used for testnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsTssReward(big.NewInt(11_000_001))
	if result {
		t.Error("tss reward upgrade height has passed")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsTssReward(big.NewInt(10_999_999))
	if result {
		t.Error("tss reward upgrade height does not reach to")
	}
}

func TestConfig_IsMantleToken(t *testing.T) {
	upgradeConfig := NewMantleUpgradeConfig(params.MantleMainnetChainID)
	result := upgradeConfig.IsMantleToken(big.NewInt(11_000_000))
	if result {
		t.Error("mantle token upgrade could not used for mainnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsMantleToken(big.NewInt(11_000_000))
	if !result {
		t.Error("mantle token upgrade should used for testnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsMantleToken(big.NewInt(11_000_001))
	if result {
		t.Error("mantle token upgrade height has passed")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsMantleToken(big.NewInt(10_999_999))
	if result {
		t.Error("mantle token upgrade height does not reach to")
	}
}

func TestConfig_IsEigenDa(t *testing.T) {
	upgradeConfig := NewMantleUpgradeConfig(params.MantleMainnetChainID)
	result := upgradeConfig.IsEigenDa(big.NewInt(8_280_000))
	if result {
		t.Error("eigen data layer upgrade could not used for mainnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsEigenDa(big.NewInt(8_280_000))
	if !result {
		t.Error("eigen data layer upgrade should used for testnet")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsEigenDa(big.NewInt(8_280_001))
	if result {
		t.Error("eigen data layer upgrade height has passed")
	}

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsEigenDa(big.NewInt(8_279_999))
	if result {
		t.Error("eigen data layer upgrade height does not reach to")
	}
}

func TestConfig_IsUpdateGasLimitBlock(t *testing.T) {
	upgradeConfig := NewMantleUpgradeConfig(params.MantleMainnetChainID)
	result := upgradeConfig.IsUpdateGasLimitBlock(big.NewInt(222_073)) && (upgradeConfig.chainID == params.MantleTestnetChainID)
	require.Equal(t, result, false, "update gaslimit upgrade could not used for mainnet")

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsUpdateGasLimitBlock(big.NewInt(222_073))
	require.Equal(t, result, true, "update gaslimit upgrade should used for testnet")

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsUpdateGasLimitBlock(big.NewInt(222_074))
	require.Equal(t, result, true, "update gaslimit upgrade take effect after 222_073")

	upgradeConfig = NewMantleUpgradeConfig(params.MantleTestnetChainID)
	result = upgradeConfig.IsUpdateGasLimitBlock(big.NewInt(222_072))
	require.Equal(t, result, false, "update gaslimit upgrade height does not reach to")
}
