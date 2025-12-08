package helpers

import (

	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
)


func OllamaConfig(diff string) (string,error){
	oc:=ollamaclient.New(usermodel.GetTextGenerationModel())
	// oc.Verbose=true


	output,err:=oc.GetOutput(GitDiffPrompt(diff))

	return output,err
}