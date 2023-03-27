package common

import (
	"bytes"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	typByte32               abi.Type
	typUint256              abi.Type
	typSlashMsg             abi.Type
	typBytes                abi.Type
	outputRootArguments     abi.Arguments
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
	typByte32, _ = abi.NewType("bytes32", "bytes32", nil)
	typUint256, _ = abi.NewType("uint256", "uint256", nil)
	typBytes, _ = abi.NewType("bytes", "bytes", nil)
	outputRootArguments = abi.Arguments{
		{
			Type: typByte32,
		}, {
			Type: typUint256,
		}, {
			Type: typByte32,
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

func StateBatchHash(output [32]byte, l2_block_number *big.Int, l1_block_hash [32]byte, l1_block_number *big.Int) ([]byte, error) {
	abiEncodedRaw, err := outputRootArguments.Pack(output, l2_block_number, l1_block_hash, l1_block_number)
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
