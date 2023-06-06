package metics

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Metrics contains metrics exposed by this package.
type Metrics struct {
	OnlineNodesCount   metrics.Gauge
	SignCount          metrics.Gauge
	RollbackCount      metrics.Gauge
	SlashCount         metrics.Gauge
	ActiveMembersCount metrics.Gauge
	ApproveNumber      metrics.Gauge
}

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {

	var labels []string
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}

	var sign = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "sign",
			Name:      "sign failed counter",
			Help:      "sign failed ",
		}, labels).With(labelsAndValues...)
	var rollback = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "sign",
			Name:      "rollback counter",
			Help:      "rollback event ",
		}, labels).With(labelsAndValues...)
	var approve = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "sign",
			Name:      "approve counter",
			Help:      "approve number ",
		}, labels).With(labelsAndValues...)

	var slash = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "slash",
			Name:      "slash counter",
			Help:      "remark slash behavior",
		}, labels).With(labelsAndValues...)
	var online = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "node",
			Name:      "online node counter",
			Help:      "online node number ",
		}, labels).With(labelsAndValues...)
	var active = prometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "node",
			Name:      "active counter",
			Help:      "active node behavior",
		}, labels).With(labelsAndValues...)

	return &Metrics{
		OnlineNodesCount:   online,
		SignCount:          sign,
		SlashCount:         slash,
		ActiveMembersCount: active,
		RollbackCount:      rollback,
		ApproveNumber:      approve,
	}

}
