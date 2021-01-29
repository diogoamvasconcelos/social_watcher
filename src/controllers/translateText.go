package controllers

import (
	"log"

	"github.com/diogoamvasconcelos/social_watcher/src/lib"
)

func TranslateToEnglish(originalText string, languageCode string) string {
	if languageCode == "en" || languageCode == "und" {
		return originalText
	}

	log.Printf("Translating text:%s", originalText)

	translateClient := lib.NewTranslateClient()
	// souceLanguage; "auto" (Amazon Comprehend not available in Stockholm region)
	// res, err := lib.TranslateText(translateClient, originalText, "auto", "en")
	res, err := lib.TranslateText(translateClient, originalText, languageCode, "en")
	if err != nil {
		log.Fatal("Failed to translate text:", err)
	}

	return res
}
