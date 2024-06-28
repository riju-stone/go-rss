package utils

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultLogFormat       = "[%level%] %time% - %mssg%"
	defaultTimestampFormat = time.RFC3339
)

type CustomLogFormat struct {
	TimestampFormat string
	LogFormat       string
}

func (f *CustomLogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	logOutput := f.LogFormat
	if logOutput == "" {
		logOutput = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	logOutput = strings.Replace(logOutput, "%time%", entry.Time.Format(timestampFormat), 1)
	logOutput = strings.Replace(logOutput, "%mssg%", entry.Message, 1)
	logLevel := strings.ToUpper(entry.Level.String())
	logOutput = strings.Replace(logOutput, "%level%", logLevel, 1)

	for key, value := range entry.Data {
		switch v := value.(type) {
		case string:
			logOutput = strings.Replace(logOutput, "%"+key+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			logOutput = strings.Replace(logOutput, "%"+key+"%", s, 1)
		case bool:
			b := strconv.FormatBool(v)
			logOutput = strings.Replace(logOutput, "%"+key+"%", b, 1)
		}
	}

	return []byte(logOutput), nil
}

func InitLogger(f io.Writer) *logrus.Logger {
	env := os.Getenv("ENVIRONMENT")

	logLevel := logrus.DebugLevel
	if env == "prod" {
		logLevel = logrus.InfoLevel
	}

	log := &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logLevel,
		Formatter: &CustomLogFormat{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "\n[%level%] %time% - %mssg%",
		},
	}

	return log
}
