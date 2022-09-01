package tssnode

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "tssnode",
		Short: "tss party node",
	}

	cmd.AddCommand(StartCmd())

	cmd.PersistentFlags().StringP("config", "c", "config", "configuration file with extension")

	return cmd
	//ctx := context.Background()
	//if err := cmd.ExecuteContext(ctx); err != nil {
	//	os.Exit(1)
	//}
}
