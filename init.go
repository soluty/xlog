package xlog

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

// https://github.com/appliedgocode/what
// 本库是受到上面的启发，对what库的作者有些看法不认同而修改得来

var debugLog *log.Logger
var errLog *log.Logger
var infoLog *log.Logger

var rootDir string
var callerDepth = 2
var bootLock sync.Mutex
var boot bool

func init() {
	debugLog = log.New(os.Stdout, "", log.LstdFlags)
	infoLog = log.New(os.Stdout, "", 0)
	errLog = log.New(os.Stderr, "", log.LstdFlags)
}

// 在main函数中调用Boot方法，并且正确填写项目根目录的相对路径，否则会显示全路径
// 如果要在某个init方法中调用,则需要在匿名函数中调用, Boot方法一般调用一次即可
// err info debug
func Boot(mainPath string, wrapDepth int, w ...io.Writer) {
	bootLock.Lock()
	defer bootLock.Unlock()
	if boot {
		return
	}
	boot = true
	if len(w) > 2 {
		errLog = log.New(w[0], "", log.LstdFlags)
		infoLog = log.New(w[1], "", 0)
		debugLog = log.New(w[2], "", log.LstdFlags)
	} else if len(w) > 1 {
		errLog = log.New(w[0], "", log.LstdFlags)
		infoLog = log.New(w[1], "", 0)
	} else if len(w) > 0 {
		errLog = log.New(w[0], "", log.LstdFlags)
	}
	callerDepth = callerDepth + wrapDepth
	file, _, _ := fileAndLine(callerDepth, true)
	rootDir = path.Join(path.Dir(file))
	if mainPath != "" {
		rootDir = path.Join(rootDir, mainPath)
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
