package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string `mapstructure:"api_key"`
	Port   string `mapstructure:"port"`
}

var Cnf *Config
var ProcessPath string

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

	ex, err := os.Executable()
	if err != nil {
		panic("Can't get process path!\n" + err.Error())
	}
	ProcessPath = filepath.Dir(ex)

	fmt.Println("Process path: " + ProcessPath)

	path := "cache"

	if _, err := os.Stat(ProcessPath + PATH_SEPARATOR + path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(ProcessPath+PATH_SEPARATOR+path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Config loaded!")
}
