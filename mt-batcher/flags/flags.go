package flags

import (
	"github.com/urfave/cli"
	"time"
)

const envVarPrefix = "MT_BATCHER"

func prefixEnvVar(prefix, suffix string) string {
	return prefix + "_" + suffix
}

var (
	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   "BUILD_ENV",
	}
	MtlNetworkNameFlag = cli.StringFlag{
		Name:     "mtl-network-name",
		Usage:    "mantle network name",
		Required: true,
		EnvVar:   "MTL_NETWORK_NAME",
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	L2MtlRpcFlag = cli.StringFlag{
		Name:     "l2-mtl-rpc",
		Usage:    "HTTP provider URL for L2 execution engine",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "L2_MTL_RPC"),
	}
	DisperserEndpointFlag = cli.StringFlag{
		Name:     "disperser",
		Usage:    "Endpoint at which disperser is available",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DISPERSER_ENDPOINT"),
	}

	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "CHAIN_ID"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "GRAPH_PROVIDER"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:     "private",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "PRIVATE_KEY"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the proposer",
		EnvVar: prefixEnvVar(envVarPrefix, "MNEMONIC"),
	}
	SequencerHDPathFlag = cli.StringFlag{
		Name: "sequencer-hd-path",
		Usage: "The HD path used to derive the sequencer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: prefixEnvVar(envVarPrefix, "SEQUENCER_HD_PATH"),
	}
	EigenContractAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_ADDRESS"),
	}
	EigenFeeContractAddressFlag = cli.StringFlag{
		Name:     "eigen-fee-address",
		Usage:    "Address of the datalayr fee contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_FEE_ADDRESS"),
	}
	BlockOffsetFlag = cli.Uint64Flag{
		Name:   "block-offset",
		Usage:  "The offset between the eigenda contract start and the L2 geth blocks",
		Value:  1,
		EnvVar: prefixEnvVar(envVarPrefix, "BLOCK_OFFSET"),
	}
	RollUpMinSizeFlag = cli.Uint64Flag{
		Name:   "rollup-min-size",
		Usage:  "Rollup transaction min size data for eigen da",
		Value:  1000,
		EnvVar: prefixEnvVar(envVarPrefix, "ROLLUP_MIN_SIZE"),
	}
	FeeSizeSecFlag = cli.StringFlag{
		Name:   "fee-size-sec",
		Usage:  "Rollup transaction fee size",
		Value:  "102400", //
		EnvVar: prefixEnvVar(envVarPrefix, "FEE_SIZE_SEC"),
	}
	RollUpMaxSizeFlag = cli.Uint64Flag{
		Name:   "rollup-max-size",
		Usage:  "Rollup transaction max size data for eigen da",
		Value:  31600, // ktz for order is 3000
		EnvVar: prefixEnvVar(envVarPrefix, "ROLLUP_MAX_SIZE"),
	}
	EigenLayerNodeFlag = cli.IntFlag{
		Name:   "eigen-layer-node",
		Usage:  "The offset between the CTC contract start and the L2 geth blocks",
		Value:  4,
		EnvVar: prefixEnvVar(envVarPrefix, "EIGEN_LAYER_NODE"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "RESUBMISSION_TIMEOUT"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "NUM_CONFIRMATIONS"),
	}
	SafeAbortNonceTooLowCountFlag = cli.Uint64Flag{
		Name: "safe-abort-nonce-too-low-count",
		Usage: "Number of ErrNonceTooLow observations required to " +
			"give up on a tx at a particular nonce without receiving " +
			"confirmation",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name:     "poll-interval",
		Usage:    "Delay between querying L2 for more transactions and creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "POLL_INTERVAL"),
	}
	DataStoreDurationFlag = cli.IntFlag{
		Name:     "duration",
		Usage:    "Duration to store blob",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DATA_STORE_DURATION"),
	}
	DataStoreTimeoutFlag = cli.IntFlag{
		Name:     "timeout",
		Usage:    "Blob timeout",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DATA_STORE_TIMEOUT"),
	}
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_TERMINAL"),
	}
	FeeModelEnableFlags = cli.BoolFlag{
		Name:   "fee-model-enable",
		Usage:  "fee model enable",
		EnvVar: prefixEnvVar(envVarPrefix, "FEE_MODEL_ENABLE"),
	}
	SentryEnableFlag = cli.BoolFlag{
		Name:   "sentry-enable",
		Usage:  "Whether or not to enable Sentry. If true, sentry-dsn must also be set",
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_ENABLE"),
	}
	SentryDsnFlag = cli.StringFlag{
		Name:   "sentry-dsn",
		Usage:  "Sentry data source name",
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_DSN"),
	}
	SentryTraceRateFlag = cli.DurationFlag{
		Name:   "sentry-trace-rate",
		Usage:  "Sentry trace rate",
		Value:  50 * time.Millisecond,
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_TRACE_RATE"),
	}
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar(envVarPrefix, "HTTP2_DISABLE"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	MtlNetworkNameFlag,
	L1EthRpcFlag,
	L2MtlRpcFlag,
	DisperserEndpointFlag,
	ChainIdFlag,
	GraphProviderFlag,
	PrivateKeyFlag,
	MnemonicFlag,
	SequencerHDPathFlag,
	EigenContractAddressFlag,
	EigenFeeContractAddressFlag,
	BlockOffsetFlag,
	RollUpMinSizeFlag,
	RollUpMaxSizeFlag,
	FeeSizeSecFlag,
	PollIntervalFlag,
	DataStoreDurationFlag,
	DataStoreTimeoutFlag,
	EigenLayerNodeFlag,
	ResubmissionTimeoutFlag,
	NumConfirmationsFlag,
	FeeModelEnableFlags,
	SafeAbortNonceTooLowCountFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	SentryEnableFlag,
	SentryDsnFlag,
	SentryTraceRateFlag,
	HTTP2DisableFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
