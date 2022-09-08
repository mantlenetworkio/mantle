package oracle

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/gas-oracle/tokenprice"
)

type L1Client struct {
	*ethclient.Client
	tokenPricer *tokenprice.Client
}

func NewL1Client(ethereumHttpUrl string, tokenPricer *tokenprice.Client) (*L1Client, error) {
	l1Client, err := ethclient.Dial(ethereumHttpUrl)
	if err != nil {
		return nil, err
	}
	return &L1Client{
		Client:      l1Client,
		tokenPricer: tokenPricer,
	}, nil
}

func (c *L1Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ratio, err := c.tokenPricer.PriceRatio()
	if err != nil {
		return nil, err
	}
	tip, err := c.Client.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if tip == nil {
		return tip, nil
	}
	tip.BaseFee = new(big.Int).Mul(tip.BaseFee, big.NewInt(int64(ratio)))
	return tip, nil
}
