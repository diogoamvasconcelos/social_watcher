package controllers

import (
	"fmt"
	"log"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

func PostHealthcheckToDiscord(healthcheck HealthcheckWebsiteResult) string {
	// Get healthcheck config
	mappingsConfig := GetMappingsConfig()
	healthcheckConfig := mappingsConfig.HealthCheck

	// Discord Client
	credentials := GetDiscordCredentials()
	discordClient, err := lib.NewDiscordClient(credentials.Token)
	if err != nil {
		return "ERROR"
	}

	err = discordClient.Open()
	if err != nil {
		fmt.Println("Discord: Error opening connection,", err)
		return "ERROR"
	}

	var message string
	if healthcheck.IsUp {
		message += fmt.Sprintf(":green_heart: <%s> is UP again :green_heart:", healthcheck.Website)
	} else {
		message += fmt.Sprintf(":rotating_light::fire: <%s> is DOWN @here :fire::rotating_light:", healthcheck.Website)
	}

	for _, channel := range healthcheckConfig.Channels {
		messageResult, err := discordClient.ChannelMessageSend(channel, message)
		if err != nil {
			log.Println("Discord: Failed to send message to Channel,", err)
			return "ERROR"
		}
		log.Print("Discord: Successfully sent message, ", messageResult)
	}

	discordClient.Close()
	return "SUCCESS"
}
