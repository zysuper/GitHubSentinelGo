package commands

import (
    "fmt"
    "github.com/zysuper/github-sentinel-go/internal/scheduler"
)

// UpdateCommand 处理更新命令
type UpdateCommand struct{}

func (c *UpdateCommand) Execute(args []string) error {
    fmt.Println("Fetching updates...")
    scheduler.ManualUpdate()
    return nil
}
