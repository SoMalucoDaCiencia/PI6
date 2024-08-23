package main

import (
	"PI6/src"
)

func main() {

	// err := src.Setup()
	// if err != nil {
	// 	panic(fmt.Errorf("an error occured while starting application: %v", err.Error()))
	// }

	// err = database.CheckDatabase()
	// if err != nil {
	// 	panic(fmt.Errorf("an error occured while starting application: %v", err.Error()))
	// }

	err := src.FilterCEPs()
	if err != nil {
		panic(err)
	}

	println("done...")
}
