package commands

import (
    "fmt"
    "strings"
)

// Command 是所有子命令的接口
type Command interface {
    Execute(args []string) error
}

// CommandRegistry 保存命令的映射
type CommandRegistry struct {
    commands map[string]Command
}

// NewCommandRegistry 创建一个命令注册表
func NewCommandRegistry() *CommandRegistry {
    return &CommandRegistry{
        commands: make(map[string]Command),
    }
}

// RegisterCommand 注册一个命令
func (r *CommandRegistry) RegisterCommand(name string, cmd Command) {
    r.commands[name] = cmd
}

// ExecuteCommand 根据输入执行对应命令
func (r *CommandRegistry) ExecuteCommand(input string) error {
    parts := strings.Fields(input)
    if len(parts) == 0 {
        return fmt.Errorf("Please enter a command.")
    }

    cmdName := parts[0]
    cmdArgs := parts[1:]

    cmd, exists := r.commands[cmdName]
    if !exists {
        return fmt.Errorf("Unknown command: %s", cmdName)
    }

    return cmd.Execute(cmdArgs)
}

// HelpCommand 列出所有可用的命令
type HelpCommand struct {
    registry *CommandRegistry
}

func (c *HelpCommand) Execute(args []string) error {
    fmt.Println("Available commands:")
    for name := range c.registry.commands {
        fmt.Printf("  %s\n", name)
    }
    return nil
}

// 注册 HelpCommand
func (r *CommandRegistry) RegisterHelpCommand() {
    helpCmd := &HelpCommand{registry: r}
    r.RegisterCommand("help", helpCmd)
}
