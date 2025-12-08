package cmd

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shivgitcode/gimi/helpers"
)


func GenerateCommitMsg(cmd string,args []string){
	generateCmd:=flag.NewFlagSet(cmd,flag.ExitOnError)

	generateCmd.Parse(args)

	s:=spinner.New(spinner.CharSets[33],100*time.Millisecond)
	s.Prefix="Getting your commit msg: "
	diff,err:=helpers.GetGitDiff()
	if err!=nil{
		color.Red(err.Error())
		return
	}
	
	if strings.TrimSpace(diff)==""{
		color.Yellow("No staged changes found")
		return 
	}
	s.Start()
	commitMessage,err:=helpers.OllamaConfig(diff)
	if err!=nil{
		color.Red(err.Error())
		s.Stop()
		return
	}
	s.Stop()

	color.Cyan("This your commit msg:\n")

	fmt.Println(commitMessage)


}