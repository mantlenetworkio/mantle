package tokenprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTokenPrice(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, false)
	ethPrice, err := tokenPricer.query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	bitPrice, err := tokenPricer.query("BITUSDT")
	require.NoError(t, err)
	t.Logf("BIT price:%v", bitPrice)

	t.Logf("ratio:%v", ethPrice/bitPrice)
	bitPrice, ethPrice = tokenPricer.getTokenPricesFromCex()
	t.Logf("ETH price:%v", ethPrice)
	t.Logf("BIT price:%v", bitPrice)
	t.Logf("ratio:%v", ethPrice/bitPrice)

	eth2bitPrice, err := tokenPricer.getTokenPriceFromUniswap(wETHAddress, bitTokenAddress, bitTokenDecimals)
	require.NoError(t, err)
	t.Logf("ETH/BIT:%v", eth2bitPrice)

	ethPrice, err = tokenPricer.getTokenPriceFromUniswap(wETHAddress, usdtAddress, usdtDecimals)
	require.NoError(t, err)
	t.Logf("ETH/USDT:%v", ethPrice)
}

func TestGetTokenPriceWithRealTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 1, false)

	ethPrice, err := tokenPricer.query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode2(t *testing.T) {
	tokenPricer := NewClient("", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 1, false)

	_, ethPrice := tokenPricer.getTokenPricesFromUniswap()
	t.Logf("ETH price:%v", ethPrice)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode3(t *testing.T) {
	tokenPricer := NewClient("", "https://mainnet.infura.io/v3", 3, 1, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	require.Equal(t, DefaultETHPrice, ratio)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithDefaultTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 2, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	require.Equal(t, DefaultTokenRatio, ratio)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithNoSource(t *testing.T) {
	// source url are both invalid, so can not access correct prices
	tokenPricer := NewClient("https://api.bybit.co", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c232", 3, 0, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	require.Equal(t, DefaultTokenRatio, ratio)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOnlySource1(t *testing.T) {
	// uniswapURL is invalid, so can not access correct prices
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c232", 3, 0, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOnlySource2(t *testing.T) {
	// only uniswapURL is valid
	tokenPricer := NewClient("https://api.bybit.co", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, false)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithMNT(t *testing.T) {
	// only uniswapURL is valid
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, true)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func Test_getMedian(t *testing.T) {
	result := getMedian([]float64{0, 0, 0})
	require.Equal(t, float64(0), result)

	result = getMedian([]float64{1.1, 0, 0})
	require.Equal(t, 1.1, result)

	result = getMedian([]float64{1.1, 2.1, 0})
	require.Equal(t, 2.1, result)

	result = getMedian([]float64{2.1, 1.1})
	require.Equal(t, 2.1, result)

	result = getMedian([]float64{1.1, 2.1, 3.1})
	require.Equal(t, 2.1, result)

	result = getMedian([]float64{1.1, 3.1, 2.1})
	require.Equal(t, 2.1, result)

	result = getMedian([]float64{1.1, 3.1, 2.1, 4.1})
	require.Equal(t, 3.1, result)
}

func Test_getMax(t *testing.T) {
	result := getMax(1.1, 2.1)
	require.Equal(t, 2.1, result)
}

func Test_getMin(t *testing.T) {
	result := getMin(1.1, 2.1)
	require.Equal(t, 1.1, result)
}
