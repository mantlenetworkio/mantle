package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitAndRegisterStats(t *testing.T) {
	InitAndRegisterStats(DefaultRegistry)

	GasOracleStats.L1BaseFeeGauge.Update(100)
	l1BaseFee := GasOracleStats.L1BaseFeeGauge.Value()
	require.Equal(t, int64(100), l1BaseFee)

	GasOracleStats.TokenRatioGauge.Update(4000)
	tokenRatio := GasOracleStats.TokenRatioGauge.Value()
	require.Equal(t, int64(4000), tokenRatio)

	GasOracleStats.FeeScalarGauge.Update(1500000)
	feeScalar := GasOracleStats.FeeScalarGauge.Value()
	require.Equal(t, int64(1500000), feeScalar)

	GasOracleStats.DaFeeGauge.Update(1500000)
	daFee := GasOracleStats.DaFeeGauge.Value()
	require.Equal(t, int64(1500000), daFee)

	GasOracleStats.OverHeadGauge.Update(1500000)
	overHead := GasOracleStats.OverHeadGauge.Value()
	require.Equal(t, int64(1500000), overHead)
}
