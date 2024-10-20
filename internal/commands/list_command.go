package commands

import (
    "fmt"
    "github.com/zysuper/github-sentinel-go/internal/repository"
)

// ListCommand 处理列表命令
type ListCommand struct{}

func (c *ListCommand) Execute(args []string) error {
    repos := repository.ListSubscriptions()
    if len(repos) == 0 {
        fmt.Println("No subscriptions found.")
        return nil
    }
    fmt.Println("Subscribed repositories:")
    for _, repo := range repos {
        fmt.Println("- " + repo)
    }
    return nil
}
