package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type eventData struct {
	Keyword string `json:"keyword to watch for"`
}

func handleRequest(ctx context.Context, event eventData) {
	log.Printf("Event %v", event)
	log.Printf("Event keyword: %s", event.Keyword)
}

func main() {
	log.SetPrefix("WatchTrigger: ")

	lambda.Start(handleRequest)
}
