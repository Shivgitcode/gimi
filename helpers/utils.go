package helpers

import (
	"os"
	"path/filepath"

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