package metrics

import "github.com/prometheus/client_golang/prometheus"

type MtBatchMetrics interface {
	MtFeeBalanceETH() prometheus.Gauge

	MtBatchBalanceETH() prometheus.Gauge

	BatchSizeBytes() prometheus.Summary

	NumTxnPerBatch() prometheus.Summary

	L2StoredBlockNumber() prometheus.Gauge

	L2ConfirmedBlockNumber() prometheus.Gauge

	RollUpBatchIndex() prometheus.Gauge

	ReRollUpBatchIndex() prometheus.Gauge

	EigenUserFee() prometheus.Gauge

	MtFeeNonce() prometheus.Gauge

	MtBatchNonce() prometheus.Gauge

	NumEigenNode() prometheus.Gauge

	RollupTimeDuration() prometheus.Gauge

	FeeTimeDuration() prometheus.Gauge

	CheckerTimeDuration() prometheus.Gauge
}
