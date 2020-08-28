package xlog

// Info方法除了时间以外不带任何额外信息,会自动加上换行符
func Info(fmt string, args ...interface{}) {
	for _, logger := range mLogs {
		logger.Printf("[INFO] "+fmt+"\n", args...)
	}
}
