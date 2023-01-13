package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	mt_meso_service "github.com/mantlenetworkio/mantle/mt-meso-service"
	"github.com/mantlenetworkio/mantle/mt-meso-service/flags"
	"github.com/urfave/cli"
	"os"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", GitVersion, params.VersionWithCommit(GitCommit, GitDate))
	app.Name = "mt-challenger"
	app.Usage = "Mt Middleware Service"
	app.Description = "Mt Middleware Service is a server for l1 and l2 communication"
	app.Action = mt_meso_service.Main(GitVersion)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Mt Middleware Service application failed", "message", err)
	}
}
