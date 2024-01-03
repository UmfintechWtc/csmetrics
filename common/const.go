package common

const (
	COLLECT_METRICS_CONFIG_PATH   = "./deploy/xconf/collect-metrics.yml"
	PARSE_CONFIG_ERROR            = 10001
	COLLECT_PROCESS_METRICS_ERROR = 10002
	COLLECT_SESSION_METRICS_ERROR = 10003
	COLLECT_TCP_METRICS_ERROR     = 10004
	FAILED_CODE                   = "Failed"
	SUCCEED_CODE                  = "Succeed"
	RUN_WITH_DEBUG                = "debug"
	RUN_WITH_RELEASE              = "release"
	RUN_WITH_TEST                 = "test"
)

var (
	BASE_LABELS                      = []string{"hostname", "ip"}
	GAUGE_PROCESS_METRICS_LABELS     = append(BASE_LABELS, "user")
	GAUGE_NETSTAT_METRICS_LABELS     = append(BASE_LABELS, "state")
	GAUGE_SESSION_METRICS_LABELS     = append(BASE_LABELS, "user")
	COUNTER_INTERFACE_METRICS_LABELS = append(BASE_LABELS, "code", "path")
	RUN_MODE                         = []string{"debug", "release", "test"}
)
