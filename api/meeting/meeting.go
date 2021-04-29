package meeting

import (
	"http/web"
	cache_meeting "huntsub/huntsub-map-server/cache/org/meeting"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/meeting"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type MeetingServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewMeetingServer() *MeetingServer {
	var s = &MeetingServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)

	return s
}

func (s *MeetingServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var p = &meeting.Meeting{}
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *MeetingServer) mustGetPost(r *http.Request) *meeting.Meeting {
	var id = r.URL.Query().Get("id")
	var p, err = cache_meeting.Get(id)
	web.AssertNil(err)
	return p
}

func (s *MeetingServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *MeetingServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newMeeting = &meeting.Meeting{}
	s.MustDecodeBody(r, &newMeeting)
	var p = s.mustGetPost(r)
	res, err := p.Update(newMeeting)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *MeetingServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := meeting.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = meeting.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *MeetingServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(meeting.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *MeetingServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	// var value = r.URL.Query().Get("value")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := meeting.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = meeting.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
