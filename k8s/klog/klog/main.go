package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"k8s.io/klog/v2"
)

// go run main.go -alsologtostderr -logtostderr=0 -v=3 -flush-interval=30 -log_file_max_size=1
func main() {
	var flushInterval int
	flag.IntVar(&flushInterval, "flush-interval", 10, "the interval(in seconds) at which logs are flushed to disk")
	klog.InitFlags(flag.CommandLine)

	// 日志目录名规范：/apps/logs/qfusion/log-operator/2024-09
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirDate := time.Now().Format("2006-01")
	logDir := fmt.Sprintf(dir + "/" + dirDate)
	if err = os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(err)
	}

	// 日志文件名规范：2024-09-03.log
	fileDate := time.Now().Format("2006-01-02")
	logFileName := logDir + "/" + fileDate + ".log"

	// 日志轮转配置
	logger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    1,     // 日志文件大小上限：1MB
		MaxAge:     30,    // 保留旧日志文件最大天数：30
		MaxBackups: 5,     // 保留旧日志文件数：5
		LocalTime:  true,  // 使用本地时间
		Compress:   false, // 是否压缩
	}

	// 将 klog 输出重定向至 lumberjack
	klog.SetOutput(logger)

	flag.Parse()

	go func() {
		for {
			time.Sleep(time.Duration(flushInterval) * time.Second)
			klog.Flush()
		}
	}()

	// 模拟日志
	for {
		klog.Infoln(time.Now())
	}
}
