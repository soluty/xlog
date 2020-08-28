package xlog

import "runtime/debug"

func Error(err error) {
	if err == nil {
		return
	}
	stack := debug.Stack()
	file, line, name := fileAndLine(callerDepth, false)
	for _, logger := range mLogs {
		logger.Printf("[ERR] %v:%v in (%v): %v\nis in %s", file, line, name, err, stack)
	}
}
