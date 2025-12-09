package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)



func InitCommand(args []string){
	initCmd:=flag.NewFlagSet("init",flag.ExitOnError)

	initCmd.Parse(args)
	homedir,err:=os.UserHomeDir()
	if err!=nil{
		fmt.Println("error finding homedir")
		return
	}
	configDirPath:=filepath.Join(homedir,".gimi")
	configFilePath:=filepath.Join(configDirPath,"config.json")
	os.MkdirAll(configDirPath,0755)
	viper.SetConfigType("json")
	viper.SetConfigFile(configFilePath)
	result:=""
	prompt:=&survey.Select{
		Message: "Select Backend",
		Options: []string{
			"openai",
			"ollama",
		},
	}

	
	survey.AskOne(prompt,&result)
	keyPrompt:=&survey.Input{
		Message: "Enter Your Openai Api Key: ",
	}
	modelPrompt:=&survey.Input{
		Message:"Enter which model you want to use (default gpt_4.1): ",
		Default: "gpt_4.1",
	}
	apiKey:=""
	model:="4.1"


	
	if result=="openai"{
		if viper.GetString("apiKey")==""{
			survey.AskOne(keyPrompt,&apiKey)
			viper.Set("apiKey",apiKey)
		}
		if viper.GetString("model")==""{
			survey.AskOne(modelPrompt,&model)
			viper.Set("model",model)

		}
	}

	viper.Set("Backend",result)
	if err:=viper.WriteConfigAs(configFilePath);err!=nil{
		color.Red("Error writing inside file")
		return
	}

	color.Cyan("✔️ Configuration Complete")



}