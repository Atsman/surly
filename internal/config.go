package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func configureConfigResolution() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath(".")
}

func readConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

type Config struct {
	HttpAddr   string
	RemoteAddr string
}

func createConfig() Config {
	config := Config{}
	config.HttpAddr = viper.GetString("http_addr")
	config.RemoteAddr = viper.GetString("remote_addr")
	return config
}

func InitConfig() Config {
	configureConfigResolution()
	readConfig()
	return createConfig()
}
