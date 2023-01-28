package rollup

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mantlenetworkio/mantle/l2geth/p2p/enode"
	"math/big"
	"testing"
	"time"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/abi/bind"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/ethclient"

	"github.com/jarcoal/httpmock"
)

const url = "http://localhost:9999"

func TestRollupClientCannotConnect(t *testing.T) {
	endpoint := fmt.Sprintf("%s/eth/context/latest", url)
	client := NewClient(url, big.NewInt(1))

	httpmock.ActivateNonDefault(client.client.GetClient())

	response, _ := httpmock.NewJsonResponder(
		400,
		map[string]interface{}{},
	)
	httpmock.RegisterResponder(
		"GET",
		endpoint,
		response,
	)

	context, err := client.GetLatestEthContext()
	if context != nil {
		t.Fatal("returned value is not nil")
	}
	if !errors.Is(err, errHTTPError) {
		t.Fatalf("Incorrect error returned: %s", err)
	}
}
func TestDecodedJSON(t *testing.T) {
	str := []byte(`
	{
		"index": 643116,
		"batchIndex": 21083,
		"blockNumber": 25954867,
		"timestamp": 1625605288,
		"gasLimit": "11000000",
		"target": "0x4200000000000000000000000000000000000005",
		"origin": null,
		"data": "0xf86d0283e4e1c08343eab8941a5245ea5210c3b57b7cfdf965990e63534a7b528901a055690d9db800008081aea019f7c6719f1718475f39fb9e5a6a897c3bd5057488a014666e5ad573ec71cf0fa008836030e686f3175dd7beb8350809b47791c23a19092a8c2fab1f0b4211a466",
		"queueOrigin": "sequencer",
		"value": "0x1a055690d9db80000",
		"queueIndex": null,
		"decoded": {
			"nonce": "2",
			"gasPrice": "15000000",
			"gasLimit": "4451000",
			"value": "0x1a055690d9db80000",
			"target": "0x1a5245ea5210c3b57b7cfdf965990e63534a7b52",
			"data": "0x",
			"sig": {
				"v": 1,
				"r": "0x19f7c6719f1718475f39fb9e5a6a897c3bd5057488a014666e5ad573ec71cf0f",
				"s": "0x08836030e686f3175dd7beb8350809b47791c23a19092a8c2fab1f0b4211a466"
			}
		},
		"confirmed": true
	}`)

	tx := new(transaction)
	json.Unmarshal(str, tx)
	cmp, _ := new(big.Int).SetString("1a055690d9db80000", 16)
	if tx.Value.ToInt().Cmp(cmp) != 0 {
		t.Fatal("Cannot decode")
	}
}

type ExtAcc struct {
	Key  *ecdsa.PrivateKey
	Addr common.Address
}

func FromHexKey(hexkey string) (ExtAcc, error) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return ExtAcc{}, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("publicKey is not of type *ecdsa.PublicKey")
		return ExtAcc{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return ExtAcc{key, addr}, nil
}

func TestBatchTransactions(t *testing.T) {
	account, _ := FromHexKey("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	//("3f75eb22760d0d5c50cd320a7289262777ba6e52c0073799407daf19e5905aca")
	txOpt := bind.NewKeyedTransactor(account.Key)

	client, err := ethclient.Dial("http://localhost:7543")
	assert.NoError(t, err)

	receiveAccount := common.HexToAddress("0x76EFAac78C011D24BC975a5dAcC9a91245397e1a")
	txOpt.GasLimit = uint64(21000)
	txOpt.GasPrice = big.NewInt(1)

	for i := 0; i < 10000; i++ {
		nonce, err := client.PendingNonceAt(context.Background(), account.Addr)
		assert.NoError(t, err)

		txOpt.Value = big.NewInt(1).Mul(big.NewInt(3e14), big.NewInt(int64(i)))
		rawTx := types.NewTransaction(nonce, receiveAccount, txOpt.Value, txOpt.GasLimit, txOpt.GasPrice, nil)

		signedTx, err := txOpt.Signer(types.HomesteadSigner{}, txOpt.From, rawTx)
		assert.NoError(t, err)

		err = client.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		t.Logf("index %d, txHash %s", i, signedTx.Hash().String())

		time.Sleep(100 * time.Millisecond)
	}
}

func TestNodeID(t *testing.T) {
	db, _ := enode.OpenDB("")
	account, _ := FromHexKey("0b230e4787022b0dead74c0f16eddc966075109c8945e7c23822bf05ef2459f6")
	ln := enode.NewLocalNode(db, account.Key)
	t.Log(ln.ID())
}
