package proposer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	l2types "github.com/mantlenetworkio/mantle/l2geth/core/types"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/ctc"
	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/scc"
	tssClient "github.com/mantlenetworkio/mantle/batch-submitter/tss-client"
	"github.com/mantlenetworkio/mantle/bss-core/drivers"
	"github.com/mantlenetworkio/mantle/bss-core/metrics"
	"github.com/mantlenetworkio/mantle/bss-core/txmgr"
	fpbindings "github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	tss_types "github.com/mantlenetworkio/mantle/tss/common"
)

// stateRootSize is the size in bytes of a state root.
const stateRootSize = 32

var bigOne = new(big.Int).SetUint64(1) //nolint:unused

type Config struct {
	Name                 string
	L1Client             *ethclient.Client
	L2Client             *l2ethclient.Client
	TssClient            *tssClient.Client
	BlockOffset          uint64
	MaxStateRootElements uint64
	MinStateRootElements uint64
	SCCAddr              common.Address
	CTCAddr              common.Address
	FPRollupAddr         common.Address
	ChainID              *big.Int
	PrivKey              *ecdsa.PrivateKey
}

type Driver struct {
	cfg            Config
	sccContract    *scc.StateCommitmentChain
	rawSccContract *bind.BoundContract
	ctcContract    *ctc.CanonicalTransactionChain
	fpRollup       *fpbindings.Rollup
	rawFPContract  *bind.BoundContract
	fpAssertion    *fpbindings.AssertionMap
	walletAddr     common.Address
	metrics        *metrics.Base
}

func NewDriver(cfg Config) (*Driver, error) {
	sccContract, err := scc.NewStateCommitmentChain(
		cfg.SCCAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	ctcContract, err := ctc.NewCanonicalTransactionChain(
		cfg.CTCAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	fpRollup, err := fpbindings.NewRollup(
		cfg.FPRollupAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	assertionAddr, err := fpRollup.Assertions(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	assertionMap, err := fpbindings.NewAssertionMap(
		assertionAddr, cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}

	parsedSCC, err := abi.JSON(strings.NewReader(
		scc.StateCommitmentChainABI,
	))
	if err != nil {
		return nil, err
	}
	parsedFP, err := abi.JSON(strings.NewReader(
		fpbindings.RollupABI,
	))
	if err != nil {
		return nil, err
	}

	rawSccContract := bind.NewBoundContract(
		cfg.SCCAddr, parsedSCC, cfg.L1Client, cfg.L1Client, cfg.L1Client,
	)

	rawFPContract := bind.NewBoundContract(
		cfg.FPRollupAddr, parsedFP, cfg.L1Client, cfg.L1Client, cfg.L1Client,
	)

	walletAddr := crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)

	return &Driver{
		cfg:            cfg,
		sccContract:    sccContract,
		rawSccContract: rawSccContract,
		ctcContract:    ctcContract,
		fpRollup:       fpRollup,
		rawFPContract:  rawFPContract,
		fpAssertion:    assertionMap,
		walletAddr:     walletAddr,
		metrics:        metrics.NewBase("batch_submitter", cfg.Name),
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

	start, err := d.sccContract.GetTotalElements(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	if err != nil {
		return nil, nil, err
	}
	start.Add(start, blockOffset)

	end, err := d.ctcContract.GetTotalElements(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	if err != nil {
		return nil, nil, err
	}
	end.Add(end, blockOffset)

	if start.Cmp(end) > 0 {
		return nil, nil, fmt.Errorf("invalid range, "+
			"end(%v) < start(%v)", end, start)
	}

	return start, end, nil
}

// CraftBatchTx transforms the L2 blocks between start and end into a batch
// transaction using the given nonce. A dummy gas price is used in the resulting
// transaction to use for size estimation.
//
// NOTE: This method SHOULD NOT publish the resulting transaction.
func (d *Driver) CraftBatchTx(
	ctx context.Context,
	start, end, nonce *big.Int,
) (*types.Transaction, error) {

	name := d.cfg.Name

	log.Info(name+" crafting batch tx", "start", start, "end", end, "nonce", nonce)

	var blocks []*l2types.Block
	var stateRoots [][stateRootSize]byte
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		// Consume state roots until reach our maximum tx size.
		if uint64(len(stateRoots)) > d.cfg.MaxStateRootElements {
			break
		}

		block, err := d.cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
		stateRoots = append(stateRoots, block.Root())
	}

	// Abort if we don't have enough state roots to meet our minimum
	// requirement.
	if uint64(len(stateRoots)) < d.cfg.MinStateRootElements {
		log.Info(name+" number of state roots  below minimum",
			"num_state_roots", len(stateRoots),
			"min_state_roots", d.cfg.MinStateRootElements)
		return nil, nil
	}

	d.metrics.NumElementsPerBatch().Observe(float64(len(stateRoots)))

	log.Info(name+" batch constructed", "num_state_roots", len(stateRoots))

	opts, err := bind.NewKeyedTransactorWithChainID(
		d.cfg.PrivKey, d.cfg.ChainID,
	)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	blockOffset := new(big.Int).SetUint64(d.cfg.BlockOffset)
	offsetStartsAtIndex := new(big.Int).Sub(start, blockOffset)
	// Assembly data request tss node signature
	tssReqParams := tss_types.SignStateRequest{
		StartBlock:          start.String(),
		OffsetStartsAtIndex: offsetStartsAtIndex.String(),
		StateRoots:          stateRoots,
	}
	tssReponseBytes, err := d.cfg.TssClient.GetSignStateBatch(tssReqParams)
	if err != nil {
		log.Error("get tss manager signature fail", "err", err)
		return nil, err
	}
	var tssResponse tssClient.TssResponse
	err = json.Unmarshal(tssReponseBytes, &tssResponse)
	if err != nil {
		log.Error("failed to unmarshal response from tss", "err", err)
		return nil, err
	}

	log.Info("append log", "stateRoots", fmt.Sprintf("%v", stateRoots), "offsetStartsAtIndex", offsetStartsAtIndex, "signature", hex.EncodeToString(tssResponse.Signature), "rollback", tssResponse.RollBack)
	var tx *types.Transaction
	if tssResponse.RollBack {
		tx, err = d.sccContract.RollBackL2Chain(opts, start, offsetStartsAtIndex, tssResponse.Signature)
	} else {
		if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
			log.Info("append state with fraud proof")
			// ##### FRAUD-PROOF modify #####
			// check stake initialised
			tx, err = d.FraudProofAppendStateBatch(
				opts, stateRoots, offsetStartsAtIndex, tssResponse.Signature, blocks,
			)
			if err != nil {
				log.Error("fraud proof append state batch in error: ", err.Error())
			}
			// ##### FRAUD-PROOF modify ##### //
		} else {
			log.Info("append state with scc")
			tx, err = d.sccContract.AppendStateBatch(
				opts, stateRoots, offsetStartsAtIndex, tssResponse.Signature,
			)
		}
	}

	switch {
	case err == nil:
		return tx, nil
	// If the transaction failed because the backend does not support
	// eth_maxPriorityFeePerGas, fallback to using the default constant.
	// Currently Alchemy is the only backend provider that exposes this method,
	// so in the event their API is unreachable we can fallback to a degraded
	// mode of operation. This also applies to our test environments, as hardhat
	// doesn't support the query either.
	case drivers.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn(d.cfg.Name + " eth_maxPriorityFeePerGas is unsupported " +
			"by current backend, using fallback gasTipCap")
		opts.GasTipCap = drivers.FallbackGasTipCap
		if tssResponse.RollBack {
			return d.sccContract.RollBackL2Chain(
				opts, start, offsetStartsAtIndex, tssResponse.Signature,
			)
		} else {
			if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
				log.Info("append state with fraud proof by gas tip cap")
				// ##### FRAUD-PROOF modify #####
				// check stake initialised
				return d.FraudProofAppendStateBatch(
					opts, stateRoots, offsetStartsAtIndex, tssResponse.Signature, blocks,
				)
				// ##### FRAUD-PROOF modify ##### //
			} else {
				log.Info("append state with scc by gas tip cap")
				return d.sccContract.AppendStateBatch(
					opts, stateRoots, offsetStartsAtIndex, tssResponse.Signature,
				)
			}
		}
	default:
		return nil, err
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
	var finalTx *types.Transaction
	var err error

	opts, err := bind.NewKeyedTransactorWithChainID(
		d.cfg.PrivKey, d.cfg.ChainID,
	)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true

	if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
		// ##### FRAUD-PROOF modify #####
		log.Info("get RawTransact from Fraud Proof")
		finalTx, err = d.rawFPContract.RawTransact(opts, tx.Data())
		// ##### FRAUD-PROOF modify ##### //
	} else {
		log.Info("get RawTransact from SCC")
		finalTx, err = d.rawSccContract.RawTransact(opts, tx.Data())
	}
	switch {
	case err == nil:
		return finalTx, nil

	// If the transaction failed because the backend does not support
	// eth_maxPriorityFeePerGas, fallback to using the default constant.
	// Currently Alchemy is the only backend provider that exposes this method,
	// so in the event their API is unreachable we can fallback to a degraded
	// mode of operation. This also applies to our test environments, as hardhat
	// doesn't support the query either.
	case drivers.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn(d.cfg.Name + " eth_maxPriorityFeePerGas is unsupported " +
			"by current backend, using fallback gasTipCap")
		opts.GasTipCap = drivers.FallbackGasTipCap

		if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
			// ##### FRAUD-PROOF modify #####
			log.Info("get RawTransact from Fraud Proof by gas tip cap")
			return d.rawFPContract.RawTransact(opts, tx.Data())
			// ##### FRAUD-PROOF modify ##### //
		} else {
			log.Info("get RawTransact from SCC by gas tip cap")
			return d.rawSccContract.RawTransact(opts, tx.Data())
		}

	default:
		return nil, err
	}
}

// SendTransaction injects a signed transaction into the pending pool for
// execution.
func (d *Driver) SendTransaction(
	ctx context.Context,
	tx *types.Transaction,
) error {
	return d.cfg.L1Client.SendTransaction(ctx, tx)
}

func (d *Driver) FraudProofAppendStateBatch(opts *bind.TransactOpts, batch [][32]byte, shouldStartAtElement *big.Int, signature []byte, blocks []*l2types.Block) (*types.Transaction, error) {

	challengeContext, _ := d.fpRollup.ChallengeCtx(&bind.CallOpts{})
	isInChallenge := challengeContext.DefenderAssertionID.Uint64() != 0 && !challengeContext.Completed
	if isInChallenge {
		log.Warn("currently in challenge, can't submit new assertion")
		return nil, nil
	}
	var latestAssertion rollupTypes.Assertion
	var staker rollupTypes.Staker
	if ret, err := d.fpRollup.Stakers(&bind.CallOpts{}, opts.From); err != nil {
		return nil, err
	} else {
		staker.IsStaked = ret.IsStaked
		staker.AmountStaked = ret.AmountStaked
		staker.AssertionID = ret.AssertionID
		staker.CurrentChallenge = ret.CurrentChallenge
	}
	if ret, err := d.fpAssertion.Assertions(&bind.CallOpts{}, staker.AssertionID); err != nil {
		return nil, err
	} else {
		latestAssertion.ID = staker.AssertionID
		latestAssertion.VmHash = ret.StateHash
		latestAssertion.InboxSize = ret.InboxSize
		latestAssertion.Parent = ret.Parent
		latestAssertion.Deadline = ret.Deadline
		latestAssertion.ProposalTime = ret.ProposalTime
	}

	txBatch := rollupTypes.NewTxBatch(blocks, uint64(len(blocks)))
	assertion := txBatch.ToAssertion(&latestAssertion)

	fmt.Println(assertion.VmHash.String())
	fmt.Println(assertion.InboxSize)
	fmt.Println(batch)
	fmt.Println(shouldStartAtElement)
	fmt.Println(signature)

	// create assertion
	return d.fpRollup.CreateAssertionWithStateBatch(
		opts, assertion.VmHash, assertion.InboxSize, batch, shouldStartAtElement, signature)
}
