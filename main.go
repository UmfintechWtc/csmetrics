package main

import (
	"collect-metrics/client/cli"
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
	"collect-metrics/common"
	"collect-metrics/handler"
	"collect-metrics/logx"
	config "collect-metrics/module"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		logx.Fatalf("got unknown gin.engine mode %v, expect %v",
			mode,
			common.RUN_MODE,
		)
	}
	r := gin.New()
	// r.Use(gin.WrapH(promhttp.Handler()))
	// r.GET("/metrics", prom.Gauge)
	r.GET("/metrics", func(c *gin.Context) {
		logx.Infof("request url %s, status: %d, client: %s", c.Request.URL.Path, c.Writer.Status(), c.ClientIP())
		prom.Counter(c)
		prom.Histogram(c, config.Histogram)
		prom.Summary(c, config.Summary)
		handler := promhttp.HandlerFor(prom.Registry, prom.PromOpts)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	// 绑定 prom.AllRegistry，仅包含Counter类型
	r.GET("/metrics/*path", func(c *gin.Context) {
		prom.Counter(c)
		handler := promhttp.HandlerFor(prom.Registry, prom.PromOpts)
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
	// 命令行参数
	var configFilePath string
	var versionFlag bool
	flag.StringVar(&configFilePath, "config", common.COLLECT_METRICS_CONFIG_PATH, "配置文件")
	flag.BoolVar(&versionFlag, "version", false, "版本信息")
	flag.Parse()
	if versionFlag {
		fmt.Println(common.APP_VERSION)
		return
	}
	// 初始化配置
	config, err := config.LoadInternalConfig(configFilePath)
	if err != nil {
		log.Panic(err)
	}
	// 创建 logx 对象
	logx.SetGlobalLogger(logx.NewLogrusLogger(config.Server.LogrusConfig))
	// 创建 ShellCli 对象
	shellCli := cli.NewCliImpl()
	// 创建 PrometheusMetricsType 对象
	prometheusMetricsType := p.NewMetricsImpl()
	// 创建 CollectorValues 对象
	collectorValues := collector.NewCollectorValuesImpl()
	// 创建 PrometheusHandler 对象
	newPrometheusHandler := handler.NewPrometheusHandler(
		prometheusMetricsType,
		collectorValues,
		shellCli,
		config,
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
	go func() {
		if err = svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Fatalf("Starting server failed with %s", err)
		}
	}()
	logx.Infof("Server started successfully with %s:%d/%s", config.Server.Listen, config.Server.Port, common.URL_PREFIX["all"])
	// 初始化Gauge
	newPrometheusHandler.Gauge()

	// 如果收到指定的信号，那么关闭 HTTP 服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logx.Info("Shutdown server ...")
	timeout := time.Duration(config.Server.ShutdownTimeoutMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		logx.Fatalf(
			fmt.Sprintf("Shutdown server failure with %s", err),
		)
	}
	// 等待服务关闭
	select {
	case <-ctx.Done():
		logx.Infof(
			fmt.Sprintf("Timeout of %dms reached", config.Server.ShutdownTimeoutMs),
		)
	}
	logx.Infof("Server exited")
}
