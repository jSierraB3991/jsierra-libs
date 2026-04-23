package eliotlibs

type NotDatabaseConfigurateError struct{}

func (NotDatabaseConfigurateError) Error() string {
	return "NOT_DATABASE_CONFIGURATE_ERROR"
}

type InvalidEncryptData struct{}

func (InvalidEncryptData) Error() string {
	return "INVALID_ENCRYPT_DATA"
}

type InvalidFunctionMigrationError struct{}

func (InvalidFunctionMigrationError) Error() string {
	return "INVALID_FUNCTION_MIGRATION_ERROR"
}

type InvalidTenantDataError struct{}

func (InvalidTenantDataError) Error() string {
	return "INVALID_TENANT_DATA_ERROR"
}

type InvalidCharactersInTenantError struct{}

func (InvalidCharactersInTenantError) Error() string {
	return "INVALID_CHARACTERS_IN_TENANT_ERROR"
}
