package logx

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func (l *LogrusLogger) Debug(a ...interface{}) {
	l.Logger.Debug(ColorFuncMap["Debug"](a...))
}

func (l *LogrusLogger) Debugf(format string, a ...interface{}) {
	l.Logger.Debugf(ColorFuncMapWithFormat["Debugf"](format, a...))
}

func (l *LogrusLogger) Info(a ...interface{}) {
	l.Logger.Info(ColorFuncMap["Info"](a...))
}

func (l *LogrusLogger) Infof(format string, a ...interface{}) {
	l.Logger.Infof(ColorFuncMapWithFormat["Infof"](format, a...))
}

func (l *LogrusLogger) Warn(a ...interface{}) {
	l.Logger.Warn(ColorFuncMap["Warn"](a...))
}

func (l *LogrusLogger) Warnf(format string, a ...interface{}) {
	l.Logger.Warnf(ColorFuncMapWithFormat["Warnf"](format, a...))
}

func (l *LogrusLogger) Error(a ...interface{}) {
	l.Logger.Error(ColorFuncMap["Error"](a...))
}

func (l *LogrusLogger) Errorf(format string, a ...interface{}) {
	l.Logger.Errorf(ColorFuncMapWithFormat["Errorf"](format, a...))
}

func (l *LogrusLogger) Fatal(a ...interface{}) {
	l.Logger.Fatal(ColorFuncMap["Fatal"](a...))
}

func (l *LogrusLogger) Fatalf(format string, a ...interface{}) {
	l.Logger.Fatalf(ColorFuncMapWithFormat["Fatalf"](format, a...))
}

func NewLogrusLogger() Logger {
	ins := logrus.New()
	ins.SetOutput(os.Stdout)
	ins.SetReportCaller(true)
	ins.SetFormatter(&CSFormattor{})
	ins.SetLevel(logrus.DebugLevel)
	return &LogrusLogger{
		Logger: ins,
	}
}
