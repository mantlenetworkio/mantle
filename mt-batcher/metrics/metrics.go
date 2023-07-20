package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MtBatchBase struct {
	mtBatchBalanceETH      prometheus.Gauge
	mtFeeBalanceETH        prometheus.Gauge
	batchSizeBytes         prometheus.Summary
	numTxnPerBatch         prometheus.Summary
	l2StoredBlockNumber    prometheus.Gauge
	l2ConfirmedBlockNumber prometheus.Gauge
	rollUpBatchIndex       prometheus.Gauge
	reRollUpBatchIndex     prometheus.Gauge
	eigenUserFee           prometheus.Gauge
	mtFeeNonce             prometheus.Gauge
	mtBatchNonce           prometheus.Gauge
	numEigenNode           prometheus.Gauge
	rollupTimeDuration     prometheus.Gauge
	checkerTimeDuration    prometheus.Gauge
	feeTimeDuration        prometheus.Gauge
}

func NewMtBatchBase() *MtBatchBase {
	return &MtBatchBase{
		mtBatchBalanceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "mt_batch_balance_eth",
			Help:      "ETH balance of the mt batch",
			Subsystem: "mtbatcher",
		}),
		mtFeeBalanceETH: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "mt_fee_balance_eth",
			Help:      "ETH balance of the mt fee",
			Subsystem: "mtbatcher",
		}),
		batchSizeBytes: promauto.NewSummary(prometheus.SummaryOpts{
			Name:       "batch_size_bytes",
			Help:       "Size of batches in bytes",
			Subsystem:  "mtbatcher",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		numTxnPerBatch: promauto.NewSummary(prometheus.SummaryOpts{
			Name:       "num_txn_per_batch",
			Help:       "Number of transaction in each batch",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			Subsystem:  "mtbatcher",
		}),

		l2StoredBlockNumber: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "l2_store_block_number",
			Help:      "eigen da store block number",
			Subsystem: "mtbatcher",
		}),

		l2ConfirmedBlockNumber: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "l2_confirmed_block_number",
			Help:      "eigen da confirmed block number",
			Subsystem: "mtbatcher",
		}),

		rollUpBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "rollup_batch_index",
			Help:      "Count of batches submitted",
			Subsystem: "mtbatcher",
		}),

		reRollUpBatchIndex: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "re_rollup_batch_index",
			Help:      "Count of batches re-submitted",
			Subsystem: "mtbatcher",
		}),

		eigenUserFee: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "eigen_user_fee",
			Help:      "user fee for eigen",
			Subsystem: "mtbatcher",
		}),

		mtFeeNonce: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "mt_fee_nonce",
			Help:      "nonce for mt address",
			Subsystem: "mtbatcher",
		}),

		mtBatchNonce: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "mt_batch_nonce",
			Help:      "nonce for mt batch address",
			Subsystem: "mtbatcher",
		}),

		numEigenNode: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "num_eigen_node",
			Help:      "total eigen layer nodes",
			Subsystem: "mtbatcher",
		}),
		rollupTimeDuration: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "rollup_time_duration",
			Help:      "time duration for da rollup",
			Subsystem: "mtbatcher",
		}),
		checkerTimeDuration: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "checker_time_duration",
			Help:      "time duration for data checker re rollup",
			Subsystem: "mtbatcher",
		}),
		feeTimeDuration: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "fee_time_duration",
			Help:      "time duration for fee submitter",
			Subsystem: "mtbatcher",
		}),
	}
}

func (mbb *MtBatchBase) MtBatchBalanceETH() prometheus.Gauge {
	return mbb.mtBatchBalanceETH
}

func (mbb *MtBatchBase) MtFeeBalanceETH() prometheus.Gauge {
	return mbb.mtFeeBalanceETH
}

func (mbb *MtBatchBase) BatchSizeBytes() prometheus.Summary {
	return mbb.batchSizeBytes
}

func (mbb *MtBatchBase) NumTxnPerBatch() prometheus.Summary {
	return mbb.numTxnPerBatch
}

func (mbb *MtBatchBase) L2StoredBlockNumber() prometheus.Gauge {
	return mbb.l2StoredBlockNumber
}

func (mbb *MtBatchBase) L2ConfirmedBlockNumber() prometheus.Gauge {
	return mbb.l2ConfirmedBlockNumber
}

func (mbb *MtBatchBase) RollUpBatchIndex() prometheus.Gauge {
	return mbb.rollUpBatchIndex
}

func (mbb *MtBatchBase) ReRollUpBatchIndex() prometheus.Gauge {
	return mbb.reRollUpBatchIndex
}

func (mbb *MtBatchBase) EigenUserFee() prometheus.Gauge {
	return mbb.eigenUserFee
}

func (mbb *MtBatchBase) MtFeeNonce() prometheus.Gauge {
	return mbb.mtFeeNonce
}

func (mbb *MtBatchBase) MtBatchNonce() prometheus.Gauge {
	return mbb.mtBatchNonce
}

func (mbb *MtBatchBase) NumEigenNode() prometheus.Gauge {
	return mbb.numEigenNode
}

func (mbb *MtBatchBase) RollupTimeDuration() prometheus.Gauge {
	return mbb.rollupTimeDuration
}

func (mbb *MtBatchBase) FeeTimeDuration() prometheus.Gauge {
	return mbb.feeTimeDuration
}

func (mbb *MtBatchBase) CheckerTimeDuration() prometheus.Gauge {
	return mbb.checkerTimeDuration
}
