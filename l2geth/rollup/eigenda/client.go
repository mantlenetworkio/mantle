package eigenda

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/pkg/errors"
)

var errTssHTTPError = errors.New("eigen da http error")

type EigenClient interface {
	GetLatestTransactionBatchIndex() (*uint64, error)
	GetRollupStoreByRollupBatchIndex(batchIndex int64) (*common.RollupStoreResponse, error)
	GetBatchTransactionByDataStoreId(storeNumber uint32) ([]*types.Transaction, error)
}

type Client struct {
	client *resty.Client
}

func NewEigenClient(url string) *Client {
	client := resty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errTssHTTPError)
		}
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetLatestTransactionBatchIndex() (batchIndex *uint64, err error) {
	var batchIndexTmp uint64
	response, err := c.client.R().
		SetResult(&batchIndexTmp).
		Get("/eigen/getLatestTransactionBatchIndex")
	if err != nil {
		return nil, fmt.Errorf("cannot get latest batch index: %w", err)
	}
	if response.StatusCode() != 200 {
		return nil, errors.New("fetch latest batch index fail")
	}
	return &batchIndexTmp, err
}

func (c *Client) GetRollupStoreByRollupBatchIndex(batchIndex int64) (*common.RollupStoreResponse, error) {
	response, err := c.client.R().
		SetBody(map[string]interface{}{"batch_index": batchIndex}).
		SetResult(&common.RollupStoreResponse{}).
		Post("/eigen/getRollupStoreByRollupBatchIndex")
	if err != nil {
		return nil, fmt.Errorf("cannot get roll store data: %w", err)
	}
	rollupStore, ok := response.Result().(*common.RollupStoreResponse)
	if response.StatusCode() != 200 || !ok {
		return nil, errors.New("fetch roll store data fail")
	}
	return rollupStore, nil
}

func (c *Client) GetBatchTransactionByDataStoreId(storeNumber uint32) ([]*types.Transaction, error) {
	var TxList []types.Transaction
	response, err := c.client.R().
		SetBody(map[string]interface{}{"store_number": storeNumber}).
		SetResult(&TxList).
		Post("/eigen/getBatchTransactionByDataStoreId")
	if err != nil {
		return nil, fmt.Errorf("cannot get tx list: %w", err)
	}
	if response.StatusCode() == 200 {
		var retTxList []*types.Transaction
		for _, tx := range TxList {
			log.Info("transaction index", "index", tx.GetMeta().Index)
			retTxList = append(retTxList, &tx)
		}
		return retTxList, nil
	} else {
		return nil, errors.New("fetch tx list fail")
	}
}
