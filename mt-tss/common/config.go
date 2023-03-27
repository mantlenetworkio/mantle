package common

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Configuration struct {
	Manager                        ManagerConfig `json:"manager"`
	Node                           NodeConfig    `json:"node"`
	L1Url                          string        `json:"l1_url" mapstructure:"l1_url"`
	SccContractAddress             string        `json:"scc_contract_address" mapstructure:"scc_contract_address"`
	TssGroupContractAddress        string        `json:"tss_group_contract_address" mapstructure:"tss_group_contract_address"`
	TssStakingSlashContractAddress string        `json:"tss_staking_slash_contract_address" mapstructure:"tss_staking_slash_contract_address"`
	TimedTaskInterval              string        `json:"timed_task_interval" mapstructure:"timed_task_interval"`
	L1ReceiptConfirmTimeout        string        `json:"l1_receipt_confirm_timeout" mapstructure:"l1_receipt_confirm_timeout"`
	L1ConfirmBlocks                int           `json:"l1_confirm_blocks" mapstructure:"l1_confirm_blocks"`
	SignedBatchesWindow            int           `json:"signed_batches_window" mapstructure:"signed_batches_window"`
	MinSignedInWindow              int           `json:"min_signed_in_window" mapstructure:"min_signed_in_window"`
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

type NodeConfig struct {
	BaseDir      string `json:"base_dir" mapstructure:"base_dir"`
	DBDir        string `json:"db_dir" mapstructure:"db_dir"`
	WsAddr       string `json:"ws_addr" mapstructure:"ws_addr"`
	HttpAddr     string `json:"http_addr" mapstructure:"http_addr"`
	L2EthRpc     string `json:"l2_eth_rpc" mapstructure:"l2_eth_rpc"`
	DisableHTTP2 bool   `json:"disable_http2" mapstructure:"disable_http2"`
	PrivateKey   string `json:"private_key" mapstructure:"private_key"`

	PreParamFile    string        `json:"pre_param_file" mapstructure:"pre_param_file"`
	P2PPort         string        `json:"p2p_port" mapstructure:"p2p_port"`
	BootstrapPeers  string        `json:"bootstrap_peers" mapstructure:"bootstrap_peers"`
	ExternalIP      string        `json:"external_ip" mapstructure:"external_ip"`
	KeyGenTimeout   time.Duration `json:"key_gen_timeout" mapstructure:"key_gen_timeout"`
	KeySignTimeout  time.Duration `json:"key_sign_timeout" mapstructure:"key_sign_timeout"`
	PreParamTimeout time.Duration `json:"pre_param_timeout" mapstructure:"pre_param_timeout"`

	Secrets SecretsManagerConfig `json:"secrets" mapstructure:"secrets"`
	Shamir  ShamirConfig         `json:"shamir" mapstructure:"shamir"`
}

type SecretsManagerConfig struct {
	Enable   bool   `json:"enable" mapstructure:"enable"`
	SecretId string `json:"secret_id" mapstructure:"secret_id"`
}

type ShamirConfig struct {
	Enable bool      `json:"enable" mapstructure:"enable"`
	Kms    KmsConfig `json:"kms" mapstructure:"kms"`
	S3     S3Config  `json:"s3" mapstructure:"s3"`
	Sm     SMConfig  `json:"sm" mapstructure:"sm"`
	Xor    string    `json:"xor" mapstructure:"xor"`
}

type KmsConfig struct {
	KeyId  string     `json:"key_id" mapstructure:"key_id"`
	Region string     `json:"region" mapstructure:"region"`
	Aksk   AKSKConfig `json:"aksk" mapstructure:"aksk"`
}

type S3Config struct {
	Region  string     `json:"region" mapstructure:"region"`
	Buckets string     `json:"buckets" mapstructure:"buckets"`
	Aksk    AKSKConfig `json:"aksk" mapstructure:"aksk"`
}

type SMConfig struct {
	Region    string     `json:"region" mapstructure:"region"`
	SecretIds string     `json:"secretIds" mapstructure:"secretIds"`
	Aksk      AKSKConfig `json:"aksk" mapstructure:"aksk"`
}

type AKSKConfig struct {
	Id     string `json:"id" mapstructure:"id"`
	Secret string `json:"secret" mapstructure:"secret"`
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
		Node: NodeConfig{
			P2PPort:         "8000",
			KeyGenTimeout:   10 * time.Second,
			KeySignTimeout:  10 * time.Second,
			PreParamTimeout: 5 * time.Minute,
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
