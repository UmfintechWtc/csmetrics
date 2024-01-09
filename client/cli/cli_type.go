package cli

import (
	"errors"
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

func executeCommandWithFloat64(cmd string) (map[string]float64, error) {
	resultStatus := make(map[string]float64)
	subCommands := strings.Split(cmd, " | ")
	for _, subCmd := range subCommands {
		cmd := exec.Command("bash", "-c", subCmd)
		// 校验上一个命令是否有输出结果，若有则作为下一个命令的输入
		if len(cmdOutput) > 0 {
			cmd.Stdin = strings.NewReader(string(cmdOutput))
		}
		// 获取subCmd标准输出
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		// 获取subCmd标准错误输出
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, err
		}
		// 执行subCmd
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		// 读取标准输出
		cmdOutput, err = ioutil.ReadAll(stdout)
		if err != nil {
			return nil, err
		}
		// 读取标准错误输出
		errOutput, err := ioutil.ReadAll(stderr)
		if err != nil {
			return nil, err
		}
		// 等待命令执行完成
		if err := cmd.Wait(); err != nil {
			return nil, errors.New(string(errOutput))
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
