package src

import (
	src "PI6"
	log "PI6/share/log"
	"time"

	gocron "github.com/go-co-op/gocron/v2"
	"github.com/spf13/viper"
)

func Setup() error {
	viper.AutomaticEnv()
	if !viper.GetBool("DOCKER") {
		viper.SetConfigFile("../.env")
		return viper.ReadInConfig()
	}
	return nil
}

func LaunchCronTasks() (s gocron.Scheduler, err error) {

	s, err = gocron.NewScheduler()
	if err != nil {
		return s, err
	}

	_, err = s.NewJob(gocron.DurationJob(3*time.Hour), gocron.NewTask(src.MainRoutine))
	if err != nil {
		return s, err
	}

	c := len(s.Jobs())
	if c <= 0 {
		log.WriteLog(log.LogInfo, "Cron has no task", "")
		return s, nil
	}

	s.Start()
	return s, nil
}
