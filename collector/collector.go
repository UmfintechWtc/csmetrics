package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 初始化 CollectorValues 接口
type CollectorValues interface {
	// GaugeCollector 调用Cli接口，设置Label 及 Metric的Value，并将Metric的值上报Prometheus
	GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdRes map[string]float64, metricType string) error
	// CounterCollector Metric的Value自增 +1，并将Metric的值上报Prometheus
	CounterCollector(counterVec *prometheus.CounterVec, labelsValues string)
	// HistogramCollector 设置Metric的Bucket且Value自增+1，并将Metric的值上报Prometheus
	HistogramCollector(histogramVec *prometheus.HistogramVec, labelsValues string, value float64)
	// HistogramCollector 设置Metric的Mediam且Value自增+1，并将Metric的值上报Prometheus
	SummaryCollector(summaryVec *prometheus.SummaryVec, labelsValues string, value float64)
}
