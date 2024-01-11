package main

import (
	"collect-metrics/logx"
)

func main() {
	logger := logx.NewLogrusLogger()
	logger.Info("this is a test")

}
