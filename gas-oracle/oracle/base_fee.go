package oracle

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"

	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
	ometrics "github.com/mantlenetworkio/mantle/gas-oracle/metrics"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"
)

func wrapUpdateBaseFee(l1Backend bind.ContractTransactor, l2Backend DeployContractBackend, cfg *Config) (func() error, error) {
	if cfg.l2ChainID == nil {
		return nil, errNoChainID
	}

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

	// initialize some metrics
	// initialize fee scalar from contract
	feeScalar, err := contract.Scalar(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return nil, err
	}
	ometrics.GasOracleStats.FeeScalarGauge.Update(feeScalar.Int64())

	return func() error {
		baseFee, err := contract.L1BaseFee(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		feeScalar, err := contract.Scalar(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		// Update fee scalar metrics
		ometrics.GasOracleStats.FeeScalarGauge.Update(feeScalar.Int64())

		// NOTE this will return base multiple with coin ratio
		log.Info("get header in l1 client", "type is", reflect.ValueOf(l1Backend).Type())
		tip, err := l1Backend.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err
		}
		if tip.BaseFee == nil {
			return errNoBaseFee
		}
		if !isDifferenceSignificant(baseFee.Uint64(), tip.BaseFee.Uint64(), cfg.l1BaseFeeSignificanceFactor) {
			log.Warn("non significant base fee update", "tip", tip.BaseFee, "current", baseFee)
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
		// set L1BaseFee to base fee + tip cap, to cover rollup tip cap
		tx, err := contract.SetL1BaseFee(opts, tip.BaseFee)
		if err != nil {
			return err
		}
		log.Info("updating L1 base fee", "tx.gasPrice", tx.GasPrice(), "tx.gasLimit", tx.Gas(),
			"tx.data", hexutil.Encode(tx.Data()), "tx.to", tx.To().Hex(), "tx.nonce", tx.Nonce())
		if err := l2Backend.SendTransaction(context.Background(), tx); err != nil {
			return fmt.Errorf("cannot update base fee: %w", err)
		}
		log.Info("L1 base fee transaction already sent", "hash", tx.Hash().Hex(), "baseFee", tip.BaseFee)
		ometrics.GasOracleStats.L1BaseFeeGauge.Update(tip.BaseFee.Int64())

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
