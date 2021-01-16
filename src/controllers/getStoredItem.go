package controllers

import (
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type StoredItemKey struct {
	// HOW-TO: do Pick<StoredItem, "PK | SK"> ??
	PK string
	SK string
}

func GetStoredItem(key StoredItemKey) (StoredItem, error) {
	dynamodbClient := lib.NewDynamoClient()
	mainTableName := os.Getenv("STORED_ITEMS_TABLE_NAME")

	result, err := lib.DynamoDBGetItem(dynamodbClient, mainTableName, key)
	if err != nil {
		log.Print("Failed to get item: ", err)
		return StoredItem{}, err
	}

	if len(result) == 0 {
		// Not found
		return StoredItem{}, nil
	}

	dbitem := DynamoDBStoredItem{}
	err = dynamodbattribute.UnmarshalMap(result, &dbitem)
	if err != nil {
		log.Fatal("Failed to unmarshall dbitem: ", err)
	}

	item, err := FromDynamoDStoredItemToMainItem(dbitem)
	if err != nil {
		log.Fatal("Failed to convert dbitem to MainItem: ", err)
	}

	return item, nil
}

func FromDynamoDStoredItemToMainItem(dbitem DynamoDBStoredItem) (StoredItem, error) {
	happenedAt, err := lib.FromISO8061(dbitem.GSI1SK)
	if err != nil {
		return StoredItem{}, err
	}

	itemIndex, err := strconv.Atoi(dbitem.SK)
	if err != nil {
		return StoredItem{}, err
	}

	return StoredItem{
		ID:         dbitem.PK,
		ItemIndex:  itemIndex,
		HappenedAt: happenedAt,
		SourceType: dbitem.GSI2PK,
		Data:       dbitem.Data,
		Link:       dbitem.Link,
	}, nil
}
