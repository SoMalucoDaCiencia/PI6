package main

import (
	"PI6/models"
	"PI6/share"
	"encoding/json"
	"os"
)

func main() {
	file, err := os.ReadFile("../misc/proxies.json")
	if err != nil {
		panic(err)
	}
	var proxies []models.ProxyObj
	err = json.Unmarshal(file, &proxies)
	if err != nil {
		panic(err)
	}

	share.GetRandomProxy(proxies)
	//l, _ := json.Marshal(proxies)
	//err = os.WriteFile("../misc/proxies.json", l, 666)
	//if err != nil {
	//    panic(err)
	//}
}
