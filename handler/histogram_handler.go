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
	histogramRequestsDelay   *prometheus.HistogramVec
	histogramBucketCondition []float64
)

func (p *PrometheusHandler) Histogram(c *gin.Context, config config.Histogram) {
	// 校验buckets边界值
	linear := config.Delay.Buckets.Linear
	slice := config.Delay.Buckets.Slice
	// 如果线性、切片都开启 或者仅开启线性，则以线性为准
	if (linear.Enabled && slice.Enabled) || (linear.Enabled && !slice.Enabled) {
		histogramBucketCondition = prometheus.LinearBuckets(
			linear.Range["start"],
			linear.Range["width"],
			int(linear.Range["count"]),
		)
		// 以切片为准
	} else if !linear.Enabled && slice.Enabled {
		histogramBucketCondition = slice.Range
		// 如果线性、列表都关闭，则返回异常
	} else if !linear.Enabled && !slice.Enabled {
		c.JSON(
			http.StatusBadRequest,
			common.NewErrorResponse(
				common.HISTOGRAM_BUCKET_ERROR,
				errors.New("Invalid bucket, need to enable linear or slice configuration..."),
			),
		)
		return
	}
	histogramMetricOnce.Do(func() {
		histogramRequestsDelay = p.MetricsType.CreateHistogram(
			"requests_delay_with_histogram",
			"Total number of HTTP requests delay with histogram",
			histogramBucketCondition,
			[]string{"code"},
		)
		p.Registry.MustRegister(histogramRequestsDelay)
	})
	p.Collect.HistogramCollector(histogramRequestsDelay, strconv.Itoa(c.Writer.Status()), common.RandomInt())
}
