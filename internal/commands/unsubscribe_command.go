package commands

import (
    "fmt"
    "github.com/zysuper/github-sentinel-go/internal/repository"
)

// UnsubscribeCommand 处理取消订阅命令
type UnsubscribeCommand struct{}

func (c *UnsubscribeCommand) Execute(args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("Usage: unsubscribe owner/repo")
    }
    repo := args[0]
    err := repository.Unsubscribe(repo)
    if err != nil {
        return err
    }
    fmt.Println("Successfully unsubscribed from:", repo)
    return nil
}
