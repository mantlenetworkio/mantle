package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	heartbeat "github.com/mantlenetworkio/mantle/mt-heartbeat"
	"github.com/mantlenetworkio/mantle/mt-heartbeat/flags"
	oplog "github.com/mantlenetworkio/mantle/mt-service/log"
	"github.com/urfave/cli"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "mt-heartbeat"
	app.Usage = "Heartbeat recorder"
	app.Description = "Service that records opt-in heartbeats from op nodes"
	app.Action = heartbeat.Main(app.Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
