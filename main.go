package main

import (
	"fmt"
	"github.com/spf13/viper"
	"rest-gorm/application"
)

func main() {
	config()
	application.NewContainer()
}

func config() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
