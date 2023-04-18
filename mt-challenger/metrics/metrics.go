package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ChallengerBase struct {
	balanceETH         prometheus.Gauge
	nonceETH           prometheus.Counter
	reRollupBatchIndex prometheus.Counter
	checkBatchIndex    prometheus.Counter
}

func NewChallengerBase() *ChallengerBase {
	return &ChallengerBase{
		balanceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "balance_eth",
			Help:      "ETH balance of the mt batch",
			Subsystem: "mt-challenger",
		}),

		nonceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "nonce_eth",
			Help:      "nonce for mt batch address",
			Subsystem: "mt-challenger",
		}),

		reRollupBatchIndex: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "re_rollup_batch_index",
			Help:      "re-rollup batch index for eigen layer",
			Subsystem: "mt-challenger",
		}),

		checkBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "checker_batch_index",
			Help:      "checker batch index for eigen layer",
			Subsystem: "mt-challenger",
		}),
	}
}

func (cb *ChallengerBase) BalanceETH() prometheus.Gauge {
	return cb.balanceETH
}

func (cb *ChallengerBase) NonceETH() prometheus.Counter {
	return cb.nonceETH
}

func (cb *ChallengerBase) ReRollUpBatchIndex() prometheus.Counter {
	return cb.reRollupBatchIndex
}

func (cb *ChallengerBase) CheckRollUpBatchIndex() prometheus.Counter {
	return cb.checkBatchIndex
}
