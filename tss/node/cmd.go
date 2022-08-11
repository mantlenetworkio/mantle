package node

import (
	"fmt"

	"github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "launch a tss node process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			config := types.GetConfigFromCmd(cmd)
			fmt.Println(config)
			return nil
		},
	}
	return cmd
}
