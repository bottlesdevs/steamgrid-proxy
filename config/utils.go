package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string `mapstructure:"api_key"`
}

var Cnf *Config

func init() {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		panic("Config error!")
	}

	err = viper.Unmarshal(&Cnf)

	if err != nil {
		panic("Config error!\n" + err.Error())
	}

	fmt.Println("Config loaded!")
}