package mt_batcher

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/getsentry/sentry-go"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/mt-batcher/l1l2client"
	"github.com/mantlenetworkio/mantle/mt-batcher/sequencer"
	"github.com/urfave/cli"
	"time"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}
		log.Info("Config parsed",
			"disperser", cfg.DisperserEndpoint,
			"mtlrpc", cfg.L2MtlRpc)

		if cfg.SentryEnable {
			defer sentry.Flush(2 * time.Second)
		}
		log.Info("Initializing batch submitter")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var logHandler log.Handler
		logLevel, err := log.LvlFromString(cfg.LogLevel)
		if err != nil {
			return err
		}
		log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

		sequencerPrivKey, eigenAddress, err := bsscore.ParseWalletPrivKeyAndContractAddr(
			"Sequencer", cfg.Mnemonic, cfg.SequencerHDPath,
			cfg.PrivateKey, cfg.RollupAddress,
		)
		if err != nil {
			return err
		}

		l1l2Config := &l1l2client.Config{
			L1RpcUrl:     cfg.L1EthRpc,
			ChainId:      cfg.ChainId,
			Private:      cfg.PrivateKey,
			DisableHTTP2: cfg.DisableHTTP2,
		}

		l1Client, err := l1l2client.NewL1ChainClient(ctx, l1l2Config)
		if err != nil {
			return err
		}
		chainID, err := l1Client.Client.ChainID(ctx)
		if err != nil {
			return err
		}
		l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		driverConfig := &sequencer.DriverConfig{
			L1Client:          l1Client,
			L2Client:          l2Client,
			EigenAddr:         eigenAddress,
			PrivKey:           sequencerPrivKey,
			BlockOffset:       cfg.BlockOffset,
			ChainID:           chainID,
			DataStoreDuration: uint64(cfg.DataStoreDuration),
			DataStoreTimeout:  cfg.DataStoreTimeout,
			DisperserSocket:   cfg.DisperserEndpoint,
			PollInterval:      cfg.PollInterval,
		}
		driver, err := sequencer.NewDriver(ctx, driverConfig)
		if err := driver.Start(); err != nil {
			return err
		}
		defer driver.Stop()
		log.Info("mt batcher started")
		return nil
	}
}
