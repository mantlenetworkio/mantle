package tokenprice

import (
	"fmt"
	"math/big"
)

type TokenPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Result struct {
	RetCode int
	Result  TokenPrice
}

func (c *Client) query(symbol string) (float64, error) {
	response, err := c.client.R().
		SetResult(&Result{}).
		SetQueryParams(map[string]string{
			"symbol": symbol,
		}).
		Get("/spot/quote/v1/ticker/price")
	if err != nil {
		return 0, fmt.Errorf("cannot fetch token price result: %w", err)
	}
	result, ok := response.Result().(*Result)
	if !ok {
		return 0, fmt.Errorf("cannot parse result")
	}
	if result.Result.Price == "" {
		return 0, fmt.Errorf("empty price")
	}
	priceBigFloat, _ := big.NewFloat(0).SetString(result.Result.Price)
	priceFloat64, _ := priceBigFloat.Float64()
	return priceFloat64, nil
}
