package clique

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"

	"github.com/stretchr/testify/assert"
)

func TestSequencerValidateBasic(t *testing.T) {
	var seed []byte
	rand.Read(seed)
	var addr common.Address
	copy(addr[:], seed)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKey := priv.PublicKey
	testCases := []struct {
		seq *Sequencer
		err bool
		msg string
	}{
		{
			seq: NewSequencer(addr, pubKey, 1),
			err: false,
			msg: "",
		},
		{
			seq: nil,
			err: true,
			msg: "nil sequencer",
		},
		{
			seq: &Sequencer{
				PubKey: ecdsa.PublicKey{
					Curve: nil,
					X:     nil,
					Y:     nil,
				},
			},
			err: true,
			msg: "sequencer does not have a public key",
		},
		{
			seq: NewSequencer(addr, pubKey, -1),
			err: true,
			msg: "sequencer has negative voting power",
		},
	}

	for _, tc := range testCases {
		err := tc.seq.SequencerBasic()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}

//------------------//

func TestCommonAddressValidity(t *testing.T) {
	var addr common.Address
	fmt.Printf("addr: %s", addr)

	var newaddress common.Address
	address := randAddress()
	rand.Read(newaddress.Bytes())
	println(address.String())
}
