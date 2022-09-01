package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	DefaultBaseDir  string
	DefaultSecretId string
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultBaseDir = filepath.Join(userHomeDir, ".tssnode")
}

type Configuration struct {
	BaseConfig `json:"base" mapstructure:"base"`
	Key        KeyConfig            `json:"key" mapstructure:"key"`
	Secrets    SecretsManagerConfig `json:"secrets" mapstructure:"secrets"`
	Shamir     ShamirConfig         `json:"shamir" mapstructure:"shamir"`
}

type BaseConfig struct {
	BaseDir                   string        `json:"base_dir" mapstructure:"base_dir"`
	DBDir                     string        `json:"db_dir" mapstructure:"db_dir"`
	PreParamFile              string        `json:"pre_param_file" mapstructure:"pre_param_file"`
	P2PPort                   string        `json:"p2p_port" mapstructure:"p2p_port"`
	BootstrapPeers            string        `json:"bootstrap_peers" mapstructure:"bootstrap_peers"`
	ExternalIP                string        `json:"external_ip" mapstructure:"external_ip"`
	KeyGenTimeout             time.Duration `json:"key_gen_timeout" mapstructure:"key_gen_timeout"`
	KeySignTimeout            time.Duration `json:"key_sign_timeout" mapstructure:"key_sign_timeout"`
	PreParamTimeout           time.Duration `json:"pre_param_timeout" mapstructure:"pre_param_timeout"`
	PauseSigning              bool          `json:"pause_signing" mapstructure:"pause_signing"`
	WsAddr                    string        `json:"ws_addr" mapstructure:"ws_addr"`
	L2EthRpc                  string        `json:"l2_eth_rpc" mapstructure:"l2_eth_rpc"`
	L1EthRpc                  string        `json:"l1_eth_rpc" mapstructure:"l1_eth_rpc"`
	DisableHTTP2              bool          `json:"disable_http2" mapstructure:"disable_http2"`
	TssGroupManagerAddress    string        `json:"tss_group_manager_address" mapstructure:"tss_group_manager_address"`
	TssStakingSlashingAddress string        `json:"tss_staking_slashing_address" mapstructure:"tss_staking_slashing_address"`
}

func DefaultBaseConfiguration() BaseConfig {
	return BaseConfig{
		BaseDir:         DefaultBaseDir,
		DBDir:           filepath.Join(DefaultBaseDir, "db"),
		P2PPort:         "8000",
		KeyGenTimeout:   10 * time.Second,
		KeySignTimeout:  10 * time.Second,
		PreParamTimeout: 5 * time.Minute,
	}
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
		BaseConfig: DefaultBaseConfiguration(),
		Key: KeyConfig{
			KeyringDir:     DefaultBaseDir,
			KeyringBackend: "os",
		},
		Secrets: SecretsManagerConfig{
			Enable:   false,
			SecretId: DefaultSecretId,
		},
		Shamir: ShamirConfig{
			Enable: false,
			Xor:    "",
		},
	}
}

type KeyConfig struct {
	PrivateKey     string `json:"private_key" mapstructure:"private_key"` //private key hex based on secp256k1
	KeyringName    string `json:"keyring_name" mapstructure:"keyring_name"`
	KeyringDir     string `json:"keyring_dir" mapstructure:"keyring_dir"`
	KeyringBackend string `json:"keyring_backend" mapstructure:"keyring_backend"`
}

type SecretsManagerConfig struct {
	Enable   bool   `json:"enable" mapstructure:"enable"`
	SecretId string `json:"secret_id" mapstructure:"secret_id"`
}

func LoadConfig(configFile string, configuration interface{}) error {
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(configFile))
	viper.SetConfigName(strings.Split(path.Base(configFile), ".")[0])
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read from config file: %w", err)
	}
	viper.SetEnvPrefix("tss")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	if err := viper.Unmarshal(configuration); err != nil {
		return fmt.Errorf("fail to unmarshal: %w", err)
	}
	return nil
}
