package auth

import (
	"net/http"
	"strings"
)

const (
	loopIP = "127.0.0.1"
)

func IsLocal(r *http.Request) bool {
	return strings.HasPrefix(r.RemoteAddr, loopIP)
}
