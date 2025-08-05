package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	auth_header := headers.Get("Authorization")
	if auth_header == "" {
		return "", fmt.Errorf("no authorization header")
	}

	token := strings.TrimPrefix(auth_header, "Bearer ")

	return token, nil
}
