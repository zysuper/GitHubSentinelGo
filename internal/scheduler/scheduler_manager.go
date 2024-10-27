package scheduler

import (
	"fmt"
	"time"

	"github.com/zysuper/github-sentinel-go/pkg/logger"
)

type SchedulerManager struct {
	scheduler *Scheduler
	days      int
	interval  time.Duration
	quit      chan struct{}
}

// NewSchedulerManager 创建一个新的 SchedulerManager 实例
func NewSchedulerManager(scheduler *Scheduler, days int, interval time.Duration) *SchedulerManager {
	return &SchedulerManager{
		scheduler: scheduler,
		days:      days,
		interval:  interval,
		quit:      make(chan struct{}),
	}
}

// Start 开始定时任务
func (sm *SchedulerManager) Start() {
	ticker := time.NewTicker(sm.interval)
	logger.Info(fmt.Sprintf("SchedulerManager started with interval: %s and days: %d", sm.interval, sm.days))

	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Info("Running scheduled update...")
				sm.scheduler.ManualUpdate(false, "closed", sm.days) // 使用传入的天数作为时间范围
			case <-sm.quit:
				ticker.Stop()
				logger.Info("SchedulerManager stopped.")
				return
			}
		}
	}()
}

// Stop 停止定时任务
func (sm *SchedulerManager) Stop() {
	close(sm.quit)
}
