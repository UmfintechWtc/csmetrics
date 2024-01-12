package logx

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

type CSFormattor struct{}

var callLogFunctionName string = "log_global.go"
var skipInt int

// 自定义日志格式
func (m *CSFormattor) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := "info   "
	switch entry.Level {
	case log.DebugLevel:
		level = log.DebugLevel.String() + "  "
	case log.WarnLevel:
		level = log.WarnLevel.String()
	case log.ErrorLevel:
		level = log.ErrorLevel.String() + "  "
	case log.FatalLevel, log.PanicLevel:
		level = log.FatalLevel.String() + "  "
	}
	buf := entry.Buffer
	if buf == nil {
		buf = &bytes.Buffer{}
	}
	for skip := 0; skip <= 20; skip++ {
		_, file, _, _ := runtime.Caller(skip)
		if filepath.Base(file) == callLogFunctionName {
			skipInt = skip + 1
			break
		}
	}
	// for skip := 0; ; skip++ {
	// 	pc, file, line, ok := runtime.Caller(skip)
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	// }
	_, file, line, _ := runtime.Caller(skipInt)
	buf.WriteString(
		fmt.Sprintf(
			"%s | %s | %s:%d | %s\n",
			timestamp,
			level,
			// 调用者
			filepath.Base(file),
			// 调用行号
			line,
			entry.Message,
		),
	)
	// buf.WriteByte('\n')
	return buf.Bytes(), nil
}
