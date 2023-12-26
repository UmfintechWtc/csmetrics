package common

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
func ConvertInterfaceToMap(data interface{}) (map[string]interface{}, bool) {
	myMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, false
	}
	return myMap, true
}

func ConvertInterfaceToList(data interface{}) ([]string, bool) {
	myList, ok := data.([]interface{})
	if !ok {
		return nil, false
	}
	var result []string
	for _, elem := range myList {
		strElem, ok := elem.(string)
		if !ok {
			return nil, false
		}
		result = append(result, strElem)
	}
	return result, true
}

func GetMetrics() {
	// 获取所有的 metric families
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		fmt.Println("Error gathering metrics:", err)
		return
	}

	// 遍历所有的 metric families
	for _, mf := range metricFamilies {
		// 输出 metric 的名称
		fmt.Println("Metric Name:", mf.GetName())

		// 遍历该 metric family 的 metric
		for _, m := range mf.GetMetric() {
			// 输出 metric 的标签值
			fmt.Print("Labels: ")
			for _, label := range m.GetLabel() {
				fmt.Printf("%s=%s ", label.GetName(), label.GetValue())
			}
			// 输出 metric 的值
			fmt.Printf("Value: %v\n", m.GetGauge().GetValue())
		}
	}
}

func MatchKeyWord(str string, sliceData []string) bool {
	for _, e := range sliceData {
		if strings.Contains(str, e) {
			return true
		}
	}
	return false
}

func CreateGaugeMetrics(metricName, serverName string, labels []string) *prometheus.GaugeVec {
	customMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricName,
			Help: fmt.Sprintf("Starx %s Server Metrics", serverName),
		},
		labels,
	)
	return customMetric
}

func ExecCmd(cmd string) (string, error) {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("ExecCmd failed: [%s]", cmd))
	}
	return string(output), nil
}

func GetHostName() string {
	hostname, _ := os.Hostname()
	return hostname
}

func CurrentTime(s time.Duration) string {
	previousTime := time.Now().Add(-s)
	formattedPreviousTime := previousTime.Format("2006-01-02 15:04:05")
	return formattedPreviousTime
}

func RandomInt() float64 {
	rand.Seed(time.Now().UnixNano())

	// 生成一个1到100的随机数
	randomNumber := rand.Intn(100) + 1
	return float64(randomNumber)
}
