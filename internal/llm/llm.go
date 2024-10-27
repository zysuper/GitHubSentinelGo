package llm

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

//go:embed report_sys_prompt.txt
var reportPrompt string

type LLMClient struct {
	client *openai.Client
}

func NewLLMClient() *LLMClient {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	return &LLMClient{client: client}
}

func (llm *LLMClient) GenerateReport(content string, dryRun bool) (string, error) {
	// 如果 dryRun 模式开启，将内容保存到文件而不调用 API
	if dryRun {
		filename := "dry_run_prompt.txt"
		err := os.WriteFile(filename, []byte(reportPrompt+"\n\n"+content), 0644)
		if err != nil {
			return "", fmt.Errorf("failed to save dry run prompt to file: %v", err)
		}
		return fmt.Sprintf("Dry run mode enabled. Prompt saved to %s", filename), nil
	}

	resp, err := llm.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: reportPrompt},
			{Role: "user", Content: content},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
