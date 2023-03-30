package mt_batcher

import (
	"context"
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-batcher/metrics"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/restorer"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/sequencer"
	"github.com/urfave/cli"
	"math/big"
	"strings"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}

		log.Info("Init Mantle Batcher success", "CurrentVersion", gitVersion)

		mantleBatch, err := NewMantleBatch(cfg)
		if err != nil {
			return err
		}
		if err := mantleBatch.Start(); err != nil {
			return err
		}
		defer mantleBatch.Stop()

		log.Info("mantle batcher started")

		<-(chan struct{})(nil)

		return nil
	}
}

type MantleBatch struct {
	ctx             context.Context
	cfg             Config
	sequencerDriver *sequencer.Driver
	daService       *restorer.DaService
	metrics         *metrics.Metrics
}

func NewMantleBatch(cfg Config) (*MantleBatch, error) {
	ctx := context.Background()

	sequencerPrivKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"MtBatcher", cfg.Mnemonic, cfg.SequencerHDPath,
		cfg.PrivateKey, cfg.EigenContractAddress,
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
	log.Info("l1Client init success", "chainID", chainID)

	l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
	if err != nil {
		return nil, err
	}
	log.Info("l2Client init success")

	mtBatherPrivateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
	if err != nil {
		return nil, err
	}

	signer := func(chainID *big.Int) sequencer.SignerFn {
		s := common2.PrivateKeySignerFn(mtBatherPrivateKey, chainID)
		return func(_ context.Context, addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
			return s(addr, tx)
		}
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
	eignenABI, err := bindings.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		log.Error("MtBatcher get eigenda contract abi fail", "err", err)
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		ethc.Address(common.HexToAddress(cfg.EigenContractAddress)), parsed, l1Client, l1Client,
		l1Client,
	)
	log.Info("contract init success", "EigenContractAddress", cfg.EigenContractAddress)

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
	eigenFeeABI, err := bindings.BVMEigenDataLayrFeeMetaData.GetAbi()
	if err != nil {
		log.Error("MtBatcher get eigen fee contract abi fail", "err", err)
		return nil, err
	}
	rawEigenFeeContract := bind.NewBoundContract(
		ethc.Address(common.HexToAddress(cfg.EigenFeeContractAddress)), feeParsed, l1Client, l1Client,
		l1Client,
	)

	driverConfig := &sequencer.DriverConfig{
		L1Client:                  l1Client,
		L2Client:                  l2Client,
		DtlClientUrl:              cfg.DtlClientUrl,
		EigenDaContract:           eigenContract,
		RawEigenContract:          rawEigenContract,
		EigenABI:                  eignenABI,
		EigenFeeContract:          eigenFeeContract,
		RawEigenFeeContract:       rawEigenFeeContract,
		EigenFeeABI:               eigenFeeABI,
		FeeModelEnable:            cfg.FeeModelEnable,
		FeeSizeSec:                cfg.FeeSizeSec,
		FeePerBytePerTime:         cfg.FeePerBytePerTime,
		Logger:                    logger,
		PrivKey:                   sequencerPrivKey,
		BlockOffset:               cfg.BlockOffset,
		RollUpMinSize:             cfg.RollUpMinSize,
		RollUpMaxSize:             cfg.RollUpMaxSize,
		EigenLayerNode:            cfg.EigenLayerNode,
		ChainID:                   chainID,
		DataStoreDuration:         uint64(cfg.DataStoreDuration),
		DataStoreTimeout:          cfg.DataStoreTimeout,
		DisperserSocket:           cfg.DisperserEndpoint,
		MainWorkerPollInterval:    cfg.MainWorkerPollInterval,
		CheckerWorkerPollInterval: cfg.CheckerWorkerPollInterval,
		DbPath:                    cfg.DbPath,
		CheckerBatchIndex:         cfg.CheckerBatchIndex,
		CheckerEnable:             cfg.CheckerEnable,
		GraphProvider:             cfg.GraphProvider,
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		SignerFn:                  signer(chainID),
	}
	driver, err := sequencer.NewDriver(ctx, driverConfig)
	if err != nil {
		log.Error("new driver fail", "err", err)
		return nil, err
	}
	daServiceConfig := &restorer.DaServiceConfig{
		EigenContract:   eigenContract,
		EigenABI:        eignenABI,
		RetrieverSocket: cfg.RetrieverSocket,
		GraphProvider:   cfg.GraphProvider,
		Timeout:         cfg.RetrieverTimeout,
		DaServicePort:   cfg.EigenDaHttpPort,
		EigenLayerNode:  cfg.EigenLayerNode,
		Debug:           cfg.EchoDebug,
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
			log.Error("metrics server failed to start", "err", err)
		}
	}()
	return nil
}

func (mb *MantleBatch) Stop() {
	mb.daService.Stop()
	mb.sequencerDriver.Stop()
}
