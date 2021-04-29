package httpserver

import (
	"net/http"
)

type ServerStatic struct {
	server *http.ServeMux
	dir    string
}

func NewServerStatic(dir string) *ServerStatic {
	var server = http.NewServeMux()
	server.Handle("/", http.FileServer(http.Dir(dir)))
	return &ServerStatic{server: server}
}

func (s *ServerStatic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.server.ServeHTTP(w, r)
}
