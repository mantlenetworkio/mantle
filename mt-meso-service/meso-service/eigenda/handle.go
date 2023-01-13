package eigenda

import (
	"bytes"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/labstack/echo/v4"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/mantlenetworkio/mantle/mt-batcher/sequencer"
	"math/big"
	"net/http"
)

const (
	SelfServiceOK     = 2000
	SelfServiceError  = 4000
	SelfInvalidParams = 4001
)

type RollupStoreRequest struct {
	batchIndex *big.Int
}

type TransactionRequest struct {
	StoreNumber uint32 `json:"store_number"`
}

func (s *DaService) GetLatestTransactionBatchIndex(c echo.Context) error {
	batchIndex, err := s.EigenDaMiddleWare.GetRollupBatchIndex()
	if err != nil {
		return c.JSON(SelfServiceError, "get roll batch index fail")
	}
	return c.JSON(SelfServiceOK, batchIndex.String())
}

func (s *DaService) GetRollupStoreByRollupBatchIndex(c echo.Context) error {
	var rsReq RollupStoreRequest
	if err := c.Bind(&rsReq); err != nil {
		retValue := BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	rollupStore, err := s.EigenDaMiddleWare.GetRollupStoreByRollupBatchIndex(rsReq.batchIndex)
	if err != nil {
		return c.JSON(SelfServiceError, "get roll batch index fail")
	}
	return c.JSON(SelfServiceOK, rollupStore)
}

func (s *DaService) GetBatchTransactionByDataStoreId(c echo.Context) error {
	var txReq TransactionRequest
	if err := c.Bind(&txReq); err != nil {
		retValue := BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	data, err := s.EigenDaMiddleWare.GetBatchTransactionByStoreNumber(txReq.StoreNumber)
	batchTxn := new([]sequencer.BatchTx)
	batchRlpStream := rlp.NewStream(bytes.NewBuffer(data), 0)
	err = batchRlpStream.Decode(batchTxn)
	if err != nil {
		log.Error("decode batch tx fail")
		retValue := BaseResource(false, SelfServiceError, nil, "decode batch tx fail")
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
	retValue := BaseResource(true, SelfServiceOK, txList, "get transaction from eigen da success")
	return c.JSON(http.StatusOK, retValue)
}
