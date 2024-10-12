package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"
)

func init() {
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(RequestCounter)
}


var RequestCounter = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_get",
        Help: "Number of GET requests made to URLs. Labels include url and response code",
    },
    []string{"url", "code"},
)

