package commands

import (
	"github.com/zysuper/github-sentinel-go/internal/scheduler"
)

// UpdateCommand 处理更新命令，并获取所有订阅仓库的最新发布信息
type UpdateCommand struct {
	scheduler *scheduler.Scheduler
}

// NewUpdateCommand 初始化 UpdateCommand
func NewUpdateCommand() *UpdateCommand {
	return &UpdateCommand{scheduler: scheduler.NewScheduler()}
}

func (c *UpdateCommand) Execute(args []string) error {
	dryRun := false
	for _, arg := range args {
		if arg == "dry-run" {
			dryRun = true
			break
		}
	}

	// 调用 Scheduler 的 ManualUpdate 方法
	c.scheduler.ManualUpdate(dryRun, "closed", 7)
	return nil
}
