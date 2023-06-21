package tokenprice

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mantlenetworkio/mantle/gas-oracle/bindings"
)

var (
	bit2wETHPairAddress  = common.HexToAddress("0x5c128d25a21f681e678cb050e551a895c9309945")
	usdt2wETHPairAddress = common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6")

	wETHAddress     = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	bitTokenAddress = common.HexToAddress("0x1A4b46696b2bB4794Eb3D4c26f1c55F9170fa4C5")
	usdtAddress     = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")

	bitTokenDecimals = floatStringToBigInt("1", 18)
	usdtDecimals     = floatStringToBigInt("1", 6)

	ContractV3Quoter = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
)

type uniswapClient struct {
	uniswapV3Quoter *bindings.Uniswapv3Quoter
}

func newUniswapClient(uniswapURL string) (*uniswapClient, error) {
	l1Client, err := ethclient.Dial(uniswapURL)
	if err != nil {
		return nil, err
	}

	uniswapV3Quoter, err := bindings.NewUniswapv3Quoter(ContractV3Quoter, l1Client)
	if err != nil {
		return nil, err
	}

	return &uniswapClient{
		uniswapV3Quoter: uniswapV3Quoter,
	}, nil
}

func (c *Client) getTokenPricesFromUniswap() (float64, float64, error) {

	eth2bitPrice, err := c.getTokenPriceFromUniswap(wETHAddress, bitTokenAddress, bitTokenDecimals)
	if err != nil {
		return 0, 0, err
	}
	eth2usdtPrice, err := c.getTokenPriceFromUniswap(wETHAddress, usdtAddress, usdtDecimals)
	if err != nil {
		return 0, 0, err
	}

	tokenRatio := determineTokenRatio(float64(eth2bitPrice))
	ethPrice, _ := determineETHPrice(big.NewFloat(float64(eth2usdtPrice))).Float64()

	return tokenRatio, ethPrice, nil
}

// getTokenPriceFromUniswap estimate to execute swapping from_token to to_token to get token price
func (c *Client) getTokenPriceFromUniswap(fromToken, toToken common.Address, decimals *big.Int) (int64, error) {
	fee := big.NewInt(3000)
	fromAmount := floatStringToBigInt("1.00", 18)
	sqrtPriceLimitX96 := big.NewInt(0)

	var out []interface{}
	rawCaller := &bindings.Uniswapv3QuoterRaw{Contract: c.uniswapQuoterClient.uniswapV3Quoter}
	err := rawCaller.Call(nil, &out, "quoteExactInputSingle", fromToken, toToken,
		fee, fromAmount, sqrtPriceLimitX96)
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(out[0].(*big.Int), decimals).Int64(), nil
}

func floatStringToBigInt(amount string, decimals int) *big.Int {
	fAmount, _ := new(big.Float).SetString(amount)
	fi, _ := new(big.Float).Mul(fAmount, big.NewFloat(math.Pow10(decimals))).Int(nil)
	return fi
}
