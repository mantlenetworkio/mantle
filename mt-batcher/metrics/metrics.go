package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/big"
	"net"
	"net/http"
	"strconv"
)

const metricsNamespace = "mt-batcher"

type MetricsConfig struct {
	Hostname string
	Port     uint64
}

type Metrics struct {
	Cfg                    *MetricsConfig
	L2StoredBlockNumber    *prometheus.GaugeVec
	L2ConfirmedBlockNumber *prometheus.GaugeVec
}

func NewMetrics(cfg *MetricsConfig) *Metrics {
	return &Metrics{
		Cfg: cfg,
		L2StoredBlockNumber: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "l1_store_block_number",
			Help:      "The height of eigen da store data l2 block",
			Namespace: metricsNamespace,
		}, []string{
			"eigen",
		}),
		L2ConfirmedBlockNumber: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "l2_confirm_block_number",
			Help:      "The height of eigen da confirmed data l2 block",
			Namespace: metricsNamespace,
		}, []string{
			"eigen",
		}),
	}
}

func (m *Metrics) SetL2StoredBlockNumber(height *big.Int) {
	m.L2StoredBlockNumber.WithLabelValues("l2").Set(float64(height.Int64()))
}

func (m *Metrics) SetL2ConfirmedBlockNumber(height *big.Int) {
	m.L2ConfirmedBlockNumber.WithLabelValues("l2").Set(float64(height.Int64()))
}

func (m *Metrics) Start() (*http.Server, error) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := new(http.Server)
	srv.Addr = net.JoinHostPort(m.Cfg.Hostname, strconv.FormatUint(m.Cfg.Port, 10))
	srv.Handler = mux
	err := srv.ListenAndServe()
	return srv, err
}
