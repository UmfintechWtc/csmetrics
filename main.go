package main

import (
	"collect-metrics/client/cli"
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
	"collect-metrics/common"
	"collect-metrics/handler"
	config "collect-metrics/module"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetRouter(prom *handler.PrometheusHandler) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.WrapH(promhttp.HandlerFor(prom.PromService.Register(), promhttp.HandlerOpts{})))
	// r.Use(gin.WrapH(promhttp.Handler()))
	// r.Use(gin.WrapH(promhttp.Handler())) // promhttp.HandlerFor(
	// TODO: not working with NewGoCollector
	// prometheus.DefaultGatherer,
	// promhttp.HandlerOpts{
	// 	EnableOpenMetrics: false,
	// },
	// )
	r.GET("/gmetrics", prom.Gauge)
	r.Any("/gmetrics/*path", prom.Counter)
	r.GET("/cmetrics", prom.Counter)
	r.Any("/cmetrics/*path", prom.Counter)
	r.NoRoute(func(c *gin.Context) {
		c.Writer.WriteHeaderNow()
		fmt.Println(c.Writer.Status())
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("No such route: %v", c.Request.URL.Path),
		})
	})
	return r
}

func main() {
	// 初始化配置
	config, err := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	if err != nil {
		log.Panic(err)
	}
	// 创建 PrometheusMetricsType 对象
	prometheusMetricsType := p.NewMetricsImpl()
	// 创建 ShellCli 对象
	shelCli := cli.NewCliImpl()
	// 创建 CollectorValues 对象
	collectorValues := collector.NewCollectorValuesImpl(shelCli)
	r := SetRouter(
		handler.NewPrometheusHandler(
			prometheusMetricsType,
			collectorValues,
		),
	)
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: r,
	}
	err = svr.ListenAndServe()
	if err != nil {
		log.Panic(
			fmt.Sprintf("Starting server failed with %s", err),
		)
	}
}
