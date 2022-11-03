package scheduler

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSequencerValidateBasic(t *testing.T) {
	var seed []byte
	rand.Read(seed)
	var addr common.Address
	copy(addr[:], seed)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	testCases := []struct {
		seq *Sequencer
		err bool
		msg string
	}{
		{
			seq: NewSequencer(addr, priv.D.Bytes(), 1),
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
				NodeID: nil,
			},
			err: true,
			msg: "sequencer does not have a node id",
		},
		{
			seq: NewSequencer(addr, priv.D.Bytes(), -1),
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

func TestMathScale(t *testing.T) {
	scale := int64(math.Pow10(18))
	require.Equal(t, int64(1000000000000000000), scale)
}
