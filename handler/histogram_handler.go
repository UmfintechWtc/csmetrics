package handler

import (
	"collect-metrics/common"
	config "collect-metrics/module"
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	histogramMetricOnce      sync.Once
	histogramRegistry        *prometheus.Registry
	histogramRequestsDelay   *prometheus.HistogramVec
	histogramBucketCondition []float64
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
	// 校验buckets边界值
	linear := config.Metrics.Histogram.Delay.Buckets.Linear
	slice := config.Metrics.Histogram.Delay.Buckets.Slice
	// 如果线性、切片都开启 或者仅开启线性，则以线性为准
	if (linear.Enabled && slice.Enabled) || (linear.Enabled && !slice.Enabled) {
		histogramBucketCondition = prometheus.LinearBuckets(linear.Range["start"], linear.Range["width"], int(linear.Range["count"]))
		// 以切片为准
	} else if !linear.Enabled && slice.Enabled {
		histogramBucketCondition = slice.Range
		// 如果线性、列表都关闭，则返回异常
	} else if !linear.Enabled && !slice.Enabled {
		c.JSON(
			http.StatusBadRequest,
			common.NewErrorResponse(
				common.HISTOGRAM_BUCKET_ERROR,
				255,
				errors.New("Invalid bucket, need to enable linear or slice configuration..."),
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
		histogramRequestsDelay = p.PromService.CreateHistogram(
			config.Metrics.Histogram.Delay.MetricName,
			config.Metrics.Histogram.Delay.MetricHelp,
			histogramBucketCondition,
			common.HISTOGRAM_DELAY_METRICS_LABELS,
		)
		histogramRegistry.MustRegister(histogramRequestsDelay)
	})
	setHistogramLabelsValue := map[string]string{
		common.HISTOGRAM_DELAY_METRICS_LABELS[0]: strconv.Itoa(c.Writer.Status()),
	}
	p.Collect.HistogramCollector(histogramRequestsDelay, setHistogramLabelsValue, common.RandomInt())
}
