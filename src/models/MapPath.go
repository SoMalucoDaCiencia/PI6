package models

import (
	"PI6/src"
	"encoding/json"
	"fmt"
)

type MapPath struct {
	origin, destiny GeoPos
}

func NewMapPath(origin, destiny GeoPos) MapPath {
	return MapPath{
		origin:  origin,
		destiny: destiny,
	}
}

func (this *MapPath) ExtractRegister() (*Register, error) {

	ret := &Register{}
	respnse := []byte{}
	headers := map[string]string{
		"accept":          "*/*",
		"accept-language": "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7",
		"dnt":             "1",
		"priority":        "u=1, i",
		"referer":         "https://www.google.com/",
		"sec-ch-ua":       `^\^Not`,
		"cookie":          src.GoogleToken,
	}

	// t := time.Now().UnixMilli()
	_, err := src.Rest("GET", this.GetURI(), &respnse, headers, nil, nil)
	if err != nil {
		return nil, err
	}
	// println(time.Now().UnixMilli()-t, "millis")

	// println(string(respnse[4:]))

	f := []any{}
	err = json.Unmarshal(respnse[4:], &f)
	if err != nil {
		return nil, err
	}

	// res := mine.AbsoluteFlatMap(f)
	// println(len(res))
	print(fmt.Sprintf("%v", f[0].([]any)[20].([]any)[0]))
	print(fmt.Sprintf("%v", f[0].([]any)[20].([]any)[2]))

	return ret, nil
}

func (this *MapPath) GetURI() string {
	return fmt.Sprintf(
		src.GoogleURI,
		this.origin.lat,
		this.origin.long,
		this.origin.lat,
		this.origin.long,
		this.destiny.lat,
		this.destiny.long,
		this.destiny.lat,
		this.destiny.long,
	)
}
