package helpers

import (
	"os"
	"path/filepath"
	"fmt"

	"github.com/spf13/viper"
)


func InitConfig(){
	homedir,_:=os.UserHomeDir()
	configFilePath:=filepath.Join(homedir,".gimi")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configFilePath)

	if err:=viper.ReadInConfig();err!=nil{
		panic("could not read config: "+err.Error())
	}
}

func FilePath()string{
	homedir,_:=os.UserHomeDir()
	configFilePath:=filepath.Join(homedir,".gimi","config.json")

	return configFilePath

}

func PrintList(){
	fmt.Println("List Of Available Commands")
	fmt.Println("		init - to generate a config where you can add your openai")
	fmt.Println("		generate - to generate the commit message")
	fmt.Println("       reset - to reset the config for gimi")
	fmt.Println("       help or -h - to see list of available commands")

}