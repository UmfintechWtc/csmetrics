package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) SummaryCollector(summaryVec *prometheus.SummaryVec, labels map[string]string, value float64) {
	summaryVec.With(labels).Observe(value)
}
