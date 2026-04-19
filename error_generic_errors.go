package eliotlibs

type MethodNotAllowedError struct {
}

func (e *MethodNotAllowedError) Error() string {
	return "METHOD_NOT_ALLOWED"
}
