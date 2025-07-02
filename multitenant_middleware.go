package eliotlibs

import (
	"context"
	"net/http"
	"regexp"
)

func getContext(req *http.Request, tenantId string) context.Context {

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, tenantId)
	if !matched {
		return req.Context()
	}

	// Inyectar tenant en el contexto de la request
	return context.WithValue(req.Context(), ContextTenantKey, tenantId)
}

func GetContextWithTenantForMultenant(req *http.Request, schemasAllow []string) context.Context {
	tenantID := req.Header.Get("X-Tenant-ID")

	for _, v := range schemasAllow {
		if v == tenantID {
			return getContext(req, tenantID)
		}
	}
	return nil
}
