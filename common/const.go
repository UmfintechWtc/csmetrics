package common

const (
	COLLECT_METRICS_CONFIG_PATH = "./deploy/xconf/collect-metrics.yml"
	FAILED_CODE                 = "Failed"
	SUCCEED_CODE                = "Succeed"
)

var (
	BASE_LABELS            = []string{"hostname", "ip"}
	PROCESS_METRICS_LABELS = append(BASE_LABELS, "user")
	NETSTAT_METRICS_LABELS = append(BASE_LABELS, "state")
	SESSION_METRICS_LABELS = append(BASE_LABELS, "user")
)
