package config

import (
	"fmt"
	"os"

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
		fmt.Println("No secrets.yaml file found, the app will look for env variables instead")
	} else {
		os.Setenv("TWILIO_ACCOUNT_SID", viper.GetString("twilio.account_sid"))
		os.Setenv("TWILIO_AUTH_TOKEN", viper.GetString("twilio.auth_token"))
		os.Setenv("TWILIO_PHONE_NUMBER", viper.GetString("twilio.phone_number"))
	}

	return nil
}
