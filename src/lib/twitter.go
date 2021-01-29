package lib

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type twitterBotCredentials struct {
	APIKey       string `json:"ApiKey"`
	APISecretKey string `json:"ApiSecretKey"`
	BearerToken  string `json:"BearerToken"`
}

func NewTwitterClient(token string) *http.Client {
	ctx := context.Background()
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}))
}
