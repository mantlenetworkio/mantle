package client

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var errDtlHTTPError = errors.New("dtl http error")

type RollupBatchIndex struct {
	BatchIndex uint64 `json:"batchIndex"`
}

type DtlClient interface {
	GetLatestTransactionBatchIndex() (uint64, error)
}

type Client struct {
	client *resty.Client
}

func NewDtlClient(url string) *Client {
	client := resty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errDtlHTTPError)
		}
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetLatestTransactionBatchIndex() (uint64, error) {
	var rollupBatchIndex RollupBatchIndex
	response, err := c.client.R().
		SetResult(&rollupBatchIndex).
		Get("/da/getLatestTransactionBatchIndex")
	if err != nil {
		return 0, fmt.Errorf("cannot get latest batch index: %w", err)
	}
	if response.StatusCode() != 200 {
		return 0, errors.New("fetch latest batch index fail")
	}
	return rollupBatchIndex.BatchIndex, err
}
