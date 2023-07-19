package mt_batcher

import (
	"time"

	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/mantlenetworkio/mantle/mt-batcher/flags"
	"github.com/urfave/cli"
)

type Config struct {
	BuildEnv                  string
	MtlNetworkName            string
	L1EthRpc                  string
	L2MtlRpc                  string
	DtlClientUrl              string
	DisperserEndpoint         string
	RetrieverSocket           string
	EigenDaHttpPort           int
	GrpcPort                  int64
	ChainId                   uint64
	GraphProvider             string
	PrivateKey                string
	Mnemonic                  string
	SequencerHDPath           string
	FeePrivateKey             string
	FeeMnemonic               string
	FeeHDPath                 string
	Passphrase                string
	EigenContractAddress      string
	EigenFeeContractAddress   string
	DataStoreDuration         uint64
	DataStoreTimeout          uint64
	SentryEnable              bool
	MainWorkerPollInterval    time.Duration
	CheckerWorkerPollInterval time.Duration
	FeeWorkerPollInterval     time.Duration
	BlockOffset               uint64
	RollUpMinTxn              uint64
	RollUpMaxSize             uint64
	EigenLayerNode            int
	EigenLogConfig            logging.Config
	MetricsServerEnable       bool
	MetricsHostname           string
	MetricsPort               uint64
	LogLevel                  string
	LogTerminal               bool
	SentryDsn                 string
	SentryTraceRate           time.Duration
	ResubmissionTimeout       time.Duration
	RetrieverTimeout          time.Duration
	PollingDuration           time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	EchoDebug                 bool
	MtlBatcherEnable          bool
	FeeSizeSec                string
	FeePerBytePerTime         uint64
	FeeModelEnable            bool
	DisableHTTP2              bool
	DbPath                    string
	CheckerBatchIndex         uint64
	CheckerEnable             bool
	EnableHsm                 bool
	HsmAPIName                string
	HsmCreden                 string
	HsmAddress                string
	HsmFeeAPIName             string
	HsmFeeAddress             string
	MinTimeoutRollupTxn       uint64
	RollupTimeout             time.Duration
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		BuildEnv:                  ctx.GlobalString(flags.BuildEnvFlag.Name),
		MtlNetworkName:            ctx.GlobalString(flags.MtlNetworkNameFlag.Name),
		L1EthRpc:                  ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2MtlRpc:                  ctx.GlobalString(flags.L2MtlRpcFlag.Name),
		DtlClientUrl:              ctx.GlobalString(flags.DtlClientUrlFlag.Name),
		DisperserEndpoint:         ctx.GlobalString(flags.DisperserEndpointFlag.Name),
		RetrieverSocket:           ctx.GlobalString(flags.RetrieverSocketFlag.Name),
		EigenDaHttpPort:           ctx.GlobalInt(flags.EigenDaHttpPortFlag.Name),
		ChainId:                   ctx.GlobalUint64(flags.ChainIdFlag.Name),
		GraphProvider:             ctx.GlobalString(flags.GraphProviderFlag.Name),
		PrivateKey:                ctx.GlobalString(flags.PrivateKeyFlag.Name),
		Mnemonic:                  ctx.GlobalString(flags.MnemonicFlag.Name),
		SequencerHDPath:           ctx.GlobalString(flags.SequencerHDPathFlag.Name),
		FeePrivateKey:             ctx.GlobalString(flags.FeePrivateKeyFlag.Name),
		FeeMnemonic:               ctx.GlobalString(flags.FeeMnemonicFlag.Name),
		FeeHDPath:                 ctx.GlobalString(flags.FeeHDPathFlag.Name),
		Passphrase:                ctx.GlobalString(flags.PassphraseFlag.Name),
		EigenContractAddress:      ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		EigenFeeContractAddress:   ctx.GlobalString(flags.EigenFeeContractAddressFlag.Name),
		DataStoreDuration:         ctx.GlobalUint64(flags.DataStoreDurationFlag.Name),
		DataStoreTimeout:          ctx.GlobalUint64(flags.DataStoreTimeoutFlag.Name),
		MainWorkerPollInterval:    ctx.GlobalDuration(flags.MainWorkerPollIntervalFlag.Name),
		CheckerWorkerPollInterval: ctx.GlobalDuration(flags.CheckerWorkerPollIntervalFlag.Name),
		FeeWorkerPollInterval:     ctx.GlobalDuration(flags.FeeWorkerPollIntervalFlag.Name),
		PollingDuration:           ctx.GlobalDuration(flags.PollingDurationFlag.Name),
		BlockOffset:               ctx.GlobalUint64(flags.BlockOffsetFlag.Name),
		RollUpMinTxn:              ctx.GlobalUint64(flags.RollUpMinTxnFlag.Name),
		RollUpMaxSize:             ctx.GlobalUint64(flags.RollUpMaxSizeFlag.Name),
		EigenLayerNode:            ctx.GlobalInt(flags.EigenLayerNodeFlag.Name),
		EigenLogConfig:            logging.ReadCLIConfig(ctx),
		RetrieverTimeout:          ctx.GlobalDuration(flags.RetrieverTimeoutFlag.Name),
		ResubmissionTimeout:       ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),
		NumConfirmations:          ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		SafeAbortNonceTooLowCount: ctx.GlobalUint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		MetricsServerEnable:       ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:           ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:               ctx.GlobalUint64(flags.MetricsPortFlag.Name),
		LogLevel:                  ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:               ctx.GlobalBool(flags.LogTerminalFlag.Name),
		SentryEnable:              ctx.GlobalBool(flags.SentryEnableFlag.Name),
		SentryDsn:                 ctx.GlobalString(flags.SentryDsnFlag.Name),
		SentryTraceRate:           ctx.GlobalDuration(flags.SentryTraceRateFlag.Name),
		FeeSizeSec:                ctx.GlobalString(flags.FeeSizeSecFlag.Name),
		FeePerBytePerTime:         ctx.GlobalUint64(flags.FeePerBytePerTimeFlag.Name),
		FeeModelEnable:            ctx.GlobalBool(flags.FeeModelEnableFlags.Name),
		DisableHTTP2:              ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
		EchoDebug:                 ctx.GlobalBool(flags.EchoDebugFlag.Name),
		MtlBatcherEnable:          ctx.GlobalBool(flags.MtlBatcherEnableFlag.Name),
		DbPath:                    ctx.GlobalString(flags.DbPathFlag.Name),
		CheckerBatchIndex:         ctx.GlobalUint64(flags.CheckerBatchIndexFlag.Name),
		CheckerEnable:             ctx.GlobalBool(flags.CheckerEnableFlag.Name),
		EnableHsm:                 ctx.GlobalBool(flags.EnableHsmFlag.Name),
		HsmAddress:                ctx.GlobalString(flags.HsmAddressFlag.Name),
		HsmAPIName:                ctx.GlobalString(flags.HsmAPINameFlag.Name),
		HsmCreden:                 ctx.GlobalString(flags.HsmCredenFlag.Name),
		HsmFeeAPIName:             ctx.GlobalString(flags.HsmFeeAPINameFlag.Name),
		HsmFeeAddress:             ctx.GlobalString(flags.HsmFeeAddressFlag.Name),
		MinTimeoutRollupTxn:       ctx.GlobalUint64(flags.MinTimeoutRollupTxnFlag.Name),
		RollupTimeout:             ctx.GlobalDuration(flags.RollupTimeoutFlag.Name),
	}
	return cfg, nil
}
