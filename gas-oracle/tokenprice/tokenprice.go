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

	// DefaultTokenRatio is eth_price / mnt_price, 4000 = $1800/$0.45
	DefaultTokenRatio = float64(4000)
	// TokenRatioMax token_ratio upper bounds
	TokenRatioMax = float64(10000)
	// TokenRatioMin token_ratio lower bounds
	TokenRatioMin = float64(500)

	// DefaultETHPrice is default eth_price
	// If SwitchOneDollarTokenRatio valid, use DefaultETHPrice to set token_ratio to make mnt_price is 1$
	DefaultETHPrice = big.NewFloat(1800)
	// ETHPriceMax eth_price upper bounds
	ETHPriceMax = big.NewFloat(20000)
	// ETHPriceMin eth_price lower bounds
	ETHPriceMin = big.NewFloat(100)

	DefaultMNTPrice = big.NewFloat(0.45)
	// MNTPriceMax mnt_price upper bounds
	MNTPriceMax = big.NewFloat(10)
	// MNTPriceMin mnt_price lower bounds
	MNTPriceMin = big.NewFloat(0.2)

	// RealTokenRatioMode use eth_price / mnt_price to set token_ratio
	RealTokenRatioMode = TokenRatioMode(0)
	// OneDollarTokenRatioMode use eth_price to set token_ratio, so mnt price is 1$
	OneDollarTokenRatioMode = TokenRatioMode(1)
	// DefaultTokenRatioMode use DefaultTokenRatio to set token_ratio
	DefaultTokenRatioMode = TokenRatioMode(2)
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
		ratio = DefaultTokenRatio
		ethPrice, _ = DefaultETHPrice.Float64()
	}

	// get token price from cex
	// if both dex and cex return token prices correctly, token price from cex will be used
	// if cex failed, token price from dex will be used
	// if both failed, default value will be used
	if ratioCex, ethPriceCex, errCex := c.priceRatio(); errCex != nil && errDex != nil {
		ratio = DefaultTokenRatio
		ethPrice, _ = DefaultETHPrice.Float64()
	} else if errCex == nil {
		ratio = ratioCex
		ethPrice = ethPriceCex
	}

	switch c.tokenRatioMode {
	case DefaultTokenRatioMode:
		// use default eth/mnt price to set token ratio
		ratio = DefaultTokenRatio
	case OneDollarTokenRatioMode:
		// supposing that mnt is 1 USD, so token_ratio is equals to eth_price
		ratio = ethPrice
	default:
		// default mode is RealTokenRatioMode which uses eth_price / mnt_price to set token_ratio
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
	mntPrice, err := c.query("BITUSDT")
	if err != nil {
		return 0, 0, err
	}

	ethPrice = determineETHPrice(ethPrice)
	mntPrice = determineMNTPrice(mntPrice)

	ratio, _ := new(big.Float).Quo(ethPrice, mntPrice).Float64()
	ratio = determineTokenRatio(ratio)

	ethPriceFloat64, _ := ethPrice.Float64()

	return ratio, ethPriceFloat64, nil
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

func determineMNTPrice(price *big.Float) *big.Float {
	if price.Cmp(MNTPriceMax) == 1 || price.Cmp(MNTPriceMin) == -1 {
		return DefaultMNTPrice
	}

	return price
}

func determineETHPrice(price *big.Float) *big.Float {
	if price.Cmp(ETHPriceMax) == 1 || price.Cmp(ETHPriceMin) == -1 {
		return DefaultETHPrice
	}

	return price
}

func determineTokenRatio(ratio float64) float64 {
	if ratio < TokenRatioMin || ratio > TokenRatioMax {
		return DefaultTokenRatio
	}

	return ratio
}
