package common

import (
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/common/hexutil"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/btcsuite/btcd/btcec"
	"math/big"
)

func NodeToAddress(publicKey string) (common.Address, error) {
	pubKeyBz, err := hexutil.Decode(publicKey)
	if err != nil {
		return common.Address{}, err
	}
	pk, err := btcec.ParsePubKey(pubKeyBz, btcec.S256())
	address := crypto.PubkeyToAddress(*pk.ToECDSA())
	return address, nil
}

func StateBatchDigestBytes(stateRoots [][32]byte, offsetStartsAtIndex *big.Int) []byte {
	rawBytes := make([]byte, 0)
	for _, sr := range stateRoots {
		rawBytes = append(rawBytes, sr[:]...)
	}
	rawBytes = append(rawBytes, offsetStartsAtIndex.Bytes()...)
	return crypto.Keccak256Hash(rawBytes).Bytes()
}
