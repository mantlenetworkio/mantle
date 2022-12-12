package scheduler

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/getsentry/sentry-go"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mantle/scheduler/service"
	"github.com/urfave/cli"
)

// Main is the entrypoint into the scheduler service. This method returns
// a closure that executes the service and blocks until the service exits. The
// use of a closure allows the parameters bound to the top-level main package,
// e.g. GitVersion, to be captured and used once the function is executed.
func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}

		log.Info("Config parsed",
			"min_tx_size", cfg.MinL1TxSize,
			"max_tx_size", cfg.MaxL1TxSize)

		// The call to defer is done here so that any errors logged from
		// this point on are posted to Sentry before exiting.
		if cfg.SentryEnable {
			defer sentry.Flush(2 * time.Second)
		}

		log.Info("Initializing scheduler")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Set up our logging. If Sentry is enabled, we will use our custom log
		// handler that logs to stdout and forwards any error messages to Sentry
		// for collection. Otherwise, logs will only be posted to stdout.
		var logHandler log.Handler
		if cfg.SentryEnable {
			err := sentry.Init(sentry.ClientOptions{
				Dsn:              cfg.SentryDsn,
				Environment:      cfg.EthNetworkName,
				Release:          "scheduler@" + gitVersion,
				TracesSampleRate: bsscore.TraceRateToFloat64(cfg.SentryTraceRate),
				Debug:            false,
			})
			if err != nil {
				return err
			}

			logHandler = bsscore.SentryStreamHandler(os.Stdout, log.JSONFormat())
		} else if cfg.LogTerminal {
			logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
		} else {
			logHandler = log.StreamHandler(os.Stdout, log.JSONFormat())
		}

		logLevel, err := log.LvlFromString(cfg.LogLevel)
		if err != nil {
			return err
		}

		log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

		var services []*service.Service

		scheduler, err := NewScheduler(ctx, cancel, services)
		if err != nil {
			log.Error("Unable to create scheduler", "error", err)
			return err
		}

		log.Info("Starting scheduler")

		if err := scheduler.Start(); err != nil {
			return err
		}
		defer scheduler.Stop()

		log.Info("Scheduler started")

		<-(chan struct{})(nil)

		return nil
	}
}

// Scheduler is a service that configures the necessary resources for
// running the scheduler services.
type Scheduler struct {
	ctx      context.Context
	services []*service.Service
	cancel   func()
}

// NewScheduler initializes the Scheduler, gathering any resources
// that will be needed by the scheduler services.
func NewScheduler(
	ctx context.Context,
	cancel func(),
	services []*service.Service,
) (*Scheduler, error) {
	return &Scheduler{
		ctx:      ctx,
		services: services,
		cancel:   cancel,
	}, nil
}

// Start starts all provided services.
func (s *Scheduler) Start() error {
	for _, service := range s.services {
		if err := service.Start(); err != nil {
			return err
		}
	}
	return nil
}

// Stop stops all provided services and blocks until shutdown.
func (s *Scheduler) Stop() {
	s.cancel()
	for _, service := range s.services {
		_ = service.Stop()
	}
}
