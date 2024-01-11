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
	summaryMetricOnce      sync.Once
	summaryRequestsDelay   *prometheus.SummaryVec
	summaryBucketCondition = make(map[float64]float64)
)

func (p *PrometheusHandler) Summary(c *gin.Context, config config.Summary) {
	// 校验中位数
	if len(config.Delay.Median) == 0 {
		c.JSON(
			http.StatusBadRequest,
			common.NewErrorResponse(
				common.SUMMARY_BUCKET_ERROR,
				errors.New("Invalid median..."),
			),
		)
		return
	} else {
		for percent, median := range config.Delay.Median {
			summaryBucketCondition[percent/100] = median
		}
	}
	summaryMetricOnce.Do(func() {
		summaryRequestsDelay = p.MetricsType.CreateSummary(
			"requests_delay_with_summary",
			"Total number of HTTP requests delay with summary",
			summaryBucketCondition,
			[]string{"code"},
		)
		p.Registry.MustRegister(summaryRequestsDelay)
	})
	p.Collect.SummaryCollector(summaryRequestsDelay, strconv.Itoa(c.Writer.Status()), common.RandomInt())
}
