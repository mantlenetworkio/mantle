package l1l2client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/bss-core/dial"
	"math/big"
)

type Config struct {
	L1Client     string
	ChainId      uint64
	Private      string
	Address      common.Address
	DisableHTTP2 bool
}

type L1ChainClient struct {
	conf   *Config
	Client *ethclient.Client
}

func NewL1ChainClient(ctx context.Context, conf *Config) *L1ChainClient {
	client, err := dial.L1EthClientWithTimeout(ctx, conf.L1Client, conf.DisableHTTP2)
	if err != nil {
		log.Error("Error. Cannot connect to provider")
	}

	privateKey, err := crypto.HexToECDSA(conf.Private)
	if err != nil {
		log.Error("Invalid key. ")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Cannot get publicKeyECDSA")
	}
	publicAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	if err != nil {
		log.Error("Cannot get nonce")
	}
	conf.Address = publicAddr
	c := &L1ChainClient{
		conf:   conf,
		Client: client,
	}
	return c
}

func (c *L1ChainClient) GetBalance(ctx context.Context) (*big.Int, error) {
	return c.Client.BalanceAt(ctx, c.conf.Address, nil)
}

func (c *L1ChainClient) GetBlockNumber() (uint64, error) {
	return c.Client.BlockNumber(context.Background())
}

func (c *L1ChainClient) PrepareAuthTransactor() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(c.conf.Private)
	if err != nil {
		log.Error("Invalid private key")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(c.conf.ChainId))
	if err != nil {
		log.Error("Error getting Transactor")
	}
	nonce, err := c.Client.PendingNonceAt(context.Background(), c.conf.Address)
	if err != nil {
		log.Error("Cannot sync nonce from provider")
	}
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(10000000)
	suggestedGasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Cannot get transaction suggested price")
		suggestedGasPrice, _ = new(big.Int).SetString("1000000000", 10) // 1Gwei
	}
	auth.GasPrice = suggestedGasPrice
	return auth
}

func (c *L1ChainClient) EnsureTransactionEvaled(tx *types.Transaction) error {
	receipt, err := bind.WaitMined(context.Background(), c.Client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("transaction failed")
	}
	return nil
}
