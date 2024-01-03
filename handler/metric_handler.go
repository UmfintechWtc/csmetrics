package handler

import (
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusHandler struct {
	PromService       p.PrometheusMetricsType
	Collect           collector.CollectorValues
	GaugeRegistry     *prometheus.Registry
	CounterRegistry   *prometheus.Registry
	SummaryRegistry   *prometheus.Registry
	HistogramRegistry *prometheus.Registry
}

// NewPrometheusHandler 用于构造 PrometheusHandler 实例
func NewPrometheusHandler(p p.PrometheusMetricsType, collector collector.CollectorValues) *PrometheusHandler {
	return &PrometheusHandler{
		PromService: p,
		Collect:     collector,
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
