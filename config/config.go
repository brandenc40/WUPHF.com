package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")

	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetConfigName("secrets.yaml")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Error merging config file, %s", err)
	}
}
