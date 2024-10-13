package main

import (
	"PI6/database"
	"PI6/share"
	"PI6/share/log"
	"PI6/share/routine"
	"fmt"
)

func main() {

	// Cria as variáveis de ambiente.
	// ######################################################
	err := share.Setup()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	log.WriteLog(log.LogOk, "environment variables set up successfully", "")

	// Checa o banco de dados e cria os modelos se necessário.
	// ######################################################
	err = database.CheckDatabase()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	log.WriteLog(log.LogOk, "SQL Server is up and running", "")

	// Cria as rotinas pra rodar a cada 3 horas.
	// ######################################################
	err = routine.LaunchCronTasks()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}
	log.WriteLog(log.LogOk, "cron launched with success", "")

	// Bloqueie indefinidamente
	// ######################################################
	<-make(chan bool)
}
