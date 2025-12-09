package ai

type AI interface{
	GenerateCommitMessage(prompt string) (string,error)

}

func GetBackend(name string,apiKey string)AI{
	if name=="openai"{
		return &OpenAIBackend{ApiKey: apiKey}
	}
	return &OllamaBackend{}

}