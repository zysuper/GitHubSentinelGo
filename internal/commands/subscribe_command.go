package commands

import (
    "fmt"
    "github.com/zysuper/github-sentinel-go/internal/repository"
)

// SubscribeCommand 处理订阅命令
type SubscribeCommand struct{}

func (c *SubscribeCommand) Execute(args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("Usage: subscribe owner/repo")
    }
    repo := args[0]
    err := repository.Subscribe(repo)
    if err != nil {
        return err
    }
    fmt.Println("Successfully subscribed to:", repo)
    return nil
}
