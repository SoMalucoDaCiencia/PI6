package src

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html/charset"

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

func Rest(method, uri string, returnJson *[]byte, headers map[string]string, query map[string]string, bodyInput []byte) (int, error) {
	return RestClient(http.DefaultClient, method, uri, returnJson, headers, query, bodyInput)
}

func RestClient(httpClient *http.Client, method string, uri string, returnJson *[]byte, headers map[string]string, query map[string]string, bodyInput []byte) (int, error) {

	if bodyInput == nil {
		bodyInput = []byte{}
	}
	var buf = *bytes.NewBuffer(bodyInput)

	var req *http.Request
	r, err := http.NewRequest(method, uri, &buf)
	req = r

	q := req.URL.Query()
	for key, value := range query {
		q.Set(key, value)
	}

	req.URL.RawQuery = q.Encode()

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	//req = req.WithContext(ctx)
	//
	resp, err := (*httpClient).Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "connect: connection refused") {
			return http.StatusBadGateway, errors.New("Failed to connect to the upstream server pointed as \"" + uri + "\"")
		}
		return http.StatusInternalServerError, err
	}
	if resp.StatusCode/100 != 2 {
		return resp.StatusCode, errors.New(resp.Status)
	}

	// Ensure the body will be closed
	defer func(Body io.ReadCloser) {
		errBody := Body.Close()
		if errBody != nil {
			return
		}
	}(resp.Body)

	var body []byte
	if (strings.Contains(resp.Header.Get("content-type"), "application/json") ||
		strings.Contains(resp.Header.Get("content-type"), "text/plain")) && resp.ContentLength > 0 {
		// Read the body content
		unicodeReader, errDe := charset.NewReader(resp.Body, "latin1")
		body, errDe = io.ReadAll(unicodeReader)
		if errDe != nil {
			return http.StatusInternalServerError, errDe
		}
	} else {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	*returnJson = body
	return http.StatusOK, nil
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
