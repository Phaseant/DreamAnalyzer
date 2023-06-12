package repository

import (
	"github.com/sirupsen/logrus"
	tr "github.com/snakesel/libretranslate"
)

func translateTo(lang string, text string) (string, error) {
	translate := tr.New(tr.Config{
		Url: "http://localhost:5001", //local use
		// Url: "http://libretranslate:5001", //docker use
	})
	translated, err := translate.Translate(text, "auto", lang)
	if err != nil {
		logrus.Error("error while translating text: ", err)
		return "", err
	}
	return translated, nil
}

func detectLang(text string) (string, error) {
	translate := tr.New(tr.Config{
		Url: "http://localhost:5001", //local use
		// Url: "http://libretranslate:5001", //docker use
	})
	_, detected_lang, err := translate.Detect(text)
	if err != nil {
		logrus.Error("error while detecting language: ", err)
		return "", err
	}
	return detected_lang, nil
}
