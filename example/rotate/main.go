package main

import (
	"os"
	"time"

	"github.com/scue/go-logrotate"
	"github.com/soluty/xlog"
)

func main() {
	writer := logrotate.New("x.log", "0 * * * * *", 1) // 最多保留3个日志文件
	xlog.Boot("", 0, os.Stderr, writer)
	go writer.CronTask() // 后台定时任务

	for {
		xlog.Debug("%v", time.Now())
		time.Sleep(time.Second)
	}
}
