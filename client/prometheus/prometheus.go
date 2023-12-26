package prometheus

import "github.com/prometheus/client_golang/prometheus"

// 初始化prometheus 数据类型
type PrometheusMetricsType interface {
	// 初始化Gauge类型Metrics
	Gauge(opts prometheus.GaugeOpts, labels []string) *prometheus.GaugeVec
	// 初始化Counter类型Metrics
	Counter(opts prometheus.CounterOpts, labels []string) *prometheus.CounterVec
	// 初始化Histogram类型Metrics
	Histogram(opts prometheus.HistogramOpts, labels []string) *prometheus.HistogramVec
	// 初始化Summary类型Metrics
	Summary(opts prometheus.SummaryOpts, labels []string) *prometheus.SummaryVec
}
