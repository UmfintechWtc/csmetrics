package collector

import (
	"collect-metrics/common"

	"github.com/prometheus/client_golang/prometheus"
)

// 初始化 CollectorValues 接口
type CollectorValues interface {
	// GaugeCollector 调用Cli接口，设置Label 及 Metric的Value，并将Metric的值上报Prometheus
	GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdTemplate string, cmdArgs, labelNames []string) *common.Response
	// CounterCollector 设置Metric的Value，并将Metric的值上报Prometheus
	CounterCollector(counterVec *prometheus.CounterVec, labels map[string]string) error
}
