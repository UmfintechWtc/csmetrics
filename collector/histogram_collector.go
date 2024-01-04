package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) HistogramCollector(counterVec *prometheus.CounterVec, labels map[string]string) {
	counterVec.With(labels).Inc()
}
