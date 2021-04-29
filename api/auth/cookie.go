package auth

import (
	"net/http"
)

const maxCookieAge = 30 * 24 * 3600

func setCookie(w http.ResponseWriter, name string, value string) {
	var cookie = http.Cookie{Name: name, Value: value, HttpOnly: false, MaxAge: maxCookieAge, Path: "/"}
	http.SetCookie(w, &cookie)
}

func clearCookie(w http.ResponseWriter, name string) {
	var cookie = http.Cookie{Name: name, Value: "", MaxAge: -1}
	http.SetCookie(w, &cookie)
}
