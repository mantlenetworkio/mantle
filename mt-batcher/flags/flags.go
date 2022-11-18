package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "SEQUENCER"

func prefixEnvVar(suffix string) string {
	return envVarPrefix + "_" + suffix
}

var (
	MtlRpcFlag = cli.StringFlag{
		Name:     "mtl-rpc",
		Usage:    "HTTP provider URL for L2 execution engine",
		Required: true,
		EnvVar:   opservice.PrefixEnvVar(envVarPrefix, "MTL_RPC"),
	}
	DisperserEndpointFlag = cli.StringFlag{
		Name:     "disperser",
		Usage:    "Endpoint at which disperser is available",
		Required: true,
		EnvVar:   prefixEnvVar("DISPERSER"),
	}
	GrpcPortFlag = cli.IntFlag{
		Name:     "grpc-port",
		Usage:    "Port at which node listens for grpc calls",
		Required: true,
		EnvVar:   prefixEnvVar("GRPC_PORT"),
	}

	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_ID"),
	}
	ChainProviderFlag = cli.StringFlag{
		Name:     "chain-provider",
		Usage:    "Ethereum chain rpc",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_PROVIDER"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar("GRAPH_PROVIDER"),
	}
	PrivateFlag = cli.StringFlag{
		Name:     "private",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar("PRIVATE"),
	}
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}
	DurationFlag = cli.IntFlag{
		Name:     "duration",
		Usage:    "Duration to store blob",
		Required: true,
		EnvVar:   prefixEnvVar("DURATION"),
	}
	TimeoutFlag = cli.IntFlag{
		Name:     "timeout",
		Usage:    "Blob timeout",
		Required: true,
		EnvVar:   prefixEnvVar("TIMEOUT"),
	}
)

var requiredFlags = []cli.Flag{
	MtlRpcFlag,
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

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
