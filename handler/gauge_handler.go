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
	processGauge   *prometheus.GaugeVec
	processChannel = make(chan map[string]float64, 1)
	sessionGauge   *prometheus.GaugeVec
	netstatGauge   *prometheus.GaugeVec
	gaugeDoOnce    sync.Once
	gaugeMetrics   map[string]GagueMetrics
)

type GagueMetrics struct {
	MetricName  string
	MetricHelp  string
	MetricCmd   string
	MetricLabel []string
}

func formatMetrics() map[string]GagueMetrics {
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

func (p *PrometheusHandler) RunCli(cmd string) (*cli.GaugeValues, error) {
	r, err := p.Cli.GaugeValues(cmd)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (p *PrometheusHandler) Gauge() {
	// gaugeChannel := make(chan map[string]float64)
	gaugeDoOnce.Do(func() {
		for k, v := range formatMetrics() {
			switch k {
			case "process":
				processGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(processGauge)
				cmdRes, err := p.RunCli(v.MetricCmd)
				if err != nil {
					fmt.Println(err)
				}
				p.Collect.GaugeCollector(processGauge, cmdRes.CmdRes)
			case "netstat":
				netstatGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(netstatGauge)
				cmdRes, err := p.RunCli(v.MetricCmd)
				if err != nil {
					fmt.Println(err)
				}
				p.Collect.GaugeCollector(netstatGauge, cmdRes.CmdRes)
			case "session":
				sessionGauge = p.MetricsType.CreateGauge(v.MetricName, v.MetricHelp, v.MetricLabel)
				p.Registry.MustRegister(sessionGauge)
				cmdRes, err := p.RunCli(v.MetricCmd)
				if err != nil {
					fmt.Println(err)
				}
				p.Collect.GaugeCollector(sessionGauge, cmdRes.CmdRes)
			}
		}
	})

	go func() {
		timeTicket := time.NewTicker(10 * time.Second)
		for {
			fmt.Println("TianCiwang -- ", time.Now())
			select {
			case <-timeTicket.C:
				for k, v := range formatMetrics() {
					switch k {
					case "process":
						cmdRes, err := p.RunCli(v.MetricCmd)
						if err != nil {
							fmt.Println(err)
						}
						p.Collect.GaugeCollector(processGauge, cmdRes.CmdRes)
					case "netstat":
						cmdRes, err := p.RunCli(v.MetricCmd)
						if err != nil {
							fmt.Println(err)
						}
						p.Collect.GaugeCollector(netstatGauge, cmdRes.CmdRes)
					case "session":
						cmdRes, err := p.RunCli(v.MetricCmd)
						if err != nil {
							fmt.Println(err)
						}
						p.Collect.GaugeCollector(sessionGauge, cmdRes.CmdRes)
					}

				}
			}
		}
	}()
}
