package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) HistogramCollector(histogramVec *prometheus.HistogramVec, labels map[string]string, value float64) {
	histogramVec.With(labels).Observe(value)
}
