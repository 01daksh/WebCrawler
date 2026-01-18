package configs

import (
	"log"

	"github.com/spf13/viper"
)

func InitializeConfig() error {
	// 1. Set the name of the config file (without extension)
	viper.SetConfigName("config")

	// 2. Set the type of the config file
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if you want defaults or env vars to work
			log.Println("No config file found, using defaults or env vars")
		} else {
			// Config file was found but another error occurred
			log.Fatalf("Fatal error reading config file: %s \n", err)
			panic(err)
		}
	}

	return nil
}

func GetServerAddress() string {
	return viper.GetString("server.httpServerAddress")
}
