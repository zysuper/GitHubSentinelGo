package llm

import (
	"context"
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

func (llm *LLMClient) GenerateReport(content string) (string, error) {
	resp, err := llm.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Summarize the following GitHub issues and pull requests in a formal report format."},
			{Role: "user", Content: content},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
