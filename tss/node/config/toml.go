package config

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

const defaultBaseConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml_lang/toml

###############################################################################
###                           Base Configuration                            ###
###############################################################################

[base]

# the base directory for storing the data, tss localSaveData etc.
base_dir = "{{ .BaseConfig.BaseDir }}"

# the preParamFile for tss job
pre_param_file = "{{ .BaseConfig.PreParamFile }}"

# listening port for p2p communication
p2p_port = "{{ .BaseConfig.P2PPort }}"

# the address ip we should advertise to the network
external_ip = "{{ .BaseConfig.ExternalIP }}"

# bootstrap peers the node is going to connecting to once the node started
# format: /ip4/${ip}/tcp/${port}/p2p/${peerID},..
# the self peerID can be checked via command of 'tssnode keys parse-peer-id -h '
bootstrap_peers = "{{ .BaseConfig.BootstrapPeers }}"

# timeout setting during tss job
key_gen_timeout = "{{ .BaseConfig.KeyGenTimeout }}"
key_sign_timeout = "{{ .BaseConfig.KeySignTimeout }}"
pre_param_timeout = "{{ .BaseConfig.PreParamTimeout }}"

# pause tss signing process when true
pause_signing = "{{ .BaseConfig.PauseSigning }}"

# websocket remote address
ws_addr = "{{ .BaseConfig.WsAddr }}"

`

const defaultKeyConfigTemplate = `
###############################################################################
###                           Key Configuration                             ###
###############################################################################

[key]

# The private key for identifying the node, it should be hex string here without '0x'.
# It is unsafe to put the raw private key here in the file, it would be nice to
# set it to environment with the prefix 'TSS', export TSS_KEY_PRIVATE_KEY="981a3e...."
# or it is recommended to store the private key into keyring, you can checkout with 'tssnode keys -h'.
private_key = "{{ .Key.PrivateKey }}"

# the name stored in the keyring
keyring_name = "{{ .Key.KeyringName }}"

# keyring directory; if omitted, the default 'base_dir' directory will be used
keyring_dir = "{{ .Key.KeyringDir }}"

# keyring backend(os|file|test) default: os
keyring_backend = "{{ .Key.KeyringBackend }}"

`
const defaultSecretsManagerTemplate = `
###############################################################################
###                         Secrets Manager Configuration                   ###
###############################################################################

[secrets]
# whether use aws secrets manager to storage important content
enable = {{ .Secrets.Enable }}

# the secret id must be need to get or pull data from aws secrets manager, if enable=true ,secret id must be config
secret_id = "{{ .Secrets.SecretId }}"

`

const defaultShamirTemplate = `
###############################################################################
###                         Shamir Manager Configuration                   ###
###############################################################################

[shamir]
# whether use shamir to storage important content,default : false
enable = {{ .Shamir.Enable }}

# xor key for XOR
xor = {{ .Shamir.Xor }}

[shamir.kms]
key_id = "{{ .Shamir.Kms.KeyId }}"
region = "{{ .Shamir.Kms.Region }}"
aksk.id = "{{ .Shamir.Kms.Aksk.Id }}"
aksk.secret = "{{ .Shamir.Kms.Aksk.Secret }}"

[shamir.s3]
region = "{{ .Shamir.S3.Region }}"
buckets = "{{ .Shamir.S3.Buckets }}"
aksk.id = "{{ .Shamir.Kms.Aksk.Id }}"
aksk.secret = "{{ .Shamir.Kms.Aksk.Secret }}"

[shamir.sm]
region = "{{ .Shamir.Sm.Region }}"
secretIds = "{{ .Shamir.Sm.SecretIds }}"
aksk.id = "{{ .Shamir.Kms.Aksk.Id }}"
aksk.secret = "{{ .Shamir.Kms.Aksk.Secret }}"

`

const DefaultConfigTemplate = defaultBaseConfigTemplate + defaultKeyConfigTemplate + defaultSecretsManagerTemplate + defaultShamirTemplate

var configTemplate *template.Template

func init() {
	var err error

	tmpl := template.New("appConfigFileTemplate")
	if configTemplate, err = tmpl.Parse(DefaultConfigTemplate); err != nil {
		panic(err)
	}

}

func WriteStartConfigFile(configFilePath string, config interface{}) {
	var buffer bytes.Buffer
	if err := configTemplate.Execute(&buffer, config); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(configFilePath, buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}
