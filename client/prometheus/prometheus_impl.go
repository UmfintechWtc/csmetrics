package prometheus

import (
	"collect-metrics/client/cli"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricsImpl struct {
	Cli cli.ShellCli
}

func (m *MetricsImpl) Register() *prometheus.Registry {
	register := prometheus.NewRegistry()
	return register
}

func (m *MetricsImpl) CreateGauge(metricName, metricHelp string, labelNames []string) *prometheus.GaugeVec {
	GaugeMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricName,
			Help: metricHelp,
		},
		labelNames,
	)
	return GaugeMetric
}

func (m *MetricsImpl) CreateCounter(metricName, metricHelp string, labelNames []string) *prometheus.CounterVec {
	CounterMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName,
			Help: metricHelp,
		},
		labelNames,
	)
	return CounterMetric
}

func (m *MetricsImpl) CreateHistogram(metricName, metricHelp string, bucket []float64, labelNames []string) *prometheus.HistogramVec {
	HistogramMetric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricName,
			Help:    metricHelp,
			Buckets: bucket,
		},
		labelNames,
	)
	return HistogramMetric
}

func (m *MetricsImpl) CreateSummary(metricName, metricHelp string, labelNames []string) *prometheus.SummaryVec {
	SummaryMetric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: metricName,
			Help: metricHelp,
		},
		labelNames,
	)
	return SummaryMetric
}

func NewMetricsImpl() PrometheusMetricsType {
	return &MetricsImpl{}
}
