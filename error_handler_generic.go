package eliotlibs

import "net/http"

func validateWithError(mesasge string, codReturn int, err error) *int {
	if mesasge == err.Error() {
		return &codReturn
	}
	return nil
}

func RunMultipleValidationCode(message string, valueReturn int, errs ...error) *int {
	for _, err := range errs {
		code := validateWithError(message, valueReturn, err)
		if code != nil {
			return code
		}
	}
	return nil
}

func GetErrorCodeMethodNotAllowed(message string) *int {
	return RunMultipleValidationCode(message, http.StatusMethodNotAllowed,
		MethodNotAllowedError{},
	)
}

func GetErrorCodeUnprocessableEntity(message string) *int {
	return runMultipleValidationCode(message, http.StatusUnprocessableEntity,
		InvalidFormatError{},
	)
}

func RunValidationCodeSpecificErrors(message string, funcs ...func(string) *int) int {
	for _, fn := range funcs {
		code := fn(message)
		if code != nil {
			return *code
		}
	}
	return http.StatusBadRequest
}
