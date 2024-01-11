package handler

import (
	"collect-metrics/client/cli"
	"collect-metrics/common"
	"collect-metrics/log"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	processGauge  *prometheus.GaugeVec
	sessionGauge  *prometheus.GaugeVec
	netstatGauge  *prometheus.GaugeVec
	gaugeDoOnce   sync.Once
	gaugeMetrics  map[string]GagueMetrics
	formatMetrics = formatMetricsFunc()
)

type GagueMetrics struct {
	MetricName  string
	MetricHelp  string
	MetricCmd   string
	MetricLabel []string
}

func formatMetricsFunc() map[string]GagueMetrics {
	metrics := make(map[string]GagueMetrics)
	for k, v := range common.GaugeMetrics {
		metric := GagueMetrics{}
		for m, n := range v {
			switch m {
			case "name":
				metric.MetricName = n.(string)
			case "help":
				metric.MetricHelp = n.(string)
			case "cmd":
				metric.MetricCmd = n.(string)
			case "labels":
				metric.MetricLabel = common.ChangeInterfaceToSlice(n)
			}
		}
		metrics[k] = metric
	}
	return metrics
}

// 执行Shell Cli，获取结果
func (p *PrometheusHandler) RunCli(metricType string, ch chan *cli.GaugeValues) error {
	cmd := formatMetrics[metricType].MetricCmd
	r, err := p.Cli.GaugeValues(cmd)
	if err != nil {
		return err
	}
	ch <- r
	return nil
}

// 定时处理数据
func (p *PrometheusHandler) BackGroundTask(k string, ch chan *cli.GaugeValues, gauge *prometheus.GaugeVec, cycle time.Duration) error {
	// 定时写入数据
	go func() {
		timeTicker := time.NewTicker(cycle)
		for {
			select {
			case <-timeTicker.C:
				if err := p.RunCli(k, ch); err == nil {
					fmt.Println(log.Red("error running"))
					return
				}
			}
		}
	}()
	// 实时接收数据
	go func() {
		for {
			select {
			case cmd := <-ch:
				p.Collect.GaugeCollector(gauge, cmd.CmdRes, k)
			}
		}
	}()
	return nil
}

func (p *PrometheusHandler) RunGlobalCycle() bool {
	if p.Config.Server.GlobalPeriodSeconds != nil {
		return true
	}
	return false
}

func (p *PrometheusHandler) Gauge() {
	gaugeDoOnce.Do(func() {
		for k, v := range formatMetrics {
			switch k {
			case "process":
				// 执行周期策略
				cycle := p.Config.Metrics.Gauge.PS.PeriodSeconds
				if cycle != nil {
					cycle = p.Config.Metrics.Gauge.PS.PeriodSeconds
				} else {
					if p.RunGlobalCycle() {
						cycle = p.Config.Server.GlobalPeriodSeconds
					} else {
						cycle = &common.RUN_COMMON_CYCLE
					}
				}
				processChannel := make(chan *cli.GaugeValues, 1)
				processGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(processGauge)
				// 定义一个局部变量，因为定义的 goroutine 调用变量是外部共享的
				localK := k
				// 程序启动时先加载一次数据
				p.RunCli(localK, processChannel)
				p.BackGroundTask(k, processChannel, processGauge, *cycle)
			case "netstat":
				cycle := p.Config.Metrics.Gauge.TCP.PeriodSeconds
				if cycle != nil {
					cycle = p.Config.Metrics.Gauge.TCP.PeriodSeconds
				} else {
					if p.RunGlobalCycle() {
						cycle = p.Config.Server.GlobalPeriodSeconds
					} else {
						cycle = &common.RUN_COMMON_CYCLE
					}
				}
				netstatChannel := make(chan *cli.GaugeValues, 1)
				netstatGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(netstatGauge)
				localK := k
				p.RunCli(localK, netstatChannel)
				p.BackGroundTask(k, netstatChannel, netstatGauge, *cycle)
			case "session":
				cycle := p.Config.Metrics.Gauge.Session.PeriodSeconds
				if cycle != nil {
					cycle = p.Config.Metrics.Gauge.Session.PeriodSeconds
				} else {
					if p.RunGlobalCycle() {
						cycle = p.Config.Server.GlobalPeriodSeconds
					} else {
						cycle = &common.RUN_COMMON_CYCLE
					}
				}
				sessionChannel := make(chan *cli.GaugeValues, 1)
				sessionGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(sessionGauge)
				localK := k
				p.RunCli(localK, sessionChannel)
				p.BackGroundTask(k, sessionChannel, sessionGauge, *cycle)
			}

		}
	})
}
