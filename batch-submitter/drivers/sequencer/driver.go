package sequencer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/ctc"
	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/da"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/bss-core/drivers"
	"github.com/mantlenetworkio/mantle/bss-core/metrics"
	"github.com/mantlenetworkio/mantle/bss-core/txmgr"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
)

const (
	appendSequencerBatchMethodName = "appendSequencerBatch"
)

var bigOne = new(big.Int).SetUint64(1)

type Config struct {
	Name                string
	L1Client            *ethclient.Client
	L2Client            *l2ethclient.Client
	BlockOffset         uint64
	CTCAddr             common.Address
	DaUpgradeBlock      uint64
	DAAddr              common.Address
	ChainID             *big.Int
	PrivKey             *ecdsa.PrivateKey
	EnableSequencerHsm  bool
	SequencerHsmAddress string
	SequencerHsmAPIName string
	SequencerHsmCreden  string
	BatchType           BatchType
	MaxRollupTxn        uint64
	MinRollupTxn        uint64
}

type Driver struct {
	cfg              Config
	ctcContract      *ctc.CanonicalTransactionChain
	daContract       *da.BVMEigenDataLayrChain
	rawCtcContract   *bind.BoundContract
	rawDaCtcContract *bind.BoundContract
	walletAddr       common.Address
	ctcABI           *abi.ABI
	DaABI            *abi.ABI
	metrics          *Metrics
}

func NewDriver(cfg Config) (*Driver, error) {
	ctcContract, err := ctc.NewCanonicalTransactionChain(
		cfg.CTCAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewCanonicalTransactionChain in error", "error", err)
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(
		ctc.CanonicalTransactionChainABI,
	))
	if err != nil {
		log.Error("Parse CanonicalTransactionChain in error", "error", err)
		return nil, err
	}

	ctcABI, err := ctc.CanonicalTransactionChainMetaData.GetAbi()
	if err != nil {
		log.Error("Get CanonicalTransactionChain ABI in error", "error", err)
		return nil, err
	}

	rawCtcContract := bind.NewBoundContract(
		cfg.CTCAddr, parsed, cfg.L1Client, cfg.L1Client,
		cfg.L1Client,
	)

	daContract, err := da.NewBVMEigenDataLayrChain(
		cfg.DAAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewBVMEigenDataLayrChain in error", "error", err)
		return nil, err
	}
	daParsed, err := abi.JSON(strings.NewReader(
		da.BVMEigenDataLayrChainABI,
	))
	if err != nil {
		log.Error("Parse BVMEigenDataLayrChain in error", "error", err)
		return nil, err
	}

	daABI, err := da.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		log.Error("Get BVMEigenDataLayrChain ABI in error", "error", err)
		return nil, err
	}

	rawDaCtcContract := bind.NewBoundContract(
		cfg.DAAddr, daParsed, cfg.L1Client, cfg.L1Client,
		cfg.L1Client,
	)

	var walletAddr common.Address
	if cfg.EnableSequencerHsm {
		walletAddr = common.HexToAddress(cfg.SequencerHsmAddress)
		log.Info("use sequencer hsm", "walletaddr", walletAddr)
	} else {
		walletAddr = crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)
		log.Info("not use sequencer hsm", "walletaddr", walletAddr)
	}

	return &Driver{
		cfg:              cfg,
		ctcContract:      ctcContract,
		daContract:       daContract,
		rawCtcContract:   rawCtcContract,
		rawDaCtcContract: rawDaCtcContract,
		walletAddr:       walletAddr,
		ctcABI:           ctcABI,
		DaABI:            daABI,
		metrics:          NewMetrics(cfg.Name),
	}, nil
}

// Name is an identifier used to prefix logs for a particular service.
func (d *Driver) Name() string {
	return d.cfg.Name
}

// WalletAddr is the wallet address used to pay for batch transaction fees.
func (d *Driver) WalletAddr() common.Address {
	return d.walletAddr
}

// Metrics returns the subservice telemetry object.
func (d *Driver) Metrics() metrics.Metrics {
	return d.metrics
}

// ClearPendingTx a publishes a transaction at the next available nonce in order
// to clear any transactions in the mempool left over from a prior running
// instance of the batch submitter.
func (d *Driver) ClearPendingTx(
	ctx context.Context,
	txMgr txmgr.TxManager,
	l1Client *ethclient.Client,
) error {

	return drivers.ClearPendingTx(
		d.cfg.Name, ctx, txMgr, l1Client, d.walletAddr, d.cfg.PrivKey,
		d.cfg.ChainID,
	)
}

// GetBatchBlockRange returns the start and end L2 block heights that need to be
// processed. Note that the end value is *exclusive*, therefore if the returned
// values are identical nothing needs to be processed.
func (d *Driver) GetBatchBlockRange(
	ctx context.Context) (*big.Int, *big.Int, error) {

	blockOffset := new(big.Int).SetUint64(d.cfg.BlockOffset)

	start, err := d.ctcContract.GetTotalElements(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	if err != nil {
		return nil, nil, err
	}
	start.Add(start, blockOffset)

	latestHeader, err := d.cfg.L2Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	var end *big.Int
	// for upgrade
	if latestHeader.Number.Cmp(big.NewInt(int64(d.cfg.DaUpgradeBlock))) < 0 {
		// Add one because end is *exclusive*.
		end = new(big.Int).Add(latestHeader.Number, bigOne)
	} else {
		end, err = d.daContract.GetL2ConfirmedBlockNumber(&bind.CallOpts{
			Context: context.Background(),
		})
		if err != nil {
			return nil, nil, err
		}
	}
	l2Txn := big.NewInt(0).Sub(end, start)
	if l2Txn.Cmp(big.NewInt(int64(d.cfg.MinRollupTxn))) < 0 {
		return start, start, nil
	}
	if l2Txn.Cmp(big.NewInt(int64(d.cfg.MaxRollupTxn))) > 0 {
		end = big.NewInt(0).Add(start, big.NewInt(int64(d.cfg.MaxRollupTxn)))
	}
	return start, end, nil
}

// CraftBatchTx transforms the L2 blocks between start and end into a batch
// transaction using the given nonce. A dummy gas price is used in the resulting
// transaction to use for size estimation. A nil transaction is returned if the
// transaction does not meet the minimum size requirements.
//
// NOTE: This method SHOULD NOT publish the resulting transaction.
func (d *Driver) CraftBatchTx(
	ctx context.Context,
	start, end, nonce *big.Int,
) (*types.Transaction, error) {

	name := d.cfg.Name

	log.Info(name+" crafting batch tx", "start", start, "end", end,
		"nonce", nonce, "type", d.cfg.BatchType.String())

	var lastTimestamp uint64
	var lastBlockNumber uint64
	numSequencedTxs := 0
	numSubsequentQueueTxs := 0

	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		block, err := d.cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, err
		}
		// For each sequencer transaction, update our running total with the
		// size of the transaction.
		batchElement := BatchElementFromBlock(block)
		if batchElement.IsSequencerTx {
			numSequencedTxs += 1
		} else {
			numSubsequentQueueTxs += 1
		}
		if i.Cmp(big.NewInt(0).Sub(end, bigOne)) == 0 {
			lastTimestamp = batchElement.Timestamp
			lastBlockNumber = batchElement.BlockNumber
		}
	}
	blocksLen := numSequencedTxs + numSubsequentQueueTxs
	shouldStartAt := start.Uint64()

	for {
		var (
			contexts []BatchContext
		)

		batchContext := BatchContext{
			NumSequencedTxs:       uint64(numSequencedTxs),
			NumSubsequentQueueTxs: uint64(numSubsequentQueueTxs),
			Timestamp:             lastTimestamp,
			BlockNumber:           lastBlockNumber,
		}

		d.metrics.BatchNumSequencedTxs().Set(float64(batchContext.NumSequencedTxs))
		d.metrics.BatchNumSubsequentQueueTxs().Set(float64(batchContext.NumSubsequentQueueTxs))
		d.metrics.BatchTimestamp().Set(float64(batchContext.Timestamp))
		d.metrics.BatchBlockNumber().Set(float64(batchContext.BlockNumber))

		contexts = append(contexts, batchContext)
		batchParams := &AppendSequencerBatchParams{
			ShouldStartAtElement:  shouldStartAt - d.cfg.BlockOffset,
			TotalElementsToAppend: uint64(blocksLen),
			Contexts:              contexts,
		}

		// Encode the batch arguments using the configured encoding type.
		batchArguments, err := batchParams.Serialize(d.cfg.BatchType)
		if err != nil {
			return nil, err
		}

		appendSequencerBatchID := d.ctcABI.Methods[appendSequencerBatchMethodName].ID
		calldata := append(appendSequencerBatchID, batchArguments...)

		log.Info(name+" testing batch size",
			"calldata_size", len(calldata))

		d.metrics.NumElementsPerBatch().Observe(float64(blocksLen))

		log.Info(name+" batch constructed",
			"num_txs", blocksLen,
			"final_size", len(calldata),
			"batch_type", d.cfg.BatchType)

		var opts *bind.TransactOpts
		if d.cfg.EnableSequencerHsm {
			seqBytes, err := hex.DecodeString(d.cfg.SequencerHsmCreden)
			apikey := option.WithCredentialsJSON(seqBytes)
			client, err := kms.NewKeyManagementClient(ctx, apikey)
			if err != nil {
				log.Info("sequencer", "create signer error", err.Error())
				return nil, err
			}
			mk := &bsscore.ManagedKey{
				KeyName:      d.cfg.SequencerHsmAPIName,
				EthereumAddr: common.HexToAddress(d.cfg.SequencerHsmAddress),
				Gclient:      client,
			}
			opts, err = mk.NewEthereumTransactorrWithChainID(ctx, d.cfg.ChainID)
			if err != nil {
				log.Info("sequencer", "create signer error", err.Error())
				return nil, err
			}
		} else {
			opts, err = bind.NewKeyedTransactorWithChainID(
				d.cfg.PrivKey, d.cfg.ChainID,
			)
			if err != nil {
				log.Info("sequencer", "create signer error", err.Error())
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}
		opts.Context = ctx
		opts.Nonce = nonce
		opts.NoSend = true

		tx, err := d.rawCtcContract.RawTransact(opts, calldata)
		switch {
		case err == nil:
			return tx, nil

		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		case drivers.IsMaxPriorityFeePerGasNotFoundError(err):
			log.Warn(d.cfg.Name + " eth_maxPriorityFeePerGas is unsupported " +
				"by current backend, using fallback gasTipCap")
			opts.GasTipCap = drivers.FallbackGasTipCap
			return d.rawCtcContract.RawTransact(opts, calldata)

		default:
			return nil, err
		}
	}
}

// UpdateGasPrice signs an otherwise identical txn to the one provided but with
// updated gas prices sampled from the existing network conditions.
//
// NOTE: Thie method SHOULD NOT publish the resulting transaction.
func (d *Driver) UpdateGasPrice(
	ctx context.Context,
	tx *types.Transaction,
) (*types.Transaction, error) {

	gasTipCap, err := d.cfg.L1Client.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		if !drivers.IsMaxPriorityFeePerGasNotFoundError(err) {
			return nil, err
		}

		log.Warn(d.cfg.Name + " eth_maxPriorityFeePerGas is unsupported " +
			"by current backend, using fallback gasTipCap")
		gasTipCap = drivers.FallbackGasTipCap
	}

	header, err := d.cfg.L1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	gasFeeCap := txmgr.CalcGasFeeCap(header.BaseFee, gasTipCap)

	// The estimated gas limits performed by RawTransact fail semi-regularly
	// with out of gas exceptions. To remedy this we extract the internal calls
	// to perform gas price/gas limit estimation here and add a buffer to
	// account for any network variability.
	gasLimit, err := d.cfg.L1Client.EstimateGas(ctx, ethereum.CallMsg{
		From:      d.walletAddr,
		To:        &d.cfg.CTCAddr,
		GasPrice:  nil,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      tx.Data(),
	})
	if err != nil {
		return nil, err
	}

	var opts *bind.TransactOpts
	if d.cfg.EnableSequencerHsm {
		seqBytes, err := hex.DecodeString(d.cfg.SequencerHsmCreden)
		apikey := option.WithCredentialsJSON(seqBytes)
		client, err := kms.NewKeyManagementClient(ctx, apikey)
		if err != nil {
			return nil, err
		}
		mk := &bsscore.ManagedKey{
			KeyName:      d.cfg.SequencerHsmAPIName,
			EthereumAddr: common.HexToAddress(d.cfg.SequencerHsmAddress),
			Gclient:      client,
		}
		opts, err = mk.NewEthereumTransactorrWithChainID(ctx, d.cfg.ChainID)
		log.Info("sequencer", "enable-hsm", true)
	} else {
		opts, err = bind.NewKeyedTransactorWithChainID(
			d.cfg.PrivKey, d.cfg.ChainID,
		)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = 6 * gasLimit / 5 // add 20% buffer to gas limit
	opts.NoSend = true

	return d.rawCtcContract.RawTransact(opts, tx.Data())
}

// SendTransaction injects a signed transaction into the pending pool for
// execution.
func (d *Driver) SendTransaction(
	ctx context.Context,
	tx *types.Transaction,
) error {
	return d.cfg.L1Client.SendTransaction(ctx, tx)
}
