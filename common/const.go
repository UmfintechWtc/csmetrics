package common

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

// Metrics 基础Label
var BASE_LABELS []string = []string{"hostname", "ip"}

// Gauge Process Metric Label
var GAUGE_PROCESS_METRICS_LABELS []string = append(BASE_LABELS, "user")

// Gauge TCP Metric Label
var GAUGE_NETSTAT_METRICS_LABELS []string = append(BASE_LABELS, "state")

// Gauge Session Metric Label
var GAUGE_SESSION_METRICS_LABELS []string = append(BASE_LABELS, "user")

// Counter Requests Metric Label
var COUNTER_REQUESTS_METRICS_LABELS []string = []string{"path"}

// Histogram Requests Delay Metric Label
var HISTOGRAM_DELAY_METRICS_LABELS []string = []string{"code"}

// Summary Reqypes Delay Metric Label
var SUMMARY_DELAY_METRICS_LABELS []string = []string{"code"}

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
