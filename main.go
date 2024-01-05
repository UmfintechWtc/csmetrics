package main

import (
	"collect-metrics/client/cli"
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
	"collect-metrics/common"
	"collect-metrics/handler"
	config "collect-metrics/module"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetRouter(
	prom *handler.PrometheusHandler,
	mode string,
	config config.CateGoryMetrics,
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
		prom.Gauge(mode, config.Gauge, c)
		prom.Counter(mode, config.Counter, c)
		prom.Histogram(mode, config.Histogram, c)
		prom.Summary(mode, config.Summary, c)
		handler := promhttp.HandlerFor(prom.AllRegistry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 prom.AllRegistry，仅包含Counter类型
	r.GET("/metrics/*path", func(c *gin.Context) {
		prom.Counter(mode, config.Counter, c)
		handler := promhttp.HandlerFor(prom.AllRegistry, prom.PromOpts)
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
	// 配置命令行参数
	var configFilePath string
	flag.StringVar(&configFilePath, "config", common.COLLECT_METRICS_CONFIG_PATH, "配置文件")
	flag.Parse()
	fmt.Println(configFilePath)
	// 初始化配置
	config, err := config.LoadInternalConfig(configFilePath)
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
		config.Metrics,
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
