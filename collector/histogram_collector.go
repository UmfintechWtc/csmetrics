package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *CollectorValuesImpl) HistogramCollector(histogramVec *prometheus.HistogramVec, labelsValues string, value float64) {
	histogramVec.WithLabelValues(labelsValues).Observe(value)
}
