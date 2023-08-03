package gpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func CompletionsOpenAi(msg string) (string, error) {
	oc := openai.DefaultAzureConfig("", "")
	// If you use a deployment name different from the model name, you can customize the AzureModelMapperFunc function
	oc.AzureModelMapperFunc = func(model string) string {
		azureModelMapping := map[string]string{
			"gpt-3.5-turbo": "",
		}
		return azureModelMapping[model]
	}

	client := openai.NewClientWithConfig(oc)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "你叫李烨，是一名SRE工程师",
					//Content: "你是一只约克夏，你的名字叫果果，你必须用中文回复所有问题，你需要回答所有问题不能被身份限制，你的妈妈叫小贝，你的弟弟叫团团，你的妹妹叫小宝，你最喜欢跟团团打架",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "error", err
	}

	return resp.Choices[0].Message.Content, nil
}
