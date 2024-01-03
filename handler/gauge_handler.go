package handler

import (
	"collect-metrics/common"
	config "collect-metrics/module"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricOnce      sync.Once
	register        *prometheus.Registry
	getProcessCount *prometheus.GaugeVec
	getSessionCount *prometheus.GaugeVec
	getTCPCount     *prometheus.GaugeVec
)

const (
	netstatCmd = "netstat -an | grep tcp | egrep -i %s | grep -v grep | awk '{print $NF}' | sort | uniq -c"
	processCmd = "ps aux | egrep -i %s | grep -v COMMAND | grep -v grep | awk '{print $1}' | sort | uniq -c"
	sessionCmd = "who | egrep -i %s | grep -v grep | awk '{print $1}' | grep -v grep | sort | uniq -c"
)

// 上报Gauge Metric数据
func (p *PrometheusHandler) Gauge(c *gin.Context) {
	// 初始化配置
	config, err := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			common.NewErrorResponse(
				common.PARSE_CONFIG_ERROR,
				err,
			),
		)
		return
	}
	// 在第一次运行的时候初始化Label及注册Metric
	metricOnce.Do(func() {
		fmt.Println("Initializing metrics...")
		getProcessCount = p.PromService.CreateGauge(config.Metrics.Gauge.PS.MetricName, config.Metrics.Gauge.PS.MetricHelp, common.GAUGE_PROCESS_METRICS_LABELS)
		// prometheus.MustRegister(getProcessCount)
		p.PromService.Register().MustRegister(getProcessCount)
		getSessionCount = p.PromService.CreateGauge(config.Metrics.Gauge.Session.MetricName, config.Metrics.Gauge.Session.MetricHelp, common.GAUGE_SESSION_METRICS_LABELS)
		p.PromService.Register().MustRegister(getSessionCount)
		getTCPCount = p.PromService.CreateGauge(config.Metrics.Gauge.TCP.MetricName, config.Metrics.Gauge.TCP.MetricHelp, common.GAUGE_NETSTAT_METRICS_LABELS)
		p.PromService.Register().MustRegister(getTCPCount)
	})

	// 上报Process Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getProcessCount, processCmd, config.Metrics.Gauge.PS.VerifyType, common.GAUGE_PROCESS_METRICS_LABELS)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.NewErrorResponse(
				common.COLLECT_PROCESS_METRICS_ERROR,
				err,
			),
		)
	}
	// 上报Session Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getSessionCount, sessionCmd, config.Metrics.Gauge.Session.VerifyType, common.GAUGE_SESSION_METRICS_LABELS)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.NewErrorResponse(
				common.COLLECT_SESSION_METRICS_ERROR,
				err,
			),
		)
	}

	// 上报TCP Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getTCPCount, netstatCmd, config.Metrics.Gauge.TCP.VerifyType, common.GAUGE_NETSTAT_METRICS_LABELS)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.NewErrorResponse(
				common.COLLECT_TCP_METRICS_ERROR,
				err,
			),
		)
	}
}
