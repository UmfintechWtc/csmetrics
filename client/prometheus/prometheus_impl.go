package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsImpl struct {
}

func (m *MetricsImpl) Gauge(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	GaugeMetric := prometheus.NewGaugeVec(
		opts,
		labelNames,
	)
	return GaugeMetric
}

func (m *MetricsImpl) Counter(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	CounterMetric := prometheus.NewCounterVec(
		opts,
		labelNames,
	)
	return CounterMetric
}

func (m *MetricsImpl) Histogram(opts prometheus.HistogramOpts, labelNames []string) *prometheus.HistogramVec {
	HistogramMetric := prometheus.NewHistogramVec(
		opts,
		labelNames,
	)
	return HistogramMetric
}

func (m *MetricsImpl) Summary(opts prometheus.SummaryOpts, labelNames []string) *prometheus.SummaryVec {
	SummaryMetric := prometheus.NewSummaryVec(
		opts,
		labelNames,
	)
	return SummaryMetric
}

func NewMetricsImpl() PrometheusMetricsType {
	return &MetricsImpl{}
}
