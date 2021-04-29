package ustatic

import (
	"strings"
)

type UStatic struct {
	Remote string
	Local  string
}

func NewUStatic(remote string, local string) *UStatic {
	u := &UStatic{
		Remote: remote,
		Local:  local,
	}
	return u
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func (u *UStatic) getLink(uri string) string {
	return singleJoiningSlash(u.Remote, uri)
}
