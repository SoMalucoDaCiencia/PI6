package src

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	mine "github.com/artking28/myGoUtils"
)

func FilterCEPs() error {
	content, err := os.Open("misc/allCEPs.csv")
	if err != nil {
		return err
	}

	records, err := csv.NewReader(content).ReadAll()
	if err != nil {
		return err
	}

	set := mine.NewSet[string]()
	list := []string{}
	for _, line := range records {
		if !set.Has(line[1]) {
			set.Add(line[1])
			list = append(list, fmt.Sprintf("%s,%s,%s", line[0], line[1], line[2]))
		}
	}

	return os.WriteFile("misc/filterCEPs.csv", []byte(strings.Join(list, "\n")), 0644)
}

// func getWeather() (models.WeatherEnum, uint8, error) {
// 	resp, err := http.Get("http://apiadvisor.climatempo.com.br/api/v1/weather/locale/3477/current?token=e07f43f4512d028ca15e3d81f0635a40")
// 	if err != nil {
// 		return models.NoneWeather, 0, err
// 	}

// 	bytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return models.NoneWeather, 0, err
// 	}

// 	return
// }
