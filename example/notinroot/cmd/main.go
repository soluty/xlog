package main

import (
	"github.com/soluty/xlog"
	"github.com/soluty/xlog/example/notinroot"
)

func main() {
	// 写根目录相对main.go的相对路径
	xlog.Boot("..", 0)
	notinroot.Hello()
	xlog.Debug("aaa")
	xlog.Info("ddd")
}
