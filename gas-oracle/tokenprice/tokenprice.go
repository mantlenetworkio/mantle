package tokenprice

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/go-resty/resty/v2"
)

var errHTTPError = errors.New("http error")

// NewClient create a new Client given a remote HTTP url and update frequency
func NewClient(url string, frequency uint64) *Client {
	client := resty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errHTTPError)
		}
		return nil
	})

	return &Client{
		client:    client,
		frequency: time.Duration(frequency) * time.Second,
	}
}

// Client is an HTTP based TokenPriceClient
type Client struct {
	client     *resty.Client
	frequency  time.Duration
	lastRatio  float64
	lastUpdate time.Time
}

type TokenPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Result struct {
	RetCode int
	Result  TokenPrice
}

func (c *Client) Query(symbol string) (*big.Float, error) {
	response, err := c.client.R().
		SetResult(&Result{}).
		SetQueryParams(map[string]string{
			"symbol": symbol,
		}).
		Get("/spot/quote/v1/ticker/price")
	if err != nil {
		return nil, fmt.Errorf("cannot fetch token price result: %w", err)
	}
	result, ok := response.Result().(*Result)
	if !ok {
		return nil, fmt.Errorf("cannot parse result")
	}
	if result.Result.Price == "" {
		return nil, fmt.Errorf("empty price")
	}
	bigPrice, _ := big.NewFloat(0).SetString(result.Result.Price)
	return bigPrice, nil
}

func (c *Client) PriceRatio() (float64, error) {
	if time.Now().Sub(c.lastUpdate) < c.frequency {
		return c.lastRatio, nil
	}
	ethPrice, err := c.Query("ETHUSDT")
	if err != nil {
		return 0, err
	}
	bitPrice, err := c.Query("BITUSDT")
	if err != nil {
		return 0, err
	}
	bigZero := big.NewFloat(0)
	if ethPrice.Cmp(bigZero) != 1 {
		return 0, fmt.Errorf("invalid eth Price")
	}
	if bitPrice.Cmp(bigZero) != 1 {
		return 0, fmt.Errorf("invalid bit Price")
	}
	ratio, _ := ethPrice.Quo(ethPrice, bitPrice).Float64()
	c.lastUpdate = time.Now()
	c.lastRatio = ratio
	return c.lastRatio, nil
}
