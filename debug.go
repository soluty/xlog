// +build debug

package xlog

import "strconv"

// Debug方法除附带额外的调用文件行号和包名函数名,并且会自动加上换行符
func Debug(fmt string, args ...interface{}) {
	file, line, name := fileAndLine(callerDepth, false)
	debugLog.Printf(file+":"+strconv.Itoa(line)+" in ("+name+"): "+fmt+"\n", args...)
}
