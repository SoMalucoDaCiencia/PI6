package src

import "github.com/spf13/viper"

func Setup() error {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}
