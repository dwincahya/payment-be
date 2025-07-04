package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return v
}
