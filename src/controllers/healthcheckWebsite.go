package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type HealthcheckWebsiteResult struct {
	IsUp bool
}

type DownforHttpCheckResponse struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	IsDown     bool   `json:"isDown"`
}

// Based on: https://downforeveryoneorjustme.com/
func HealthcheckWebsite(website string) HealthcheckWebsiteResult {
	log.Printf("Healthcheck website: %v", website)

	resp, err := http.Get(fmt.Sprintf("https://api-prod.downfor.cloud/httpcheck/%s", website))
	if err != nil {
		log.Fatal("Failed to check 'downfor'", err)
	}

	bodyBinary, err := lib.GetBodyBinaryFromHttpResp(resp)
	if err != nil {
		log.Fatal(err)
	}

	var body DownforHttpCheckResponse
	err = json.Unmarshal(bodyBinary, &body)
	if err != nil {
		log.Fatal(err)
	}

	return HealthcheckWebsiteResult{
		IsUp: !body.IsDown,
	}
}
