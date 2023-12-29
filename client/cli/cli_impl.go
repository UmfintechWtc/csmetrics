package cli

import (
	"collect-metrics/common"
)

type CliImpl struct{}

func (c *CliImpl) GaugeValues(kwargs []string, cmdTemplate string) (*GaugeValues, error) {
	cliK, err := common.ChangeInterfaceToCustomFormat(kwargs)
	if err != nil {
		return nil, err
	}
	n, err := executeCommandWithFloat64(cliK, cmdTemplate)
	if err != nil {
		return nil, err
	}
	gauge.CmdRes = n
	return gauge, nil

}

func NewCliImpl() ShellCli {
	return &CliImpl{}
}
