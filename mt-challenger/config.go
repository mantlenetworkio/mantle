package challenger

import (
	"time"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/datalayr/common/logging"

	"github.com/mantlenetworkio/mantle/mt-challenger/challenger"
	"github.com/mantlenetworkio/mantle/mt-challenger/flags"
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
	DtlClientUrl              string
	KzgConfig                 challenger.KzgConfig
	FromStoreNumber           uint64
	PollInterval              time.Duration
	CompensatePollInterval    time.Duration
	DbPath                    string
	Passphrase                string
	CheckerBatchIndex         uint64
	UpdateBatchIndexStep      uint64
	DisableHTTP2              bool
	NeedReRollupBatch         string
	ChallengerCheckEnable     bool
	ReRollupToolEnable        bool
	DataCompensateEnable      bool
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	LoggingConfig             logging.Config
	MetricsServerEnable       bool
	MetricsHostname           string
	MetricsPort               uint64
	EnableHsm                 bool
	HsmAPIName                string
	HsmCreden                 string
	HsmAddress                string
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		L1EthRpc:             ctx.GlobalString(flags.L1EthRpcFlag.Name),
		ChainId:              ctx.GlobalUint64(flags.ChainIdFlag.Name),
		GraphProvider:        ctx.GlobalString(flags.GraphProviderFlag.Name),
		L2MtlRpc:             ctx.GlobalString(flags.L2MtlRpcFlag.Name),
		PrivateKey:           ctx.GlobalString(flags.PrivateKeyFlag.Name),
		Mnemonic:             ctx.GlobalString(flags.MnemonicFlag.Name),
		Passphrase:           ctx.GlobalString(flags.PassphraseFlag.Name),
		SequencerHDPath:      ctx.GlobalString(flags.SequencerHDPathFlag.Name),
		EigenContractAddress: ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		RetrieverSocket:      ctx.GlobalString(flags.RetrieverSocketFlag.Name),
		DtlClientUrl:         ctx.GlobalString(flags.DtlClientUrlFlag.Name),
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
		CompensatePollInterval:    ctx.GlobalDuration(flags.CompensatePollIntervalFlag.Name),
		DbPath:                    ctx.GlobalString(flags.DbPathFlag.Name),
		CheckerBatchIndex:         ctx.GlobalUint64(flags.CheckerBatchIndexFlag.Name),
		UpdateBatchIndexStep:      ctx.GlobalUint64(flags.UpdateBatchIndexStepFlag.Name),
		NeedReRollupBatch:         ctx.GlobalString(flags.NeedReRollupBatchFlag.Name),
		ChallengerCheckEnable:     ctx.GlobalBool(flags.ChallengerCheckEnableFlag.Name),
		ReRollupToolEnable:        ctx.GlobalBool(flags.ReRollupToolEnableFlag.Name),
		DataCompensateEnable:      ctx.GlobalBool(flags.DataCompensateEnableFlag.Name),
		DisableHTTP2:              ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
		MetricsServerEnable:       ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:           ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:               ctx.GlobalUint64(flags.MetricsPortFlag.Name),
		EnableHsm:                 ctx.GlobalBool(flags.EnableHsmFlag.Name),
		HsmAddress:                ctx.GlobalString(flags.HsmAddressFlag.Name),
		HsmAPIName:                ctx.GlobalString(flags.HsmAPINameFlag.Name),
		HsmCreden:                 ctx.GlobalString(flags.HsmCredenFlag.Name),
	}
	return cfg, nil
}
