package restorer

import (
	"bytes"
	"context"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceRetrieverServer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	gecho "github.com/labstack/echo/v4"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/pkg/errors"
	"github.com/shurcooL/graphql"
	"google.golang.org/grpc"
	"math/big"
	"net/http"
	"strings"
)

type RollupStoreRequest struct {
	BatchIndex int64 `json:"batch_index"`
}

type TransactionRequest struct {
	StoreNumber uint32 `json:"store_number"`
}

type DataStoreRequest struct {
	FromStoreNumber   string `json:"from_store_number"`
	EigenContractAddr string `json:"eigen_contract_addr"`
}

func (s *DaService) GetLatestTransactionBatchIndex(c gecho.Context) error {
	batchIndex, err := s.Cfg.EigenContract.RollupBatchIndex(&bind.CallOpts{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("fail to get batch index"))
	}
	log.Info("rollup batch index", "batchIndex", batchIndex.Uint64())
	return c.JSON(http.StatusOK, batchIndex.Uint64())
}

func (s *DaService) GetRollupStoreByRollupBatchIndex(c gecho.Context) error {
	var rsReq RollupStoreRequest
	if err := c.Bind(&rsReq); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid request params"))
	}
	rollupStore, err := s.Cfg.EigenContract.GetRollupStoreByRollupBatchIndex(&bind.CallOpts{}, big.NewInt(rsReq.BatchIndex))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("get rollup store fail"))
	}
	rsRep := &common.RollupStoreResponse{
		DataStoreId: rollupStore.DataStoreId,
		ConfirmAt:   rollupStore.ConfirmAt,
		Status:      rollupStore.Status,
	}
	return c.JSON(http.StatusOK, rsRep)
}

func (s *DaService) GetBatchTransactionByDataStoreId(c gecho.Context) error {
	var txReq TransactionRequest
	if err := c.Bind(&txReq); err != nil {
		log.Error("invalid request params", "err", err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid request params"))
	}
	log.Info("GetBatchTransactionByDataStoreId Request para", "StoreNumber", txReq.StoreNumber)
	conn, err := grpc.Dial(s.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("disperser Cannot connect to", "err", err)
		return c.JSON(http.StatusBadRequest, errors.New("disperser Cannot connect to"))
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)

	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: txReq.StoreNumber,
	}
	reply, err := client.RetrieveFramesAndData(s.Ctx, request, opt)
	if err != nil {
		log.Error("retrieve frames and data error", "err", err)
		return c.JSON(http.StatusBadRequest, errors.New("recovery data fail"))
	}
	batchTxn := new([]common.BatchTx)
	batchRlpStream := rlp.NewStream(bytes.NewBuffer(reply.GetData()), 0)
	err = batchRlpStream.Decode(batchTxn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("decode batch tx fail"))
	}
	newBatchTxn := *batchTxn
	var txList []types.Transaction
	for i := 0; i < len(newBatchTxn); i++ {
		var l2Tx types.Transaction
		rlpStream := l2rlp.NewStream(bytes.NewBuffer(newBatchTxn[i].RawTx), 0)
		if err := l2Tx.DecodeRLP(rlpStream); err != nil {
			log.Error("Decode RLP fail")
		}
		log.Info("Transaction", "TxHash", l2Tx.Hash().Hex(), "Index", l2Tx.GetMeta().Index)
		txList = append(txList, l2Tx)
	}
	return c.JSON(http.StatusOK, txList)
}

func (s *DaService) GetDataStore(c gecho.Context) error {
	var dsReq DataStoreRequest
	if err := c.Bind(&dsReq); err != nil {
		log.Error("invalid request params", "err", err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid request params"))
	}
	var query struct {
		DataStores []graphView.DataStoreGql `graphql:"dataStores(first:1,where:{storeNumber_gt: $lastStoreNumber,confirmer: $confirmer,confirmed:true})"`
	}
	variables := map[string]interface{}{
		"lastStoreNumber": graphql.String(dsReq.FromStoreNumber),
		"confirmer":       graphql.String(strings.ToLower(dsReq.EigenContractAddr)),
	}
	err := s.GraphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		log.Error("GetExpiringDataStores error")
		return c.JSON(http.StatusBadRequest, errors.New("iGetExpiringDataStores error"))
	}
	if len(query.DataStores) == 0 {
		return c.JSON(http.StatusBadRequest, errors.New("no new stores"))
	}
	store, err := query.DataStores[0].Convert()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("conversion error"))
	}
	return c.JSON(http.StatusOK, store)
}

func (s *DaService) GetTransactionList(c gecho.Context) error {
	var txReq TransactionRequest
	if err := c.Bind(&txReq); err != nil {
		log.Error("invalid request params", "err", err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid request params"))
	}
	conn, err := grpc.Dial(s.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("disperser Cannot connect to"))
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.Timeout)
	defer cancel()
	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: txReq.StoreNumber,
	}
	reply, err := client.RetrieveFramesAndData(ctx, request, opt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("RetrieveFramesAndData error"))
	}
	data := reply.GetData()
	batchTxn := new([]common.BatchTx)
	batchRlpStream := rlp.NewStream(bytes.NewBuffer(data), uint64(len(data)))
	err = batchRlpStream.Decode(batchTxn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("decode data fail"))
	}
	var TxHash []string
	newBatchTxn := *batchTxn
	for i := 0; i < len(newBatchTxn); i++ {
		l2Tx := new(types.Transaction)
		rlpStream := l2rlp.NewStream(bytes.NewBuffer(newBatchTxn[i].RawTx), 0)
		if err := l2Tx.DecodeRLP(rlpStream); err != nil {
			log.Error("Decode RLP fail")
			continue
		}
		log.Info("transaction", "hash", l2Tx.Hash().Hex())
		TxHash = append(TxHash, l2Tx.Hash().String())
	}
	return c.JSON(http.StatusOK, TxHash)
}
