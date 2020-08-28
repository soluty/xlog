package main

import (
	"errors"

	"github.com/soluty/xlog"
)

func main() {
	xlog.Boot("", 0)
	xlog.Info("info")
	xlog.Debug("debug")
	xlog.Error(errors.New("aa"))
}
