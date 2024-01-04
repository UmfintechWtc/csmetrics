package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) CounterCollector(counterVec *prometheus.CounterVec, labels map[string]string) {
	counterVec.With(labels).Inc()
}
