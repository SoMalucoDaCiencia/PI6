package routine

import (
	"PI6/database"
	"PI6/models"
	"PI6/models/entity"
	log "PI6/share/log"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

func worker(wg *sync.WaitGroup, atg models.AppleTokenGetter, from *models.Address, to []*models.Address) error {
	defer wg.Done()

	db, err := database.GetConn()
	if err != nil {
		return err
	}

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
		destinies = append(destinies, models.NewGeoPos(add.Lat, add.Long))
	}

	result, err := entity.ExtractRegister(atg.AccessToken, origin, destinies)
	if err != nil {
		err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
		log.WriteLog(log.LogErr, err.Error(), "apple")
		panic(err.Error())
	}

	return db.Create(result).Error
}

func MainRoutine() error {
	start := time.Now()

	file, err := os.OpenFile("../misc/distritos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("err", err.Error())
		panic(err)
	}
	defer file.Close()

	var all []*models.Address
	err = gocsv.UnmarshalFile(file, &all)
	if err != nil {
		return err
	}

	db, err := database.GetConn()
	if err != nil {
		return err
	}

	var tokens []entity.Token
	db = db.Model(&entity.Token{}).Where("createdAt >= DATEADD(week, -1, GETDATE())").First(&tokens)
	if db.Error != nil {
		return db.Error
	}

	atg, err := entity.ExtractToken(tokens[0].Value)
	if db.Error != nil {
		return db.Error
	}

	var wg sync.WaitGroup
	for i, ads0 := range all {
		vec := append(all[:i], all[i+1:]...)
		if !atg.IsValid() {
			err := atg.Renew()
			if err != nil {
				return err
			}
		}
		for i := 0; i < len(vec); i += 10 {
			wg.Add(1)
			if i+9 < len(vec) {
				go func() {
					err = worker(&wg, atg, ads0, vec[i:i+9])
				}()
				if err != nil {
					return err
				}
				continue
			}
			go func() {
				err = worker(&wg, atg, ads0, vec[i:])
			}()
			if err != nil {
				return err
			}
		}
	}

	wg.Wait()
	log.WriteLog(log.LogOk, "routine completed in "+time.Since(start).String(), "")
	return nil
}
