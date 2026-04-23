package eliotlibs

import (
	"context"
	"regexp"
	"strings"
)

type contextKey string

const (
	ContextTenantKey contextKey = "tenant_id"
	HEADER_TENAT_ID  string     = "X-Tenant-ID"
)

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(ContextTenantKey).(string)
	return v, ok
}

func QuoteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}

func WithTenant(ctx context.Context) (*string, error) {
	tenant, isOk := FromContext(ctx)
	if !isOk {
		return nil, InvalidTenantDataError{}
	}

	err := ValidateTenant(tenant)
	if err != nil {
		return nil, err
	}
	// Quoteo seguro
	quotedSchema := QuoteIdentifier(tenant)
	return &quotedSchema, nil

}

func ValidateTenant(tenant string) error {
	// Validar de tenant ID
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, tenant); !ok {
		return InvalidCharactersInTenantError{}
	}

	// Validación del nombre del schema
	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`).MatchString(tenant) {
		return InvalidCharactersInTenantError{}
	}
	return nil

}
