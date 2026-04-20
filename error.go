package eliotlibs

func FinsihApp(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrorWithParams interface {
	GetData() map[string]any
}

func GetParamsByError(err error) map[string]any {
	he, ok := err.(ErrorWithParams)
	if ok {
		return he.GetData()
	}
	return nil
}

type ErrorResponse struct {
	Code    string `json:"code" example:"UNAUTHORIZED_ERROR"`
	Message string `json:"message" example:"Unauthorized"`
}
