package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoamvasconcelos/social_watcher/src/controllers"
	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

func handleRequest(ctx context.Context, event events.DynamoDBEvent) (string, error) {
	log.Printf("Initial Event: %#v", event)

	for _, record := range event.Records {
		if record.EventName == "REMOVE" {
			log.Printf("Skipping 'REMOVE' event: %#v", record.Change)
			continue
		}

		var dynamoDbItem controllers.DynamoDBMainItem
		err := lib.UnmarshalStreamImage(record.Change.NewImage, &dynamoDbItem)
		if err != nil {
			log.Fatalf("Failed to unmarshal Record, %v", err)
		}

		mainItem, err := controllers.FromDynamoDBMainItemToMainItem(dynamoDbItem)
		if err != nil {
			log.Fatalf("Failed to convert DynamoDBItem to MainItem, %v", err)
		}

		log.Print("Processing the following MainItem, ", mainItem)

		// Filter item
		switch mainItem.Type {
		case "healthcheck":
			{
				healthCheckData := controllers.HealthcheckWebsiteResult{}
				err = json.Unmarshal([]byte(mainItem.Data), &healthCheckData)
				if err != nil {
					log.Fatal("Failed to convert mainItem.Data to HealthcheckWebsiteResult, ", err)
				}

				res := controllers.PostHealthcheckToDiscord(healthCheckData)
				if res == "ERROR" {
					log.Fatalf("Failed to post item to discord: %#v", mainItem)
				}
			}
		}
	}

	return "OK", nil
}

func main() {
	log.SetPrefix("MainDdbStreamConsumer: ")
	lambda.Start(handleRequest)
}
