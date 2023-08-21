package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Global *viper.Viper

func init() {
	Global = viper.New()

	configEnv := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if configEnv == "" {
		configEnv = "development"
	}
	Global.SetConfigName(configEnv)

	configFilePath := strings.ToLower(os.Getenv("CONFIG_FILE_PATH"))
	if configFilePath == "" {
		configFilePath = "./config"
	}
	Global.AddConfigPath(configFilePath)

	configFileType := strings.ToLower(os.Getenv("CONFIG_FILE_TYPE"))
	if configFileType == "" {
		configFileType = "yaml"
	}
	Global.SetConfigType(configFileType)

	Global.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Global.AutomaticEnv()

	if err := Global.ReadInConfig(); err != nil {
		log.Printf("Error unable to load config file: %v", err)
	}
}
