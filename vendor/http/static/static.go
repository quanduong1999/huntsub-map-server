package static

import (
	"net/http"
)

var cacheControl = "public,maxage=900"

func addHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", cacheControl)
		h.ServeHTTP(w, r)
	})
}

func NewStatic(dir string) http.Handler {
	var withoutGzip = http.FileServer(http.Dir(dir))
	return addHeader(withoutGzip)
}

var offStatic = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "off static", http.StatusNotFound)
})
