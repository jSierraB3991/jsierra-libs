package eliotlibs

func RunValidations(validations ...func() error) error {
	for _, validation := range validations {
		err := validation()
		if err != nil {
			return err
		}
	}
	return nil
}
