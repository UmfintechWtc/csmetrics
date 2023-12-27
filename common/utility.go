package common

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func MatchKeyWord(str string, sliceData []string) bool {
	for _, e := range sliceData {
		if strings.Contains(str, e) {
			return true
		}
	}
	return false
}

func ExecCmd(cmd string) (string, error) {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("ExecCmd failed: [%s]", cmd))
	}
	return string(output), nil
}

func ChangeInterfaceToCustomFormat(kwargs interface{}) (string, error) {
	switch v := kwargs.(type) {
	case string:
		return v, nil
	case []string:
		fmtV := ""
		if len(v) == 1 {
			fmtV = `"(` + v[0] + `)"`
		} else if len(v) == 2 {
			fmtV = `"(` + v[0] + "|" + v[1] + `)"`

		} else {
			for i, s := range v {
				if i == 0 {
					fmtV += `"(` + s
				} else if i == len(v)-1 {
					fmtV += "|" + s + `)"`
				} else {
					fmtV += "|" + s
				}
			}
		}
		return fmtV, nil
	case map[string]string:
		fmtK := ""
		for k := range v {
			fmtK += k + "|"
		}
		return fmtK, nil
	default:
		return "", errors.New(
			fmt.Sprintf("Unknown Cli kwargs: %#v", kwargs),
		)
	}
}

func HostInfo() (string, string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", "", err
	}
	address, err := net.LookupIP(hostname)
	if err != nil {
		return "", "", err
	}
	return hostname, address[0].String(), nil
}

func RandomInt() float64 {
	rand.Seed(time.Now().UnixNano())

	// 生成一个1到100的随机数
	randomNumber := rand.Intn(100) + 1
	return float64(randomNumber)
}
