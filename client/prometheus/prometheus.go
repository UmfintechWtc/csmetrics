package prometheus

import "github.com/prometheus/client_golang/prometheus"

// 初始化prometheus 数据类型
type PrometheusMetricsType interface {
	// 初始化Gauge类型Metrics
	CreateGauge(opts prometheus.GaugeOpts, labels []string) *prometheus.GaugeVec
	SetGaugeValues(vec *prometheus.GaugeVec, labels map[string]string, value float64)
	// 初始化Counter类型Metrics
	CreateCounter(opts prometheus.CounterOpts, labels []string) *prometheus.CounterVec
	// 初始化Histogram类型Metrics
	CreateHistogram(opts prometheus.HistogramOpts, labels []string) *prometheus.HistogramVec
	// 初始化Summary类型Metrics
	CreateSummary(opts prometheus.SummaryOpts, labels []string) *prometheus.SummaryVec
}
