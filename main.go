package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/shivgitcode/gimi/cmd"
)


func main(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading .env")
	}
	if len(os.Args)<1{
		color.Red("no command provided")
		return
	}
	subCmd:=os.Args[1];
	args:=os.Args[2:];

	switch(subCmd){
	case "generate":
		cmd.GenerateCommitMsg(subCmd,args)
	default:
		color.Cyan("List of available command")
		


	}

}