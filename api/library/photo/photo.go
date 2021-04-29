package photo

import (
	"fmt"
	"http/web"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/user"
	"huntsub/huntsub-map-server/x/timer"
	"net/http"
	"os"
	// "os/user"
)

var obj = event.ObjectEventSource

type PhotoServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewPhotoServer() *PhotoServer {
	var s = &PhotoServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/getall", s.HandleGetall)

	return s
}

func (s *PhotoServer) HandleGetall(w http.ResponseWriter, r *http.Request) {
	arr := []string{}
	var id = r.URL.Query().Get("id")
	var _type = r.URL.Query().Get("type")

	u, err := user.GetByID(id)
	web.AssertNil(err)
	var path string
	if _type == "img" {
		path = fmt.Sprintf("static/upload/user/%s/%s/img", timer.TimeToDay(u.MTime), id)
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
