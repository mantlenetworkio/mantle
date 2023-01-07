package txmgr

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type UpdateGasPriceFunc = func(ctx context.Context) (*types.Transaction, error)

type SendTransactionFunc = func(ctx context.Context, tx *types.Transaction) error

type Config struct {
	ResubmissionTimeout       time.Duration
	ReceiptQueryInterval      time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
}

type TxManager interface {
	Send(ctx context.Context, updateGasPrice UpdateGasPriceFunc, sendTxn SendTransactionFunc) (*types.Receipt, error)
}

type ReceiptSource interface {
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type SimpleTxManager struct {
	cfg     Config
	backend ReceiptSource
	l       log.Logger
}

func NewSimpleTxManager(cfg Config, backend ReceiptSource) *SimpleTxManager {
	if cfg.NumConfirmations == 0 {
		panic("txmgr: NumConfirmations cannot be zero")
	}
	return &SimpleTxManager{
		cfg:     cfg,
		backend: backend,
	}
}

func (m *SimpleTxManager) Send(ctx context.Context, updateGasPrice UpdateGasPriceFunc, sendTx SendTransactionFunc) (*types.Receipt, error) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ctxc, cancel := context.WithCancel(ctx)
	defer cancel()

	sendState := NewSendState(m.cfg.SafeAbortNonceTooLowCount)

	receiptChan := make(chan *types.Receipt, 1)
	sendTxAsync := func() {
		defer wg.Done()

		tx, err := updateGasPrice(ctxc)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return
			}
			log.Error("MtBatcher update txn gas price fail", "err", err)
			return
		}

		txHash := tx.Hash()
		nonce := tx.Nonce()
		gasTipCap := tx.GasTipCap()
		gasFeeCap := tx.GasFeeCap()
		log.Info("MtBatcher publishing transaction", "txHash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap)

		err = sendTx(ctxc, tx)
		sendState.ProcessSendError(err)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return
			}
			log.Error("MtBatcher unable to publish transaction", "err", err)
			if sendState.ShouldAbortImmediately() {
				cancel()
			}
			return
		}

		log.Info("MtBatcher transaction published successfully", "hash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap)

		receipt, err := waitMined(
			ctxc, m.backend, tx, m.cfg.ReceiptQueryInterval,
			m.cfg.NumConfirmations, sendState,
		)
		if err != nil {
			log.Debug("MtBatcher send tx failed", "hash", txHash, "nonce", nonce, "gasTipCap", gasTipCap, "gasFeeCap", gasFeeCap, "err", err)
		}
		if receipt != nil {
			select {
			case receiptChan <- receipt:
				log.Trace("MtBatcher send tx succeeded", "hash", txHash,
					"nonce", nonce, "gasTipCap", gasTipCap,
					"gasFeeCap", gasFeeCap)
			default:
			}
		}
	}

	wg.Add(1)
	go sendTxAsync()

	ticker := time.NewTicker(m.cfg.ResubmissionTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if sendState.IsWaitingForConfirmation() {
				continue
			}
			wg.Add(1)
			go sendTxAsync()

		case <-ctxc.Done():
			return nil, ctxc.Err()

		case receipt := <-receiptChan:
			return receipt, nil
		}
	}
}

func WaitMined(
	ctx context.Context,
	backend ReceiptSource,
	tx *types.Transaction,
	queryInterval time.Duration,
	numConfirmations uint64,
) (*types.Receipt, error) {
	return waitMined(ctx, backend, tx, queryInterval, numConfirmations, nil)
}

func waitMined(
	ctx context.Context,
	backend ReceiptSource,
	tx *types.Transaction,
	queryInterval time.Duration,
	numConfirmations uint64,
	sendState *SendState,
) (*types.Receipt, error) {

	queryTicker := time.NewTicker(queryInterval)
	defer queryTicker.Stop()

	txHash := tx.Hash()

	for {
		receipt, err := backend.TransactionReceipt(ctx, txHash)
		switch {
		case receipt != nil:
			if sendState != nil {
				sendState.TxMined(txHash)
			}

			txHeight := receipt.BlockNumber.Uint64()
			tipHeight, err := backend.BlockNumber(ctx)
			if err != nil {
				log.Error("MtBatcher Unable to fetch block number", "err", err)
				break
			}

			log.Trace("MtBatcher Transaction mined, checking confirmations",
				"txHash", txHash, "txHeight", txHeight,
				"tipHeight", tipHeight,
				"numConfirmations", numConfirmations)

			if txHeight+numConfirmations <= tipHeight+1 {
				log.Info("MtBatcher Transaction confirmed", "txHash", txHash)
				return receipt, nil
			}

			confsRemaining := (txHeight + numConfirmations) - (tipHeight + 1)
			log.Info("MtBatcher Transaction not yet confirmed", "txHash", txHash,
				"confsRemaining", confsRemaining)

		case err != nil:
			log.Trace("MtBatcher Receipt retrievel failed", "hash", txHash,
				"err", err)

		default:
			if sendState != nil {
				sendState.TxNotMined(txHash)
			}
			log.Trace("MtBatcher Transaction not yet mined", "hash", txHash)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func CalcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}
