package flags

import (
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/urfave/cli"
)

const envVarPrefix = "DA_CHALLENGER"

func prefixEnvVar(suffix string) string {
	return envVarPrefix + "_" + suffix
}

var (
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2MtlRpcFlag = cli.StringFlag{
		Name:     "l2-mtl-rpc",
		Usage:    "HTTP provider URL for L2 execution engine",
		Required: true,
		EnvVar:   prefixEnvVar("L2_MTL_RPC"),
	}
	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_ID"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar("GRAPH_PROVIDER"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar("PRIVATE_KEY"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the proposer",
		EnvVar: prefixEnvVar("MNEMONIC"),
	}
	SequencerHDPathFlag = cli.StringFlag{
		Name: "sequencer-hd-path",
		Usage: "The HD path used to derive the sequencer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: prefixEnvVar("SEQUENCER_HD_PATH"),
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
	G1PathFlag = cli.StringFlag{
		Name:     "g1-path",
		Usage:    "Path to G1 SRS",
		Required: true,
		EnvVar:   prefixEnvVar("G1_PATH"),
	}
	G2PathFlag = cli.StringFlag{
		Name:     "g2-path",
		Usage:    "Path to G2 SRS",
		Required: true,
		EnvVar:   prefixEnvVar("G2_PATH"),
	}
	SrsTablePathFlag = cli.StringFlag{
		Name:     "srs-table-path",
		Usage:    "Path to SRS Table directory",
		Required: true,
		EnvVar:   prefixEnvVar("SRS_TABLE_PATH"),
	}
	OrderFlag = cli.StringFlag{
		Name:     "order",
		Usage:    "Order of the SRS",
		Required: true,
		EnvVar:   prefixEnvVar("ORDER"),
	}
	KzgWorkersFlag = cli.IntFlag{
		Name:     "kzg-num-workers",
		Usage:    "Order of the SRS",
		Required: false,
		Value:    4,
		EnvVar:   prefixEnvVar("KZG_NUM_WORKERS"),
	}
	StartStoreNumFlag = cli.Uint64Flag{
		Name:     "starting-store-numer",
		Usage:    "Store number from which challenger should pull",
		Required: false,
		Value:    4,
		EnvVar:   prefixEnvVar("STARTING_STORE_NUMER"),
	}
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar("HTTP2_DISABLE"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2MtlRpcFlag,
	ChainIdFlag,
	GraphProviderFlag,
	PrivateKeyFlag,
	EigenContractAddressFlag,
	G1PathFlag,
	G2PathFlag,
	SrsTablePathFlag,
	OrderFlag,
	RetrieverSocketFlag,
}

var optionalFlags = []cli.Flag{
	KzgWorkersFlag,
	HTTP2DisableFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
	Flags = append(Flags, logging.CLIFlags(envVarPrefix)...)
}

var Flags []cli.Flag
