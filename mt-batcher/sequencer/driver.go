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
	"unsafe"
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
	PrivKey                   *ecdsa.PrivateKey
	BlockOffset               uint64
	EigenLayerNode            int
	ChainID                   *big.Int
	DataStoreDuration         uint64
	DataStoreTimeout          uint64
	DisperserSocket           string
	PollInterval              time.Duration
	GraphProvider             string
	EigenLogConfig            logging.Config
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	SignerFn                  SignerFn
}

type Driver struct {
	Ctx              context.Context
	Cfg              *DriverConfig
	EigenDaContract  *bindings.BVMEigenDataLayrChain
	RawEigenContract *bind.BoundContract
	WalletAddr       common.Address
	EigenABI         *abi.ABI
	GraphClient      *graphView.GraphClient
	logger           *logging.Logger
	txMgr            txmgr.TxManager
	cancel           func()
	wg               sync.WaitGroup
}

var bigOne = new(big.Int).SetUint64(1)

func NewDriver(ctx context.Context, cfg *DriverConfig) (*Driver, error) {
	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		cfg.EigenContractAddr, cfg.L1Client,
	)
	if err != nil {
		log.Error("binding eigenda contract fail", "err", err)
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
		log.Error("parse eigenda contract abi fail", "err", err)
		return nil, err
	}
	eignenABI, err := bindings.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		log.Error("get eigenda contract abi fail", "err", err)
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		cfg.EigenContractAddr, parsed, cfg.L1Client, cfg.L1Client,
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
		Cfg:              cfg,
		Ctx:              ctx,
		EigenDaContract:  eigenContract,
		RawEigenContract: rawEigenContract,
		WalletAddr:       walletAddr,
		EigenABI:         eignenABI,
		GraphClient:      graphClient,
		logger:           logger,
		txMgr:            txMgr,
	}, nil
}

func (d *Driver) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	opts := &bind.TransactOpts{
		From: d.WalletAddr,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return d.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   new(big.Int).SetUint64(tx.Nonce()),
		NoSend:  true,
	}
	finalTx, err := d.RawEigenContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	case d.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtBatcher eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return d.RawEigenContract.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

func (d *Driver) GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, error) {
	blockOffset := new(big.Int).SetUint64(d.Cfg.BlockOffset)
	var end *big.Int
	log.Info("MtBatcher GetBatchBlockRange", "blockOffset", blockOffset)
	start, err := d.EigenDaContract.GetL2SubmitBlockNumber(&bind.CallOpts{
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
	var batchTxList []BatchTx
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		block, err := d.Cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, big.NewInt(0), big.NewInt(0)
		}
		txs := block.Transactions()
		if len(txs) != 1 {
			panic(fmt.Sprintf("attempting to create batch element from block %d, "+
				"found %d txs instead of 1", block.Number(), len(txs)))
		}
		log.Info("Origin Transactions", "txs[0]", txs[0], "Transaction l2BlockNumber", block.Number(), "txs[0].QueueOrigin()", txs[0].QueueOrigin())
		//isSequencerTx := txs[0].QueueOrigin() == l2types.QueueOriginSequencer
		//if !isSequencerTx || txs[0] == nil {
		//	continue
		//}
		var txBuf bytes.Buffer
		if err := txs[0].EncodeRLP(&txBuf); err != nil {
			panic(fmt.Sprintf("Unable to encode tx: %v", err))
		}
		log.Info("Rlp Transactions", "txBuf", txBuf.Bytes(), "txs[0].QueueOrigin()", txs[0].QueueOrigin())
		batchTx := BatchTx{
			BlockNumber: i,
			rawTx:       txBuf.Bytes(),
		}
		batchTxList = append(batchTxList, batchTx)
	}
	batchTxVector := &BatchTxVector{batchTxList}
	Len := unsafe.Sizeof(*batchTxVector)
	batchTxVectorBytes := &SliceBatchTx{
		addr: uintptr(unsafe.Pointer(batchTxVector)),
		cap:  int(Len),
		len:  int(Len),
	}
	var transactionByte []byte
	txBufBytes := *(*[]byte)(unsafe.Pointer(batchTxVectorBytes))
	if len(txBufBytes) > 31*d.Cfg.EigenLayerNode {
		transactionByte = txBufBytes
	} else {
		paddingBytes := make([]byte, (31*d.Cfg.EigenLayerNode)-len(txBufBytes))
		transactionByte = append(txBufBytes, paddingBytes...)
	}
	log.Info("transaction data len", "dataLen", len(transactionByte))
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
	log.Info("WalletAddr Balance", "balance", balance)
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
	log.Info("WalletAddr Balance", "balance", balance)
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

func (d *Driver) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return d.Cfg.L1Client.SendTransaction(ctx, tx)
}

func (d *Driver) DisperseStoreAndConfirmData(data []byte, startl2BlockNumber *big.Int, endl2BlockNumber *big.Int) error {
	params, err := d.callEncode(data)
	if err != nil {
		return err
	}
	uploadHeader, err := common2.CreateUploadHeader(params)
	if err != nil {
		return err
	}
	tx, err := d.StoreData(
		d.Ctx, uploadHeader, uint8(params.Duration), params.BlockNumber, startl2BlockNumber, endl2BlockNumber, params.TotalOperatorsIndex,
	)
	if err != nil {
		log.Error("MtBatcher StoreData tx", "err", err)
		return err
	} else if tx == nil {
		return errors.New("tx is nil")
	}
	log.Info("d.StoreData", "txHash", tx.Hash().String())
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return d.UpdateGasPrice(ctx, tx)
	}
	log.Info("updateGasPrice", "gasPrice", updateGasPrice)
	receipt, err := d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	log.Info(" d.txMgr.Send", "receipt", receipt.Status, "err", err)

	if err != nil {
		log.Error("MtBatcher unable to StoreData", "err", err)
		return err
	}
	log.Info("MtBatcher StoreData successfully", "tx_hash", receipt.TxHash)
	event, ok := graphView.PollingInitDataStore(
		d.GraphClient,
		tx.Hash().Bytes()[:],
		d.logger,
		12,
	)
	if !ok {
		log.Error("could not get initDataStore")
		return errors.New("could not get initDataStore")
	}
	log.Info("PollingInitDataStore", "MsgHash", event.MsgHash, "StoreNumber", event.StoreNumber)
	meta, err := d.callDisperse(
		params.HeaderHash,
		event.MsgHash[:],
	)
	if err != nil {
		log.Error("callDisperse fail", "err", err)
		return err
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
	log.Info("Event", "obj", string(obj))
	log.Info("Calldata", "calldata", hexutil.Encode(callData))
	obj, _ = json.Marshal(params)
	log.Info("Params", "obj", string(obj))
	obj, _ = json.Marshal(meta)
	log.Info("Meta", "obj", string(obj))
	obj, _ = json.Marshal(searchData)
	log.Info("SearchData ", "obj", string(obj))
	log.Info("HeaderHash: ", "DataCommitment", hex.EncodeToString(event.DataCommitment[:]))
	log.Info("MsgHash", "event", hex.EncodeToString(event.MsgHash[:]))
	tx, err = d.ConfirmData(d.Ctx, callData, searchData, startl2BlockNumber, endl2BlockNumber)
	if err != nil {
		return err
	}
	updateGasPrice = func(ctx context.Context) (*types.Transaction, error) {
		log.Info("MtBatcher ConfirmData update gas price")
		return d.UpdateGasPrice(ctx, tx)
	}
	receipt, err = d.txMgr.Send(
		d.Ctx, updateGasPrice, d.SendTransaction,
	)
	if err != nil {
		log.Error("MtBatcher unable to ConfirmData tx", "err", err)
		return err
	}
	log.Info("MtBatcher ConfirmData successfully", "tx_hash", receipt.TxHash)
	return nil
}

func (d *Driver) callEncode(data []byte) (common2.StoreParams, error) {
	conn, err := grpc.Dial(d.Cfg.DisperserSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Err. Disperser Cannot connect to", "DisperserSocket", d.Cfg.DisperserSocket)
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
	log.Info("get store", "reply", reply)
	if err != nil {
		log.Error("get store err", err)
		return common2.StoreParams{}, err
	}
	log.Info("get store end")
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
		log.Error("d.cfg.DisperserSocket", "err", err)
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
			log.Info("EigenDa Sequencer fetching current block range")
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
			if new(big.Int).Sub(end, start).Cmp(big.NewInt(int64(d.Cfg.BlockOffset))) < 0 {
				log.Info("end sub start must bigger than block offset", "start", start, "end", end, "BlockOffset", d.Cfg.BlockOffset)
				continue
			}
			aggregateTxData, startL2BlockNumber, endL2BlockNumber := d.TxAggregator(
				d.Ctx, start, end,
			)
			if err != nil {
				log.Error("EigenDa Sequencer unable to craft batch tx", "err", err)
				continue
			}
			err = d.DisperseStoreAndConfirmData(aggregateTxData, startL2BlockNumber, endL2BlockNumber)
			if err != nil {
				log.Error("DisperseStoreAndConfirmData fail", "err", err)
				continue
			}
		case err := <-d.Ctx.Done():
			log.Error("EigenDa Sequencer service shutting down", "err", err)
			return
		}
	}
}
