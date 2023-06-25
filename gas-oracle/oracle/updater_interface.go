package oracle

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
	ometrics "github.com/mantlenetworkio/mantle/gas-oracle/metrics"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"
)

// getLatestBlockNumberFn is used by the GasPriceUpdater
// to get the latest block number. The outer function binds the
// inner function to a `bind.ContractBackend` which is implemented
// by the `ethclient.Client`
func wrapGetLatestBlockNumberFn(backend bind.ContractBackend) func() (uint64, error) {
	return func() (uint64, error) {
		tip, err := backend.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return 0, err
		}
		return tip.Number.Uint64(), nil
	}
}

// wrapGetGasUsedByBlock is used by the GasPriceUpdater to get
// the amount of gas used by a particular block. This is used to
// track gas usage over time
func wrapGetGasUsedByBlock(backend bind.ContractBackend) func(*big.Int) (uint64, error) {
	return func(number *big.Int) (uint64, error) {
		block, err := backend.HeaderByNumber(context.Background(), number)
		if err != nil {
			return 0, err
		}
		return block.GasUsed, nil
	}
}

// DeployContractBackend represents the union of the
// DeployBackend and the ContractBackend
type DeployContractBackend interface {
	bind.DeployBackend
	bind.ContractBackend
}

// updateL2GasPriceFn is used by the GasPriceUpdater
// to update the L2 gas price
// perhaps this should take an options struct along with the backend?
// how can this continue to be decomposed?
func wrapUpdateL2GasPriceFn(backend DeployContractBackend, cfg *Config) (func(uint64) error, error) {
	var opts *bind.TransactOpts
	var err error
	if !cfg.EnableHsm {
		if cfg.privateKey == nil {
			return nil, errNoPrivateKey
		}
		if cfg.l2ChainID == nil {
			return nil, errNoChainID
		}

		opts, err = bind.NewKeyedTransactorWithChainID(cfg.privateKey, cfg.l2ChainID)
		if err != nil {
			return nil, err
		}
	} else {
		seqBytes, err := hex.DecodeString(cfg.HsmCreden)
		apikey := option.WithCredentialsJSON(seqBytes)
		client, err := kms.NewKeyManagementClient(context.Background(), apikey)
		if err != nil {
			log.Crit("gasoracle", "create signer error", err.Error())
		}
		mk := &bsscore.ManagedKey{
			KeyName:      cfg.HsmAPIName,
			EthereumAddr: common.HexToAddress(cfg.HsmAddress),
			Gclient:      client,
		}
		opts, err = mk.NewEthereumTransactorrWithChainID(context.Background(), cfg.l2ChainID)
		if err != nil {
			log.Crit("gasoracle", "create signer error", err.Error())
			return nil, err
		}
	}

	// Don't send the transaction using the `contract` so that we can inspect
	// it beforehand
	opts.NoSend = true

	// Create a new contract bindings in scope of the updateL2GasPriceFn
	// that is returned from this function
	contract, err := bindings.NewBVMGasPriceOracle(cfg.gasPriceOracleAddress, backend)
	if err != nil {
		return nil, err
	}

	return func(updatedGasPrice uint64) error {
		log.Trace("UpdateL2GasPriceFn", "gas-price", updatedGasPrice)
		if cfg.gasPrice == nil {
			// Set the gas price manually to use legacy transactions
			gasPrice, err := backend.SuggestGasPrice(context.Background())
			if err != nil {
				log.Error("cannot fetch gas price", "message", err)
				return err
			}
			log.Trace("fetched L2 tx.gasPrice", "gas-price", gasPrice)
			opts.GasPrice = gasPrice
		} else {
			// Allow a configurable gas price to be set
			opts.GasPrice = cfg.gasPrice
		}

		// Query the current L2 gas price
		currentPrice, err := contract.GasPrice(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			log.Error("cannot fetch current gas price", "message", err)
			return err
		}

		// no need to update when they are the same
		if currentPrice.Uint64() == updatedGasPrice {
			log.Info("gas price did not change", "gas-price", updatedGasPrice)
			ometrics.GasOracleStats.TxNotSignificantCounter.Inc(1)
			return nil
		}

		// Only update the gas price when it must be changed by at least
		// a paramaterizable amount.
		if !isDifferenceSignificant(currentPrice.Uint64(), updatedGasPrice, cfg.l2GasPriceSignificanceFactor) {
			log.Info("gas price did not significantly change", "min-factor", cfg.l2GasPriceSignificanceFactor,
				"current-price", currentPrice, "next-price", updatedGasPrice)
			ometrics.GasOracleStats.TxNotSignificantCounter.Inc(1)
			return nil
		}

		// Set the gas price by sending a transaction
		tx, err := contract.SetGasPrice(opts, new(big.Int).SetUint64(updatedGasPrice))
		if err != nil {
			return err
		}

		log.Debug("updating L2 gas price", "tx.gasPrice", tx.GasPrice(), "tx.gasLimit", tx.Gas(),
			"tx.data", hexutil.Encode(tx.Data()), "tx.to", tx.To().Hex(), "tx.nonce", tx.Nonce())
		pre := time.Now()
		if err := backend.SendTransaction(context.Background(), tx); err != nil {
			return err
		}
		ometrics.GasOracleStats.TxSendTimer.Update(time.Since(pre))
		log.Info("L2 gas price transaction sent", "hash", tx.Hash().Hex())

		ometrics.GasOracleStats.L2GasPriceGauge.Update(int64(updatedGasPrice))
		ometrics.GasOracleStats.TxSendCounter.Inc(1)

		if cfg.waitForReceipt {
			// Keep track of the time it takes to confirm the transaction
			pre := time.Now()
			// Wait for the receipt
			receipt, err := waitForReceipt(backend, tx)
			if err != nil {
				return err
			}
			ometrics.GasOracleStats.TxConfTimer.Update(time.Since(pre))

			log.Info("L2 gas price transaction confirmed", "hash", tx.Hash().Hex(),
				"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
		}
		return nil
	}, nil
}

// Only update the gas price when it must be changed by at least
// a paramaterizable amount. If the param is greater than the result
// of 1 - (min/max) where min and max are the gas prices then do not
// update the gas price
func isDifferenceSignificant(a, b uint64, c float64) bool {
	max := max(a, b)
	min := min(a, b)
	factor := 1 - (float64(min) / float64(max))
	return c <= factor
}

// Wait for the receipt by polling the backend
func waitForReceipt(backend DeployContractBackend, tx *types.Transaction) (*types.Receipt, error) {
	t := time.NewTicker(300 * time.Millisecond)
	receipt := new(types.Receipt)
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

func max(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b uint64) uint64 {
	if a >= b {
		return b
	}
	return a
}
