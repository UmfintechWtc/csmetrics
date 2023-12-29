package common

const (
	COLLECT_METRICS_CONFIG_PATH   = "./deploy/xconf/collect-metrics.yml"
	PARSE_CONFIG_ERROR            = 10001
	COLLECT_PROCESS_METRICS_ERROR = 10002
	COLLECT_SESSION_METRICS_ERROR = 10003
	COLLECT_TCP_METRICS_ERROR     = 10004
	FAILED_CODE                   = "Failed"
	SUCCEED_CODE                  = "Succeed"
)

var (
	BASE_LABELS            = []string{"hostname", "ip"}
	PROCESS_METRICS_LABELS = append(BASE_LABELS, "user")
	NETSTAT_METRICS_LABELS = append(BASE_LABELS, "state")
	SESSION_METRICS_LABELS = append(BASE_LABELS, "user")
)
