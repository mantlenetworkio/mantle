package signer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mantlenetworkio/mantle/bss-core/dial"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	"github.com/mantlenetworkio/mantle/tss/bindings/tsh"
	"github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/l1chain"
	managertypes "github.com/mantlenetworkio/mantle/tss/manager/types"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib"
	"github.com/mantlenetworkio/mantle/tss/node/types"
	"github.com/mantlenetworkio/mantle/tss/ws/client"
)

type Processor struct {
	localPubkey               string
	localPubKeyByte           []byte
	address                   ethc.Address
	privateKey                *ecdsa.PrivateKey
	chainId                   *big.Int
	tssServer                 tsslib.Server
	wsClient                  *client.WSClients
	l2Client                  *l2ethclient.Client
	l1Client                  *ethclient.Client
	ctx                       context.Context
	cancel                    func()
	stopChan                  chan struct{}
	wg                        *sync.WaitGroup
	askRequestChan            chan tdtypes.RPCRequest
	signRequestChan           chan tdtypes.RPCRequest
	askSlashChan              chan tdtypes.RPCRequest
	signSlashChan             chan tdtypes.RPCRequest
	keygenRequestChan         chan tdtypes.RPCRequest
	askRollBackChan           chan tdtypes.RPCRequest
	signRollBackChan          chan tdtypes.RPCRequest
	waitSignLock              *sync.RWMutex
	waitSignMsgs              map[string]common.SignStateRequest
	waitSignSlashLock         *sync.RWMutex
	waitSignSlashMsgs         map[string]map[uint64]common.SlashRequest
	cacheVerifyLock           *sync.RWMutex
	cacheVerify               *types.Cache[string, bool]
	cacheSignLock             *sync.RWMutex
	cacheSign                 *types.Cache[string, []byte]
	nodeStore                 types.NodeStore
	logger                    zerolog.Logger
	tssGroupManagerAddress    string
	tssStakingSlashingAddress string
	taskInterval              time.Duration
	tssStakingSlashingCaller  *tsh.TssStakingSlashingCaller
	tssQueryService           managertypes.TssQueryService
	l1ConfirmBlocks           int
	confirmReceiptTimeout     time.Duration
	metrics                   *Metrics
}

func NewProcessor(cfg common.Configuration, contx context.Context, tssInstance tsslib.Server, privKey *ecdsa.PrivateKey, pubkeyByte []byte, pubKeyHex string, nodeStore types.NodeStore, address ethc.Address) (*Processor, error) {
	taskIntervalDur, err := time.ParseDuration(cfg.TimedTaskInterval)
	if err != nil {
		return nil, err
	}
	receiptConfirmTimeoutDur, err := time.ParseDuration(cfg.L1ReceiptConfirmTimeout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(contx)
	l1Cli, err := dial.L1EthClientWithTimeout(ctx, cfg.L1Url, cfg.Node.DisableHTTP2)
	if err != nil {
		return nil, err
	}

	chainId, err := l1Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	wsClient, err := client.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubKeyHex)
	if err != nil {
		return nil, err
	}
	l2Client, err := DialL2EthClientWithTimeout(ctx, cfg.Node.L2EthRpc, cfg.Node.DisableHTTP2)
	tssStakingSlashingCaller, err := tsh.NewTssStakingSlashingCaller(ethc.HexToAddress(cfg.TssStakingSlashContractAddress), l1Cli)
	if err != nil {
		return nil, err
	}

	queryService := l1chain.NewQueryService(cfg.L1Url, cfg.TssGroupContractAddress, cfg.L1ConfirmBlocks, nodeStore)

	processor := Processor{
		localPubkey:               pubKeyHex,
		localPubKeyByte:           pubkeyByte,
		address:                   address,
		privateKey:                privKey,
		chainId:                   chainId,
		tssServer:                 tssInstance,
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
		askRollBackChan:           make(chan tdtypes.RPCRequest, 1),
		signRollBackChan:          make(chan tdtypes.RPCRequest, 1),
		waitSignLock:              &sync.RWMutex{},
		waitSignMsgs:              make(map[string]common.SignStateRequest),
		waitSignSlashLock:         &sync.RWMutex{},
		waitSignSlashMsgs:         make(map[string]map[uint64]common.SlashRequest),
		cacheVerifyLock:           &sync.RWMutex{},
		cacheVerify:               types.NewCache[string, bool](1000),
		cacheSignLock:             &sync.RWMutex{},
		cacheSign:                 types.NewCache[string, []byte](10),
		nodeStore:                 nodeStore,
		tssGroupManagerAddress:    cfg.TssGroupContractAddress,
		tssStakingSlashingAddress: cfg.TssStakingSlashContractAddress,
		taskInterval:              taskIntervalDur,
		tssStakingSlashingCaller:  tssStakingSlashingCaller,
		tssQueryService:           queryService,
		l1ConfirmBlocks:           cfg.L1ConfirmBlocks,
		confirmReceiptTimeout:     receiptConfirmTimeoutDur,
		metrics:                   PrometheusMetrics("tssnode"),
	}
	return &processor, nil
}

func (p *Processor) Start() {
	p.logger.Info().Msg("Signer is starting")
	p.wg.Add(9)
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
	go p.deleteSlashing()
	go p.SignRollBack()
	go p.VerifyRollBack()
}
