package oracle

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
)

func wrapUpdateDaFee(daBackend *bindings.BVMEigenDataLayrFee, l2Backend DeployContractBackend, cfg *Config) (func() error, error) {
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

		currentDaFee, err := contract.DaGasPrice(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		daFee, err := daBackend.GetRollupFee(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		if !isDifferenceSignificant(currentDaFee.Uint64(), daFee.Uint64(), cfg.daFeeSignificanceFactor) {
			log.Debug("non significant da fee update", "da", daFee, "current", currentDaFee)
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

		tx, err := contract.SetDAGasPrice(opts, daFee)
		if err != nil {
			return err
		}
		log.Debug("updating da fee", "tx.gasPrice", tx.GasPrice(), "tx.gasLimit", tx.Gas(),
			"tx.data", hexutil.Encode(tx.Data()), "tx.to", tx.To().Hex(), "tx.nonce", tx.Nonce())
		if err := l2Backend.SendTransaction(context.Background(), tx); err != nil {
			return fmt.Errorf("cannot update base fee: %w", err)
		}
		log.Info("L1 base fee transaction sent", "hash", tx.Hash().Hex(), "baseFee", daFee)

		if cfg.waitForReceipt {
			// Wait for the receipt
			receipt, err := waitForReceipt(l2Backend, tx)
			if err != nil {
				return err
			}

			log.Info("da-fee transaction confirmed", "hash", tx.Hash().Hex(),
				"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
		}
		return nil
	}, nil
}
