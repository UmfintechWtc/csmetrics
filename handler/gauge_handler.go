package handler

// import (
// 	"collect-metrics/common"
// 	config "collect-metrics/module"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/prometheus/client_golang/prometheus"
// )

// var (
// 	gaugeMetricOnce sync.Once
// 	getProcessCount *prometheus.GaugeVec
// 	processChan     chan map[string]float64
// 	getSessionCount *prometheus.GaugeVec
// 	ttyChan         chan map[string]float64
// 	getTCPCount     *prometheus.GaugeVec
// 	tcpChan         chan map[string]float64
// )

// // 格式化执行命令
// func fmtCmdFunc(queryCondition []string, cmdTemplate string) (string, error) {
// 	cliK, err := common.ChangeInterfaceToCustomFormat(queryCondition)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf(cmdTemplate, cliK), nil
// }

// // 执行命令获取结果
// func (p *PrometheusHandler) getCmdRes(cmdRes chan map[string]float64, cmd string, cycle time.Duration) error {
// 	ticker := time.NewTicker(10 * time.Second)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			r, err := p.Cli.GaugeValues(cmd)
// 			fmt.Println(r.CmdRes)
// 			if err != nil {
// 				return err
// 			}
// 			cmdRes <- r.CmdRes
// 		default:
// 			fmt.Println("Error getting command")
// 		}
// 	}
// }

// // 初始化命令执行
// func (p *PrometheusHandler) InitializeCmd(config config.Gauge) error {
// 	processChan = make(chan map[string]float64, 1)
// 	processCmd := "ps aux | egrep -i '(zabbix|starcross|root)' | grep -v COMMAND | grep -v grep | awk '{print $1}' | sort | uniq -c"
// 	go func() {
// 		p.getCmdRes(processChan, processCmd, *config.PS.PeriodSeconds)
// 	}()
// 	return nil
// }

// // 上报Gauge Metric数据
// func (p *PrometheusHandler) Gauge(mode string, config config.Gauge, c *gin.Context) {
// 	// 仅在第一次运行的时候初始化Label及注册Metric
// 	gaugeMetricOnce.Do(func() {
// 		gaugeRegistry := p.Registry(p.AllRegistry, mode)
// 		getProcessCount = p.PromService.CreateGauge(config.PS.MetricName, config.PS.MetricHelp, common.GAUGE_PROCESS_METRICS_LABELS)
// 		gaugeRegistry.MustRegister(getProcessCount)
// 	})
// 	// 上报Process Label的Value及Metric的值
// 	for {
// 		select {
// 		case res := <-processChan:
// 			rsp := p.Collect.GaugeCollector(getProcessCount, res, config.PS.VerifyType, common.GAUGE_PROCESS_METRICS_LABELS)
// 			if rsp != nil {
// 				c.JSON(
// 					http.StatusInternalServerError,
// 					rsp,
// 				)
// 				return
// 			}
// 		default:
// 			fmt.Println("processChan is empty")
// 		}
// 		time.Sleep(3 * time.Second)
// 	}
// }
