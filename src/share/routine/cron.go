package routine

import (
	log "PI6/share/log"
	"time"

	"github.com/go-co-op/gocron"
)

func LaunchCronTasks() (err error) {

	s := gocron.NewScheduler(time.UTC)

	if _, err = s.Every(1).Minutes().Do(func() {
		if err := MainRoutine(); err != nil {
			log.WriteLog(log.LogErr, err.Error(), "database")
		}
	}); err != nil {
		println("ta aqui")
		return err
	}

	if s.Len() <= 0 {
		log.WriteLog(log.LogInfo, "Cron has no task", "")
		return nil
	}
	s.StartAsync()
	return nil
}
