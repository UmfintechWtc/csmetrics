package handler

import (
	"collect-metrics/common"
	"collect-metrics/config"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	configOnce      sync.Once
	getProcessCount *prometheus.GaugeVec
)

func processCount(config *config.CollectMetricsConfiguration) prometheus.GaugeOpts {
	processOpts := prometheus.GaugeOpts{
		Name: config.Metrics.PS.MetricName,
		Help: fmt.Sprintf(config.Metrics.PS.MetricName + " processes"),
	}
	return processOpts
}

func (i *PrometheusHandler) Gauge(c *gin.Context) {
	res1, _ := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	fmt.Println(res1.Server.GlobalPeriodSeconds, "==========")
	res := prometheus.GaugeOpts{
		Name: "TianCiwang",
		Help: "TianCiwang",
	}
	configOnce.Do(func() {
		// only in the first run to register the metrics
		getProcessCount = i.service.Gauge(res, []string{"name", "age"})
		prometheus.MustRegister(getProcessCount)
	})
	getProcessCount.WithLabelValues("TianCiwang", "31").Set(common.RandomInt())
}
