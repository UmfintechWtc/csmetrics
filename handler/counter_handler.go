package handler

import (
	"collect-metrics/common"
	config "collect-metrics/module"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterMetricOnce sync.Once
	counterRegistry   *prometheus.Registry
	getRequestsCount  *prometheus.CounterVec
	counterCode       string
)

func (p *PrometheusHandler) Counter(mode string, c *gin.Context) {
	// 初始化配置
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
	counterMetricOnce.Do(func() {
		if mode == common.RUN_WITH_DEBUG {
			// 当为debug 模式时，开启内置Go 运行时相关指标
			p.CounterRegistry.MustRegister(
				prometheus.NewGoCollector(),
			)
			// 当为debug 模式时，开启内置当前进程相关指标
			p.CounterRegistry.MustRegister(
				prometheus.NewProcessCollector(
					prometheus.ProcessCollectorOpts{},
				),
			)
			counterRegistry = p.CounterRegistry
		} else {
			counterRegistry = p.AllRegistry
		}
		getRequestsCount = p.PromService.CreateCounter(
			config.Metrics.Counter.Request.MetricName,
			config.Metrics.Counter.Request.MetricHelp,
			common.COUNTER_REQUESTS_METRICS_LABELS,
		)
		counterRegistry.MustRegister(getRequestsCount)

	})

	if mode == common.RUN_WITH_DEBUG {
		num1 := common.RandomInt()
		if num1 <= 20 {
			counterCode = "100"
		} else if 20 < num1 && num1 <= 40 {
			counterCode = "200"
		} else if 40 < num1 && num1 <= 60 {
			counterCode = "300"
		} else if 60 < num1 && num1 <= 80 {
			counterCode = "400"
		} else if 80 < num1 && num1 <= 100 {
			counterCode = "500"
		}
	} else {
		counterCode = strconv.Itoa(c.Writer.Status())
	}
	setCounterLabelsValue := map[string]string{
		common.COUNTER_REQUESTS_METRICS_LABELS[0]: c.Request.URL.Path,
		common.COUNTER_REQUESTS_METRICS_LABELS[1]: counterCode,
	}
	p.Collect.CounterCollector(getRequestsCount, setCounterLabelsValue)
}
