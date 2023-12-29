package cli

// 初始化Shell CLI接口
type ShellCli interface {
	// 获取Gauge类型的值
	GaugeValues(kwargs []string, cmdTemplate string) (*GaugeValues, error)
}
