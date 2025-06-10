package eliotlibs

import (
	"context"
	"net/http"
	"regexp"
)

func GetContextWithTenantForMultenant(req *http.Request) context.Context {
	tenantID := req.Header.Get("X-Tenant-ID")
	if tenantID == "" {
		tenantID = "public"
	}

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, tenantID)
	if !matched {
		return req.Context()
	}

	// Inyectar tenant en el contexto de la request
	return context.WithValue(req.Context(), ContextTenantKey, tenantID)
}
