package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

type discordBotCredentials struct {
	Username string `json:"Username"`
	Token    string `json:"Token"`
}

func PostToDiscord(item StoredItem) string {
	ssmClient := lib.NewSSMClient()
	ssmResult, ssmErr := ssmClient.Param("discord_bot_keys", true).GetValue()
	if ssmErr != nil {
		log.Fatal(ssmErr)
	}
	var credentials discordBotCredentials
	err := json.Unmarshal([]byte(ssmResult), &credentials)
	if err != nil {
		log.Fatal(err)
	}

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
	message += "> New Twitter message"
	message += "\n"
	message += item.Link
	messageResult, err := discordClient.ChannelMessageSend("794147209976741918", message)
	if err != nil {
		log.Println("Discord: Failed to send message to Channel,", err)
		return "ERROR"
	}

	discordClient.Close()
	log.Print("Discord: Successfully sent message, ", messageResult)
	return "SUCCESS"
}
