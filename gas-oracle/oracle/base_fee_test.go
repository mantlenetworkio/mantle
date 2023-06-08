package oracle

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"sync"
	"testing"
)

//func TestBaseFeeUpdate(t *testing.T) {
//	L1ClientUrl := ""
//	L1ClientUrl := ""
//
//	key, _ := crypto.GenerateKey()
//	sim, _ := newSimulatedBackend(key)
//	chain := sim.Blockchain()
//	tokenPricer := tokenprice.NewClient("https://api.bybit.com", 3)
//	opts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
//	addr, _, gpo, err := bindings.DeployGasPriceOracle(opts, sim, opts.From)
//	if err != nil {
//		t.Fatal(err)
//	}
//	sim.Commit()
//
//	cfg := &Config{
//		privateKey:            key,
//		l2ChainID:             big.NewInt(1337),
//		gasPriceOracleAddress: addr,
//		gasPrice:              big.NewInt(784637584),
//	}
//
//	update, err := wrapUpdateBaseFee(sim, sim, cfg)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// Get the initial base fee
//	l1BaseFee, err := gpo.L1BaseFee(&bind.CallOpts{})
//	if err != nil {
//		t.Fatal(err)
//	}
//	// base fee should start at 0
//	if l1BaseFee.Cmp(common.Big0) != 0 {
//		t.Fatal("does not start at 0")
//	}
//	// get the header to know what the base fee
//	// should be updated to
//	tip := chain.CurrentHeader()
//	if tip.BaseFee == nil {
//		t.Fatal("no base fee found")
//	}
//	ratio, err := tokenPricer.PriceRatio()
//	if err != nil {
//		t.Fatal(err)
//	}
//	tip.BaseFee = new(big.Int).Mul(tip.BaseFee, big.NewInt(int64(ratio)))
//	// Ensure that there is no false negative by
//	// checking that the values don't start out the same
//	if l1BaseFee.Cmp(tip.BaseFee) == 0 {
//		t.Fatal("values are already the same")
//	}
//	// Call the update function to do the update
//	if err := update(); err != nil {
//		t.Fatalf("cannot update base fee: %s", err)
//	}
//	sim.Commit()
//	// Check the updated base fee
//	l1BaseFee, err = gpo.L1BaseFee(&bind.CallOpts{})
//	if err != nil {
//		t.Fatal(err)
//	}
//	// the base fee should be equal to the value
//	// on the header
//	if tip.BaseFee.Cmp(l1BaseFee) != 0 {
//		t.Fatal("base fee not updated")
//	}
//}

func TestFee(t *testing.T) {
	client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/vX39fRJe5UvWip7sAj_nmSO9p7Z5Y2KI")
	if err != nil {
		log.Fatalln(err.Error())
	}
	//tip, err := rpcData.L1Client.HeaderByNumber(context.Background(), big.NewInt(8986262))
	ethPrice := big.NewFloat(1900.29)
	bitPrice := big.NewFloat(0.4716)
	ratio, _ := ethPrice.Quo(ethPrice, bitPrice).Float64()

	tip, err := client.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	// get tip
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	require.NoError(t, err)

	// get history 20 block best base
	bestBaseFee := getHistoryBestPrice(client, tip.Number, tip.BaseFee, 20)
	tip.BaseFee = new(big.Int).Mul(new(big.Int).Add(bestBaseFee, gasTipCap), big.NewInt(int64(ratio)))

	log.Println("L1BaseFee->", tip.BaseFee)
}

func getHistoryBestPrice(c *ethclient.Client, endHeight *big.Int, lastBaseFee *big.Int, countWindow int) *big.Int {
	var baseFees = make([]*big.Int, 0)
	var bestPrice = new(big.Int)
	var wg = sync.WaitGroup{}
	// get base fee
	for i := 0; i < countWindow; i++ {
		wg.Add(1)
		go func() {
			header, err := c.HeaderByNumber(context.Background(), endHeight.Sub(endHeight, new(big.Int).SetInt64(int64(i))))
			if err == nil && header.BaseFee != nil {
				baseFees = append(baseFees, header.BaseFee)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	// get best base fee, append last base fee again, incase get base fees all in error
	baseFees = append(baseFees, lastBaseFee)
	for j := 0; j < len(baseFees); j++ {
		if bestPrice.Cmp(baseFees[j]) < 0 {
			bestPrice = baseFees[j]
		}
	}
	return bestPrice
}
