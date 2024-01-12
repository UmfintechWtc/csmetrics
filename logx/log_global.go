package logx

var globalLogger Logger

// 初始化全局日志接口
func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}

/*
定义全局日志，模块级别
logx.Debug		-> 调用Logger.Debug 接口
logx.Debugf		-> 调用Logger.Debugf 接口
logx.Info		-> 调用Logger.Info 接口
logx.Infof		-> 调用Logger.Infof 接口
logx.Warn		-> 调用Logger.Warn 接口
logx.Warnf		-> 调用Logger.Warnf 接口
logx.Error		-> 调用Logger.Error 接口
logx.Errorf		-> 调用Logger.Errorf 接口
logx.Fatal		-> 调用Logger.Fatal 接口
logx.Fatalf		-> 调用Logger.Fatalf 接口
*/
func Debug(a ...interface{}) {
	globalLogger.Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	globalLogger.Debugf(format, a...)
}

func Info(a ...interface{}) {
	globalLogger.Info(a...)
}

func Infof(format string, a ...interface{}) {
	globalLogger.Infof(format, a...)
}

func Warn(a ...interface{}) {
	globalLogger.Warn(a...)
}

func Warnf(format string, a ...interface{}) {
	globalLogger.Warnf(format, a...)
}

func Error(a ...interface{}) {
	globalLogger.Error(a...)
}

func Errorf(format string, a ...interface{}) {
	globalLogger.Errorf(format, a...)
}

func Fatal(a ...interface{}) {
	globalLogger.Fatal(a...)
}

func Fatalf(format string, a ...interface{}) {
	globalLogger.Fatalf(format, a...)
}
