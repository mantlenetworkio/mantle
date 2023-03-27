package tss_client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var errTssHTTPError = errors.New("tss http error")

type TssClient interface {
	GetSignOutput(output common.SignOutputRequest) ([]byte, error)
}

type Client struct {
	client *resty.Client
}

type TssResponse struct {
	Signature []byte `json:"signature"`
	RollBack  bool   `json:"roll_back"`
}

func NewClient(url string) *Client {
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

func (c *Client) GetSignStateBatch(output common.SignOutputRequest) ([]byte, error) {
	var signature []byte
	response, err := c.client.R().
		SetBody(map[string]interface{}{"output_root": output.OutputRoot, "l2_block_number": output.L2BlockNumber, "l1_block_number": output.L1BlockNumber, "l1_block_hash": output.L1BlockHash}).
		SetResult(signature).
		Post("/api/v1/sign/state")
	if err != nil {
		return nil, fmt.Errorf("cannot get signature: %w", err)
	}
	if response.StatusCode() == 200 {
		return response.Body(), nil
	} else {
		return nil, errors.New("fetch tss manager signature faill")
	}
}
