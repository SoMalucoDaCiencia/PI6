package main

import (
	"PI6/database"
	"PI6/models"
	src "PI6/share"
	"fmt"
)

func main() {

	err := src.Setup()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}

	err = database.CheckDatabase()
	if err != nil {
		panic(fmt.Errorf("an error occurred while starting application: %v", err.Error()))
	}

	pos1 := models.NewGeoPos(37.331423, -122.030503)
	pos2 := []models.GeoPos{
		models.NewGeoPos(37.32556561130194, -121.94635203581443),
		models.NewGeoPos(37.44176585512703, -122.17259315798667),
	}

	all, err := models.ExtractRegister(pos1, pos2)
	if err != nil {
		panic(fmt.Errorf("an error occurred while extracting register: %v", err.Error()))
	}

	println()
	for _, r := range all {
		println(r.AsString())
	}
}
