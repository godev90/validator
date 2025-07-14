package faults

import "golang.org/x/text/language"

type LanguageTag string

var (
	Bahasa  LanguageTag = LanguageTag(language.Indonesian.String())
	English LanguageTag = LanguageTag(language.English.String())

	DefaultLocale = English
)

type LangPackage struct {
	Tag     LanguageTag
	Message string
}
