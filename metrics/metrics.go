package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics keeps track of metrics in our app.
type Metrics interface {
	IncRequests()
	IncCacheHits()
}

// PrometheusMetric implements metrics with prometheus.
type PrometheusMetric struct {
	Requests  prometheus.Counter
	CacheHits prometheus.Counter
}

// NewPrometheusMetric creates a new prometheus metric.
func NewPrometheusMetric() *PrometheusMetric {
	return &PrometheusMetric{
		Requests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "The total number of api requests",
		}),
		CacheHits: promauto.NewCounter(prometheus.CounterOpts{
			Name: "cache_hits",
			Help: "Number of cache hits",
		}),
	}
}

// IncRequests increment total requests.
func (m *PrometheusMetric) IncRequests() {
	m.Requests.Inc()
}

// IncCacheHits increment cache hits.
func (m *PrometheusMetric) IncCacheHits() {
	m.CacheHits.Inc()
}
