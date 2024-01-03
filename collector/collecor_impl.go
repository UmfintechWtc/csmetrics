package collector

import "collect-metrics/client/cli"

type CollectorValuesImpl struct {
	Cli cli.ShellCli
}

// NewCollectorValuesImpl 用于构造 CollectorValuesImpl 实例
func NewCollectorValuesImpl(cli cli.ShellCli) CollectorValues {
	return &CollectorValuesImpl{
		Cli: cli,
	}
}
