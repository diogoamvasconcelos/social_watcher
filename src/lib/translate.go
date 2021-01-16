package lib

import (
	"log"

	"github.com/aws/aws-sdk-go/service/translate"
)

func NewTranslateClient() *translate.Translate {
	// Create AWS Session
	sess, err := AWSSessions()
	if err != nil {
		log.Println(err)
		return nil
	}
	translatesvc := translate.New(sess)
	return translatesvc
}

func TranslateText(translatesvc *translate.Translate, text string, sourceLanguageCode string, targetLanguageCode string) (string, error) {
	res, err := translatesvc.Text(&translate.TextInput{
		Text:               &text,
		SourceLanguageCode: &sourceLanguageCode,
		TargetLanguageCode: &targetLanguageCode,
	})
	if err != nil {
		return "", err
	}

	return *res.TranslatedText, nil
}
