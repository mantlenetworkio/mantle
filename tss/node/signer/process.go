package signer

import (
	"context"
	"crypto/ecdsa"
	l2ethclient "github.com/bitdao-io/bitnetwork/l2geth/ethclient"
	"github.com/bitdao-io/bitnetwork/tss/node/config"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib"
	"github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"sync"
)

type Processor struct {
	localPubkey       string
	tssServer         tsslib.Server
	wsClient          *client.WSClients
	l2Client          *l2ethclient.Client
	ctx               context.Context
	cancel            func()
	srcChainId        string
	dstChainId        string
	dstChainName      string
	pause             bool
	pauseWg           *sync.WaitGroup
	stopChan          chan struct{}
	wg                *sync.WaitGroup
	askRequestChan    chan tdtypes.RPCRequest
	signRequestChan   chan tdtypes.RPCRequest
	keygenRequestChan chan tdtypes.RPCRequest
	waitSignLock      *sync.Mutex
	waitSignMsgs      map[string]types.SignStateRequest
	logger            zerolog.Logger

	metrics *Metrics
}

func NewProcessor(cfg config.Configuration, contx context.Context, tssInstance tsslib.Server, privKey *ecdsa.PrivateKey, pubKey string) (*Processor, error) {
	ctx, cancel := context.WithCancel(contx)

	wsClient, err := client.NewWSClient(cfg.WsAddr, "/ws", privKey, pubKey)
	if err != nil {
		return nil, err
	}
	l2Client, err := DialL2EthClientWithTimeout(ctx, cfg.BaseConfig.L2EthRpc, cfg.BaseConfig.DisableHTTP2)

	processor := Processor{
		localPubkey:       pubKey,
		tssServer:         tssInstance,
		pauseWg:           &sync.WaitGroup{},
		stopChan:          make(chan struct{}),
		wg:                &sync.WaitGroup{},
		logger:            log.With().Str("module", "signer").Logger(),
		wsClient:          wsClient,
		l2Client:          l2Client,
		ctx:               ctx,
		cancel:            cancel,
		askRequestChan:    make(chan tdtypes.RPCRequest, 100),
		signRequestChan:   make(chan tdtypes.RPCRequest, 100),
		keygenRequestChan: make(chan tdtypes.RPCRequest, 1),
		waitSignLock:      &sync.Mutex{},
		waitSignMsgs:      make(map[string]types.SignStateRequest),
		metrics:           PrometheusMetrics("tssnode"),
	}
	return &processor, nil
}

func (p *Processor) waitIfPause() {
	if p.pause {
		p.logger.Info().Msg("signing process is paused, waiting for the wake up signal")
		p.pauseWg.Wait()
		p.logger.Info().Msg("signing process is waked up, continue working......")
	}
}

func (p *Processor) Start() {
	p.logger.Info().Msg("Signer is starting")
	p.waitIfPause()
	p.wg.Add(4)
	p.run()
}

func (p *Processor) Stop() {
	p.logger.Info().Msg("going to stop signer")
	defer p.logger.Info().Msg("signer stopped")
	close(p.stopChan)
	p.wsClient.Cli.Stop()
	p.cancel()
	p.l2Client.Close()
	p.wg.Wait()
}

func (p *Processor) run() {
	go p.ProcessMessage()
	go p.Verify()
	go p.Sign()
	go p.Keygen()
}

func (p *Processor) PauseStatus() bool {
	return p.pause
}

func (p *Processor) Pause() {
	if p.pause {
		p.logger.Warn().Msg("tss node signing process is already paused ")
		return
	}
	p.pauseWg.Add(1)
	p.pause = true
	p.logger.Info().Msg("Paused tss node signing process ")
}

func (p *Processor) Wakeup() {
	if !p.pause {
		p.logger.Warn().Msg("tss node signing process is already wake up ")
		return
	}
	p.pause = false
	p.pauseWg.Done()
	p.logger.Info().Msg("Wake up tss node signing process ")
}
