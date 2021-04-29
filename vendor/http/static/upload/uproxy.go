package upload

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type FallbackHandler struct {
	Normal   http.Handler
	Path     string
	Fallback http.Handler
	FPath    string
}

func NewFallbackHandler(normal http.Handler, fallback http.Handler, path string) *FallbackHandler {
	var s = &FallbackHandler{
		Normal:   normal,
		Fallback: fallback,
		FPath:    path + "/__",
	}
	return s
}

func (s *FallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var useFallback = strings.HasPrefix(r.URL.Path, s.FPath)
	if useFallback {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, s.FPath)
		w.Header().Add("Fallback", "true")
		s.Fallback.ServeHTTP(w, r)
	} else {
		w.Header().Add("Fallback", "false")
		s.Normal.ServeHTTP(w, r)
	}
}

type UploadProxy struct {
	Remote *url.URL
	Local  string
}

func NewUploadProxy(remote *url.URL, folder string, path string) http.Handler {
	// /data/file
	var localFile = NewUploadFileServer(folder, 0)
	if remote == nil {
		return http.StripPrefix(path, localFile)
	}
	// logger.Infof(0, "proxy /upload to %s", remoteHost)
	var proxy = httputil.NewSingleHostReverseProxy(remote)
	return NewFallbackHandler(proxy, localFile, path)
}
