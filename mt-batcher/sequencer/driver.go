package sequencer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceDL"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	"github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	rc "github.com/mantlenetworkio/mantle/mt-batcher/bindings"
	common2 "github.com/mantlenetworkio/mantle/mt-batcher/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/big"
	"strings"
	"sync"
	"time"
)

type DriverConfig struct {
	L1Client          *l1l2client.L1ChainClient
	L2Client          *l2ethclient.Client
	EigenAddr         common.Address
	PrivKey           *ecdsa.PrivateKey
	BlockOffset       uint64
	ChainID           *big.Int
	DataStoreDuration uint64
	DataStoreTimeout  uint64
	DisperserSocket   string
	PollInterval      time.Duration
	GraphProvider     string
}

type Driver struct {
	Ctx              context.Context
	Cfg              *DriverConfig
	EigenDaContract  *bindings.BVMEigenDataLayrChain
	RawEigenContract *bind.BoundContract
	WalletAddr       common.Address
	EigenABI         *abi.ABI
	L1ChainClient    *l1l2client.L1ChainClient
	GraphClient      *graphView.GraphClient
	cancel           func()
	wg               sync.WaitGroup
}

var bigOne = new(big.Int).SetUint64(1)

func NewDriver(ctx context.Context, cfg *DriverConfig) (*Driver, error) {
	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		cfg.EigenAddr, cfg.L1Client.Client,
	)
	if err != nil {
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrChainABI,
	))
	if err != nil {
		return nil, err
	}

	eignenABI, err := bindings.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		cfg.EigenAddr, parsed, cfg.L1Client.Client, cfg.L1Client.Client,
		cfg.L1Client.Client,
	)

	graphClient := graphView.NewGraphClient(cfg.GraphProvider, nil)

	walletAddr := crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)
	return &Driver{
		Cfg:              cfg,
		Ctx:              ctx,
		EigenDaContract:  eigenContract,
		RawEigenContract: rawEigenContract,
		WalletAddr:       walletAddr,
		EigenABI:         eignenABI,
		L1ChainClient:    cfg.L1Client,
		GraphClient:      graphClient,
	}, nil
}

func (d *Driver) GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, *big.Int, error) {
	blockOffset := new(big.Int).SetUint64(d.Cfg.BlockOffset)
	start, err := d.EigenDaContract.L2BlockNumber(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	start.Add(start, blockOffset)
	latestHeader, err := d.Cfg.L2Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	end := new(big.Int).Add(latestHeader.Number, bigOne)
	if start.Cmp(end) > 0 {
		return nil, nil, nil, fmt.Errorf("invalid range, "+
			"end(%v) < start(%v)", end, start)
	}
	return start, end, latestHeader.Number, nil
}

func (d *Driver) CraftBatchTx(
	ctx context.Context,
	start, end, blockNumber *big.Int,
) (*types.Transaction, error) {
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		block, err := d.Cfg.L2Client.BlockByNumber(ctx, i)
		if err != nil {
			return nil, err
		}
		txs := block.Transactions()
		if len(txs) != 1 {
			panic(fmt.Sprintf("attempting to create batch element from block %d, "+
				"found %d txs instead of 1", block.Number(), len(txs)))
		}
		var txBuf bytes.Buffer
		if err := txs[0].EncodeRLP(&txBuf); err != nil {
			panic(fmt.Sprintf("Unable to encode tx: %v", err))
		}
		err = d.Disperse(txBuf.Bytes(), blockNumber)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (d *Driver) Disperse(data []byte, l2BlockNumber *big.Int) error {
	params, err := d.callEncode(data)
	if err != nil {
		return err
	}
	uploadHeader, err := common2.CreateUploadHeader(params)
	if err != nil {
		return err
	}
	auth := d.L1ChainClient.PrepareAuthTransactor()
	tx, err := d.EigenDaContract.StoreData(auth, uploadHeader, uint8(params.Duration), params.BlockNumber, l2BlockNumber, params.TotalOperatorsIndex)
	if err != nil {
		return err
	}
	err = d.L1ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
	event, ok := graphView.PollingInitDataStore(
		d.GraphClient,
		tx.Hash().Bytes()[:],
		nil,
		12,
	)
	if !ok {
		return errors.New("could not get initDataStore")
	}
	meta, err := d.callDisperse(
		params.HeaderHash,
		event.MsgHash[:],
	)
	if err != nil {
		return err
	}
	calldata := common2.MakeCalldata(params, meta, event.StoreNumber, event.MsgHash)
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
	auth = d.L1ChainClient.PrepareAuthTransactor()
	tx, err = d.EigenDaContract.ConfirmData(auth, calldata, searchData)
	if err != nil {
		return err
	}
	fmt.Printf("ConfirmDataStore tx sent. TxHash: %v\n", tx.Hash().Hex())
	err = d.L1ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
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
	log.Info("get store")
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

func (s *Driver) Start() error {
	s.wg.Add(1)
	go s.eventLoop()
	return nil
}

func (s *Driver) Stop() {
	s.cancel()
	s.wg.Wait()
}

func (s *Driver) eventLoop() {
	defer s.wg.Done()
	ticker := time.NewTicker(s.Cfg.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("EigenDa Sequencer  fetching current block range")
			start, end, l2block, err := s.GetBatchBlockRange(s.Ctx)
			if err != nil {
				log.Error("EigenDa Sequencer unable to get block range", "err", err)
				continue
			}
			if start.Cmp(end) == 0 {
				log.Info("EigenDa Sequencer no updates", "start", start, "end", end)
				continue
			}
			log.Info("EigenDa Sequencer block range", "start", start, "end", end)
			tx, err := s.CraftBatchTx(
				s.Ctx, start, end, l2block,
			)
			if err != nil {
				log.Error("EigenDa Sequencer unable to craft batch tx",
					"err", err)
				continue
			} else if tx == nil {
				continue
			}
		case err := <-s.Ctx.Done():
			log.Error("EigenDa Sequencer service shutting down", "err", err)
			return
		}
	}
}
