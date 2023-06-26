package tokenprice

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/go-resty/resty/v2"
)

type TokenRatioMode uint64

// Client is an HTTP based TokenPriceClient
type Client struct {
	client               *resty.Client
	uniswapQuoterClient  *uniswapClient
	frequency            time.Duration
	lastRatio            float64
	lastEthPrice         float64
	lastMntPrice         float64
	lastUpdate           time.Time
	tokenRatioMode       TokenRatioMode
	tokenPairForMNTPrice string
}

var (
	errHTTPError = errors.New("http error")

	// DefaultTokenRatio is eth_price / mnt_price, 4000 = $1800/$0.45
	DefaultTokenRatio = float64(4000)
	// TokenRatioMax token_ratio upper bounds
	TokenRatioMax = float64(100000)
	// TokenRatioMin token_ratio lower bounds
	TokenRatioMin = float64(100)

	// DefaultETHPrice is default eth_price
	// If SwitchOneDollarTokenRatio valid, use DefaultETHPrice to set token_ratio to make mnt_price is 1$
	DefaultETHPrice = float64(1800)
	// ETHPriceMax eth_price upper bounds
	ETHPriceMax = float64(1000000)
	// ETHPriceMin eth_price lower bounds
	ETHPriceMin = float64(100)

	DefaultMNTPrice = 0.45
	// MNTPriceMax mnt_price upper bounds
	MNTPriceMax = float64(100)
	// MNTPriceMin mnt_price lower bounds
	MNTPriceMin = 0.01

	// RealTokenRatioMode use eth_price / mnt_price to set token_ratio
	RealTokenRatioMode = TokenRatioMode(0)
	// OneDollarTokenRatioMode use eth_price to set token_ratio, so mnt price is 1$
	OneDollarTokenRatioMode = TokenRatioMode(1)
	// DefaultTokenRatioMode use DefaultTokenRatio to set token_ratio
	DefaultTokenRatioMode = TokenRatioMode(2)

	// token pairs, used to query token pairs price
	// ETHUSDT used to query eth/usdt price
	ETHUSDT = "ETHUSDT"
	// BITUSDT used to query bit/usdt price
	BITUSDT = "BITUSDT"
	// MNTUSDT used to query mnt/usdt price
	MNTUSDT = "MNTUSDT"
)

// NewClient create a new Client given a remote HTTP url, update frequency and different mode_switch for token ratio
// tokenPairMNTMode(true/false) to choose if mnt_price is in production
func NewClient(url, uniswapURL string, frequency uint64, tokenRatioMode uint64, tokenPairMNTMode bool) *Client {
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

	uniswapQuoterClient, err := newUniswapClient(uniswapURL, tokenPairMNTMode)
	if err != nil {
		return nil
	}

	tokenPairForMNTPrice := determineTokenPairForMNT(tokenPairMNTMode)

	return &Client{
		client:               client,
		uniswapQuoterClient:  uniswapQuoterClient,
		frequency:            time.Duration(frequency) * time.Second,
		lastRatio:            DefaultTokenRatio,
		lastEthPrice:         DefaultETHPrice,
		lastMntPrice:         DefaultMNTPrice,
		tokenRatioMode:       TokenRatioMode(tokenRatioMode),
		tokenPairForMNTPrice: tokenPairForMNTPrice,
	}
}

func (c *Client) PriceRatioWithMode() (float64, error) {
	if time.Now().Sub(c.lastUpdate) < c.frequency {
		return c.lastRatio, nil
	}

	// Todo query token prices concurrent
	var mntPrices, ethPrices []float64
	// get token price from oracle1(dex)
	mntPrice1, ethPrice1 := c.getTokenPricesFromUniswap()
	mntPrices = append(mntPrices, mntPrice1)
	ethPrices = append(ethPrices, ethPrice1)
	log.Info("query prices from oracle1", "mnt_price", mntPrice1, "eth_price", ethPrice1)

	// get token price from oracle2(cex)
	mntPrice2, ethPrice2 := c.getTokenPricesFromCex()
	mntPrices = append(mntPrices, mntPrice2)
	ethPrices = append(ethPrices, ethPrice2)
	log.Info("query prices from oracle2", "mnt_price", mntPrice2, "eth_price", ethPrice2)

	// get token price from oracle3(cex)
	// Todo add a third oracle to query prices
	mntPrice3, ethPrice3 := c.getTokenPricesFromCex()
	mntPrices = append(mntPrices, mntPrice3)
	ethPrices = append(ethPrices, ethPrice3)
	log.Info("query prices from oracle3", "mnt_price", mntPrice3, "eth_price", ethPrice3)

	// median price for eth & mnt
	medianMNTPrice := getMedian(mntPrices)
	medianETHPrice := getMedian(ethPrices)

	// determine mnt_price, eth_price
	mntPrice := c.determineMNTPrice(medianMNTPrice)
	ethPrice := c.determineETHPrice(medianETHPrice)
	log.Info("prices after determine", "mnt_price", mntPrice, "eth_price", ethPrice)

	// calculate ratio
	ratio := c.determineTokenRatio(mntPrice, ethPrice)

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

	log.Info("token ratio", "token ratio", ratio)

	c.lastUpdate = time.Now()
	c.lastRatio = ratio
	c.lastEthPrice = ethPrice
	c.lastMntPrice = mntPrice

	return ratio, nil
}

func (c *Client) getTokenPricesFromCex() (float64, float64) {
	ethPrice, err := c.queryV5(ETHUSDT)
	if err != nil {
		log.Warn("get token prices", "query eth price error", err)
		return 0, 0
	}
	mntPrice, err := c.queryV5(c.tokenPairForMNTPrice)
	if err != nil {
		log.Warn("get token prices", "query mnt price error", err)
		return 0, ethPrice
	}

	return mntPrice, ethPrice
}

func (c *Client) determineMNTPrice(price float64) float64 {
	if price > MNTPriceMax || price < MNTPriceMin {
		return c.lastMntPrice
	}

	return price
}

func (c *Client) determineETHPrice(price float64) float64 {
	if price > ETHPriceMax || price < ETHPriceMin {
		return c.lastEthPrice
	}

	return price
}

func (c *Client) determineTokenRatio(mntPrice, ethPrice float64) float64 {
	// calculate [tokenRatioMin, tokenRatioMax]
	tokenRatioMin := getMax(c.lastRatio*0.95, TokenRatioMin)
	tokenRatioMax := getMin(c.lastRatio*1.05, TokenRatioMax)

	ratio := ethPrice / mntPrice
	if ratio <= tokenRatioMin {
		return tokenRatioMin
	}
	if ratio >= tokenRatioMax {
		return tokenRatioMax
	}

	return ratio
}

func determineTokenPairForMNT(tokenPairMNTMode bool) string {
	if tokenPairMNTMode {
		return MNTUSDT
	} else {
		return BITUSDT
	}
}

func getMedian(nums []float64) float64 {
	nonZeros := make([]float64, 0)
	for _, num := range nums {
		if num != 0 {
			nonZeros = append(nonZeros, num)
		}
	}
	sort.Float64s(nonZeros)
	if len(nonZeros) == 0 {
		return 0
	}
	return nonZeros[len(nonZeros)/2]
}

func getMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
