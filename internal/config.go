package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	HttpAddr string
}

func InitConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := Config{}
	config.HttpAddr = viper.GetString("http_addr")

	return config
}
