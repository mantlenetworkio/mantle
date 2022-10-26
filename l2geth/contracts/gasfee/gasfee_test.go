package gasfee

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPacketData(t *testing.T) {
	data, err := PacketData()
	require.NoError(t, err)
	require.NotNil(t, data)

}
