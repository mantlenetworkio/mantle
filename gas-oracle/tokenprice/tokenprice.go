package tokenprice

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/go-resty/resty/v2"
)

type TokenRatioMode uint64

// Client is an HTTP based TokenPriceClient
type Client struct {
	client              *resty.Client
	uniswapQuoterClient *uniswapClient
	frequency           time.Duration
	lastRatio           float64
	lastEthPrice        float64
	lastUpdate          time.Time
	tokenRatioMode      TokenRatioMode
}

type TokenPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Result struct {
	RetCode int
	Result  TokenPrice
}

var (
	errHTTPError = errors.New("http error")

	// DefaultTokenRatio is eth_price / mnt_price
	DefaultTokenRatio = 4000
	// DefaultETHPrice is default eth_price
	// If SwitchOneDollarTokenRatio valid, use DefaultETHPrice to set token_ratio to mnt price is 1$
	DefaultETHPrice = 1800

	// RealTokenRatioMode use eth_price / mnt_price to set token_ratio
	RealTokenRatioMode = TokenRatioMode(0)
	// DefaultTokenRatioMode use DefaultTokenRatio to set token_ratio
	DefaultTokenRatioMode = TokenRatioMode(1)
	// OneDollarTokenRatioMode use eth_price to set token_ratio, so mnt price is 1$
	OneDollarTokenRatioMode = TokenRatioMode(2)
)

// NewClient create a new Client given a remote HTTP url, update frequency and different mode_switch for token ratio
func NewClient(url, uniswapURL string, frequency uint64, tokenRatioMode uint64) *Client {
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

	uniswapQuoterClient, err := newUniswapClient(uniswapURL)
	if err != nil {
		return nil
	}

	return &Client{
		client:              client,
		uniswapQuoterClient: uniswapQuoterClient,
		frequency:           time.Duration(frequency) * time.Second,
		tokenRatioMode:      TokenRatioMode(tokenRatioMode),
	}
}

func (c *Client) PriceRatioWithMode() (float64, error) {
	if time.Now().Sub(c.lastUpdate) < c.frequency {
		return c.lastRatio, nil
	}

	// Todo query token prices concurrent
	var ratio, ethPrice float64
	var errDex error
	// get token price from dex
	if ratio, ethPrice, errDex = c.getTokenPricesFromUniswap(); errDex != nil {
		ratio = float64(DefaultTokenRatio)
		ethPrice = float64(DefaultETHPrice)
	}

	// get token price from cex
	// if both dex and cex return token prices correctly, token price from cex will be used
	// if both failed, default value will be used
	if ratioCex, ethPriceCex, errCex := c.priceRatio(); errCex != nil && errDex != nil {
		ratio = float64(DefaultTokenRatio)
		ethPrice = float64(DefaultETHPrice)
	} else if errCex == nil {
		ratio = ratioCex
		ethPrice = ethPriceCex
	}

	switch c.tokenRatioMode {
	case DefaultTokenRatioMode:
		// use default eth/mnt price to set token ratio
		ratio = float64(DefaultTokenRatio)
	case OneDollarTokenRatioMode:
		// supposing that mnt is 1 USD, so token_ratio is equals to eth_price
		ratio = ethPrice
	}

	c.lastUpdate = time.Now()
	c.lastRatio = ratio
	c.lastEthPrice = ethPrice

	return ratio, nil
}

func (c *Client) priceRatio() (float64, float64, error) {
	ethPrice, err := c.query("ETHUSDT")
	if err != nil {
		return 0, 0, err
	}
	bitPrice, err := c.query("BITUSDT")
	if err != nil {
		return 0, 0, err
	}
	bigZero := big.NewFloat(0)
	if ethPrice.Cmp(bigZero) != 1 {
		return 0, 0, fmt.Errorf("invalid eth Price")
	}
	if bitPrice.Cmp(bigZero) != 1 {
		return 0, 0, fmt.Errorf("invalid mnt Price")
	}
	ratio, _ := new(big.Float).Quo(ethPrice, bitPrice).Float64()
	ethPriceInt64, _ := ethPrice.Float64()
	return ratio, ethPriceInt64, nil
}

func (c *Client) query(symbol string) (*big.Float, error) {
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
