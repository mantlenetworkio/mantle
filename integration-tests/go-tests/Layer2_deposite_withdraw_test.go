package go_tests

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	l1bit "github.com/mantlenetworkio/mantle/go-test/contracts/L1/local/LocalBitToken.sol"
	l1bridge "github.com/mantlenetworkio/mantle/go-test/contracts/L1/messaging/L1StandardBridge.sol"
	l2bridge "github.com/mantlenetworkio/mantle/go-test/contracts/L2/messaging/L2StandardBridge.sol"
	l2eth "github.com/mantlenetworkio/mantle/go-test/contracts/L2/predeploys/BVM_ETH.sol"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

const (
	l1url           = "http://localhost:9545"
	l2url           = "http://localhost:8545"
	l2BitAddress    = "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
	l2EthAddress    = "0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111"
	l1BridgeAddress = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788"
	l2BridgeAddress = "0x4200000000000000000000000000000000000010"

	l1BitAddress = "0x59b670e9fA9D0A427751Af201D676719a970857b"

	userPrivateKey = "ddf04c9058d6fac4fea241820f2fbc3b36868d33b80894ba5ff9a9baf8793e10"
	userAddress    = "0xeE3e7d56188ae7af8d5bab980908E3e91c0d7384"

	DECIMAL5    = 5000000000000000000
	DECIMAL1    = 1000000000000000000
	DECIMAL0_5  = 500000000000000000
	DECIMAL0_1  = 100000000000000000
	DECIMAL00_1 = 10000000000000000
)

func TestEnv(t *testing.T) {
	// check l1 bit token
	t.Log("check l1 bit token.....")
	checkTokenAddress(t)

	t.Log("check token bridge.....")
	checkTokenBridge(t)

	t.Log("check balance.....")
	checkBalance(t)
}

func TestDepositAndWithdraw(t *testing.T) {
	t.Log("check balance.....")
	checkBalance(t)

	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// query eth erc20 token
	l1Bridge, err := l1bridge.NewL1StandardBridge(common.HexToAddress(l1BridgeAddress), l1Client)
	require.NoError(t, err)
	l2Bridge, err := l2bridge.NewL2StandardBridge(common.HexToAddress(l2BridgeAddress), l2Client)
	require.NoError(t, err)

	// TEST deposit ETH
	t.Log("----------------")
	t.Log("ETH DEPOSIT TEST")
	t.Log("----------------")
	t.Log("ETH before deposit...\\")
	t.Log("l1 eth balance: ", getETHBalanceFromL1(t, userAddress))
	t.Log("l2 eth balance: ", getETHBalanceFromL2(t, userAddress))
	// do deposit
	auth := buildAuth(t, l1Client, userPrivateKey, big.NewInt(DECIMAL0_1))
	tx, err := l1Bridge.DepositETH(auth, 2_000_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("deposit eth tx hash is: ", tx.Hash())
	t.Log("ETH after deposit...\\")
	t.Log("l1 eth balance: ", getETHBalanceFromL1(t, userAddress))
	//require.Equal(t, getETHBalanceFromL1(t, userAddress), 0)
	// wait for l2 confirmation
	time.Sleep(10 * time.Second)
	t.Log("l2 eth balance: ", getETHBalanceFromL2(t, userAddress))
	//require.Equal(t, getETHBalanceFromL2(t, userAddress), 0)
	t.Log("eth deposit amount: ", DECIMAL0_1)

	// TEST deposit BIT
	t.Log("----------------")
	t.Log("BIT DEPOSIT TEST")
	t.Log("----------------")
	t.Log("BIT before deposit.....\\")
	setL1BitApprove(t)
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))
	auth = buildAuth(t, l1Client, userPrivateKey, big.NewInt(0))
	tx, err = l1Bridge.DepositERC20(auth, common.HexToAddress(l1BitAddress), common.HexToAddress(l2BitAddress), big.NewInt(DECIMAL0_1), 2_000_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("deposit bit tx hash is: ", tx.Hash())
	t.Log("BIT after deposit.....\\")
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	time.Sleep(10 * time.Second)
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))
	t.Log("bit deposit amount: ", DECIMAL0_1)

	// TEST withdraw ETH
	t.Log("-----------------")
	t.Log("ETH WITHDRAW TEST")
	t.Log("-----------------")
	t.Log("ETH before withdraw.....\\")
	setL2EthApprove(t)
	t.Log("l1 eth balance: ", getETHBalanceFromL1(t, userAddress))
	t.Log("l2 eth balance: ", getETHBalanceFromL2(t, userAddress))
	auth = buildAuth(t, l2Client, userPrivateKey, big.NewInt(0))
	tx, err = l2Bridge.Withdraw(auth, common.HexToAddress(l2EthAddress), big.NewInt(DECIMAL0_1), 300_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("withdraw eth tx hash is: ", tx.Hash())
	t.Log("ETH after withdraw.....\\")
	time.Sleep(10 * time.Second)
	t.Log("l1 eth balance: ", getETHBalanceFromL1(t, userAddress))
	t.Log("l2 eth balance: ", getETHBalanceFromL2(t, userAddress))
	t.Log("eth withdraw amount: ", DECIMAL0_1)

	// TEST withdraw BIT
	t.Log("-----------------")
	t.Log("BIT WITHDRAW TEST")
	t.Log("-----------------")
	t.Log("BIT before withdraw.....\\")
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))
	auth = buildAuth(t, l2Client, userPrivateKey, big.NewInt(0))
	tx, err = l2Bridge.Withdraw(auth, common.HexToAddress(l2BitAddress), big.NewInt(DECIMAL0_1), 300_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("withdraw bit tx hash is: ", tx.Hash())
	t.Log("BIT after withdraw.....\\")
	time.Sleep(10 * time.Second)
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))

	t.Log("bit withdraw amount: ", DECIMAL0_1)
}

func TestShowL1L2Balance(t *testing.T) {
	l1Eth := getETHBalanceFromL1(t, userAddress)
	l2Eth := getETHBalanceFromL2(t, userAddress)
	t.Log("l1 eth balance: ", l1Eth)
	t.Log("l2 eth balance: ", l2Eth)
	sumEth := big.NewInt(0)
	t.Log("sum balance is: ", sumEth.Add(l1Eth, l2Eth))

	l1Bit := getBITBalanceFromL1(t, userAddress)
	l2Bit := getBITBalanceFromL2(t, userAddress)
	t.Log("l1 bit balance: ", l1Bit)
	t.Log("l2 bit balance: ", l2Bit)
	sumBit := big.NewInt(0)
	t.Log("sum balance is: ", sumBit.Add(l1Bit, l2Bit))
}

func checkTokenAddress(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// check l1 token address
	code, err := l1Client.CodeAt(context.Background(), common.HexToAddress(l1BitAddress), nil)
	require.NoError(t, err)
	require.True(t, len(code) > 0)
	t.Log("L1 ADDRESS INFO")
	t.Log("L1 Bit Address: ", l1BitAddress)

	// check l2 token address
	code, err = l2Client.CodeAt(context.Background(), common.HexToAddress(l2BitAddress), nil)
	require.NoError(t, err)
	require.True(t, len(code) > 0)
	code, err = l2Client.CodeAt(context.Background(), common.HexToAddress(l2EthAddress), nil)
	require.NoError(t, err)
	require.True(t, len(code) > 0)
	t.Log("L2 ADDRESS INFO")
	t.Log("L2 Bit Address: ", l2BitAddress)
	t.Log("L2 ETH Address: ", l2EthAddress)
}

func checkTokenBridge(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// check l1 token bridge
	code, err := l1Client.CodeAt(context.Background(), common.HexToAddress(l1BridgeAddress), nil)
	require.NoError(t, err)
	require.NotEmpty(t, code)
	t.Log("TOKEN BRIDGE INFO")
	t.Log("find l1 token bridge at: ", l1BridgeAddress)
	// check l2 token bridge
	code, err = l2Client.CodeAt(context.Background(), common.HexToAddress(l2BridgeAddress), nil)
	require.NoError(t, err)
	require.NotEmpty(t, code)
	t.Log("find l2 token bridge at: ", l2BridgeAddress)
}

func checkBalance(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// init balance
	l1Eth := getETHBalanceFromL1(t, userAddress)
	l1Bit := getBITBalanceFromL1(t, userAddress)
	decimal1 := big.NewInt(DECIMAL1)
	if l1Eth.Cmp(decimal1) < 0 {
		delta := big.NewInt(0)
		transferETH(t, l1Client, common.HexToAddress(userAddress), delta.Sub(decimal1, l1Eth).Int64())
		l1Eth = getETHBalanceFromL1(t, userAddress)
	}
	if l1Bit.Cmp(decimal1) < 0 {
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
	require.Equal(t, l1Eth.Int64(), int64(DECIMAL1))
	require.Equal(t, l1Bit.Int64(), int64(DECIMAL1))
	t.Log("balance eth: ", l1Eth)
	t.Log("balance bit: ", l1Bit)
}

func setL1BitApprove(t *testing.T) {
	client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l1BitInstance, err := l1bit.NewBitTokenERC20(common.HexToAddress(l1BitAddress), client)
	require.NoError(t, err)
	auth := buildAuth(t, client, userPrivateKey, big.NewInt(0))
	tx, err := l1BitInstance.Approve(auth, common.HexToAddress(l1BridgeAddress), big.NewInt(DECIMAL5))
	require.NoError(t, err)
	require.NotNil(t, tx)
	l1BitAllowance, err := l1BitInstance.Allowance(&bind.CallOpts{}, common.HexToAddress(userAddress), common.HexToAddress(l1BridgeAddress))
	require.NoError(t, err)
	require.Equal(t, int64(DECIMAL5), l1BitAllowance.Int64())
}

func setL2EthApprove(t *testing.T) {
	client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l2EthInstance, err := l2eth.NewBVMETH(common.HexToAddress(l2EthAddress), client)
	require.NoError(t, err)
	auth := buildAuth(t, client, userPrivateKey, big.NewInt(0))
	tx, err := l2EthInstance.Approve(auth, common.HexToAddress(l2BridgeAddress), big.NewInt(DECIMAL5))
	require.NoError(t, err)
	require.NotNil(t, tx)
	l1BitAllowance, err := l2EthInstance.Allowance(&bind.CallOpts{}, common.HexToAddress(userAddress), common.HexToAddress(l2BridgeAddress))
	require.NoError(t, err)
	require.Equal(t, int64(DECIMAL5), l1BitAllowance.Int64())
}

func getETHBalanceFromL1(t *testing.T, address string) *big.Int {
	client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, client)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	require.NoError(t, err)
	require.NotNil(t, balance)
	return balance
}

func getBITBalanceFromL1(t *testing.T, address string) *big.Int {
	client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l1BitInstance, err := l1bit.NewBitTokenERC20(common.HexToAddress(l1BitAddress), client)
	require.NoError(t, err)
	bal, err := l1BitInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	require.NoError(t, err)
	require.NotNil(t, bal)
	return bal
}

func getETHBalanceFromL2(t *testing.T, address string) *big.Int {
	client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l2EthInstance, err := l2eth.NewBVMETH(common.HexToAddress(l2EthAddress), client)
	require.NoError(t, err)
	balance, err := l2EthInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	require.NoError(t, err)
	require.NotNil(t, balance)
	return balance
}

func getBITBalanceFromL2(t *testing.T, address string) *big.Int {
	client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, client)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	require.NoError(t, err)
	require.NotNil(t, balance)
	return balance
}

func buildAuth(t *testing.T, client *ethclient.Client, privateKey string, amount *big.Int) *bind.TransactOpts {
	privKey, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	require.NoError(t, err)

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//require.NoError(t, err)

	auth := bind.NewKeyedTransactor(privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = amount             // in wei
	auth.GasLimit = uint64(3000000) // in units
	//auth.GasPrice = gasPrice
	auth.GasPrice = big.NewInt(1)
	return auth
}

func transferETH(t *testing.T, client *ethclient.Client, address common.Address, amount int64) {
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

func mintBIT(t *testing.T, client *ethclient.Client, privateKey string, amount int64) {
	l1bitToken, err := l1bit.NewBitTokenERC20(common.HexToAddress(l1BitAddress), client)
	require.NoError(t, err)
	auth := buildAuth(t, client, privateKey, big.NewInt(0))
	tx, err := l1bitToken.Mint(auth, big.NewInt(amount))
	require.NoError(t, err)
	require.NotNil(t, tx)
}

func TestDecimal(t *testing.T) {
	client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l2EthInstance, err := l2eth.NewBVMETH(common.HexToAddress(l2EthAddress), client)
	require.NoError(t, err)

	decimal, err := l2EthInstance.Decimals(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, decimal, uint8(0x12))

	symble, err := l2EthInstance.Symbol(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, symble, "WETH")

	t.Log(decimal)
	t.Log(symble)
}
