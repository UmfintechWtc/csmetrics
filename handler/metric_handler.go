package handler

import "collect-metrics/service"

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
