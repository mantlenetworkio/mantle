package migration_action

import (
	"context"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/mt-chain-ops/genesis"
	"github.com/mantlenetworkio/mantle/mt-chain-ops/genesis/migration"
)

type Config struct {
	DeployConfig          *genesis.DeployConfig
	BVMAddressesPath      string
	EVMAddressesPath      string
	BVMAllowancesPath     string
	BVMMessagesPath       string
	EVMMessagesPath       string
	Network               string
	HardhatDeployments    []string
	L1URL                 string
	StartingL1BlockNumber uint64
	L2DBPath              string
	DryRun                bool
	NoCheck               bool
}

func Migrate(cfg *Config) (*genesis.MigrationResult, error) {
	deployConfig := cfg.DeployConfig

	bvmAddresses, err := migration.NewAddresses(cfg.BVMAddressesPath)
	if err != nil {
		return nil, err
	}
	evmAddresess, err := migration.NewAddresses(cfg.EVMAddressesPath)
	if err != nil {
		return nil, err
	}
	bvmAllowances, err := migration.NewAllowances(cfg.BVMAllowancesPath)
	if err != nil {
		return nil, err
	}
	bvmMessages, err := migration.NewSentMessage(cfg.BVMMessagesPath)
	if err != nil {
		return nil, err
	}
	evmMessages, err := migration.NewSentMessage(cfg.EVMMessagesPath)
	if err != nil {
		return nil, err
	}

	migrationData := migration.MigrationData{
		BvmAddresses:  bvmAddresses,
		EvmAddresses:  evmAddresess,
		BvmAllowances: bvmAllowances,
		BvmMessages:   bvmMessages,
		EvmMessages:   evmMessages,
	}

	l1Client, err := ethclient.Dial(cfg.L1URL)
	if err != nil {
		return nil, err
	}
	var blockNumber *big.Int
	bnum := cfg.StartingL1BlockNumber
	if bnum != 0 {
		blockNumber = new(big.Int).SetUint64(bnum)
	}

	block, err := l1Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	chaindataPath := filepath.Join(cfg.L2DBPath, "geth", "chaindata")
	ancientPath := filepath.Join(chaindataPath, "ancient")
	ldb, err := rawdb.Open(
		rawdb.OpenOptions{
			Type:              "leveldb",
			Directory:         chaindataPath,
			Cache:             4096,
			Handles:           120,
			AncientsDirectory: ancientPath,
			Namespace:         "",
			ReadOnly:          false,
		})
	if err != nil {
		return nil, err
	}
	defer ldb.Close()

	return genesis.MigrateDB(ldb, deployConfig, block, &migrationData, !cfg.DryRun, cfg.NoCheck)
}
