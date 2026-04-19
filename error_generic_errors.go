package eliotlibs

type MethodNotAllowedError struct {
}

func (MethodNotAllowedError) Error() string {
	return "METHOD_NOT_ALLOWED"
}

type InvalidFormatError struct {
}

func (InvalidFormatError) Error() string {
	return "INVALID_FORMAT"
}
