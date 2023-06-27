package tokenprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getTokenPriceFromUniswap(t *testing.T) {
	tokenPricer := NewClient("", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, false)
	ethPrice, err := tokenPricer.getTokenPriceFromUniswap(wETHAddress, usdtAddress, usdtDecimals)
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	eth2bitPrice, err := tokenPricer.getTokenPriceFromUniswap(wETHAddress, bitTokenAddress, bitTokenDecimals)
	require.NoError(t, err)
	t.Logf("BIT price:%v", ethPrice/eth2bitPrice)
}
