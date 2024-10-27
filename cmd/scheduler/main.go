package main

import (
	"flag"
	"time"

	"github.com/zysuper/github-sentinel-go/internal/scheduler"
	"github.com/zysuper/github-sentinel-go/pkg/logger"
)

func main() {
	days := flag.Int("d", 7, "Number of days in the past to fetch project updates")
	// 解析 interval 参数，默认为 24 小时
	intervalStr := flag.String("i", "24h", "Update interval, e.g., '24h', '1h30m', '15m'")

	flag.Parse()

	// 解析 interval 字符串为 Duration 类型
	interval, err := time.ParseDuration(*intervalStr)
	if err != nil {
		logger.Error("Invalid interval format. Please use formats like '24h', '1h30m', or '15m'.")
		return
	}

	schedulerInstance := scheduler.NewScheduler()
	schedulerManager := scheduler.NewSchedulerManager(schedulerInstance, *days, interval)

	logger.Info("Scheduler service started")
	schedulerManager.Start()

	// 阻塞主线程以保持任务运行
	select {}
}
