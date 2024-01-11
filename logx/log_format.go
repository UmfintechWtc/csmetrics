package logx

import (
	"bytes"
	"encoding/json"
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
	levelCode := 0
	switch entry.Level {
	case log.DebugLevel:
		level = "DEBUG"
		levelCode = 1
	case log.WarnLevel:
		level = "WARN"
		levelCode = 2
	case log.ErrorLevel:
		level = "ERROR"
		levelCode = 3
	case log.FatalLevel, log.PanicLevel:
		level = "FATAL"
		levelCode = 4
	}
	buf := entry.Buffer
	if buf == nil {
		buf = &bytes.Buffer{}
	}
	buf.WriteString(
		ColorFuncMap[level](
			"%s|%s|%d|%s:%d|%s",
			timestamp,
			level,
			levelCode,
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
