package logging

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	defaultLogFormat       = "[%level%] %time% - %mssg%"
	defaultTimestampFormat = time.RFC3339
)

var log *logrus.Logger

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

func init() {
	godotenv.Load(".env")
	env := os.Getenv("ENVIRONMENT")
	logFile := os.Getenv("LOGDIR")

	fmt.Printf("Logfile: %v", logFile)
	fmt.Printf("Env: %v", env)
	if _, err := os.Stat("logs"); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir("logs", os.ModePerm); err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logLevel := logrus.DebugLevel
	if env == "prod" {
		logLevel = logrus.InfoLevel
	}

	log = &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logLevel,
		Formatter: &CustomLogFormat{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "\n[%level%] %time% - %mssg%",
		},
	}
	log.SetReportCaller(true)
}

// Debug
func Debug(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Info
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warning
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Errors
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

// panic
func Panic(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
