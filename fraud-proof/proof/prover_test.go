package proof

import (
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestExecutionStateHash(t *testing.T) {
	gasUsed, ok := new(big.Int).SetString("1eaed", 16)
	require.True(t, ok)
	es := ExecutionState{
		VMHash:            common.HexToHash("0x73ea1d2ee379cd800d4006953888d2e23a57675dfc44de09e1d0194fcd866b1f"),
		CumulativeGasUsed: gasUsed,
		BlockGasUsed:      gasUsed,
	}
	t.Log(es.Hash().String())
}
