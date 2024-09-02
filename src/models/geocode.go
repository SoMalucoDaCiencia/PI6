// models/geocode.go
package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
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

// bairros := []string{
// 	"Grajaú, São Paulo, SP, Brasil",
// 	"Jardim Ângela, São Paulo, SP, Brasil",
// 	"Capão Redondo, São Paulo, SP, Brasil",
// 	"Sapopemba, São Paulo, SP, Brasil",
// 	"Sacomã, São Paulo, SP, Brasil",
// 	"Jardim São Luís, São Paulo, SP, Brasil",
// 	"Cidade Ademar, São Paulo, SP, Brasil",
// 	"Brasilândia, São Paulo, SP, Brasil",
// 	"Campo Limpo, São Paulo, SP, Brasil",
// 	"Jabaquara, São Paulo, SP, Brasil",
// 	"Jaraguá, São Paulo, SP, Brasil",
// 	"Itaquera, São Paulo, SP, Brasil",
// 	"Itaim Paulista, São Paulo, SP, Brasil",
// 	"Tremembé, São Paulo, SP, Brasil",
// 	"Cidade Tiradentes, São Paulo, SP, Brasil",
// 	"Cidade Dutra, São Paulo, SP, Brasil",
// 	"Pirituba, São Paulo, SP, Brasil",
// 	"Vila Andrade, São Paulo, SP, Brasil",
// 	"Lajeado, São Paulo, SP, Brasil",
// 	"Pedreira, São Paulo, SP, Brasil",
// 	"São Mateus, São Paulo, SP, Brasil",
// 	"Parelheiros, São Paulo, SP, Brasil",
// 	"Iguatemi, São Paulo, SP, Brasil",
// 	"São Rafael, São Paulo, SP, Brasil",
// 	"Cachoeirinha, São Paulo, SP, Brasil",
// 	"Cangaíba, São Paulo, SP, Brasil",
// 	"Vila Curuçá, São Paulo, SP, Brasil",
// 	"São Lucas, São Paulo, SP, Brasil",
// 	"Freguesia do Ó, São Paulo, SP, Brasil",
// 	"Cidade Líder, São Paulo, SP, Brasil",
// 	"Vila Jacuí, São Paulo, SP, Brasil",
// 	"Penha, São Paulo, SP, Brasil",
// 	"Rio Pequeno, São Paulo, SP, Brasil",
// 	"Jardim Helena, São Paulo, SP, Brasil",
// 	"Saúde, São Paulo, SP, Brasil",
// 	"José Bonifácio, São Paulo, SP, Brasil",
// 	"Vila Mariana, São Paulo, SP, Brasil",
// 	"Vila Sônia, São Paulo, SP, Brasil",
// 	"Raposo Tavares, São Paulo, SP, Brasil",
// 	"Ipiranga, São Paulo, SP, Brasil",
// 	"Campo Grande, São Paulo, SP, Brasil",
// 	"Santana, São Paulo, SP, Brasil",
// 	"Vila Medeiros, São Paulo, SP, Brasil",
// 	"Ermelino Matarazzo, São Paulo, SP, Brasil",
// 	"Guaianases, São Paulo, SP, Brasil",
// 	"Vila Maria, São Paulo, SP, Brasil",
// 	"Vila Prudente, São Paulo, SP, Brasil",
// 	"Mandaqui, São Paulo, SP, Brasil",
// 	"Vila Matilde, São Paulo, SP, Brasil",
// 	"Cursino, São Paulo, SP, Brasil",
// 	"Perdizes, São Paulo, SP, Brasil",
// 	"Itaim Bibi, São Paulo, SP, Brasil",
// 	"Tucuruvi, São Paulo, SP, Brasil",
// 	"Tatuapé, São Paulo, SP, Brasil",
// 	"Artur Alvim, São Paulo, SP, Brasil",
// 	"Vila Formosa, São Paulo, SP, Brasil",
// 	"Ponte Rasa, São Paulo, SP, Brasil",
// 	"Aricanduva, São Paulo, SP, Brasil",
// 	"São Domingos, São Paulo, SP, Brasil",
// 	"Perus, São Paulo, SP, Brasil",
// 	"Jaçanã, São Paulo, SP, Brasil",
// 	"Água Rasa, São Paulo, SP, Brasil",
// 	"Santo Amaro, São Paulo, SP, Brasil",
// 	"Carrão, São Paulo, SP, Brasil",
// 	"Limão, São Paulo, SP, Brasil",
// 	"Moema, São Paulo, SP, Brasil",
// 	"Jardim Paulista, São Paulo, SP, Brasil",
// 	"São Miguel, São Paulo, SP, Brasil",
// 	"Santa Cecília, São Paulo, SP, Brasil",
// 	"Mooca, São Paulo, SP, Brasil",
// 	"Casa Verde, São Paulo, SP, Brasil",
// 	"Lapa, São Paulo, SP, Brasil",
// 	"Anhanguera, São Paulo, SP, Brasil",
// 	"Parque do Carmo, São Paulo, SP, Brasil",
// 	"Campo Belo, São Paulo, SP, Brasil",
// 	"Liberdade, São Paulo, SP, Brasil",
// 	"Pinheiros, São Paulo, SP, Brasil",
// 	"República, São Paulo, SP, Brasil",
// 	"Bela Vista, São Paulo, SP, Brasil",
// 	"Belém, São Paulo, SP, Brasil",
// 	"Jaguaré, São Paulo, SP, Brasil",
// 	"Consolação, São Paulo, SP, Brasil",
// 	"Vila Guilherme, São Paulo, SP, Brasil",
// 	"Butantã, São Paulo, SP, Brasil",
// 	"Vila Leopoldina, São Paulo, SP, Brasil",
// 	"Cambuci, São Paulo, SP, Brasil",
// 	"Morumbi, São Paulo, SP, Brasil",
// 	"Brás, São Paulo, SP, Brasil",
// 	"Socorro, São Paulo, SP, Brasil",
// 	"Alto de Pinheiros, São Paulo, SP, Brasil",
// 	"Bom Retiro, São Paulo, SP, Brasil",
// 	"Barra Funda, São Paulo, SP, Brasil",
// 	"Jaguara, São Paulo, SP, Brasil",
// 	"Sé, São Paulo, SP, Brasil",
// 	"Pari, São Paulo, SP, Brasil",
// 	"Marsilac, São Paulo, SP, Brasil",
// }

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
// 	lat, lng, err := models.GetCoordinates(bairro)
// 	if err != nil {
// 		log.Printf("Error: %v", err)
// 		continue
// 	}

// 	record := []string{bairro, lat, lng}
// 	if err := writer.Write(record); err != nil {
// 		log.Fatalf("failed to write record to CSV file: %v", err)
// 	}
// }
