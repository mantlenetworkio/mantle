package tokenprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTokenPrice(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", 3)
	ethPrice, err := tokenPricer.Query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", ethPrice)

	bitPrice, err := tokenPricer.Query("BITUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", bitPrice)

	t.Log(ethPrice.Quo(ethPrice, bitPrice))
	ratio, err := tokenPricer.PriceRatio()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)

}
