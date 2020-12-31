package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

		var dynamoDbItem controllers.DynamoDBItem
		err := UnmarshalStreamImage(record.Change.NewImage, &dynamoDbItem)
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
	log.SetPrefix("DDbStreeamConsumer: ")
	lambda.Start(handleRequest)
}

// UnmarshalStreamImage converts events.DynamoDBAttributeValue to struct
// from: https://stackoverflow.com/questions/49129534/unmarshal-mapstringdynamodbattributevalue-into-a-struct
func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)
	for k, v := range attribute {
		var dbAttr dynamodb.AttributeValue
		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
}

func fromDynamoDBItemToStoredItem(item controllers.DynamoDBItem) (controllers.StoredItem, error) {
	itemIndex, err := strconv.Atoi(item.SK)
	if err != nil {
		return controllers.StoredItem{}, err
	}

	happenedAt, err := lib.FromISO8061(item.GSI1SK)
	if err != nil {
		return controllers.StoredItem{}, err
	}

	return controllers.StoredItem{
		ID:         item.PK,
		ItemIndex:  itemIndex,
		HappenedAt: happenedAt,
		SourceType: item.GSI2PK,
		Link:       item.Link,
		Data:       item.Data,
	}, nil
}
