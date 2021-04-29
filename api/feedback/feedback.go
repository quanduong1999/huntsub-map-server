package feedback

import (
	"http/web"
	cache_feedback "huntsub/huntsub-map-server/cache/org/feedback"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/feedback"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type FeedbackServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewFeedbackServer() *FeedbackServer {
	var s = &FeedbackServer{
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

func (s *FeedbackServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var p = &feedback.Feedback{}
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *FeedbackServer) mustGetPost(r *http.Request) *feedback.Feedback {
	var id = r.URL.Query().Get("id")
	var p, err = cache_feedback.Get(id)
	web.AssertNil(err)
	return p
}

func (s *FeedbackServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *FeedbackServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newFb = &feedback.Feedback{}
	s.MustDecodeBody(r, &newFb)
	var p = s.mustGetPost(r)
	res, err := p.Update(newFb)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *FeedbackServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := feedback.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = feedback.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *FeedbackServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(feedback.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *FeedbackServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	// var value = r.URL.Query().Get("value")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := feedback.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = feedback.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
