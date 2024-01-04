package cli

import (
	"collect-metrics/common"
)

type CliImpl struct{}

func (c *CliImpl) GaugeValues(kwargs []string, cmdTemplate string) (*GaugeValues, *common.Response) {
	cliK, err := common.ChangeInterfaceToCustomFormat(kwargs)
	if err != nil {
		return nil, common.NewErrorResponse(
			common.FORMAT_CLI_QUERY_ERROR,
			255,
			err,
		)
	}
	n, rsp := executeCommandWithFloat64(cliK, cmdTemplate)
	if rsp != nil {
		return nil, rsp
	}
	gauge.CmdRes = n
	return gauge, nil

}

func NewCliImpl() ShellCli {
	return &CliImpl{}
}
