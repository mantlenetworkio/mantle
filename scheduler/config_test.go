package scheduler_test

import (
	"fmt"
	"testing"

	scheduler "github.com/bitdao-io/bitnetwork/scheduler"
	"github.com/stretchr/testify/require"
)

var validateConfigTests = []struct {
	name   string
	cfg    scheduler.Config
	expErr error
}{
	{
		name: "bad log level",
		cfg: scheduler.Config{
			LogLevel: "unknown",
		},
		expErr: fmt.Errorf("unknown level: unknown"),
	},
	{
		name: "sequencer priv key or mnemonic none set",
		cfg: scheduler.Config{
			LogLevel: "info",

			SequencerPrivateKey: "",
			Mnemonic:            "",
			SequencerHDPath:     "",
		},
		expErr: scheduler.ErrSequencerPrivKeyOrMnemonic,
	},
	{
		name: "sequencer priv key or mnemonic both set",
		cfg: scheduler.Config{
			LogLevel: "info",

			SequencerPrivateKey: "sequencer-privkey",
			Mnemonic:            "mnemonic",
			SequencerHDPath:     "sequencer-path",
		},
		expErr: scheduler.ErrSequencerPrivKeyOrMnemonic,
	},
	{
		name: "sequencer priv key or mnemonic only mnemonic set",
		cfg: scheduler.Config{
			LogLevel: "info",

			SequencerPrivateKey: "",
			Mnemonic:            "mnemonic",
			SequencerHDPath:     "",
		},
		expErr: scheduler.ErrSequencerPrivKeyOrMnemonic,
	},
	{
		name: "sequencer priv key or mnemonic only hdpath set",
		cfg: scheduler.Config{
			LogLevel: "info",

			SequencerPrivateKey: "",
			Mnemonic:            "",
			SequencerHDPath:     "sequencer-path",
		},
		expErr: scheduler.ErrSequencerPrivKeyOrMnemonic,
	},
	{
		name: "proposer priv key or mnemonic none set",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",

			ProposerPrivateKey: "",
			Mnemonic:           "",
			ProposerHDPath:     "",
		},
		expErr: scheduler.ErrProposerPrivKeyOrMnemonic,
	},
	{
		name: "proposer priv key or mnemonic both set",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",

			ProposerPrivateKey: "proposer-privkey",
			Mnemonic:           "mnemonic",
			ProposerHDPath:     "proposer-path",
		},
		expErr: scheduler.ErrProposerPrivKeyOrMnemonic,
	},
	{
		name: "proposer priv key or mnemonic only mnemonic set",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",

			ProposerPrivateKey: "",
			Mnemonic:           "mnemonic",
			ProposerHDPath:     "",
		},
		expErr: scheduler.ErrProposerPrivKeyOrMnemonic,
	},
	{
		name: "proposer priv key or mnemonic only hdpath set",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",

			ProposerPrivateKey: "",
			Mnemonic:           "",
			ProposerHDPath:     "proposer-path",
		},
		expErr: scheduler.ErrProposerPrivKeyOrMnemonic,
	},
	{
		name: "same sequencer and proposer hd path",
		cfg: scheduler.Config{
			LogLevel: "info",

			Mnemonic:        "mnemonic",
			SequencerHDPath: "path",
			ProposerHDPath:  "path",
		},
		expErr: scheduler.ErrSameSequencerAndProposerHDPath,
	},
	{
		name: "same sequencer and proposer privkey",
		cfg: scheduler.Config{
			LogLevel: "info",

			SequencerPrivateKey: "privkey",
			ProposerPrivateKey:  "privkey",
		},
		expErr: scheduler.ErrSameSequencerAndProposerPrivKey,
	},
	{
		name: "sentry-dsn not set when sentry-enable is true",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",
			ProposerPrivateKey:  "proposer-privkey",

			SentryEnable: true,
			SentryDsn:    "",
		},
		expErr: scheduler.ErrSentryDSNNotSet,
	},
	// Valid configs
	{
		name: "valid config with privkeys and no sentry",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",
			ProposerPrivateKey:  "proposer-privkey",
			SentryEnable:        false,
			SentryDsn:           "",
		},
		expErr: nil,
	},
	{
		name: "valid config with mnemonic and no sentry",
		cfg: scheduler.Config{
			LogLevel:        "info",
			Mnemonic:        "mnemonic",
			SequencerHDPath: "sequencer-path",
			ProposerHDPath:  "proposer-path",
			SentryEnable:    false,
			SentryDsn:       "",
		},
		expErr: nil,
	},
	{
		name: "valid config with privkeys and sentry",
		cfg: scheduler.Config{
			LogLevel:            "info",
			SequencerPrivateKey: "sequencer-privkey",
			ProposerPrivateKey:  "proposer-privkey",
			SentryEnable:        true,
			SentryDsn:           "scheduler",
		},
		expErr: nil,
	},
	{
		name: "valid config with mnemonic and sentry",
		cfg: scheduler.Config{
			LogLevel:        "info",
			Mnemonic:        "mnemonic",
			SequencerHDPath: "sequencer-path",
			ProposerHDPath:  "proposer-path",
			SentryEnable:    true,
			SentryDsn:       "scheduler",
		},
		expErr: nil,
	},
}

// TestValidateConfig asserts the behavior of ValidateConfig by testing expected
// error and success configurations.
func TestValidateConfig(t *testing.T) {
	for _, test := range validateConfigTests {
		t.Run(test.name, func(t *testing.T) {
			err := scheduler.ValidateConfig(&test.cfg)
			require.Equal(t, err, test.expErr)
		})
	}
}
