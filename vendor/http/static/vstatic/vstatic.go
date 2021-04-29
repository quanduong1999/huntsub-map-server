package vstatic

import (
	"net/http"
	"strings"
)

type VersionStatic struct {
	internal *internalVersionStatic
}

func NewVersionStatic(dir string) *VersionStatic {
	vs := &VersionStatic{}
	vs.internal = newInternalVersionStatic(dir)
	vs.internal.reactivate()
	return vs
}

const internalVersionStaticPrefix = "/__"

func (v *VersionStatic) isInternalRequest(r *http.Request) bool {
	return strings.HasPrefix(r.URL.Path, internalVersionStaticPrefix)
}

func (v *VersionStatic) serveInternalRequest(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, internalVersionStaticPrefix)
	v.internal.ServeHTTP(w, r)
}

func (v *VersionStatic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if v.isInternalRequest(r) {
		v.serveInternalRequest(w, r)
		return
	}
	if v.internal.current == nil {
		http.Error(w, "vstatic not found", http.StatusNotFound)
		return
	}
	v.internal.current.ServeHTTP(w, r)
}
