package metrics

import (
	"context"
	"fmt"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

type Manager struct {
	promauto.Factory
	ctx           context.Context
	cancel        context.CancelFunc
	mutex         sync.Mutex
	handler       http.Handler
	gaugeVecs     map[string]*prometheus.GaugeVec
	counterVecs   map[string]*prometheus.CounterVec
	summaryVecs   map[string]*prometheus.SummaryVec
	histogramVecs map[string]*prometheus.HistogramVec
	// append types after here
}

func NewMetricsManager(ctx context.Context) *Manager {
	reg := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	factory := promauto.With(reg)
	ctx, cancel := context.WithCancel(ctx)
	return &Manager{
		ctx:           ctx,
		cancel:        cancel,
		handler:       handler,
		Factory:       factory,
		gaugeVecs:     make(map[string]*prometheus.GaugeVec, 0),
		counterVecs:   make(map[string]*prometheus.CounterVec, 0),
		summaryVecs:   make(map[string]*prometheus.SummaryVec, 0),
		histogramVecs: make(map[string]*prometheus.HistogramVec, 0),
	}
}

func (m *Manager) SetGaugeVec(name string, vec *prometheus.GaugeVec) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.gaugeVecs[name] = vec
}

func (m *Manager) setDefaultGaugeVec(name string) *prometheus.GaugeVec {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.gaugeVecs[name] = m.Factory.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      name,
			Help:      "default gauge vec",
			Subsystem: "default",
		},
		[]string{name},
	)
	return m.gaugeVecs[name]
}

func (m *Manager) MustGetGaugeVec(name string) *prometheus.GaugeVec {
	vec, ok := m.gaugeVecs[name]
	if !ok {
		return m.setDefaultGaugeVec(name)
	}
	return vec
}

func (m *Manager) SetCounterVec(name string, vec *prometheus.CounterVec) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counterVecs[name] = vec
}

func (m *Manager) setDefaultCounterVec(name string) *prometheus.CounterVec {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counterVecs[name] = m.Factory.NewCounterVec(
		prometheus.CounterOpts{
			Name:      name,
			Help:      "default counter vec",
			Subsystem: "default",
		},
		[]string{name},
	)
	return m.counterVecs[name]
}

func (m *Manager) MustGetCounterVec(name string) *prometheus.CounterVec {
	vec, ok := m.counterVecs[name]
	if !ok {
		return m.setDefaultCounterVec(name)
	}
	return vec
}

func (m *Manager) SetSummaryVec(name string, vec *prometheus.SummaryVec) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.summaryVecs[name] = vec
}

func (m *Manager) setDefaultSummaryVec(name string) *prometheus.SummaryVec {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.summaryVecs[name] = m.Factory.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:      name,
			Help:      "default summary vec",
			Subsystem: "default",
		},
		[]string{name},
	)
	return m.summaryVecs[name]
}

func (m *Manager) MustGetSummaryVec(name string) *prometheus.SummaryVec {
	vec, ok := m.summaryVecs[name]
	if !ok {
		return m.setDefaultSummaryVec(name)
	}
	return vec
}

func (m *Manager) SetHistogramVec(name string, vec *prometheus.HistogramVec) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.histogramVecs[name] = vec
}

func (m *Manager) setDefaultHistogramVec(name string) *prometheus.HistogramVec {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.histogramVecs[name] = m.Factory.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      name,
			Help:      "default histogram vec",
			Subsystem: "default",
		},
		[]string{name},
	)
	return m.histogramVecs[name]
}

func (m *Manager) MustGetHistogramVec(name string) *prometheus.HistogramVec {
	vec, ok := m.histogramVecs[name]
	if !ok {
		return m.setDefaultHistogramVec(name)
	}
	return vec
}

func (m *Manager) Start(module string, hostname string, port string) error {
	log.Info(fmt.Sprintf("%s metrics started", module))
	metricsAddr := fmt.Sprintf("%s:%s", hostname, port)
	go func() {
		http.Handle("/metrics", m.handler)
		_ = http.ListenAndServe(metricsAddr, m.handler)
	}()
	return nil
}

func (m *Manager) Stop(module string) error {
	log.Info(fmt.Sprintf("%s metrics started", module))
	m.cancel()
	return nil
}
