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

func SetRouter(
	prom *handler.PrometheusHandler,
	mode string,
) *gin.Engine {
	if mode == "" {
		gin.SetMode(gin.ReleaseMode)
	} else if common.CheckKey(mode, common.RUN_MODE) {
		gin.SetMode(mode)
	} else {
		log.Panic(
			fmt.Sprintf(
				"got unknown gin.engine mode %v, expect %v",
				mode,
				common.RUN_MODE,
			),
		)
	}
	r := gin.New()
	// 绑定 Gauge Metric 路由
	r.GET("/gmetrics", gin.WrapH(
		promhttp.HandlerFor(
			prom.Gauge(
				mode,
				&gin.Context{},
			),
			prom.PromOpts,
		),
	))
	r.Any("/gmetrics/*path", prom.Counter)
	// 绑定 Counter Metric 路由
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
	// 创建 ShellCli 对象
	shellCli := cli.NewCliImpl()
	// 创建 PrometheusMetricsType 对象
	prometheusMetricsType := p.NewMetricsImpl()
	// 创建 CollectorValues 对象
	collectorValues := collector.NewCollectorValuesImpl(shellCli)
	// 创建 PrometheusHandler 对象
	newPrometheusHandler := handler.NewPrometheusHandler(
		prometheusMetricsType,
		collectorValues,
	)
	// 设置 Gin 路由
	r := SetRouter(
		newPrometheusHandler,
		config.Server.Mode,
	)
	// 启动http服务
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
