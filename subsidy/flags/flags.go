package flags

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var (
	UserDir, _          = os.UserHomeDir()
	EthereumHttpUrlFlag = cli.StringFlag{
		Name:   "ethereum-http-url",
		Value:  "http://127.0.0.1:8545",
		Usage:  "L1 HTTP Endpoint",
		EnvVar: "SUBSIDY_ETHEREUM_HTTP_URL",
	}
	SCCAddressFlag = cli.StringFlag{
		Name:   "scc-address",
		Usage:  "Address of SCC_CONTRACT",
		Value:  "",
		EnvVar: "SUBSIDY_SCC_ADDRESS",
	}

	SCCTopicFlag = cli.StringFlag{
		Name:   "scc-topic",
		Usage:  "Address of SCC_CONTRACT",
		Value:  "StateBatchAppended(uint256,bytes32,uint256,uint256,bytes)",
		EnvVar: "SUBSIDY_SCC_TOPIC",
	}

	CTCAddressFlag = cli.StringFlag{
		Name:   "scc-address",
		Usage:  "Address of CTC_CONTRACT",
		Value:  "",
		EnvVar: "SUBSIDY_CTC_ADDRESS",
	}

	CTCTopicFlag = cli.StringFlag{
		Name:   "scc-topic",
		Usage:  "Address of CTC_CONTRACT",
		Value:  "SequencerBatchAppended(uint256,uint256,uint256)",
		EnvVar: "SUBSIDY_CTC_TOPIC",
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "Private Key corresponding to SUBSIDY Owner",
		EnvVar: "SUBSIDY_PAYER_PRIVATE_KEY",
	}
	LogLevelFlag = cli.IntFlag{
		Name:   "loglevel",
		Value:  3,
		Usage:  "log level to emit to the screen",
		EnvVar: "SUBSIDY_LOG_LEVEL",
	}
	L1QueryEpochLengthSecondsFlag = cli.Uint64Flag{
		Name:   "l1-query-epoch-length-seconds",
		Value:  60 * 60 * 24,
		Usage:  "query rollup cost epoch length seconds",
		EnvVar: "SUBSIDY_L1_QUERY_EPOCH_LENGTH_SECONDS",
	}
	WaitForReceiptFlag = cli.BoolFlag{
		Name:   "wait-for-receipt",
		Usage:  "wait for receipts when sending transactions",
		EnvVar: "SUBSIDY_WAIT_FOR_RECEIPT",
	}
	HomeDirFlag = cli.StringFlag{
		Name:   "home-dir",
		Usage:  "subsidy work home dir",
		Value:  filepath.Join(UserDir, ".subsidy"),
		EnvVar: "SUBSIDY_HOME_DIR",
	}
	CacheDirFlag = cli.StringFlag{
		Name:   "cache-dir",
		Usage:  "subsidy work cache dir",
		Value:  ".cache",
		EnvVar: "SUBSIDY_CACHE_DIR",
	}
	FileNameFlag = cli.StringFlag{
		Name:   "file-name",
		Usage:  "subsidy file dir for record payer state",
		Value:  "payerState",
		EnvVar: "SUBSIDY_FILE_NAME",
	}
	StartBlockFlag = cli.Uint64Flag{
		Name:   "start-block",
		Usage:  "payer start block",
		EnvVar: "SUBSIDY_START_BLOCK",
	}
	RevisedBlockFlag = cli.Uint64Flag{
		Name:   "revised-block",
		Usage:  "payer revised block,if not empty block will be reset",
		EnvVar: "SUBSIDY_REVISED_BLOCk",
	}
	MetricsEnabledFlag = cli.BoolFlag{
		Name:   "metrics",
		Usage:  "Enable metrics collection and reporting",
		EnvVar: "SUBSIDY_METRICS_ENABLE",
	}
	MetricsHTTPFlag = cli.StringFlag{
		Name:   "metrics.addr",
		Usage:  "Enable stand-alone metrics HTTP server listening interface",
		Value:  "127.0.0.1",
		EnvVar: "SUBSIDY_METRICS_HTTP",
	}
	MetricsPortFlag = cli.IntFlag{
		Name:   "metrics.port",
		Usage:  "Metrics HTTP server listening port",
		Value:  6060,
		EnvVar: "SUBSIDY_METRICS_PORT",
	}
	MetricsEnableInfluxDBFlag = cli.BoolFlag{
		Name:   "metrics.influxdb",
		Usage:  "Enable metrics export/push to an external InfluxDB database",
		EnvVar: "SUBSIDY_METRICS_ENABLE_INFLUX_DB",
	}
	MetricsInfluxDBEndpointFlag = cli.StringFlag{
		Name:   "metrics.influxdb.endpoint",
		Usage:  "InfluxDB API endpoint to report metrics to",
		Value:  "http://localhost:8086",
		EnvVar: "SUBSIDY_METRICS_INFLUX_DB_ENDPOINT",
	}
	MetricsInfluxDBDatabaseFlag = cli.StringFlag{
		Name:   "metrics.influxdb.database",
		Usage:  "InfluxDB database name to push reported metrics to",
		Value:  "gas-oracle",
		EnvVar: "SUBSIDY_METRICS_INFLUX_DB_DATABASE",
	}
	MetricsInfluxDBUsernameFlag = cli.StringFlag{
		Name:   "metrics.influxdb.username",
		Usage:  "Username to authorize access to the database",
		Value:  "test",
		EnvVar: "SUBSIDY_METRICS_INFLUX_DB_USERNAME",
	}
	MetricsInfluxDBPasswordFlag = cli.StringFlag{
		Name:   "metrics.influxdb.password",
		Usage:  "Password to authorize access to the database",
		Value:  "test",
		EnvVar: "SUBSIDY_METRICS_INFLUX_DB_PASSWORD",
	}
)

var Flags = []cli.Flag{
	EthereumHttpUrlFlag,
	SCCAddressFlag,
	CTCAddressFlag,
	SCCTopicFlag,
	CTCTopicFlag,
	PrivateKeyFlag,
	LogLevelFlag,
	L1QueryEpochLengthSecondsFlag,
	WaitForReceiptFlag,
	HomeDirFlag,
	CacheDirFlag,
	FileNameFlag,
	StartBlockFlag,
	RevisedBlockFlag,
	MetricsEnabledFlag,
	MetricsHTTPFlag,
	MetricsPortFlag,
	MetricsEnableInfluxDBFlag,
	MetricsInfluxDBEndpointFlag,
	MetricsInfluxDBDatabaseFlag,
	MetricsInfluxDBUsernameFlag,
	MetricsInfluxDBPasswordFlag,
}
