package src

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
)

func Setup() error {
	viper.AutomaticEnv()
	viper.SetConfigFile("../.env")
	return viper.ReadInConfig()
}

func LaunchCronTasks() (err error) {

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(2).Hours().Do(func() (ret []error) {

		return ret
	})
	if err != nil {
		return err
	}

	if s.Len() <= 0 {
		WriteLog(LogInfo, "Cron has no task", "")
		return nil
	}
	s.StartAsync()
	return nil
}
