package notification

import (
	"http/web"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/notification"
	"net/http"
)

var obj = event.ObjectEventSource

type NotificationManagementServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewNotificationManagementServer() *NotificationManagementServer {
	var s = &NotificationManagementServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)

	return s
}

func (s *NotificationManagementServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var noti = &notification.NotifiationManagement{}
	s.MustDecodeBody(r, noti)
	web.AssertNil(noti.Create())
	obj.EmitCreate(noti)
	s.SendData(w, noti)
}

func (s *NotificationManagementServer) mustGetNoti(r *http.Request) *notification.NotifiationManagement {
	var id = r.URL.Query().Get("id")
	var p, err = notification.GetNotification(map[string]interface{}{
		"dtime":  0,
		"userid": id,
	})
	web.AssertNil(err)
	return p
}

func (s *NotificationManagementServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetNoti(r)
	s.SendData(w, p)
}

func (s *NotificationManagementServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newNoti = &notification.NotifiationManagement{}
	s.MustDecodeBody(r, &newNoti)
	println(newNoti.NotificationNumber)
	println(newNoti.MessageNumber)
	var p = s.mustGetNoti(r)
	res, err := p.Update(newNoti)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}
