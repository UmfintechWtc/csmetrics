package handler

import (
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
)

type PrometheusHandler struct {
	PromService p.PrometheusMetricsType
	Collect     collector.CollectorValues
}

// NewPrometheusHandler 用于构造 PrometheusHandler 实例
func NewPrometheusHandler(prometheus p.PrometheusMetricsType, collector collector.CollectorValues) *PrometheusHandler {
	return &PrometheusHandler{
		PromService: prometheus,
		Collect:     collector,
	}
}
