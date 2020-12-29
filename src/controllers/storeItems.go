package controllers

import (
	"fmt"
	"log"
	"time"
)

type StoredItem struct {
	ID         string                   `json:"id"`        // Source|Hash
	ItemIndex  int                      `json:"itemIndex"` // 0 = initial, >0 = aggregate
	HappenedAt time.Time                `json:"happenedAt"`
	SourceType string                   `json:"sourceType"` // twitter
	Link       string                   `json:"link"`
	Data       TwitterSearchResultTweet `json:"data"`
}

func StoreItems(items TwitterSearchResult) string {
	for _, item := range items.Data {
		storedItem := fromTwitterSearchItemToStoredItem(item)

		log.Printf("Item to store: %#v", storedItem)
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
		Data: TwitterSearchResultTweet{
			ID:             item.ID,
			Text:           item.Text,
			ConversationID: item.ConversationID,
		},
	}
}
