package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func StartServer(hostname string, port uint64) {
	metricsPortStr := strconv.FormatUint(port, 10)
	metricsAddr := fmt.Sprintf("%s:%s", hostname, metricsPortStr)

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(metricsAddr, nil)
}
