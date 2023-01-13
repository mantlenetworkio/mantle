package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "MT_MESO_SERVICE"

func prefixEnvVar(suffix string) string {
	return envVarPrefix + "_" + suffix
}

var (
	EthRpcFlag = cli.StringFlag{
		Name:     "eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("ETH_RPC"),
	}
	EigenDaHttpPortFlag = cli.IntFlag{
		Name:     "eigen-da-http-port",
		Usage:    "Eigen da service port",
		Required: true,
		EnvVar:   prefixEnvVar("EIGEN_DA_HTTP_PORT"),
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
	EchoDebugFlag = cli.BoolFlag{
		Name:   "echo-debug",
		Usage:  "Echo log debug",
		EnvVar: prefixEnvVar("ECHO-DEBUG"),
	}
)

var requiredFlags = []cli.Flag{
	EthRpcFlag,
	EigenDaHttpPortFlag,
	EigenContractAddressFlag,
	RetrieverSocketFlag,
}

var optionalFlags = []cli.Flag{
	HTTP2DisableFlag,
	EchoDebugFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
