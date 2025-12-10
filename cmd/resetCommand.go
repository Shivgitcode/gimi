package cmd

import (
	"flag"


	"github.com/fatih/color"
	"github.com/shivgitcode/gimi/helpers"
	"github.com/spf13/viper"
)


func ResetCommand(args []string){
	resetCmd:=flag.NewFlagSet("reset",flag.ExitOnError)
	
	resetCmd.Parse(args)

	viper.Set("backend","")
	viper.Set("apiKey","")
	viper.Set("model","")
	if err:=viper.WriteConfigAs(helpers.FilePath());err!=nil{
		color.Red("⨯ cannot write in file")
		return
	}

	color.Cyan("Configuration reset successfull ✔︎")
}