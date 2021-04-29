package whfeedback

import (
	"http/web"
	cache_whfeedback "huntsub/huntsub-map-server/cache/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type WHFeedbackServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewWHFeedbackServer() *WHFeedbackServer {
	var s = &WHFeedbackServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)

	return s
}

func (s *WHFeedbackServer) mustGetPost(r *http.Request) *whfeedback.WHFeedback {
	var id = r.URL.Query().Get("id")
	var p, err = cache_whfeedback.Get(id)
	web.AssertNil(err)
	return p
}

func (s *WHFeedbackServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *WHFeedbackServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newWHFb = &whfeedback.WHFeedback{}
	s.MustDecodeBody(r, &newWHFb)
	var p = s.mustGetPost(r)
	res, err := p.Update(newWHFb)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *WHFeedbackServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := whfeedback.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = whfeedback.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *WHFeedbackServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(whfeedback.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *WHFeedbackServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	// var value = r.URL.Query().Get("value")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := whfeedback.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = whfeedback.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
