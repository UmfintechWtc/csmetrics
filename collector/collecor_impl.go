package collector

import "collect-metrics/client/cli"

type CollectorValuesImpl struct {
	Cli cli.ShellCli
}

func NewCollectorValuesImpl(cli cli.ShellCli) CollectorValues {
	return &CollectorValuesImpl{
		Cli: cli,
	}
}
