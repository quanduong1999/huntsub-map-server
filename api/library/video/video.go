package video

import (
	"fmt"
	"http/web"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/user"
	"huntsub/huntsub-map-server/x/timer"
	"net/http"
	"os"
)

var obj = event.ObjectEventSource

type VideoServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewVideoServer() *VideoServer {
	var s = &VideoServer{
		ServeMux: http.NewServeMux(),
	}

	s.HandleFunc("/getall", s.HanldeGetAll)

	return s
}

func (s *VideoServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var _type = r.URL.Query().Get("type")
	arr := []string{}

	u, err := user.GetByID("id")
	web.AssertNil(err)

	var path string
	if _type == "video" {
		path = fmt.Sprintf("static/upload/user/%s/%s/video", timer.TimeToDay(u.MTime), id)
	} else {
		path = ""
	}

	x, err := os.Open(path)
	web.AssertNil(err)
	y, err := x.Readdir(0)
	web.AssertNil(err)
	for _, i := range y {
		fmt.Println(i.Name())
		arr = append(arr, i.Name())
	}
	s.SendData(w, arr)
}
