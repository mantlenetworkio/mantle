package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "BATCH_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   "BUILD_ENV",
	}
	EthNetworkNameFlag = cli.StringFlag{
		Name:     "eth-network-name",
		Usage:    "Ethereum network name",
		Required: true,
		EnvVar:   "ETH_NETWORK_NAME",
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   "L1_ETH_RPC",
	}
	L2EthRpcFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVar:   "L2_ETH_RPC",
	}
	TssClientUrl = cli.StringFlag{
		Name:     "tss-client-url",
		Usage:    "HTTP provider URL for tss",
		Required: true,
		EnvVar:   "TSS_CLIENT_RPC",
	}
	JwtSecret = cli.StringFlag{
		Name:     "jwt-secret",
		Usage:    "jet access secret",
		Required: false,
		EnvVar:   "JWT_SECRET",
	}
	DaAddressFlag = cli.StringFlag{
		Name:     "da-address",
		Usage:    "Address of the da contract",
		Required: true,
		EnvVar:   "DA_ADDRESS",
	}
	DaUpgradeBlockFlag = cli.Uint64Flag{
		Name:     "da-upgrade-block",
		Usage:    "eigen layer upgrade block",
		Required: true,
		EnvVar:   prefixEnvVar("DA_UPGRADE_BLOCK"),
	}
	CTCAddressFlag = cli.StringFlag{
		Name:     "ctc-address",
		Usage:    "Address of the CTC contract",
		Required: true,
		EnvVar:   "CTC_ADDRESS",
	}
	SCCAddressFlag = cli.StringFlag{
		Name:     "scc-address",
		Usage:    "Address of the SCC contract",
		Required: true,
		EnvVar:   "SCC_ADDRESS",
	}
	FPRollupAddressFlag = cli.StringFlag{
		Name:     "fraud-proof-rollup-address",
		Usage:    "Address of the FraudProof Rollup contract",
		Required: true,
		EnvVar:   "FP_ROLLUP_ADDRESS",
	}
	MinStateRootElementsFlag = cli.Uint64Flag{
		Name: "min-state-root-elements",
		Usage: "Minimum number of elements required to submit a state " +
			"root batch",
		Required: true,
		EnvVar:   prefixEnvVar("MIN_STATE_ROOT_ELEMENTS"),
	}
	MaxStateRootElementsFlag = cli.Uint64Flag{
		Name: "max-state-root-elements",
		Usage: "Maximum number of elements required to submit a state " +
			"root batch",
		Required: true,
		EnvVar:   prefixEnvVar("MAX_STATE_ROOT_ELEMENTS"),
	}
	RollupTimeoutFlag = cli.DurationFlag{
		Name:     "rollup-timeout",
		Usage:    "Delay between rollup timeout transactions ",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_TIMEOUT"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("POLL_INTERVAL"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("NUM_CONFIRMATIONS"),
	}
	SafeAbortNonceTooLowCountFlag = cli.Uint64Flag{
		Name: "safe-abort-nonce-too-low-count",
		Usage: "Number of ErrNonceTooLow observations required to " +
			"give up on a tx at a particular nonce without receiving " +
			"confirmation",
		Required: true,
		EnvVar:   prefixEnvVar("SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   prefixEnvVar("RESUBMISSION_TIMEOUT"),
	}
	FinalityConfirmationsFlag = cli.Uint64Flag{
		Name: "finality-confirmations",
		Usage: "Number of confirmations that we should wait before " +
			"submitting state roots for CTC elements",
		Required: true,
		EnvVar:   prefixEnvVar("FINALITY_CONFIRMATIONS"),
	}
	RunTxBatchSubmitterFlag = cli.BoolFlag{
		Name:     "run-tx-batch-submitter",
		Usage:    "Determines whether or not to run the tx batch submitter",
		Required: true,
		EnvVar:   prefixEnvVar("RUN_TX_BATCH_SUBMITTER"),
	}
	RunStateBatchSubmitterFlag = cli.BoolFlag{
		Name:     "run-state-batch-submitter",
		Usage:    "Determines whether or not to run the state batch submitter",
		Required: true,
		EnvVar:   prefixEnvVar("RUN_STATE_BATCH_SUBMITTER"),
	}
	SafeMinimumEtherBalanceFlag = cli.Uint64Flag{
		Name: "safe-minimum-ether-balance",
		Usage: "Safe minimum amount of ether the batch submitter key " +
			"should hold before it starts to log errors",
		Required: true,
		EnvVar:   prefixEnvVar("SAFE_MINIMUM_ETHER_BALANCE"),
	}
	ClearPendingTxsFlag = cli.BoolFlag{
		Name: "clear-pending-txs",
		Usage: "Whether or not to clear pending transaction in the " +
			"mempool on startup",
		Required: true,
		EnvVar:   prefixEnvVar("CLEAR_PENDING_TXS"),
	}

	/* Optional Flags */

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar("LOG_TERMINAL"),
	}
	SentryEnableFlag = cli.BoolFlag{
		Name:   "sentry-enable",
		Usage:  "Whether or not to enable Sentry. If true, sentry-dsn must also be set",
		EnvVar: prefixEnvVar("SENTRY_ENABLE"),
	}
	SentryDsnFlag = cli.StringFlag{
		Name:   "sentry-dsn",
		Usage:  "Sentry data source name",
		EnvVar: prefixEnvVar("SENTRY_DSN"),
	}
	SentryTraceRateFlag = cli.DurationFlag{
		Name:   "sentry-trace-rate",
		Usage:  "Sentry trace rate",
		Value:  50 * time.Millisecond,
		EnvVar: prefixEnvVar("SENTRY_TRACE_RATE"),
	}
	BlockOffsetFlag = cli.Uint64Flag{
		Name:   "block-offset",
		Usage:  "The offset between the CTC contract start and the L2 geth blocks",
		Value:  1,
		EnvVar: prefixEnvVar("BLOCK_OFFSET"),
	}
	SequencerPrivateKeyFlag = cli.StringFlag{
		Name:   "sequencer-private-key",
		Usage:  "The private key to use for sending to the sequencer contract",
		EnvVar: prefixEnvVar("SEQUENCER_PRIVATE_KEY"),
	}
	ProposerPrivateKeyFlag = cli.StringFlag{
		Name:   "proposer-private-key",
		Usage:  "The private key to use for sending to the proposer contract",
		EnvVar: prefixEnvVar("PROPOSER_PRIVATE_KEY"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the proposer",
		EnvVar: prefixEnvVar("MNEMONIC"),
	}
	SequencerHDPathFlag = cli.StringFlag{
		Name: "sequencer-hd-path",
		Usage: "The HD path used to derive the sequencer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: prefixEnvVar("SEQUENCER_HD_PATH"),
	}
	ProposerHDPathFlag = cli.StringFlag{
		Name: "proposer-hd-path",
		Usage: "The HD path used to derive the proposer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: prefixEnvVar("PROPOSER_HD_PATH"),
	}
	SequencerBatchType = cli.StringFlag{
		Name:   "sequencer-batch-type",
		Usage:  "The type of sequencer batch to be submitted. Valid arguments are legacy or zlib.",
		Value:  "legacy",
		EnvVar: prefixEnvVar("SEQUENCER_BATCH_TYPE"),
	}
	MetricsServerEnableFlag = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Whether or not to run the embedded metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostnameFlag = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "The hostname of the metrics server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPortFlag = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "The port of the metrics server",
		Value:  7300,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar("HTTP2_DISABLE"),
	}
	SccRollbackFlag = cli.BoolFlag{
		Name:   "EnableSccRollbackFlag",
		Usage:  "Whether or not to enable scc rollback.",
		EnvVar: prefixEnvVar("SCC_ROLLBACK"),
	}
	EnableSequencerHsmFlag = cli.BoolFlag{
		Name:   "enable-sequencer-hsm",
		Usage:  "Whether or not to use cloudhsm for sequencer",
		EnvVar: prefixEnvVar("ENABLE_SEQUENCER_HSM"),
	}
	SequencerHsmAddressFlag = cli.StringFlag{
		Name:   "sequencer-hsm-address",
		Usage:  "The address of private-key in hsm for sequencer",
		Value:  "",
		EnvVar: prefixEnvVar("SEQUENCER_HSM_ADDRESS"),
	}
	SequencerHsmAPIName = cli.StringFlag{
		Name:   "sequencer-hsm-api-name",
		Usage:  "The api-name of private-key in hsm for sequencer",
		Value:  "",
		EnvVar: prefixEnvVar("SEQUENCER_HSM_API_NAME"),
	}
	SequencerHsmCreden = cli.StringFlag{
		Name:   "sequencer-hsm-creden",
		Usage:  "The creden of private-key in hsm for sequencer",
		Value:  "",
		EnvVar: prefixEnvVar("SEQUENCER_HSM_CREDEN"),
	}
	EnableProposerHsmFlag = cli.BoolFlag{
		Name:   "enable-proposer-hsm",
		Usage:  "Whether or not to use cloudhsm for proposer",
		EnvVar: prefixEnvVar("ENABLE_PROPOSER_HSM"),
	}
	ProposerHsmAddressFlag = cli.StringFlag{
		Name:   "proposer-hsm-address",
		Usage:  "The address of private-key in hsm for proposer",
		Value:  "",
		EnvVar: prefixEnvVar("PROPOSER_HSM_ADDRESS"),
	}
	ProposerHsmAPIName = cli.StringFlag{
		Name:   "proposer-hsm-api-name",
		Usage:  "The api-name of private-key in hsm for proposer",
		Value:  "",
		EnvVar: prefixEnvVar("PROPOSER_HSM_API_NAME"),
	}
	ProposerHsmCreden = cli.StringFlag{
		Name:   "proposer-hsm-creden",
		Usage:  "The creden of private-key in hsm for proposer",
		Value:  "",
		EnvVar: prefixEnvVar("PROPOSER_HSM_CREDEN"),
	}
	RollupClientHttpFlag = cli.StringFlag{
		Name:   "rollup.clienthttp",
		Usage:  "HTTP endpoint for the rollup client",
		Value:  "http://localhost:7878",
		EnvVar: "ROLLUP_CLIENT_HTTP",
	}
	AllowL2AutoRollback = cli.BoolFlag{
		Name:     "rollup.allow-l2-auto-rollback",
		Usage:    "Trigger for allowing layer2 auto rollback",
		Required: false,
		EnvVar:   "ROLLUP_ALLOW_L2_AUTO_ROLLBACK",
	}

	MinRollupTxnFlag = cli.Uint64Flag{
		Name:     "min-rollup-txn",
		Usage:    "Minimum number of transaction from l2geth which is used to submit to rollup",
		Required: true,
		EnvVar:   prefixEnvVar("MIN_ROLLUP_TXN"),
	}

	MaxRollupTxnFlag = cli.Uint64Flag{
		Name:     "max-rollup-txn",
		Usage:    "Maximum number of transaction from l2geth which is used to submit to rollup",
		Required: true,
		EnvVar:   prefixEnvVar("MAX_ROLLUP_TXN"),
	}

	MinTimeoutStateRootElementsFlag = cli.Uint64Flag{
		Name: "min-timeout-state-root-elements",
		Usage: "Minimum number of elements required to submit a state " +
			"root batch",
		Required: true,
		EnvVar:   prefixEnvVar("MIN_TIMEOUT_STATE_ROOT_ELEMENTS"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	EthNetworkNameFlag,
	L1EthRpcFlag,
	L2EthRpcFlag,
	TssClientUrl,
	JwtSecret,
	DaAddressFlag,
	DaUpgradeBlockFlag,
	CTCAddressFlag,
	SCCAddressFlag,
	FPRollupAddressFlag,
	MinStateRootElementsFlag,
	MaxStateRootElementsFlag,
	RollupTimeoutFlag,
	PollIntervalFlag,
	NumConfirmationsFlag,
	SafeAbortNonceTooLowCountFlag,
	ResubmissionTimeoutFlag,
	FinalityConfirmationsFlag,
	RunTxBatchSubmitterFlag,
	RunStateBatchSubmitterFlag,
	SafeMinimumEtherBalanceFlag,
	ClearPendingTxsFlag,
	MaxRollupTxnFlag,
	MinRollupTxnFlag,
	MinTimeoutStateRootElementsFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	SentryEnableFlag,
	SentryDsnFlag,
	SentryTraceRateFlag,
	BlockOffsetFlag,
	SequencerBatchType,
	SequencerPrivateKeyFlag,
	ProposerPrivateKeyFlag,
	MnemonicFlag,
	SequencerHDPathFlag,
	ProposerHDPathFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	HTTP2DisableFlag,
	EnableProposerHsmFlag,
	ProposerHsmAddressFlag,
	ProposerHsmAPIName,
	ProposerHsmCreden,
	EnableSequencerHsmFlag,
	SequencerHsmAddressFlag,
	SequencerHsmAPIName,
	SequencerHsmCreden,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
