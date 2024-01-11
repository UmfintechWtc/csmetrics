package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) SummaryCollector(summaryVec *prometheus.SummaryVec, labelsValues string, value float64) {
	summaryVec.WithLabelValues(labelsValues).Observe(value)
}
