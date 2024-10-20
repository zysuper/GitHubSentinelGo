package commands

import (
	"github.com/zysuper/github-sentinel-go/internal/scheduler"
)

// UpdateCommand 处理更新命令，并获取所有订阅仓库的最新版本信息
type UpdateCommand struct{}

func (c *UpdateCommand) Execute(args []string) error {
	// 直接调用 scheduler 的 ManualUpdate 方法
	scheduler.ManualUpdate()
	return nil
}
