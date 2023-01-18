package clique

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer"
	binding "github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer/bindings"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
)

var (
	amount     = int64(1e18)
	sequencer0 = binding.SequencerSequencerInfo{
		Owner:       ethc.BigToAddress(big.NewInt(1)),
		MintAddress: ethc.BigToAddress(big.NewInt(1)),
		NodeID:      []byte{1},
		Amount:      big.NewInt(amount),
		KeyIndex:    big.NewInt(0),
	}
	sequencer1 = binding.SequencerSequencerInfo{
		Owner:       ethc.BigToAddress(big.NewInt(2)),
		MintAddress: ethc.BigToAddress(big.NewInt(2)),
		NodeID:      []byte{1},
		Amount:      big.NewInt(amount),
		KeyIndex:    big.NewInt(0),
	}
	sequencer2 = binding.SequencerSequencerInfo{
		Owner:       ethc.BigToAddress(big.NewInt(3)),
		MintAddress: ethc.BigToAddress(big.NewInt(3)),
		NodeID:      []byte{1},
		Amount:      big.NewInt(amount),
		KeyIndex:    big.NewInt(0),
	}
)

func TestSchedulerCompareSequencerSet(t *testing.T) {
	testCases := []struct {
		newSeqs synchronizer.SequencerSequencerInfos
		msg     string
		check   func(changes []*Sequencer)
	}{
		{
			newSeqs: synchronizer.SequencerSequencerInfos{
				sequencer0,
				sequencer1,
				sequencer2,
			},
			msg: "Add sequencer",
			check: func(changes []*Sequencer) {
				seqSet := NewSequencerSet(changes)
				// only add sequencer2
				require.Equal(t, seqSet.Len(), 1)
				require.Equal(t, seqSet.Sequencers[0].Address.String(), sequencer2.MintAddress.String())
			},
		},
		{
			newSeqs: synchronizer.SequencerSequencerInfos{},
			msg:     "Delete all sequencer",
			check: func(changes []*Sequencer) {
				require.Equal(t, len(changes), 2)
				for _, v := range changes {
					require.Equal(t, v.Power, int64(0))
				}
			},
		},
		{
			newSeqs: synchronizer.SequencerSequencerInfos{
				sequencer0,
				binding.SequencerSequencerInfo{
					Owner:       sequencer1.Owner,
					MintAddress: sequencer1.MintAddress,
					NodeID:      sequencer1.NodeID,
					Amount:      big.NewInt(0),
					KeyIndex:    sequencer1.KeyIndex,
				},
			},
			msg: "Delete sequencer0",
			check: func(changes []*Sequencer) {
				require.Equal(t, len(changes), 1)
				require.Equal(t, changes[0].Power, int64(0))
			},
		},
		{
			newSeqs: synchronizer.SequencerSequencerInfos{
				sequencer0, sequencer1,
			},
			msg: "Do nothing",
			check: func(changes []*Sequencer) {
			},
		},
	}
	for _, tc := range testCases {
		// init sequencer set
		var seqz []*Sequencer
		seqs := synchronizer.SequencerSequencerInfos{
			sequencer0, sequencer1,
		}
		for _, item := range seqs {
			var addrTemp common.Address
			copy(addrTemp[:], item.MintAddress[:])
			votingPower := big.NewInt(0).Div(item.Amount, scale)
			seqz = append(seqz, NewSequencer(addrTemp, votingPower.Int64(), item.NodeID))
		}
		seqSet := NewSequencerSet(seqz)
		// get changes from test case
		changes := compareSequencerSet(seqSet.Sequencers, tc.newSeqs)
		tc.check(changes)
	}
}

func TestSigAndVerify(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)
	pub := prv.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pub).Hex()
	bps := &types.BatchPeriodStartMsg{
		RollbackStates: nil,
		BatchIndex:     0,
		StartHeight:    1,
		MaxHeight:      100,
		ExpireTime:     100,
		Sequencer:      common.HexToAddress("0xe86c354b11bdc9f295eb2aca01640727dc332d43"),
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
