# xlog
极简的日志框架, 仅提供info, debug和 error级别的日志, error级别提供堆栈信息

可以通过xlog.Boot()来设置特定的启动, 具体可以参考example中调用的例子

debug消息包含日期,文件名,行号,包名,函数名

info消息只包含日期

error消息除了debug消息包含的以外,还包含当时gorutine的堆栈信息

debug消息可以通过启动参数打开, 默认情况下是关闭的.