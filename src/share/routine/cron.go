package routine

import (
	"time"
	log "PI6/share/log"


	gocron "github.com/go-co-op/gocron/v2"
)

func LaunchCronTasks() (s gocron.Scheduler, err error) {

	s, err = gocron.NewScheduler()
	if err != nil {
		return s, err
	}

	_, err = s.NewJob(gocron.DurationJob(3*time.Hour), gocron.NewTask(MainRoutine))
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