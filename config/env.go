package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go_starter/logs"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ERROR_READING_CONFIG_FILE", err)
		logs.Error("ERROR_READING_CONFIG_FILE")
		return
	}
	fmt.Println("SUCCESS_READING_CONFIG_FILE")
}

func GetEnv(key, defaultValue string) string {
	readValue := viper.GetString(key)
	if readValue == "" {
		return defaultValue
	}
	return readValue
}

func Env(key string) string {
	readValue := viper.GetString(key)
	return readValue
}
