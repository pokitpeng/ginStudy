package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields = logrus.Fields

// Config define
type Config struct {
	Level      string
	Filename   string
	MaxSize    int // MB
	MaxBackups int
	MaxAge     int //days
}

// init 初始化日志
func InitLogger() {
	logger.Formatter = &logrus.TextFormatter{DisableColors: true}
	SetConfig(&Config{
		Level:      viper.GetString("log.level"),
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),
	})
}

// SetConfig set log config
func SetConfig(config *Config) {
	logger.SetLevel(parseLogLevel(config.Level))
	multiWriter := io.MultiWriter(os.Stdout,
		&lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
		})
	logger.SetOutput(multiWriter)
}

// SetLogLevel set log level
func SetLogLevel(l string) {
	logger.Level = parseLogLevel(l)
}

func parseLogLevel(l string) logrus.Level {
	switch strings.ToUpper(l) {
	default:
		return logrus.InfoLevel
	case "DEBUG", "5":
		return logrus.DebugLevel
	case "INFO", "4":
		return logrus.InfoLevel
	case "WARN", "3":
		return logrus.WarnLevel
	case "ERROR", "2":
		return logrus.ErrorLevel
	case "FATAL", "1":
		return logrus.FatalLevel
	case "PANIC", "0":
		return logrus.PanicLevel
	}
}

// Entry define
type Entry struct {
	e *logrus.Entry
}

// WithFields return entry
func WithFields(fields Fields) *Entry {
	return &Entry{
		e: logger.WithFields(fields),
	}
}

// Debug logs a message at level Debug on the standard logger.
func (entry *Entry) Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Debug(args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func (entry *Entry) Debugf(format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Debugf(format, args...)
	}
}

// Info logs a message at level Info on the standard logger.
func (entry *Entry) Info(args ...interface{}) {
	entry.e.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.e.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (entry *Entry) Warn(args ...interface{}) {
	entry.e.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.e.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (entry *Entry) Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Error(args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func (entry *Entry) Errorf(format string, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Errorf(format, args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func (entry *Entry) Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Fatal(args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger.
func (entry *Entry) Fatalf(format string, args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry.e.Data["file"] = fileInfo(2)
		entry.e.Fatalf(format, args...)
	}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Debug(args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Debugf(format, args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Error(args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Errorf(format, args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Fatal(args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{"file": fileInfo(2)})
		entry.Fatalf(format, args...)
	}
}

func fileInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	function := runtime.FuncForPC(pc).Name()
	slash := strings.LastIndex(function, "/")
	if slash >= 0 {
		function = function[slash+1:]
	}
	return fmt.Sprintf("%s:%d:%s", file, line, function)
}
