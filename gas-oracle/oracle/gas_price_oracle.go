package oracle

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
	"github.com/mantlenetworkio/mantle/gas-oracle/gasprices"
	ometrics "github.com/mantlenetworkio/mantle/gas-oracle/metrics"
	"github.com/mantlenetworkio/mantle/gas-oracle/tokenprice"
	"github.com/mantlenetworkio/mantle/l2geth/core/rawdb"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
)

const GAS_ORACLE_SYNC_HEIGHT = "GAS_ORACLE_SYNC_HEIGHT"

var (
	// errInvalidSigningKey represents the error when the signing key used
	// is not the Owner of the contract and therefore cannot update the gasprice
	errInvalidSigningKey = errors.New("invalid signing key")
	// errNoChainID represents the error when the chain id is not provided
	// and it cannot be remotely fetched
	errNoChainID = errors.New("no chain id provided")
	// errNoPrivateKey represents the error when the private key is not provided to
	// the application
	errNoPrivateKey = errors.New("no private key provided")
	// errWrongChainID represents the error when the configured chain id is not
	// correct
	errWrongChainID = errors.New("wrong chain id provided")
	// errNoBaseFee represents the error when the base fee is not found on the
	// block. This means that the block being queried is pre eip1559
	errNoBaseFee = errors.New("base fee not found on block")
)

// GasPriceOracle manages a hot key that can update the L2 Gas Price
type GasPriceOracle struct {
	l1ChainID       *big.Int
	l2ChainID       *big.Int
	ctx             context.Context
	stop            chan struct{}
	contract        *bindings.BVMGasPriceOracle
	l2Backend       DeployContractBackend
	l1Backend       bind.ContractTransactor
	daBackend       *bindings.BVMEigenDataLayrFee
	sccBackend      *bindings.StateCommitmentChain
	ctcBackend      *bindings.CanonicalTransactionChain
	gasPriceUpdater *gasprices.GasPriceUpdater
	config          *Config
}

// Start runs the GasPriceOracle
func (g *GasPriceOracle) Start() error {
	if g.config.l1ChainID == nil {
		return fmt.Errorf("layer-one: %w", errNoChainID)
	}
	if g.config.l2ChainID == nil {
		return fmt.Errorf("layer-two: %w", errNoChainID)
	}
	var address common.Address
	if !g.config.EnableHsm {
		if g.config.privateKey == nil {
			return errNoPrivateKey
		}
		address = crypto.PubkeyToAddress(g.config.privateKey.PublicKey)
	} else {
		address = common.HexToAddress(g.config.HsmAddress)
	}

	log.Info("Starting Gas Price Oracle", "l1-chain-id", g.l1ChainID,
		"l2-chain-id", g.l2ChainID, "address", address.Hex())

	price, err := g.contract.GasPrice(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return err
	}
	ometrics.GasOracleStats.L2GasPriceGauge.Update(int64(price.Uint64()))

	log.Info("Starting Gas Price Oracle enableL1BaseFee", "enableL1BaseFee",
		g.config.enableL1BaseFee, "enableL2GasPrice", g.config.enableL2GasPrice, "enableDaFee", g.config.enableDaFee)

	if g.config.enableL1BaseFee {
		go g.BaseFeeLoop()
	}
	if g.config.enableL1Overhead {
		go g.OverHeadLoop()
	}
	if g.config.enableDaFee {
		go g.DaFeeLoop()
	}
	if g.config.enableL2GasPrice {
		go g.Loop()
	}

	return nil
}

func (g *GasPriceOracle) Stop() {
	close(g.stop)
}

func (g *GasPriceOracle) Wait() {
	<-g.stop
}

// ensure makes sure that the configured private key is the owner
// of the `BVM_GasPriceOracle`. If it is not the owner, then it will
// not be able to make updates to the L2 gas price.
func (g *GasPriceOracle) ensure() error {
	owner, err := g.contract.Owner(&bind.CallOpts{
		Context: g.ctx,
	})
	if err != nil {
		return err
	}
	var address common.Address
	if g.config.EnableHsm {
		address = common.HexToAddress(g.config.HsmAddress)
	} else {
		address = crypto.PubkeyToAddress(g.config.privateKey.PublicKey)
	}
	if address != owner {
		log.Error("Signing key does not match contract owner", "signer", address.Hex(), "owner", owner.Hex())
		return errInvalidSigningKey
	}
	return nil
}

// Loop is the main logic of the gas-oracle
func (g *GasPriceOracle) Loop() {
	timer := time.NewTicker(time.Duration(g.config.epochLengthSeconds) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			log.Trace("polling", "time", time.Now())
			if err := g.Update(); err != nil {
				log.Error("cannot update gas price", "message", err)
			}

		case <-g.ctx.Done():
			g.Stop()
		}
	}
}

func (g *GasPriceOracle) BaseFeeLoop() {
	timer := time.NewTicker(time.Duration(g.config.l1BaseFeeEpochLengthSeconds) * time.Second)
	defer timer.Stop()

	updateBaseFee, err := wrapUpdateBaseFee(g.l1Backend, g.l2Backend, g.config)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-timer.C:
			if err := updateBaseFee(); err != nil {
				log.Error("cannot update l1 base fee", "message", err)
			}
		case <-g.ctx.Done():
			g.Stop()
		}
	}
}

func (g *GasPriceOracle) DaFeeLoop() {
	timer := time.NewTicker(time.Duration(g.config.daFeeEpochLengthSeconds) * time.Second)
	defer timer.Stop()

	updateDaFee, err := wrapUpdateDaFee(g.daBackend, g.l2Backend, g.config)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-timer.C:
			if err := updateDaFee(); err != nil {
				log.Error("cannot update da fee", "messgae", err)
			}

		case <-g.ctx.Done():
			g.Stop()
		}
	}
}

func (g *GasPriceOracle) OverHeadLoop() {
	// set ticker
	ticker := time.NewTicker(5 * time.Second)

	// read gas-oracle synced height
	db, height := readGasOracleSyncHeight()
	log.Info("ReadGasOracleSyncHeight", "height", height)
	// set channel
	stateBatchAppendChan := make(chan *bindings.StateCommitmentChainStateBatchAppended, 1)

	// set context
	ctcTotalBatches, err := g.ctcBackend.GetTotalBatches(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	updateOverhead, err := wrapUpdateOverhead(g.l2Backend, g.config)
	if err != nil {
		panic(err)
	}

	var end uint64
	for {
		select {
		case <-ticker.C:
			log.Info("OverHeadLoop is living, HeartBeat It!")
			latestHeader, err := g.l1Backend.HeaderByNumber(g.ctx, nil)
			if err != nil {
				log.Warn("get latest header in error", "error", err)
				continue
			}
			// repeat query latest block is not allowed
			if height != nil && height.Uint64() != 0 && height.Uint64() >= latestHeader.Number.Uint64() {
				continue
			}
			if height == nil || height.Uint64() == 0 {
				height = latestHeader.Number
			}
			end = latestHeader.Number.Uint64()

			iter, err := g.sccBackend.FilterStateBatchAppended(&bind.FilterOpts{
				Start:   height.Uint64(),
				End:     &end,
				Context: g.ctx,
			}, nil)
			for iter.Next() {
				select {
				case stateBatchAppendChan <- iter.Event:
					log.Info("write event into channel", "channel length is", len(stateBatchAppendChan))
				default:
					// discard additional event to process the last one
					<-stateBatchAppendChan
					stateBatchAppendChan <- iter.Event
					log.Warn("write additional event into channel, will ignore prev events")
				}
			}
			_ = writeGasOracleSyncHeight(db, latestHeader.Number)
			height = latestHeader.Number.Add(latestHeader.Number, big.NewInt(1))
			log.Info("Update synced height", "height", height)
		case ev := <-stateBatchAppendChan:
			currentCtcBatches, err := g.ctcBackend.GetTotalBatches(&bind.CallOpts{})
			if err != nil {
				continue
			}
			log.Info("current scc batch size", "size", ev.BatchSize)
			log.Info("CTC circle num in SCC circle", "count", new(big.Int).Sub(currentCtcBatches, ctcTotalBatches))
			if err := updateOverhead(new(big.Int).Sub(currentCtcBatches, ctcTotalBatches), ev.BatchSize); err != nil {
				log.Error("cannot update overhead", "message", err)
			}
			ctcTotalBatches = currentCtcBatches
		case <-g.ctx.Done():
			db.Close()
			g.Stop()
		}
	}
}

// Update will update the gas price
func (g *GasPriceOracle) Update() error {
	l2GasPrice, err := g.contract.GasPrice(&bind.CallOpts{
		Context: g.ctx,
	})
	if err != nil {
		return fmt.Errorf("cannot get gas price: %w", err)
	}

	if err := g.gasPriceUpdater.UpdateGasPrice(); err != nil {
		return fmt.Errorf("cannot update gas price: %w", err)
	}

	newGasPrice, err := g.contract.GasPrice(&bind.CallOpts{
		Context: g.ctx,
	})
	if err != nil {
		return fmt.Errorf("cannot get gas price: %w", err)
	}

	local := g.gasPriceUpdater.GetGasPrice()
	log.Info("Update", "original", l2GasPrice, "current", newGasPrice, "local", local)
	return nil
}

// NewGasPriceOracle creates a new GasPriceOracle based on a Config
func NewGasPriceOracle(cfg *Config) (*GasPriceOracle, error) {
	tokenPricer := tokenprice.NewClient(cfg.PriceBackendURL, cfg.PriceBackendUniswapURL,
		cfg.tokenPricerUpdateFrequencySecond, cfg.tokenRatioMode, cfg.tokenPairMNTMode)
	if tokenPricer == nil {
		return nil, fmt.Errorf("invalid token price client")
	}
	// Create the L2 client
	l2Client, err := ethclient.Dial(cfg.layerTwoHttpUrl)
	if err != nil {
		return nil, err
	}

	l1Client, err := NewL1Client(cfg.ethereumHttpUrl, tokenPricer)
	if err != nil {
		return nil, err
	}
	daFeeClient, err := bindings.NewBVMEigenDataLayrFee(cfg.daFeeContractAddress, l1Client.Client)
	if err != nil {
		return nil, err
	}
	sccBackend, err := bindings.NewStateCommitmentChain(cfg.sccContractAddress, l1Client.Client)
	if err != nil {
		return nil, err
	}
	ctcBackend, err := bindings.NewCanonicalTransactionChain(cfg.ctcContractAddress, l1Client.Client)
	if err != nil {
		return nil, err
	}
	// Ensure that we can actually connect to both backends
	log.Info("Connecting to layer two")
	if err := ensureConnection(l2Client); err != nil {
		log.Error("Unable to connect to layer two")
		return nil, err
	}
	log.Info("Connecting to layer one")
	if err := ensureConnection(l1Client.Client); err != nil {
		log.Error("Unable to connect to layer one")
		return nil, err
	}

	address := cfg.gasPriceOracleAddress
	contract, err := bindings.NewBVMGasPriceOracle(address, l2Client)
	if err != nil {
		return nil, err
	}

	// Fetch the current gas price to use as the current price
	currentPrice, err := contract.GasPrice(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return nil, err
	}

	// Create a gas pricer for the gas price updater
	log.Info("Creating GasPricer", "currentPrice", currentPrice,
		"floorPrice", cfg.floorPrice, "targetGasPerSecond", cfg.targetGasPerSecond,
		"maxPercentChangePerEpoch", cfg.maxPercentChangePerEpoch)

	gasPricer, err := gasprices.NewGasPricer(
		currentPrice.Uint64(),
		cfg.floorPrice,
		tokenPricer,
		func() float64 {
			return float64(cfg.targetGasPerSecond)
		},
		cfg.maxPercentChangePerEpoch,
	)
	if err != nil {
		return nil, err
	}

	l2ChainID, err := l2Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	l1ChainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	if cfg.l2ChainID != nil {
		if cfg.l2ChainID.Cmp(l2ChainID) != 0 {
			return nil, fmt.Errorf("%w: L2: configured with %d and got %d",
				errWrongChainID, cfg.l2ChainID, l2ChainID)
		}
	} else {
		cfg.l2ChainID = l2ChainID
	}

	if cfg.l1ChainID != nil {
		if cfg.l1ChainID.Cmp(l1ChainID) != 0 {
			return nil, fmt.Errorf("%w: L1: configured with %d and got %d",
				errWrongChainID, cfg.l1ChainID, l1ChainID)
		}
	} else {
		cfg.l1ChainID = l1ChainID
	}

	if !cfg.EnableHsm && cfg.privateKey == nil {
		return nil, errNoPrivateKey
	}

	tip, err := l2Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Start at the tip
	epochStartBlockNumber := tip.Number.Uint64()
	// getLatestBlockNumberFn is used by the GasPriceUpdater
	// to get the latest block number
	getLatestBlockNumberFn := wrapGetLatestBlockNumberFn(l2Client)
	// updateL2GasPriceFn is used by the GasPriceUpdater to
	// update the gas price
	updateL2GasPriceFn, err := wrapUpdateL2GasPriceFn(l2Client, cfg)
	if err != nil {
		return nil, err
	}
	// getGasUsedByBlockFn is used by the GasPriceUpdater
	// to fetch the amount of gas that a block has used
	getGasUsedByBlockFn := wrapGetGasUsedByBlock(l2Client)

	log.Info("Creating GasPriceUpdater", "epochStartBlockNumber", epochStartBlockNumber,
		"averageBlockGasLimitPerEpoch", cfg.averageBlockGasLimitPerEpoch,
		"epochLengthSeconds", cfg.epochLengthSeconds)

	gasPriceUpdater, err := gasprices.NewGasPriceUpdater(
		gasPricer,
		epochStartBlockNumber,
		cfg.averageBlockGasLimitPerEpoch,
		cfg.epochLengthSeconds,
		getLatestBlockNumberFn,
		getGasUsedByBlockFn,
		updateL2GasPriceFn,
	)

	if err != nil {
		return nil, err
	}

	gpo := GasPriceOracle{
		l2ChainID:       l2ChainID,
		l1ChainID:       l1ChainID,
		ctx:             context.Background(),
		stop:            make(chan struct{}),
		contract:        contract,
		gasPriceUpdater: gasPriceUpdater,
		config:          cfg,
		l2Backend:       l2Client,
		l1Backend:       l1Client,
		daBackend:       daFeeClient,
		sccBackend:      sccBackend,
		ctcBackend:      ctcBackend,
	}

	if err := gpo.ensure(); err != nil {
		return nil, err
	}

	return &gpo, nil
}

// Ensure that we can actually connect
func ensureConnection(client *ethclient.Client) error {
	t := time.NewTicker(1 * time.Second)
	retries := 0
	defer t.Stop()
	for ; true; <-t.C {
		_, err := client.ChainID(context.Background())
		if err == nil {
			break
		} else {
			retries += 1
			if retries > 90 {
				return err
			}
		}
	}
	return nil
}

func readGasOracleSyncHeight() (ethdb.Database, *big.Int) {
	// read synced height
	db, err := rawdb.NewLevelDBDatabase("gas-oracle-data", 0, 0, "")
	if err != nil {
		log.Error("NewLevelDBDatabase in error", "err", err)
		panic(err)
	}
	// will close connection at once
	//defer db.Close()

	has, err := db.Has([]byte(GAS_ORACLE_SYNC_HEIGHT))
	if err != nil {
		log.Error("check db has GAS_ORACLE_SYNC_HEIGHT in error", "err", err)
		panic(err)
	}
	if !has {
		return db, nil
	}

	height, err := db.Get([]byte(GAS_ORACLE_SYNC_HEIGHT))
	if err != nil {
		log.Error("check db Get GAS_ORACLE_SYNC_HEIGHT in error", "err", err)
		panic(err)
	}
	return db, big.NewInt(0).SetUint64(binary.BigEndian.Uint64(height))
}

func writeGasOracleSyncHeight(db ethdb.Database, height *big.Int) error {
	// will close connection at once
	//defer db.Close()

	var indexBz = make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, height.Uint64())
	err := db.Put([]byte(GAS_ORACLE_SYNC_HEIGHT), indexBz)
	if err != nil {
		log.Error("put GAS_ORACLE_SYNC_HEIGHT in error", "err", err)
		return err
	}
	return nil
}
