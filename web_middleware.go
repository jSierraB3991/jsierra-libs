package eliotlibs

import (
	"strings"
	"unicode"
)

func PublicMiddleWare(route string, method string) bool {

	if (route == "/" && method == "GET") || method == "OPTIONS" {
		return true
	}

	var NO_AUTH_NEED = []string{"public"}

	for _, value := range NO_AUTH_NEED {
		if strings.Contains(route, value) {
			return true
		}
	}
	return false
}

func RemovePrefixInString(authHeader, prefix string) string {
	if strings.HasPrefix(authHeader, prefix) {
		return strings.TrimPrefix(authHeader, prefix)
	}
	return authHeader
}

func isNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func PathsEqual(pattern, path string) bool {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := range patternParts {
		if strings.HasPrefix(patternParts[i], ":") {
			// Permitir cualquier número para los parámetros dinámicos
			if !isNumber(pathParts[i]) {
				return false
			}
			continue
		}
		if patternParts[i] != pathParts[i] {
			return false
		}
	}
	return true
}
