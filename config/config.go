package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName("./config/dev/config")
	config.SetConfigType("yaml")
	if err:=config.ReadInConfig();err!=nil{
		panic(err)
	}
}
