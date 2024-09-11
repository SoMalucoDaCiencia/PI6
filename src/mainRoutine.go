package src

import (
	"sync"
	"time"

	log "PI6/share/log"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	// Simula uma tarefa que demora 1 segundo
	//
	//
	// origin := models.NewGeoPos(37.331423, -122.030503)
	// destinies := []models.GeoPos{
	// 	models.NewGeoPos(37.32556561130194, -121.94635203581443),
	// 	models.NewGeoPos(37.44176585512703, -122.17259315798667),
	// }
	//
	// all, err := models.ExtractRegister(origin, destinies)
	// if err != nil {
	// 	err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
	// 	src.WriteLog(src.LogErr, err.Error(), "apple")
	// 	panic(err.Error())
	// }

	// println()
	// for _, r := range all {
	// 	println(r.AsString())
	// }
	time.Sleep(1 * time.Second)
}

func MainRoutine() {
	start := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	log.WriteLog(log.LogOk, "routine completed in "+time.Since(start).String(), "")
}
