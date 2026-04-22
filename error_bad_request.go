package eliotlibs

type BadRequestError struct {
	Err error
}

func (e BadRequestError) Error() string {
	return e.Err.Error()
}

func (e BadRequestError) Unwrap() error {
	return e.Err
}

func ErrorBadRequest(err error) error {
	return BadRequestError{Err: err}
}
