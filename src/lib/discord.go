package lib

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewDiscordClient(token string) (*discordgo.Session, error) {
	// Create a new Discord session using the provided bot token.
	discordGo, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return nil, err
	}

	return discordGo, nil
}
