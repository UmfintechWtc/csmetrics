package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) CounterCollector(counterVec *prometheus.CounterVec, labels map[string]string) error {
	counterVec.With(labels).Inc()
	return nil
}
