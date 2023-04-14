package rollup

import (
	"errors"
	"fmt"
	gresty "github.com/go-resty/resty/v2"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"math/big"
	"strconv"
)

var errDtlHTTPError = errors.New("dtl da http error")

type RollupBatchIndex struct {
	BatchIndex uint64 `json:"batchIndex"`
}

type StoreResponse struct {
	OriginDataStoreId uint32 `json:"origin_data_store_id"`
	DataStoreId       uint32 `json:"data_store_id"`
	ConfirmAt         uint32 `json:"confirm_at"`
	Status            uint8  `json:"status"`
}

type DataStore struct {
	Index       uint64 `json:"index"`
	DataStoreId uint32 `json:"data_store_id"`
	Status      uint8  `json:"status"`
	ConfirmAt   uint32 `json:"confirm_at"`
}

type DataStoreId struct {
	BatchIndex          uint64    `json:"batchIndex"`
	BatchIndexDataStore DataStore `json:"dataStore"`
}

type TransactionBatchResponseV2 struct {
	DsId         uint64        `json:"dsId"`
	Transactions []transaction `json:"batchTx"`
}

type DtlEigenClient interface {
	GetDtlLatestBatchIndex() (*uint64, error)
	GetDtlRollupStoreByBatchIndex(batchIndex int64) (*StoreResponse, error)
	GetDtlBatchTransactionByDataStoreId(storeNumber uint32) ([]*types.Transaction, error)
}

type DtlClient struct {
	client  *gresty.Client
	chainID *big.Int
}

func NewDtlEigenClient(url string, chainID *big.Int) *DtlClient {
	client := gresty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *gresty.Client, r *gresty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errDtlHTTPError)
		}
		return nil
	})
	return &DtlClient{
		client:  client,
		chainID: chainID,
	}
}

func (dc *DtlClient) GetDtlLatestBatchIndex() (batchIndex *uint64, err error) {
	var rollupBatchIndex RollupBatchIndex
	response, err := dc.client.R().
		SetResult(&rollupBatchIndex).
		Get("/da/getLatestTransactionBatchIndex")
	if err != nil {
		return nil, fmt.Errorf("cannot get latest batch index: %w", err)
	}
	if response.StatusCode() != 200 {
		return nil, errors.New("fetch latest batch index fail")
	}
	return &rollupBatchIndex.BatchIndex, err
}

func (dc *DtlClient) GetDtlRollupStoreByBatchIndex(batchIndex int64) (*StoreResponse, error) {
	var dataStoreId DataStoreId
	response, err := dc.client.R().
		SetPathParams(map[string]string{"batchIndex": strconv.FormatInt(batchIndex, 10)}).
		SetResult(&dataStoreId).
		Get("/da/getDataStoreListByBatchIndex/{batchIndex}")
	if err != nil {
		return nil, fmt.Errorf("cannot get roll store data: %w", err)
	}
	if response.StatusCode() != 200 {
		return nil, errors.New("fetch roll store data fail")
	}
	rollupStore := &StoreResponse{
		OriginDataStoreId: dataStoreId.BatchIndexDataStore.DataStoreId,
		DataStoreId:       dataStoreId.BatchIndexDataStore.DataStoreId,
		ConfirmAt:         dataStoreId.BatchIndexDataStore.ConfirmAt,
		Status:            dataStoreId.BatchIndexDataStore.Status,
	}
	return rollupStore, nil
}

func (dc *DtlClient) GetDtlBatchTransactionByDataStoreId(storeNumber uint32) ([]*types.Transaction, error) {
	response, err := dc.client.R().
		SetPathParams(map[string]string{"dsId": strconv.FormatInt(int64(storeNumber), 10)}).
		SetResult(&TransactionBatchResponseV2{}).
		Get("/da/getBatchTxsByDataStoreId/{dsId}")
	if err != nil {
		return nil, fmt.Errorf("cannot get tx list: %w", err)
	}
	if response.StatusCode() == 200 {
		txBatch, ok := response.Result().(*TransactionBatchResponseV2)
		if !ok {
			return nil, fmt.Errorf("cannot parse transaction batch response")
		}
		return parseTransactionBatchResponseV2(txBatch, dc.chainID)
	} else {
		return nil, errors.New("fetch tx list fail")
	}
}

func parseTransactionBatchResponseV2(txBatch *TransactionBatchResponseV2, chainID *big.Int) ([]*types.Transaction, error) {
	if txBatch == nil {
		return nil, errors.New("txBatch is nil")
	}
	var txs []*types.Transaction
	for i := 0; i < len(txBatch.Transactions); i++ {
		transaction, err := batchedTransactionToTransaction(&txBatch.Transactions[i], chainID)
		if err != nil {
			return nil, err
		}
		txs = append(txs, transaction)
	}
	return txs, nil
}
