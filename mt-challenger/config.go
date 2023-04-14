package challenger

import (
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/mantlenetworkio/mantle/mt-challenger/challenger"
	"github.com/mantlenetworkio/mantle/mt-challenger/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	L1EthRpc                  string
	L2MtlRpc                  string
	ChainId                   uint64
	PrivateKey                string
	Mnemonic                  string
	SequencerHDPath           string
	EigenContractAddress      string
	GraphProvider             string
	RetrieverSocket           string
	KzgConfig                 challenger.KzgConfig
	FromStoreNumber           uint64
	PollInterval              time.Duration
	DbPath                    string
	CheckerBatchIndex         uint64
	DisableHTTP2              bool
	NeedReRollupBatch         string
	ReRollupToolEnable        bool
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	LoggingConfig             logging.Config
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		L1EthRpc:             ctx.GlobalString(flags.L1EthRpcFlag.Name),
		ChainId:              ctx.GlobalUint64(flags.ChainIdFlag.Name),
		GraphProvider:        ctx.GlobalString(flags.GraphProviderFlag.Name),
		L2MtlRpc:             ctx.GlobalString(flags.L2MtlRpcFlag.Name),
		PrivateKey:           ctx.GlobalString(flags.PrivateKeyFlag.Name),
		Mnemonic:             ctx.GlobalString(flags.MnemonicFlag.Name),
		SequencerHDPath:      ctx.GlobalString(flags.SequencerHDPathFlag.Name),
		EigenContractAddress: ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		RetrieverSocket:      ctx.GlobalString(flags.RetrieverSocketFlag.Name),
		KzgConfig: challenger.KzgConfig{
			G1Path:    ctx.GlobalString(flags.G1PathFlag.Name),
			G2Path:    ctx.GlobalString(flags.G2PathFlag.Name),
			TableDir:  ctx.GlobalString(flags.SrsTablePathFlag.Name),
			Order:     ctx.GlobalUint64(flags.OrderFlag.Name),
			NumWorker: ctx.GlobalInt(flags.KzgWorkersFlag.Name),
		},
		ResubmissionTimeout:       ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),
		NumConfirmations:          ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		SafeAbortNonceTooLowCount: ctx.GlobalUint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		LoggingConfig:             logging.ReadCLIConfig(ctx),
		FromStoreNumber:           ctx.GlobalUint64(flags.StartStoreNumFlag.Name),
		PollInterval:              ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		DbPath:                    ctx.GlobalString(flags.DbPathFlag.Name),
		CheckerBatchIndex:         ctx.GlobalUint64(flags.CheckerBatchIndexFlag.Name),
		NeedReRollupBatch:         ctx.GlobalString(flags.NeedReRollupBatchFlag.Name),
		ReRollupToolEnable:        ctx.GlobalBool(flags.ReRollupToolEnableFlag.Name),
		DisableHTTP2:              ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
	}
	return cfg, nil
}
