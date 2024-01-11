package main

import (
	"collect-metrics/logx"
	"runtime"
)

func logFatal(logger logx.Logger) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		logger.Fatal("err", "|", file, ":", line)
	} else {
		logger.Fatal("err", "| Caller information not available")
	}
}

func main() {
	logger := logx.NewLogrusLogger()
	logFatal(logger)
}
