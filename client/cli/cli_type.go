package cli

import (
	"collect-metrics/common"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

var (
	gauge     = &GaugeValues{}
	cmdOutput []byte
)

type GaugeValues struct {
	CmdRes map[string]float64
}

func executeCommandWithFloat64(cliK string, cmdTemplate string) (map[string]float64, *common.Response) {
	resultStatus := make(map[string]float64)
	cmd := fmt.Sprintf(cmdTemplate, cliK)
	subCommands := strings.Split(cmd, " | ")
	for _, subCmd := range subCommands {
		cmd := exec.Command("bash", "-c", subCmd)
		subCode := cmd.ProcessState.ExitCode()
		// 校验上一个命令是否有输出结果，若有则作为下一个命令的输入
		if len(cmdOutput) > 0 {
			cmd.Stdin = strings.NewReader(string(cmdOutput))
		}
		// 获取subCmd标准输出
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				err,
			)
		}
		// 获取subCmd标准错误输出
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				err,
			)
		}
		// 执行subCmd
		if err := cmd.Start(); err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				err,
			)
		}
		// 读取标准输出
		cmdOutput, err = ioutil.ReadAll(stdout)
		if err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				err,
			)
		}
		// 读取标准错误输出
		errOutput, err := ioutil.ReadAll(stderr)
		if err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				err,
			)
		}
		// 等待命令执行完成
		if err := cmd.Wait(); err != nil {
			return nil, common.NewErrorResponse(
				common.EXECUTE_CLI_ERROR,
				subCode,
				errors.New(string(errOutput)),
			)
		}
	}
	fmtOutput := strings.Split(string(cmdOutput), "\n")
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
