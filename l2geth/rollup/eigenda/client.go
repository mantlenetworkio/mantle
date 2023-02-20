package eigenda

import (
	"bytes"
	"encoding/json"
	"fmt"
	gresty "github.com/go-resty/resty/v2"
	common2 "github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/pkg/errors"
)

var errTssHTTPError = errors.New("eigen da http error")

type EigenClient interface {
	GetLatestTransactionBatchIndex() (*uint64, error)
	GetRollupStoreByRollupBatchIndex(batchIndex int64) (*RollupStoreResponse, error)
	GetBatchTransactionByDataStoreId(storeNumber uint32, l1MsgSender string) ([]*types.Transaction, error)
}

type Client struct {
	client *gresty.Client
}

func NewEigenClient(url string) *Client {
	client := gresty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *gresty.Client, r *gresty.Response) error {
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

func (c *Client) GetRollupStoreByRollupBatchIndex(batchIndex int64) (*RollupStoreResponse, error) {
	response, err := c.client.R().
		SetBody(map[string]interface{}{"batch_index": batchIndex}).
		SetResult(&RollupStoreResponse{}).
		Post("/eigen/getRollupStoreByRollupBatchIndex")
	if err != nil {
		return nil, fmt.Errorf("cannot get roll store data: %w", err)
	}
	rollupStore, ok := response.Result().(*RollupStoreResponse)
	if response.StatusCode() != 200 || !ok {
		return nil, errors.New("fetch roll store data fail")
	}
	return rollupStore, nil
}

func (c *Client) GetBatchTransactionByDataStoreId(storeNumber uint32, l1MsgSender string) ([]*types.Transaction, error) {
	var TxListBuf []byte
	response, err := c.client.R().
		SetBody(map[string]interface{}{"store_number": storeNumber}).
		SetResult(&TxListBuf).
		Post("/eigen/getBatchTransactionByDataStoreId")
	if err != nil {
		return nil, fmt.Errorf("cannot get tx list: %w", err)
	}
	if response.StatusCode() == 200 {
		var retTxList []*types.Transaction
		batchTxn := new([]BatchTx)
		batchRlpStream := l2rlp.NewStream(bytes.NewBuffer(TxListBuf), 0)
		err = batchRlpStream.Decode(batchTxn)
		if err != nil {
			return nil, fmt.Errorf("decode batch tx fail: %w", err)
		}
		newBatchTxn := *batchTxn
		for i := 0; i < len(newBatchTxn); i++ {
			var l2Tx types.Transaction
			rlpStream := l2rlp.NewStream(bytes.NewBuffer(newBatchTxn[i].RawTx), 0)
			if err := l2Tx.DecodeRLP(rlpStream); err != nil {
				log.Error("Decode RLP fail")
			}
			txDecodeMetaData := new(TransactionMeta)
			err := json.Unmarshal(newBatchTxn[i].TxMeta, txDecodeMetaData)
			if err != nil {
				log.Error("Unmarshal json fail")
			}
			var queueOrigin types.QueueOrigin
			var l1MessageSender *common2.Address
			if txDecodeMetaData.QueueIndex == nil {
				queueOrigin = types.QueueOriginSequencer
				l1MessageSender = nil
			} else {
				queueOrigin = types.QueueOriginL1ToL2
				addrLs := common2.HexToAddress(l1MsgSender)
				l1MessageSender = &addrLs
			}
			realTxMeta := &types.TransactionMeta{
				L1BlockNumber:   txDecodeMetaData.L1BlockNumber,
				L1Timestamp:     txDecodeMetaData.L1Timestamp,
				L1MessageSender: l1MessageSender,
				QueueOrigin:     queueOrigin,
				Index:           txDecodeMetaData.Index,
				QueueIndex:      txDecodeMetaData.QueueIndex,
				RawTransaction:  txDecodeMetaData.RawTransaction,
			}
			l2Tx.SetTransactionMeta(realTxMeta)
			retTxList = append(retTxList, &l2Tx)
		}
		return retTxList, nil
	} else {
		return nil, errors.New("fetch tx list fail")
	}
}
