package eliotlibs

type MethodNotAllowedError struct {
}

func (e *MethodNotAllowedError) Error() string {
	return "METHOD_NOT_ALLOWED"
}

type InvalidFormatError struct {
}

func (e *InvalidFormatError) Error() string {
	return "INVALID_FORMAT"
}
