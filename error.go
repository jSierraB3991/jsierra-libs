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
