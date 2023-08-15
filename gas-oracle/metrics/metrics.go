package metrics

import (
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	GasOracleStats struct {
		// metrics for L2 gas price
		TxSendCounter           metrics.Counter
		TxNotSignificantCounter metrics.Counter
		L2GasPriceGauge         metrics.Gauge
		TxConfTimer             metrics.Timer
		TxSendTimer             metrics.Timer

		// metrics for L1 base fee, L1 bas price, da fee
		// TokenRatioGauge token_ratio = eth_price / mnt_price
		TokenRatioGauge metrics.GaugeFloat64
		// L1BaseFeeGauge (l1_base_fee + l1_priority_fee) * token_ratio
		L1BaseFeeGauge metrics.Gauge
		// FeeScalarGauge value to scale the fee up by
		FeeScalarGauge metrics.Gauge
		// DaFeeGauge da_fee shows da gas price
		DaFeeGauge metrics.Gauge
		// OverHeadGauge over_head, amortized cost of batch submission per transaction
		OverHeadGauge metrics.Gauge
		// OverHeadGauge over_head, amortized cost of batch submission per transaction
		OverHeadUpdateGauge metrics.Gauge
		// L1GasPriceGauge l1_base_fee + l1_priority_fee
		L1GasPriceGauge metrics.Gauge

		// metrics for gas oracle version
		// PublishVersionGauge publish_version
		PublishVersionGauge metrics.Gauge
	}
)

func InitAndRegisterStats(r metrics.Registry) {
	metrics.Enabled = true

	// stats for L2 gas price
	GasOracleStats.TxSendCounter = metrics.NewRegisteredCounter("tx/send", r)
	GasOracleStats.TxNotSignificantCounter = metrics.NewRegisteredCounter("tx/not_significant", r)
	GasOracleStats.L2GasPriceGauge = metrics.NewRegisteredGauge("l2_gas_price", r)
	GasOracleStats.TxConfTimer = metrics.NewRegisteredTimer("tx/confirmed", r)
	GasOracleStats.TxSendTimer = metrics.NewRegisteredTimer("tx/send", r)

	// stats for L1 base fee, L1 bas price, da fee
	GasOracleStats.TokenRatioGauge = metrics.NewRegisteredGaugeFloat64("token_ratio", r)
	GasOracleStats.L1BaseFeeGauge = metrics.NewRegisteredGauge("l1_base_fee", r)
	GasOracleStats.FeeScalarGauge = metrics.NewRegisteredGauge("fee_scalar", r)
	GasOracleStats.DaFeeGauge = metrics.NewRegisteredGauge("da_fee", r)
	GasOracleStats.OverHeadGauge = metrics.NewRegisteredGauge("over_head", r)
	GasOracleStats.OverHeadUpdateGauge = metrics.NewRegisteredGauge("over_head_update", r)
	GasOracleStats.L1GasPriceGauge = metrics.NewRegisteredGauge("l1_gas_price", r)

	// stats for gas oracle version
	GasOracleStats.PublishVersionGauge = metrics.NewRegisteredGauge("publish_version", r)
}
