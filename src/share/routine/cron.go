package routine

import (
	log "PI6/share/log"
	"time"

	"github.com/go-co-op/gocron"
)

func LaunchCronTasks() (err error) {

	s := gocron.NewScheduler(time.UTC)

	// Roda a MainRoutine a cada 3 horas.
	// ######################################################
	if _, err = s.Every(3).Hours().Do(func() {
		if err := MainRoutine(true); err != nil {
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
