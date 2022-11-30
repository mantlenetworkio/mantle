package flags

import (
	"github.com/urfave/cli"
	"time"
)

const envVarPrefix = "MT-BATCHER"

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
		EnvVar:   prefixEnvVar(envVarPrefix, "DISPERSER"),
	}
	GrpcPortFlag = cli.IntFlag{
		Name:     "grpc-port",
		Usage:    "Port at which node listens for grpc calls",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "GRPC_PORT"),
	}

	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "CHAIN_ID"),
	}
	ChainProviderFlag = cli.StringFlag{
		Name:     "chain-provider",
		Usage:    "Ethereum chain rpc",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "CHAIN_PROVIDER"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "GRAPH_PROVIDER"),
	}
	PrivateFlag = cli.StringFlag{
		Name:     "private",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "PRIVATE"),
	}
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_ADDRESS"),
	}
	DurationFlag = cli.IntFlag{
		Name:     "duration",
		Usage:    "Duration to store blob",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DURATION"),
	}
	TimeoutFlag = cli.IntFlag{
		Name:     "timeout",
		Usage:    "Blob timeout",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "TIMEOUT"),
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
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	MtlNetworkNameFlag,
	L1EthRpcFlag,
	L2MtlRpcFlag,
	DisperserEndpointFlag,
	GrpcPortFlag,
	ChainIdFlag,
	ChainProviderFlag,
	GraphProviderFlag,
	PrivateFlag,
	RollupAddressFlag,
	DurationFlag,
	TimeoutFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	SentryEnableFlag,
	SentryDsnFlag,
	SentryTraceRateFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
