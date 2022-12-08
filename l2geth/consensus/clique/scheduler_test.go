package clique

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
)

func TestProducerSerializeB(t *testing.T) {
	//address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	bps := &types.BatchPeriodStartMsg{
		ReorgIndex:   0,
		BatchIndex:   0,
		StartHeight:  1,
		MaxHeight:    100,
		ExpireTime:   1669787879,
		MinerAddress: common.HexToAddress("0xe86c354b11bdc9f295eb2aca01640727dc332d43"),
		SequencerSet: []common.Address{
			common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
			common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8"),
			common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc"),
			common.HexToAddress("0x90f79bf6eb2c4f870365e785982e1f101e93b906"),
			common.HexToAddress("0x15d34aaf54267db7d7c367839aaf71a00a2c6a65"),
			common.HexToAddress("0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"),
			common.HexToAddress("0x976ea74026e726554db657fa54763abd0c3a0aa9"),
		},
		Signature: common.Hex2Bytes("2020a0bbf67b08b1df594333c1ead3a771d9742d2f33798e050da744b1255bb67860d672a5055429cc53d17e6c57550989b39cf997e2fb58d1ec6aae198a471501"),
	}
	buf := bps.SerializeBatchPeriodStartMsg()
	t.Log(hex.EncodeToString(buf))
	str := make([]byte, extraVanity)
	end := make([]byte, extraSeal)
	buf = append(str, buf...)
	buf = append(buf, end...)
	t.Log(hex.EncodeToString(buf))
}

func TestSigAndVerify(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)
	pub := prv.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pub).Hex()
	bps := &types.BatchPeriodStartMsg{
		ReorgIndex:   0,
		BatchIndex:   0,
		StartHeight:  1,
		MaxHeight:    100,
		ExpireTime:   1669787879,
		MinerAddress: common.HexToAddress("0xe86c354b11bdc9f295eb2aca01640727dc332d43"),
		SequencerSet: []common.Address{
			common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
			common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8"),
			common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc"),
			common.HexToAddress("0x90f79bf6eb2c4f870365e785982e1f101e93b906"),
			common.HexToAddress("0x15d34aaf54267db7d7c367839aaf71a00a2c6a65"),
			common.HexToAddress("0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"),
			common.HexToAddress("0x976ea74026e726554db657fa54763abd0c3a0aa9"),
		},
	}

	bpsHash := bps.GetData()
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
