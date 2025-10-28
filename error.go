package eliotlibs

func FinsihApp(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrorWithParam struct {
	ErrorString string
}

func NewErrorParam(code string, locale string, params map[string]interface{}) ErrorWithParam {
	errorString := GetMessageMultipleParamParam(code, locale, params)
	return ErrorWithParam{ErrorString: errorString}
}

func NewError(code string, locale string) ErrorWithParam {
	return NewErrorParam(code, locale, nil)
}

func (e ErrorWithParam) Error() string {
	return e.ErrorString
}
