package eliotlibs

func RunValidations(req interface{}, validations ...func(req interface{}) error) error {
	for _, validation := range validations {
		err := validation(req)
		if err != nil {
			return err
		}
	}
	return nil
}
