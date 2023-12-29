package collector

import "github.com/prometheus/client_golang/prometheus"

type CollectorValues interface {
	// GaugeCollector 调用Cli接口，设置Label 及 Metric的Value，并将Metric的值上报Prometheus
	GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdTemplate string, cmdArgs, labelNames []string) error
	// CounterCollector 调用Cli接口，设置Label 及 Metric的Value，并将Metric的值上报Prometheus
	CounterCollector(gaugeVec *prometheus.CounterVec, cmdArgs, labelNames []string) error
}
