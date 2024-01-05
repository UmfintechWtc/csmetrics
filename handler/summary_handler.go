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
	summaryRegistry        *prometheus.Registry
	summaryRequestsDelay   *prometheus.SummaryVec
	summaryBucketCondition = make(map[float64]float64)
)

func (p *PrometheusHandler) Summary(mode string, c *gin.Context) {
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
	// 校验中位数
	if len(config.Metrics.Summary.Delay.Median) == 0 {
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
		for percent, median := range config.Metrics.Summary.Delay.Median {
			summaryBucketCondition[percent/100] = median
		}
	}
	summaryMetricOnce.Do(func() {
		if mode == common.RUN_WITH_DEBUG {
			// 当为debug 模式时，开启内置Go 运行时相关指标
			p.SummaryRegistry.MustRegister(
				prometheus.NewGoCollector(),
			)
			// 当为debug 模式时，开启内置当前进程相关指标
			p.SummaryRegistry.MustRegister(
				prometheus.NewProcessCollector(
					prometheus.ProcessCollectorOpts{},
				),
			)
			summaryRegistry = p.SummaryRegistry
		} else {
			summaryRegistry = p.AllRegistry
		}
		summaryRequestsDelay = p.PromService.CreateSummary(
			config.Metrics.Summary.Delay.MetricName,
			config.Metrics.Summary.Delay.MetricHelp,
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
