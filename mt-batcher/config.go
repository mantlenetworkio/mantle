package mt_batcher

import (
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/mantlenetworkio/mantle/mt-batcher/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	BuildEnv                  string
	MtlNetworkName            string
	L1EthRpc                  string
	L2MtlRpc                  string
	DisperserEndpoint         string
	GrpcPort                  int64
	ChainId                   uint64
	GraphProvider             string
	PrivateKey                string
	Mnemonic                  string
	SequencerHDPath           string
	EigenContractAddress      string
	EigenFeeContractAddress   string
	DataStoreDuration         uint64
	DataStoreTimeout          uint64
	SentryEnable              bool
	PollInterval              time.Duration
	BlockOffset               uint64
	RollUpMinSize             uint64
	RollUpMaxSize             uint64
	EigenLayerNode            int
	EigenLogConfig            logging.Config
	LogLevel                  string
	LogTerminal               bool
	SentryDsn                 string
	SentryTraceRate           time.Duration
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	FeeSizeSec                string
	FeeModelEnable            bool
	DisableHTTP2              bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		BuildEnv:                  ctx.GlobalString(flags.BuildEnvFlag.Name),
		MtlNetworkName:            ctx.GlobalString(flags.MtlNetworkNameFlag.Name),
		L1EthRpc:                  ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2MtlRpc:                  ctx.GlobalString(flags.L2MtlRpcFlag.Name),
		DisperserEndpoint:         ctx.GlobalString(flags.DisperserEndpointFlag.Name),
		ChainId:                   ctx.GlobalUint64(flags.ChainIdFlag.Name),
		GraphProvider:             ctx.GlobalString(flags.GraphProviderFlag.Name),
		PrivateKey:                ctx.GlobalString(flags.PrivateKeyFlag.Name),
		Mnemonic:                  ctx.GlobalString(flags.MnemonicFlag.Name),
		SequencerHDPath:           ctx.GlobalString(flags.SequencerHDPathFlag.Name),
		EigenContractAddress:      ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		EigenFeeContractAddress:   ctx.GlobalString(flags.EigenFeeContractAddressFlag.Name),
		DataStoreDuration:         ctx.GlobalUint64(flags.DataStoreDurationFlag.Name),
		DataStoreTimeout:          ctx.GlobalUint64(flags.DataStoreTimeoutFlag.Name),
		PollInterval:              ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		BlockOffset:               ctx.GlobalUint64(flags.BlockOffsetFlag.Name),
		RollUpMinSize:             ctx.GlobalUint64(flags.RollUpMinSizeFlag.Name),
		RollUpMaxSize:             ctx.GlobalUint64(flags.RollUpMaxSizeFlag.Name),
		EigenLayerNode:            ctx.GlobalInt(flags.EigenLayerNodeFlag.Name),
		EigenLogConfig:            logging.ReadCLIConfig(ctx),
		ResubmissionTimeout:       ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),
		NumConfirmations:          ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		SafeAbortNonceTooLowCount: ctx.GlobalUint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		LogLevel:                  ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:               ctx.GlobalBool(flags.LogTerminalFlag.Name),
		SentryEnable:              ctx.GlobalBool(flags.SentryEnableFlag.Name),
		SentryDsn:                 ctx.GlobalString(flags.SentryDsnFlag.Name),
		SentryTraceRate:           ctx.GlobalDuration(flags.SentryTraceRateFlag.Name),
		FeeSizeSec:                ctx.GlobalString(flags.FeeSizeSecFlag.Name),
		FeeModelEnable:            ctx.GlobalBool(flags.FeeModelEnableFlags.Name),
		DisableHTTP2:              ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
	}
	return cfg, nil
}
