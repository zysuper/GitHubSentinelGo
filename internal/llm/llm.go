package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

type LLMClient struct {
	client *openai.Client
}

func NewLLMClient() *LLMClient {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	return &LLMClient{client: client}
}

func (llm *LLMClient) GenerateReport(content string, dryRun bool) (string, error) {
	sys_prompt := `请使用中文，根据以下某个 github 项目的 pull requests 和 Issues 内容，生成一份简报。
简报需要生成以下内容：
- 修复了哪些问题，以及问题的严重性
- 新增了哪些特性，功能与性能增强
- 产生了不兼容的系统改变`
	// 如果 dryRun 模式开启，将内容保存到文件而不调用 API
	if dryRun {
		filename := "dry_run_prompt.txt"
		err := os.WriteFile(filename, []byte(sys_prompt+"\n\n"+content), 0644)
		if err != nil {
			return "", fmt.Errorf("failed to save dry run prompt to file: %v", err)
		}
		return fmt.Sprintf("Dry run mode enabled. Prompt saved to %s", filename), nil
	}

	resp, err := llm.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: sys_prompt},
			{Role: "user", Content: content},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
