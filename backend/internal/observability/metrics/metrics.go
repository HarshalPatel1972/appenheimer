package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTP
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status"},
	)

	// Database
	DBConnectionActive = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_connections_active",
			Help: "Current active database connections",
		},
	)
	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_name"},
	)

	// Worker
	WorkerQueueDepth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "worker_queue_depth",
			Help: "Number of pending events in the outbox.",
		},
		[]string{"consumer"},
	)
	WorkerProcessingDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "worker_processing_duration_seconds",
			Help:    "Duration of event processing.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"consumer"},
	)

	// Build Info
	BuildInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "app_build_info",
			Help: "Build metadata",
		},
		[]string{"version", "commit", "environment"},
	)
)

func RegisterAll() {
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(DBConnectionActive)
	prometheus.MustRegister(DBQueryDuration)
	prometheus.MustRegister(WorkerQueueDepth)
	prometheus.MustRegister(WorkerProcessingDuration)
	prometheus.MustRegister(BuildInfo)
}
