package upgrade

import "math/big"

type Config struct {
	// GasPriceOracleL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	GasPriceOracleL2Block *big.Int

	// TssRewardL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	TssRewardL2Block *big.Int

	// MantleTokenL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	MantleTokenL2Block *big.Int

	// UpdateGasLimitL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	UpdateGasLimitL2Block *big.Int

	// EigenDaL2Block
	// Once the Layer 2 block height reaches this threshold, the upgrade operation will be executed
	EigenDaL2Block *big.Int

	// MockUpgradeL1Block
	// Once the Layer 1 block height reaches this threshold, the upgrade operation will be executed
	MockUpgradeL1Block *big.Int
	// MantleTokenL2BlockUpgradeStatus record upgrade status for MockUpgradeL1Block
	// Record the upgrade status only for upgrades based on L1 block number, as L2 blocks corresponding
	// to L1 blocks may not have transactions.
	// By checking if the L1 block number is greater than or equal to a certain value, we can ensure
	// that the upgrade operation will be triggered.
	// Also, keep a record of the upgrade to prevent it from being executed repeatedly for subsequent L2 heights.
	MockUpgradeL1BlockUpgradeStatus bool
}

var (
	MainnetConfig = &Config{
		GasPriceOracleL2Block: big.NewInt(-1),
		TssRewardL2Block:      big.NewInt(-1),
		MantleTokenL2Block:    big.NewInt(-1),
		EigenDaL2Block:        big.NewInt(-1),
		UpdateGasLimitL2Block: big.NewInt(-1),
		MockUpgradeL1Block:    big.NewInt(-1),
	}

	TestnetConfig = &Config{
		GasPriceOracleL2Block: big.NewInt(0),
		TssRewardL2Block:      big.NewInt(11_000_000),
		MantleTokenL2Block:    big.NewInt(11_000_000),
		EigenDaL2Block:        big.NewInt(0),
		UpdateGasLimitL2Block: big.NewInt(0),
		MockUpgradeL1Block:    big.NewInt(-1),
	}

	QAConfig = &Config{
		GasPriceOracleL2Block: big.NewInt(-1),
		TssRewardL2Block:      big.NewInt(-1),
		MantleTokenL2Block:    big.NewInt(-1),
		EigenDaL2Block:        big.NewInt(-1),
		UpdateGasLimitL2Block: big.NewInt(-1),
		MockUpgradeL1Block:    big.NewInt(-1),
	}
	LocalConfig = &Config{
		GasPriceOracleL2Block: big.NewInt(-1),
		TssRewardL2Block:      big.NewInt(-1),
		MantleTokenL2Block:    big.NewInt(-1),
		EigenDaL2Block:        big.NewInt(-1),
		UpdateGasLimitL2Block: big.NewInt(-1),
		MockUpgradeL1Block:    big.NewInt(-1),
	}
)

func init() {

}

// IsGasPriceOracleV1 returns whether num is either equal to the GasPriceOracleV1 fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsGasPriceOracleV1(num *big.Int) bool {
	return isExactBlockForked(c.GasPriceOracleL2Block, num)
}

// IsTssReward returns whether num is either equal to the TssReward fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsTssReward(num *big.Int) bool {
	return isExactBlockForked(c.TssRewardL2Block, num)
}

// IsMantleToken returns whether num is either equal to the IsMantleToken fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsMantleToken(num *big.Int) bool {
	return isExactBlockForked(c.MantleTokenL2Block, num)
}

// IsEigenDa returns whether num is either equal to the IsEigenDa fork block or greater.
// Compare with L2 BlockNumber
func (c *Config) IsEigenDa(num *big.Int) bool {
	return isExactBlockForked(c.EigenDaL2Block, num)
}

// IsMockUpgradeBasedOnL1BlockNumber returns whether num is either equal to the IsMantleToken fork block or greater.
// Compare with L1 BlockNumber
func (c *Config) IsMockUpgradeBasedOnL1BlockNumber(num *big.Int) bool {
	return isBlockForked(c.MantleTokenL2Block, num)
}

// isBlockForked returns whether a fork scheduled at block s is active at the
// given head block.
// isBlockForked is used to compare with Layer1 block height
func isBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
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
	return s.Cmp(head) == 0
}
