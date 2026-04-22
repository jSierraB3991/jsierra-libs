package eliotlibs

import (
	"errors"
	"net/http"
)

func StabilityCodeError(codeError string) string {
	if codeError == INVALID_FORMAT_ERROR_LITERAL_CONST {
		return INVALID_FORMAT_ERROR_CONST
	}
	if codeError == METHOD_NOT_ALLOWED_LITERAL_CONST {
		return METHOD_NOT_ALLOWED_CONST
	}
	return codeError
}

func GetMessageByError(codeError string, params map[string]any, request *http.Request) string {
	headerLang := GetLanguageHeader(request)
	return GetMessageMultipleParamParam(codeError, headerLang, params)
}

func UnwrapRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}
