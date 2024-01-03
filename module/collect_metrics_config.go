package config

import (
	"collect-metrics/common"
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
	// time.Duration 的零值是 0s, *time.Duration 的零值是 nil
	GlobalPeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
}

type CateGoryMetrics struct {
	Gauge   Gauge   `mapstructure:"gauge" binding:"required"`
	Counter Counter `mapstructure:"counter" binding:"required"`
}
type Gauge struct {
	// 包含time.Duration类型，避免转换失败不使用json、yaml
	TCP     Netstat `mapstructure:"netstat" binding:"required"`
	PS      Process `mapstructure:"process" binding:"required"`
	Session Tty     `mapstructure:"session" binding:"required"`
}

type Counter struct{}

type Netstat struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	MetricName    string         `mapstructure:"metric_name" binding:"required"`
	MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `mapstructure:"metric_labels" binding:"min=1"`
}

type Process struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	MetricName    string         `mapstructure:"metric_name" binding:"required"`
	MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `mapstructure:"metric_labels" binding:"min=1"`
}
type Tty struct {
	PeriodSeconds *time.Duration `mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `mapstructure:"verify_type" binding:"omitempty"`
	MetricName    string         `mapstructure:"metric_name" binding:"required"`
	MetricHelp    string         `mapstructure:"metric_help" binding:"omitempty"`
	// MetricLabels  []string       `json:"metric_labels" json:"metric_labels" binding:"min=1"`
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
