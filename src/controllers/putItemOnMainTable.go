package controllers

import (
	"log"
	"os"
	"time"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type MainItem struct {
	ID        string // Type|ID
	Type      string // health_check
	UpdatedAt time.Time
	Data      string
}

type DynamoDBMainItem struct {
	PK        string // ID
	SK        string
	GSI1PK    string // Type
	GSI1SK    string
	UpdatedAt string
	Data      string
}

func PutMainItem(item MainItem) string {
	dynamodbClient := lib.NewDynamoClient()
	mainTableName := os.Getenv("MAIN_TABLE_NAME")

	log.Printf("Item to put: %#v", item)
	dynamodbItem := fromMainItemToDynamoDBMainItem(item)

	result := lib.DynamoDBPutItem(dynamodbClient, mainTableName, dynamodbItem, "")
	log.Printf("Put DynamoDB item result: %v", result)
	if result == "ERROR" {
		log.Fatal("Failed to store item in:", mainTableName)
	}

	return "OK"
}

func fromMainItemToDynamoDBMainItem(item MainItem) DynamoDBMainItem {
	updatedAtString := lib.ToISO8061(item.UpdatedAt)

	return DynamoDBMainItem{
		PK:     item.ID,
		SK:     "N/A",
		GSI1PK: item.Type,
		GSI1SK: updatedAtString,
		Data:   item.Data,
	}
}
