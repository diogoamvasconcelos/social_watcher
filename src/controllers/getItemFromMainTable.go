package controllers

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type MainItemKey struct {
	// HOW-TO do Pick<MainItem, "PK | SK"> ??
	PK string
	SK string
}

func GetMainItem(key MainItemKey) (MainItem, error) {
	dynamodbClient := lib.NewDynamoClient()
	mainTableName := os.Getenv("MAIN_TABLE_NAME")

	result, err := lib.DynamoDBGetItem(dynamodbClient, mainTableName, key)
	if err != nil {
		log.Print("Failed to get item: ", err)
		return MainItem{}, err
	}

	if len(result) == 0 {
		// Not found
		return MainItem{}, nil
	}

	dbitem := DynamoDBMainItem{}
	err = dynamodbattribute.UnmarshalMap(result, &dbitem)
	if err != nil {
		log.Fatal("Failed to unmarshall dbitem: ", err)
	}

	item, err := FromDynamoDBMainItemToMainItem(dbitem)
	if err != nil {
		log.Fatal("Failed to convert dbitem to MainItem: ", err)
	}

	return item, nil
}

func FromDynamoDBMainItemToMainItem(dbitem DynamoDBMainItem) (MainItem, error) {
	updatedAt, err := lib.FromISO8061(dbitem.GSI1SK)
	if err != nil {
		return MainItem{}, err
	}

	return MainItem{
		ID:        dbitem.PK,
		Type:      dbitem.GSI1PK,
		UpdatedAt: updatedAt,
		Data:      dbitem.Data,
	}, nil
}
