package metrics

import (
	"context"
	"github.com/mantlenetworkio/mantle/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = metrics.NewMetricsManager(context.Background())

func init() {
	initName()
	Metrics.SetSummaryVec(NameSize.Name(), Metrics.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       NameSize.Name(),
			Help:       "Size of elements",
			Subsystem:  "FraudProof",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{"LabelSize"},
	))

	Metrics.SetGaugeVec(NameBalance.Name(), Metrics.Factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      NameBalance.Name(),
			Help:      "Balance of EOAs",
			Subsystem: "FraudProof",
		},
		[]string{"LabelBalance"},
	))

	Metrics.SetGaugeVec(NameFee.Name(), Metrics.Factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      NameFee.Name(),
			Help:      "Fee used of EOAs",
			Subsystem: "FraudProof",
		},
		[]string{"LabelFee"},
	))

	Metrics.SetGaugeVec(NameIndex.Name(), Metrics.Factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      NameIndex.Name(),
			Help:      "Index of rollup",
			Subsystem: "FraudProof",
		},
		[]string{"LabelIndex"},
	))

	Metrics.SetCounterVec(NameAlert.Name(), Metrics.Factory.NewCounterVec(
		prometheus.CounterOpts{
			Name:      NameAlert.Name(),
			Help:      "Alert of challenge status",
			Subsystem: "FraudProof",
		},
		[]string{"LabelAlert"},
	))
}
