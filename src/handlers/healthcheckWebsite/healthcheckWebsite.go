package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoamvasconcelos/social_watcher/src/controllers"
	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type eventData struct {
	Website string `json:"website"`
}

func handleRequest(ctx context.Context, event eventData) (string, error) {
	log.Printf("Event data: %v", event)
	healthcheckResult := controllers.HealthcheckWebsite(event.Website)

	log.Printf("Result %#v", healthcheckResult)

	resultStringified, err := lib.JSONStringify(healthcheckResult)
	if err != nil {
		log.Fatal(err)
	}

	lastHealthCheck, err := controllers.GetMainItem(controllers.MainItemKey{PK: toMainItemID(event.Website), SK: "N/A"})
	if err != nil {
		log.Fatal(nil)
	}

	lastHealthCheckData := controllers.HealthcheckWebsiteResult{}
	err = json.Unmarshal([]byte(lastHealthCheck.Data), &lastHealthCheckData)
	if err != nil {
		log.Fatal(err)
	}

	if healthcheckResult.IsUp != lastHealthCheckData.IsUp {
		log.Printf("Updating healthcheck for website: %s -> new value:%v", event.Website, healthcheckResult.IsUp)

		controllers.PutMainItem(controllers.MainItem{
			ID:        toMainItemID(event.Website),
			Type:      "health_check",
			UpdatedAt: lib.MinutesAgo(0),
			Data:      resultStringified,
		})

		// TODO: post to Discord
	} else {
		log.Printf("No need to update healthcheck for website: %s -> still:%v", event.Website, healthcheckResult.IsUp)
	}

	return "OK", nil
}

func main() {
	log.SetPrefix("HealthcheckWebsite: ")
	lambda.Start(handleRequest)
}

func toMainItemID(website string) string {
	return fmt.Sprintf("health_check|%s", website)
}
