package logx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func InitializeLog() {
	// 配置日志格式
	log.SetReportCaller(true)
	log.SetFormatter(&CSFormattor{})
	log.SetLevel(log.DebugLevel)
}

type CSFormattor struct{}

func (m *CSFormattor) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := "INFO"
	switch entry.Level {
	case log.DebugLevel:
		level = log.DebugLevel.String()
	case log.WarnLevel:
		level = log.WarnLevel.String()
	case log.ErrorLevel:
		level = log.ErrorLevel.String()
	case log.FatalLevel, log.PanicLevel:
		level = log.FatalLevel.String()
	}
	buf := entry.Buffer
	if buf == nil {
		buf = &bytes.Buffer{}
	}
	// _, file, line, _ := runtime.Caller(2)
	buf.WriteString(
		fmt.Sprintf(
			"%s | %s | %s:%d | %s",
			timestamp,
			level,
			// 调用者
			filepath.Base(entry.Caller.File),
			// 调用行号
			entry.Caller.Line,
			entry.Message,
		),
	)
	if len(entry.Data) > 0 {
		data, err := json.Marshal(entry.Data)
		if err == nil {
			buf.WriteByte('\t')
			buf.Write(data)
		}
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}
