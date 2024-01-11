package logx

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func (l *LogrusLogger) Debug(a ...interface{}) {
	l.Debug(ColorFuncMap["Debug"](a...))
}

func (l *LogrusLogger) Debugf(format string, a ...interface{}) {
	l.Debugf(ColorFuncMapWithFormat["Debugf"](format, a...))
}

func (l *LogrusLogger) Info(a ...interface{}) {
	l.Info(ColorFuncMap["Info"](a...))
}

func (l *LogrusLogger) Infof(format string, a ...interface{}) {
	l.Infof(ColorFuncMapWithFormat["Infof"](format, a...))
}

func (l *LogrusLogger) Warn(a ...interface{}) {
	l.Warn(ColorFuncMap["Warn"](a...))
}

func (l *LogrusLogger) Warnf(format string, a ...interface{}) {
	l.Warnf(ColorFuncMapWithFormat["Warnf"](format, a...))
}

func (l *LogrusLogger) Error(a ...interface{}) {
	l.Error(ColorFuncMap["Error"](a...))
}

func (l *LogrusLogger) Errorf(format string, a ...interface{}) {
	l.Errorf(ColorFuncMapWithFormat["Errorf"](format, a...))
}

func (l *LogrusLogger) Fatal(a ...interface{}) {
	l.Fatal(ColorFuncMap["Fatal"](a...))
}

func (l *LogrusLogger) Fatalf(format string, a ...interface{}) {
	l.Fatalf(ColorFuncMapWithFormat["Fatalf"](format, a...))
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
