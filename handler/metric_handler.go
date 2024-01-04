package handler

import (
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHandler struct {
	PromService p.PrometheusMetricsType
	PromOpts    promhttp.HandlerOpts
	Collect     collector.CollectorValues
	AllRegistry *prometheus.Registry
	// Debug 调试 Gauge
	GaugeRegistry *prometheus.Registry
	// Debug 调试 Counter
	CounterRegistry *prometheus.Registry
	// Debug 调试 Summary
	SummaryRegistry *prometheus.Registry
	// Debug 调试 Histogram
	HistogramRegistry *prometheus.Registry
}

// NewPrometheusHandler 用于构造 PrometheusHandler 实例
func NewPrometheusHandler(p p.PrometheusMetricsType, collector collector.CollectorValues) *PrometheusHandler {
	return &PrometheusHandler{
		PromService: p,
		Collect:     collector,
		AllRegistry: prometheus.NewRegistry(),
		// 初始化 Gauge Metric 注册表
		GaugeRegistry: prometheus.NewRegistry(),
		// 初始化 Counter Metric 注册表
		CounterRegistry: prometheus.NewRegistry(),
		// 初始化 Summary Metric 注册表
		SummaryRegistry: prometheus.NewRegistry(),
		// 初始化 Histogram Metric 注册表
		HistogramRegistry: prometheus.NewRegistry(),
	}
}
