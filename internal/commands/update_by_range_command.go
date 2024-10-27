package commands

import (
	"fmt"
	"strconv"

	"github.com/zysuper/github-sentinel-go/internal/scheduler"
)

// UpdateByRangeCommand 处理按状态和时间范围更新的命令
type UpdateByRangeCommand struct {
	scheduler *scheduler.Scheduler
}

// NewUpdateByRangeCommand 初始化 UpdateByRangeCommand
func NewUpdateByRangeCommand() *UpdateByRangeCommand {
	return &UpdateByRangeCommand{scheduler: scheduler.NewScheduler()}
}

func (c *UpdateByRangeCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: update_by_range <state> <days>")
	}

	dryRun := false
	for _, arg := range args {
		if arg == "dry-run" {
			dryRun = true
			break
		}
	}

	state := args[0]
	days, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("Invalid number of days: %s", args[1])
	}

	// 调用 Scheduler 的 DailyProgress 方法，传入状态和天数
	c.scheduler.ManualUpdate(dryRun, state, days)
	return nil
}
