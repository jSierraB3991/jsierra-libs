package eliotlibs

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/golang-jwt/jwt"
)

func PublicMiddleWare(route string, method string) bool {

	var NO_AUTH_NEED = []string{"public"}

	if method == "OPTIONS" {
		return true
	}

	for _, value := range NO_AUTH_NEED {
		if strings.Contains(route, value) {
			return true
		}
	}
	return false
}

func GetClaimByToken(tokenString, claim string) (interface{}, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(removeBearerPrefix(tokenString), jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	// Access the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims[claim], nil
	}
	return nil, errors.New("INVALID_CLAIMS")
}

// Función para eliminar el prefijo "Bearer " del encabezado de autorización
func removeBearerPrefix(authHeader string) string {
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return authHeader
}

func GetHeaderMail(requet *http.Request) (string, error) {
	var email string
	emailInterface, err := GetClaimByToken(requet.Header[HeaderAuthorization][0], "email")
	if err != nil {
		return "", err
	}

	if emailInterface != nil {
		return emailInterface.(string), nil
	}

	return email, nil
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
