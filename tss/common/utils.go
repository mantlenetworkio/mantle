package common

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

var (
	typByte32Array          abi.Type
	typUint256              abi.Type
	typSlashMsg             abi.Type
	stateBatchArguments     abi.Arguments
	slashMsgArguments       abi.Arguments
	groupPublicKeyArguments abi.Arguments
	slashArguments          abi.Arguments
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
