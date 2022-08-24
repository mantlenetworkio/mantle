package common

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
)

type Configuration struct {
	Name string `json:"name"`
}

func DefaultConfiguration() Configuration {
	return Configuration{}
}

func LoadConfig(file string) (Configuration, error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(file))
	viper.SetConfigName(strings.Split(path.Base(file), ".")[0])
	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, fmt.Errorf("fail to read from config file: %w", err)
	}
	viper.SetEnvPrefix("tss")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	cfg := DefaultConfiguration()
	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, fmt.Errorf("fail to unmarshal: %w", err)
	}
	return cfg, nil
}

func GetConfigFromCmd(cmd *cobra.Command) Configuration {
	if v := cmd.Context().Value("config"); v != nil {
		clientCtxPtr := v.(*Configuration)
		return *clientCtxPtr
	}
	return Configuration{}
}

func SetCmdConfig(cmd *cobra.Command, config Configuration) error {
	v := cmd.Context().Value("config")
	if v == nil {
		return errors.New("client context not set")
	}

	configPtr := v.(*Configuration)
	*configPtr = config

	return nil
}
