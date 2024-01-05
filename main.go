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
	// 绑定 prom.AllRegistry，包含4种类型数据
	r.GET("/metrics", func(c *gin.Context) {
		prom.Gauge(mode, c)
		prom.Counter(mode, c)
		prom.Histogram(mode, c)
		prom.Summary(mode, c)
		handler := promhttp.HandlerFor(prom.AllRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 prom.CounterRegistry，统计路径
	r.GET("/metrics/*path", func(c *gin.Context) {
		prom.Counter(mode, c)
		handler := promhttp.HandlerFor(prom.AllRegistry, prom.PromOpts)
		if mode == common.RUN_WITH_DEBUG {
			handler = promhttp.HandlerFor(prom.CounterRegistry, prom.PromOpts)
		}
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 /gmetrics 路由，调用Gauge方法
	r.GET(common.URL_PREFIX["gauge"], func(c *gin.Context) {
		prom.Gauge(mode, c)
		handler := promhttp.HandlerFor(prom.GaugeRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 /cmetrics 路由，调用Counter方法
	r.GET(common.URL_PREFIX["counter"], func(c *gin.Context) {
		prom.Counter(mode, c)
		handler := promhttp.HandlerFor(prom.CounterRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 /hmetrics 路由，调用Histogram方法
	r.GET(common.URL_PREFIX["histogram"], func(c *gin.Context) {
		prom.Histogram(mode, c)
		handler := promhttp.HandlerFor(prom.HistogramRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 /smetrics 路由，调用Summary方法
	r.GET(common.URL_PREFIX["summary"], func(c *gin.Context) {
		prom.Summary(mode, c)
		handler := promhttp.HandlerFor(prom.SummaryRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	r.NoRoute(func(c *gin.Context) {
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
