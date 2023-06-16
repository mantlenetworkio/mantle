package upgrade

import (
	"math/big"

	"github.com/mantlenetworkio/mantle/l2geth/params"
)

type Config struct {
	chainID *big.Int

	// tssRewardL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	tssRewardL2Block *big.Int

	// mantleTokenL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	mantleTokenL2Block *big.Int

	// updateGasLimitL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	updateGasLimitL2Block *big.Int

	// eigenDaL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	eigenDaL2Block *big.Int

	// mockUpgradeL1Block
	// Once the Layer 1 block height reaches this threshold, the upgrade operation will be executed
	mockUpgradeL1Block *big.Int
	// mockUpgradeL1BlockUpgradeStatus record upgrade status for mockUpgradeL1Block
	// Record the upgrade status only for upgrades based on L1 block number, as L2 blocks corresponding
	// to L1 blocks may not have transactions.
	// By checking if the L1 block number is greater than or equal to a certain value, we can ensure
	// that the upgrade operation will be triggered.
	// Also, keep a record of the upgrade to prevent it from being executed repeatedly for subsequent L2 heights.
	mockUpgradeL1BlockUpgradeStatus bool
}

var (
	MainnetConfig = &Config{
		chainID:               params.MantleMainnetChainID,
		tssRewardL2Block:      big.NewInt(-1),
		mantleTokenL2Block:    big.NewInt(-1),
		eigenDaL2Block:        big.NewInt(-1),
		updateGasLimitL2Block: big.NewInt(-1),
		mockUpgradeL1Block:    big.NewInt(-1),
	}

	TestnetConfig = &Config{
		chainID:               params.MantleTestnetChainID,
		tssRewardL2Block:      big.NewInt(11_000_000),
		mantleTokenL2Block:    big.NewInt(11_000_000),
		eigenDaL2Block:        big.NewInt(8_280_000),
		updateGasLimitL2Block: big.NewInt(222_073),
		mockUpgradeL1Block:    big.NewInt(-1),
	}

	QAConfig = &Config{
		chainID:               params.MantleQAChainID,
		tssRewardL2Block:      big.NewInt(-1),
		mantleTokenL2Block:    big.NewInt(-1),
		eigenDaL2Block:        big.NewInt(-1),
		updateGasLimitL2Block: big.NewInt(-1),
		mockUpgradeL1Block:    big.NewInt(-1),
	}

	LocalConfig = &Config{
		chainID:               params.MantleLocalChainID,
		tssRewardL2Block:      big.NewInt(-1),
		mantleTokenL2Block:    big.NewInt(-1),
		eigenDaL2Block:        big.NewInt(-1),
		updateGasLimitL2Block: big.NewInt(-1),
		mockUpgradeL1Block:    big.NewInt(-1),
	}
)

func NewMantleUpgradeConfig(chainID *big.Int) *Config {
	switch chainID.Int64() {
	case params.MantleMainnetChainID.Int64():
		return MainnetConfig
	case params.MantleTestnetChainID.Int64():
		return TestnetConfig
	case params.MantleQAChainID.Int64():
		return QAConfig
	case params.MantleLocalChainID.Int64():
		return LocalConfig
	default:
		return LocalConfig
	}
}

// IsTssReward returns whether num is either equal to the TssReward fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsTssReward(num *big.Int) bool {
	return isExactBlockForked(c.tssRewardL2Block, num)
}

// IsMantleToken returns whether num is either equal to the IsMantleToken fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsMantleToken(num *big.Int) bool {
	return isExactBlockForked(c.mantleTokenL2Block, num)
}

// IsEigenDa returns whether num is either equal to the IsEigenDa fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsEigenDa(num *big.Int) bool {
	return isExactBlockForked(c.eigenDaL2Block, num)
}

// IsUpdateGasLimitBlock returns whether num is either equal to the IsUpdateGasLimitBlock fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsUpdateGasLimitBlock(num *big.Int) bool {
	return isBlockForked(c.updateGasLimitL2Block, num)
}

// IsMockUpgradeBasedOnL1BlockNumber returns whether num is either equal to the IsMantleToken fork block or greater.
// Compare with L1 BlockNumber
func (c *Config) IsMockUpgradeBasedOnL1BlockNumber(num *big.Int) bool {
	return isBlockForked(c.mockUpgradeL1Block, num)
}

// isBlockForked returns whether a fork scheduled at block s is active at the
// given head block.
// isBlockForked is used to compare with Layer1 block height
func isBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}

	if s.Cmp(big.NewInt(0)) <= 0 {
		return false
	}

	return s.Cmp(head) <= 0
}

// isExactBlockForked returns whether a fork scheduled at block s is active at the
// given head block.
// isExactBlockForked is used to compare with Layer2 block height
func isExactBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}

	if s.Cmp(big.NewInt(0)) <= 0 {
		return false
	}

	return s.Cmp(head) == 0
}
