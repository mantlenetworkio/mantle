package payer

import (
	"context"
	"errors"
	"fmt"
	"github.com/mantlenetworkio/mantle/subsidy/types"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/subsidy/cache-file"
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
	sccAddrStr                string
	ctcTopic                  string
	ctcAddrStr                string
}

func NewPayer(cfg *Config) *Payer {
	queryClient, err := ethclient.Dial(cfg.ethereumHttpUrl)
	if err != nil {
		panic(err)
	}
	payClient, err := ethclient.Dial(cfg.ethereumHttpUrl)
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
		sccAddrStr:                cfg.SCCAddress.Hex(),
		ctcTopic:                  cfg.SCCTopic,
		ctcAddrStr:                cfg.CTCAddress.Hex(),
		waitForReceipt:            cfg.waitForReceipt,
		stop:                      make(chan struct{}),
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
	if fromBlock <= toBlock {
		payerState := types.PayerState{
			LastPayTime: time.Now(),
			EndBlock:    toBlock,
			//PayTxHash:   hash,
		}
		if err := ob.payerStateFileWriter.Write(&payerState); err != nil {
			panic(err)
		}
		fmt.Println("fromBlock <= toBlock")
		fmt.Println("fromBlock:", fromBlock)
		fmt.Println("endBlock:", toBlock)
		return nil
	}
	sccAddress := common.HexToAddress(ob.sccAddrStr)
	sccLogs, err := ob.getLogs(sccAddress, ob.sccTopic, fromBlock, toBlock)
	if err != nil {
		return err
	}
	totalCost := big.NewInt(0)
	for _, l := range sccLogs {
		tx, err := ob.queryClient.TransactionInBlock(context.Background(), l.BlockHash, l.TxIndex)
		if err != nil {
			return err
		}
		totalCost = totalCost.Add(totalCost, tx.Cost())
	}
	ctcAddress := common.HexToAddress(ob.sccAddrStr)
	ctcLogs, err := ob.getLogs(ctcAddress, ob.ctcTopic, fromBlock, toBlock)
	for _, l := range ctcLogs {
		tx, err := ob.queryClient.TransactionInBlock(context.Background(), l.BlockHash, l.TxIndex)
		if err != nil {
			return err
		}
		totalCost = totalCost.Add(totalCost, tx.Cost())
	}
	fmt.Println("cost", totalCost)
	if totalCost.Cmp(big.NewInt(0)) == 1 {
		fmt.Println("total cost = 0 ")
		return nil
	}
	hash, err := ob.Transfer(totalCost)
	if err != nil {
		return err
	}
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

func (ob *Payer) EndBlock() uint64 {
	endBlock := ob.payerStateFileWriter.LoadCache().EndBlock
	if endBlock == 0 {
		endBlock = ob.config.StartBlock
	}
	return ob.payerStateFileWriter.LoadCache().EndBlock
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
	senderAddr := ethcrypto.PubkeyToAddress(ob.config.privateKey.PublicKey)
	nonce, err := ob.payClient.PendingNonceAt(context.Background(), senderAddr)
	if err != nil {
		return "", err
	}
	gasLimit := uint64(21000) // in units
	gasPrice, err := ob.payClient.SuggestGasPrice(context.Background())
	if err != nil {
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
		return "", err
	}

	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), ob.config.privateKey)
	if err != nil {
		return "", err
	}

	err = ob.payClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
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
