package metrics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMetrics(t *testing.T) {
	go Metrics.Start("test", "0.0.0.0", "9190")
	require.NotEmpty(t, NameBalance.Name())
	require.NotEmpty(t, NameAlert.Name())
	require.NotEmpty(t, NameIndex.Name())
	require.NotEmpty(t, NameFee.Name())
	require.NotEmpty(t, NameSize.Name())

	// test Counter
	Metrics.MustGetCounterVec(NameSize.Name()).WithLabelValues(NameSize.LabelAssertionSize()).Inc()
	Metrics.MustGetCounterVec(NameSize.Name()).WithLabelValues(NameSize.LabelRollupSize()).Inc()

	// test Gauge
	Metrics.MustGetGaugeVec(NameAlert.Name()).WithLabelValues(NameAlert.LabelAlertChallengeStart()).Set(1)
	Metrics.MustGetGaugeVec(NameAlert.Name()).WithLabelValues(NameAlert.LabelAlertChallengeEnd()).Set(1)

	// test Histogram
	Metrics.MustGetHistogramVec(NameIndex.Name()).WithLabelValues(NameIndex.LabelAssertionIndex()).Observe(1)
	Metrics.MustGetHistogramVec(NameIndex.Name()).WithLabelValues(NameIndex.LabelBatchIndex()).Observe(1)
	Metrics.MustGetHistogramVec(NameIndex.Name()).WithLabelValues(NameIndex.LabelVerifiedIndex()).Observe(1)

	// test counter
	Metrics.MustGetSummaryVec(NameBalance.Name()).WithLabelValues(NameBalance.LabelValidatorBalance()).Observe(1)
	Metrics.MustGetSummaryVec(NameBalance.Name()).WithLabelValues(NameBalance.LabelProposerBalance()).Observe(1)

	t.Logf("wait")
	c := make(chan struct{}, 0)
	<-c
}
