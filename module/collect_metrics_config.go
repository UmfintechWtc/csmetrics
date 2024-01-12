package config

import (
	"collect-metrics/common"
	"collect-metrics/logx"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type CollectMetricsConfiguration struct {
	Server  Server          `yaml:"server" json:"server" binding:"required"`
	Metrics CateGoryMetrics `yaml:"metrics" yaml:"metrics" binding:"required"`
}
type Server struct {
	Listen string `yaml:"listen" json:"listen"`
	Port   int    `yaml:"port" json:"port" binding:"required"`
	Mode   string `mapstructure:"mode" binding:"omitempty"`
	// time.Duration 的零值是 0s, *time.Duration 的零值是 nil
	GlobalPeriodSeconds *time.Duration     `mapstructure:"periodSeconds" binding:"omitempty"`
	LogrusConfig        *logx.LogrusConfig `mapstructure:"logrus_config" json:"logrus_config" yaml:"logrus_config" binding:"omitempty"`
}
type CateGoryMetrics struct {
	Gauge Gauge `mapstructure:"gauge" binding:"required"`
	// Counter   Counter   `mapstructure:"counter" binding:"required"`
	Histogram Histogram `mapstructure:"histogram" binding:"required"`
	Summary   Summary   `mapstructure:"summary" binding:"required"`
}
type Gauge struct {
	// 包含time.Duration类型，避免转换失败不使用json、yaml
	TCP     Netstat `mapstructure:"netstat" binding:"required"`
	PS      Process `mapstructure:"process" binding:"required"`
	Session Tty     `mapstructure:"session" binding:"required"`
}

//	type Counter struct {
//		Request Requests `mapstructure:"requests" binding:"required"`
//	}
type Histogram struct {
	Delay HistogramDelay `mapstructure:"delay" binding:"required"`
}
type Summary struct {
	Delay SummaryDelay `mapstructure:"delay" binding:"required"`
}
type Netstat struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	// VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	// MetricName    string         `mapstructure:"metric_name" binding:"required"`
	// MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `mapstructure:"metric_labels" binding:"min=1"`
}
type Process struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	// VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	// MetricName    string         `mapstructure:"metric_name" binding:"required"`
	// MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `mapstructure:"metric_labels" binding:"min=1"`
}
type Tty struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	// VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	// MetricName    string         `mapstructure:"metric_name" binding:"required"`
	// MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `json:"metric_labels" json:"metric_labels" binding:"min=1"`
}

//	type Requests struct {
//		MetricName string `mapstructure:"metric_name" binding:"required"`
//		MetricHelp string `mapstructure:"metric_help" binding:"omitempty"`
//	}
type HistogramDelay struct {
	// MetricName string              `mapstructure:"metric_name" binding:"required"`
	// MetricHelp string              `mapstructure:"metric_help" binding:"omitempty"`
	Buckets HistogramDelayRange `mapstructure:"buckets" binding:"omitempty"`
}
type SummaryDelay struct {
	// MetricName string              `mapstructure:"metric_name" binding:"required"`
	// MetricHelp string              `mapstructure:"metric_help" binding:"omitempty"`
	Median map[float64]float64 `mapstructure:"median" binding:"omitempty"`
}
type HistogramDelayRange struct {
	Linear struct {
		Enabled bool               `mapstructure:"enabled"`
		Range   map[string]float64 `mapstructure:"range"`
	} `mapstructure:"linear"`
	Slice struct {
		Enabled bool      `mapstructure:"enabled"`
		Range   []float64 `mapstructure:"range"`
	} `mapstructure:"slice"`
}

func (C *CollectMetricsConfiguration) parse(path string) error {
	if len(path) == 0 {
		path = common.COLLECT_METRICS_CONFIG_PATH
	}
	if !common.FileExists(path) {
		return errors.New(
			fmt.Sprintf(
				"Config file not found %s", path,
			),
		)
	}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(C); err != nil {
		return err
	}
	return nil
}

func LoadInternalConfig(filePath string) (*CollectMetricsConfiguration, error) {
	var config *CollectMetricsConfiguration = &CollectMetricsConfiguration{}
	var err error
	var configOnce sync.Once
	configOnce.Do(func() {
		err = config.parse(filePath)
	})
	return config, err
}
