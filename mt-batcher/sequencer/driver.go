package sequencer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceDL"
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	rc "github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/txmgr"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/big"
	"strings"
	"sync"
	"time"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

type DriverConfig struct {
	L1Client                  *ethclient.Client
	L2Client                  *l2ethclient.Client
	EigenContractAddr         common.Address
	EigenFeeContractAddress   common.Address
	PrivKey                   *ecdsa.PrivateKey
	BlockOffset               uint64
	RollUpMinSize             uint64
	RollUpMaxSize             uint64
	EigenLayerNode            int
	ChainID                   *big.Int
	DataStoreDuration         uint64
	DataStoreTimeout          uint64
	DisperserSocket           string
	FeeSizeSec                string
	PollInterval              time.Duration
	GraphProvider             string
	EigenLogConfig            logging.Config
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	FeeModelEnable            bool
	SignerFn                  SignerFn
}

type Driver struct {
	Ctx                 context.Context
	Cfg                 *DriverConfig
	EigenDaContract     *bindings.BVMEigenDataLayrChain
	RawEigenContract    *bind.BoundContract
	WalletAddr          common.Address
	EigenABI            *abi.ABI
	EigenFeeContract    *bindings.BVMEigenDataLayrFee
	RawEigenFeeContract *bind.BoundContract
	EigenFeeABI         *abi.ABI
	GraphClient         *graphView.GraphClient
	logger              *logging.Logger
	txMgr               txmgr.TxManager
	cancel              func()
	wg                  sync.WaitGroup
}

var bigOne = new(big.Int).SetUint64(1)

func NewDriver(ctx context.Context, cfg *DriverConfig) (*Driver, error) {
	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		cfg.EigenContractAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("MtBatcher binding eigenda contract fail", "err", err)
		return nil, err
	}
	logger, err := logging.GetLogger(cfg.EigenLogConfig)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrChainABI,
	))
	if err != nil {
		log.Error("MtBatcher parse eigenda contract abi fail", "err", err)
		return nil, err
	}
	eigenABI, err := bindings.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		log.Error("MtBatcher get eigenda contract abi fail", "err", err)
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		cfg.EigenContractAddr, parsed, cfg.L1Client, cfg.L1Client,
		cfg.L1Client,
	)
	// for use eigen fee model
	eigenFeeContract, err := bindings.NewBVMEigenDataLayrFee(
		cfg.EigenFeeContractAddress, cfg.L1Client,
	)
	if err != nil {
		log.Error("MtBatcher binding eigen fee contract fail", "err", err)
		return nil, err
	}

	feeParsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrFeeABI,
	))
	if err != nil {
		log.Error("MtBatcher parse eigen fee contract abi fail", "err", err)
		return nil, err
	}
	eigenFeeABI, err := bindings.BVMEigenDataLayrFeeMetaData.GetAbi()
	if err != nil {
		log.Error("MtBatcher get eigen fee contract abi fail", "err", err)
		return nil, err
	}
	rawEigenFeeContract := bind.NewBoundContract(
		cfg.EigenFeeContractAddress, feeParsed, cfg.L1Client, cfg.L1Client,
		cfg.L1Client,
	)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, cfg.L1Client)

	graphClient := graphView.NewGraphClient(cfg.GraphProvider, logger)

	walletAddr := crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)
	return &Driver{
		Cfg:                 cfg,
		Ctx:                 ctx,
		EigenDaContract:     eigenContract,
		RawEigenContract:    rawEigenContract,
		WalletAddr:          walletAddr,
		EigenABI:            eigenABI,
		EigenFeeContract:    eigenFeeContract,
		EigenFeeABI:         eigenFeeABI,
		RawEigenFeeContract: rawEigenFeeContract,
		GraphClient:         graphClient,
		logger:              logger,
		txMgr:               txMgr,
	}, nil
}

func (d *Driver) UpdateGasPrice(ctx context.Context, tx *types.Transaction, feeModelEnable bool) (*types.Transaction, error) {
	var finalTx *types.Transaction
	var err error
	opts := &bind.TransactOpts{
		From: d.WalletAddr,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return d.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   new(big.Int).SetUint64(tx.Nonce()),
		NoSend:  true,
	}
	if feeModelEnable {
		log.Info("update eigen da use fee", "FeeModelEnable", d.Cfg.FeeModelEnable)
		finalTx, err = d.RawEigenFeeContract.RawTransact(opts, tx.Data())
	} else {
		log.Info("rollup date", "FeeModelEnable", d.Cfg.FeeModelEnable)
		finalTx, err = d.RawEigenContract.RawTransact(opts, tx.Data())
	}
	switch {
	case err == nil:
		return finalTx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		if feeModelEnable {
			log.Info("update eigen da use fee", "FeeModelEnable", d.Cfg.FeeModelEnable)
			return d.RawEigenFeeContract.RawTransact(opts, tx.Data())
		} else {
			log.Info("rollup date", "FeeModelEnable", d.Cfg.FeeModelEnable)
			return d.RawEigenContract.RawTransact(opts, tx.Data())
		}
	default:
		return nil, err
	}
}

func (d *Driver) GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, error) {
	blockOffset := new(big.Int).SetUint64(d.Cfg.BlockOffset)
	var end *big.Int
	log.Info("MtBatcher GetBatchBlockRange", "blockOffset", blockOffset)
	start, err := d.EigenDaContract.GetL2StoredBlockNumber(&bind.CallOpts{
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

func (d *Driver) CalcUserFeeByRules(rollupDateSize *big.Int) (*big.Int, error) {
	seconds := new(big.Int).Div(big.NewInt(int64(d.Cfg.PollInterval)), big.NewInt(1000000000))
	rollSizeSec := new(big.Int).Div(rollupDateSize, seconds)
	feeSs, ok := new(big.Int).SetString(d.Cfg.FeeSizeSec, 10)
	if !ok {
		log.Error("FeeSizeSec from string to big.int fail")
		return big.NewInt(0), nil
	}
	if rollSizeSec.Cmp(feeSs) < 0 {
		return big.NewInt(0), nil
	}
	return big.NewInt(10000), nil
}

func (d *Driver) TxAggregator(ctx context.Context, start, end *big.Int) (transactionData []byte, startL2BlockNumber *big.Int, endL2BlockNumber *big.Int) {
	var batchTxList []BatchTx
	var transactionByte []byte
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		block, err := d.Cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, big.NewInt(0), big.NewInt(0)
		}
		txs := block.Transactions()
		if len(txs) != 1 {
			panic(fmt.Sprintf("MtBatcher attempting to create batch element from block %d, "+
				"found %d txs instead of 1", block.Number(), len(txs)))
		}
		log.Info("MtBatcher origin transactions", "TxHash", txs[0].Hash().String(), "l2BlockNumber", block.Number(), "QueueOrigin", txs[0].QueueOrigin())
		//isSequencerTx := txs[0].QueueOrigin() == l2types.QueueOriginSequencer
		//if !isSequencerTx || txs[0] == nil {
		//	continue
		//}
		var txBuf bytes.Buffer
		if err := txs[0].EncodeRLP(&txBuf); err != nil {
			panic(fmt.Sprintf("MtBatcher Unable to encode tx: %v", err))
		}
		batchTx := BatchTx{
			BlockNumber: i,
			RawTx:       txBuf.Bytes(),
		}
		batchTxList = append(batchTxList, batchTx)
		txnBufBytes, err := rlp.EncodeToBytes(batchTxList)
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
	txnBufBytes, err := rlp.EncodeToBytes(batchTxList)
	if err != nil {
		panic(fmt.Sprintf("MtBatcher Unable to encode txn: %v", err))
	}
	// order=3000 can send about 31600 byte data
	if len(txnBufBytes) > 31*d.Cfg.EigenLayerNode {
		transactionByte = txnBufBytes
	} else {
		paddingBytes := make([]byte, (31*d.Cfg.EigenLayerNode)-len(txnBufBytes))
		transactionByte = append(txnBufBytes, paddingBytes...)
	}
	return transactionByte, start, end
}

func (d *Driver) StoreData(ctx context.Context, uploadHeader []byte, duration uint8, blockNumber uint32, startL2BlockNumber *big.Int, endL2BlockNumber *big.Int, totalOperatorsIndex uint32) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current balance", "err", err)
		return nil, err
	}
	log.Info("MtBatcher WalletAddr Balance", "balance", balance)
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	opts := &bind.TransactOpts{
		From: d.WalletAddr,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return d.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   nonce,
		NoSend:  true,
	}
	tx, err := d.EigenDaContract.StoreData(opts, uploadHeader, duration, blockNumber, startL2BlockNumber, endL2BlockNumber, totalOperatorsIndex)
	switch {
	case err == nil:
		return tx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBather eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return d.EigenDaContract.StoreData(opts, uploadHeader, duration, blockNumber, startL2BlockNumber, endL2BlockNumber, totalOperatorsIndex)

	default:
		return nil, err
	}
}

func (d *Driver) ConfirmData(ctx context.Context, callData []byte, searchData rc.IDataLayrServiceManagerDataStoreSearchData, startL2BlockNumber, endL2BlockNumber *big.Int) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current balance", "err", err)
		return nil, err
	}
	log.Info("MtBatcher wallet address balance", "balance", balance)
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get current nonce", "err", err)
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	opts := &bind.TransactOpts{
		From: d.WalletAddr,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return d.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   nonce,
		NoSend:  true,
	}
	tx, err := d.EigenDaContract.ConfirmData(opts, callData, searchData, startL2BlockNumber, endL2BlockNumber)
	switch {
	case err == nil:
		return tx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBather eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return d.EigenDaContract.ConfirmData(opts, callData, searchData, startL2BlockNumber, endL2BlockNumber)

	default:
		return nil, err
	}
}

func (d *Driver) UpdateFee(ctx context.Context, l2Block, daFee *big.Int) (*types.Transaction, error) {
	balance, err := d.Cfg.L1Client.BalanceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get fee wallet address current balance", "err", err)
		return nil, err
	}
	log.Info("MtBatcher fee wallet address balance", "balance", balance)
	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, d.WalletAddr, nil,
	)
	if err != nil {
		log.Error("MtBatcher unable to get fee wallet nonce", "err", err)
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	opts := &bind.TransactOpts{
		From: d.WalletAddr,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return d.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   nonce,
		NoSend:  true,
	}
	tx, err := d.EigenFeeContract.SetUserRollupFee(opts, l2Block, daFee)
	switch {
	case err == nil:
		return tx, nil
	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBather eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return d.EigenFeeContract.SetUserRollupFee(opts, l2Block, daFee)
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

func (d *Driver) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return d.Cfg.L1Client.SendTransaction(ctx, tx)
}

func (d *Driver) DisperseStoreData(data []byte, startl2BlockNumber *big.Int, endl2BlockNumber *big.Int) (common2.StoreParams, *types.Receipt, error) {
	params, err := d.callEncode(data)
	if err != nil {
		return params, nil, err
	}
	uploadHeader, err := common2.CreateUploadHeader(params)
	if err != nil {
		return params, nil, err
	}
	tx, err := d.StoreData(
		d.Ctx, uploadHeader, uint8(params.Duration), params.BlockNumber, startl2BlockNumber, endl2BlockNumber, params.TotalOperatorsIndex,
	)
	if err != nil {
		log.Error("MtBatcher StoreData tx", "err", err)
		return params, nil, err
	} else if tx == nil {
		return params, nil, errors.New("tx is nil")
	}
	log.Info("MtBatcher store data success", "txHash", tx.Hash().String())
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return d.UpdateGasPrice(ctx, tx, false)
	}
	log.Info("MtBatcher updateGasPrice", "gasPrice", updateGasPrice)
	receipt, err := d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	if err != nil {
		log.Error("MtBatcher unable to StoreData", "err", err)
		return params, nil, err
	}
	return params, receipt, nil
}

func (d *Driver) ConfirmStoredData(txHash []byte, params common2.StoreParams, startl2BlockNumber, endl2BlockNumber *big.Int) (*types.Receipt, error) {
	event, ok := graphView.PollingInitDataStore(
		d.GraphClient,
		txHash[:],
		d.logger,
		12,
	)
	if !ok {
		log.Error("MtBatcher could not get initDataStore")
		return nil, errors.New("MtBatcher could not get initDataStore")
	}
	log.Info("PollingInitDataStore", "MsgHash", event.MsgHash, "StoreNumber", event.StoreNumber)
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
			HeaderHash:          event.DataCommitment,
			DurationDataStoreId: event.DurationDataStoreId,
			GlobalDataStoreId:   event.StoreNumber,
			BlockNumber:         event.StakesFromBlockNumber,
			Fee:                 event.Fee,
			Confirmer:           common.HexToAddress(event.Confirmer),
			SignatoryRecordHash: [32]byte{},
		},
	}
	obj, _ := json.Marshal(event)
	log.Info("MtBatcher Event", "obj", string(obj))
	log.Info("MtBatcher Calldata", "calldata", hexutil.Encode(callData))
	obj, _ = json.Marshal(params)
	log.Info("MtBatcher Params", "obj", string(obj))
	obj, _ = json.Marshal(meta)
	log.Info("MtBatcher Meta", "obj", string(obj))
	obj, _ = json.Marshal(searchData)
	log.Info("MtBatcher SearchData ", "obj", string(obj))
	log.Info("MtBatcher HeaderHash: ", "DataCommitment", hex.EncodeToString(event.DataCommitment[:]))
	log.Info("MtBatcher MsgHash", "event", hex.EncodeToString(event.MsgHash[:]))
	tx, err := d.ConfirmData(d.Ctx, callData, searchData, startl2BlockNumber, endl2BlockNumber)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		log.Info("MtBatcher ConfirmData update gas price")
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
	conn, err := grpc.Dial(d.Cfg.DisperserSocket, grpc.WithInsecure())
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
	log.Info("MtBatcher get store", "reply", reply)
	if err != nil {
		log.Error("MtBatcher get store err", err)
		return common2.StoreParams{}, err
	}
	log.Info("MtBatcher get store end")
	g := reply.GetStore()
	feeBigInt := new(big.Int).SetBytes(g.Fee)
	params := common2.StoreParams{
		BlockNumber:         g.BlockNumber,
		TotalOperatorsIndex: g.TotalOperatorsIndex,
		OrigDataSize:        g.OrigDataSize,
		NumTotal:            g.NumTotal,
		Quorum:              g.Quorum,
		NumSys:              g.NumSys,
		NumPar:              g.NumPar,
		Duration:            g.Duration,
		KzgCommit:           g.KzgCommit,
		LowDegreeProof:      g.LowDegreeProof,
		Degree:              g.Degree,
		TotalSize:           g.TotalSize,
		Order:               g.Order,
		Fee:                 feeBigInt,
		HeaderHash:          g.HeaderHash,
		Disperser:           g.Disperser,
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
		AggSig:           sigs.AggSig,
		AggPubKey:        sigs.AggPubKey,
		NonSignerPubKeys: sigs.NonSignerPubKeys,
	}
	meta := common2.DisperseMeta{
		Sigs:            aggSig,
		ApkIndex:        reply.GetApkIndex(),
		TotalStakeIndex: reply.GetTotalStakeIndex(),
	}
	return meta, nil
}

func (d *Driver) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), errMaxPriorityFeePerGasNotFound.Error(),
	)
}

func (d *Driver) Start() error {
	d.wg.Add(1)
	go d.eventLoop()
	return nil
}

func (d *Driver) Stop() {
	// d.cancel()
	d.wg.Wait()
}

func (d *Driver) eventLoop() {
	defer d.wg.Done()
	ticker := time.NewTicker(d.Cfg.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("MtBatcher eigen da sequencer fetching current block range")
			start, end, err := d.GetBatchBlockRange(d.Ctx)
			if err != nil {
				log.Error("MtBatcher Sequencer unable to get block range", "err", err)
				continue
			}
			log.Info("MtBatcher get batch block range", "start", start, "end", end)
			if start.Cmp(end) == 0 {
				log.Info("MtBatcher Sequencer no updates", "start", start, "end", end)
				continue
			}
			//if new(big.Int).Sub(end, start).Cmp(big.NewInt(int64(d.Cfg.BlockOffset))) < 0 {
			//	log.Info("MtBatcher end sub start must bigger than block offset", "start", start, "end", end, "BlockOffset", d.Cfg.BlockOffset)
			//	continue
			//}
			aggregateTxData, startL2BlockNumber, endL2BlockNumber := d.TxAggregator(
				d.Ctx, start, end,
			)
			if err != nil {
				log.Error("MtBatcher eigenDa sequencer unable to craft batch tx", "err", err)
				continue
			}
			params, receipt, err := d.DisperseStoreData(aggregateTxData, startL2BlockNumber, endL2BlockNumber)
			if err != nil {
				log.Error("MtBatcher disperse store data fail", "err", err)
				continue
			}
			log.Info("MtBatcher disperse store data success", "txHash", receipt.TxHash.String())
			csdReceipt, err := d.ConfirmStoredData(receipt.TxHash.Bytes(), params, startL2BlockNumber, endL2BlockNumber)
			if err != nil {
				log.Error("MtBatcher confirm store data fail", "err", err)
				continue
			}
			log.Info("MtBatcher confirm store data success", "txHash", csdReceipt.TxHash.String())
			if d.Cfg.FeeModelEnable {
				daFee, _ := d.CalcUserFeeByRules(big.NewInt(int64(len(aggregateTxData))))
				chainFee, err := d.EigenFeeContract.GetUserRollupFee(&bind.CallOpts{})
				if err != nil {
					log.Error("get chain fee fail", "err", err)
					continue
				}
				log.Info("chainFee and daFee", "chainFee", chainFee, "daFee", daFee)
				if chainFee.Cmp(daFee) != 0 {
					txfRpt, err := d.UpdateUserDaFee(endL2BlockNumber, daFee)
					if err != nil {
						log.Error("update user da fee fail", "err", err)
						continue
					}
					log.Info("update user fee success", "Hash", txfRpt.TxHash.String())
				}
			}
		case err := <-d.Ctx.Done():
			log.Error("MtBatcher eigenDa sequencer service shutting down", "err", err)
			return
		}
	}
}
