package tss

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
)

type Configuration struct {
}

func DefaultConfiguration() Configuration {
	return Configuration{}
}

func LoadConfig(file string) (*Configuration, error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(file))
	viper.SetConfigName(strings.Split(path.Base(file), ".")[0])
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fail to read from config file: %w", err)
	}
	viper.SetEnvPrefix("tss")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	cfg := DefaultConfiguration()
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fail to unmarshal: %w", err)
	}
	return &cfg, nil
}
