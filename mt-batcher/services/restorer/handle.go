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
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/big"
	"net/http"
)

type RollupStoreRequest struct {
	BatchIndex int64 `json:"batch_index"`
}

type TransactionRequest struct {
	StoreNumber uint32 `json:"store_number"`
}

func (s *DaService) GetLatestTransactionBatchIndex(c echo.Context) error {
	batchIndex, err := s.Cfg.EigenContract.RollupBatchIndex(&bind.CallOpts{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("fail to get batch index"))
	}
	return c.JSON(http.StatusOK, batchIndex.Uint64())
}

func (s *DaService) GetRollupStoreByRollupBatchIndex(c echo.Context) error {
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

func (s *DaService) GetBatchTransactionByDataStoreId(c echo.Context) error {
	var txReq TransactionRequest
	if err := c.Bind(&txReq); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid request params"))
	}
	conn, err := grpc.Dial(s.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
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
		return c.JSON(http.StatusBadRequest, errors.New("recovery data fail"))
	}
	batchTxn := new([]common.BatchTx)
	batchRlpStream := rlp.NewStream(bytes.NewBuffer(reply.GetData()), 0)
	err = batchRlpStream.Decode(batchTxn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("decode batch tx fail"))
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
	return c.JSON(http.StatusOK, txList)
}
