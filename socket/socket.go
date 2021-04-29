package socket

import (
	"http/socket"
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/user"
	"net/http"
)

var obj = event.ObjectEventSource

type SocketServer struct {
	web.JsonServer
	*http.ServeMux
	Ws *socket.Hub
}

func NewSocketServer(Ws *socket.Hub) *SocketServer {
	var s = &SocketServer{
		ServeMux: http.NewServeMux(),
		Ws:       Ws,
	}
	s.HandleFunc("/join", s.HandleJoinSocket)
	return s
}

func (s *SocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Huntsub-Token"
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", allowedHeaders)
	header.Set("Access-Control-Expose-Headers", "Authorization")
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}

func (s *SocketServer) HandleJoinSocket(w http.ResponseWriter, r *http.Request) {
	var roomId = r.URL.Query().Get("id")
	var sess = session.MustGet(r)
	var u, err = cache_user.Get(sess.UserID)
	web.AssertNil(err)
	var newUser = &user.User{}
	newUser.StatusActive.Online = true
	newUser.StatusActive.RoomID = roomId
	newUser.StatusActive.TimeOut = 0
	res, err := u.Update(newUser)
	web.AssertNil(err)
	obj.EmitUpdate(res)

	socket.ServeWs(s.Ws, w, r, roomId, res)
}
