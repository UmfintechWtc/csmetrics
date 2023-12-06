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
	watchInterval = 3 * time.Second
	configOnce    sync.Once
)

type Config struct {
	Environment EnvironmentConfig `mapstructure:"environment"`
	Internal    InternalConfig    `mapstructure:"internal"`
}

type EnvironmentConfig struct {
	MySQL MySQLConfig `mapstructure:"internal_mysql"`
}
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type InternalConfig struct {
	Server struct {
		Listen         string        `mapstructure:"listen"`
		Port           int           `mapstructure:"port"`
		ScrapeInterval time.Duration `mapstructure:"scrapeInterval"`
	} `mapstructure:"server"`
	Supervisor struct {
		State    string `mapstructure:"state"`
		LogsPath string `mapstructure:"logs"`
		Listen   string `mapstructure:"listen"`
		Port     int    `mapstructure:"port"`
		Version  string `mapstructure:"version"`
	} `mapstructure:"supervisor"`
	Metrics struct {
		Node struct {
			Cpu  map[string]interface{} `mapstructure:"cpu"`
			Mem  map[string]interface{} `mapstructure:"mem"`
			Disk map[string]interface{} `mapstructure:"disk"`
		} `mapstructure:"node"`
		Server map[string]map[string]interface{} `mapstructure:"server"`
	} `mapstructure:"metrics"`
}

func (C *Config) parse(path string) error {
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

func LoadInternalConfig(filePath string) (*Config, error) {
	var config *Config = &Config{}
	var err error
	err = config.parse(filePath)
	return config, err
}
