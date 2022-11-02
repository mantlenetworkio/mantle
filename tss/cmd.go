package main

import (
	"context"
	"os"

	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager"
	"github.com/mantlenetworkio/mantle/tss/node/cmd/tssnode"
	"github.com/spf13/cobra"
)

func main() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)

	cmd := &cobra.Command{
		Use:   "tss",
		Short: "Tss Command",
	}

	subCmd := &cobra.Command{
		Use:   "start",
		Short: "Tss Start Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, _ := cmd.Flags().GetString("config")
			loadedCfg, err := common.LoadConfig(cfgFile)
			if err != nil {
				log.Error("fail to load config", err)
				return err
			}

			return common.SetCmdConfig(cmd, loadedCfg)
		},
	}

	subCmd.AddCommand(
		manager.Command(),
		tssnode.Command(),
	)

	cmd.AddCommand(
		subCmd,
		tssnode.PeerIDCommand(),
	)

	subCmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", &common.Configuration{})
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
