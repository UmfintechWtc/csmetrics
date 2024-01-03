package collector

import (
	"collect-metrics/common"

	"github.com/prometheus/client_golang/prometheus"
)

// 实现 GaugeCollector 接口方法
func (c *CollectorValuesImpl) GaugeCollector(gaugeVec *prometheus.GaugeVec, cmdTemplate string, cmdArgs, labelNames []string) error {
	hostname, ip, err := common.HostInfo()
	if err != nil {
		return err
	}
	cmd, err := c.Cli.GaugeValues(cmdArgs, cmdTemplate)
	if err != nil {
		return err
	}
	setLabelsValue := map[string]string{
		labelNames[0]: hostname,
		labelNames[1]: ip,
	}
	if len(cmdArgs) == 0 {
		for statusType, value := range cmd.CmdRes {
			setLabelsValue[labelNames[2]] = statusType
			gaugeVec.With(setLabelsValue).Set(value)
		}
	} else {
		for _, statusType := range cmdArgs {
			setLabelsValue[labelNames[2]] = statusType
			value, ok := cmd.CmdRes[statusType]
			if ok {
				gaugeVec.With(setLabelsValue).Set(value)
			} else {
				gaugeVec.With(setLabelsValue).Set(value)
			}
		}
	}
	return nil
}
