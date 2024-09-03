// models/geocode.go
package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Result struct {
	Lat string `json:"lat"`
	Lng string `json:"lon"`
}

func GetCoordinates(distrito string) (string, string, error) {
	encodedDistrito := url.QueryEscape(distrito)
	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&addressdetails=1", encodedDistrito)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to get data for %s: %w", distrito, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response body for %s: %w", distrito, err)
	}

	fmt.Printf("Response Body: %s\n", body)

	var results []Result
	if err := json.Unmarshal(body, &results); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal JSON for %s: %w", distrito, err)
	}

	for _, result := range results {
		if result.Lat != "" && result.Lng != "" {
			return result.Lat, result.Lng, nil
		}
	}

	return "", "", fmt.Errorf("no valid results for distrito: %s", distrito)
}

// file, err := os.Create("misc/distritos.csv")
// if err != nil {
// 	log.Fatalf("failed to create CSV file: %v", err)
// }
// defer file.Close()

// writer := csv.NewWriter(file)
// defer writer.Flush()

// if err := writer.Write([]string{"Bairro", "Latitude", "Longitude"}); err != nil {
// 	log.Fatalf("failed to write header to CSV file: %v", err)
// }

// for _, bairro := range bairros {
// lat, lng, err := models.GetCoordinates(bairro)
// 	if err != nil {
// 		log.Printf("Error: %v", err)
// 		continue
// 	}

// 	record := []string{bairro, lat, lng}
// 	if err := writer.Write(record); err != nil {
// 		log.Fatalf("failed to write record to CSV file: %v", err)
// 	}
// }
