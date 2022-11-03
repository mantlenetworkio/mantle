package clique

import (
	"encoding/hex"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/stretchr/testify/require"
)

func TestProducerSerialize(t *testing.T) {
	address := common.HexToAddress("5e7f8869877D473821862A2b1075e129946eB8b4")
	Producer := Sequencer{
		Address:          address,
		Power:            100,
		ProducerPriority: -100,
	}
	producers := Producers{
		Epoch:       100,
		SchedulerID: address.Bytes(),
		SequencerSet: SequencerSet{
			Sequencers: []*Sequencer{
				&Producer,
			},
			Producer:   &Producer,
			totalPower: 100,
		},
	}

	buf := producers.serialize()
	println(hex.EncodeToString(buf))
	producers2 := deserialize(buf)
	require.Equal(t, producers, *producers2)
}
