package go_tests

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	l1bit "github.com/mantlenetworkio/mantle/go-test/contracts/L1/local/LocalBitToken.sol"
	l1bridge "github.com/mantlenetworkio/mantle/go-test/contracts/L1/messaging/L1StandardBridge.sol"
	l2bridge "github.com/mantlenetworkio/mantle/go-test/contracts/L2/messaging/L2StandardBridge.sol"
	factory "github.com/mantlenetworkio/mantle/go-test/contracts/L2/messaging/L2StandardTokenFactory.sol"
	l1USDT "github.com/mantlenetworkio/mantle/go-test/contracts/fiat"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestCreateNewCoinPair(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	// deploy USDT at local
	l1Auth := buildAuth(t, l1Client, userPrivateKey, big.NewInt(0))
	addr, tx, USDT, err := l1USDT.DeployUSDT(l1Auth, l1Client, big.NewInt(10000000000), "USDT", "USDT", big.NewInt(6))
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.NotNil(t, USDT)
	require.Equal(t, addr, "USDT_ADDR")

	// check token factory
	tokenFactoryAddress := "0x4200000000000000000000000000000000000012"
	code, err := l2Client.CodeAt(context.Background(), common.HexToAddress(tokenFactoryAddress), nil)
	require.NoError(t, err)
	require.True(t, len(code) > 0)
	tokenFactory, err := factory.NewL2StandardTokenFactory(common.HexToAddress(tokenFactoryAddress), l2Client)
	require.NoError(t, err)

	// create l2 ERC020 binding
	l2Auth := buildAuth(t, l2Client, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", big.NewInt(0))
	tx, err = tokenFactory.CreateStandardL2Token(l2Auth, addr, "Test USDT Token", "TUSDT", 6)
	require.NoError(t, err)
	require.NotNil(t, tx)

	t.Log("tx hash: ", tx.Hash())
}

func TestNewTokenDepositAndWithdraw(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)
	l2Client, err := ethclient.Dial(l2url)
	require.NoError(t, err)
	require.NotNil(t, l2Client)

	t.Log("check balance.....")
	checkBalance(t)

	// query eth erc20 token
	l1Bridge, err := l1bridge.NewL1StandardBridge(common.HexToAddress(l1BridgeAddress), l1Client)
	require.NoError(t, err)
	l2Bridge, err := l2bridge.NewL2StandardBridge(common.HexToAddress(l2BridgeAddress), l2Client)
	require.NoError(t, err)

	// TEST deposit BIT
	t.Log("----------------")
	t.Log("BIT DEPOSIT TEST")
	t.Log("----------------")
	t.Log("BIT before deposit.....\\")
	setL1BitApprove(t)
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))
	auth := buildAuth(t, l1Client, userPrivateKey, big.NewInt(0))
	tx, err := l1Bridge.DepositERC20(auth, common.HexToAddress(l1BitAddress), common.HexToAddress(l2BitAddress), big.NewInt(DECIMAL1), 2_000_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("deposit bit tx hash is: ", tx.Hash())
	t.Log("BIT after deposit.....\\")
	t.Log("l1 bit balance: ", getBITBalanceFromL1(t, userAddress))
	time.Sleep(10 * time.Second)
	t.Log("l2 bit balance: ", getBITBalanceFromL2(t, userAddress))
	t.Log("bit deposit amount: ", DECIMAL1)

	// TEST deposit TT
	t.Log("----------------")
	t.Log("TEST-TOKEN DEPOSIT TEST")
	t.Log("----------------")
	t.Log("TT before deposit.....\\")
	setL1BitApprove(t)
	t.Log("l1 tt balance: ", getBITBalanceFromL1(t, userAddress))
	t.Log("l2 tt balance: ", getBITBalanceFromL2(t, userAddress))
	auth = buildAuth(t, l1Client, userPrivateKey, big.NewInt(0))
	tx, err = l1Bridge.DepositERC20(auth, common.HexToAddress(l1BitAddress), common.HexToAddress(l2BitAddress), big.NewInt(DECIMAL0_1), 2_000_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("deposit tt tx hash is: ", tx.Hash())
	t.Log("TT after deposit.....\\")
	t.Log("l1 tt balance: ", getBITBalanceFromL1(t, userAddress))
	time.Sleep(10 * time.Second)
	t.Log("l2 tt balance: ", getBITBalanceFromL2(t, userAddress))
	t.Log("tt deposit amount: ", DECIMAL0_1)

	// TEST withdraw TT
	t.Log("-----------------")
	t.Log("TT WITHDRAW TEST")
	t.Log("-----------------")
	t.Log("TT before withdraw.....\\")
	//setL2EthApprove(t)
	t.Log("l1 tt balance: ", getETHBalanceFromL1(t, userAddress))
	t.Log("l2 tt balance: ", getETHBalanceFromL2(t, userAddress))
	auth = buildAuth(t, l2Client, userPrivateKey, big.NewInt(0))
	tx, err = l2Bridge.Withdraw(auth, common.HexToAddress(l2EthAddress), big.NewInt(DECIMAL0_1), 300_000, []byte("0x"))
	require.NoError(t, err)
	t.Log("withdraw eth tx hash is: ", tx.Hash())
	t.Log("ETH after withdraw.....\\")
	time.Sleep(10 * time.Second)
	t.Log("l1 eth balance: ", getETHBalanceFromL1(t, userAddress))
	t.Log("l2 eth balance: ", getETHBalanceFromL2(t, userAddress))
	t.Log("eth withdraw amount: ", DECIMAL0_1)

}

func getERC20TokenFromL1(t *testing.T, address string, token string) *big.Int {
	client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, client)

	l1BitInstance, err := l1bit.NewBitTokenERC20(common.HexToAddress(token), client)
	require.NoError(t, err)
	bal, err := l1BitInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	require.NoError(t, err)
	require.NotNil(t, bal)
	return bal
}

func prepareAccount(t *testing.T) {
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
