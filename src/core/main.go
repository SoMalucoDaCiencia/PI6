package main

import (
	"PI6/src"
	"PI6/src/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	OLAT := -23.7799713
	OLONG := -46.6737655

	DLAT := -23.7125278
	DLONG := -46.7687195

	url := fmt.Sprintf("https://www.google.com.br/maps/preview/directions?authuser=0&hl=pt-BR&gl=br&pb=!1m5!1s%f%%2C%f!3m2!3d%f!4d%f!6e2!1m5!1s%f%%2C%f!3m2!3d%f!4d%f!6e2!3m12!1m3!1d290274.0087345833!2d-46.85383994709557!3d-23.63742370585772!2m3!1f0!2f0!3f0!3m2!1i859!2i953!4f13.1!6m38!1m2!18b1!30b1!2m3!5m1!6e2!20e3!6m12!4b1!49b1!66b1!74i150000!75b1!85b1!89b1!91b1!114b1!149b1!166f1.35!196b1!10b1!12b1!13b1!14b1!16b1!17m2!3e1!3e1!20m6!1e0!2e3!5e2!6b1!8b1!14b1!46m1!1b0!94b1!8m1!1e0!15m4!1siinXZtn9D_XK1sQP5_3n8Qc!4m1!2i10317!7e81!20m28!1m6!1m2!1i0!2i0!2m2!1i530!2i953!1m6!1m2!1i809!2i0!2m2!1i859!2i953!1m6!1m2!1i0!2i0!2m2!1i859!2i20!1m6!1m2!1i0!2i933!2m2!1i859!2i953!27b1!28m0!40i704!47m1!8b1", OLAT, OLONG, OLAT, OLONG, DLAT, DLONG, DLAT, DLONG)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("dnt", "1")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.google.com/")
	req.Header.Set("sec-ch-ua", `^\^Not`)
	req.Header.Set("cookie", `S=billing-ui-v3=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU:billing-ui-v3-efe=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU; receive-cookie-deprecation=1; HSID=AOO_sy-FmiCLRHi3i; SSID=ALvioYklsTgxGHHS2; APISID=9TaUF1BxDCyDv0gR/Aoj0Wl4WhUKNJFpFg; SAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-1PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-3PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; SID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BiadvdwGMYhGSJF6v6zUvFwACgYKAW0SARESFQHGX2MiyIgn3KDFeJA3YZGhD8ihhxoVAUF8yKpsZWEmUrj_AnE8dXeui5uY0076; __Secure-1PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BaAdle1TIFKsylBDp2NJwZAACgYKAR8SARESFQHGX2MioJKkn_9gGMcHm4eQ_Yud0hoVAUF8yKojKjMLjR_PC9FgMB33n2FY0076; __Secure-3PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BL6D2zb2ahaKjZntLSynErAACgYKAQESARESFQHGX2MiRN17UwaTTiZQsMByCOYgsBoVAUF8yKqe6uGmwERQcq9s7AjOsfwo0076; SEARCH_SAMESITE=CgQI6JsB; OGPC=19031986-1:; AEC=AVYB7cq9-ILZK5QjVxddGunI6ECPelIx7n_hIjYcQ39pXtN-MAeV66KLL44; NID=516=A83enFQcl1QUFtvWJ8diSkBi0XdhIfwylQGoDbUya2ZuBeKKW--PXv1ZsdEhI4EOKKZHEFaaW5hKN1VbPdABay9f2f18nNojdq_6DqprJZDdE-1Q3u8xhvIvyWvaPHmn9cfP0owY9wNHabC_C9CKfjbankvO2cOhm3fGpQKDtFwlZCoJIU5b3mQrhrdBImwx-zUCCAnAI0tSmL9k7r3JluZ93Gmw2quLfBUgJWDYSpx2pFSogqfkZAlbXq2eHB9UUTL3MyRAxO1nbfrGDsF4LysVkOA606bm0jKh8SKGC3UDqUyPNqbMve3k_DtX4dZ-UsQmrai0C9BT1Hy4TbrFWXI7bLnj7JFaLqdoSAzD_OjRWadzgrDDkDsHO-0icufSvOLWGgdnjT7wCVdp3HMqxZS9aXnFLrYPOyKjQf4UrIen0QZEp-645SC9jcKbujNfPTSot4CAMZeu0_jhiA8FfC7RZny8EtPrQy0tTmJ_YoIj5oLoH1bWTrLajddFBUglChSfJ1rwCQ-FB4A3u_Iop9nB6Etso2Afab8; DV=s9QN8ddaB_lSoH49I6th8VpDpUZ3F9ku3TJpNEbf1gAAAJBbb6P7S5CgeAAAACTWtiY5pKI3JgAAAILUrVF03KrXCwAAAA; __Secure-1PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; __Secure-3PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; SIDCC=AKEyXzVkqwOtb_ZeisFTG16QhOk__PpmW5x40Oo3tsGi8X2RU60-Q5gZglrd6PiaV_3hDcphtw; __Secure-1PSIDCC=AKEyXzWl-d7vAAFHftQtOo7hNqpghJpygezDpJzsW3h5gEgxCT6nReyMiZy1hUIfyZLh8vu--w; __Secure-3PSIDCC=AKEyXzWI-lehfd3RXGL74LvQ0YBbCh3QJPcj18G4AjFqE492NVIKpP9agrwYbpUTGTPracbnNw`)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}

	cleanedBody := cleanJSON(body)

	file, err := os.Create("misc/response.json")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(cleanedBody))
	if err != nil {
		log.Fatalf("Erro ao escrever no arquivo: %v", err)
	}

	fmt.Println("Resposta salva em response.json")

	file, err = os.Open("misc/response.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	var jsonData interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jsonData); err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	data, ok := jsonData.([]interface{})
	if !ok {
		log.Fatal("Estrutura de dados JSON inesperada")
	}

	nestedData, err := getNestedData(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	data, ok = nestedData.([]interface{})
	if !ok {
		log.Fatal("Estrutura esperada não encontrada na posição 0")
	}

	nestedData, err = getNestedData(data, 20)
	if err != nil {
		log.Fatal(err)
	}

	data, ok = nestedData.([]interface{})
	if !ok {
		log.Fatal("Estrutura esperada não encontrada na posição 20")
	}

	if len(data) < 3 {
		log.Fatal("Não há dados suficientes na posição 20")
	}

	carTime, ok := data[0].([]interface{})
	if !ok || len(carTime) < 3 {
		log.Fatal("Estrutura esperada não encontrada para tempo de carro")
	}

	carTimeNested, ok := carTime[2].([]interface{})
	if !ok || len(carTimeNested) < 2 {
		log.Fatal("Estrutura esperada não encontrada dentro do tempo de carro")
	}

	carTimeValue := carTimeNested[1]

	publicTransportTime, ok := data[2].([]interface{})
	if !ok || len(publicTransportTime) < 3 {
		log.Fatal("Estrutura esperada não encontrada para tempo de transporte público")
	}

	publicTransportTimeNested, ok := publicTransportTime[2].([]interface{})
	if !ok || len(publicTransportTimeNested) < 2 {
		log.Fatal("Estrutura esperada não encontrada dentro do tempo de transporte público")
	}

	publicTransportTimeValue := publicTransportTimeNested[1]

	fmt.Printf("Tempo de carro: %v\n", carTimeValue)
	fmt.Printf("Tempo de transporte público: %v\n", publicTransportTimeValue)
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
