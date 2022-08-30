package tokenprice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"time"
)

var (
	bybitUrl   = "https://api.bybit.com/spot/quote/v1/ticker/price"
	lastRatio  uint64
	lastUpdate time.Time
	frequency  = 3 * time.Second
)

type TokenPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Result struct {
	RetCode int
	Result  TokenPrice
}

func Query(symbol string) (*big.Float, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?symbol=%s", bybitUrl, symbol), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	cli := http.Client{
		Timeout: 45 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Result
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}
	if result.Result.Price == "" {
		return nil, fmt.Errorf("empty price")
	}
	bigPrice, _ := big.NewFloat(0).SetString(result.Result.Price)
	return bigPrice, nil
}

func PriceRatio() (uint64, error) {
	if time.Now().Sub(lastUpdate) < frequency {
		return lastRatio, nil
	}
	ethPrice, err := Query("ETHUSDT")
	if err != nil {
		return 0, err
	}
	bitPrice, err := Query("BITUSDT")
	if err != nil {
		return 0, err
	}
	bigZero := big.NewFloat(0)
	if ethPrice.Cmp(bigZero) != 1 {
		return 0, fmt.Errorf("invalid eth Price")
	}
	if bitPrice.Cmp(bigZero) != 1 {
		return 0, fmt.Errorf("invalid bit Price")
	}
	ratio, _ := ethPrice.Quo(ethPrice, bitPrice).Float64()
	lastUpdate = time.Now()
	lastRatio = uint64(math.Ceil(ratio))
	return lastRatio, nil
}
