package main

import (
	"PI6/src"
	"PI6/src/database"
	"fmt"
)

func main() {

	err := src.Setup()
	if err != nil {
		panic(fmt.Errorf("an error occured while starting application: %v", err.Error()))
	}

	err = database.CheckDatabase()
	if err != nil {
		panic(fmt.Errorf("an error occured while starting application: %v", err.Error()))
	}

	println("done...")
}
