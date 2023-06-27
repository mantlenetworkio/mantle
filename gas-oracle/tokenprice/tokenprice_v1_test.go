package tokenprice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTokenPriceV1(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", "https://mainnet.infura.io/v3/4f4692085f1340c2a645ae04d36c2321", 3, 0, false)
	ethPrice, err := tokenPricer.query(ETHUSDT)
	require.NoError(t, err)
	t.Logf("ETH price:%v", ethPrice)

	bitPrice, err := tokenPricer.query(BITUSDT)
	require.NoError(t, err)
	t.Logf("BIT price:%v", bitPrice)
}
