package restorer

import (
	"bytes"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceRetrieverServer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/labstack/echo/v4"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"google.golang.org/grpc"
	"math/big"
	"net/http"
)

const (
	SelfServiceOK     = 2000
	SelfServiceError  = 4000
	SelfInvalidParams = 4001
)

type RollupStoreRequest struct {
	BatchIndex int64 `json:"batch_index"`
}

type RollupStoreResponse struct {
	DataStoreId uint32 `json:"data_store_id"`
	ConfirmAt   uint32 `json:"confirm_at"`
	Status      uint8  `json:"status"`
}

type TransactionRequest struct {
	StoreNumber uint32 `json:"store_number"`
}

func (s *DaService) GetLatestTransactionBatchIndex(c echo.Context) error {
	batchIndex, err := s.Cfg.EigenContract.RollupBatchIndex(&bind.CallOpts{})
	if err != nil {
		retValue := common.BaseResource(false, SelfServiceError, nil, "get roll batch index fail")
		return c.JSON(http.StatusOK, retValue)
	}
	retValue := common.BaseResource(true, SelfServiceOK, batchIndex, "get batch index success")
	return c.JSON(http.StatusOK, retValue)
}

func (s *DaService) GetRollupStoreByRollupBatchIndex(c echo.Context) error {
	var rsReq RollupStoreRequest
	if err := c.Bind(&rsReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	rollupStore, err := s.Cfg.EigenContract.GetRollupStoreByRollupBatchIndex(&bind.CallOpts{}, big.NewInt(rsReq.BatchIndex))
	if err != nil {
		retValue := common.BaseResource(false, SelfServiceError, nil, "get roll batch index fail")
		return c.JSON(http.StatusOK, retValue)
	}
	rsRep := &RollupStoreResponse{
		DataStoreId: rollupStore.DataStoreId,
		ConfirmAt:   rollupStore.ConfirmAt,
		Status:      rollupStore.Status,
	}
	retValue := common.BaseResource(true, SelfServiceOK, rsRep, "get roll store success")
	return c.JSON(http.StatusOK, retValue)
}

func (s *DaService) GetBatchTransactionByDataStoreId(c echo.Context) error {
	var txReq TransactionRequest
	if err := c.Bind(&txReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "Params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	conn, err := grpc.Dial(s.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "Disperser Cannot connect to")
		return c.JSON(http.StatusOK, retValue)
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)

	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: txReq.StoreNumber,
	}
	reply, err := client.RetrieveFramesAndData(s.Ctx, request, opt)
	if err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "Recovery data fail")
		return c.JSON(http.StatusOK, retValue)
	}
	batchTxn := new([]common.BatchTx)
	batchRlpStream := rlp.NewStream(bytes.NewBuffer(reply.GetData()), 0)
	err = batchRlpStream.Decode(batchTxn)
	if err != nil {
		log.Error("decode batch tx fail")
		retValue := common.BaseResource(false, SelfServiceError, nil, "decode batch tx fail")
		return c.JSON(http.StatusOK, retValue)
	}
	newBatchTxn := *batchTxn
	var txList []*types.Transaction
	for i := 0; i < len(newBatchTxn); i++ {
		l2Tx := new(types.Transaction)
		rlpStream := l2rlp.NewStream(bytes.NewBuffer(newBatchTxn[i].RawTx), 0)
		if err := l2Tx.DecodeRLP(rlpStream); err != nil {
			log.Error("Decode RLP fail")
		}
		log.Info("tx hash:" + l2Tx.Hash().Hex())
		txList = append(txList, l2Tx)
	}
	retValue := common.BaseResource(true, SelfServiceOK, txList, "get transaction from eigen da success")
	return c.JSON(http.StatusOK, retValue)
}
