package eliotlibs

func FinsihApp(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrorWithParam interface {
	Error() string
	Params() []string
}
