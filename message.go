package eliotlibs

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func GetMessageParam(code string, accept string, name string) string {
	params := map[string]any{"Name": name}
	return GetMessageMultipleParamParam(code, accept, params)
}
func GetMessageAccept(code string, accept string) string {
	return GetMessageMultipleParamParam(code, accept, map[string]any{})
}

func GetMessageMultipleParamParam(code string, accept string, params map[string]any) string {
	//TODO challenge english by default
	bundle := i18n.NewBundle(language.Spanish)
	bundle.RegisterUnmarshalFunc(JSON, json.Unmarshal)
	resurcePath := os.Getenv(RESOURCE_PATH)
	if accept == ENGLISH {
		bundle.MustLoadMessageFile(resurcePath + "/" + ENGLISH + ".json")
	} else {
		bundle.MustLoadMessageFile(resurcePath + "/" + SPANISH + ".json")
	}

	localizer := i18n.NewLocalizer(bundle, accept)
	result, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: code, TemplateData: params})
	if err != nil {
		log.Println(err)
		return code
	}
	return result
}

func GetMessage(code string) string {
	return GetMessageParam(code, SPANISH, "")
}
