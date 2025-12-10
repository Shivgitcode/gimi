package ai

type AI interface{
	GenerateCommitMessage(prompt string) (string,error)
}

func GetBackend(name string,apiKey string,model ...string)AI{
	switch name{
	case "openai":
		return &OpenAIBackend{ApiKey: apiKey,Model: model[0]}
	case"gemini":
		return &GeminiBackend{ApiKey: apiKey,Model: model[0]}
	}
	return &OllamaBackend{}

}