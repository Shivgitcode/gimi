package main

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/shivgitcode/gimi/cmd"
	"github.com/shivgitcode/gimi/helpers"
)


func main(){
	if _,err:=os.Stat(helpers.FilePath());err==nil{
			helpers.InitConfig()
	}
	
	if len(os.Args)<=1{
		fmt.Println("Welcome to GIMI , one click cli command to generate commit message using ai")
		fmt.Println("List Of Available Commands")
		fmt.Println("		init - to generate a config where you can add your openai")
		fmt.Println("		generate - to generate the commit message")
		return
	}
	
	subCmd:=os.Args[1];
	args:=os.Args[2:];

	switch(subCmd){
	case "init":
		cmd.InitCommand(args)
	case "generate":
		cmd.GenerateCommitMsg(subCmd,args)
	default:
		color.Cyan("List of available command")
		fmt.Println("	init - to generate a config where you can add your openai")
		fmt.Println("	generate - to generate the commit message")
	}

}