package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsImpl struct{}

func (m *MetricsImpl) CreateGauge(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	GaugeMetric := prometheus.NewGaugeVec(
		opts,
		labelNames,
	)
	return GaugeMetric
}

func (i *MetricsImpl) SetGaugeValues(vec *prometheus.GaugeVec, labels map[string]string, value float64) {
	vec.With(labels).Set(value)
}

func (m *MetricsImpl) CreateCounter(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	CounterMetric := prometheus.NewCounterVec(
		opts,
		labelNames,
	)
	return CounterMetric
}

func (m *MetricsImpl) CreateHistogram(opts prometheus.HistogramOpts, labelNames []string) *prometheus.HistogramVec {
	HistogramMetric := prometheus.NewHistogramVec(
		opts,
		labelNames,
	)
	return HistogramMetric
}

func (m *MetricsImpl) CreateSummary(opts prometheus.SummaryOpts, labelNames []string) *prometheus.SummaryVec {
	SummaryMetric := prometheus.NewSummaryVec(
		opts,
		labelNames,
	)
	return SummaryMetric
}

func NewMetricsImpl() PrometheusMetricsType {
	return &MetricsImpl{}
}
