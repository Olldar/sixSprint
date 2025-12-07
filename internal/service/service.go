package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

//функция распознает тип текста и вызывает соотв.
//ToText(morse string) string или ToMorse(text string) string
func AutoDetectAndConvert(text string) string {
	text = strings.TrimSpace(text)
	if text == "" {
		return ""
	}

	firstSimbol := text[:1]

	if firstSimbol == "."|| firstSimbol == "-" {
		return morse.ToText(text)
	} else {
		return morse.ToMorse(text)
	}
}