package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "MT_MESO_SERVICE"

func prefixEnvVar(suffix string) string {
	return envVarPrefix + "_" + suffix
}

var (
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar("GRAPH_PROVIDER"),
	}
	EigenContractAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar("EIGEN_CONTRACT_ADDRESS"),
	}
	RetrieverSocketFlag = cli.StringFlag{
		Name:     "retriever-socket",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar("RETRIEVER_SOCKET"),
	}
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar("HTTP2_DISABLE"),
	}
)

var requiredFlags = []cli.Flag{
	GraphProviderFlag,
	EigenContractAddressFlag,
	RetrieverSocketFlag,
}

var optionalFlags = []cli.Flag{
	HTTP2DisableFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
