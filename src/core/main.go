package main

import (
	"PI6/src"
	"PI6/src/database"
	"PI6/src/models"
	"fmt"
	"time"
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

	t := time.Now().UnixMilli()

	mapPath := models.NewMapPath(
		models.NewGeoPos(-23.7799713, -46.6737655),
		models.NewGeoPos(-23.7125278, -46.7687195),
	)

	_, err = mapPath.ExtractRegister()
	if err != nil {
		panic(fmt.Errorf("an error occurred while extracting register: %v", err.Error()))
	}
	println(time.Now().UnixMilli()-t, "millis")

}

func getNestedData(data []interface{}, indices ...int) (interface{}, error) {
	for _, index := range indices {
		if index >= len(data) {
			return nil, fmt.Errorf("índice %d fora dos limites", index)
		}
		var ok bool
		data, ok = data[index].([]interface{})
		if !ok {
			return nil, fmt.Errorf("estrutura esperada não encontrada no índice %d", index)
		}
	}
	return data, nil
}

func cleanJSON(body []byte) string {
	cleaned := string(body)
	for len(cleaned) > 0 && (cleaned[0] == ')' || cleaned[0] == ']' || cleaned[0] == '}' || cleaned[0] == '\'') {
		cleaned = cleaned[1:]
	}
	return cleaned
}
