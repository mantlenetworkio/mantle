package main

import (
	"context"
	"github.com/bitdao-io/bitnetwork/tss/node/cmd/tssnode"
	"os"

	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss/manager"
	"github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "tss",
		Short: "Tss Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, _ := cmd.Flags().GetString("config")
			loadedCfg, err := types.LoadConfig(cfgFile)
			if err != nil {
				log.Error("fail to load config", err)
				return err
			}

			return types.SetCmdConfig(cmd, loadedCfg)
		},
	}

	rootCmd.AddCommand(
		manager.Command(),
		tssnode.Command(),
	)

	rootCmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", &types.Configuration{})
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
