package signer

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Metrics contains metrics exposed by this package.
type Metrics struct {
	AskChannelCount  metrics.Gauge
	SignChannelCount metrics.Gauge
}

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {

	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}

	var ask = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "signer",
			Name:      "ask_channel_count",
			Help:      "ask channel backlog number",
		}, labels).With(labelsAndValues...)

	var sign = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "signer",
			Name:      "sign_channel_count",
			Help:      "sign channel backlog number",
		}, labels).With(labelsAndValues...)

	return &Metrics{
		AskChannelCount:  ask,
		SignChannelCount: sign,
	}

}
