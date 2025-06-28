package logger

import (
	"log"
	"os"
)

var L Logger

type Logger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
}

type logger struct {
	inner *log.Logger
}

func NewLogger() *logger {
	return &logger{
		inner: log.New(os.Stdout, "", log.LstdFlags),
	}
}

var (
	red   = "\033[31m"
	blue  = "\033[34m"
	green = "\033[32m"
	reset = "\033[0m"
)

var _ Logger = (*logger)(nil)

func (l *logger) Info(message string) {
	l.inner.Printf("%s[INFO]%s %s", blue, reset, message)
}

func (l *logger) Debug(message string) {
	l.inner.Printf("%s[DEBUG]%s %s", green, reset, message)
}

func (l *logger) Error(message string) {
	l.inner.Printf("%s[ERROR]%s %s", red, reset, message)
}
