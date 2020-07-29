package log

import (
	"io"
	"os"
)

var (
	NewLogger     = newLogger
	DefaultLogger = newLogger(WithCallDepth(3))
)

type InnerLogger interface {
	Output(calldepth int, s string) error
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	SetFlags(flag int)
}

type Logger interface {
	Init(opts ...Option) error
	Log(level Level, v ...interface{})
	Logf(level Level, format string, v ...interface{})
	String() string
}

func Log(level Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level Level, template string, v ...interface{}) {
	DefaultLogger.Logf(level, template, v...)
}

func Info(args ...interface{}) {
	DefaultLogger.Log(InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	DefaultLogger.Logf(InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	DefaultLogger.Log(TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Log(DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Log(WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Log(ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger.Logf(ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Log(FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Logf(FatalLevel, template, args...)
	os.Exit(1)
}
