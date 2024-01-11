package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	aleradyLabelsValue = map[string][]string{
		"netstat": {},
		"process": {},
		"session": {},
	}
)

func extractLabels(cmdRes map[string]float64) []string {
	labels := make([]string, 0, len(cmdRes))
	for k := range cmdRes {
		labels = append(labels, k)
	}
	return labels
}

// 实现 GaugeCollector 接口方法
func (c *CollectorValuesImpl) GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdRes map[string]float64, metricType string) error {
	for k, v := range cmdRes {
		gaugeVec.WithLabelValues(k).Set(v)
	}
	for _, v := range aleradyLabelsValue[metricType] {
		if _, ok := cmdRes[v]; !ok {
			// 若labelValue不在标签集合中，Value 设置为零值
			gaugeVec.WithLabelValues(v).Set(0)
		}
	}
	aleradyLabelsValue[metricType] = extractLabels(cmdRes)
	return nil
}
