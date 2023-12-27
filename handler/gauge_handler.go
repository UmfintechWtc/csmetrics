package handler

import (
	"collect-metrics/common"
	"collect-metrics/config"
	"fmt"
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

func processCountOpts(ps config.Process) prometheus.GaugeOpts {
	processOpts := prometheus.GaugeOpts{
		Name: ps.MetricName,
		Help: ps.MetricHelp,
	}
	return processOpts
}

func sessionCountOpts(tty config.Tty) prometheus.GaugeOpts {
	ttyOpts := prometheus.GaugeOpts{
		Name: tty.MetricName,
		Help: tty.MetricHelp,
	}
	return ttyOpts
}

func tcpCountOpts(tcp config.Netstat) prometheus.GaugeOpts {
	tcpOtps := prometheus.GaugeOpts{
		Name: tcp.MetricName,
		Help: tcp.MetricHelp,
	}
	return tcpOtps
}

func (i *PrometheusHandler) Gauge(c *gin.Context) {
	// 初始化配置
	config, err := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    common.FAILED_CODE,
			"message": fmt.Sprintf("err: %v", err),
		})
		return
	}
	// hostname && ip address
	hostname, ip, err := common.HostInfo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    common.FAILED_CODE,
			"message": fmt.Sprintf("err: %v", err),
		})
		return
	}
	// 在第一次运行的时候初始化Label及注册Metric
	metricOnce.Do(func() {
		getProcessCount = i.service.CreateGauge(processCountOpts(config.Metrics.PS), common.PROCESS_METRICS_LABELS)
		prometheus.MustRegister(getProcessCount)
		getSessionCount = i.service.CreateGauge(sessionCountOpts(config.Metrics.Session), common.SESSION_METRICS_LABELS)
		prometheus.MustRegister(getSessionCount)
		getTCPCount = i.service.CreateGauge(tcpCountOpts(config.Metrics.TCP), common.NETSTAT_METRICS_LABELS)
		prometheus.MustRegister(getTCPCount)
	})
	// 设置Process Label的Value及Metric的值
	for _, user := range config.Metrics.PS.VerifyType {
		labels := map[string]string{
			common.PROCESS_METRICS_LABELS[0]: hostname,
			common.PROCESS_METRICS_LABELS[1]: ip,
			common.PROCESS_METRICS_LABELS[2]: user,
		}
		i.service.SetGaugeValues(getProcessCount, labels, common.RandomInt())
	}
	// 设置Session Label的Value及Metric的值
	for _, user := range config.Metrics.Session.VerifyType {
		labels := map[string]string{
			common.SESSION_METRICS_LABELS[0]: hostname,
			common.SESSION_METRICS_LABELS[1]: ip,
			common.SESSION_METRICS_LABELS[2]: user,
		}
		i.service.SetGaugeValues(getSessionCount, labels, common.RandomInt())
	}
	// 设置TCP Label的Value及Metric的值
	for _, state := range config.Metrics.TCP.VerifyType {
		labels := map[string]string{
			common.NETSTAT_METRICS_LABELS[0]: hostname,
			common.NETSTAT_METRICS_LABELS[1]: ip,
			common.NETSTAT_METRICS_LABELS[2]: state,
		}
		i.service.SetGaugeValues(getTCPCount, labels, common.RandomInt())
	}

}
