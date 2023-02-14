package common

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"

	"github.com/decred/dcrd/hdkeychain/v3"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/tyler-smith/go-bip39"
)

var (
	ErrCannotGetPrivateKey = errors.New("invalid combination of privkey or mnemonic+hdpath")
)

func ParseAddress(address string) (common.Address, error) {
	if common.IsHexAddress(address) {
		return common.HexToAddress(address), nil
	}
	return common.Address{}, fmt.Errorf("invalid address: %v", address)
}

func GetConfiguredPrivateKey(mnemonic, hdPath, privKeyStr string) (*ecdsa.PrivateKey, error) {

	useMnemonic := mnemonic != "" && hdPath != ""
	usePrivKeyStr := privKeyStr != ""

	switch {
	case useMnemonic && !usePrivKeyStr:
		return DerivePrivateKey(mnemonic, hdPath)

	case usePrivKeyStr && !useMnemonic:
		return ParsePrivateKeyStr(privKeyStr)

	default:
		return nil, ErrCannotGetPrivateKey
	}
}

type fakeNetworkParams struct{}

func (f fakeNetworkParams) HDPrivKeyVersion() [4]byte {
	return [4]byte{}
}

func (f fakeNetworkParams) HDPubKeyVersion() [4]byte {
	return [4]byte{}
}

func DerivePrivateKey(mnemonic, hdPath string) (*ecdsa.PrivateKey, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	privKey, err := hdkeychain.NewMaster(seed, fakeNetworkParams{})
	if err != nil {
		return nil, err
	}

	derivationPath, err := accounts.ParseDerivationPath(hdPath)
	if err != nil {
		return nil, err
	}

	for _, child := range derivationPath {
		privKey, err = privKey.Child(child)
		if err != nil {
			return nil, err
		}
	}

	rawPrivKey, err := privKey.SerializedPrivKey()
	if err != nil {
		return nil, err
	}

	return crypto.ToECDSA(rawPrivKey)
}

func ParsePrivateKeyStr(privKeyStr string) (*ecdsa.PrivateKey, error) {
	hex := strings.TrimPrefix(privKeyStr, "0x")
	return crypto.HexToECDSA(hex)
}

func ParseWalletPrivKeyAndContractAddr(name string, mnemonic string, hdPath string, privKeyStr string, contractAddrStr string) (*ecdsa.PrivateKey, common.Address, error) {

	privKey, err := GetConfiguredPrivateKey(mnemonic, hdPath, privKeyStr)
	if err != nil {
		return nil, common.Address{}, err
	}

	contractAddress, err := ParseAddress(contractAddrStr)
	if err != nil {
		return nil, common.Address{}, err
	}

	walletAddress := crypto.PubkeyToAddress(privKey.PublicKey)

	log.Info(name+" wallet params parsed successfully", "wallet_address",
		walletAddress, "contract_address", contractAddress)

	return privKey, contractAddress, nil
}

func PrivateKeySignerFn(key *ecdsa.PrivateKey, chainID *big.Int) bind.SignerFn {
	from := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.LatestSignerForChainID(chainID)
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
}
