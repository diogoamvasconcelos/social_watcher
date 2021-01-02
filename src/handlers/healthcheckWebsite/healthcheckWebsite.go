package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoamvasconcelos/social_watcher/src/controllers"
)

type eventData struct {
	Website string `json:"website"`
}

func handleRequest(ctx context.Context, event eventData) (string, error) {
	log.Printf("Event data: %v", event)
	healthcheckResult := controllers.HealthcheckWebsite(event.Website)

	log.Printf("Result %#v", healthcheckResult)
	//TODO check stored, if different update and post to slack :)
	return "OK", nil
}

func main() {
	log.SetPrefix("HealthcheckWebsite: ")
	lambda.Start(handleRequest)
}
