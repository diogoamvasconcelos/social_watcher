package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type HealthcheckWebsiteResult struct {
	Website string
	IsUp    bool
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
		log.Fatal("Failed to check 'downfor': ", err)
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

	isUp := !body.IsDown
	// Allow `Unauthorized` exception
	if !isUp && body.StatusCode == 401 && body.StatusText == "Unauthorized" {
		isUp = true
	}

	return HealthcheckWebsiteResult{
		Website: website,
		IsUp:    isUp,
	}
}
