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

func (p *PrometheusHandler) Summary(mode string, config config.Summary, c *gin.Context) {
	// 校验中位数
	if len(config.Delay.Median) == 0 {
		c.JSON(
			http.StatusBadRequest,
			common.NewErrorResponse(
				common.SUMMARY_BUCKET_ERROR,
				255,
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
		summaryRegistry := p.Registry(p.AllRegistry, mode)
		summaryRequestsDelay = p.PromService.CreateSummary(
			config.Delay.MetricName,
			config.Delay.MetricHelp,
			summaryBucketCondition,
			common.SUMMARY_DELAY_METRICS_LABELS,
		)
		summaryRegistry.MustRegister(summaryRequestsDelay)
	})
	setSummaryLabelsValue := map[string]string{
		// common.SUMMARY_DELAY_METRICS_LABELS[0]: c.Request.URL.Path,
		common.SUMMARY_DELAY_METRICS_LABELS[0]: strconv.Itoa(c.Writer.Status()),
	}
	p.Collect.SummaryCollector(summaryRequestsDelay, setSummaryLabelsValue, common.RandomInt())
}
