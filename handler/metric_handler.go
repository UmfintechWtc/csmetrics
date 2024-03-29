package handler

import (
	"collect-metrics/client/cli"
	p "collect-metrics/client/prometheus"
	"collect-metrics/collector"
	"collect-metrics/common"
	config "collect-metrics/module"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	internalMetrics sync.Once
)

type PrometheusHandler struct {
	MetricsType p.PrometheusMetricsType
	Collect     collector.CollectorValues
	Config      *config.CollectMetricsConfiguration
	Cli         cli.ShellCli
	PromOpts    promhttp.HandlerOpts
	Registry    *prometheus.Registry
}

// 初始化注册表
func (p *PrometheusHandler) NewRegistry(registry *prometheus.Registry, mode string) *prometheus.Registry {
	internalMetrics.Do(func() {
		if mode == common.RUN_WITH_DEBUG {
			// 当为debug 模式时，开启内置Go 运行时相关指标
			registry.MustRegister(
				prometheus.NewGoCollector(),
			)
			// 当为debug 模式时，开启内置当前进程相关指标
			registry.MustRegister(
				prometheus.NewProcessCollector(
					prometheus.ProcessCollectorOpts{},
				),
			)
		}
	})
	return registry
}

// NewPrometheusHandler 用于构造 PrometheusHandler 实例
func NewPrometheusHandler(p p.PrometheusMetricsType, collector collector.CollectorValues, cli cli.ShellCli, config *config.CollectMetricsConfiguration) *PrometheusHandler {
	return &PrometheusHandler{
		MetricsType: p,
		Collect:     collector,
		Config:      config,
		Cli:         cli,
		Registry:    prometheus.NewRegistry(),
	}
}
