package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type TwitterGetUserDetailResult struct {
	Data struct {
		PublicMetrics struct {
			Followers int `json:"followers_count"`
		} `json:"public_metrics"`
	} `json:"data"`
}

type TwitterUserDetails struct {
	ID        string `json:"id"`
	Followers int    `json:"followers"`
}

func GetTwitterUserDetails(userId string) (TwitterUserDetails, error) {
	log.Printf("GetTwitterUserDetails id: %v", userId)

	credentials := GetTwitterCredentials()
	client := lib.NewTwitterClient(credentials.BearerToken)

	// https://developer.twitter.com/en/docs/labs/tweets-and-users/api-reference/get-users-id
	queryParams := "user.fields=public_metrics"
	resp, err := client.Get(fmt.Sprintf("https://api.twitter.com/labs/2/users/%s?%s", userId, queryParams))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetUserDetails resp.status: %v", resp.Status)

	bodyBinary, err := lib.GetBodyBinaryFromHttpResp(resp)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("GetUserDetails failed: %v", string(bodyBinary))
		return TwitterUserDetails{}, errors.New(("Failed to GetUserDetails"))
	}

	var body TwitterGetUserDetailResult
	err = json.Unmarshal(bodyBinary, &body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetUserDetails resp.body: %#v", body)

	return TwitterUserDetails{ID: userId, Followers: body.Data.PublicMetrics.Followers}, nil
}
