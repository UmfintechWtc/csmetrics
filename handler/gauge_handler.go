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
	gaugeMetricOnce sync.Once
	getProcessCount *prometheus.GaugeVec
	getSessionCount *prometheus.GaugeVec
	getTCPCount     *prometheus.GaugeVec
)

const (
	netstatCmd = "netstat -an | grep tcp | egrep -i %s | grep -v grep | awk '{print $NF}' | sort | uniq -c"
	processCmd = "ps aux | egrep -i %s | grep -v COMMAND | grep -v grep | awk '{print $1}' | sort | uniq -c"
	sessionCmd = "who | egrep -i %s | awk '{print $1}' | sort | uniq -c"
)

// 上报Gauge Metric数据
func (p *PrometheusHandler) Gauge(mode string, config config.Gauge, c *gin.Context) {
	// 初始化配置
	// 仅在第一次运行的时候初始化Label及注册Metric
	gaugeMetricOnce.Do(func() {
		gaugeRegistry := p.Registry(p.AllRegistry, mode)
		getProcessCount = p.PromService.CreateGauge(
			config.PS.MetricName,
			config.PS.MetricHelp,
			common.GAUGE_PROCESS_METRICS_LABELS,
		)
		gaugeRegistry.MustRegister(getProcessCount)
		getSessionCount = p.PromService.CreateGauge(
			config.Session.MetricName,
			config.Session.MetricHelp,
			common.GAUGE_SESSION_METRICS_LABELS,
		)
		gaugeRegistry.MustRegister(getSessionCount)
		getTCPCount = p.PromService.CreateGauge(
			config.TCP.MetricName,
			config.TCP.MetricHelp,
			common.GAUGE_NETSTAT_METRICS_LABELS,
		)
		gaugeRegistry.MustRegister(getTCPCount)
	})
	// 上报Process Label的Value及Metric的值
	rsp := p.Collect.GaugeCollector(
		getProcessCount, processCmd,
		config.PS.VerifyType,
		common.GAUGE_PROCESS_METRICS_LABELS,
	)
	if rsp != nil {
		c.JSON(
			http.StatusInternalServerError,
			rsp,
		)
		return
	}
	// 上报Session Label的Value及Metric的值
	rsp = p.Collect.GaugeCollector(
		getSessionCount,
		sessionCmd,
		config.Session.VerifyType,
		common.GAUGE_SESSION_METRICS_LABELS,
	)
	if rsp != nil {
		c.JSON(
			http.StatusInternalServerError,
			rsp,
		)
		return
	}

	// 上报TCP Label的Value及Metric的值
	rsp = p.Collect.GaugeCollector(
		getTCPCount,
		netstatCmd,
		config.TCP.VerifyType,
		common.GAUGE_NETSTAT_METRICS_LABELS,
	)
	if rsp != nil {
		c.JSON(
			http.StatusInternalServerError,
			rsp,
		)
		return
	}
}
