package common

const (
	COLLECT_METRICS_CONFIG_PATH   = "./deploy/xconf/collect-metrics.yml"
	PARSE_CONFIG_ERROR            = 10001
	COLLECT_PROCESS_METRICS_ERROR = 10002
	COLLECT_SESSION_METRICS_ERROR = 10003
	COLLECT_TCP_METRICS_ERROR     = 10004
	EXECUTE_CLI_ERROR             = 10005
	FORMAT_CLI_QUERY_ERROR        = 10006
	GET_HOSTINFO_ERROR            = 10007
	FAILED_CODE                   = "Failed"
	SUCCEED_CODE                  = "Succeed"
	RUN_WITH_DEBUG                = "debug"
	RUN_WITH_RELEASE              = "release"
	RUN_WITH_TEST                 = "test"
)

var (
	BASE_LABELS                     = []string{"hostname", "ip"}
	GAUGE_PROCESS_METRICS_LABELS    = append(BASE_LABELS, "user")
	GAUGE_NETSTAT_METRICS_LABELS    = append(BASE_LABELS, "state")
	GAUGE_SESSION_METRICS_LABELS    = append(BASE_LABELS, "user")
	COUNTER_REQUESTS_METRICS_LABELS = []string{"path", "code"}
	RUN_MODE                        = []string{"debug", "release"}
	URL_PREFIX                      = map[string]string{
		"gauge":     "/gmetrics",
		"counter":   "/cmetrics",
		"summary":   "/smetrics",
		"histogram": "/hmetrics",
		"all":       "/metrics",
	}
)
