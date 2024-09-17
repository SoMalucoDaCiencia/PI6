package routine

import (
	"os"
	"sync"
	"time"
	"PI6/models"
	log "PI6/share/log"

	"github.com/gocarina/gocsv"
)

func worker(wg *sync.WaitGroup, from Address, to []*Address) error {
	defer wg.Done()
	// Simula uma tarefa que demora 1 segundo
	// 
	// origin := models.NewGeoPos(37.331423, -122.030503)
	// destinies := []models.GeoPos{
	// 	models.NewGeoPos(37.32556561130194, -121.94635203581443),
	// 	models.NewGeoPos(37.44176585512703, -122.17259315798667),
	// }
	
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

	origin := models.NewGeoPos(from.Lat, from.Long)
	var destinies []models.GeoPos
	for _, add := range to {
		destinies = append(destinies, models.NewGeoPos(*add.Lat, *add.Long))
	}
	
	all, err := models.ExtractRegister(origin, destinies)
	if err != nil {
		err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
		src.WriteLog(src.LogErr, err.Error(), "apple")
		panic(err.Error())
	}

	println()
	for _, r := range all {
		println(r.AsString())
	}

	time.Sleep(1 * time.Second)
	return nil
}

func MainRoutine() {
	start := time.Now()

	file, err := os.OpenFile("misc/distritos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	var all []*models.Address
	err = gocsv.UnmarshalFile(file, &all); 
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, ads0 := range all {
		wg.Add(1)
		for i;=0; i<len(all); i+=10 {
			if i+9 < len(all) {
				go worker(&wg, all[i:i+9])
				continue
			}
			go worker(&wg, all[i:])
		}
		
	}

	wg.Wait()
	log.WriteLog(log.LogOk, "routine completed in "+time.Since(start).String(), "")
}