package tssnode

import (
	"fmt"
	"github.com/bitdao-io/bitnetwork/tss/node/config"
	"github.com/spf13/cobra"
)

func InitConfigTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-config",
		Short: "initialize config template file",
		RunE: func(cmd *cobra.Command, args []string) error {
			output, _ := cmd.Flags().GetString("output")
			cfg := config.DefaultConfiguration()

			baseDir, _ := cmd.Flags().GetString("base-dir")
			if len(baseDir) > 0 {
				cfg.BaseDir = baseDir
				cfg.Key.KeyringDir = baseDir
			}
			config.WriteStartConfigFile(output, cfg)
			fmt.Println(fmt.Sprintf("configuration is generated in %s, please check", output))
			return nil
		},
	}
	cmd.Flags().StringP("output", "o", "config.toml", "output directory of the initialized configuration template file")
	cmd.Flags().String("base-dir", "", "specified base directory, default '~/.tssnode'")

	return cmd
}
