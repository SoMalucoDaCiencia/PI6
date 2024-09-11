package main

import (
	"PI6/database"
	share "PI6/share"
	log "PI6/share/log"
	"fmt"
)

func main() {

	err := share.Setup()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	log.WriteLog(log.LogOk, "environment variables set up successfully", "")

	err = database.CheckDatabase()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	log.WriteLog(log.LogOk, "SQL Server is up and running", "")

	s, err := share.LaunchCronTasks()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	if len(s.Jobs()) > 0 {
		log.WriteLog(log.LogOk, "cron launched with success", "")
	}

	// Bloqueie indefinidamente
	<-make(chan struct{})
}
