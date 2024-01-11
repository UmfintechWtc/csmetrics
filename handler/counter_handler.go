package handler

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// import (
// 	"collect-metrics/common"
// 	config "collect-metrics/module"
// 	"strconv"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// 	"github.com/prometheus/client_golang/prometheus"
// )

var (
	counterMetricOnce sync.Once
	getRequestsCount  *prometheus.CounterVec
	counterCode       string
)

type CounterMetrics struct {
	MetricName  string
	MetricHelp  string
	MetricCmd   string
	MetricLabel []string
}

func (p *PrometheusHandler) Counter(c *gin.Context) {
	// 初始化配置
	counterMetricOnce.Do(func() {
		getRequestsCount = p.MetricsType.CreateCounter(
			"requests_url_total",
			"get requests order by url",
			[]string{"url"},
		)
		p.Registry.MustRegister(getRequestsCount)
	})
	p.Collect.CounterCollector(getRequestsCount, c.Request.URL.Path)
}
