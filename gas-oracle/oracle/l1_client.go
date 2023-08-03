package oracle

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	ometrics "github.com/mantlenetworkio/mantle/gas-oracle/metrics"
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
	ratio, err := c.tokenPricer.PriceRatioWithMode()
	if err != nil {
		ratio = tokenprice.DefaultTokenRatio
	}
	tip, err := c.Client.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if tip == nil {
		return nil, fmt.Errorf("get tip is nil")
	}
	log.Info("show base fee original", "tip.BaseFee", tip.BaseFee, "number", tip.Number, "ratio", ratio)
	// get tip
	gasTipCap, err := c.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}
	// get history 20 block best base
	bestBaseFee := c.getHistoryBestPrice(tip.Number, tip.BaseFee, 20)
	tip.BaseFee = new(big.Int).Mul(new(big.Int).Add(bestBaseFee, gasTipCap), big.NewInt(int64(ratio)))
	log.Info("show base fee context", "bestBaseFee", bestBaseFee, "gasTipCap", gasTipCap, "ratio", ratio)
	ometrics.GasOracleStats.L1GasPriceGauge.Update(new(big.Int).Add(bestBaseFee, gasTipCap).Int64())
	ometrics.GasOracleStats.TokenRatioGauge.Update(ratio)
	return tip, nil
}

func (c *L1Client) getHistoryBestPrice(endHeight *big.Int, lastBaseFee *big.Int, countWindow int) *big.Int {
	var baseFees = make([]*big.Int, 0)
	var bestPrice = new(big.Int)
	var wg = sync.WaitGroup{}
	// get base fee
	for i := 0; i < countWindow; i++ {
		wg.Add(1)
		go func(i int) {
			header, err := c.Client.HeaderByNumber(context.Background(), new(big.Int).Sub(endHeight, new(big.Int).SetInt64(int64(i))))
			if err == nil && header.BaseFee != nil {
				baseFees = append(baseFees, header.BaseFee)
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	// get best base fee, append last base fee again, incase get base fees all in error
	baseFees = append(baseFees, lastBaseFee)
	for j := 0; j < len(baseFees); j++ {
		if bestPrice.Cmp(baseFees[j]) < 0 {
			bestPrice.Set(baseFees[j])
		}
	}
	return bestPrice
}
