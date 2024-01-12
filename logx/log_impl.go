package logx

import (
	"collect-metrics/common"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogrusLogger 用于构造 Logger实例
type LogrusLogger struct {
	*logrus.Logger
}

func (l *LogrusLogger) Debug(a ...interface{}) {
	// 调用 logrus.Logger 的 Debug 接口
	l.Logger.Debug(ColorFuncMap["Debug"](a...))
}

func (l *LogrusLogger) Debugf(format string, a ...interface{}) {
	// 调用 logrus.Logger 的 Debugf 接口
	l.Logger.Debugf(ColorFuncMapWithFormat["Debugf"](format, a...))
}

func (l *LogrusLogger) Info(a ...interface{}) {
	// 调用 logrus.Logger 的 Info 接口
	l.Logger.Info(ColorFuncMap["Info"](a...))
}

func (l *LogrusLogger) Infof(format string, a ...interface{}) {
	// 调用 logrus.Logger 的 Infof 接口
	l.Logger.Infof(ColorFuncMapWithFormat["Infof"](format, a...))
}

func (l *LogrusLogger) Warn(a ...interface{}) {
	// 调用 logrus.Logger 的 Warn 接口
	l.Logger.Warn(ColorFuncMap["Warn"](a...))
}

func (l *LogrusLogger) Warnf(format string, a ...interface{}) {
	// 调用 logrus.Logger 的 Warnf 接口
	l.Logger.Warnf(ColorFuncMapWithFormat["Warnf"](format, a...))
}

func (l *LogrusLogger) Error(a ...interface{}) {
	// 调用 logrus.Logger 的 Error 接口
	l.Logger.Error(ColorFuncMap["Error"](a...))
}

func (l *LogrusLogger) Errorf(format string, a ...interface{}) {
	// 调用 logrus.Logger 的 Errorf 接口
	l.Logger.Errorf(ColorFuncMapWithFormat["Errorf"](format, a...))
}

func (l *LogrusLogger) Fatal(a ...interface{}) {
	// 调用 logrus.Logger 的 Fatal 接口
	l.Logger.Fatal(ColorFuncMap["Fatal"](a...))
}

func (l *LogrusLogger) Fatalf(format string, a ...interface{}) {
	// 调用 logrus.Logger 的 Fatalf 接口
	l.Logger.Fatalf(ColorFuncMapWithFormat["Fatalf"](format, a...))
}

// LumberjackLoggerConfig 代表 lumberjack.Logger配置
type LumberjackLoggerConfig struct {
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename" validate:"omitempty,min=1"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size" validate:"omitempty,min=1"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups" validate:"omitempty,min=0"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age" validate:"omitempty,min=0"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress" validate:"omitempty"`
}

// LogrusConfig 代表 Logrus 配置
type LogrusConfig struct {
	LumberjackLogger *LumberjackLoggerConfig `mapstructure:"lumberjack_logger" json:"lumberjack_logger" yaml:"lumberjack_logger" validate:"omitempty"`
	LogLevel         string                  `mapstructure:"log_level" json:"log_level" yaml:"log_level" validate:"omitempty"`
	ReportCaller     bool                    `mapstructure:"report_caller" json:"report_caller" yaml:"report_caller" validate:"omitempty"`
	Stdout           bool                    `mapstructure:"stdout" json:"stdout" yaml:"stdout" validate:"omitempty"`
	EnabledWrite     bool                    `mapstructure:"enabled_write" json:"enabled_write" yaml:"enabled_write" validate:"omitempty"`
}

// GetDefaultLogrusConfig 获取默认的 LogrusConfig 实例
func GetDefaultLogrusConfig() *LogrusConfig {
	return &LogrusConfig{
		ReportCaller: true,
		Stdout:       true,
		LogLevel:     "debug",
		EnabledWrite: false,
	}
}

// GetDeefaultLumberjackConfig
func GetDeefaultLumberjackConfig() *LumberjackLoggerConfig {
	return &LumberjackLoggerConfig{
		Filename:   "./logs/" + common.APP_NAME,
		MaxSize:    100,
		MaxAge:     7,
		MaxBackups: 5,
		Compress:   true,
	}
}

func NewLogrusLogger(config *LogrusConfig) Logger {
	if config == nil {
		config = GetDefaultLogrusConfig()
	}

	// 初始化logrus对象
	ins := logrus.New()
	if config.Stdout {
		ins.SetOutput(os.Stdout)
	}
	// 回调
	if config.ReportCaller {
		ins.SetReportCaller(config.ReportCaller)
	}
	// 日志落地文件
	if config.EnabledWrite {
		if config.LumberjackLogger == nil {
			config.LumberjackLogger = GetDeefaultLumberjackConfig()
		}
		ins.SetOutput(
			&lumberjack.Logger{
				Filename:   config.LumberjackLogger.Filename,
				MaxSize:    config.LumberjackLogger.MaxSize,
				MaxBackups: config.LumberjackLogger.MaxBackups,
				MaxAge:     config.LumberjackLogger.MaxAge,
				Compress:   config.LumberjackLogger.Compress,
			},
		)
	}
	// 自定义日志格式
	ins.SetFormatter(&CSFormattor{})
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	ins.SetLevel(logLevel)

	return &LogrusLogger{
		Logger: ins,
	}
}
