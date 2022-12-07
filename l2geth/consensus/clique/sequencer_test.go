package clique

import (
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
	testCases := []struct {
		seq *Sequencer
		err bool
		msg string
	}{
		{
			seq: NewSequencer(addr, 1),
			err: false,
			msg: "",
		},
		{
			seq: nil,
			err: true,
			msg: "nil sequencer",
		},
		{
			seq: &Sequencer{},
			err: true,
			msg: "sequencer does not have a public key",
		},
		{
			seq: NewSequencer(addr, -1),
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
