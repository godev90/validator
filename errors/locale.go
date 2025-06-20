package errors

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

func SupportedTags() []LanguageTag {
	return []LanguageTag{
		Bahasa,
		English,
	}
}

func IsSupported(tag LanguageTag) bool {
	for _, t := range SupportedTags() {
		if t == tag {
			return true
		}
	}
	return false
}

var langPackages = make(map[string]LangPackage)

func RegisterLangErrorPackage(key string, tag LanguageTag, message string) {
	langPackages[key] = LangPackage{
		Tag:     tag,
		Message: message,
	}
}
