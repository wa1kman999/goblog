package logger

import (
	"context"
	"os"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

const (
	// Type 日志类型
	Type        = "app"
	LogLevelKey = "MESHER_LOG_LEVEL"
	LogFormat   = "MESHER_LOG_FORMAT"
)

var (
	// logger 应用程序日志句柄
	logger *logrus.Logger = newLogger(logLevel())
)

func logLevel() logrus.Level {
	lv := os.Getenv(LogLevelKey)
	if len(lv) > 0 {
		llv, err := logrus.ParseLevel(lv)
		if err == nil {
			return llv
		}
	}
	return logrus.InfoLevel
}

// newLogger 创建新的logger
func newLogger(lv logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(lv)

	var formatter logrus.Formatter
	if os.Getenv(LogFormat) == "text" {
		formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	} else {
		formatter = &logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableHTMLEscape: true,
		}
	}
	// 设置日志格式
	logger.SetFormatter(formatter)

	return logger
}

type Logger struct {
	logger       *logrus.Logger
	globalFields map[string]interface{}
	module       string
}

func NewLogger(module string, globalFields map[string]interface{}) *Logger {
	l := &Logger{
		logger: logger,
		module: module,
	}
	if globalFields != nil {
		l.globalFields = globalFields
	}
	return l
}

func (l *Logger) WithContext(ctx context.Context) *Entry {

	var fields map[string]interface{}
	if l.globalFields != nil {
		fields = l.globalFields
	} else {
		fields = make(map[string]interface{})
	}
	if len(l.module) > 0 {
		fields["mod"] = l.module
	}
	fields["type"] = Type

	entry := logrus.NewEntry(logger).
		WithFields(fields).
		WithContext(ctx)

	return &Entry{entry: entry}
}

type Entry struct {
	entry *logrus.Entry
}

func WithContext(ctx context.Context) *Entry {
	entry := logrus.NewEntry(logger).
		WithField("type", Type).
		WithContext(ctx)
	return &Entry{entry: entry}
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{entry: e.entry.WithField(key, value)}
}

func (e *Entry) WithFields(fields logrus.Fields) *Entry {
	return &Entry{entry: e.entry.WithFields(logrus.Fields(fields))}
}

func (e *Entry) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.entry.Info(args...)
}

func (e *Entry) Warning(args ...interface{}) {
	e.entry.Warning(args...)
}

func (e *Entry) Error(err error, args ...interface{}) {
	if er, ok := err.(*errors.Error); ok {
		e.entry.WithError(err).
			WithField("stack", er.ErrorStack()).
			Error(args...)
		return
	}
	e.entry.WithError(err).Error(args...)
}

func (e *Entry) Fatal(err error, args ...interface{}) {
	e.entry.WithError(err).Fatal(args...)
}

func (e *Entry) Panic(err error, args ...interface{}) {
	e.entry.WithError(err).Panic(args...)
}

// Entry Printf family functions

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

func (e *Entry) Warningf(format string, args ...interface{}) {
	e.entry.Warningf(format, args...)
}

func (e *Entry) Errorf(err error, format string, args ...interface{}) {
	if er, ok := err.(*errors.Error); ok {
		e.entry.WithError(err).
			WithField("stack", er.ErrorStack()).
			Errorf(format, args...)
		return
	}
	e.entry.WithError(err).Errorf(format, args...)
}

func (e *Entry) Fatalf(err error, format string, args ...interface{}) {
	e.entry.WithError(err).Fatalf(format, args...)
}

func (e *Entry) Panicf(err error, format string, args ...interface{}) {
	e.entry.WithError(err).Panicf(format, args...)
}
