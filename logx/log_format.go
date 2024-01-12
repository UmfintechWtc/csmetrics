package logx

import (
	"bytes"
	"collect-metrics/common"
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type CSFormattor struct{}

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
	// _, file, line, _ := runtime.Caller(9)
	file, line := common.GetCallerInfo()
	buf.WriteString(
		fmt.Sprintf(
			"%s | %s | %s:%d | %s",
			timestamp,
			level,
			// 调用者
			filepath.Base(file),
			// 调用行号
			line,
			entry.Message,
		),
	)
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}
