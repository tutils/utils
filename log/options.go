package log

import (
	"io"
	"log"
	"os"
)

type Options struct {
	out       io.Writer
	flags     int
	prefix    string
	level     Level
	calldepth int
}

// Option is option setter for logger
type Option func(*Options)

// default options
var (
	DefaultOut    = os.Stderr
	DefaultFlags  = log.Ldate | log.Ltime | log.Lshortfile
	DefaultPrefix = ""
	DefaultLevel  = InfoLevel
)

func newOptions(opts ...Option) *Options {
	opt := &Options{
		out:       DefaultOut,
		flags:     DefaultFlags,
		prefix:    DefaultPrefix,
		level:     DefaultLevel,
		calldepth: 2,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// WithOutput sets default output for logger
func WithOutput(out io.Writer) Option {
	return func(opts *Options) {
		opts.out = out
	}
}

// WithFlags sets default flags for logger
func WithFlags(flags int) Option {
	return func(opts *Options) {
		opts.flags = flags
	}
}

// WithPrefix sets default prefix for logger
func WithPrefix(prefix string) Option {
	return func(opts *Options) {
		opts.prefix = prefix
	}
}

// WithLevel sets default level for logger
func WithLevel(level Level) Option {
	return func(opts *Options) {
		opts.level = level
	}
}

// WithCallDepth sets default call depth for logger
func WithCallDepth(calldepth int) Option {
	return func(opts *Options) {
		opts.calldepth = calldepth
	}
}
