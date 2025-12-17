package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/shivgitcode/gimi/helpers"
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
	
	survey.AskOne(helpers.BackendPrompt,&result)
	keyPrompt:=&survey.Input{
		Message: "Enter Your Openai Api Key: ",
	}
	openaiModelPrompt:=&survey.Input{
		Message:"Enter which model you want to use (default gpt_4.1): ",
		Default: "gpt-4.1",
	}
	apiKey:=""
	model:=""


	
	switch result{
	case "openai":
		if viper.GetString("apiKey")==""{
			survey.AskOne(keyPrompt,&apiKey)
			viper.Set("apiKey",apiKey)
		}
		if viper.GetString("model")==""{
			survey.AskOne(openaiModelPrompt,&model)
			viper.Set("model",model)
		}
	case "gemini":
		if viper.GetString("apiKey")==""{
			survey.AskOne(helpers.GeminiKeyPrompt,&apiKey)
			viper.Set("apiKey",apiKey)
		}
		if viper.GetString("model")==""{
			survey.AskOne(helpers.GeminiModelPrompt,&model)
			viper.Set("model",model)

		}
	}

	viper.Set("backend",result)
	if err:=viper.WriteConfigAs(configFilePath);err!=nil{
		color.Red("Error writing inside file")
		return
	}

	color.Cyan("✔️ Configuration Complete")



}