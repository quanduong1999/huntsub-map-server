package share

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_share "huntsub/huntsub-map-server/cache/org/action/share"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/action/share"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type ShareServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewShareServer() *ShareServer {
	var s = &ShareServer{
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

func (s *ShareServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &share.Share{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *ShareServer) mustGetPost(r *http.Request) *share.Share {
	var id = r.URL.Query().Get("id")
	var p, err = cache_share.Get(id)
	web.AssertNil(err)
	return p
}

func (s *ShareServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *ShareServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &share.Share{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *ShareServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := share.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = share.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *ShareServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(share.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *ShareServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime":  0,
		"userid": sess.UserID,
	}
	res, ok := share.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = share.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
