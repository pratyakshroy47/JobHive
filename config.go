package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	MongoURI       string
	LogLevel       string
	PlaygroundPath string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		Port:           viper.GetString("port"),
		MongoURI:       viper.GetString("mongo.uri"),
		LogLevel:       viper.GetString("log.level"),
		PlaygroundPath: viper.GetString("playground.path"),
	}

	return config, nil
}
