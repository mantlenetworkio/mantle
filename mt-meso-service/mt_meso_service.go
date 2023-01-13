package mt_meso_service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-meso-service/meso-service/eigenda"
	middleware2 "github.com/mantlenetworkio/mantle/mt-meso-service/middleware"
	"github.com/urfave/cli"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}
		log.Info("Initializing mt meso service", "version", gitVersion)
		mesoService, err := NewMtMesoService(cfg)
		if err != nil {
			log.Error("Unable to create indexer", "error", err)
			return err
		}
		log.Info("Starting mt meso service")
		if err := mesoService.Start(); err != nil {
			return err
		}
		defer mesoService.Stop()
		log.Info("mt meso service started")

		<-(chan struct{})(nil)

		return nil
	}
}

type MtMesoService struct {
	ctx       context.Context
	cfg       Config
	ethClient *ethclient.Client
	daService *eigenda.DaService
}

func NewMtMesoService(cfg Config) (*MtMesoService, error) {
	ctx := context.Background()
	ethClient, err := l1l2client.L1EthClientWithTimeout(ctx, cfg.EthRpc, cfg.DisableHTTP2)
	if err != nil {
		return nil, err
	}
	eigenMiddleWareConfig := &middleware2.EigenDaMiddleWareConfig{
		EthClient:         ethClient,
		EigenContractAddr: common.HexToAddress(cfg.EigenContractAddress),
		RetrieverSocket:   cfg.RetrieverSocket,
		Timeout:           cfg.Timeout,
	}
	daConfig := &eigenda.DaServiceConfig{
		EigenMiddleWareConfig: eigenMiddleWareConfig,
		DaServicePort:         cfg.EigenDaHttpPort,
		Debug:                 cfg.EchoDebug,
	}
	daService, err := eigenda.NewDaService(ctx, daConfig)
	if err != nil {
		log.Error("new da service fail", "err", err)
	}
	return &MtMesoService{
		ctx:       ctx,
		cfg:       cfg,
		ethClient: ethClient,
		daService: daService,
	}, nil
}

func (mms *MtMesoService) Start() error {
	mms.daService.Run()
	log.Info("start service success")
	return nil
}

func (mms *MtMesoService) Stop() {
	log.Info("stop service success")
}
