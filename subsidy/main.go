package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics/influxdb"
	"github.com/ethereum/go-ethereum/params"
	ometrics "github.com/mantlenetworkio/mantle/gas-oracle/metrics"
	flags "github.com/mantlenetworkio/mantle/subsidy/flags"
	"github.com/mantlenetworkio/mantle/subsidy/payer"
	"github.com/urfave/cli"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {

	app := cli.NewApp()
	app.Flags = flags.Flags

	app.Version = GitVersion + "-" + params.VersionWithCommit(GitCommit, GitDate)
	app.Name = "subsidize"
	// TODO
	app.Usage = "subsidize Submitter"
	app.Description = "Configure with a private key and an Mantle HTTP endpoint " +
		"to send transactions that transfer Bit to address."

	// Configure the logging
	app.Before = func(ctx *cli.Context) error {
		loglevel := ctx.GlobalUint64(flags.LogLevelFlag.Name)
		log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(loglevel), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
		return nil
	}

	// Define the functionality of the application
	app.Action = func(ctx *cli.Context) error {
		config := payer.NewConfig(ctx)
		if args := ctx.Args(); len(args) > 0 {
			return fmt.Errorf("invalid command: %q", args[0])
		}
		ob := payer.NewPayer(config)
		if err := ob.Start(); err != nil {
			return err
		}
		if config.MetricsEnabled {
			address := fmt.Sprintf("%s:%d", config.MetricsHTTP, config.MetricsPort)
			log.Info("Enabling stand-alone metrics HTTP endpoint", "address", address)
			ometrics.Setup(address)
		}
		if config.MetricsEnableInfluxDB {
			endpoint := config.MetricsInfluxDBEndpoint
			database := config.MetricsInfluxDBDatabase
			username := config.MetricsInfluxDBUsername
			password := config.MetricsInfluxDBPassword
			log.Info("Enabling metrics export to InfluxDB", "endpoint", endpoint, "username", username, "database", database)
			go influxdb.InfluxDBWithTags(ometrics.DefaultRegistry, 10*time.Second, endpoint, database, username, password, "geth.", make(map[string]string))
		}

		ob.Wait()

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("application failed", "message", err)
	}
}
