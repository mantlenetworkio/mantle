package main

import (
	"os"

	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss"
	"github.com/bitdao-io/bitnetwork/tss/manager"
	"github.com/bitdao-io/bitnetwork/tss/node"
	"github.com/spf13/cobra"
)

func main() {
	var config tss.Configuration
	rootCmd := &cobra.Command{
		Use:   "tss",
		Short: "Tss Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, _ := cmd.Flags().GetString("config")
			loadedCfg, err := tss.LoadConfig(cfgFile)
			if err != nil {
				log.Error("fail to load config", err)
				return err
			}
			config = *loadedCfg
			return nil
		},
	}

	rootCmd.AddCommand(
		manager.Command(config),
		node.Command(config),
	)

	rootCmd.Flags().StringP("config", "c", "config", "configuration file with extension")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
