package scheduler

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/stretchr/testify/require"
)

func TestConvertAddr(t *testing.T) {
	pubKey, err := hex.DecodeString("04e68acfc0253a10620dff706b0a1b1f1f5833ea3beb3bde2250d5f271f3563606672ebc45e0b7ea2e816ecb70ca03137b1c9476eec63d4632e990020b7b6fba39")
	require.NoError(t, err)

	ecdsaPubKey, err := crypto.UnmarshalPubkey(pubKey)
	require.NoError(t, err)

	fmt.Println(crypto.PubkeyToAddress(*ecdsaPubKey).String())
}
