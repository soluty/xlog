package main

import (
	"errors"

	"github.com/soluty/xlog/example/wrap/logger"
)

func main() {
	logger.Debug("debug")
	logger.Info("info")
	logger.Error(errors.New("err"))
	abc()
}
