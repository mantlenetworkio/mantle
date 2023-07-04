package common

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/btcec"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mantlenetworkio/mantle/l2geth/log"
)

var (
	typByte32Array          abi.Type
	typUint256              abi.Type
	typSlashMsg             abi.Type
	typBytes                abi.Type
	stateBatchArguments     abi.Arguments
	slashMsgArguments       abi.Arguments
	groupPublicKeyArguments abi.Arguments
	slashArguments          abi.Arguments
	rollBackArguments       abi.Arguments
)

type SlashMsg struct {
	BatchIndex *big.Int
	JailNode   common.Address
	TssNodes   []common.Address
	SlashType  *big.Int
}

func init() {
	typByte32Array, _ = abi.NewType("bytes32[]", "bytes32[]", nil)
	typUint256, _ = abi.NewType("uint256", "uint256", nil)
	typBytes, _ = abi.NewType("bytes", "bytes", nil)
	stateBatchArguments = abi.Arguments{
		{
			Type: typByte32Array,
		}, {
			Type: typUint256,
		},
	}

	typSlashMsg, _ = abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "batch_index", Type: "uint256"},
			{Name: "jail_node", Type: "address"},
			{Name: "tss_nodes", Type: "address[]"},
			{Name: "slash_type", Type: "uint256"},
		},
	)

	slashMsgArguments = abi.Arguments{{Type: typSlashMsg}}

	groupPublicKeyArguments = abi.Arguments{
		{
			Type: typBytes,
		}, {
			Type: typBytes,
		},
	}
	slashArguments = abi.Arguments{
		{
			Type: typBytes,
		}, {
			Type: typBytes,
		},
	}
	rollBackArguments = abi.Arguments{{Type: typUint256}}

}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func NodeToAddress(publicKey string) (common.Address, error) {
	if !has0xPrefix(publicKey) {
		publicKey = "0x" + publicKey
	}
	pubKeyBz, err := hexutil.Decode(publicKey)
	if err != nil {
		return common.Address{}, err
	}
	pk, err := btcec.ParsePubKey(pubKeyBz, btcec.S256())
	address := crypto.PubkeyToAddress(*pk.ToECDSA())
	return address, nil
}

func StateBatchHash(stateRoots [][32]byte, offsetStartsAtIndex *big.Int) ([]byte, error) {
	abiEncodedRaw, err := stateBatchArguments.Pack(stateRoots, offsetStartsAtIndex)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(abiEncodedRaw).Bytes(), nil
}

func SlashMsgBytes(batchIndex uint64, jailNode common.Address, tssNodes []common.Address, slashType byte) ([]byte, error) {
	return slashMsgArguments.Pack(SlashMsg{
		SlashType:  new(big.Int).SetUint64(uint64(slashType)),
		JailNode:   jailNode,
		TssNodes:   tssNodes,
		BatchIndex: new(big.Int).SetUint64(batchIndex),
	})
}

func SlashMsgHash(batchIndex uint64, jailNode common.Address, tssNodes []common.Address, slashType byte) ([]byte, error) {
	abiEncodedRaw, err := SlashMsgBytes(batchIndex, jailNode, tssNodes, slashType)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(abiEncodedRaw).Bytes(), nil
}

func RollBackHash(startBlock *big.Int) ([]byte, error) {
	abiEncodeRaw, err := rollBackArguments.Pack(startBlock)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(abiEncodeRaw).Bytes(), nil
}

func SetGroupPubKeyBytes(localKey, poolPubKey []byte) ([]byte, error) {
	return groupPublicKeyArguments.Pack(localKey, poolPubKey)
}

func SlashBytes(msg, sig []byte) ([]byte, error) {
	return slashArguments.Pack(msg, sig)
}

func IsAddrExist(set []common.Address, find common.Address) bool {
	for _, s := range set {
		if bytes.Compare(s.Bytes(), find.Bytes()) == 0 {
			return true
		}
	}
	return false
}

// Ensure we can actually connect l1
func EnsureConnection(client *ethclient.Client) error {
	t := time.NewTicker(1 * time.Second)
	retries := 0
	defer t.Stop()
	for ; true; <-t.C {
		_, err := client.ChainID(context.Background())
		if err == nil {
			break
		} else {
			retries += 1
			if retries > 90 {
				return err
			}
		}
	}
	return nil
}

func EstimateGas(client *ethclient.Client, prikey *ecdsa.PrivateKey, chainId *big.Int, ctx context.Context, tx *types.Transaction, rawContract *bind.BoundContract, to common.Address) (*types.Transaction, error) {
	from := crypto.PubkeyToAddress(prikey.PublicKey)

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error("failed to get header by l1client", "err", err.Error())
		return nil, err
	}
	var gasPrice *big.Int
	var gasTipCap *big.Int
	var gasFeeCap *big.Int
	if header.BaseFee == nil {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			log.Error("cannot fetch gas price ", "err", err.Error())
			return nil, err
		}
	} else {
		gasTipCap, err = client.SuggestGasTipCap(ctx)
		if err != nil {
			log.Warn("failed to SuggestGasTipCap, FallbackGasTipCap = big.NewInt(1500000000) ")
			gasTipCap = big.NewInt(1500000000)
		}
		gasFeeCap = new(big.Int).Add(
			gasTipCap,
			new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
		)
	}

	msg := ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      tx.Data(),
	}

	gasLimit, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Error("failed to EstimateGas", "err", err.Error())
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(prikey, chainId)
	if err != nil {
		log.Error("failed to new ops in estimate gas function", "err", err.Error())
		return nil, err
	}
	opts.Context = ctx
	opts.NoSend = true
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())

	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = 120 * gasLimit / 100 //add 20% buffer to gas limit

	return rawContract.RawTransact(opts, tx.Data())
}
