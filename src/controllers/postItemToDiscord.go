package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
	"gopkg.in/yaml.v2"
)

type KeywordConfig struct {
	Value    string   `yaml:"value"`
	Channels []string `yaml:"channels"`
}

type GenericConfig struct {
	Channels []string `yaml:"channels"`
}
type MappingsConfig struct {
	Keyword     []KeywordConfig `yaml:"keywords"`
	HealthCheck GenericConfig   `yaml:"healthcheck"`
}

type DiscordBotCredentials struct {
	Username string `json:"Username"`
	Token    string `json:"Token"`
}

var keywordMappingConfigPath = "configuration/mappings.yaml"

func PostToDiscord(item StoredItem) string {
	// Get keywordConfig
	mappingsConfig := GetMappingsConfig()

	keyword := "pureref"
	keywordConfig := KeywordConfig{}
	for _, mapping := range mappingsConfig.Keyword {
		if mapping.Value == keyword {
			keywordConfig = mapping
			break
		}
	}

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
	message += fmt.Sprintf("> New `%s` Twitter message (lang=%s)", keyword, item.Data.Lang)
	message += "\n"
	message += item.Link

	for _, channel := range keywordConfig.Channels {
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

func GetMappingsConfig() MappingsConfig {
	data, err := ioutil.ReadFile(keywordMappingConfigPath)
	if err != nil {
		log.Fatalf("Failed to open '%s': %v", keywordMappingConfigPath, err)
	}

	mappingsConfig := MappingsConfig{}
	err = yaml.Unmarshal([]byte(data), &mappingsConfig)
	if err != nil {
		log.Fatalf("Failed to Unmarshal yaml: %v", err)
	}

	return mappingsConfig
}

func GetDiscordCredentials() DiscordBotCredentials {
	ssmClient := lib.NewSSMClient()
	ssmResult, ssmErr := ssmClient.Param("discord_bot_keys", true).GetValue()
	if ssmErr != nil {
		log.Fatal(ssmErr)
	}
	var credentials DiscordBotCredentials
	err := json.Unmarshal([]byte(ssmResult), &credentials)
	if err != nil {
		log.Fatal(err)
	}

	return credentials
}
