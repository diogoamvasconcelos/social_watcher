package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type SearchResultTweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type SearchResultMetadata struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type SearchResult struct {
	Data []SearchResultTweet  `json:"data"`
	Meta SearchResultMetadata `json:"meta"`
}

type twitterBotCredentials struct {
	APIKey       string `json:"ApiKey"`
	APISecretKey string `json:"ApiSecretKey"`
	BearerToken  string `json:"BearerToken"`
}

func SearchTwitter(keyword string) SearchResult {
	log.Printf("SearchTwitter keyword: %v", keyword)

	ssmClient := lib.NewSSMClient()
	ssmResult, ssmErr := ssmClient.Param("twitter_bot_keys", true).GetValue()
	if ssmErr != nil {
		log.Fatal(ssmErr)
	}
	var credentials twitterBotCredentials
	err := json.Unmarshal([]byte(ssmResult), &credentials)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: credentials.BearerToken,
		TokenType:   "Bearer",
	}))

	// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
	resp, err := client.Get(fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?query=%s", keyword))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Printf("Search Response status: %v", resp.Status)

	bodyBinary, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var body SearchResult
	err = json.Unmarshal(bodyBinary, &body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Search Result Body", body)

	return body

}
