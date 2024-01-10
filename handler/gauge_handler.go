package handler

import (
	"collect-metrics/client/cli"
	"collect-metrics/common"
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

func (p *PrometheusHandler) RunCli(metricType string, ch chan *cli.GaugeValues) {
	cmd := formatMetrics[metricType].MetricCmd
	r, _ := p.Cli.GaugeValues(cmd)
	fmt.Println(metricType, " -- ", cmd, " -- ", r.CmdRes)
	ch <- r
}

func (p *PrometheusHandler) BackGroundTask(k string, ch chan *cli.GaugeValues, gauge *prometheus.GaugeVec) {
	// 定时写入数据
	go func() {
		timeTicker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-timeTicker.C:
				p.RunCli(k, ch)

			}
		}
	}()
	// 实时接收数据
	go func() {
		for {
			select {
			case cmd := <-ch:
				p.Collect.GaugeCollector(gauge, cmd.CmdRes)
			}
		}
	}()
}

func (p *PrometheusHandler) Gauge() {
	gaugeDoOnce.Do(func() {
		for k, v := range formatMetrics {
			switch k {
			case "process":
				processChannel := make(chan *cli.GaugeValues, 1)
				processGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(processGauge)
				// 定义一个局部变量，因为定义的 goroutine 调用变量是外部共享的
				localK := k
				// 程序启动时先加载一次数据
				p.RunCli(localK, processChannel)
				// 定时写入数据
				go func() {
					timeTicker := time.NewTicker(3 * time.Second)
					for {
						select {
						case <-timeTicker.C:
							p.RunCli(localK, processChannel)

						}
					}
				}()
				// 实时接收数据
				go func() {
					for {
						select {
						case cmd := <-processChannel:
							p.Collect.GaugeCollector(processGauge, cmd.CmdRes)
						}
					}
				}()
			case "netstat":
				netstatChannel := make(chan *cli.GaugeValues, 1)
				netstatGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(netstatGauge)
				localK := k
				p.RunCli(localK, netstatChannel)
				// 定时写入数据
				go func() {
					timeTicker := time.NewTicker(3 * time.Second)
					for {
						select {
						case <-timeTicker.C:
							p.RunCli(localK, netstatChannel)

						}
					}
				}()
				// 实时接收数据
				go func() {
					for {
						select {
						case cmd := <-netstatChannel:
							p.Collect.GaugeCollector(netstatGauge, cmd.CmdRes)
						}
					}
				}()
			case "session":
				sessionChannel := make(chan *cli.GaugeValues, 1)
				sessionGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(sessionGauge)
				localK := k
				p.RunCli(localK, sessionChannel)
				// 定时写入数据
				go func() {
					timeTicker := time.NewTicker(3 * time.Second)
					for {
						select {
						case <-timeTicker.C:
							p.RunCli(localK, sessionChannel)

						}
					}
				}()
				// 实时接收数据
				go func() {
					for {
						select {
						case cmd := <-sessionChannel:
							p.Collect.GaugeCollector(sessionGauge, cmd.CmdRes)
						}
					}
				}()
			}

		}
	})
}
