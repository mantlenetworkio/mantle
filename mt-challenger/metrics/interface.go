package metrics

import "github.com/prometheus/client_golang/prometheus"

type ChallengerMetrics interface {
	BalanceETH() prometheus.Gauge

	NonceETH() prometheus.Counter

	ReRollupBatchIndex() prometheus.Counter

	CheckBatchIndex() prometheus.Counter
}
