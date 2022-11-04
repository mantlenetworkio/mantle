package payer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

func TestPayer(t *testing.T) {
	key, _ := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	UserDir, _ := os.UserHomeDir()
	p := NewPayer(&Config{
		ethereumHttpUrl:           "http://localhost:8545",
		SCCAddress:                common.HexToAddress(""),
		SCCTopic:                  "",
		CTCAddress:                common.HexToAddress(""),
		CTCTopic:                  "",
		privateKey:                key,
		receiverAddr:              common.HexToAddress("0x00000398232E2064F896018496b4b44b3D62751F"),
		l1QueryEpochLengthSeconds: 5,
		waitForReceipt:            true,
		StartBlock:                1,

		HomeDir:  filepath.Join(UserDir, ".subsidy"),
		CacheDir: "payer",
		FileName: "state.txt",
		// Metrics config
		//MetricsEnabled          bool
		//MetricsHTTP             string
		//MetricsPort             int
		//MetricsEnableInfluxDB   bool
		//MetricsInfluxDBEndpoint string
		//MetricsInfluxDBDatabase string
		//MetricsInfluxDBUsername string
		//MetricsInfluxDBPassword string
	})
	//if err := p.PayRollupCost(); err != nil {
	//	panic(err)
	//}
	tx, err := p.Transfer(big.NewInt(1999999999999999999))
	if err != nil {
		panic(err)
	}
	fmt.Println("tx:", tx)
}
