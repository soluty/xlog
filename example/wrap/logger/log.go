package logger

import (
	"os"

	"github.com/soluty/xlog"
)

func init() {
	func() {
		xlog.Boot("..", 1, os.Stderr)
	}()
}

func Info(fmt string, args ...interface{}) {
	xlog.Info(fmt, args...)
}

func Debug(fmt string, args ...interface{}) {
	xlog.Debug(fmt, args...)
}

func Error(err error) {
	xlog.Error(err)
}
