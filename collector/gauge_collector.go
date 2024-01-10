package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 实现 GaugeCollector 接口方法
func (c *CollectorValuesImpl) GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdRes map[string]float64) error {
	for k, v := range cmdRes {
		gaugeVec.WithLabelValues(k).Set(v)
	}
	return nil
}
