package flags

import (
	"github.com/urfave/cli"
)

var (
	EthereumHttpUrlFlag = cli.StringFlag{
		Name:   "ethereum-http-url",
		Value:  "http://127.0.0.1:9545",
		Usage:  "L1 HTTP Endpoint",
		EnvVar: "GAS_PRICE_ORACLE_ETHEREUM_HTTP_URL",
	}
	EthereumWssUrlFlag = cli.StringFlag{
		Name:   "ethereum-wss-url",
		Value:  "ws://127.0.0.1:9545",
		Usage:  "L1 WSS Endpoint",
		EnvVar: "GAS_PRICE_ORACLE_ETHEREUM_WSS_URL",
	}
	LayerTwoHttpUrlFlag = cli.StringFlag{
		Name:   "layer-two-http-url",
		Value:  "http://127.0.0.1:8545",
		Usage:  "Sequencer HTTP Endpoint",
		EnvVar: "GAS_PRICE_ORACLE_LAYER_TWO_HTTP_URL",
	}
	L1ChainIDFlag = cli.Uint64Flag{
		Name:   "l1-chain-id",
		Usage:  "L1 Chain ID",
		EnvVar: "GAS_PRICE_ORACLE_L1_CHAIN_ID",
	}
	L2ChainIDFlag = cli.Uint64Flag{
		Name:   "l2-chain-id",
		Usage:  "L2 Chain ID",
		EnvVar: "GAS_PRICE_ORACLE_L2_CHAIN_ID",
	}
	GasPriceOracleAddressFlag = cli.StringFlag{
		Name:   "gas-price-oracle-address",
		Usage:  "Address of BVM_GasPriceOracle",
		Value:  "0x420000000000000000000000000000000000000F",
		EnvVar: "GAS_PRICE_ORACLE_GAS_PRICE_ORACLE_ADDRESS",
	}
	DaFeeContractAddressFlag = cli.StringFlag{
		Name:   "da-fee-contract-address",
		Usage:  "Address of DA-Fee-Contract",
		Value:  "0x9700cA34e333BCfa83ee7eA9de998a876474Dc2c",
		EnvVar: "GAS_PRICE_ORACLE_DA_FEE_CONTRACT_ADDRESS",
	}
	SCCContractAddressFlag = cli.StringFlag{
		Name:   "scc-contract-address",
		Usage:  "Address of StateCommitChain",
		Value:  "0x82e130FF187E787D5DdDFAa4f36CB59e6B1Da6dd",
		EnvVar: "GAS_PRICE_ORACLE_SCC_CONTRACT_ADDRESS",
	}
	CTCContractAddressFlag = cli.StringFlag{
		Name:   "ctc-contract-address",
		Usage:  "Address of CanonicalTransactionChain",
		Value:  "0xEd5166f12FCb48a0804B62FDccB37f59F1F1bc3B",
		EnvVar: "GAS_PRICE_ORACLE_CTC_CONTRACT_ADDRESS",
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "Private Key corresponding to BVM_GasPriceOracle Owner",
		EnvVar: "GAS_PRICE_ORACLE_PRIVATE_KEY",
	}
	TransactionGasPriceFlag = cli.Uint64Flag{
		Name:   "transaction-gas-price",
		Usage:  "Hardcoded tx.gasPrice, not setting it uses gas estimation",
		EnvVar: "GAS_PRICE_ORACLE_TRANSACTION_GAS_PRICE",
	}
	EnableL1BaseFeeFlag = cli.BoolFlag{
		Name:   "enable-l1-base-fee",
		Usage:  "Enable updating the L1 base fee",
		EnvVar: "GAS_PRICE_ORACLE_ENABLE_L1_BASE_FEE",
	}
	EnableL1OverheadFlag = cli.BoolFlag{
		Name:   "enable-l1-overhead",
		Usage:  "Enable updating the L1 overhead",
		EnvVar: "GAS_PRICE_ORACLE_ENABLE_L1_OVERHEAD",
	}
	EnableL2GasPriceFlag = cli.BoolFlag{
		Name:   "enable-l2-gas-price",
		Usage:  "Enable updating the L2 gas price",
		EnvVar: "GAS_PRICE_ORACLE_ENABLE_L2_GAS_PRICE",
	}
	EnableDaFeeFlag = cli.BoolFlag{
		Name:   "enable-da-gas-price",
		Usage:  "Enable updating the da gas price",
		EnvVar: "GAS_PRICE_ORACLE_ENABLE_DA_FEE",
	}
	BatchSizeCap = cli.IntFlag{
		Name:   "set-batch-size-cap",
		Value:  1000,
		Usage:  "Setup batch size cap",
		EnvVar: "GAS_PRICE_ORACLE_BATCH_SIZE_CAP",
	}
	BatchSizeBottom = cli.IntFlag{
		Name:   "set-batch-size-bottom",
		Value:  100,
		Usage:  "Setup batch size bottom",
		EnvVar: "GAS_PRICE_ORACLE_BATCH_SIZE_BOTTOM",
	}
	SizeGap = cli.IntFlag{
		Name:   "set-batch-size-gap",
		Value:  100,
		Usage:  "Setup batch size gap",
		EnvVar: "GAS_PRICE_ORACLE_SIZE_GAP",
	}
	StateRollupGasUsed = cli.IntFlag{
		Name:   "set-state-rollup-gas-used",
		Value:  2521687,
		Usage:  "Setup StateRollupGasUsed",
		EnvVar: "GAS_PRICE_ORACLE_STATE_ROLLUP_GAS_USED",
	}
	StateHashGasUsed = cli.IntFlag{
		Name:   "set-state-hash-gas-used",
		Value:  1412,
		Usage:  "Setup StateHashGasUsed",
		EnvVar: "GAS_PRICE_ORACLE_STATE_HASH_GAS_USED",
	}
	DataRollupGasUsed = cli.IntFlag{
		Name:   "set-data-rollup-gas-used",
		Value:  137893,
		Usage:  "Setup DataRollupGasUsed",
		EnvVar: "GAS_PRICE_ORACLE_DATA_ROLLUP_GAS_USED",
	}
	LogLevelFlag = cli.IntFlag{
		Name:   "loglevel",
		Value:  3,
		Usage:  "log level to emit to the screen",
		EnvVar: "GAS_PRICE_ORACLE_LOG_LEVEL",
	}
	FloorPriceFlag = cli.Uint64Flag{
		Name:   "floor-price",
		Value:  1,
		Usage:  "gas price floor",
		EnvVar: "GAS_PRICE_ORACLE_FLOOR_PRICE",
	}
	TargetGasPerSecondFlag = cli.Uint64Flag{
		Name:   "target-gas-per-second",
		Value:  11_000_000,
		Usage:  "target gas per second",
		EnvVar: "GAS_PRICE_ORACLE_TARGET_GAS_PER_SECOND",
	}
	MaxPercentChangePerEpochFlag = cli.Float64Flag{
		Name:   "max-percent-change-per-epoch",
		Value:  0.1,
		Usage:  "max percent change of gas price per second",
		EnvVar: "GAS_PRICE_ORACLE_MAX_PERCENT_CHANGE_PER_EPOCH",
	}
	AverageBlockGasLimitPerEpochFlag = cli.Uint64Flag{
		Name:   "average-block-gas-limit-per-epoch",
		Value:  11_000_000,
		Usage:  "average block gas limit per epoch",
		EnvVar: "GAS_PRICE_ORACLE_AVERAGE_BLOCK_GAS_LIMIT_PER_EPOCH",
	}
	EpochLengthSecondsFlag = cli.Uint64Flag{
		Name:   "epoch-length-seconds",
		Value:  10,
		Usage:  "length of epochs in seconds",
		EnvVar: "GAS_PRICE_ORACLE_EPOCH_LENGTH_SECONDS",
	}
	L1BaseFeeEpochLengthSecondsFlag = cli.Uint64Flag{
		Name:   "l1-base-fee-epoch-length-seconds",
		Value:  15,
		Usage:  "polling time for updating the L1 base fee",
		EnvVar: "GAS_PRICE_ORACLE_L1_BASE_FEE_EPOCH_LENGTH_SECONDS",
	}
	DaFeeEpochLengthSecondsFlag = cli.Uint64Flag{
		Name:   "da-fee-epoch-length-seconds",
		Value:  15,
		Usage:  "polling time for updating the Da fee",
		EnvVar: "GAS_PRICE_ORACLE_DA_FEE_EPOCH_LENGTH_SECONDS",
	}
	L1BaseFeeSignificanceFactorFlag = cli.Float64Flag{
		Name:   "l1-base-fee-significant-factor",
		Value:  0.05,
		Usage:  "only update when the L1 base fee changes by more than this factor",
		EnvVar: "GAS_PRICE_ORACLE_L1_BASE_FEE_SIGNIFICANT_FACTOR",
	}
	DaFeeSignificanceFactorFlag = cli.Float64Flag{
		Name:   "da-fee-significant-factor",
		Value:  0.05,
		Usage:  "only update when the L1 base fee changes by more than this factor",
		EnvVar: "GAS_PRICE_ORACLE_DA_FEE_SIGNIFICANT_FACTOR",
	}
	L2GasPriceSignificanceFactorFlag = cli.Float64Flag{
		Name:   "significant-factor",
		Value:  0.05,
		Usage:  "only update when the gas price changes by more than this factor",
		EnvVar: "GAS_PRICE_ORACLE_SIGNIFICANT_FACTOR",
	}
	PriceBackendURL = cli.StringFlag{
		Name:     "price-backend-url",
		Usage:    "price exchange backend url",
		EnvVar:   "PRICE_BACKEND_URL",
		Required: true,
	}
	PriceBackendUniswapURL = cli.StringFlag{
		Name:     "price-backend-uniswap-url",
		Usage:    "price backend uniswap url",
		EnvVar:   "PRICE_BACKEND_UNISWAP_URL",
		Required: true,
	}
	TokenPricerUpdateFrequencySecond = cli.Uint64Flag{
		Name:   "token-pricer-update-frequency-second",
		Value:  3,
		Usage:  "token pricer update frequency",
		EnvVar: "TOKEN_PRICER_UPDATE_FREQUENCY",
	}
	TokenRatioMode = cli.Uint64Flag{
		Name:   "token-ratio-mode",
		Value:  0,
		Usage:  "token ratio mode",
		EnvVar: "TOKEN_RATIO_MODE",
	}
	TokenPairMNTMode = cli.BoolFlag{
		Name:     "token-pair-mnt-mode",
		Usage:    "use mnt price to calculate token ratio",
		EnvVar:   "TOKEN_PAIR_MNT_MODE",
		Required: true,
	}
	WaitForReceiptFlag = cli.BoolFlag{
		Name:   "wait-for-receipt",
		Usage:  "wait for receipts when sending transactions",
		EnvVar: "GAS_PRICE_ORACLE_WAIT_FOR_RECEIPT",
	}
	MetricsEnabledFlag = cli.BoolFlag{
		Name:   "metrics",
		Usage:  "Enable metrics collection and reporting",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_ENABLE",
	}
	MetricsHTTPFlag = cli.StringFlag{
		Name:   "metrics.addr",
		Usage:  "Enable stand-alone metrics HTTP server listening interface",
		Value:  "127.0.0.1",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_HTTP",
	}
	MetricsPortFlag = cli.IntFlag{
		Name:   "metrics.port",
		Usage:  "Metrics HTTP server listening port",
		Value:  9107,
		EnvVar: "GAS_PRICE_ORACLE_METRICS_PORT",
	}
	MetricsEnableInfluxDBFlag = cli.BoolFlag{
		Name:   "metrics.influxdb",
		Usage:  "Enable metrics export/push to an external InfluxDB database",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_ENABLE_INFLUX_DB",
	}
	MetricsInfluxDBEndpointFlag = cli.StringFlag{
		Name:   "metrics.influxdb.endpoint",
		Usage:  "InfluxDB API endpoint to report metrics to",
		Value:  "http://localhost:8086",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_INFLUX_DB_ENDPOINT",
	}
	MetricsInfluxDBDatabaseFlag = cli.StringFlag{
		Name:   "metrics.influxdb.database",
		Usage:  "InfluxDB database name to push reported metrics to",
		Value:  "gas-oracle",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_INFLUX_DB_DATABASE",
	}
	MetricsInfluxDBUsernameFlag = cli.StringFlag{
		Name:   "metrics.influxdb.username",
		Usage:  "Username to authorize access to the database",
		Value:  "test",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_INFLUX_DB_USERNAME",
	}
	MetricsInfluxDBPasswordFlag = cli.StringFlag{
		Name:   "metrics.influxdb.password",
		Usage:  "Password to authorize access to the database",
		Value:  "test",
		EnvVar: "GAS_PRICE_ORACLE_METRICS_INFLUX_DB_PASSWORD",
	}
	EnableHsmFlag = cli.BoolFlag{
		Name:   "enable-hsm",
		Usage:  "Enalbe the hsm",
		EnvVar: "GAS_PRICE_ORACLE_ENABLE_HSM",
	}
	HsmAPINameFlag = cli.StringFlag{
		Name:   "hsm-api-name",
		Usage:  "the api name of hsm",
		EnvVar: "GAS_PRICE_ORACLE_HSM_API_NAME",
	}
	HsmAddressFlag = cli.StringFlag{
		Name:   "hsm-address",
		Usage:  "the address of hsm key",
		EnvVar: "GAS_PRICE_ORACLE_HSM_ADDRESS",
	}
	HsmCredenFlag = cli.StringFlag{
		Name:   "hsm-creden",
		Usage:  "the creden of hsm key",
		EnvVar: "GAS_PRICE_ORACLE_HSM_CREDEN",
	}
)

var Flags = []cli.Flag{
	EthereumHttpUrlFlag,
	EthereumWssUrlFlag,
	LayerTwoHttpUrlFlag,
	L1ChainIDFlag,
	L2ChainIDFlag,
	L1BaseFeeSignificanceFactorFlag,
	DaFeeSignificanceFactorFlag,
	GasPriceOracleAddressFlag,
	DaFeeContractAddressFlag,
	PrivateKeyFlag,
	TransactionGasPriceFlag,
	LogLevelFlag,
	FloorPriceFlag,
	TargetGasPerSecondFlag,
	MaxPercentChangePerEpochFlag,
	AverageBlockGasLimitPerEpochFlag,
	EpochLengthSecondsFlag,
	L1BaseFeeEpochLengthSecondsFlag,
	DaFeeEpochLengthSecondsFlag,
	L2GasPriceSignificanceFactorFlag,
	PriceBackendURL,
	PriceBackendUniswapURL,
	TokenPricerUpdateFrequencySecond,
	TokenRatioMode,
	TokenPairMNTMode,
	WaitForReceiptFlag,
	EnableL1BaseFeeFlag,
	EnableL1OverheadFlag,
	EnableL2GasPriceFlag,
	EnableDaFeeFlag,
	SCCContractAddressFlag,
	CTCContractAddressFlag,
	BatchSizeBottom,
	BatchSizeCap,
	SizeGap,
	StateRollupGasUsed,
	StateHashGasUsed,
	DataRollupGasUsed,
	EnableHsmFlag,
	HsmAddressFlag,
	HsmAPINameFlag,
	HsmCredenFlag,
	MetricsEnabledFlag,
	MetricsHTTPFlag,
	MetricsPortFlag,
	MetricsEnableInfluxDBFlag,
	MetricsInfluxDBEndpointFlag,
	MetricsInfluxDBDatabaseFlag,
	MetricsInfluxDBUsernameFlag,
	MetricsInfluxDBPasswordFlag,
}
