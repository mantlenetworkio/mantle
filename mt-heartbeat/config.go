package mt-heartbeat

import (
"errors"

"github.com/mantlenetworkio/mantle/mt-heartbeat/flags"
oplog "github.com/mantlenetworkio/mantle/mt-service/log"
mtmetrics "github.com/mantlenetworkio/mantle/mt-service/metrics"
mtpprof "github.com/mantlenetworkio/mantle/mt-service/pprof"
"github.com/urfave/cli"
)
type Config struct {
	HTTPAddr string
	HTTPPort int

	Log oplog.CLIConfig

	Metrics mtmetrics.CLIConfig

	Pprof mtpprof.CLIConfig
}

func (c Config) Check() error {
	if c.HTTPAddr == "" {
		return errors.New("must specify a valid HTTP address")
	}
	if c.HTTPPort <= 0 {
		return errors.New("must specify a valid HTTP port")
	}
	if err := c.Log.Check(); err != nil {
		return err
	}
	if err := c.Metrics.Check(); err != nil {
		return err
	}
	if err := c.Pprof.Check(); err != nil {
		return err
	}
	return nil
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		HTTPAddr: ctx.GlobalString(flags.HTTPAddrFlag.Name),
		HTTPPort: ctx.GlobalInt(flags.HTTPPortFlag.Name),
		Log:      oplog.ReadCLIConfig(ctx),
		Metrics:  mtmetrics.ReadCLIConfig(ctx),
		Pprof:    mtpprof.ReadCLIConfig(ctx),
	}
}
