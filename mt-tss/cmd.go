package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/manager"
	"github.com/mantlenetworkio/mantle/mt-tss/node/cmd/tssnode"
	"github.com/spf13/cobra"
)

func main() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)

	rootCmd := &cobra.Command{
		Use:   "tss",
		Short: "Tss Start Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, _ := cmd.Flags().GetString("config")
			loadedCfg, err := common.LoadConfig(cfgFile)
			if err == nil {
				return common.SetCmdConfig(cmd, loadedCfg)
			} else {
				return nil
			}
		},
	}

	rootCmd.AddCommand(
		manager.Command(),
		tssnode.Command(),
		tssnode.PeerIDCommand(),
	)

	rootCmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", &common.Configuration{})
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
