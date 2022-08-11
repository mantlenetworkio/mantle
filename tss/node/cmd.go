package node

import (
	"github.com/bitdao-io/bitnetwork/tss"
	"github.com/spf13/cobra"
)

func Command(config tss.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "launch a tss  node process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	return cmd
}
