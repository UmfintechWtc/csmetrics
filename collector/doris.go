package collector

// import (
// 	"collect-metrics/client"
// 	"collect-metrics/common"
// 	"collect-metrics/config"
// 	cmlog "collect-metrics/log"
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"sync"

// 	"github.com/prometheus/client_golang/prometheus"
// )

// var (
// 	ManagedServer       []string
// 	configOnce          sync.Once
// 	createRestartMetric *prometheus.GaugeVec
// 	createTabletMetric  *prometheus.GaugeVec
// 	countRestartValue   int
// )

// type ExportMetrics struct {
// 	RestartCount *prometheus.GaugeVec
// 	TabletCount  *prometheus.GaugeVec
// }

// type DorisMetric struct {
// 	MetricName string
// 	LabelName  []string
// }

// func NewDorisMetric(metricType string, metrics map[string]interface{}) (*DorisMetric, error) {
// 	// get metric name
// 	metricInfo, ok := common.ConvertInterfaceToMap(metrics[metricType])
// 	if !ok {
// 		return nil, errors.New(
// 			fmt.Sprintf("Invalid %s metric: %v", metricType, metrics[metricType]),
// 		)
// 	}
// 	// get metric label
// 	labels, ok := common.ConvertInterfaceToList(metricInfo["labels"])
// 	if !ok {
// 		return nil, errors.New(
// 			fmt.Sprintf("Invalid %s labels: %v", metricType, metricInfo["labels"]),
// 		)
// 	}
// 	return &DorisMetric{
// 		MetricName: metricInfo["name"].(string),
// 		LabelName:  labels,
// 	}, nil
// }

// func Doris(config *config.CollectMetricsConfiguration) (*ExportMetrics, error) {
// 	allProcess, err := client.NewSupervisor(
// 		config.Internal.Supervisor.Listen,
// 		config.Internal.Supervisor.Port,
// 		config.Internal.Supervisor.Version,
// 		config.Internal.Supervisor.State,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	metrics, ok := common.ConvertInterfaceToMap(
// 		config.Internal.Metrics.Server["doris"]["metrics"],
// 	)
// 	if !ok {
// 		return nil, errors.New(
// 			fmt.Sprintf("Invalid metrics: %v", config.Internal.Metrics.Server["doris"]["metrics"]),
// 		)
// 	}
// 	subServer, ok := common.ConvertInterfaceToList(
// 		config.Internal.Metrics.Server["doris"]["subserver"],
// 	)
// 	if !ok {
// 		return nil, errors.New(
// 			fmt.Sprintf(
// 				"Invalid subServer: %v",
// 				config.Internal.Metrics.Server["doris"]["subserver"],
// 			),
// 		)
// 	}

// 	for _, server := range subServer {
// 		if common.MatchKeyWord(server, allProcess.AllProcessInfo) {
// 			ManagedServer = append(ManagedServer, server)
// 		} else {
// 			cmlog.Warning(fmt.Sprintf("supervisor未托管%s服务", ManagedServer[0]))
// 		}
// 	}
// 	// doris restart metrics
// 	restartMetrics, err := NewDorisMetric("restart", metrics)
// 	if err != nil {
// 		return nil, err
// 	}
// 	restartCmd := fmt.Sprintf(
// 		"awk -v FS=',' '$1 > \"%s\" {print $0}' %s",
// 		common.CurrentTime(config.Internal.Server.ScrapeInterval),
// 		config.Internal.Supervisor.LogsPath,
// 	)
// 	cmdRes, err := common.ExecCmd(restartCmd)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(cmdRes) == 0 {
// 		countRestartValue = 0
// 	} else {
// 		countRestartValue = strings.Count(cmdRes, config.Internal.Metrics.Server["doris"]["keyword"].(string))
// 	}
// 	createRestartMetric = common.CreateGaugeMetrics(restartMetrics.MetricName, subServer[0], restartMetrics.LabelName)
// 	createRestartMetric.WithLabelValues(common.GetHostName(), subServer[0]).Set(float64(countRestartValue))
// 	// doris tablet metrics
// 	tabletMetrics, err := NewDorisMetric("tablet", metrics)
// 	if err != nil {
// 		return nil, err
// 	}
// 	createTabletMetric = common.CreateGaugeMetrics(tabletMetrics.MetricName, subServer[0], tabletMetrics.LabelName)
// 	configOnce.Do(func() {
// 		// only in the first run to registration the metrics
// 		prometheus.MustRegister(createRestartMetric)
// 		prometheus.MustRegister(createTabletMetric)
// 	})

// 	// createRestartMetric.WithLabelValues("0.0.0.0", common.GetHostName(), subServer[0]).Set(float64(1))
// 	createTabletMetric.WithLabelValues("1", "2", "3", "4", "5", "6").Set(1)

// 	return &ExportMetrics{
// 		RestartCount: createRestartMetric,
// 		TabletCount:  createTabletMetric,
// 	}, nil
// }
