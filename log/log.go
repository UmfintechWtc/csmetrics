package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/rifflock/lfshook"
	cmlog "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = cmlog.New()

func init() {
	log.SetFormatter(&cmlog.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(cmlog.InfoLevel)

	// Create an LFShook instance to write logs to a file
	logFilePath := "./collect-metrics.log"
	log.AddHook(createLfsHook(logFilePath))
}

func createLfsHook(logFilePath string) cmlog.Hook {
	writer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    8,
		MaxBackups: 1,
		MaxAge:     1,
		Compress:   true,
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		cmlog.InfoLevel:  writer,
		cmlog.WarnLevel:  writer,
		cmlog.ErrorLevel: writer,
		cmlog.FatalLevel: writer,
	}, &cmlog.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return lfsHook
}

func logMessage(level cmlog.Level, message string) {
	// pc, file, line, _ := runtime.Caller(2)
	// callerName := runtime.FuncForPC(pc).Name()
	// logMessage := fmt.Sprintf("%s | %s | %s | %s -> %s",
	logMessage := fmt.Sprintf("%s | %s | %s",
		(time.Now()).Format("2006-01-02 15:04:05"),
		strings.ToUpper(level.String()),
		message,
		// callerName,
		// fmt.Sprintf("%s:%d", file, line),
	)
	// log.Info(logMessage)
	fmt.Println(logMessage)
	// if level == cmlog.ErrorLevel {
	// 	os.Exit(1)
	// }
}

func Info(format string) {
	logMessage(cmlog.InfoLevel, fmt.Sprintf(format))
}

func Warning(format string) {
	logMessage(cmlog.WarnLevel, fmt.Sprintf(format))
}

func Error(format string) {
	logMessage(cmlog.ErrorLevel, fmt.Sprintf(format))
}
