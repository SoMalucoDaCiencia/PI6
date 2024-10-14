package share

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

func Rest(method, uri string, returnJson *[]byte, headers map[string]string, query map[string]string, bodyInput []byte) (int, error) {
	client := &http.Client{Timeout: time.Second * 30}
	return RestClient(client, method, uri, returnJson, headers, query, bodyInput)
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
