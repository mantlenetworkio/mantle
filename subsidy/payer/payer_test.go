package payer

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestPayer(t *testing.T) {
	key, _ := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	UserDir, _ := os.UserHomeDir()
	ob := NewPayer(&Config{
		ethereumHttpUrl:           "https://rpc.ankr.com/eth_goerli",
		l2gethHttpUrl:             "http://127.0.0.1:8545",
		SCCAddress:                common.HexToAddress("0x56Fab8B6bceB262fC6E17cA142d1b3e611aE076F"),
		SCCTopic:                  "StateBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)",
		CTCAddress:                common.HexToAddress("0x2E816dC5A21868f160bDad407a740a580245251C"),
		CTCTopic:                  "SequencerBatchAppended(uint256,uint256,uint256)",
		gpoAddress:                common.HexToAddress("0x420000000000000000000000000000000000000F"),
		privateKey:                key,
		receiverAddr:              common.HexToAddress("0x00000398232E2064F896018496b4b44b3D62751F"),
		l1QueryEpochLengthSeconds: 5,
		waitForReceipt:            true,
		StartBlock:                7933454,
		HomeDir:                   filepath.Join(UserDir, ".subsidy"),
		CacheDir:                  "payer",
		FileName:                  "state.txt",
	})
	// testing transfer
	txHash, err := ob.Transfer(big.NewInt(1999999999999999999))
	require.NoError(t, err)
	require.True(t, txHash != "")
	// testing
	var fromBlock uint64 = 7933268
	var toBlock uint64 = 7933459
	totalFee := big.NewInt(0)
	sccLogs, err := ob.getLogs(ob.sccAddrStr, ob.sccTopic, fromBlock, toBlock)
	require.NoError(t, err)
	require.True(t, len(sccLogs) > 0)
	ctcLogs, err := ob.getLogs(ob.ctcAddrStr, ob.ctcTopic, fromBlock, toBlock)
	require.NoError(t, err)
	require.True(t, len(sccLogs) > 0)
	totalFee = totalFee.Add(ob.CalculateCost(sccLogs), ob.CalculateCost(ctcLogs))
	require.True(t, totalFee.Cmp(big.NewInt(0)) == 1)
}
