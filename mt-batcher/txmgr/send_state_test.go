package txmgr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/mantlenetworkio/mantle/mt-batcher/txmgr"
)

const testSafeAbortNonceTooLowCount = 3

var (
	testHash = common.HexToHash("0x01")
)

func newSendState() *txmgr.SendState {
	return txmgr.NewSendState(testSafeAbortNonceTooLowCount)
}

func processNSendErrors(sendState *txmgr.SendState, err error, n int) {
	for i := 0; i < n; i++ {
		sendState.ProcessSendError(err)
	}
}

func TestSendStateNoAbortAfterInit(t *testing.T) {
	sendState := newSendState()
	require.False(t, sendState.ShouldAbortImmediately())
	require.False(t, sendState.IsWaitingForConfirmation())
}

func TestSendStateNoAbortAfterProcessNilError(t *testing.T) {
	sendState := newSendState()

	processNSendErrors(sendState, nil, testSafeAbortNonceTooLowCount)
	require.False(t, sendState.ShouldAbortImmediately())
}

func TestSendStateNoAbortAfterProcessOtherError(t *testing.T) {
	sendState := newSendState()

	otherError := errors.New("other error")
	processNSendErrors(sendState, otherError, testSafeAbortNonceTooLowCount)
	require.False(t, sendState.ShouldAbortImmediately())
}

func TestSendStateAbortSafelyAfterNonceTooLowButNoTxMined(t *testing.T) {
	sendState := newSendState()

	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.False(t, sendState.ShouldAbortImmediately())
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.False(t, sendState.ShouldAbortImmediately())
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.True(t, sendState.ShouldAbortImmediately())
}

func TestSendStateMiningTxCancelsAbort(t *testing.T) {
	sendState := newSendState()

	sendState.ProcessSendError(core.ErrNonceTooLow)
	sendState.ProcessSendError(core.ErrNonceTooLow)
	sendState.TxMined(testHash)
	require.False(t, sendState.ShouldAbortImmediately())
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.False(t, sendState.ShouldAbortImmediately())
}

func TestSendStateReorgingTxResetsAbort(t *testing.T) {
	sendState := newSendState()

	sendState.ProcessSendError(core.ErrNonceTooLow)
	sendState.ProcessSendError(core.ErrNonceTooLow)
	sendState.TxMined(testHash)
	sendState.TxNotMined(testHash)
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.False(t, sendState.ShouldAbortImmediately())
}

func TestSendStateNoAbortEvenIfNonceTooLowAfterTxMined(t *testing.T) {
	sendState := newSendState()

	sendState.TxMined(testHash)
	processNSendErrors(
		sendState, core.ErrNonceTooLow, testSafeAbortNonceTooLowCount,
	)
	require.False(t, sendState.ShouldAbortImmediately())
}

func TestSendStateSafeAbortIfNonceTooLowPersistsAfterUnmine(t *testing.T) {
	sendState := newSendState()

	sendState.TxMined(testHash)
	sendState.TxNotMined(testHash)
	sendState.ProcessSendError(core.ErrNonceTooLow)
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.False(t, sendState.ShouldAbortImmediately())
	sendState.ProcessSendError(core.ErrNonceTooLow)
	require.True(t, sendState.ShouldAbortImmediately())
}

func TestSendStateSafeAbortWhileCallingNotMinedOnUnminedTx(t *testing.T) {
	sendState := newSendState()

	processNSendErrors(
		sendState, core.ErrNonceTooLow, testSafeAbortNonceTooLowCount,
	)
	sendState.TxNotMined(testHash)
	require.True(t, sendState.ShouldAbortImmediately())
}

func TestSendStateIsWaitingForConfirmationAfterTxMined(t *testing.T) {
	sendState := newSendState()

	testHash2 := common.HexToHash("0x02")

	sendState.TxMined(testHash)
	require.True(t, sendState.IsWaitingForConfirmation())
	sendState.TxMined(testHash2)
	require.True(t, sendState.IsWaitingForConfirmation())
}

func TestSendStateIsNotWaitingForConfirmationAfterTxUnmined(t *testing.T) {
	sendState := newSendState()

	sendState.TxMined(testHash)
	sendState.TxNotMined(testHash)
	require.False(t, sendState.IsWaitingForConfirmation())
}
