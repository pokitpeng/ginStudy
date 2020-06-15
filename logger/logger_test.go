package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	SetConfig(&Config{
		Level:      "debug",
		Filename:   "./test.log",
		MaxSize:    500,
		MaxBackups: 1,
		MaxAge:     1,
	})
	Debug("debug")
	Debugf("debug %s", "str")
	Info("info")
	Infof("info %d", 1)
	Warn("warn")
	Warnf("warn %d", 1)
	Error("error")
	Errorf("error %s", "str")
	// Fatal("fatal")
	// Fatalf("fatal %s", "str")

	WithFields(Fields{"field": "2"}).Debug("debug")
	WithFields(Fields{"field": "2"}).Debugf("debug %d", 1)
	WithFields(Fields{"field": "2"}).Info("info")
	WithFields(Fields{"field": "2"}).Infof("info %d", 1)
	WithFields(Fields{"field": "2"}).Warn("warn")
	WithFields(Fields{"field": "2"}).Warnf("warn %d", 1)
	WithFields(Fields{"field": "2"}).Error("error")
	WithFields(Fields{"field": "2"}).Errorf("error %s", "str")
	// WithFields(Fields{"field": "2"}).Fatal("fatal")
	// WithFields(Fields{"field": "2"}).Fatalf("fatal %s", "str")
}
