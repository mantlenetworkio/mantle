package tokenprice

import (
	"fmt"
	"math/big"
)

var (
	noHTTPError = 0
)

type Response struct {
	RetCode     int         `json:"retCode"`
	RetMsg      string      `json:"retMsg"`
	PriceResult PriceResult `json:"result"`
	RetExtInfo  struct{}    `json:"retExtInfo"`
	Time        int64       `json:"time"`
}

type PriceResult struct {
	Category string `json:"category"`
	List     []Item `json:"list"`
}

type Item struct {
	Symbol     string `json:"symbol"`
	LastPrice  string `json:"lastPrice"`
	IndexPrice string `json:"indexPrice"`
}

func (c *Client) queryV5(symbol string) (float64, error) {
	response, err := c.client.R().
		SetResult(&Response{}).
		SetQueryParams(map[string]string{
			"symbol": symbol,
		}).
		Get("v5/market/tickers?category=linear&")
	if err != nil {
		return 0, fmt.Errorf("cannot fetch token price result: %w", err)
	}
	result, ok := response.Result().(*Response)
	if !ok {
		return 0, fmt.Errorf("cannot parse result")
	}
	if result.RetCode != noHTTPError {
		return 0, fmt.Errorf("query error")
	}
	if len(result.PriceResult.List) == 0 {
		return 0, fmt.Errorf("empty price in result")
	}
	priceBigFloat, _ := big.NewFloat(0).SetString(result.PriceResult.List[0].IndexPrice)
	priceFloat64, _ := priceBigFloat.Float64()
	return priceFloat64, nil
}
