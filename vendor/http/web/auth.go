package web

import (
	"net/http"
	"strings"
)

const bearerHeader = "Bearer "
const accessToken = "token"
const errInvalidAccessToken = Unauthorized("invalid access token")
const errMissingAccessToken = Unauthorized("missing access token")

func GetToken(r *http.Request) string {
	var authHeader = r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, bearerHeader) {
		return strings.TrimPrefix(authHeader, bearerHeader)
	}
	return r.URL.Query().Get(accessToken)
}

func MustHaveToken(token string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := GetToken(r)
		if len(current) < 1 {
			panic(errMissingAccessToken)
		}
		if token != current {
			panic(errInvalidAccessToken)
		}
		handler.ServeHTTP(w, r)
	})
}
