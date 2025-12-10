package helpers

import "github.com/AlecAivazis/survey/v2"

var GeminiKeyPrompt = &survey.Input{
	Message: "Enter gemini api key: ",
}

var GeminiModelPrompt = &survey.Input{
	Message: "Enter model name for gemini (default gemini-2.5-flash): ",
	Default: "gemini-2.5-flash",
}
