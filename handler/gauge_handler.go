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
	metricOnce      sync.Once
	getProcessCount *prometheus.GaugeVec
	getSessionCount *prometheus.GaugeVec
	getTCPCount     *prometheus.GaugeVec
)

const (
	netstatCmd = "netstat -an | grep tcp | egrep -i %s | grep -v grep | awk '{print $NF}' | sort | uniq -c"
	processCmd = "ps aux | egrep -i %s | grep -v COMMAND | grep -v grep | awk '{print $1}' | sort | uniq -c"
	sessionCmd = "who | egrep -i %s | grep -v grep | awk '{print $1}' | grep -v grep | sort | uniq -c"
)

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
		getProcessCount = p.PromService.CreateGauge(config.Metrics.PS.MetricName, config.Metrics.PS.MetricHelp, common.PROCESS_METRICS_LABELS)
		prometheus.MustRegister(getProcessCount)
		getSessionCount = p.PromService.CreateGauge(config.Metrics.Session.MetricName, config.Metrics.Session.MetricHelp, common.SESSION_METRICS_LABELS)
		prometheus.MustRegister(getSessionCount)
		getTCPCount = p.PromService.CreateGauge(config.Metrics.TCP.MetricName, config.Metrics.TCP.MetricHelp, common.NETSTAT_METRICS_LABELS)
		prometheus.MustRegister(getTCPCount)
	})

	// 注册Process Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getProcessCount, processCmd, config.Metrics.PS.VerifyType, common.PROCESS_METRICS_LABELS)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.NewErrorResponse(
				common.COLLECT_PROCESS_METRICS_ERROR,
				err,
			),
		)
	}
	// 注册Session Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getSessionCount, sessionCmd, config.Metrics.Session.VerifyType, common.SESSION_METRICS_LABELS)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.NewErrorResponse(
				common.COLLECT_SESSION_METRICS_ERROR,
				err,
			),
		)
	}

	// 注册TCP Label的Value及Metric的值
	err = p.Collect.GaugeCollector(getTCPCount, netstatCmd, config.Metrics.TCP.VerifyType, common.NETSTAT_METRICS_LABELS)
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
