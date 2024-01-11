package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) CounterCollector(counterVec *prometheus.CounterVec, labels string) {
	counterVec.WithLabelValues(labels).Inc()
}
