package config

import (
	"collect-metrics/common"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	configOnce sync.Once
)

type CollectMetricsConfiguration struct {
	// IsContainer *bool   `json:"is_container" yaml:"is_container" mapstructure:"server" binding:"omitempty"`
	Server  Server  `json:"server"`
	Metrics Metrics `json:"metrics"`
}

type Server struct {
	Listen string `json:"listen" yaml:"listen" binding:"required"`
	Port   int    `json:"port" yaml:"port" binding:"required"`
	// time.Duration 的零值是 0s, *time.Duration 的零值是 nil
	GlobalPeriodSeconds *time.Duration `json:"periodSeconds" yaml:"periodSeconds" mapstructure:"periodSeconds" binding:"omitempty"`
}
type Metrics struct {
	TCP     *Netstat `json:"netstat" yaml:"netstat" binding:"omitempty"`
	PS      *Process `json:"process" yaml:"process" binding:"omitempty"`
	Session *Tty     `json:"tty" yaml:"tty" binding:"omitempty"`
}
type Netstat struct {
	PeriodSeconds *time.Duration `json:"periodSeconds" yaml:"periodSeconds" mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `json:"verify_type" yaml:"verify_type" binding:"required, dive, min=1"`
	MetricName    string         `json:"metric_name" yaml:"metric_name" binding:"required"`
	MetricLabels  []string       `json:"metric_labels" yaml:"metric_labels" binding:"required, dive, min=1"`
}

type Process struct {
	PeriodSeconds *time.Duration `json:"periodSeconds" yaml:"periodSeconds" mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `json:"verify_type" yaml:"verify_type" binding:"required, dive, min=1"`
	MetricName    string         `json:"metric_name" yaml:"metric_name" binding:"required"`
	MetricLabels  []string       `json:"metric_labels" yaml:"metric_labels" binding:"required, dive, min=1"`
}
type Tty struct {
	PeriodSeconds *time.Duration `json:"periodSeconds" yaml:"periodSeconds" mapstructure:"periodSeconds" binding:"omitempty"`
	VerifyType    []string       `json:"verify_type" yaml:"verify_type" binding:"required, dive, min=1"`
	MetricName    string         `json:"metric_name" yaml:"metric_name" binding:"required"`
	MetricLabels  []string       `json:"metric_labels" yaml:"metric_labels" binding:"required, dive, min=1"`
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
	configOnce.Do(func() {
		err = config.parse(filePath)
	})
	return config, err
}
