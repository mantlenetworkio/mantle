package dial

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type ChainSettings struct {
	Provider string
	ChainId  uint64
	Private  string
	Address  common.Address
}

type ChainClient struct {
	ChainSettings
	Client       *ethclient.Client
	NonceManager *NonceManager
}

func MakeChainClient(settings ChainSettings) *ChainClient {
	client, err := ethclient.Dial(settings.Provider)
	if err != nil {
		log.Error("Error. Cannot connect to provider")
	}

	privateKey, err := crypto.HexToECDSA(settings.Private)
	if err != nil {
		log.Error("Invalid key. ")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Error("Cannot get publicKeyECDSA")
	}
	publicAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), publicAddr)
	if err != nil {
		log.Error("Cannot get nonce")
	}

	settings.Address = publicAddr

	nonceManager := &NonceManager{
		client:      client,
		Nonce:       nonce,
		AccountAddr: publicAddr,
	}

	c := &ChainClient{
		ChainSettings: settings,
		Client:        client,
		NonceManager:  nonceManager,
	}

	return c
}

func (c *ChainClient) GetBlockNumber() (uint64, error) {
	return c.Client.BlockNumber(context.Background())
}

func (c *ChainClient) PrepareAuthTransactor() *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(c.Private)
	if err != nil {
		log.Error("Invalid private key")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(c.ChainId))
	if err != nil {
		log.Error("Error getting Transactor")
	}

	nonce, err := c.Client.PendingNonceAt(context.Background(), c.ChainSettings.Address)
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

	//logger.Println("[ChainIO] total price is ", suggestedGasPrice.Uint64() * auth.GasLimit)

	return auth
}

func (c *ChainClient) EnsureTransactionEvaled(tx *types.Transaction) error {
	receipt, err := bind.WaitMined(context.Background(), c.Client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("transaction failed")
	}
	return nil
}

type NonceManager struct {
	client      *ethclient.Client
	Nonce       uint64
	AccountAddr common.Address
}

func (n *NonceManager) TakeNonce() uint64 {
	nonce := n.Nonce
	n.Nonce += 1
	return nonce
}

func (n *NonceManager) SyncProvider() uint64 {
	time.Sleep(500 * time.Millisecond)
	log.Info("Sync Nonce with Provider")
	nonce, err := n.client.PendingNonceAt(context.Background(), n.AccountAddr)
	if err != nil {
		log.Error("Cannot sync nonce from provider")
	}
	return nonce
}
