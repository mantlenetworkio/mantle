package manager

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestSendState(t *testing.T) {
	prik, err := crypto.GenerateKey()
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(prik.PublicKey)
	var batchIndex uint64 = 11
	status := "init"
	sendState.set(address, batchIndex, status)
	actualStatus := sendState.get(address, batchIndex)

	status = "done"
	sendState.set(address, batchIndex, status)
	actualStatus = sendState.get(address, batchIndex)
	require.EqualValues(t, status, actualStatus)
}
