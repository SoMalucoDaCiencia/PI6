package models

import (
	"PI6/share"
	"encoding/json"
	"time"
)

type AppleTokenGetter struct {
	Seed             string    `json:"seed"`
	CreatedAt        time.Time `json:"createdAt"`
	AccessToken      string    `json:"accessToken"`
	ExpiresInSeconds uint      `json:"expiresInSeconds"`
}

func (this *AppleTokenGetter) IsValid() bool {
	return this.CreatedAt.Add(time.Second * time.Duration(this.ExpiresInSeconds)).After(time.Now())
}

func (this *AppleTokenGetter) Renew() error {

	response := []byte{}
	headers := map[string]string{
		"Host":          "go-app",
		"Authorization": "Bearer " + this.Seed,
	}

	_, err := share.Rest("GET", "https://maps-api.apple.com/v1/token", &response, headers, nil, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(response, this)
}
