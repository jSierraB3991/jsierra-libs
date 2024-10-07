package eliotlibs

import (
	"net/http"
)

func GetHeader(r *http.Request, header string) []string {
	return r.Header.Values(header)
}

func GetLanguageHeader(r *http.Request) string {
	headerLanguage := GetHeader(r, "Accept-Language")
	idioma := "es"
	if len(headerLanguage) >= 1 {
		idioma = headerLanguage[0]
	}
	return idioma
}
