package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shivgitcode/gimi/ai"
	"github.com/shivgitcode/gimi/helpers"
	"github.com/spf13/viper"
)


func GenerateCommand(cmd string,args []string){
	generateCmd:=flag.NewFlagSet(cmd,flag.ExitOnError)
	stagecmd:=exec.Command("git","add",".")

	err:=stagecmd.Run()

	if err!=nil{
		color.Red("cannot stage your changes")
	}
	

	if _,err:=os.Stat(helpers.FilePath()); err!=nil{
		color.Red("cannot generate file path do not exist use gimi init to initialize config")
		return
	}

	backendVar:=viper.GetString("backend")
	apiKey:=viper.GetString("apiKey")
	model:=viper.GetString("model")


	generateCmd.Parse(args)



	switch backendVar{
		case "openai":
			if apiKey==""{
				color.Yellow("export your own openai api key to use open ai backend\nexport OPENAI_API_KEY=sk-xxx\n else you can use ollama which is free and much faster")
				return
			}
			if model==""{
				color.Yellow("Cannot find model")
				return
			}
		case "gemini":
			if apiKey==""{
				color.Yellow("export your own openai api key to use open ai backend\nexport OPENAI_API_KEY=sk-xxx\n else you can use ollama which is free and much faster")
				return
			}
			if model==""{
				color.Yellow("Cannot find model")
				return
			}
	}



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
	prompt:=helpers.GitDiffPrompt(diff)

	backend:=ai.GetBackend(backendVar,apiKey,model)

	s.Start()
	starTime:=time.Now()
	done:=make(chan bool)
	go func ()  {
		ticker:=time.NewTicker(100*time.Millisecond)
		defer ticker.Stop()
		for{
			select{
			case<-done:
				return
			case<-ticker.C:
				elapsed:=time.Since(starTime)
				s.Suffix=fmt.Sprintf(" :%v",elapsed.Round(10*time.Millisecond))
			}
		}
	}()

	commitMessage,err:=backend.GenerateCommitMessage(prompt)
	if err!=nil{
		color.Red(err.Error())
		s.Stop()
		return
	}
	close(done)
	s.Stop()

	color.Cyan("This your commit msg :\n")

	fmt.Println(commitMessage)

	color.RGB(0, 255, 255).Printf("Processing complete %v\n",time.Since(starTime).Round(100*time.Millisecond))



}