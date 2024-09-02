package main

import (
	"PI6/src"
	"PI6/src/database"
	"fmt"
	"log"
	"net/http"
	"time"
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

	url := "https://www.google.com/maps/preview/directions?authuser=0&hl=pt-BR&gl=br&pb=!1m2!1sCasa%2Bda%2BPizza%2B-%2BAvenida%2BAtl%C3%A2ntica%2B-%2BInterlagos%2C%2BS%C3%A3o%2BPaulo%2B-%2BSP!2s0x94ce4e31947dd573%3A0x209bb115e9bc5236!1m5!1sPDKBikes%2B-%2BR.%2Bda%2BPaz%2C%2B2069%2B-%2BCh%C3%A1cara%2BSanto%2BAnt%C3%B4nio%2B(Zona%2BSul)%2C%2BS%C3%A3o%2BPaulo%2B-%2BSP%2C%2B04713-002!2s0x94ce512254a963bb%3A0x6230b9b9d57f824d!3m2!3d-23.6262784!4d-46.707612499999996!3m12!1m3!1d4497.841894012286!2d-46.710161123786264!3d-23.626237263980634!2m3!1f0!2f0!3f0!3m2!1i1920!2i945!4f13.1!6m24!1m2!18b1!30b1!2m3!5m1!6e2!20e3!10b1!12b1!13b1!14b1!16b1!17m1!3e1!20m6!1e0!2e3!5e2!6b1!8b1!14b1!46m1!1b0!94b1!8m0!15m4!1sWIbGZsbZKJrL1sQPjqvquAI!4m1!2i10147!7e81!20m28!1m6!1m2!1i0!2i0!2m2!1i530!2i945!1m6!1m2!1i1870!2i0!2m2!1i1920!2i945!1m6!1m2!1i0!2i0!2m2!1i1920!2i20!1m6!1m2!1i0!2i925!2m2!1i1920!2i945!27b1!28m0!40i703!47m1!8b1"

	// Criação do cliente HTTP
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Criação da requisição
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	// Adicionando os cabeçalhos
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("dnt", "1")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.google.com/")
	req.Header.Set("sec-ch-ua", `^\^Not`)
	req.Header.Set("cookie", `S=billing-ui-v3=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU:billing-ui-v3-efe=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU; receive-cookie-deprecation=1; HSID=AOO_sy-FmiCLRHi3i; SSID=ALvioYklsTgxGHHS2; APISID=9TaUF1BxDCyDv0gR/Aoj0Wl4WhUKNJFpFg; SAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-1PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-3PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; SID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BiadvdwGMYhGSJF6v6zUvFwACgYKAW0SARESFQHGX2MiyIgn3KDFeJA3YZGhD8ihhxoVAUF8yKpsZWEmUrj_AnE8dXeui5uY0076; __Secure-1PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BaAdle1TIFKsylBDp2NJwZAACgYKAR8SARESFQHGX2MioJKkn_9gGMcHm4eQ_Yud0hoVAUF8yKojKjMLjR_PC9FgMB33n2FY0076; __Secure-3PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BL6D2zb2ahaKjZntLSynErAACgYKAQESARESFQHGX2MiRN17UwaTTiZQsMByCOYgsBoVAUF8yKqe6uGmwERQcq9s7AjOsfwo0076; SEARCH_SAMESITE=CgQI6JsB; OGPC=19031986-1:; AEC=AVYB7cq9-ILZK5QjVxddGunI6ECPelIx7n_hIjYcQ39pXtN-MAeV66KLL44; NID=516=A83enFQcl1QUFtvWJ8diSkBi0XdhIfwylQGoDbUya2ZuBeKKW--PXv1ZsdEhI4EOKKZHEFaaW5hKN1VbPdABay9f2f18nNojdq_6DqprJZDdE-1Q3u8xhvIvyWvaPHmn9cfP0owY9wNHabC_C9CKfjbankvO2cOhm3fGpQKDtFwlZCoJIU5b3mQrhrdBImwx-zUCCAnAI0tSmL9k7r3JluZ93Gmw2quLfBUgJWDYSpx2pFSogqfkZAlbXq2eHB9UUTL3MyRAxO1nbfrGDsF4LysVkOA606bm0jKh8SKGC3UDqUyPNqbMve3k_DtX4dZ-UsQmrai0C9BT1Hy4TbrFWXI7bLnj7JFaLqdoSAzD_OjRWadzgrDDkDsHO-0icufSvOLWGgdnjT7wCVdp3HMqxZS9aXnFLrYPOyKjQf4UrIen0QZEp-645SC9jcKbujNfPTSot4CAMZeu0_jhiA8FfC7RZny8EtPrQy0tTmJ_YoIj5oLoH1bWTrLajddFBUglChSfJ1rwCQ-FB4A3u_Iop9nB6Etso2Afab8; DV=s9QN8ddaB_lSoH49I6th8VpDpUZ3F9ku3TJpNEbf1gAAAJBbb6P7S5CgeAAAACTWtiY5pKI3JgAAAILUrVF03KrXCwAAAA; __Secure-1PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; __Secure-3PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; SIDCC=AKEyXzVkqwOtb_ZeisFTG16QhOk__PpmW5x40Oo3tsGi8X2RU60-Q5gZglrd6PiaV_3hDcphtw; __Secure-1PSIDCC=AKEyXzWl-d7vAAFHftQtOo7hNqpghJpygezDpJzsW3h5gEgxCT6nReyMiZy1hUIfyZLh8vu--w; __Secure-3PSIDCC=AKEyXzWI-lehfd3RXGL74LvQ0YBbCh3QJPcj18G4AjFqE492NVIKpP9agrwYbpUTGTPracbnNw`)

	// Realiza a requisição
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	// Exibe o status da resposta
	fmt.Printf("Status: %s\n", resp.Status)

	// Lê e exibe o corpo da resposta
	body := make([]byte, 1024)
	n, err := resp.Body.Read(body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}

	fmt.Printf("Corpo da resposta:\n%s\n", string(body[:n]))
}

// curl --request GET \
//   --url 'https://www.google.com/maps/preview/directions?authuser=0&hl=pt-BR&gl=br&pb=!1m2!1sCasa%2Bda%2BPizza%2B-%2BAvenida%2BAtl%C3%A2ntica%2B-%2BInterlagos%2C%2BS%C3%A3o%2BPaulo%2B-%2BSP!2s0x94ce4e31947dd573%3A0x209bb115e9bc5236!1m5!1sPDKBikes%2B-%2BR.%2Bda%2BPaz%2C%2B2069%2B-%2BCh%C3%A1cara%2BSanto%2BAnt%C3%B4nio%2B(Zona%2BSul)%2C%2BS%C3%A3o%2BPaulo%2B-%2BSP%2C%2B04713-002!2s0x94ce512254a963bb%3A0x6230b9b9d57f824d!3m2!3d-23.6262784!4d-46.707612499999996!3m12!1m3!1d4497.841894012286!2d-46.710161123786264!3d-23.626237263980634!2m3!1f0!2f0!3f0!3m2!1i1920!2i945!4f13.1!6m24!1m2!18b1!30b1!2m3!5m1!6e2!20e3!10b1!12b1!13b1!14b1!16b1!17m1!3e1!20m6!1e0!2e3!5e2!6b1!8b1!14b1!46m1!1b0!94b1!8m0!15m4!1sWIbGZsbZKJrL1sQPjqvquAI!4m1!2i10147!7e81!20m28!1m6!1m2!1i0!2i0!2m2!1i530!2i945!1m6!1m2!1i1870!2i0!2m2!1i1920!2i945!1m6!1m2!1i0!2i0!2m2!1i1920!2i20!1m6!1m2!1i0!2i925!2m2!1i1920!2i945!27b1!28m0!40i703!47m1!8b1' \
//   --header '^sec-ch-ua: ^\^Not' \
//   --header 'accept: */*' \
//   --header 'accept-language: en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7' \
//   --header 'cookie: S=billing-ui-v3=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU:billing-ui-v3-efe=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU; receive-cookie-deprecation=1; HSID=AOO_sy-FmiCLRHi3i; SSID=ALvioYklsTgxGHHS2; APISID=9TaUF1BxDCyDv0gR/Aoj0Wl4WhUKNJFpFg; SAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-1PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-3PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; SID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BiadvdwGMYhGSJF6v6zUvFwACgYKAW0SARESFQHGX2MiyIgn3KDFeJA3YZGhD8ihhxoVAUF8yKpsZWEmUrj_AnE8dXeui5uY0076; __Secure-1PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BaAdle1TIFKsylBDp2NJwZAACgYKAR8SARESFQHGX2MioJKkn_9gGMcHm4eQ_Yud0hoVAUF8yKojKjMLjR_PC9FgMB33n2FY0076; __Secure-3PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BL6D2zb2ahaKjZntLSynErAACgYKAQESARESFQHGX2MiRN17UwaTTiZQsMByCOYgsBoVAUF8yKqe6uGmwERQcq9s7AjOsfwo0076; SEARCH_SAMESITE=CgQI6JsB; OGPC=19031986-1:; AEC=AVYB7cq9-ILZK5QjVxddGunI6ECPelIx7n_hIjYcQ39pXtN-MAeV66KLL44; NID=516=A83enFQcl1QUFtvWJ8diSkBi0XdhIfwylQGoDbUya2ZuBeKKW--PXv1ZsdEhI4EOKKZHEFaaW5hKN1VbPdABay9f2f18nNojdq_6DqprJZDdE-1Q3u8xhvIvyWvaPHmn9cfP0owY9wNHabC_C9CKfjbankvO2cOhm3fGpQKDtFwlZCoJIU5b3mQrhrdBImwx-zUCCAnAI0tSmL9k7r3JluZ93Gmw2quLfBUgJWDYSpx2pFSogqfkZAlbXq2eHB9UUTL3MyRAxO1nbfrGDsF4LysVkOA606bm0jKh8SKGC3UDqUyPNqbMve3k_DtX4dZ-UsQmrai0C9BT1Hy4TbrFWXI7bLnj7JFaLqdoSAzD_OjRWadzgrDDkDsHO-0icufSvOLWGgdnjT7wCVdp3HMqxZS9aXnFLrYPOyKjQf4UrIen0QZEp-645SC9jcKbujNfPTSot4CAMZeu0_jhiA8FfC7RZny8EtPrQy0tTmJ_YoIj5oLoH1bWTrLajddFBUglChSfJ1rwCQ-FB4A3u_Iop9nB6Etso2Afab8; DV=s9QN8ddaB_lSoH49I6th8VpDpUZ3F9ku3TJpNEbf1gAAAJBbb6P7S5CgeAAAACTWtiY5pKI3JgAAAILUrVF03KrXCwAAAA; __Secure-1PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; __Secure-3PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; SIDCC=AKEyXzVkqwOtb_ZeisFTG16QhOk__PpmW5x40Oo3tsGi8X2RU60-Q5gZglrd6PiaV_3hDcphtw; __Secure-1PSIDCC=AKEyXzWl-d7vAAFHftQtOo7hNqpghJpygezDpJzsW3h5gEgxCT6nReyMiZy1hUIfyZLh8vu--w; __Secure-3PSIDCC=AKEyXzWI-lehfd3RXGL74LvQ0YBbCh3QJPcj18G4AjFqE492NVIKpP9agrwYbpUTGTPracbnNw' \
//   --header 'dnt: 1' \
//   --header 'priority: u=1, i' \
//   --header 'referer: https://www.google.com/' \
//   --cookie 'S=billing-ui-v3=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU:billing-ui-v3-efe=eO83sU_kgRH6elkZY1_YMuPqpzPpq9iU; receive-cookie-deprecation=1; HSID=AOO_sy-FmiCLRHi3i; SSID=ALvioYklsTgxGHHS2; APISID=9TaUF1BxDCyDv0gR/Aoj0Wl4WhUKNJFpFg; SAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-1PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; __Secure-3PAPISID=5umMX-bwvWQAvuSs/A6ULpcdrEdhrwKppW; SID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BiadvdwGMYhGSJF6v6zUvFwACgYKAW0SARESFQHGX2MiyIgn3KDFeJA3YZGhD8ihhxoVAUF8yKpsZWEmUrj_AnE8dXeui5uY0076; __Secure-1PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BaAdle1TIFKsylBDp2NJwZAACgYKAR8SARESFQHGX2MioJKkn_9gGMcHm4eQ_Yud0hoVAUF8yKojKjMLjR_PC9FgMB33n2FY0076; __Secure-3PSID=g.a000mgil_1-XR4TQh8ZHh5x3FagjhQd2sUZ3COEcSWApAq46Pa9BL6D2zb2ahaKjZntLSynErAACgYKAQESARESFQHGX2MiRN17UwaTTiZQsMByCOYgsBoVAUF8yKqe6uGmwERQcq9s7AjOsfwo0076; SEARCH_SAMESITE=CgQI6JsB; OGPC=19031986-1:; AEC=AVYB7cq9-ILZK5QjVxddGunI6ECPelIx7n_hIjYcQ39pXtN-MAeV66KLL44; NID=516=A83enFQcl1QUFtvWJ8diSkBi0XdhIfwylQGoDbUya2ZuBeKKW--PXv1ZsdEhI4EOKKZHEFaaW5hKN1VbPdABay9f2f18nNojdq_6DqprJZDdE-1Q3u8xhvIvyWvaPHmn9cfP0owY9wNHabC_C9CKfjbankvO2cOhm3fGpQKDtFwlZCoJIU5b3mQrhrdBImwx-zUCCAnAI0tSmL9k7r3JluZ93Gmw2quLfBUgJWDYSpx2pFSogqfkZAlbXq2eHB9UUTL3MyRAxO1nbfrGDsF4LysVkOA606bm0jKh8SKGC3UDqUyPNqbMve3k_DtX4dZ-UsQmrai0C9BT1Hy4TbrFWXI7bLnj7JFaLqdoSAzD_OjRWadzgrDDkDsHO-0icufSvOLWGgdnjT7wCVdp3HMqxZS9aXnFLrYPOyKjQf4UrIen0QZEp-645SC9jcKbujNfPTSot4CAMZeu0_jhiA8FfC7RZny8EtPrQy0tTmJ_YoIj5oLoH1bWTrLajddFBUglChSfJ1rwCQ-FB4A3u_Iop9nB6Etso2Afab8; DV=s9QN8ddaB_lSoH49I6th8VpDpUZ3F9ku3TJpNEbf1gAAAJBbb6P7S5CgeAAAACTWtiY5pKI3JgAAAILUrVF03KrXCwAAAA; __Secure-1PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; __Secure-3PSIDTS=sidts-CjIBUFGohxVuMrOpM54cYoZxcfUw9-eIzWr1mu4hQrsvNLhD8ZwFkpp3HbVtb5ADuVPLrxAA; SIDCC=AKEyXzVkqwOtb_ZeisFTG16QhOk__PpmW5x40Oo3tsGi8X2RU60-Q5gZglrd6PiaV_3hDcphtw; __Secure-1PSIDCC=AKEyXzWl-d7vAAFHftQtOo7hNqpghJpygezDpJzsW3h5gEgxCT6nReyMiZy1hUIfyZLh8vu--w; __Secure-3PSIDCC=AKEyXzWI-lehfd3RXGL74LvQ0YBbCh3QJPcj18G4AjFqE492NVIKpP9agrwYbpUTGTPracbnNw'
