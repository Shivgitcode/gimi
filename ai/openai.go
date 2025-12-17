package ai

import (
	"context"
	"strings"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIBackend struct{
	ApiKey string
	Model string
}

func (o *OpenAIBackend) GenerateCommitMessage(prompt string)(string,error){
	client:= openai.NewClient(option.WithAPIKey(o.ApiKey))
	resp,err := client.Chat.Completions.New(context.TODO(),openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(prompt),
		},
		Model: o.Model,
		Temperature: openai.Float(0.2),
		MaxTokens: openai.Int(100),
		Seed: openai.Int(42),

	})
	if err!=nil{
		return "",err
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content),nil

}