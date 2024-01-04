package handler

import (
	"collect-metrics/common"
	config "collect-metrics/module"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	histogramMetricOnce sync.Once
	histogramRegistry   *prometheus.Registry
	getRequestsDelay    *prometheus.HistogramVec
	histogramCode       string
)

func (p *PrometheusHandler) Histogram(mode string, c *gin.Context) {
	config, err := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.NewErrorResponse(
				common.PARSE_CONFIG_ERROR,
				255,
				err,
			),
		)
		return
	}
	histogramMetricOnce.Do(func() {
		if mode == common.RUN_WITH_DEBUG {
			// 当为debug 模式时，开启内置Go 运行时相关指标
			p.HistogramRegistry.MustRegister(
				prometheus.NewGoCollector(),
			)
			// 当为debug 模式时，开启内置当前进程相关指标
			p.HistogramRegistry.MustRegister(
				prometheus.NewProcessCollector(
					prometheus.ProcessCollectorOpts{},
				),
			)
			histogramRegistry = p.HistogramRegistry
		} else {
			histogramRegistry = p.AllRegistry
		}
		getRequestsDelay = p.PromService.CreateHistogram(
			config.Metrics.Counter.Request.MetricName,
			config.Metrics.Counter.Request.MetricHelp,
			[]float64{32},
			common.COUNTER_REQUESTS_METRICS_LABELS,
		)
		counterRegistry.MustRegister(getRequestsCount)

	})
	return
}
