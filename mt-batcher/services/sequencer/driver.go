package sequencer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceDL"
	"github.com/Layr-Labs/datalayr/common/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	l2gethcommon "github.com/mantlenetworkio/mantle/l2geth/common"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	common3 "github.com/mantlenetworkio/mantle/l2geth/rollup/eigenda"
	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	rc "github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/metrics"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/client"
	common4 "github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/sequencer/db"
	"github.com/mantlenetworkio/mantle/mt-batcher/txmgr"
)

var (
	pollingInterval = 1000 * time.Millisecond
)

type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

type DriverConfig struct {
	L1Client                  *ethclient.Client
	L2Client                  *l2ethclient.Client
	L1ChainID                 *big.Int
	DtlClientUrl              string
	EigenDaContract           *bindings.BVMEigenDataLayrChain
	RawEigenContract          *bind.BoundContract
	EigenFeeContract          *bindings.BVMEigenDataLayrFee
	RawEigenFeeContract       *bind.BoundContract
	Logger                    *logging.Logger
	PrivKey                   *ecdsa.PrivateKey
	FeePrivKey                *ecdsa.PrivateKey
	BlockOffset               uint64
	RollUpMinTxn              uint64
	RollUpMaxSize             uint64
	EigenLayerNode            int
	DataStoreDuration         uint64
	DataStoreTimeout          uint64
	DisperserSocket           string
	FeeWorkerPollInterval     time.Duration
	MainWorkerPollInterval    time.Duration
	CheckerWorkerPollInterval time.Duration
	GraphPollingDuration      time.Duration
	GraphProvider             string
	EigenLogConfig            logging.Config
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	DbPath                    string
	CheckerBatchIndex         uint64
	CheckerEnable             bool
	FeeSizeSec                string
	FeePerBytePerTime         uint64
	FeeModelEnable            bool
	MinTimeoutRollupTxn       uint64
	RollupTimeout             time.Duration
	Metrics                   metrics.MtBatchMetrics

	EnableHsm     bool
	HsmAddress    string
	HsmFeeAddress string
	HsmAPIName    string
	HsmFeeAPIName string
	HsmCreden     string
}

type FeePipline struct {
	RollUpFee        *big.Int
	EndL2BlockNumber *big.Int
}

type Driver struct {
	Ctx           context.Context
	Cfg           *DriverConfig
	WalletAddr    common.Address
	FeeWalletAddr common.Address
	GraphClient   *graphView.GraphClient
	DtlClient     client.DtlClient
	txMgr         txmgr.TxManager
	LevelDBStore  *db.Store
	FeeCh         chan *FeePipline
	cancel        func()
	wg            sync.WaitGroup
}

var bigOne = new(big.Int).SetUint64(1)

func NewDriver(ctx context.Context, cfg *DriverConfig) (*Driver, error) {
	_, cancel := context.WithTimeout(ctx, common4.DefaultTimeout)
	defer cancel()
	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, cfg.L1Client)

	graphClient := graphView.NewGraphClient(cfg.GraphProvider, cfg.Logger)

	levelDBStore, err := db.NewStore(cfg.DbPath)
	if err != nil {
		log.Error("init leveldb fail", "err", err)
		return nil, err
	}
	dtlClient := client.NewDtlClient(cfg.DtlClientUrl)
	var walletAddr, feeWalletAddr common.Address
	if cfg.EnableHsm {
		walletAddr = common.HexToAddress(cfg.HsmAddress)
		feeWalletAddr = common.HexToAddress(cfg.HsmFeeAddress)
	} else {
		walletAddr = crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)
		feeWalletAddr = crypto.PubkeyToAddress(cfg.FeePrivKey.PublicKey)
	}
	return &Driver{
		Cfg:           cfg,
		Ctx:           ctx,
		WalletAddr:    walletAddr,
		FeeWalletAddr: feeWalletAddr,
		GraphClient:   graphClient,
		DtlClient:     dtlClient,
		txMgr:         txMgr,
		LevelDBStore:  levelDBStore,
		FeeCh:         make(chan *FeePipline),
		cancel:        cancel,
	}, nil
}

func (d *Driver) UpdateGasPrice(ctx context.Context, tx *types.Transaction, feeModelEnable bool) (*types.Transaction, error) {
	var finalTx *types.Transaction
	var err error
	var opts *bind.TransactOpts
	if feeModelEnable {
		if !d.Cfg.EnableHsm {
			opts, err = bind.NewKeyedTransactorWithChainID(
				d.Cfg.FeePrivKey, d.Cfg.L1ChainID,
			)
		} else {
			opts, err = common4.NewHSMTransactOpts(ctx, d.Cfg.HsmFeeAPIName,
				d.Cfg.HsmFeeAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
		}
		if err != nil {
			return nil, err
		}
		opts.Context = ctx
		opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
		opts.NoSend = true
	} else {
		if !d.Cfg.EnableHsm {
			opts, err = bind.NewKeyedTransactorWithChainID(
				d.Cfg.PrivKey, d.Cfg.L1ChainID,
			)
		} else {
			opts, err = common4.NewHSMTransactOpts(ctx, d.Cfg.HsmAPIName,
				d.Cfg.HsmAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
		}
		if err != nil {
			return nil, err
		}
		opts.Context = ctx
		opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
		opts.NoSend = true
	}
	if feeModelEnable {
		log.Debug("MtBatcher update eigen da use fee", "FeeModelEnable", d.Cfg.FeeModelEnable)
		finalTx, err = d.Cfg.RawEigenFeeContract.RawTransact(opts, tx.Data())
	} else {
		log.Debug("MtBatcher rollup data", "FeeModelEnable", d.Cfg.FeeModelEnable)
		finalTx, err = d.Cfg.RawEigenContract.RawTransact(opts, tx.Data())
	}
	switch {
	case err == nil:
		return finalTx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = common4.FallbackGasTipCap
		if feeModelEnable {
			log.Debug("update eigen da use fee", "FeeModelEnable", d.Cfg.FeeModelEnable)
			return d.Cfg.RawEigenFeeContract.RawTransact(opts, tx.Data())
		} else {
			log.Debug("rollup date", "FeeModelEnable", d.Cfg.FeeModelEnable)
			return d.Cfg.RawEigenContract.RawTransact(opts, tx.Data())
		}
	default:
		return nil, err
	}
}

func (d *Driver) GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, error) {
	blockOffset := new(big.Int).SetUint64(d.Cfg.BlockOffset)
	var end *big.Int
	log.Debug("MtBatcher GetBatchBlockRange", "blockOffset", blockOffset)
	start, err := d.Cfg.EigenDaContract.GetL2ConfirmedBlockNumber(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return nil, nil, err
	}
	latestHeader, err := d.Cfg.L2Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	end = new(big.Int).Add(start, big.NewInt(int64(d.Cfg.BlockOffset)))
	if start.Cmp(end) > 0 {
		return nil, nil, fmt.Errorf("invalid range, end(%v) < start(%v)", end, start)
	}
	if end.Cmp(latestHeader.Number) > 0 {
		end = latestHeader.Number
	}
	return start, end, nil
}

func (d *Driver) TxAggregator(ctx context.Context, start, end *big.Int) (transactionData []byte, startL2BlockNumber *big.Int, endL2BlockNumber *big.Int) {
	var batchTxList []common3.BatchTx
	var transactionByte []byte
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		block, err := d.Cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			log.Error("get blockNumber from l2 fail", "err", err)
			continue
		}
		txs := block.Transactions()
		if len(txs) != 1 {
			panic(fmt.Sprintf("MtBatcher attempting to create batch element from block %d, "+
				"found %d txs instead of 1", block.Number(), len(txs)))
		}
		log.Debug("MtBatcher origin transactions", "TxHash", txs[0].Hash().String(), "l2BlockNumber", block.Number(), "QueueOrigin", txs[0].QueueOrigin(), "Index", *txs[0].GetMeta().Index, "QueueIndex", txs[0].GetMeta().QueueIndex, "i", i)
		var txBuf bytes.Buffer
		if err := txs[0].EncodeRLP(&txBuf); err != nil {
			panic(fmt.Sprintf("MtBatcher Unable to encode tx: %v", err))
		}
		var l1MessageSender *l2gethcommon.Address
		if txs[0].GetMeta().QueueIndex != nil {
			l1Origin, err := d.DtlClient.GetEnqueueByIndex(*txs[0].GetMeta().QueueIndex)
			if err != nil {
				l1MessageSender = txs[0].GetMeta().L1MessageSender
			} else {
				originAddress := l2gethcommon.HexToAddress(l1Origin)
				l1MessageSender = &originAddress
			}
		} else {
			l1MessageSender = txs[0].GetMeta().L1MessageSender
		}
		txMeta := &common3.TransactionMeta{
			L1BlockNumber:   txs[0].GetMeta().L1BlockNumber,
			L1Timestamp:     txs[0].GetMeta().L1Timestamp,
			L1MessageSender: l1MessageSender,
			Index:           txs[0].GetMeta().Index,
			QueueIndex:      txs[0].GetMeta().QueueIndex,
			RawTransaction:  txs[0].GetMeta().RawTransaction,
		}
		txMetaByte, err := json.Marshal(txMeta)
		if err != nil {
			log.Error("tx meta json marshal error", "err", err)
			continue
		}
		batchTx := common3.BatchTx{
			BlockNumber: i.Bytes(),
			TxMeta:      txMetaByte,
			RawTx:       txBuf.Bytes(),
		}
		batchTxList = append(batchTxList, batchTx)
		txnBufBytes, err := l2rlp.EncodeToBytes(batchTxList)
		if err != nil {
			panic(fmt.Sprintf("MtBatcher Unable to encode txn: %v", err))
		}
		if uint64(len(txnBufBytes)) >= d.Cfg.RollUpMaxSize {
			log.Info("MtBatcher batch size more than RollUpMaxSize, real rollup data", "RollUpMaxSize", d.Cfg.RollUpMaxSize, "start", start, "end", i)
			return transactionByte, start, i
		} else {
			transactionByte = txnBufBytes
		}
	}
	txnBufBytes, err := l2rlp.EncodeToBytes(batchTxList)
	if err != nil {
		panic(fmt.Sprintf("MtBatcher Unable to encode txn: %v", err))
	}
	var totalNode int
	daNodes, err := d.GetEigenLayerNode()
	if err != nil {
		log.Error("get da node fail", "err", err)
		totalNode = d.Cfg.EigenLayerNode
	} else {
		log.Info("MtBatcher current da node", "totalNode", daNodes)
		totalNode = daNodes
	}
	d.Cfg.Metrics.NumEigenNode().Set(float64(daNodes))
	if len(txnBufBytes) > 31*totalNode {
		transactionByte = txnBufBytes
	} else {
		paddingBytes := make([]byte, (31*d.Cfg.EigenLayerNode)-len(txnBufBytes))
		transactionByte = append(txnBufBytes, paddingBytes...)
	}
	return transactionByte, start, end
}

func (d *Driver) StoreData(ctx context.Context, uploadHeader []byte, duration uint8, blockNumber uint32, startL2BlockNumber *big.Int, endL2BlockNumber *big.Int, totalOperatorsIndex uint32, isReRollup bool) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current balance", "err", err)
		return nil, err
	}
	d.Cfg.Metrics.MtBatchBalanceETH().Set(common4.WeiToEth64(balance))
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if !d.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			d.Cfg.PrivKey, d.Cfg.L1ChainID,
		)
	} else {
		opts, err = common4.NewHSMTransactOpts(ctx, d.Cfg.HsmAPIName,
			d.Cfg.HsmAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	tx, err := d.Cfg.EigenDaContract.StoreData(opts, uploadHeader, duration, blockNumber, startL2BlockNumber, endL2BlockNumber, totalOperatorsIndex, isReRollup)
	switch {
	case err == nil:
		return tx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = common4.FallbackGasTipCap
		return d.Cfg.EigenDaContract.StoreData(opts, uploadHeader, duration, blockNumber, startL2BlockNumber, endL2BlockNumber, totalOperatorsIndex, isReRollup)

	default:
		return nil, err
	}
}

func (d *Driver) ConfirmData(ctx context.Context, callData []byte, searchData rc.IDataLayrServiceManagerDataStoreSearchData, startL2BlockNumber, endL2BlockNumber *big.Int, originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current balance", "err", err)
		return nil, err
	}
	d.Cfg.Metrics.MtBatchBalanceETH().Set(common4.WeiToEth64(balance))
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current nonce", "err", err)
		return nil, err
	}
	d.Cfg.Metrics.MtBatchNonce().Set(float64(nonce64))
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if !d.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			d.Cfg.PrivKey, d.Cfg.L1ChainID,
		)
	} else {
		opts, err = common4.NewHSMTransactOpts(ctx, d.Cfg.HsmAPIName,
			d.Cfg.HsmAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	tx, err := d.Cfg.EigenDaContract.ConfirmData(opts, callData, searchData, startL2BlockNumber, endL2BlockNumber, originDataStoreId, reConfirmedBatchIndex, isReRollup)
	switch {
	case err == nil:
		return tx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = common4.FallbackGasTipCap
		return d.Cfg.EigenDaContract.ConfirmData(opts, callData, searchData, startL2BlockNumber, endL2BlockNumber, originDataStoreId, reConfirmedBatchIndex, isReRollup)

	default:
		return nil, err
	}
}

func (d *Driver) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return d.Cfg.L1Client.SendTransaction(ctx, tx)
}

func (d *Driver) DisperseStoreData(data []byte, startl2BlockNumber *big.Int, endl2BlockNumber *big.Int, isReRollup bool) (common2.StoreParams, *types.Receipt, error) {
	params, err := d.callEncode(data)
	if err != nil {
		return params, nil, err
	}
	uploadHeader, err := common2.CreateUploadHeader(params)
	if err != nil {
		return params, nil, err
	}
	log.Info("Operator Info", "NumSys", params.NumSys, "NumPar", params.NumPar, "TotalOperatorsIndex", params.TotalOperatorsIndex, "NumTotal", params.NumTotal)
	tx, err := d.StoreData(
		d.Ctx, uploadHeader, uint8(params.Duration), params.ReferenceBlockNumber, startl2BlockNumber, endl2BlockNumber, params.TotalOperatorsIndex, isReRollup,
	)
	if err != nil {
		log.Error("MtBatcher StoreData tx", "err", err)
		return params, nil, err
	} else if tx == nil {
		return params, nil, errors.New("tx is nil")
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return d.UpdateGasPrice(ctx, tx, false)
	}
	receipt, err := d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	if err != nil {
		log.Error("MtBatcher unable to StoreData", "err", err)
		return params, nil, err
	}
	return params, receipt, nil
}

func (d *Driver) GetEigenLayerNode() (int, error) {
	operators, err := d.GraphClient.QueryOperatorsByStatus()
	if err != nil {
		log.Error("MtBatcher query operators fail", "err", err)
		return 0, err
	}
	return len(operators), nil
}

func (d *Driver) ConfirmStoredData(txHash []byte, params common2.StoreParams, startl2BlockNumber, endl2BlockNumber *big.Int, originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool) (*types.Receipt, error) {
	event, ok := d.GraphClient.PollingInitDataStore(
		d.Ctx,
		txHash[:],
		d.Cfg.GraphPollingDuration,
	)
	if !ok {
		log.Error("MtBatcher could not get initDataStore", "ok", ok)
		return nil, errors.New("MtBatcher could not get initDataStore")
	}
	log.Debug("PollingInitDataStore", "MsgHash", event.MsgHash, "StoreNumber", event.StoreNumber)
	meta, err := d.callDisperse(
		params.HeaderHash,
		event.MsgHash[:],
	)
	if err != nil {
		log.Error("MtBatcher call Disperse fail", "err", err)
		return nil, err
	}
	callData := common2.MakeCalldata(params, meta, event.StoreNumber, event.MsgHash)
	searchData := rc.IDataLayrServiceManagerDataStoreSearchData{
		Duration:  event.Duration,
		Timestamp: new(big.Int).SetUint64(uint64(event.InitTime)),
		Index:     event.Index,
		Metadata: rc.IDataLayrServiceManagerDataStoreMetadata{
			HeaderHash:           event.DataCommitment,
			DurationDataStoreId:  event.DurationDataStoreId,
			GlobalDataStoreId:    event.StoreNumber,
			ReferenceBlockNumber: event.ReferenceBlockNumber,
			BlockNumber:          uint32(event.InitBlockNumber.Uint64()),
			Fee:                  event.Fee,
			Confirmer:            common.HexToAddress(event.Confirmer),
			SignatoryRecordHash:  [32]byte{},
		},
	}
	var tx *types.Transaction
	if !isReRollup {
		tx, err = d.ConfirmData(d.Ctx, callData, searchData, startl2BlockNumber, endl2BlockNumber, originDataStoreId, reConfirmedBatchIndex, isReRollup)
		if err != nil {
			return nil, err
		}
	} else {
		// originDataStoreId uint32, reConfirmedBatchIndex *big.Int, isReRollup bool
		tx, err = d.ConfirmData(d.Ctx, callData, searchData, startl2BlockNumber, endl2BlockNumber, originDataStoreId, reConfirmedBatchIndex, isReRollup)
		if err != nil {
			return nil, err
		}
	}

	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return d.UpdateGasPrice(ctx, tx, false)
	}
	receipt, err := d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	if err != nil {
		log.Error("MtBatcher unable to ConfirmData tx", "err", err)
		return nil, err
	}
	return receipt, nil
}

func (d *Driver) callEncode(data []byte) (common2.StoreParams, error) {
	conn, err := grpc.Dial(d.Cfg.DisperserSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("MtBatcher Disperser Cannot connect to", "DisperserSocket", d.Cfg.DisperserSocket)
		return common2.StoreParams{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.Cfg.DataStoreTimeout))
	defer cancel()
	request := &pb.EncodeStoreRequest{
		Duration: d.Cfg.DataStoreDuration,
		Data:     data,
	}
	opt := grpc.MaxCallSendMsgSize(1024 * 1024 * 300)
	reply, err := c.EncodeStore(ctx, request, opt)
	if err != nil {
		log.Error("MtBatcher get store err", err)
		return common2.StoreParams{}, err
	}
	g := reply.GetStore()
	feeBigInt := new(big.Int).SetBytes(g.Fee)
	params := common2.StoreParams{
		ReferenceBlockNumber: g.ReferenceBlockNumber,
		TotalOperatorsIndex:  g.TotalOperatorsIndex,
		OrigDataSize:         g.OrigDataSize,
		NumTotal:             g.NumTotal,
		Quorum:               g.Quorum,
		NumSys:               g.NumSys,
		NumPar:               g.NumPar,
		Duration:             g.Duration,
		KzgCommit:            g.KzgCommit,
		LowDegreeProof:       g.LowDegreeProof,
		Degree:               g.Degree,
		TotalSize:            g.TotalSize,
		Order:                g.Order,
		Fee:                  feeBigInt,
		HeaderHash:           g.HeaderHash,
		Disperser:            g.Disperser,
	}
	return params, nil
}

func (d *Driver) callDisperse(headerHash []byte, messageHash []byte) (common2.DisperseMeta, error) {
	conn, err := grpc.Dial(d.Cfg.DisperserSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("MtBatcher Dial DisperserSocket", "err", err)
		return common2.DisperseMeta{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.Cfg.DataStoreTimeout))
	defer cancel()
	request := &pb.DisperseStoreRequest{
		HeaderHash:  headerHash,
		MessageHash: messageHash,
	}
	reply, err := c.DisperseStore(ctx, request)
	if err != nil {
		return common2.DisperseMeta{}, err
	}
	sigs := reply.GetSigs()
	aggSig := common2.AggregateSignature{
		AggSig:            sigs.AggSig,
		StoredAggPubkeyG1: sigs.StoredAggPubkeyG1,
		UsedAggPubkeyG2:   sigs.UsedAggPubkeyG2,
		NonSignerPubkeys:  sigs.NonSignerPubkeys,
	}
	meta := common2.DisperseMeta{
		Sigs:            aggSig,
		ApkIndex:        reply.GetApkIndex(),
		TotalStakeIndex: reply.GetTotalStakeIndex(),
	}
	return meta, nil
}

func (d *Driver) CalcUserFeeByRules(rollupDateSize *big.Int) (*big.Int, error) {
	seconds := new(big.Int).Div(big.NewInt(int64(d.Cfg.MainWorkerPollInterval)), big.NewInt(1000000000))
	rollSizeSec := new(big.Int).Div(rollupDateSize, seconds)
	feeSs, ok := new(big.Int).SetString(d.Cfg.FeeSizeSec, 10)
	if !ok {
		log.Error("FeeSizeSec from string to big.int fail")
		return big.NewInt(0), nil
	}
	if rollSizeSec.Cmp(feeSs) < 0 {
		return big.NewInt(0), nil
	}
	return new(big.Int).Mul(new(big.Int).Div(rollSizeSec, feeSs), new(big.Int).SetUint64(d.Cfg.FeePerBytePerTime)), nil
}

func (d *Driver) UpdateFee(ctx context.Context, l2Block, daFee *big.Int) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.FeeWalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get fee wallet address current balance", "err", err)
		return nil, err
	}
	d.Cfg.Metrics.MtFeeBalanceETH().Set(common4.WeiToEth64(balance))
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.FeeWalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get fee wallet nonce", "err", err)
		return nil, err
	}
	d.Cfg.Metrics.MtFeeNonce().Set(float64(nonce64))
	nonce := new(big.Int).SetUint64(nonce64)
	var opts *bind.TransactOpts
	if !d.Cfg.EnableHsm {
		opts, err = bind.NewKeyedTransactorWithChainID(
			d.Cfg.FeePrivKey, d.Cfg.L1ChainID,
		)
	} else {
		opts, err = common4.NewHSMTransactOpts(ctx, d.Cfg.HsmFeeAPIName,
			d.Cfg.HsmFeeAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
	}
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.NoSend = true

	tx, err := d.Cfg.EigenFeeContract.SetRollupFee(opts, l2Block, daFee)
	switch {
	case err == nil:
		return tx, nil
	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = common4.FallbackGasTipCap
		return d.Cfg.EigenFeeContract.SetRollupFee(opts, l2Block, daFee)
	default:
		return nil, err
	}
}

func (d *Driver) UpdateUserDaFee(l2Block, daFee *big.Int) (*types.Receipt, error) {
	tx, err := d.UpdateFee(
		d.Ctx, l2Block, daFee,
	)
	if err != nil {
		log.Error("MtBatcher Update User fee fail", "err", err)
		return nil, err
	} else if tx == nil {
		return nil, errors.New("tx is nil")
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return d.UpdateGasPrice(ctx, tx, true)
	}
	receipt, err := d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	if err != nil {
		log.Error("MtBatcher unable to StoreData", "err", err)
		return nil, err
	}
	return receipt, nil
}

func (d *Driver) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), common4.ErrMaxPriorityFeePerGasNotFound.Error(),
	)
}

func (d *Driver) ServiceInit() error {
	rollupWalletBalance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Warn("Get rollup wallet address balance fail", "err", err)
		return err
	}
	d.Cfg.Metrics.MtBatchBalanceETH().Set(common4.WeiToEth64(rollupWalletBalance))

	rollupNonce, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Warn("Get rollup wallet address nonce fail", "err", err)
		return err
	}
	d.Cfg.Metrics.MtBatchNonce().Set(float64(rollupNonce))

	feeWalletBalance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.FeeWalletAddr, nil,
	)
	if err != nil {
		log.Warn("Get rollup fee wallet address balance fail", "err", err)
		return err
	}
	d.Cfg.Metrics.MtFeeBalanceETH().Set(common4.WeiToEth64(feeWalletBalance))

	feeNonce, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Warn("Get rollup fee wallet address nonce fail", "err", err)
		return err
	}
	d.Cfg.Metrics.MtFeeNonce().Set(float64(feeNonce))
	return nil
}

func (d *Driver) Start() error {
	d.wg.Add(1)
	go d.RollupMainWorker()
	err := d.ServiceInit()
	if err != nil {
		log.Error("init metrics fail", "err", err)
		return err
	}
	d.Cfg.Metrics.RollupTimeDuration().Set(float64(d.Cfg.MainWorkerPollInterval))
	if d.Cfg.CheckerEnable {
		batchIndex, ok := d.LevelDBStore.GetReRollupBatchIndex()
		if batchIndex == 0 || !ok {
			d.LevelDBStore.SetReRollupBatchIndex(1)
		}
		d.wg.Add(1)
		d.Cfg.Metrics.CheckerTimeDuration().Set(float64(d.Cfg.CheckerWorkerPollInterval))
		go d.CheckConfirmedWorker()
	}
	if d.Cfg.FeeModelEnable {
		d.wg.Add(1)
		d.Cfg.Metrics.FeeTimeDuration().Set(float64(d.Cfg.FeeWorkerPollInterval))
		go d.RollUpFeeWorker()
	}
	return nil
}

func (d *Driver) Stop() {
	d.cancel()
	d.wg.Wait()
}

func (d *Driver) GetBatchBlockRangeWithTimeout(ctx context.Context) (*big.Int, *big.Int, error) {
	log.Debug("RollupTimeInterval start")
	rollupTimout := d.Cfg.RollupTimeout
	exit := time.NewTimer(rollupTimout)
	ticker := time.NewTicker(pollingInterval)
	for {
		start, end, err := d.GetBatchBlockRange(d.Ctx)
		if err != nil {
			return nil, nil, err
		}
		select {
		case <-ticker.C:
			if big.NewInt(0).Sub(end, start).Cmp(big.NewInt(int64(d.Cfg.RollUpMinTxn))) >= 0 {
				return start, end, nil
			}
			if start.Cmp(end) == 0 {
				log.Info("MtBatcher Sequencer no updates", "start", start, "end", end)
				continue
			}
		case <-exit.C:
			if big.NewInt(0).Sub(end, start).Cmp(big.NewInt(int64(d.Cfg.MinTimeoutRollupTxn))) >= 0 {
				return start, end, nil
			}
			return nil, nil, errors.Errorf("error: Timeout txn less than MinTimeoutRollupTxn")
		case err := <-d.Ctx.Done():
			log.Error("MtBatcher get block range timeout error", "err", err)
			return nil, nil, errors.New("MtBatcher get block range timeout error")
		}
	}
}

func (d *Driver) RollupMainWorker() {
	defer d.wg.Done()
	ticker := time.NewTicker(d.Cfg.MainWorkerPollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			start, end, err := d.GetBatchBlockRangeWithTimeout(d.Ctx)
			if err != nil {
				log.Warn("MtBatcher Sequencer unable to get block range", "err", err)
				continue
			}
			log.Info("MtBatcher get batch block range", "start", start, "end", end)
			aggregateTxData, startL2BlockNumber, endL2BlockNumber := d.TxAggregator(
				d.Ctx, start, end,
			)
			if err != nil {
				log.Error("MtBatcher eigenDa sequencer unable to craft batch tx", "err", err)
				continue
			}
			d.Cfg.Metrics.NumTxnPerBatch().Observe(float64((new(big.Int).Sub(endL2BlockNumber, startL2BlockNumber)).Uint64()))
			d.Cfg.Metrics.BatchSizeBytes().Observe(float64(len(aggregateTxData)))
			params, receipt, err := d.DisperseStoreData(aggregateTxData, startL2BlockNumber, endL2BlockNumber, false)
			if err != nil {
				log.Error("MtBatcher disperse store data fail", "err", err)
				continue
			}
			log.Info("MtBatcher disperse store data success", "txHash", receipt.TxHash.String())
			d.Cfg.Metrics.L2StoredBlockNumber().Set(float64(start.Uint64()))
			time.Sleep(10 * time.Second) // sleep for data into graph node
			csdReceipt, err := d.ConfirmStoredData(receipt.TxHash.Bytes(), params, startL2BlockNumber, endL2BlockNumber, 0, big.NewInt(0), false)
			if err != nil {
				log.Error("MtBatcher confirm store data fail", "err", err)
				continue
			}
			log.Info("MtBatcher confirm store data success", "txHash", csdReceipt.TxHash.String())
			d.Cfg.Metrics.L2ConfirmedBlockNumber().Set(float64(start.Uint64()))
			if d.Cfg.FeeModelEnable {
				daFee, _ := d.CalcUserFeeByRules(big.NewInt(int64(len(aggregateTxData))))
				feePip := &FeePipline{
					RollUpFee:        daFee,
					EndL2BlockNumber: endL2BlockNumber,
				}
				d.FeeCh <- feePip
			}
			batchIndex, _ := d.Cfg.EigenDaContract.RollupBatchIndex(&bind.CallOpts{})
			d.Cfg.Metrics.RollUpBatchIndex().Set(float64(batchIndex.Uint64()))

		case err := <-d.Ctx.Done():
			log.Error("MtBatcher eigenDa sequencer service shutting down", "err", err)
			return
		}
	}
}

func (d *Driver) RollUpFeeWorker() {
	defer d.wg.Done()
	ticker := time.NewTicker(d.Cfg.FeeWorkerPollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if d.Cfg.FeeModelEnable {
				chainFee, err := d.Cfg.EigenFeeContract.GetRollupFee(&bind.CallOpts{})
				if err != nil {
					log.Error("MtBatcher RollUpFeeWorker get chain fee fail", "err", err)
					continue
				}
				daFee := <-d.FeeCh
				log.Info("MtBatcher RollUpFeeWorker chainFee and daFee", "chainFee", chainFee, "daFee", *daFee)
				if chainFee.Cmp(daFee.RollUpFee) != 0 {
					txfRpt, err := d.UpdateUserDaFee(daFee.EndL2BlockNumber, daFee.RollUpFee)
					if err != nil {
						log.Error("MtBatcher RollUpFeeWorker update user da fee fail", "err", err)
						continue
					}
					d.Cfg.Metrics.EigenUserFee().Set(float64(daFee.RollUpFee.Uint64()))
					log.Info("MtBatcher RollUpFeeWorker update user fee success", "Hash", txfRpt.TxHash.String())
				}
			}
		case err := <-d.Ctx.Done():
			log.Error("MtBatcher RollUpFeeWorker eigenDa sequencer service shutting down", "err", err)
			return
		}
	}
}

func (d *Driver) CheckConfirmedWorker() {
	defer d.wg.Done()
	ticker := time.NewTicker(d.Cfg.CheckerWorkerPollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestReRollupBatchIndex, err := d.Cfg.EigenDaContract.ReRollupIndex(&bind.CallOpts{})
			if err != nil {
				log.Error("Checker get batch index fail", "err", err)
				continue
			}
			d.Cfg.Metrics.ReRollUpBatchIndex().Set(float64(latestReRollupBatchIndex.Uint64()))
			batchIndex, ok := d.LevelDBStore.GetReRollupBatchIndex()
			if !ok {
				log.Error("Checker get batch index from db fail", "err", err)
				continue
			}

			if batchIndex >= latestReRollupBatchIndex.Uint64() {
				log.Info("Checker db batch index and contract batch idnex is equal", "DbBatchIndex", batchIndex, "latestReRollupBatchIndex", latestReRollupBatchIndex.Uint64())
				continue
			}

			log.Debug("Checker db batch index and contract batch idnex", "DbBatchIndex", batchIndex, "ContractBatchIndex", latestReRollupBatchIndex.Uint64())
			for i := batchIndex; i < latestReRollupBatchIndex.Uint64(); i++ {
				log.Info("Checker batch confirm data index", "batchIndex", i)
				reConfirmedBatchIndex, err := d.Cfg.EigenDaContract.ReRollupBatchIndex(&bind.CallOpts{}, big.NewInt(int64(i)))
				if err != nil {
					log.Info("Checker get batch index by re rollup index fail", "err", err)
					continue
				}

				rollupStore, err := d.Cfg.EigenDaContract.RollupBatchIndexRollupStores(&bind.CallOpts{}, reConfirmedBatchIndex)
				if err != nil {
					log.Info("Checker get rollup store fail", "err", err)
					continue
				}
				if rollupStore.DataStoreId > 0 {
					rollupBlock, err := d.Cfg.EigenDaContract.DataStoreIdToL2RollUpBlock(&bind.CallOpts{}, rollupStore.DataStoreId)
					if err != nil {
						log.Info("Checker get l2 rollup block fail", "err", err)
						continue
					}
					log.Info("Checker DataStoreIdToL2RollUpBlock", "rollupBlock.StartL2BlockNumber", rollupBlock.StartL2BlockNumber, "rollupBlock.EndBL2BlockNumber", rollupBlock.EndBL2BlockNumber)

					aggregateTxData, startL2BlockNumber, endL2BlockNumber := d.TxAggregator(
						d.Ctx, rollupBlock.StartL2BlockNumber, rollupBlock.EndBL2BlockNumber,
					)
					if err != nil {
						log.Error("Checker eigenDa sequencer unable to craft batch tx", "err", err)
						continue
					}
					log.Info("Checker tx aggregator", "startL2BlockNumber", startL2BlockNumber, "endL2BlockNumber", endL2BlockNumber)

					params, receipt, err := d.DisperseStoreData(aggregateTxData, startL2BlockNumber, endL2BlockNumber, true)
					if err != nil {
						log.Error("Checker disperse store data fail", "err", err)
						continue
					}
					time.Sleep(10 * time.Second) // sleep for data into graph node
					csdReceipt, err := d.ConfirmStoredData(receipt.TxHash.Bytes(), params, startL2BlockNumber, endL2BlockNumber, rollupStore.DataStoreId, reConfirmedBatchIndex, true)
					if err != nil {
						log.Error("Checker confirm store data fail", "err", err)
						continue
					}
					log.Info("Checker confirm re-rollup store data success", "txHash", csdReceipt.TxHash.String())
				}
				d.LevelDBStore.SetReRollupBatchIndex(i)
			}
		case err := <-d.Ctx.Done():
			log.Error("MtBatcher eigenDa sequencer service shutting down", "err", err)
			return
		}
	}
}
