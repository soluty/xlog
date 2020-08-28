package xlog

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync/atomic"
)

// https://github.com/appliedgocode/what
// 本库是受到上面的启发，对what库的作者有些看法不认同而修改得来

var mLogs []*log.Logger
var rootDir string
var callerDepth = 2
var boot int32

func init() {
	mLogs = append(mLogs, log.New(os.Stdout, "", log.LstdFlags))
}

// 在main函数中调用Boot方法，并且正确填写项目根目录的相对路径，否则会显示全路径
// 如果要在某个init方法中调用,则需要在匿名函数中调用, Boot方法一般调用一次即可
func Boot(mainPath string, wrapDepth int, w ...io.Writer) {
	if atomic.AddInt32(&boot, 1) != 1 {
		return
	}
	callerDepth = callerDepth + wrapDepth
	file, _, _ := fileAndLine(callerDepth, true)
	rootDir = path.Join(path.Dir(file))
	if mainPath != "" {
		rootDir = path.Join(rootDir, mainPath)
	}
	if len(w) > 0 {
		mLogs = nil
		for _, value := range w {
			mLogs = append(mLogs, log.New(value, "", log.LstdFlags))
		}
	}
}

func fileAndLine(depth int, isCallInInit bool) (file string, line int, name string) {
	pc, file, line, ok := runtime.Caller(depth)
	if !isCallInInit && ok && rootDir != "" {
		file = strings.TrimPrefix(file, rootDir+"/")
	}
	if ok {
		name = runtime.FuncForPC(pc).Name()
	}
	return file, line, name
}
