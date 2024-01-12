package common

import "time"

const APP_NAME string = "collect-metrics"

// 默认加载配置文件路径
const COLLECT_METRICS_CONFIG_PATH string = "./deploy/xconf/collect-metrics.yml"

// 解析配置异常错误码
const PARSE_CONFIG_ERROR int = 10001

// Process Metric 错误码
const COLLECT_PROCESS_METRICS_ERROR int = 10002

// Session Metric 错误码
const COLLECT_SESSION_METRICS_ERROR int = 10003

// TCP Metric 错误码
const COLLECT_TCP_METRICS_ERROR int = 10004

// 执行Shell Cli 错误码
const EXECUTE_CLI_ERROR int = 10005

// 格式化Shell Cli 查询条件错误码
const FORMAT_CLI_QUERY_ERROR int = 10006

// 获取主机名、IP 地址错误码
const GET_HOSTINFO_ERROR int = 10007

// 初始化Histogram Bucket 错误码
const HISTOGRAM_BUCKET_ERROR int = 10008

// 初始化Summary Median 错误码
const SUMMARY_BUCKET_ERROR int = 10009

// 运行模式 Debug
const RUN_WITH_DEBUG string = "debug"

// 运行模式 Release
const RUN_WITH_RELEASE string = "release"

// 默认运行周期
var RUN_COMMON_CYCLE time.Duration = 60 * time.Second

// Gauge Metrics 定义
var GaugeMetrics map[string]map[string]interface{} = map[string]map[string]interface{}{
	"netstat": {
		"name":   "netstat_group",
		"help":   "get netstat order by state",
		"labels": []string{"state"},
		"cmd":    "netstat -an | grep tcp | grep -v grep | awk '{print $NF}' | sort | uniq -c",
	},
	"session": {
		"name":   "session_group",
		"help":   "get tty order by user",
		"labels": []string{"user"},
		"cmd":    "who | awk '{print $1}' | sort | uniq -c",
	},
	"process": {
		"name":   "process_group",
		"help":   "get process order by user",
		"labels": []string{"user"},
		"cmd":    "ps aux | grep -v COMMAND | grep -v grep | awk '{print $1}' | sort | uniq -c",
	},
}

// Counter Metrics 定义
var CounterMetrics map[string]map[string]interface{} = map[string]map[string]interface{}{
	"requests": {
		"name":   "requests_url_total",
		"help":   "get requests order by url",
		"labels": []string{"url"},
	},
}

// Histogram Metrics 定义
var HistogramMetrics map[string]map[string]interface{} = map[string]map[string]interface{}{
	"delay": {
		"name":   "requests_delay_with_histogram",
		"help":   "Total number of HTTP requests delay with histogram",
		"labels": []string{"code"},
	},
}

// Summary Metrics 定义
var Summary map[string]map[string]interface{} = map[string]map[string]interface{}{
	"delay": {
		"name":   "requests_delay_with_summary",
		"help":   "Total number of HTTP requests delay with summary",
		"labels": []string{"code"},
	},
}

// 支持的运行模式
var RUN_MODE []string = []string{"debug", "release"}

// 声明GIN路由
var URL_PREFIX map[string]string = map[string]string{
	"gauge":     "/gmetrics",
	"counter":   "/cmetrics",
	"summary":   "/smetrics",
	"histogram": "/hmetrics",
	"all":       "/metrics",
	"noroute":   "/metrics/*path",
}
