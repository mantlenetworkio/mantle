package genesis

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/mt-bindings/bindings"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func initTestAccountForEthBit(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)
}

func TestCheckBalance(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// init balance
	l1Eth := getETHBalanceFromL1(t, userAddress)
	l1Bit := getBITBalanceFromL1(t, userAddress)
	t.Log("find l1 bit balance : ", l1Bit)
	t.Log("find l1 eth balance : ", l1Eth)

	decimal1 := big.NewInt(DECIMAL1)

	if l1Eth.Cmp(decimal1) < 0 {
		delta := big.NewInt(0)
		transferETH(t, l1Client, common.HexToAddress(userAddress), delta.Sub(decimal1, l1Eth).Int64())
		l1Eth = getETHBalanceFromL1(t, userAddress)
	}
	if l1Bit.Cmp(decimal1) < 0 {
		t.Log("start mint bit")
		delta := big.NewInt(0)
		mintBIT(t, l1Client, userPrivateKey, delta.Sub(decimal1, l1Bit).Int64())
	}
	l1Eth = getETHBalanceFromL1(t, userAddress)
	if l1Eth.Cmp(decimal1) < 0 {
		delta := big.NewInt(0)
		transferETH(t, l1Client, common.HexToAddress(userAddress), delta.Sub(decimal1, l1Eth).Int64())
		l1Eth = getETHBalanceFromL1(t, userAddress)
	}

	t.Log("L1 BALANCE INFO")
	l1Eth = getETHBalanceFromL1(t, userAddress)
	l1Bit = getBITBalanceFromL1(t, userAddress)
	require.Equal(t, int64(DECIMAL1), l1Eth.Int64())
	require.Equal(t, int64(DECIMAL1), l1Bit.Int64())
	t.Log("balance eth: ", l1Eth.Int64())
	t.Log("balance bit: ", l1Bit.Int64())
}

func TransferETH(t *testing.T, client *ethclient.Client, address common.Address, amount int64) {
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	require.NoError(t, err)

	value := big.NewInt(amount) // in wei (1 eth)
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	var data []byte
	tx := types.NewTransaction(nonce, address, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	require.NoError(t, err)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	require.NoError(t, err)

	err = client.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
}

func MintBIT(t *testing.T, client *ethclient.Client, privateKey string, amount int64) {
	l1bitToken, err := bindings.NewBitTokenERC20(common.HexToAddress(l1BitAddress), client)
	require.NoError(t, err)
	auth := buildL1Auth(t, client, privateKey, big.NewInt(0))
	tx, err := l1bitToken.Mint(auth, big.NewInt(amount))
	require.NoError(t, err)
	require.NotNil(t, tx)
}
