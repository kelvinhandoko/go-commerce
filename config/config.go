package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	var cfg Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./files/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	return cfg
}
