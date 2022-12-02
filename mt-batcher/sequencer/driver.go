package sequencer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceDL"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	rc "github.com/mantlenetworkio/mantle/mt-batcher/bindings"

	common2 "github.com/mantlenetworkio/mantle/mt-batcher/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/dial"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"google.golang.org/grpc"
	"math/big"
	"time"
)

type DataStoreSettings struct {
	Duration uint64
	Timeout  uint64
}

type ServerSettings struct {
	Port int
}

type DisperserSettings struct {
	Socket string
}

type RollupSettings struct {
	Address common.Address
}

type SequencerSettings struct {
	ServerSettings    ServerSettings
	DataStoreSettings DataStoreSettings
	DisperserSettings DisperserSettings
	RollupSettings    RollupSettings
	ChainSettings     dial.ChainSettings
	GraphEndpoint     string
}

type EigenSequencer struct {
	ctx context.Context
	SequencerSettings
	ChainClient *dial.ChainClient
	L2MtlCilent *l2ethclient.Client
	GraphClient *graphView.GraphClient
}

func NewEigenSequencer(
	ctx context.Context,
	chainClient *dial.ChainClient,
	graphClient *graphView.GraphClient,
	settings SequencerSettings,
	l2cli *l2ethclient.Client,
) *EigenSequencer {
	return &EigenSequencer{
		ctx:               ctx,
		SequencerSettings: settings,
		ChainClient:       chainClient,
		L2MtlCilent:       l2cli,
		GraphClient:       graphClient,
	}
}

func (es *EigenSequencer) Start() error {
	err := es.FetchBlock(es.ctx)
	if err != nil {
		log.Error("")
	}
	return nil
}

func (es *EigenSequencer) FetchBlock(ctx context.Context) error {
	return nil
}

func (es *EigenSequencer) Disperse(data []byte) error {
	params, err := es.callEncode(data)
	if err != nil {
		return err
	}
	log.Info("after encode")

	uploadHeader, err := common2.CreateUploadHeader(params)
	if err != nil {
		return err
	}
	rollup, err := es.getRollupContractBinding()
	if err != nil {
		return err
	}
	auth := es.ChainClient.PrepareAuthTransactor()
	tx, err := rollup.StoreData(auth, uploadHeader, uint8(params.Duration), params.BlockNumber, params.TotalOperatorsIndex)
	if err != nil {
		return err
	}
	err = es.ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
	event, ok := graphView.PollingInitDataStore(
		es.GraphClient,
		tx.Hash().Bytes()[:],
		nil,
		12,
	)
	if !ok {
		return errors.New("could not get initDataStore")
	}
	meta, err := es.callDisperse(
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
	obj, _ := json.Marshal(event)
	log.Info("Event: " + string(obj))
	log.Info("Calldata: " + hexutil.Encode(calldata))
	obj, _ = json.Marshal(params)
	log.Info("Params: " + string(obj))
	obj, _ = json.Marshal(meta)
	log.Info("Meta: " + string(obj))
	obj, _ = json.Marshal(searchData)
	log.Info("SearchData: " + string(obj))
	log.Info("HeaderHash: " + hex.EncodeToString(event.DataCommitment[:]))
	log.Info("MsgHash: " + hex.EncodeToString(event.MsgHash[:]))
	auth = es.ChainClient.PrepareAuthTransactor()
	tx, err = rollup.ConfirmData(auth, calldata, searchData)
	if err != nil {
		return err
	}
	fmt.Printf("ConfirmDataStore tx sent. TxHash: %v\n", tx.Hash().Hex())
	err = es.ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
	return nil
}

func (es *EigenSequencer) callEncode(data []byte) (common2.StoreParams, error) {
	// ToDo divide file to chunks and send via stream if file is too large
	conn, err := grpc.Dial(es.DisperserSettings.Socket, grpc.WithInsecure())
	if err != nil {
		log.Error("Err. Disperser Cannot connect to", "socket", es.DisperserSettings.Socket)
		return common2.StoreParams{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(es.DataStoreSettings.Timeout))
	defer cancel()
	request := &pb.EncodeStoreRequest{
		Duration: es.DataStoreSettings.Duration,
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

func (es *EigenSequencer) callDisperse(headerHash []byte, messageHash []byte) (common2.DisperseMeta, error) {
	conn, err := grpc.Dial(es.DisperserSettings.Socket, grpc.WithInsecure())
	if err != nil {
		log.Error("es.DisperserSettings.Socket", "err", err)
		return common2.DisperseMeta{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(es.DataStoreSettings.Timeout))
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

func (es *EigenSequencer) getRollupContractBinding() (*rc.BVMEigenDataLayrChain, error) {
	rollup, err := rc.NewBVMEigenDataLayrChain(es.RollupSettings.Address, es.ChainClient.Client)
	if err != nil {
		return nil, err
	}
	return rollup, nil
}
