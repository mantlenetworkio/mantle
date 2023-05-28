package oracle

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
)

func wrapUpdateBaseFee(l1Backend bind.ContractTransactor, l2Backend DeployContractBackend, cfg *Config) (func() error, error) {
	if cfg.privateKey == nil {
		return nil, errNoPrivateKey
	}
	if cfg.l2ChainID == nil {
		return nil, errNoChainID
	}

	opts, err := bind.NewKeyedTransactorWithChainID(cfg.privateKey, cfg.l2ChainID)
	if err != nil {
		return nil, err
	}
	// Once https://github.com/ethereum/go-ethereum/pull/23062 is released
	// then we can remove setting the context here
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	// Don't send the transaction using the `contract` so that we can inspect
	// it beforehand
	opts.NoSend = true

	// Create a new contract bindings in scope of the updateL2GasPriceFn
	// that is returned from this function
	contract, err := bindings.NewBVMGasPriceOracle(cfg.gasPriceOracleAddress, l2Backend)
	if err != nil {
		return nil, err
	}
	return func() error {
		baseFee, err := contract.L1BaseFee(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		tip, err := l1Backend.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err
		}
		if tip.BaseFee == nil {
			return errNoBaseFee
		}
		if !isDifferenceSignificant(baseFee.Uint64(), tip.BaseFee.Uint64(), cfg.l1BaseFeeSignificanceFactor) {
			log.Debug("non significant base fee update", "tip", tip.BaseFee, "current", baseFee)
			return nil
		}

		// Use the configured gas price if it is set,
		// otherwise use gas estimation
		if cfg.gasPrice != nil {
			opts.GasPrice = cfg.gasPrice
		} else {
			gasPrice, err := l2Backend.SuggestGasPrice(opts.Context)
			if err != nil {
				return err
			}
			opts.GasPrice = gasPrice
		}
		gasTipCap, err := l1Backend.SuggestGasTipCap(opts.Context)
		if err != nil {
			return err
		}
		// get history 20 block best gasprice
		bestBaseFee := getHistoryBestPrice(l1Backend, tip.Number, tip.BaseFee, 20)
		// set L1BaseFee to base fee + tip cap, to cover rollup tip cap
		tx, err := contract.SetL1BaseFee(opts, new(big.Int).Add(bestBaseFee, gasTipCap))
		if err != nil {
			return err
		}
		log.Debug("updating L1 base fee", "tx.gasPrice", tx.GasPrice(), "tx.gasLimit", tx.Gas(),
			"tx.data", hexutil.Encode(tx.Data()), "tx.to", tx.To().Hex(), "tx.nonce", tx.Nonce())
		if err := l2Backend.SendTransaction(context.Background(), tx); err != nil {
			return fmt.Errorf("cannot update base fee: %w", err)
		}
		log.Info("L1 base fee transaction sent", "hash", tx.Hash().Hex(), "baseFee", tip.BaseFee)

		if cfg.waitForReceipt {
			// Wait for the receipt
			receipt, err := waitForReceipt(l2Backend, tx)
			if err != nil {
				return err
			}

			log.Info("base-fee transaction confirmed", "hash", tx.Hash().Hex(),
				"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
		}
		return nil
	}, nil
}

func getHistoryBestPrice(l1Backend bind.ContractTransactor, endHeight *big.Int, lastBaseFee *big.Int, countWindow int) *big.Int {
	var baseFees = make([]*big.Int, 0)
	var bestPrice = new(big.Int)
	var wg = sync.WaitGroup{}
	// get base fee
	for i := 0; i < countWindow; i++ {
		wg.Add(1)
		go func() {
			header, err := l1Backend.HeaderByNumber(context.Background(), endHeight.Sub(endHeight, new(big.Int).SetInt64(int64(i))))
			if err == nil && header.BaseFee != nil {
				baseFees = append(baseFees, header.BaseFee)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	// get best base fee, append last base fee again, incase get base fees all in error
	baseFees = append(baseFees, lastBaseFee)
	for j := 0; j < len(baseFees); j++ {
		if bestPrice.Cmp(baseFees[j]) < 0 {
			bestPrice = baseFees[j]
		}
	}
	return bestPrice
}
