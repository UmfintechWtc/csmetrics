package cli

import (
	"collect-metrics/common"
	"fmt"
	"strconv"
	"strings"
)

var (
	netstat = &Netstat{}
	process = &Process{}
	session = &Session{}
	gauge   = &GaugeValues{}
)

type GaugeValues struct {
	CmdRes map[string]float64
}
type Netstat struct {
	CmdRes map[string]float64
}
type Process struct {
	CmdRes map[string]float64
}

type Session struct {
	CmdRes map[string]float64
}

func executeCommandWithFloat64(cliK string, cmdTemplate string) (map[string]float64, error) {
	resultStatus := make(map[string]float64)
	cmd := fmt.Sprintf(cmdTemplate, cliK)
	output, err := common.ExecCmd(cmd)
	if err != nil {
		return nil, err
	}
	fmtOutput := strings.Split(output, "\n")
	for _, line := range fmtOutput {
		parts := strings.Split(strings.TrimLeft(line, " "), " ")
		if len(parts) != 1 {
			if floatValue, err := strconv.ParseFloat(parts[0], 64); err == nil {
				resultStatus[parts[1]] = floatValue
			}
		}
	}
	return resultStatus, nil
}
