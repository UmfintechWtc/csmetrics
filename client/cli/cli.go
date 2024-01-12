package cli

import "collect-metrics/logx"

// 初始化Shell CLI接口
type ShellCli interface {
	// 获取Gauge类型的值
	GaugeValues(cmd string) (*GaugeValues, error)
}

func TC() {
	logx.Debug("no")
}
