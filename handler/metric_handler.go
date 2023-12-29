package handler

import (
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
)

type PrometheusHandler struct {
	PromService p.PrometheusMetricsType
	Collect     collector.CollectorValues
}

func NewPrometheusHandler(prometheus p.PrometheusMetricsType, collector collector.CollectorValues) *PrometheusHandler {
	return &PrometheusHandler{
		PromService: prometheus,
		Collect:     collector,
	}
}
