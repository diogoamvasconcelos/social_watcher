package main

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

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

		var dynamoDbItem controllers.DynamoDBStoredItem
		err := lib.UnmarshalStreamImage(record.Change.NewImage, &dynamoDbItem)
		if err != nil {
			log.Fatalf("Failed to unmarshal Record, %v", err)
		}

		storedItem, err := fromDynamoDBItemToStoredItem(dynamoDbItem)
		if err != nil {
			log.Fatalf("Failed to convert DynamoDBItem to StoredItem, %v", err)
		}

		res := controllers.PostToDiscord(storedItem)
		if res == "ERROR" {
			log.Fatalf("Failed to post item to discord: %#v", storedItem)
		}
	}

	return "OK", nil
}

func main() {
	log.SetPrefix("DdbStreamConsumer: ")
	lambda.Start(handleRequest)
}

func fromDynamoDBItemToStoredItem(item controllers.DynamoDBStoredItem) (controllers.StoredItem, error) {
	itemIndex, err := strconv.Atoi(item.SK)
	if err != nil {
		return controllers.StoredItem{}, err
	}

	happenedAt, err := lib.FromISO8061(item.GSI1SK)
	if err != nil {
		return controllers.StoredItem{}, err
	}

	gsi2pkUnpacked := strings.Split(item.GSI2PK, "|")
	if len(gsi2pkUnpacked) != 2 {
		return controllers.StoredItem{}, errors.New("Failed to unpack GSI2PK")
	}

	return controllers.StoredItem{
		ID:         item.PK,
		ItemIndex:  itemIndex,
		HappenedAt: happenedAt,
		SourceType: gsi2pkUnpacked[0],
		Keyword:    gsi2pkUnpacked[1],
		Link:       item.Link,
		Data:       item.Data,
	}, nil
}
