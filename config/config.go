package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".") // for tests, look in current direcotry
	viper.SetConfigType("yaml")

	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName("secrets.yaml")
	if err := viper.MergeInConfig(); err != nil {
		return err
	}
	return nil
}
