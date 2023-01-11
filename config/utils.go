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

	var imageTypes []string = []string{"grids", "heroes", "logos", "icons"}

	for _, imageType := range imageTypes {
		if _, err := os.Stat(filepath.Join(ProcessPath, path, imageType)); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(filepath.Join(ProcessPath, path, imageType), os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
	}

	fmt.Println("Config loaded!")
}
