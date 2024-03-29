package prometheus

import "github.com/prometheus/client_golang/prometheus"

// 初始化 Prometheus 数据类型接口
type PrometheusMetricsType interface {
	CreateGauge(metricName, metricHelp string, labels []string) *prometheus.GaugeVec
	// CreateCounter 创建Counter Metric类型方法
	CreateCounter(metricName, metricHelp string, labels []string) *prometheus.CounterVec
	// CreateHistogram 创建Histogram Metric类型方法
	CreateHistogram(metricName, metricHelp string, bucket []float64, labels []string) *prometheus.HistogramVec
	// CreateSummary 创建Summary Metric类型方法
	CreateSummary(metricName, metricHelp string, median map[float64]float64, labels []string) *prometheus.SummaryVec
}
