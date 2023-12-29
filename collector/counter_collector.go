package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) CounterCollector(gaugeVec *prometheus.CounterVec, cmdArgs, labelNames []string) error {
	return nil
}
