package tokenprice

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTokenPrice(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0)
	ethPrice, err := tokenPricer.query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	bitPrice, err := tokenPricer.query("BITUSDT")
	require.NoError(t, err)
	t.Logf("BIT price:%v", bitPrice)

	t.Log(ethPrice.Quo(ethPrice, bitPrice))
	ratio, _, err := tokenPricer.priceRatio()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)

	ratioFromUniswap, err := tokenPricer.getTokenPriceFromUniswap(wETHAddress, bitTokenAddress, bitTokenDecimals)
	require.NoError(t, err)
	t.Logf("ETH/BIT:%v", ratioFromUniswap)

	ratioFromUniswap, err = tokenPricer.getTokenPriceFromUniswap(wETHAddress, usdtAddress, usdtDecimals)
	require.NoError(t, err)
	t.Logf("ETH/USDT:%v", ratioFromUniswap)
}

func TestGetTokenPriceWithRealTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 1)

	ethPrice, err := tokenPricer.query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode2(t *testing.T) {
	tokenPricer := NewClient("", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 1)

	_, ethPrice, err := tokenPricer.getTokenPricesFromUniswap()
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithOneDollarTokenRatioMode3(t *testing.T) {
	tokenPricer := NewClient("", "https://mainnet.infura.io/v3", 3, 1)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	require.Equal(t, DefaultETHPrice, big.NewFloat(ratio))
	t.Logf("ratio:%v", ratio)
}

func TestGetTokenPriceWithDefaultTokenRatioMode(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 2)

	ratio, err := tokenPricer.PriceRatioWithMode()
	require.NoError(t, err)
	require.Equal(t, DefaultTokenRatio, ratio)
	t.Logf("ratio:%v", ratio)
}
