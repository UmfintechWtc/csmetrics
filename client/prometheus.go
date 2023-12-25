package client

import (
	"github.com/prometheus/client_golang/prometheus"
)

// type MetricsType interface {
// 	Gauge(labels []string) *prometheus.GaugeVec
// 	Counter(labels []string) *prometheus.CounterVec
// 	Histogram(labels []string) *prometheus.HistogramVec
// 	Summary(labels []string) *prometheus.SummaryVec
// }

func Gauge(labels []string) *prometheus.GaugeVec {
	GaugeMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "",
			Help: "",
		},
		labels,
	)
	return GaugeMetric
}

func Counter(labels []string) *prometheus.CounterVec {
	CounterMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "",
			Help: "",
		},
		labels,
	)
	return CounterMetric
}

func Histogram(labels []string) *prometheus.HistogramVec {
	HistogramMetric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "",
			Help: "",
		},
		labels,
	)
	return HistogramMetric
}

func Summary(labels []string) *prometheus.SummaryVec {
	SummaryMetric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "",
			Help: "",
		},
		labels,
	)
	return SummaryMetric
}
