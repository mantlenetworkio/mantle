package payer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/subsidy/cache-file"
	"github.com/mantlenetworkio/mantle/subsidy/types"
)

type Payer struct {
	ctx                       context.Context
	config                    *Config
	queryClient               *ethclient.Client
	payClient                 *ethclient.Client
	payerStateFileWriter      *cache_file.PayerStateFileWriter
	l1QueryEpochLengthSeconds uint64
	waitForReceipt            bool
	stop                      chan struct{}
	sccTopic                  string
	sccAddrStr                common.Address
	ctcTopic                  string
	ctcAddrStr                common.Address
	receiveAddress            common.Address
}

func NewPayer(cfg *Config) *Payer {
	queryClient, err := ethclient.Dial(cfg.queryerHttpUrl)
	if err != nil {
		panic(err)
	}
	payClient, err := ethclient.Dial(cfg.payerHttpUrl)
	if err != nil {
		panic(err)
	}
	state := cache_file.NewPayerStateFileWriter(cfg.HomeDir, cfg.CacheDir, cfg.FileName)
	return &Payer{
		payClient:                 payClient,
		queryClient:               queryClient,
		config:                    cfg,
		l1QueryEpochLengthSeconds: cfg.l1QueryEpochLengthSeconds,
		ctx:                       context.Background(),
		payerStateFileWriter:      state,
		sccTopic:                  cfg.SCCTopic,
		sccAddrStr:                cfg.SCCAddress,
		ctcTopic:                  cfg.CTCTopic,
		ctcAddrStr:                cfg.CTCAddress,
		waitForReceipt:            cfg.waitForReceipt,
		stop:                      make(chan struct{}),
		receiveAddress:            cfg.receiverAddr,
	}
}

func (ob *Payer) getLogs(address common.Address, topic string, fromBlock, toBlock uint64) ([]ethtypes.Log, error) {
	filter := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{address},
		Topics:    [][]common.Hash{{ethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return ob.queryClient.FilterLogs(context.Background(), filter)
}

// PayRollupCost from block
func (ob *Payer) PayRollupCost() error {
	endBlock := ob.EndBlock()
	fromBlock := endBlock + 1
	tip, err := ob.queryClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	toBlock := tip.Number.Uint64()
	if fromBlock > toBlock {
		log.Info(fmt.Sprintf("to:%v less than from:%v,no new block\n", toBlock, fromBlock))
		return nil
	}
	if toBlock-fromBlock > 1000 {
		toBlock = fromBlock + 1000
	}
	totalFee := big.NewInt(0)
	sccLogs, err := ob.getLogs(ob.sccAddrStr, ob.sccTopic, fromBlock, toBlock)
	if err != nil {
		return err
	}
	ctcLogs, err := ob.getLogs(ob.ctcAddrStr, ob.ctcTopic, fromBlock, toBlock)
	if err != nil {
		return err
	}
	totalFee = totalFee.Add(ob.CalculateCost(sccLogs), ob.CalculateCost(ctcLogs))
	var hash string
	if totalFee.Cmp(big.NewInt(0)) == 0 {
		log.Info(fmt.Sprintf("block height form %v to %v totalFee is zero", fromBlock, toBlock))
	} else {
		hash, err = ob.Transfer(totalFee)
		if err != nil {
			return err
		}
	}
	log.Info(fmt.Sprintf("block height form %v to %v,amount:%v,transfer hash:%v", fromBlock, toBlock, totalFee, hash))
	payerState := types.PayerState{
		LastPayTime: time.Now(),
		EndBlock:    toBlock,
		PayTxHash:   hash,
	}
	if err := ob.payerStateFileWriter.Write(&payerState); err != nil {
		panic(err)
	}
	return nil
}

func (ob *Payer) CalculateCost(logs []ethtypes.Log) *big.Int {
	totalFee := big.NewInt(0)
	for _, l := range logs {
		tx, _, err := ob.queryClient.TransactionByHash(context.Background(), l.TxHash)
		if err != nil {
			return big.NewInt(0)
		}
		totalFee = totalFee.Add(totalFee, tx.Cost())
	}
	return totalFee
}

func (ob *Payer) EndBlock() uint64 {
	endBlock := ob.payerStateFileWriter.LoadCache().EndBlock
	if endBlock == 0 {
		endBlock = ob.config.StartBlock
	}
	return endBlock
}

// Start runs the Payer
func (ob *Payer) Start() error {
	ob.payLoop()
	return nil
}

func (ob *Payer) Stop() {
	close(ob.stop)
}

func (ob *Payer) Wait() {
	<-ob.stop
}

func (ob *Payer) Transfer(amount *big.Int) (string, error) {
	senderAddr := ob.receiveAddress
	nonce, err := ob.payClient.PendingNonceAt(context.Background(), senderAddr)
	if err != nil {
		log.Error("PendingNonceAt error:", err)
		return "", err
	}
	gasLimit := uint64(21000) // in units
	gasPrice, err := ob.payClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("SuggestGasPrice error:", err)
		return "", err
	}
	baseTx := &ethtypes.LegacyTx{
		To:       &ob.config.receiverAddr,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    amount,
		Data:     nil,
	}
	tx := ethtypes.NewTx(baseTx)
	chainID, err := ob.payClient.NetworkID(context.Background())
	if err != nil {
		log.Error("payClient.NetworkID error:", err)
		return "", err
	}
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewLondonSigner(chainID), ob.config.privateKey)
	if err != nil {
		log.Error("ethtypes.SignTx error:", err)
		return "", err
	}
	err = ob.payClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("SendTransaction error:", err)
		return "", err
	}

	log.Info("tx sent: %s", signedTx.Hash().Hex())
	if ob.waitForReceipt {
		// Wait for the receipt
		receipt, err := waitForReceipt(ob.payClient, tx)
		if err != nil {
			return signedTx.Hash().Hex(), err
		}
		log.Info("L1 transaction confirmed", "hash", tx.Hash().Hex(),
			"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
	}
	return signedTx.Hash().Hex(), nil
}

func (ob *Payer) payLoop() {
	timer := time.NewTicker(time.Duration(ob.l1QueryEpochLengthSeconds) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			if err := ob.PayRollupCost(); err != nil {
				log.Error("cannot pay rollup cost", "messgae", err)
			}

		case <-ob.ctx.Done():
			ob.Stop()
		}
	}
}

func waitForReceipt(backend *ethclient.Client, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	t := time.NewTicker(300 * time.Millisecond)
	receipt := new(ethtypes.Receipt)
	var err error
	for range t.C {

		receipt, err = backend.TransactionReceipt(context.Background(), tx.Hash())
		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt != nil {
			t.Stop()
			break
		}
	}
	return receipt, nil
}
