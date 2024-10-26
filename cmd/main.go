package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zysuper/github-sentinel-go/internal/commands"
)

func main() {
	fmt.Println("Welcome to GitHub Sentinel Go!")
	fmt.Println("Type 'help' for a list of commands.")

	registry := commands.NewCommandRegistry()

	// 注册命令
	registry.RegisterCommand("subscribe", &commands.SubscribeCommand{})
	registry.RegisterCommand("unsubscribe", &commands.UnsubscribeCommand{})
	registry.RegisterCommand("list", &commands.ListCommand{})
	registry.RegisterCommand("update", commands.NewUpdateCommand())
	registry.RegisterHelpCommand()

	repl(registry)
}

func repl(registry *commands.CommandRegistry) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Exiting GitHub Sentinel Go. Goodbye!")
			break
		}

		err := registry.ExecuteCommand(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
