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
	typByte32Array      abi.Type
	typUint256          abi.Type
	typSlashType        abi.Type
	typAddress          abi.Type
	typAddresses        abi.Type
	stateBatchArguments abi.Arguments
	slashMsgArguments   abi.Arguments
)

func init() {
	typByte32Array, _ = abi.NewType("bytes32[]", "bytes32[]", nil)
	typUint256, _ = abi.NewType("uint256", "uint256", nil)
	typSlashType, _ = abi.NewType("uint8", "enumTssStakingSlashing.SlashType", nil)
	typAddress, _ = abi.NewType("address", "address", nil)
	typAddresses, _ = abi.NewType("address[]", "address[]", nil)
	stateBatchArguments = abi.Arguments{
		{
			Type: typByte32Array,
		}, {
			Type: typUint256,
		},
	}

	slashMsgArguments = abi.Arguments{
		{
			Type: typUint256,
		},
		{
			Type: typAddress,
		},
		{
			Type: typAddresses,
		},
		{
			Type: typSlashType,
		},
	}
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
	return slashMsgArguments.Pack(new(big.Int).SetUint64(batchIndex), jailNode, tssNodes, slashType)
}

func SlashMsgHash(batchIndex uint64, jailNode common.Address, tssNodes []common.Address, slashType byte) ([]byte, error) {
	abiEncodedRaw, err := SlashMsgBytes(batchIndex, jailNode, tssNodes, slashType)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(abiEncodedRaw).Bytes(), nil
}
