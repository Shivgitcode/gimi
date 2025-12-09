package ai

import (
	"context"
	"strings"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIBackend struct{
	ApiKey string
}

func (o *OpenAIBackend) GenerateCommitMessage(prompt string)(string,error){
	client:= openai.NewClient(option.WithAPIKey(o.ApiKey))
	resp,err := client.Chat.Completions.New(context.TODO(),openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4_1,
	})
	if err!=nil{
		return "",err
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content),nil

}