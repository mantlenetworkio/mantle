package mt_batcher

import (
	"github.com/mantlenetworkio/mantle/mt-batcher/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	BuildEnv          string
	MtlNetworkName    string
	L1EthRpc          string
	L2MtlRpc          string
	DisperserEndpoint string
	GrpcPort          int64
	ChainId           uint64
	ChainProvider     string
	GraphProvider     string
	PrivateKey        string
	Mnemonic          string
	SequencerHDPath   string
	RollupAddress     string
	DataStoreDuration uint64
	DataStoreTimeout  uint64
	SentryEnable      bool
	PollInterval      time.Duration

	LogLevel        string
	LogTerminal     bool
	SentryDsn       string
	SentryTraceRate time.Duration

	DisableHTTP2 bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		BuildEnv:          ctx.GlobalString(flags.BuildEnvFlag.Name),
		MtlNetworkName:    ctx.GlobalString(flags.MtlNetworkNameFlag.Name),
		L1EthRpc:          ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2MtlRpc:          ctx.GlobalString(flags.L2MtlRpcFlag.Name),
		DisperserEndpoint: ctx.GlobalString(flags.DisperserEndpointFlag.Name),
		GrpcPort:          ctx.GlobalInt64(flags.GrpcPortFlag.Name),
		ChainId:           ctx.GlobalUint64(flags.GrpcPortFlag.Name),
		ChainProvider:     ctx.GlobalString(flags.GrpcPortFlag.Name),
		GraphProvider:     ctx.GlobalString(flags.GrpcPortFlag.Name),
		PrivateKey:        ctx.GlobalString(flags.PrivateKeyFlag.Name),
		Mnemonic:          ctx.GlobalString(flags.MnemonicFlag.Name),
		SequencerHDPath:   ctx.GlobalString(flags.SequencerHDPathFlag.Name),
		RollupAddress:     ctx.GlobalString(flags.GrpcPortFlag.Name),
		DataStoreDuration: ctx.GlobalUint64(flags.DataStoreDurationFlag.Name),
		DataStoreTimeout:  ctx.GlobalUint64(flags.DataStoreTimeoutFlag.Name),
		PollInterval:      ctx.GlobalDuration(flags.PollIntervalFlag.Name),

		LogLevel:        ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:     ctx.GlobalBool(flags.LogTerminalFlag.Name),
		SentryEnable:    ctx.GlobalBool(flags.SentryEnableFlag.Name),
		SentryDsn:       ctx.GlobalString(flags.SentryDsnFlag.Name),
		SentryTraceRate: ctx.GlobalDuration(flags.SentryTraceRateFlag.Name),
		DisableHTTP2:    ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
	}
	return cfg, nil
}
