package calendar

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_calendar "huntsub/huntsub-map-server/cache/org/calendar"
	cache_usr "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/config/cons"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/calendar"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type CalendarServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewCalendarServer() *CalendarServer {
	var s = &CalendarServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/getday", s.HanldeGetDay)
	s.HandleFunc("/confirm", s.HanldeConfirmDeleting)
	s.HandleFunc("/get_by_month", s.HanldeGetByMonth)

	return s
}

func (s *CalendarServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var u, err = cache_usr.Get(sess.UserID)
	web.AssertNil(err)
	var p = &calendar.Calendar{}
	p.Priority = cons.Level_0
	p.UserID = u.ID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *CalendarServer) mustGetPost(r *http.Request) *calendar.Calendar {
	var id = r.URL.Query().Get("id")
	var p, err = cache_calendar.Get(id)
	web.AssertNil(err)
	return p
}

func (s *CalendarServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *CalendarServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newC = &calendar.Calendar{}
	s.MustDecodeBody(r, &newC)
	var p = s.mustGetPost(r)
	res, err := p.Update(newC)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *CalendarServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
		"$or": []interface{}{
			bson.M{
				"service_provider.userid": sess.UserID,
			},
			bson.M{
				"service_caller.userid": sess.UserID,
			},
		},
	}
	res, ok := calendar.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)

	data := []*CalendarForm{}
	for _, c := range res {
		data = append(data, NewCalendarForm(c))
	}

	var count, _ok = calendar.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  data,
	})
}

func (s *CalendarServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(calendar.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *CalendarServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	// var key = r.URL.Query().Get("key")
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime":  0,
		"userid": sess.UserID,
	}
	res, ok := calendar.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = calendar.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
func (s *CalendarServer) HanldeGetDay(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var day = r.URL.Query().Get("day")

	var query = map[string]interface{}{
		"dtime":      0,
		"userid":     sess.UserID,
		"meetingday": day,
	}
	res, ok := calendar.GetMany(query)
	web.AssertNil(ok)
	var count, _ok = calendar.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *CalendarServer) HanldeConfirmDeleting(w http.ResponseWriter, r *http.Request) {
	var cl = &calendar.Calendar{}
	s.MustDecodeBody(r, cl)
	obj.EmitConfirm(cl)
	s.SendData(w, nil)
}

func (s *CalendarServer) HanldeGetByMonth(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var from = r.URL.Query().Get("from")
	var to = r.URL.Query().Get("to")

	var query = map[string]interface{}{
		"dtime":      0,
		"userid":     sess.UserID,
		"meetingday": bson.M{"$gte": from, "$lte": to},
	}
	res, err := calendar.GetMany(query)
	web.AssertNil(err)

	s.SendData(w, res)
}
