package mt_batcher

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-batcher/metrics"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/restorer"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/sequencer"

	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/urfave/cli"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}

		mantleBatch, err := NewMantleBatch(cfg)
		if err != nil {
			return err
		}
		if err := mantleBatch.Start(); err != nil {
			return err
		}
		defer mantleBatch.Stop()

		log.Debug("mantle batcher started")

		<-(chan struct{})(nil)

		return nil
	}
}

type MantleBatch struct {
	ctx             context.Context
	cfg             Config
	sequencerDriver *sequencer.Driver
	daService       *restorer.DaService
	metrics         *metrics.MtBatchMetrics
}

func NewMantleBatch(cfg Config) (*MantleBatch, error) {
	ctx := context.Background()

	sequencerPrivKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"MtBatcher", cfg.Mnemonic, cfg.SequencerHDPath,
		cfg.PrivateKey, cfg.EigenContractAddress, cfg.Passphrase,
	)
	if err != nil {
		return nil, err
	}

	mtFeePrivateKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"MtBatcher", cfg.FeeMnemonic, cfg.FeeHDPath,
		cfg.FeePrivateKey, cfg.EigenFeeContractAddress, cfg.Passphrase,
	)
	if err != nil {
		return nil, err
	}

	l1Client, err := l1l2client.L1EthClientWithTimeout(ctx, cfg.L1EthRpc, cfg.DisableHTTP2)
	if err != nil {
		return nil, err
	}

	chainID, err := l1Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
	if err != nil {
		return nil, err
	}

	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		ethc.Address(common.HexToAddress(cfg.EigenContractAddress)), l1Client,
	)
	if err != nil {
		log.Error("MtBatcher binding eigenda contract fail", "err", err)
		return nil, err
	}

	logger, err := logging.GetLogger(cfg.EigenLogConfig)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrChainABI,
	))
	if err != nil {
		log.Error("MtBatcher parse eigenda contract abi fail", "err", err)
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		ethc.Address(common.HexToAddress(cfg.EigenContractAddress)), parsed, l1Client, l1Client,
		l1Client,
	)
	log.Debug("contract init success", "EigenContractAddress", cfg.EigenContractAddress)

	eigenFeeContract, err := bindings.NewBVMEigenDataLayrFee(
		ethc.Address(common.HexToAddress(cfg.EigenFeeContractAddress)),
		l1Client,
	)
	if err != nil {
		log.Error("MtBatcher binding eigen fee contract fail", "err", err)
		return nil, err
	}

	feeParsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrFeeABI,
	))
	if err != nil {
		log.Error("MtBatcher parse eigen fee contract abi fail", "err", err)
		return nil, err
	}
	rawEigenFeeContract := bind.NewBoundContract(
		ethc.Address(common.HexToAddress(cfg.EigenFeeContractAddress)), feeParsed, l1Client, l1Client,
		l1Client,
	)

	driverConfig := &sequencer.DriverConfig{
		L1Client:                  l1Client,
		L2Client:                  l2Client,
		L1ChainID:                 chainID,
		DtlClientUrl:              cfg.DtlClientUrl,
		EigenDaContract:           eigenContract,
		RawEigenContract:          rawEigenContract,
		EigenFeeContract:          eigenFeeContract,
		RawEigenFeeContract:       rawEigenFeeContract,
		FeeModelEnable:            cfg.FeeModelEnable,
		FeeSizeSec:                cfg.FeeSizeSec,
		FeePerBytePerTime:         cfg.FeePerBytePerTime,
		Logger:                    logger,
		PrivKey:                   sequencerPrivKey,
		FeePrivKey:                mtFeePrivateKey,
		BlockOffset:               cfg.BlockOffset,
		RollUpMinTxn:              cfg.RollUpMinTxn,
		RollUpMaxSize:             cfg.RollUpMaxSize,
		EigenLayerNode:            cfg.EigenLayerNode,
		DataStoreDuration:         uint64(cfg.DataStoreDuration),
		DataStoreTimeout:          cfg.DataStoreTimeout,
		DisperserSocket:           cfg.DisperserEndpoint,
		MainWorkerPollInterval:    cfg.MainWorkerPollInterval,
		CheckerWorkerPollInterval: cfg.CheckerWorkerPollInterval,
		FeeWorkerPollInterval:     cfg.FeeWorkerPollInterval,
		GraphPollingDuration:      cfg.PollingDuration,
		DbPath:                    cfg.DbPath,
		CheckerBatchIndex:         cfg.CheckerBatchIndex,
		CheckerEnable:             cfg.CheckerEnable,
		GraphProvider:             cfg.GraphProvider,
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		Metrics:                   metrics.NewMtBatchBase(),
		EnableHsm:                 cfg.EnableHsm,
		HsmAddress:                cfg.HsmAddress,
		HsmAPIName:                cfg.HsmAPIName,
		HsmCreden:                 cfg.HsmCreden,
		HsmFeeAPIName:             cfg.HsmFeeAPIName,
		HsmFeeAddress:             cfg.HsmFeeAddress,
		MinTimeoutRollupTxn:       cfg.MinTimeoutRollupTxn,
		RollupTimeout:             cfg.RollupTimeout,
	}
	if cfg.MinTimeoutRollupTxn >= cfg.RollUpMinTxn {
		log.Error("new driver fail", "err", "config value error : MinTimeoutRollupTxn should less than RollUpMinTxn  MinTimeoutRollupTxn(%v)>RollUpMinTxn(%v)", cfg.MinTimeoutRollupTxn, cfg.RollUpMinTxn)
		return nil, errors.New("config value error : MinTimeoutRollupTxn should less than RollUpMinTxn")
	}
	log.Debug("hsm",
		"enablehsm", driverConfig.EnableHsm, "hsmaddress", driverConfig.HsmAddress,
		"hsmapiname", driverConfig.HsmAPIName, "HsmFeeAPIName", driverConfig.HsmFeeAPIName, "HsmFeeAddress", driverConfig.HsmFeeAddress)
	driver, err := sequencer.NewDriver(ctx, driverConfig)
	if err != nil {
		log.Error("new driver fail", "err", err)
		return nil, err
	}
	daServiceConfig := &restorer.DaServiceConfig{
		EigenContract:   eigenContract,
		RetrieverSocket: cfg.RetrieverSocket,
		GraphProvider:   cfg.GraphProvider,
		DaServicePort:   cfg.EigenDaHttpPort,
		EigenLayerNode:  cfg.EigenLayerNode,
	}
	daService, err := restorer.NewDaService(ctx, daServiceConfig)
	if err != nil {
		log.Error("new da http service fail", "err", err)
		return nil, err
	}
	return &MantleBatch{
		ctx:             ctx,
		cfg:             cfg,
		sequencerDriver: driver,
		daService:       daService,
	}, nil
}

func (mb *MantleBatch) Start() error {
	if mb.cfg.MtlBatcherEnable {
		if err := mb.sequencerDriver.Start(); err != nil {
			return err
		}
	}
	go func() {
		err := mb.daService.Start()
		if err != nil {
			log.Error("da server failed to start", "err", err)
		}
	}()
	if mb.cfg.MetricsServerEnable {
		go metrics.StartServer(mb.cfg.MetricsHostname, mb.cfg.MetricsPort)
	}
	return nil
}

func (mb *MantleBatch) Stop() {
	mb.daService.Stop()
	mb.sequencerDriver.Stop()
}
