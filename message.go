package eliotlibs

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func GetMessageParam(code string, accept string, name string) string {

	//TODO challenge english by default
	bundle := i18n.NewBundle(language.Spanish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	resurcePath := os.Getenv("RESOURCE_PATH")
	if accept == "en" {
		bundle.MustLoadMessageFile(resurcePath + "/en.json")
	} else {
		bundle.MustLoadMessageFile(resurcePath + "/es.json")
	}

	localizer := i18n.NewLocalizer(bundle, accept)
	result, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: code, TemplateData: map[string]interface{}{"Name": name}})
	if err != nil {
		log.Println(err)
		return code
	}
	return result
}

func GetMessage(code string) string {
	return GetMessageParam(code, "es", "")
}
