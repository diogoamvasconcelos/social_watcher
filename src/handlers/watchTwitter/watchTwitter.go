package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoamvasconcelos/social_watcher/src/controllers"
)

type eventData struct {
	Keyword string `json:"keyword"`
}

func handleRequest(ctx context.Context, event eventData) (string, error) {
	log.Printf("Event data: %v", event)
	items := controllers.SearchTwitter(event.Keyword)

	result := controllers.StoreItems(items, event.Keyword)
	return result, nil
}

func main() {
	log.SetPrefix("WatchTwitter: ")
	lambda.Start(handleRequest)
}
