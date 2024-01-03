package prometheus

import "github.com/prometheus/client_golang/prometheus"

// 初始化 Prometheus 数据类型接口
type PrometheusMetricsType interface {
	// Register 创建自定义注册表
	Register() *prometheus.Registry
	// CreateGauge 创建Guage Metric类型方法
	CreateGauge(metricName, metricHelp string, labels []string) *prometheus.GaugeVec
	// CreateCounter 创建Counter Metric类型方法
	CreateCounter(metricName, metricHelp string, labels []string) *prometheus.CounterVec
	// CreateHistogram 创建Histogram Metric类型方法
	CreateHistogram(metricName, metricHelp string, labels []string) *prometheus.HistogramVec
	// CreateSummary 创建Summary Metric类型方法
	CreateSummary(metricName, metricHelp string, labels []string) *prometheus.SummaryVec
}
