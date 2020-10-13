package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics keeps track of metrics in our app.
type Metrics interface {
	IncCacheHits()
}

// PrometheusMetric implements metrics with prometheus.
type PrometheusMetric struct {
	CacheHits prometheus.Counter
}

// NewPrometheusMetric creates a new prometheus metric.
func NewPrometheusMetric() *PrometheusMetric {
	return &PrometheusMetric{
		CacheHits: promauto.NewCounter(prometheus.CounterOpts{
			Name: "cache_hits",
			Help: "Number of cache hits",
		}),
	}
}

// IncCacheHits increment cache hits.
func (m *PrometheusMetric) IncCacheHits() {
	m.CacheHits.Inc()
}
