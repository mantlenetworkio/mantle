package metrics

import "github.com/prometheus/client_golang/prometheus"

type ChallengerMetrics interface {
	BalanceETH() prometheus.Gauge

	NonceETH() prometheus.Gauge

	ReRollupBatchIndex() prometheus.Gauge

	CheckBatchIndex() prometheus.Gauge

	DataStoreId() prometheus.Gauge
}
