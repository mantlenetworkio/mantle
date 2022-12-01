package mt_batcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/getsentry/sentry-go"
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
		fmt.Println(ctx)
		return nil
	}
}
