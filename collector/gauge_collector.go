package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// 实现 GaugeCollector 接口方法
func (c *CollectorValuesImpl) GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdRes map[string]float64, cmdArgs, labelNames []string) error {
	if len(cmdArgs) == 0 {
		fmt.Println(cmdRes)
	} else {
		fmt.Println(cmdRes)
	}
	return nil
}
