package service

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/bitdao-io/bitnetwork/bss-core/metrics"
	"github.com/bitdao-io/bitnetwork/bss-core/txmgr"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// weiToEth is the conversion rate from wei to ether.
var weiToEth = new(big.Float).SetFloat64(1e-18)

// Driver is an interface for creating and submitting batch transactions for a
// specific contract.
type Driver interface {
	// Name is an identifier used to prefix logs for a particular service.
	Name() string

	// Metrics returns the subservice telemetry object.
	Metrics() metrics.Metrics

	// ExecuteProcess executes a process on this driver.
	ExecuteProcess(ctx context.Context) error
}

type ServiceConfig struct {
	Context         context.Context
	Driver          Driver
	PollInterval    time.Duration
	L1Client        *ethclient.Client
	TxManagerConfig txmgr.Config
}

type Service struct {
	cfg    ServiceConfig
	ctx    context.Context
	cancel func()

	txMgr   txmgr.TxManager
	metrics metrics.Metrics

	wg sync.WaitGroup
}

func NewService(cfg ServiceConfig) *Service {
	ctx, cancel := context.WithCancel(cfg.Context)

	txMgr := txmgr.NewSimpleTxManager(
		cfg.Driver.Name(), cfg.TxManagerConfig, cfg.L1Client,
	)

	return &Service{
		cfg:     cfg,
		ctx:     ctx,
		cancel:  cancel,
		txMgr:   txMgr,
		metrics: cfg.Driver.Metrics(),
	}
}

func (s *Service) Start() error {
	s.wg.Add(1)
	go s.eventLoop()
	return nil
}

func (s *Service) Stop() error {
	s.cancel()
	s.wg.Wait()
	return nil
}

func (s *Service) eventLoop() {
	defer s.wg.Done()

	name := s.cfg.Driver.Name()

	ticker := time.NewTicker(s.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.cfg.Driver.ExecuteProcess(s.ctx)

		case err := <-s.ctx.Done():
			log.Error(name+" service shutting down", "err", err)
			return
		}
	}
}
