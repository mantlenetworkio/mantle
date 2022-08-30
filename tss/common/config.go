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
	Manager                 ManagerConfig `json:"manager"`
	L1Url                   string        `json:"l1_url" mapstructure:"l1_url"`
	SccContractAddress      string        `json:"scc_contract_address" mapstructure:"scc_contract_address"`
	TimedTaskInterval       string        `json:"timed_task_interval" mapstructure:"timed_task_interval"`
	L1ReceiptConfirmTimeout string        `json:"l1_receipt_confirm_timeout" mapstructure:"l1_receipt_confirm_timeout"`
	L1ConfirmBlocks         int           `json:"l1_confirm_blocks" mapstructure:"l1_confirm_blocks"`
	SignedBatchesWindow     int           `json:"signed_batches_window" mapstructure:"signedBatchesWindow"`
	MinSignedInWindow       int           `json:"min_signed_in_window" mapstructure:"minSignedInWindow"`
}

type ManagerConfig struct {
	WsAddr   string `json:"ws_addr" mapstructure:"ws_addr"`
	HttpAddr string `json:"http_addr" mapstructure:"http_addr"`
	DBDir    string `json:"db_dir" mapstructure:"db_dir"`

	KeygenTimeout     string `json:"keygen_timeout" mapstructure:"keygen_timeout"`
	CPKConfirmTimeout string `json:"cpk_confirm_timeout" mapstructure:"cpk_confirm_timeout"`
	AskTimeout        string `json:"ask_timeout" mapstructure:"ask_timeout"`
	SignTimeout       string `json:"sign_timeout" mapstructure:"sign_timeout"`
}

func DefaultConfiguration() Configuration {
	return Configuration{
		TimedTaskInterval:       "10s",
		L1ReceiptConfirmTimeout: "20m",
		L1ConfirmBlocks:         10,
		SignedBatchesWindow:     100,
		MinSignedInWindow:       50,
		Manager: ManagerConfig{
			KeygenTimeout:     "120s",
			CPKConfirmTimeout: "2h",
			AskTimeout:        "60s",
			SignTimeout:       "60s",
		},
	}
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
