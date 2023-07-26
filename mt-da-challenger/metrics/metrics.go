package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ChallengerBase struct {
	balanceETH         prometheus.Gauge
	nonceETH           prometheus.Gauge
	reRollupBatchIndex prometheus.Gauge
	checkBatchIndex    prometheus.Gauge
	dataStoreId        prometheus.Gauge
}

func NewChallengerBase() *ChallengerBase {
	return &ChallengerBase{
		balanceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "balance_eth",
			Help:      "ETH balance of the mt batch",
			Subsystem: "challenger",
		}),

		nonceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "nonce_eth",
			Help:      "nonce for mt batch address",
			Subsystem: "challenger",
		}),

		reRollupBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "re_rollup_batch_index",
			Help:      "re-rollup batch index for eigen layer",
			Subsystem: "challenger",
		}),

		checkBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "checker_batch_index",
			Help:      "checker batch index for eigen layer",
			Subsystem: "challenger",
		}),
		dataStoreId: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "data_store_id",
			Help:      "current rollup da data_store_id",
			Subsystem: "mtbatcher",
		}),
	}
}

func (cb *ChallengerBase) BalanceETH() prometheus.Gauge {
	return cb.balanceETH
}

func (cb *ChallengerBase) NonceETH() prometheus.Gauge {
	return cb.nonceETH
}

func (cb *ChallengerBase) ReRollupBatchIndex() prometheus.Gauge {
	return cb.reRollupBatchIndex
}

func (cb *ChallengerBase) CheckBatchIndex() prometheus.Gauge {
	return cb.checkBatchIndex
}

func (cb *ChallengerBase) DataStoreId() prometheus.Gauge {
	return cb.dataStoreId
}
