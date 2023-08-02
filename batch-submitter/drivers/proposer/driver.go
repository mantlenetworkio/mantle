package proposer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/ctc"
	"github.com/mantlenetworkio/mantle/batch-submitter/bindings/scc"
	tssClient "github.com/mantlenetworkio/mantle/batch-submitter/tss-client"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/bss-core/drivers"
	"github.com/mantlenetworkio/mantle/bss-core/metrics"
	"github.com/mantlenetworkio/mantle/bss-core/txmgr"
	fpbindings "github.com/mantlenetworkio/mantle/fraud-proof/bindings"
	rollupTypes "github.com/mantlenetworkio/mantle/fraud-proof/rollup/types"
	l2types "github.com/mantlenetworkio/mantle/l2geth/core/types"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	tss_types "github.com/mantlenetworkio/mantle/tss/common"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"
)

// stateRootSize is the size in bytes of a state root.
const stateRootSize = 32

// block number buffer for dtl to sync data
const blockBuffer = 2

var bigOne = new(big.Int).SetUint64(1) //nolint:unused

type Config struct {
	Name                        string
	L1Client                    *ethclient.Client
	L2Client                    *l2ethclient.Client
	TssClient                   *tssClient.Client
	BlockOffset                 uint64
	MaxStateRootElements        uint64
	MinStateRootElements        uint64
	SCCAddr                     common.Address
	CTCAddr                     common.Address
	FPRollupAddr                common.Address
	ChainID                     *big.Int
	PrivKey                     *ecdsa.PrivateKey
	SccRollback                 bool
	RollupTimeout               time.Duration
	PollInterval                time.Duration
	FinalityConfirmations       uint64
	EnableProposerHsm           bool
	ProposerHsmCreden           string
	ProposerHsmAddress          string
	ProposerHsmAPIName          string
	AllowL2AutoRollback         bool
	MinTimeoutStateRootElements uint64
}

type Driver struct {
	cfg                  Config
	sccContract          *scc.StateCommitmentChain
	rawSccContract       *bind.BoundContract
	ctcContract          *ctc.CanonicalTransactionChain
	fpRollup             *fpbindings.Rollup
	rawFPContract        *bind.BoundContract
	fpAssertion          *fpbindings.AssertionMap
	walletAddr           common.Address
	rollbackEndBlock     *big.Int
	rollbackEndStateRoot [stateRootSize]byte
	once                 sync.Once
	lastCommitTime       time.Time
	lastStart            *big.Int
	metrics              *metrics.Base
}

func NewDriver(cfg Config) (*Driver, error) {
	log.Info("Show configration", "cfg.SCCAddr", cfg.SCCAddr, "cfg.CTCAddr", cfg.CTCAddr, "cfg.FPRollupAddr", cfg.FPRollupAddr)
	sccContract, err := scc.NewStateCommitmentChain(
		cfg.SCCAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewStateCommitmentChain in error", "error", err)
		return nil, err
	}

	ctcContract, err := ctc.NewCanonicalTransactionChain(
		cfg.CTCAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewCanonicalTransactionChain in error", "error", err)
		return nil, err
	}

	fpRollup, err := fpbindings.NewRollup(
		cfg.FPRollupAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewRollup in error", "error", err)
		return nil, err
	}

	assertionAddr, err := fpRollup.Assertions(&bind.CallOpts{})
	if err != nil {
		log.Error("fpRollup get Assertions in error", "error", err)
		return nil, err
	}

	assertionMap, err := fpbindings.NewAssertionMap(
		assertionAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("NewAssertionMap in error", "error", err)
		return nil, err
	}

	parsedSCC, err := abi.JSON(strings.NewReader(
		scc.StateCommitmentChainABI,
	))
	if err != nil {
		log.Error("Parse StateCommitmentChain ABI in error", "error", err)
		return nil, err
	}
	parsedFP, err := abi.JSON(strings.NewReader(
		fpbindings.RollupABI,
	))
	if err != nil {
		log.Error("Parse Rollup ABI in error", "error", err)
		return nil, err
	}

	rawSccContract := bind.NewBoundContract(
		cfg.SCCAddr, parsedSCC, cfg.L1Client, cfg.L1Client, cfg.L1Client,
	)

	rawFPContract := bind.NewBoundContract(
		cfg.FPRollupAddr, parsedFP, cfg.L1Client, cfg.L1Client, cfg.L1Client,
	)

	var walletAddr common.Address
	if cfg.EnableProposerHsm {
		walletAddr = common.HexToAddress(cfg.ProposerHsmAddress)
		log.Info("use proposer hsm", "walletaddr", walletAddr)
	} else {
		walletAddr = crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)
		log.Info("not use proposer hsm", "walletaddr", walletAddr)
	}

	return &Driver{
		cfg:                  cfg,
		sccContract:          sccContract,
		rawSccContract:       rawSccContract,
		ctcContract:          ctcContract,
		fpRollup:             fpRollup,
		rawFPContract:        rawFPContract,
		fpAssertion:          assertionMap,
		walletAddr:           walletAddr,
		rollbackEndBlock:     big.NewInt(0),
		rollbackEndStateRoot: [stateRootSize]byte{},
		once:                 sync.Once{},
		lastStart:            big.NewInt(0),
		metrics:              metrics.NewBase("batch_submitter", cfg.Name),
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

	currentHeader, err := d.cfg.L1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	finality := new(big.Int).SetUint64(d.cfg.FinalityConfirmations)
	finality.Add(finality, new(big.Int).SetInt64(blockBuffer)) // add 2 block number buffer to dtl sync data
	currentNumber := currentHeader.Number
	currentNumber.Sub(currentNumber, finality)

	end, err := d.ctcContract.GetTotalElements(&bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: currentNumber,
	})
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

	if start.Cmp(d.lastStart) > 0 {
		d.lastStart = start
		d.lastCommitTime = time.Now().Add(-d.cfg.PollInterval)
	}
	//If the waiting time has not been reached, then check whether the minimum stateroot number
	//is met. if not, return nil
	rollupTxn := end.Uint64() - start.Uint64()
	if rollupTxn < d.cfg.MinStateRootElements && (d.lastCommitTime.Add(d.cfg.RollupTimeout).After(time.Now()) || rollupTxn < d.cfg.MinTimeoutStateRootElements) {
		if rollupTxn < d.cfg.MinStateRootElements {
			log.Info(name+" number of state roots  below minimum",
				"num_state_roots", rollupTxn,
				"min_state_roots", d.cfg.MinStateRootElements)
			return nil, nil
		}
		log.Info(name+" number of timeout state roots below minimum or timeout can't satisfy the constrain",
			"num_state_roots", rollupTxn,
			"min_timeout_state_roots", d.cfg.MinTimeoutStateRootElements)
		return nil, nil
	}

	var blocks []*l2types.Block
	var stateRoots [][stateRootSize]byte
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		// Consume state roots until reach our maximum tx size.
		if uint64(len(stateRoots)) > d.cfg.MaxStateRootElements {
			end = i
			log.Info("range is big than max stateroot elements", "elements", d.cfg.MaxStateRootElements, "start", start, "new end", end)
			break
		}

		block, err := d.cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
		stateRoots = append(stateRoots, block.Root())
	}

	d.metrics.NumElementsPerBatch().Observe(float64(len(stateRoots)))

	log.Info(name+" batch constructed", "num_state_roots", len(stateRoots))

	var opts *bind.TransactOpts
	var err error
	if d.cfg.EnableProposerHsm {
		proBytes, err := hex.DecodeString(d.cfg.ProposerHsmCreden)
		if err != nil {
			return nil, err
		}
		apikey := option.WithCredentialsJSON(proBytes)
		client, err := kms.NewKeyManagementClient(ctx, apikey)
		if err != nil {
			return nil, err
		}
		mk := &bsscore.ManagedKey{
			KeyName:      d.cfg.ProposerHsmAPIName,
			EthereumAddr: common.HexToAddress(d.cfg.ProposerHsmAddress),
			Gclient:      client,
		}
		opts, err = mk.NewEthereumTransactorrWithChainID(ctx, d.cfg.ChainID)
		if err != nil {
			return nil, err
		}
		log.Info("proposer", "enable-hsm", true)
	} else {
		opts, err = bind.NewKeyedTransactorWithChainID(
			d.cfg.PrivKey, d.cfg.ChainID,
		)
		if err != nil {
			return nil, err
		}
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	blockOffset := new(big.Int).SetUint64(d.cfg.BlockOffset)
	offsetStartsAtIndex := new(big.Int).Sub(start, blockOffset)
	// Assembly data request tss node signature
	tssResponse, err := d.RequestTssSignature(0, start, offsetStartsAtIndex, "", stateRoots)
	if err != nil {
		log.Error(name+" get tss manager signature fail", "err", err)
		return nil, err
	}
	log.Info(name+" append log", "stateRoots size ", len(stateRoots), "offsetStartsAtIndex", offsetStartsAtIndex, "signature", hex.EncodeToString(tssResponse.Signature), "rollback", tssResponse.RollBack)
	log.Info(name+" signature ", "len", len(tssResponse.Signature))
	var tx *types.Transaction
	if tssResponse.RollBack {
		d.Metrics().TssRollbackSignal().Inc()
		log.Error("tssResponse indicate a layer2 rollback, look up this!!!")
		if d.cfg.AllowL2AutoRollback {
			log.Info("l2geth trigger auto rollback")
			if d.rollbackEndBlock.Cmp(end) <= 0 && d.rollbackEndBlock.Cmp(start) > 0 {
				tempS := stateRoots[d.rollbackEndBlock.Uint64()-start.Uint64()-1]
				if bytes.Equal(tempS[:], d.rollbackEndStateRoot[:]) {
					err = errors.New("l2geth is still rollback")
					log.Error(name + " still waiting l2geth rollback result")
					return nil, err
				}
			}
			d.rollbackEndStateRoot = stateRoots[len(stateRoots)-1]
			d.rollbackEndBlock = end
			log.Info("sending l2geth rollback transaction")
			tx, err = d.fpRollup.RollbackL2Chain(opts, start, offsetStartsAtIndex, tssResponse.Signature)
		}
	} else {
		if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
			log.Info("append state with fraud proof")
			// ##### FRAUD-PROOF modify #####
			// check stake initialised
			owner, _ := d.fpRollup.Owner(&bind.CallOpts{})
			assertionInit, _ := d.fpAssertion.Assertions(&bind.CallOpts{}, new(big.Int).SetUint64(0))
			if !bytes.Equal(owner.Bytes(), opts.From.Bytes()) || bytes.Equal(assertionInit.Deadline.Bytes(), common.Big0.Bytes()) {
				log.Error("fraud proof not init", "owner", owner, "expect", opts.From)
				return nil, nil
			}
			// Append state batch
			tx, err = d.FraudProofAppendStateBatch(
				opts, stateRoots, offsetStartsAtIndex, tssResponse.Signature, blocks,
			)
			if err != nil {
				log.Error("fraud proof append state batch is failed", "err", err)
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
			if d.cfg.AllowL2AutoRollback {
				return d.fpRollup.RollbackL2Chain(opts, start, offsetStartsAtIndex, tssResponse.Signature)
			} else {
				return nil, nil
			}
		} else {
			if len(d.cfg.FPRollupAddr.Bytes()) != 0 {
				log.Info("append state with fraud proof by gas tip cap")
				// ##### FRAUD-PROOF modify #####
				// check stake initialised
				// check is in challenge status
				challengeContext, err := d.fpRollup.ChallengeCtx(&bind.CallOpts{})
				if err != nil {
					return nil, err
				}
				isInChallenge := challengeContext.DefenderAssertionID.Uint64() != 0 && !challengeContext.Completed
				if isInChallenge {
					log.Warn("currently in challenge, can't submit new assertion")
					return nil, nil
				}
				// check rollback status
				if d.cfg.SccRollback {
					fpChallenge, err := fpbindings.NewChallenge(
						challengeContext.ChallengeAddress, d.cfg.L1Client,
					)
					if err != nil {
						return nil, err
					}
					alreadyRollback, err := fpChallenge.Rollback(&bind.CallOpts{})
					if err != nil {
						return nil, err
					}
					challenger, err := fpChallenge.Challenger(&bind.CallOpts{})
					if err != nil {
						return nil, err
					}
					winner, err := fpChallenge.Winner(&bind.CallOpts{})
					if err != nil {
						return nil, err
					}
					startInboxSize, err := fpChallenge.StartInboxSize(&bind.CallOpts{})
					if err != nil {
						return nil, err
					}
					if challengeContext.Completed && !alreadyRollback && bytes.Equal(challenger[:], winner[:]) {
						tssResponse, err = d.RequestTssSignature(1, startInboxSize, offsetStartsAtIndex, challengeContext.ChallengeAddress.String(), nil)
						if err != nil {
							return nil, err
						}
						if tssResponse.RollBack {
							// delete scc batch states one by one
							totalBatches, err := d.sccContract.GetTotalBatches(&bind.CallOpts{})
							if err != nil {
								return nil, err
							}

							filter, err := d.sccContract.FilterStateBatchAppended(&bind.FilterOpts{}, []*big.Int{totalBatches})
							if err != nil {
								return nil, err
							}
							if filter.Event.PrevTotalElements.Cmp(startInboxSize) <= 0 {
								var rollbackTx *types.Transaction
								var rollbackErr error
								// must ensure all those action done properly until rollback finished
								// or RollBackL2Chain will happen multiple times
								d.once.Do(
									func() {
										rollbackTx, rollbackErr = d.fpRollup.RollbackL2Chain(
											opts, startInboxSize, offsetStartsAtIndex, tssResponse.Signature,
										)
									},
								)
								if rollbackTx != nil || rollbackErr != nil {
									return rollbackTx, rollbackErr
								} else {
									return fpChallenge.SetRollback(opts)
								}
							}
							return d.fpRollup.RejectLatestCreatedAssertionWithBatch(opts, fpbindings.LibBVMCodecChainBatchHeader{
								BatchIndex:        filter.Event.BatchIndex,
								BatchRoot:         filter.Event.BatchRoot,
								BatchSize:         filter.Event.BatchSize,
								PrevTotalElements: filter.Event.PrevTotalElements,
								Signature:         filter.Event.Signature,
								ExtraData:         filter.Event.ExtraData,
							})
						}
					}
				}
				// rollup assertion
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

	var opts *bind.TransactOpts
	if d.cfg.EnableProposerHsm {
		proBytes, err := hex.DecodeString(d.cfg.ProposerHsmCreden)
		apikey := option.WithCredentialsJSON(proBytes)
		client, err := kms.NewKeyManagementClient(ctx, apikey)
		if err != nil {
			return nil, err
		}
		mk := &bsscore.ManagedKey{
			KeyName:      d.cfg.ProposerHsmAPIName,
			EthereumAddr: common.HexToAddress(d.cfg.ProposerHsmAddress),
			Gclient:      client,
		}
		opts, err = mk.NewEthereumTransactorrWithChainID(ctx, d.cfg.ChainID)
		log.Info("proposer", "enable-hsm", true)
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
	var latestAssertion rollupTypes.Assertion
	var stakerInfo rollupTypes.Staker
	var stakerAddr common.Address
	if ret, err := d.fpRollup.Registers(&bind.CallOpts{}, opts.From); err != nil {
		return nil, err
	} else {
		stakerAddr = ret
	}
	if ret, err := d.fpRollup.Stakers(&bind.CallOpts{}, stakerAddr); err != nil {
		return nil, err
	} else {
		stakerInfo.IsStaked = ret.IsStaked
		stakerInfo.AmountStaked = ret.AmountStaked
		stakerInfo.AssertionID = ret.AssertionID
		stakerInfo.CurrentChallenge = ret.CurrentChallenge
	}
	if ret, err := d.fpAssertion.Assertions(&bind.CallOpts{}, stakerInfo.AssertionID); err != nil {
		return nil, err
	} else {
		latestAssertion.ID = stakerInfo.AssertionID
		latestAssertion.VmHash = ret.StateHash
		latestAssertion.InboxSize = ret.InboxSize
		latestAssertion.Parent = ret.Parent
		latestAssertion.Deadline = ret.Deadline
		latestAssertion.ProposalTime = ret.ProposalTime
	}

	txBatch := rollupTypes.NewTxBatch(blocks, uint64(len(blocks)))

	// First assertion check
	lastCreatedAssertionID, err := d.fpRollup.LastCreatedAssertionID(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	if lastCreatedAssertionID.Uint64() != 0 && latestAssertion.InboxSize.Uint64()+uint64(len(txBatch.Txs)) != txBatch.LastBlockNumber() {
		log.Error("Online total InboxSize not match with local batch's LatestBlockNumber")
		log.Info(fmt.Sprintf("show proposer error, currenyInboxSize: %d, batchLength: %d, lastBlockNumber: %d", latestAssertion.InboxSize.Uint64(), len(txBatch.Txs), txBatch.LastBlockNumber()))
		return nil, errors.New("Online total InboxSize not match with local batch's LatestBlockNumber")
	}

	assertion := txBatch.ToAssertion(&latestAssertion)

	log.Debug("show assertion", "VmHash", assertion.VmHash.String())
	log.Debug("show assertion", "InboxSize", assertion.InboxSize)
	log.Debug("show assertion", "Batch", batch)
	log.Debug("show assertion", "ShouldStartAtElement", shouldStartAtElement.String())
	log.Debug("show assertion", "Signature", hex.EncodeToString(signature))

	// create assertion
	return d.fpRollup.CreateAssertionWithStateBatch(
		opts, assertion.VmHash, assertion.InboxSize, batch, shouldStartAtElement, signature)
}

func (d *Driver) RequestTssSignature(requestType uint64, start, offsetStartsAtIndex *big.Int, challenge string, stateRoots [][stateRootSize]byte) (*tssClient.TssResponse, error) {
	var tssResponse tssClient.TssResponse
	tssReqParams := tss_types.SignStateRequest{
		Type:                requestType,
		StartBlock:          start,
		OffsetStartsAtIndex: offsetStartsAtIndex,
		Challenge:           challenge,
		StateRoots:          stateRoots,
	}
	tssReponseBytes, err := d.cfg.TssClient.GetSignStateBatch(tssReqParams)
	if err != nil {
		log.Error("get tss manager signature fail", "err", err)
		return nil, err
	}
	err = json.Unmarshal(tssReponseBytes, &tssResponse)
	if err != nil {
		log.Error("failed to unmarshal response from tss", "err", err)
		return nil, err
	}
	return &tssResponse, nil
}
