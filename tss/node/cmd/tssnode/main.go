package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Use:   "tssnode",
		Short: "tss party node",
	}

	cmd.AddCommand(StartCmd())
	cmd.AddCommand(InitConfigTemplate())

	cmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	ctx := context.Background()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
