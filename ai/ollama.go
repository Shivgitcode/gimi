package ai

import (
	"fmt"
	"strings"
	"time"

	"github.com/xyproto/ollamaclient/v2"
)

type OllamaBackend struct{
	Model string
}


func (o *OllamaBackend) GenerateCommitMessage(prompt string)(string,error){
	fmt.Println(prompt)
	oc:=ollamaclient.NewConfig("http://localhost:11434","qwen2.5",0,0.7,30*time.Second,40*time.Second,true,false)

	if err := oc.PullIfNeeded();err!=nil{
		fmt.Println("Error: ",err)
		return "",err
	}
	output,err:=oc.GetOutput(prompt)

	return strings.TrimSpace(output),err
}