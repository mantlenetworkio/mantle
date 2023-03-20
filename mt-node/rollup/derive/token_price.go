package derive

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	errHTTPError             = errors.New("http error")
	defaultFrequency  uint64 = 10
	defaultSourceName        = "bybit"
)

// NewTPClient create a new TPClient given a remote HTTP url and update frequency
func NewTPClient(url string, sourceName string, frequency uint64) *TPClient {
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
	if frequency == 0 {
		frequency = defaultFrequency
	}
	if sourceName == "" {
		sourceName = defaultSourceName
	}

	return &TPClient{
		client:    client,
		frequency: time.Duration(frequency) * time.Second,
	}
}

// TPClient is an HTTP based TokenPriceClient
type TPClient struct {
	client     *resty.Client
	frequency  time.Duration
	sourceName string
	lastRatio  uint64
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

func (c *TPClient) Query(tokenA, tokenB string) (*big.Float, error) {
	switch c.sourceName {
	case defaultSourceName:
		return c.bybitQuery(tokenA, tokenB)

	default:
		return nil, errors.New("no support token price source")
	}
}

func (c *TPClient) bybitQuery(tokenA, tokenB string) (*big.Float, error) {
	response, err := c.client.R().
		SetResult(&Result{}).
		SetQueryParams(map[string]string{
			"symbol": strings.ToUpper(tokenA) + strings.ToUpper(tokenB),
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

func (c *TPClient) PriceRatio() (uint64, error) {
	if time.Now().Sub(c.lastUpdate) < c.frequency {
		return c.lastRatio, nil
	}
	ethPrice, err := c.Query("ETH", "USDT")
	if err != nil {
		return 0, err
	}
	bitPrice, err := c.Query("BIT", "USDT")
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
	c.lastRatio = uint64(math.Ceil(ratio))
	return c.lastRatio, nil
}
