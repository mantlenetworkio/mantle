package signer

import (
	"context"
	"crypto/ecdsa"
	"github.com/bitdao-io/bitnetwork/bss-core/dial"
	l2ethclient "github.com/bitdao-io/bitnetwork/l2geth/ethclient"
	"github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib"
	"github.com/bitdao-io/bitnetwork/tss/node/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/client"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"sync"
)

type Processor struct {
	localPubkey               string
	privateKey                *ecdsa.PrivateKey
	tssServer                 tsslib.Server
	wsClient                  *client.WSClients
	l2Client                  *l2ethclient.Client
	l1Client                  *ethclient.Client
	ctx                       context.Context
	cancel                    func()
	srcChainId                string
	dstChainId                string
	dstChainName              string
	pause                     bool
	pauseWg                   *sync.WaitGroup
	stopChan                  chan struct{}
	wg                        *sync.WaitGroup
	askRequestChan            chan tdtypes.RPCRequest
	signRequestChan           chan tdtypes.RPCRequest
	askSlashChan              chan tdtypes.RPCRequest
	signSlashChan             chan tdtypes.RPCRequest
	keygenRequestChan         chan tdtypes.RPCRequest
	waitSignLock              *sync.Mutex
	waitSignMsgs              map[string]common.SignStateRequest
	waitSignSlashLock         *sync.Mutex
	waitSignSlashMsgs         map[string]map[uint64]common.SlashRequest
	nodeStore                 types.NodeStore
	logger                    zerolog.Logger
	tssGroupManagerAddress    string
	tssStakingSlashingAddress string
	metrics                   *Metrics
}

func NewProcessor(cfg common.Configuration, contx context.Context, tssInstance tsslib.Server, privKey *ecdsa.PrivateKey, pubKey string, nodeStore types.NodeStore) (*Processor, error) {
	ctx, cancel := context.WithCancel(contx)
	l1Cli, err := dial.L1EthClientWithTimeout(ctx, cfg.L1Url, cfg.Node.DisableHTTP2)
	if err != nil {
		return nil, err
	}
	wsClient, err := client.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubKey)
	if err != nil {
		return nil, err
	}
	l2Client, err := DialL2EthClientWithTimeout(ctx, cfg.Node.L2EthRpc, cfg.Node.DisableHTTP2)

	processor := Processor{
		localPubkey:               pubKey,
		privateKey:                privKey,
		tssServer:                 tssInstance,
		pauseWg:                   &sync.WaitGroup{},
		stopChan:                  make(chan struct{}),
		wg:                        &sync.WaitGroup{},
		logger:                    log.With().Str("module", "signer").Logger(),
		wsClient:                  wsClient,
		l2Client:                  l2Client,
		l1Client:                  l1Cli,
		ctx:                       ctx,
		cancel:                    cancel,
		askRequestChan:            make(chan tdtypes.RPCRequest, 100),
		signRequestChan:           make(chan tdtypes.RPCRequest, 100),
		askSlashChan:              make(chan tdtypes.RPCRequest, 1),
		signSlashChan:             make(chan tdtypes.RPCRequest, 1),
		keygenRequestChan:         make(chan tdtypes.RPCRequest, 1),
		waitSignLock:              &sync.Mutex{},
		waitSignMsgs:              make(map[string]common.SignStateRequest),
		nodeStore:                 nodeStore,
		tssGroupManagerAddress:    cfg.Node.TssGroupManagerAddress,
		tssStakingSlashingAddress: cfg.Node.TssStakingSlashingAddress,
		metrics:                   PrometheusMetrics("tssnode"),
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
	p.wg.Add(6)
	p.run()
}

func (p *Processor) Stop() {
	p.logger.Info().Msg("going to stop signer")
	defer p.logger.Info().Msg("signer stopped")
	close(p.stopChan)
	p.wsClient.Cli.Stop()
	p.cancel()
	p.l2Client.Close()
	p.l1Client.Close()
	p.wg.Wait()
}

func (p *Processor) run() {
	go p.ProcessMessage()
	go p.Verify()
	go p.VerifySlash()
	go p.Sign()
	go p.SignSlash()
	go p.Keygen()
}
