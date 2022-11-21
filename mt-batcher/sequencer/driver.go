package sequencer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceDL"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	rc "github.com/mantlenetworkio/mantle/mt-batcher/bindings/DataLayrRollup"
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

	GraphEndpoint string
}

type MtBatcher struct {
	ctx context.Context
	SequencerSettings
	ChainClient *dial.ChainClient
	GraphClient *graphView.GraphClient
}

func NewMtBatcher(
	ctx context.Context,
	chainClient *dial.ChainClient,
	graphClient *graphView.GraphClient,
	settings SequencerSettings,
) *MtBatcher {
	return &MtBatcher{
		ctx:               ctx,
		SequencerSettings: settings,
		ChainClient:       chainClient,
		GraphClient:       graphClient,
	}
}

func (mt *MtBatcher) Start() error {
	err := mt.Stake()
	if err != nil {
		return err
	}
	err = mt.FetchBlock(mt.ctx)
	if err != nil {
		log.Error("")
	}
	return nil
}

func (mt *MtBatcher) FetchBlock(ctx context.Context) error {
	return nil
}

func (s *MtBatcher) Stake() error {
	rollup, err := s.getRollupContractBinding()
	if err != nil {
		return err
	}
	status, err := rollup.SequencerStatus(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if status == 0 {
		auth := s.ChainClient.PrepareAuthTransactor()
		tx, err := rollup.Stake(auth)
		if err != nil {
			return err
		}
		log.Info("Stake Tx: " + tx.Hash().Hex())
		err = s.ChainClient.EnsureTransactionEvaled(tx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mt *MtBatcher) Disperse(data []byte) error {
	params, err := mt.callEncode(data)
	if err != nil {
		return err
	}
	log.Info("after encode")

	uploadHeader, err := createUploadHeader(params)
	if err != nil {
		return err
	}
	rollup, err := mt.getRollupContractBinding()
	if err != nil {
		return err
	}
	auth := mt.ChainClient.PrepareAuthTransactor()
	tx, err := rollup.StoreData(auth, uploadHeader, uint8(params.Duration), params.BlockNumber, params.TotalOperatorsIndex)
	if err != nil {
		return err
	}
	err = mt.ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
	event, ok := graphView.PollingInitDataStore(
		s.GraphClient,
		tx.Hash().Bytes()[:],
		s.Logger,
		12,
	)
	if !ok {
		return errors.New("could not get initDataStore")
	}
	meta, err := mt.callDisperse(
		params.HeaderHash,
		event.MsgHash[:],
	)
	if err != nil {
		return err
	}
	calldata := makeCalldata(params, meta, event.StoreNumber, event.MsgHash)
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
	auth = mt.ChainClient.PrepareAuthTransactor()
	tx, err = rollup.ConfirmData(auth, calldata, searchData)
	if err != nil {
		return err
	}
	fmt.Printf("ConfirmDataStore tx sent. TxHash: %v\n", tx.Hash().Hex())
	err = mt.ChainClient.EnsureTransactionEvaled(tx)
	if err != nil {
		return err
	}
	return nil
}

func (mt *MtBatcher) callEncode(data []byte) (StoreParams, error) {
	// ToDo divide file to chunks and send via stream if file is too large
	conn, err := grpc.Dial(s.DisperserSettings.Socket, grpc.WithInsecure())
	if err != nil {
		log.Error("Err. Disperser Cannot connect to", "socket", mt.DisperserSettings.Socket)
		return StoreParams{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.DataStoreSettings.Timeout))
	defer cancel()
	request := &pb.EncodeStoreRequest{
		Duration: mt.DataStoreSettings.Duration,
		Data:     data,
	}
	opt := grpc.MaxCallSendMsgSize(1024 * 1024 * 300)
	reply, err := c.EncodeStore(ctx, request, opt)
	log.Info("get store")
	if err != nil {
		log.Error("get store err", err)
		return StoreParams{}, err
	}
	log.Info("get store end")
	g := reply.GetStore()
	feeBigInt := new(big.Int).SetBytes(g.Fee)
	params := StoreParams{
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

func (mt *MtBatcher) callDisperse(headerHash []byte, messageHash []byte) (DisperseMeta, error) {
	conn, err := grpc.Dial(mt.DisperserSettings.Socket, grpc.WithInsecure())
	if err != nil {
		log.Error("mt.DisperserSettings.Socket", "err", err)
		return DisperseMeta{}, err
	}
	defer conn.Close()
	c := pb.NewDataDispersalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.DataStoreSettings.Timeout))
	defer cancel()
	request := &pb.DisperseStoreRequest{
		HeaderHash:  headerHash,
		MessageHash: messageHash,
	}
	reply, err := c.DisperseStore(ctx, request)
	if err != nil {
		return DisperseMeta{}, err
	}
	sigs := reply.GetSigs()
	aggSig := AggregateSignature{
		AggSig:           sigs.AggSig,
		AggPubKey:        sigs.AggPubKey,
		NonSignerPubKeys: sigs.NonSignerPubKeys,
	}
	meta := DisperseMeta{
		Sigs:            aggSig,
		ApkIndex:        reply.GetApkIndex(),
		TotalStakeIndex: reply.GetTotalStakeIndex(),
	}
	return meta, nil
}

func (mt *MtBatcher) getRollupContractBinding() (*rc.ContractDataLayrRollup, error) {
	rollup, err := rc.NewContractDataLayrRollup(mt.RollupSettings.Address, mt.ChainClient.Client)
	if err != nil {
		return nil, err
	}
	return rollup, nil
}
