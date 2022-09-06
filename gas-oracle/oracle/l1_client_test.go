package oracle

import (
	"context"
	"github.com/bitdao-io/bitnetwork/gas-oracle/tokenprice"
	"testing"
)

func TestNewClient(t *testing.T) {
	tokenPricer := tokenprice.NewClient("http://127.0.0.1:8000", 3)
	client, _ := NewL1Client("https://goerli.infura.io/v3/5d207effd0bb4c718cee75a49dbddfee", tokenPricer)
	tip, err := client.HeaderByNumber(context.Background(), nil)
	t.Logf("err:%+v", err)
	t.Logf("tip:%v,%v", tip.BaseFee, tip.Number)
}
