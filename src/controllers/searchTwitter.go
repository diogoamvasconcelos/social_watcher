package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type TwitterSearchResultTweet struct {
	ID             string    `json:"id"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
	ConversationID string    `json:"conversation_id"`
	AuthorID       string    `json:"author_id"`
	Lang           string    `json:"lang"`
}

type TwitterSearchResultMetadata struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type TwitterSearchResult struct {
	Data []TwitterSearchResultTweet  `json:"data"`
	Meta TwitterSearchResultMetadata `json:"meta"`
}

func SearchTwitter(keyword string) TwitterSearchResult {
	log.Printf("SearchTwitter keyword: %v", keyword)

	credentials := GetTwitterCredentials()
	client := lib.NewTwitterClient(credentials.BearerToken)

	// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
	// https://developer.twitter.com/en/docs/twitter-api/tweets/search/integrate/build-a-rule (for filter retweets)
	queryParams := fmt.Sprintf("query=%s", keyword)
	queryParams += "%20-is:retweet"
	queryParams += "&max_results=100" //max
	queryParams += fmt.Sprintf("&start_time=%s", lib.ToISO8061(lib.MinutesAgo(60*4)))
	queryParams += "&place.fields=country"
	queryParams += "&tweet.fields=created_at,lang,conversation_id,author_id"
	resp, err := client.Get(fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?%s", queryParams))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Search Response status: %v", resp.Status)

	bodyBinary, err := lib.GetBodyBinaryFromHttpResp(resp)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("Search Result (RAW): %#v", string(bodyBinary))

	var body TwitterSearchResult
	err = json.Unmarshal(bodyBinary, &body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Search Result: %#v", body)
	log.Printf("Search Result count: %#v", body.Meta.ResultCount)

	return body
}

type TwitterBotCredentials struct {
	APIKey       string `json:"ApiKey"`
	APISecretKey string `json:"ApiSecretKey"`
	BearerToken  string `json:"BearerToken"`
}

func GetTwitterCredentials() TwitterBotCredentials {
	ssmClient := lib.NewSSMClient()
	ssmResult, ssmErr := ssmClient.Param("twitter_bot_keys", true).GetValue()
	if ssmErr != nil {
		log.Fatal(ssmErr)
	}
	var credentials TwitterBotCredentials
	err := json.Unmarshal([]byte(ssmResult), &credentials)
	if err != nil {
		log.Fatal(err)
	}

	return credentials
}
