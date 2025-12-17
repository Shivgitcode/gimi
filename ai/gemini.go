package ai

import (
	"context"
	"strings"

	"google.golang.org/genai"
)

type GeminiBackend struct{
	Model string
	ApiKey string
}


func(o *GeminiBackend) GenerateCommitMessage(prompt string)(string,error){
	ctx:= context.Background()
	client,err:= genai.NewClient(ctx,&genai.ClientConfig{
		APIKey: o.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err!=nil{
		return "",err
	}
	parts:=[]*genai.Part{
		{Text: prompt},
	}

	contents:=[]*genai.Content{{Parts: parts}}
	temp := float32(0.2)
	seed := int32(40)
	config:=&genai.GenerateContentConfig{
		MaxOutputTokens: 500,
		Temperature: &temp,
		Seed:&seed,
	}

	result,err:=client.Models.GenerateContent(ctx,o.Model,contents,config)
	if err!=nil{
		return "",err
	}




	return strings.TrimSpace(string(result.Text())),nil

}