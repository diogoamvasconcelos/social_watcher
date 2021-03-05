package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("https://api-prod.downfor.cloud/httpcheck/%s", website))
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

	// Allow `Unauthorized` exception
	isUp := !body.IsDown || body.StatusCode == http.StatusUnauthorized

	if !isUp {
		// Maybe should just use this instead: https://github.com/hashicorp/go-retryablehttp
		currentTry := 0
		for currentTry < 5 {
			log.Printf("downfor says it's down, going to try directly to check the website. Try number:%d", currentTry)
			resp, err := client.Get(website)
			if err != nil {
				log.Printf("Failed to check '%s': %v", website, err)
			} else {
				isUp = resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusUnauthorized
				if isUp {
					break
				}
			}

			currentTry++
		}
	}

	return HealthcheckWebsiteResult{
		Website: website,
		IsUp:    isUp,
	}
}
