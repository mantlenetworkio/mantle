package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/mantlenetworkio/mantle/mt-challenger"
	"github.com/mantlenetworkio/mantle/mt-challenger/flags"
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
	app.Usage = "MtChallenger EigenDA Challenger Service"
	app.Description = "MtChallenger service for eigen da challenger check eigen da data store right or wrong"
	app.Action = challenger.Main(GitVersion)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("MtChallenger application failed", "message", err)
	}
}
