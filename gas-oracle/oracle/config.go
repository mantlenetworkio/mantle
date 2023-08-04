package oracle

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/gas-oracle/flags"
	"github.com/urfave/cli"
)

// Config represents the configuration options for the gas oracle
type Config struct {
	l1ChainID                        *big.Int
	l2ChainID                        *big.Int
	ethereumHttpUrl                  string
	ethereumWssUrl                   string
	layerTwoHttpUrl                  string
	gasPriceOracleAddress            common.Address
	daFeeContractAddress             common.Address
	sccContractAddress               common.Address
	ctcContractAddress               common.Address
	privateKey                       *ecdsa.PrivateKey
	gasPrice                         *big.Int
	waitForReceipt                   bool
	floorPrice                       uint64
	targetGasPerSecond               uint64
	maxPercentChangePerEpoch         float64
	averageBlockGasLimitPerEpoch     uint64
	epochLengthSeconds               uint64
	l1BaseFeeEpochLengthSeconds      uint64
	daFeeEpochLengthSeconds          uint64
	l2GasPriceSignificanceFactor     float64
	PriceBackendURL                  string
	PriceBackendUniswapURL           string
	tokenPricerUpdateFrequencySecond uint64
	tokenRatioMode                   uint64
	tokenPairMNTMode                 bool
	l1BaseFeeSignificanceFactor      float64
	daFeeSignificanceFactor          float64
	enableL1BaseFee                  bool
	enableL1Overhead                 bool
	enableL2GasPrice                 bool
	enableDaFee                      bool
	// hsm config
	EnableHsm  bool
	HsmAPIName string
	HsmCreden  string
	HsmAddress string
	// overhead
	batchSizeBottom    int
	batchSizeCap       int
	sizeGap            int
	stateRollupGasUsed *big.Int
	stateHashGasUsed   *big.Int
	dataRollupGasUsed  *big.Int
	// Metrics config
	MetricsEnabled          bool
	MetricsHTTP             string
	MetricsPort             int
	MetricsEnableInfluxDB   bool
	MetricsInfluxDBEndpoint string
	MetricsInfluxDBDatabase string
	MetricsInfluxDBUsername string
	MetricsInfluxDBPassword string
}

// NewConfig creates a new Config
func NewConfig(ctx *cli.Context) *Config {
	cfg := Config{}
	cfg.ethereumHttpUrl = ctx.GlobalString(flags.EthereumHttpUrlFlag.Name)
	cfg.ethereumWssUrl = ctx.GlobalString(flags.EthereumWssUrlFlag.Name)
	cfg.layerTwoHttpUrl = ctx.GlobalString(flags.LayerTwoHttpUrlFlag.Name)
	addr := ctx.GlobalString(flags.GasPriceOracleAddressFlag.Name)
	cfg.gasPriceOracleAddress = common.HexToAddress(addr)
	daFeeContractAddress := ctx.GlobalString(flags.DaFeeContractAddressFlag.Name)
	cfg.daFeeContractAddress = common.HexToAddress(daFeeContractAddress)
	cfg.sccContractAddress = common.HexToAddress(ctx.GlobalString(flags.SCCContractAddressFlag.Name))
	cfg.ctcContractAddress = common.HexToAddress(ctx.GlobalString(flags.CTCContractAddressFlag.Name))
	cfg.targetGasPerSecond = ctx.GlobalUint64(flags.TargetGasPerSecondFlag.Name)
	cfg.maxPercentChangePerEpoch = ctx.GlobalFloat64(flags.MaxPercentChangePerEpochFlag.Name)
	cfg.averageBlockGasLimitPerEpoch = ctx.GlobalUint64(flags.AverageBlockGasLimitPerEpochFlag.Name)
	cfg.epochLengthSeconds = ctx.GlobalUint64(flags.EpochLengthSecondsFlag.Name)
	cfg.l1BaseFeeEpochLengthSeconds = ctx.GlobalUint64(flags.L1BaseFeeEpochLengthSecondsFlag.Name)
	cfg.daFeeEpochLengthSeconds = ctx.GlobalUint64(flags.DaFeeEpochLengthSecondsFlag.Name)
	cfg.l2GasPriceSignificanceFactor = ctx.GlobalFloat64(flags.L2GasPriceSignificanceFactorFlag.Name)
	cfg.PriceBackendURL = ctx.GlobalString(flags.PriceBackendURL.Name)
	cfg.PriceBackendUniswapURL = ctx.GlobalString(flags.PriceBackendUniswapURL.Name)
	cfg.tokenPricerUpdateFrequencySecond = ctx.GlobalUint64(flags.TokenPricerUpdateFrequencySecond.Name)
	cfg.tokenRatioMode = ctx.GlobalUint64(flags.TokenRatioMode.Name)
	cfg.tokenPairMNTMode = ctx.GlobalBool(flags.TokenPairMNTMode.Name)
	cfg.floorPrice = ctx.GlobalUint64(flags.FloorPriceFlag.Name)
	cfg.l1BaseFeeSignificanceFactor = ctx.GlobalFloat64(flags.L1BaseFeeSignificanceFactorFlag.Name)
	cfg.daFeeSignificanceFactor = ctx.GlobalFloat64(flags.DaFeeSignificanceFactorFlag.Name)
	cfg.enableL1BaseFee = ctx.GlobalBool(flags.EnableL1BaseFeeFlag.Name)
	cfg.enableL1Overhead = ctx.GlobalBool(flags.EnableL1OverheadFlag.Name)
	cfg.enableL2GasPrice = ctx.GlobalBool(flags.EnableL2GasPriceFlag.Name)
	cfg.enableDaFee = ctx.GlobalBool(flags.EnableDaFeeFlag.Name)
	cfg.EnableHsm = ctx.GlobalBool(flags.EnableHsmFlag.Name)
	cfg.HsmAddress = ctx.GlobalString(flags.HsmAddressFlag.Name)
	cfg.HsmAPIName = ctx.GlobalString(flags.HsmAPINameFlag.Name)
	cfg.HsmCreden = ctx.GlobalString(flags.HsmCredenFlag.Name)

	cfg.batchSizeCap = ctx.GlobalInt(flags.BatchSizeCap.Name)
	cfg.batchSizeBottom = ctx.GlobalInt(flags.BatchSizeBottom.Name)
	cfg.sizeGap = ctx.GlobalInt(flags.SizeGap.Name)
	cfg.stateRollupGasUsed = big.NewInt(ctx.GlobalInt64(flags.StateRollupGasUsed.Name))
	cfg.stateHashGasUsed = big.NewInt(ctx.GlobalInt64(flags.StateHashGasUsed.Name))
	cfg.dataRollupGasUsed = big.NewInt(ctx.GlobalInt64(flags.DataRollupGasUsed.Name))

	if cfg.EnableHsm {
		log.Info("gasoracle", "enable hsm", cfg.EnableHsm,
			"hsm address", cfg.HsmAddress)
	} else {
		if ctx.GlobalIsSet(flags.PrivateKeyFlag.Name) {
			hex := ctx.GlobalString(flags.PrivateKeyFlag.Name)
			hex = strings.TrimPrefix(hex, "0x")
			key, err := crypto.HexToECDSA(hex)
			if err != nil {
				log.Error(fmt.Sprintf("Option %q: %v", flags.PrivateKeyFlag.Name, err))
			}
			cfg.privateKey = key
		} else {
			log.Crit("No private key configured")
		}
	}

	if ctx.GlobalIsSet(flags.L1ChainIDFlag.Name) {
		chainID := ctx.GlobalUint64(flags.L1ChainIDFlag.Name)
		cfg.l1ChainID = new(big.Int).SetUint64(chainID)
	}
	if ctx.GlobalIsSet(flags.L2ChainIDFlag.Name) {
		chainID := ctx.GlobalUint64(flags.L2ChainIDFlag.Name)
		cfg.l2ChainID = new(big.Int).SetUint64(chainID)
	}

	if ctx.GlobalIsSet(flags.TransactionGasPriceFlag.Name) {
		gasPrice := ctx.GlobalUint64(flags.TransactionGasPriceFlag.Name)
		cfg.gasPrice = new(big.Int).SetUint64(gasPrice)
	}

	if ctx.GlobalIsSet(flags.WaitForReceiptFlag.Name) {
		cfg.waitForReceipt = true
	}

	cfg.MetricsEnabled = ctx.GlobalBool(flags.MetricsEnabledFlag.Name)
	cfg.MetricsHTTP = ctx.GlobalString(flags.MetricsHTTPFlag.Name)
	cfg.MetricsPort = ctx.GlobalInt(flags.MetricsPortFlag.Name)
	cfg.MetricsEnableInfluxDB = ctx.GlobalBool(flags.MetricsEnableInfluxDBFlag.Name)
	cfg.MetricsInfluxDBEndpoint = ctx.GlobalString(flags.MetricsInfluxDBEndpointFlag.Name)
	cfg.MetricsInfluxDBDatabase = ctx.GlobalString(flags.MetricsInfluxDBDatabaseFlag.Name)
	cfg.MetricsInfluxDBUsername = ctx.GlobalString(flags.MetricsInfluxDBUsernameFlag.Name)
	cfg.MetricsInfluxDBPassword = ctx.GlobalString(flags.MetricsInfluxDBPasswordFlag.Name)

	return &cfg
}
