package payer

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/subsidy/flags"
	"github.com/urfave/cli"
)

// Config represents the configuration options for the gas oracle
type Config struct {
	queryerHttpUrl            string
	payerHttpUrl              string
	l2HttpUrl                 string
	SCCAddress                common.Address
	SCCTopic                  string
	CTCAddress                common.Address
	CTCTopic                  string
	gpoAddress                common.Address
	privateKey                *ecdsa.PrivateKey
	receiverAddr              common.Address
	l1QueryEpochLengthSeconds uint64
	waitForReceipt            bool
	HomeDir                   string
	CacheDir                  string
	FileName                  string
	RevisedBlock              uint64
	StartBlock                uint64
	//receivable                common.Address
	// Metrics config
	MetricsEnabled          bool
	MetricsHTTP             string
	MetricsPort             int
	MetricsEnableInfluxDB   bool
	MetricsInfluxDBEndpoint string
	MetricsInfluxDBDatabase string
	MetricsInfluxDBUsername string
	MetricsInfluxDBPassword string
}

// NewConfig creates a new Config
func NewConfig(ctx *cli.Context) *Config {
	cfg := Config{}
	cfg.queryerHttpUrl = ctx.GlobalString(flags.QueryerHttpUrlFlag.Name)
	cfg.payerHttpUrl = ctx.GlobalString(flags.PayerHttpUrlFlag.Name)
	sccAddr := ctx.GlobalString(flags.SCCAddressFlag.Name)
	cfg.SCCAddress = common.HexToAddress(sccAddr)
	cfg.SCCTopic = ctx.GlobalString(flags.SCCTopicFlag.Name)
	ctcAddr := ctx.GlobalString(flags.CTCAddressFlag.Name)
	cfg.CTCAddress = common.HexToAddress(ctcAddr)
	cfg.CTCTopic = ctx.GlobalString(flags.CTCTopicFlag.Name)
	cfg.l2HttpUrl = ctx.GlobalString(flags.L2HttpUrlFlag.Name)
	gpoAddr := ctx.GlobalString(flags.GPOAddressFlag.Name)
	cfg.gpoAddress = common.HexToAddress(gpoAddr)
	cfg.l1QueryEpochLengthSeconds = ctx.GlobalUint64(flags.L1QueryEpochLengthSecondsFlag.Name)
	cfg.HomeDir = ctx.GlobalString(flags.HomeDirFlag.Name)
	cfg.CacheDir = ctx.GlobalString(flags.CacheDirFlag.Name)
	cfg.FileName = ctx.GlobalString(flags.FileNameFlag.Name)
	cfg.StartBlock = ctx.GlobalUint64(flags.StartBlockFlag.Name)
	cfg.RevisedBlock = ctx.GlobalUint64(flags.RevisedBlockFlag.Name)

	if ctx.GlobalIsSet(flags.PrivateKeyFlag.Name) {
		hex := ctx.GlobalString(flags.PrivateKeyFlag.Name)
		hex = strings.TrimPrefix(hex, "0x")
		key, err := crypto.HexToECDSA(hex)
		if err != nil {
			log.Error(fmt.Sprintf("Option %q: %v", flags.PrivateKeyFlag.Name, err))
		}
		cfg.privateKey = key
	} else {
		log.Crit("No private key configured")
	}

	receiveHex := ctx.GlobalString(flags.ReceiveAddressFlag.Name)
	cfg.receiverAddr = common.HexToAddress(receiveHex)

	if ctx.GlobalIsSet(flags.WaitForReceiptFlag.Name) {
		cfg.waitForReceipt = true
	}

	cfg.MetricsEnabled = ctx.GlobalBool(flags.MetricsEnabledFlag.Name)
	cfg.MetricsHTTP = ctx.GlobalString(flags.MetricsHTTPFlag.Name)
	cfg.MetricsPort = ctx.GlobalInt(flags.MetricsPortFlag.Name)
	cfg.MetricsEnableInfluxDB = ctx.GlobalBool(flags.MetricsEnableInfluxDBFlag.Name)
	cfg.MetricsInfluxDBEndpoint = ctx.GlobalString(flags.MetricsInfluxDBEndpointFlag.Name)
	cfg.MetricsInfluxDBDatabase = ctx.GlobalString(flags.MetricsInfluxDBDatabaseFlag.Name)
	cfg.MetricsInfluxDBUsername = ctx.GlobalString(flags.MetricsInfluxDBUsernameFlag.Name)
	cfg.MetricsInfluxDBPassword = ctx.GlobalString(flags.MetricsInfluxDBPasswordFlag.Name)

	return &cfg
}
