package log

import (
	"fmt"
	"log"
)

type defaultLogger struct {
	logger InnerLogger
	opts   Options
}

func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) Log(level Level, v ...interface{}) {
	if !l.opts.level.Enabled(level) {
		return
	}
	l.logger.Output(l.opts.calldepth, fmt.Sprint(v...))
}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {
	if !l.opts.level.Enabled(level) {
		return
	}
	l.logger.Output(l.opts.calldepth, fmt.Sprintf(format, v...))
}

func (l *defaultLogger) String() string {
	return "default"
}

func newLogger(opts ...Option) Logger {
	opt := newOptions(opts...)

	innerLogger := log.New(opt.out, opt.prefix, opt.flags)

	return &defaultLogger{
		logger: innerLogger,
		opts:   *opt,
	}
}
