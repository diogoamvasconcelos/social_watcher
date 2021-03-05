package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type TwitterData struct {
	TwitterSearchResultTweet        // important: don't use a comma here
	TranslatedText           string `json:"translated_text"`
}

// TODO: Create a common base struct for StoredItem and NewKeywordItem

type StoredItem struct {
	ID         string // Source|Hash
	ItemIndex  int    // 0 = initial, >0 = aggregate
	HappenedAt time.Time
	SourceType string // twitter
	Keyword    string
	Link       string
	Data       TwitterData
}

type DynamoDBStoredItem struct {
	PK     string // ID
	SK     string // ItemIndex
	GSI1PK string // "ALL"
	GSI1SK string // HappenedAt
	GSI2PK string // SourceType|Keyword
	GSI2SK string // HappenedAt
	Link   string
	Data   TwitterData
}

type NewKeywordItem struct {
	ID         string // Source|Hash
	ItemIndex  int    // 0 = initial, >0 = aggregate
	HappenedAt time.Time
	SourceType string // twitter
	Keyword    string
	Type       string // NewKeyword
}

type DynamoDBNewKeywordItem struct {
	PK     string // ID
	SK     string // ItemIndex
	GSI1PK string // "ALL"
	GSI1SK string // HappenedAt
	GSI2PK string // SourceType|Keyword
	GSI2SK string // HappenedAt
}

func StoreItems(items TwitterSearchResult, keyword string) string {
	dynamodbClient := lib.NewDynamoClient()
	storedItemsTableName := os.Getenv("STORED_ITEMS_TABLE_NAME")

	nofItems := len(items.Data)
	errors := make(chan error, nofItems)

	for _, item := range items.Data {
		go func(item TwitterSearchResultTweet) {
			storedItem := fromTwitterSearchItemToStoredItem(item, keyword)
			log.Printf("Item to store: %#v", storedItem)

			prevStoredItem, err := GetStoredItem(StoredItemKey{PK: storedItem.ID, SK: strconv.Itoa(storedItem.ItemIndex)})
			if err != nil {
				errors <- err
				return
			}

			if (prevStoredItem != StoredItem{}) {
				log.Printf("Skipping as item is already stored")
				errors <- nil
				return
				/*
					// check if new keyword or should skip
					if prevStoredItem.Keyword == keyword {
						log.Printf("Skipping as item is already stored")
					} else {
						dynamodbNewKeywordItem := toDynamoDBNewKeywordItem(fromStoredItemToNewKeywordItem(storedItem))

						// TODO: Deduplicate this
						result := lib.DynamoDBPutItem(dynamodbClient, storedItemsTableName, dynamodbNewKeywordItem, "attribute_not_exists(PK)")
						log.Printf("Put DynamoDB item result: %v", result)
						if result == "ERROR" {
							log.Fatal("Failed to store item in:", storedItemsTableName)
						}
					}
				*/
			}

			dynamodbItem := fromStoredItemToDynamoDBItem(storedItem)

			result := lib.DynamoDBPutItem(dynamodbClient, storedItemsTableName, dynamodbItem, "attribute_not_exists(PK)")
			log.Printf("Put DynamoDB item result: %v", result)
			if result == "ERROR" {
				errors <- fmt.Errorf("Failed to store item in: %v", storedItemsTableName)
			} else {
				errors <- nil
			}

		}(item)
	}

	for i := 0; i < nofItems; i++ {
		if err := <-errors; err != nil {
			log.Fatal(err)
		}
	}

	return "OK"
}

func fromTwitterSearchItemToStoredItem(item TwitterSearchResultTweet, keyword string) StoredItem {
	sourceType := "twitter"

	/*
			// HOW-TO: spread `data` like this
			twitterData := TwitterData{
				...item,
				TranslatedText: TranslateToEnglish(item.Text, item.Lang),
			}

		or even this
		twitterData := TwitterData{
			ID:             item.ID,
			Text:           item.Text,
			CreatedAt:      item.CreatedAt,
			ConversationID: item.ConversationID,
			Lang:           item.Lang,
			TranslatedText: TranslateToEnglish(item.Text, item.Lang),
		}
	*/

	// NICE!!
	twitterData := TwitterData{TwitterSearchResultTweet: item, TranslatedText: TranslateToEnglish(item.Text, item.Lang)}

	return StoredItem{
		ID:         toStoredItemID(item.ID, sourceType),
		ItemIndex:  0,
		HappenedAt: item.CreatedAt,
		SourceType: sourceType,
		Keyword:    keyword,
		Link:       fmt.Sprintf("https://twitter.com/x/status/%s", item.ID),
		Data:       twitterData,
	}
}

func fromStoredItemToDynamoDBItem(item StoredItem) DynamoDBStoredItem {
	happenedAtString := lib.ToISO8061(item.HappenedAt)

	return DynamoDBStoredItem{
		PK:     item.ID,
		SK:     strconv.Itoa(item.ItemIndex),
		GSI1PK: "All",
		GSI1SK: happenedAtString,
		GSI2PK: fmt.Sprintf("%s|%s", item.SourceType, item.Keyword),
		GSI2SK: happenedAtString,
		Link:   item.Link,
		Data:   item.Data,
	}
}

func toStoredItemID(uniqueID string, sourceType string) string {
	return fmt.Sprintf("%s|%s", sourceType, uniqueID)
}

func toDynamoDBNewKeywordItem(item NewKeywordItem) DynamoDBNewKeywordItem {
	happenedAtString := lib.ToISO8061(item.HappenedAt)

	return DynamoDBNewKeywordItem{
		PK:     item.ID,
		SK:     strconv.Itoa(item.ItemIndex),
		GSI1PK: "All",
		GSI1SK: happenedAtString,
		GSI2PK: fmt.Sprintf("%s|%s", item.SourceType, item.Keyword),
		GSI2SK: happenedAtString,
	}
}

func fromStoredItemToNewKeywordItem(item StoredItem) NewKeywordItem {
	// HOW-TO: omit just some fields (Link, Data)
	// Maybe use the struct coverter, like for `json/xml`
	return NewKeywordItem{
		ID:         item.ID,
		ItemIndex:  1,
		HappenedAt: time.Now(),
		SourceType: item.SourceType,
		Keyword:    item.Keyword,
		Type:       "NewKeyword",
	}
}
