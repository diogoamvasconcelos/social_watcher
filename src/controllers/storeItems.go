package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type StoredItem struct {
	ID         string // Source|Hash
	ItemIndex  int    // 0 = initial, >0 = aggregate
	HappenedAt time.Time
	SourceType string // twitter
	Link       string
	Data       TwitterSearchResultTweet
}

type DynamoDBItem struct {
	PK     string // ID
	SK     string // ItemIndex
	GSI1PK string // "ALL"
	GSI1SK string // HappenedAt
	GSI2PK string // SourceType
	GSI2SK string // HappenedAt
	Link   string
	Data   TwitterSearchResultTweet
}

func StoreItems(items TwitterSearchResult) string {
	dynamodbClient := lib.NewDynamoClient()
	storedItemsTableName := os.Getenv("STORED_ITEMS_TABLE_NAME")

	for _, item := range items.Data {
		storedItem := fromTwitterSearchItemToStoredItem(item)
		log.Printf("Item to store: %#v", storedItem)

		dynamodbItem := fromStoredItemToDynamoDBItem(storedItem)

		result := lib.DynamoDBPutItem(dynamodbClient, storedItemsTableName, dynamodbItem, "attribute_not_exists(PK)")
		log.Printf("Put DynamoDB item result: %v", result)
		if result == "ERROR" {
			log.Fatal("Failed to store item in:", storedItemsTableName)
		}
	}

	return "OK"
}

func fromTwitterSearchItemToStoredItem(item TwitterSearchResultTweet) StoredItem {
	sourceType := "twitter"

	return StoredItem{
		ID:         fmt.Sprintf("%s|%s", sourceType, item.ID),
		ItemIndex:  0,
		HappenedAt: item.CreatedAt,
		SourceType: sourceType,
		Link:       fmt.Sprintf("https://twitter.com/x/status/%s", item.ID),
		Data:       item,
	}
}

func fromStoredItemToDynamoDBItem(item StoredItem) DynamoDBItem {
	happenedAtString := lib.ToISO8061(item.HappenedAt)

	return DynamoDBItem{
		PK:     item.ID,
		SK:     strconv.Itoa(item.ItemIndex),
		GSI1PK: "All",
		GSI1SK: happenedAtString,
		GSI2PK: item.SourceType,
		GSI2SK: happenedAtString,
		Link:   item.Link,
		Data:   item.Data,
	}
}
