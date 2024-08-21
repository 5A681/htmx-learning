package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST   string
	DB_NAME   string
	DB_USER   string
	DB_PASS   string
	DB_PORT   int
	HTTP_PORT int
}

func NewConfig() Config {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config := Config{}
	// Get values from the .env file
	config.DB_HOST = viper.GetString("DB_HOST")
	config.DB_NAME = viper.GetString("DB_NAME")
	config.DB_USER = viper.GetString("DB_USER")
	config.DB_PASS = viper.GetString("DB_PASS")
	config.DB_PORT = viper.GetInt("DB_PORT")
	config.HTTP_PORT = viper.GetInt("HTTP_PORT")
	return config
}
