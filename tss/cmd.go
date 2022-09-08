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
	rootCmd := &cobra.Command{
		Use:   "tss",
		Short: "Tss Daemon",
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

	rootCmd.AddCommand(
		manager.Command(),
		tssnode.Command(),
	)

	rootCmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", &common.Configuration{})
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
