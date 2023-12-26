package handler

import (
	p "collect-metrics/client/prometheus"
	"collect-metrics/service"
)

type MetricHandler struct {
	service service.MetricService
	name    string
	age     int
}

func NewMetricHandler(service service.MetricService, name string, age int) *MetricHandler {
	return &MetricHandler{
		service: service,
		name:    name,
		age:     age,
	}
}

type PrometheusHandler struct {
	service p.PrometheusMetricsType
}

func NewPrometheusHandler(service p.PrometheusMetricsType) *PrometheusHandler {
	return &PrometheusHandler{
		service: service,
	}
}
