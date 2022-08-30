package tokenprice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTokenPrice(t *testing.T) {
	ethPrice, err := Query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", ethPrice)

	bitPrice, err := Query("BITUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", bitPrice)

	t.Log(ethPrice.Quo(ethPrice, bitPrice))
	ratio, err := PriceRatio()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)

}
