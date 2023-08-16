package oracle

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"

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

var jumpTable = make(map[int]*big.Int, 0)
var orderedSizes = make([]int, 0)

func wrapUpdateOverhead(l2Backend DeployContractBackend, cfg *Config) (func(*big.Int, *big.Int) error, error) {
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
	return func(diff *big.Int, size *big.Int) error {
		calculateJumpTable(diff, cfg)
		newOverheadLevel, err := getOverheadLevelInJumpTable(size)
		if err != nil {
			return err
		}
		overhead, err := contract.Overhead(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		ometrics.GasOracleStats.OverHeadUpdateGauge.Inc(1)
		// skip update if overhead is not changed
		if overhead.Cmp(newOverheadLevel) == 0 {
			log.Info("skip update overhead", "overhead", overhead)
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
		tx, err := contract.SetOverhead(opts, newOverheadLevel)
		if err != nil {
			return err
		}
		log.Info("updating L1 overhead", "tx.gasPrice", tx.GasPrice(), "tx.gasLimit", tx.Gas(),
			"tx.data", hexutil.Encode(tx.Data()), "tx.to", tx.To().Hex(), "tx.nonce", tx.Nonce())
		if err := l2Backend.SendTransaction(context.Background(), tx); err != nil {
			return fmt.Errorf("cannot update base fee: %w", err)
		}
		log.Info("L2 overhead transaction sent", "hash", tx.Hash().Hex(), "old overhead", overhead, "new overhead", newOverheadLevel)
		ometrics.GasOracleStats.OverHeadGauge.Update(newOverheadLevel.Int64())

		if cfg.waitForReceipt {
			// Wait for the receipt
			receipt, err := waitForReceipt(l2Backend, tx)
			if err != nil {
				return err
			}

			log.Info("overhead transaction confirmed", "hash", tx.Hash().Hex(),
				"gas-used", receipt.GasUsed, "blocknumber", receipt.BlockNumber)
		}
		return nil
	}, nil
}

func calculateJumpTable(diff *big.Int, cfg *Config) {
	// data rollup cost
	var sequencerOverhead = new(big.Int).Mul(diff, cfg.dataRollupGasUsed)
	// sum up
	var OverheadGasUsedOnL1 = new(big.Int).Add(cfg.stateRollupGasUsed, sequencerOverhead)
	// calculate jump table
	for levelSize := cfg.batchSizeBottom; levelSize <= cfg.batchSizeCap; {
		orderedSizes = append(orderedSizes, levelSize)
		jumpTable[levelSize] = new(big.Int).Add(new(big.Int).Div(OverheadGasUsedOnL1, new(big.Int).SetUint64(uint64(levelSize))), cfg.stateHashGasUsed)
		levelSize += cfg.sizeGap
	}
}

func getOverheadLevelInJumpTable(batchSize *big.Int) (*big.Int, error) {
	size := int(batchSize.Int64())
	// check jumpTable is initialised
	if len(jumpTable) <= 0 {
		return nil, fmt.Errorf("jump table is not initialised")
	}
	// init into memory
	if len(orderedSizes) == 0 {
		i := 0
		for key := range jumpTable {
			orderedSizes = append(orderedSizes, key)
			i++
		}
		sort.Ints(orderedSizes)
	}
	// search Size in table
	if orderedSizes[0] >= size {
		return jumpTable[orderedSizes[0]], nil
	}
	if orderedSizes[len(orderedSizes)-1] <= size {
		return jumpTable[orderedSizes[len(orderedSizes)-1]], nil
	}
	for i := 1; i < len(orderedSizes); i++ {
		if orderedSizes[i] >= size {
			return jumpTable[orderedSizes[i-1]], nil
		}
	}
	return nil, fmt.Errorf("cant find overhead in jump table with given batch size: %d", size)
}
