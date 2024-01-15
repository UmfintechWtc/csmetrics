package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
//	aleradyLabelsValue = map[string][]string{
//		"netstat": {},
//		"process": {},
//		"session": {},
//	}
)

func extractLabels(cmdRes map[string]float64) []string {
	labels := make([]string, 0, len(cmdRes))
	for k := range cmdRes {
		labels = append(labels, k)
	}
	return labels
}

// 实现 GaugeCollector 接口方法
func (c *CollectorValuesImpl) GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdRes map[string]float64, metricType string) {
	for k, v := range cmdRes {
		gaugeVec.WithLabelValues(k).Set(v)
	}
	// for _, v := range aleradyLabelsValue[metricType] {
	// 	if _, ok := cmdRes[v]; !ok {
	// 		// 若labelValue不在标签集合中，Value 设置为零值
	// 		gaugeVec.WithLabelValues(v).Set(0)
	// 	}
	// }
	// aleradyLabelsValue[metricType] = extractLabels(cmdRes)

	// 避免map并发写时产生竞争，抛出concurrent map writes异常，并发读不会产生竞争
	// 这里不使用CVMap.Range是因为遍历整个sync.Map，目标是获取最后层级value
	lastLabelValues, _ := c.CVMap.Load(metricType)
	if lastLabelValues != nil {
		for _, v := range lastLabelValues.([]string) {
			if _, ok := cmdRes[v]; !ok {
				// 若labelValue不在标签集合中，Value 设置为零值
				gaugeVec.WithLabelValues(v).Set(0)
			}
		}
	}
	c.CVMap.Store(metricType, extractLabels(cmdRes))

}
