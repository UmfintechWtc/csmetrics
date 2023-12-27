package main

import (
	"collect-metrics/common"
	"collect-metrics/config"
	_ "collect-metrics/docs"
	"collect-metrics/handler"
	"fmt"
	"log"
	"net/http"

	"collect-metrics/client/cli"
	p "collect-metrics/client/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetRouter(prom *handler.PrometheusHandler) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.WrapH(promhttp.Handler())) // promhttp.HandlerFor(
	// TODO: not working with NewGoCollector
	// prometheus.DefaultGatherer,
	// promhttp.HandlerOpts{
	// 	EnableOpenMetrics: false,
	// },
	// )
	r.GET("/gmetrics", prom.Gauge)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("No such route: %v", c.Request.URL.Path),
		})
	})
	return r
}

func main() {
	newPrometheus := p.NewMetricsImpl()
	newCli := cli.NewCliImpl()
	config, _ := config.LoadInternalConfig(common.COLLECT_METRICS_CONFIG_PATH)
	a, b := newCli.Netstat(config.Metrics.TCP.VerifyType)
	fmt.Println(a, b)
	r := SetRouter(handler.NewPrometheusHandler(newPrometheus))
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(
			fmt.Sprintf("Starting server failed with %s", err),
		)
	}
}
