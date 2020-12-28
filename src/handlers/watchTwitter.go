package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoamvasconcelos/social_watcher/src/controllers"
)

type eventData struct {
	Keyword string `json:"keyword"`
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	// https://blog.golang.org/json
	var data eventData
	err := json.Unmarshal(event.Detail, &data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Event data: %v", data)
	controllers.SearchTwitter(data.Keyword)

	return "Success", nil
}

func main() {
	log.SetPrefix("WatchTrigger: ")

	lambda.Start(handleRequest)
}
