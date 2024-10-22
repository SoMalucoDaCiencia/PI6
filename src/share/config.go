package share

import (
	"github.com/spf13/viper"
)

func Setup() error {
	viper.AutomaticEnv()
	//if !viper.GetBool("DOCKER") {
	//viper.SetConfigFile(".env")
	//return viper.ReadInConfig()
	//}
	return nil
}
