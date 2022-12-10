package clique

import (
	"crypto/ecdsa"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/stretchr/testify/require"
)

func TestSigAndVerify(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)
	pub := prv.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pub).Hex()
	bps := &types.BatchPeriodStartMsg{
		ReorgIndex:  0,
		BatchIndex:  0,
		StartHeight: 1,
		MaxHeight:   100,
		ExpireTime:  1669787879,
		Sequencer:   common.HexToAddress("0xe86c354b11bdc9f295eb2aca01640727dc332d43"),
	}

	bpsHash := bps.GetSignData()
	sig, err := crypto.Sign(crypto.Keccak256(bpsHash), prv)
	require.NoError(t, err)
	bps.Signature = sig
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(bpsHash), bps.GetSignature())
	require.NoError(t, err)
	addressEcr := crypto.PubkeyToAddress(*pubEcr).Hex()
	require.Equal(t, address, addressEcr)

}

func TestSigAndEcrecover(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)
	pub := prv.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pub).Hex()

	testText := []byte("GenerateKey")
	sig, err := crypto.Sign(crypto.Keccak256(testText), prv)
	require.NoError(t, err)
	pubEcr, err := crypto.SigToPub(crypto.Keccak256(testText), sig)
	addressEcr := crypto.PubkeyToAddress(*pubEcr).Hex()

	require.NoError(t, err)
	require.Equal(t, address, addressEcr)
}
