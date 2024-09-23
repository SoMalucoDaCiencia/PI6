package main

import (
	"PI6/models"
	"PI6/share"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

func main() {

	file, err := os.OpenFile("../misc/distritos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("err", err.Error())
		panic(err)
	}
	defer file.Close()

	var all []*models.Address
	err = gocsv.UnmarshalFile(file, &all)
	if err != nil {
		panic(err)
	}

	for k, ad0 := range all {
		vec := make([]*models.Address, len(all))
		copy(vec[:], all[:])
		vec = append(vec[:k], vec[k+1:]...)
		u := share.FloatsAsUUID(ad0.Lat, ad0.Long)
		for i := 0; i < len(vec); i += 10 {
			if i+10 < len(vec) {
				fmt.Printf("%s => \n", u)
				for _, sub := range vec[i : i+10] {
					fmt.Printf("\t%s\n", share.FloatsAsUUID(sub.Lat, sub.Long))
				}
				continue
			}
			fmt.Printf("%s => \n", u)
			for _, sub := range vec[i:] {
				fmt.Printf("\t%s\n", share.FloatsAsUUID(sub.Lat, sub.Long))
			}
		}
	}
}
