package main

import (
	"flag"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"k8s.io/klog/v2"
)

// go run main.go -alsologtostderr -logtostderr=0 -v=3
func main() {
	klog.InitFlags(flag.CommandLine)

	// 日志轮转配置
	logger := &lumberjack.Logger{
		Filename:   "log-operator.log",
		MaxSize:    1,     // 日志文件大小上限：1MB
		MaxAge:     30,    // 保留旧日志文件最大天数：30
		MaxBackups: 30,    // 保留旧日志文件数：30
		LocalTime:  true,  // 使用本地时间
		Compress:   false, // 是否压缩
	}

	// 将 klog 输出重定向至 lumberjack
	klog.SetOutput(logger)

	flag.Parse()

	// 模拟日志
	for {
		klog.Infoln(time.Now())
	}
}
