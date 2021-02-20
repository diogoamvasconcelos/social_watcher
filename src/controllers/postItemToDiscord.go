package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
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

	keywordConfig := KeywordConfig{}
	for _, mapping := range mappingsConfig.Keyword {
		if mapping.Value == item.Keyword {
			keywordConfig = mapping
			break
		}
	}

	// Get author data
	twitterUserDetails, err := GetTwitterUserDetails(item.Data.AuthorID)
	if err != nil {
		fmt.Println("GetTwitterUserDetails Error: ", err)
		twitterUserDetails = TwitterUserDetails{ID: item.Data.AuthorID, Followers: -1}
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

	var messageData discordgo.MessageSend
	messageData.Content += fmt.Sprintf("> New `%s` Twitter message (author followers: %d)", item.Keyword, twitterUserDetails.Followers)
	messageData.Content += "\n"
	messageData.Content += item.Link

	if item.Data.Text != item.Data.TranslatedText {
		/*
			messageData.Embed = &discordgo.MessageEmbed{
				Color:       1146986, // DARK_AQUA
				Title:       fmt.Sprintf("Translated message (lang: %s)", item.Data.Lang),
				Description: item.Data.TranslatedText,
			}
		*/
		messageData.Content += "\n"
		messageData.Content += fmt.Sprintf("Translated message (lang: %s)", item.Data.Lang)
		messageData.Content += "```"
		messageData.Content += item.Data.TranslatedText
		messageData.Content += "```"
	}

	for _, channel := range keywordConfig.Channels {
		messageResult, err := discordClient.ChannelMessageSendComplex(channel, &messageData)
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
