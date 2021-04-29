package follow

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_follow "huntsub/huntsub-map-server/cache/org/action/follow"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/action/follow"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type FollowServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewFollowServer() *FollowServer {
	var s = &FollowServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/follow_number", s.HanldeCountFollowNumber)
	s.HandleFunc("/check", s.HanldeCheckFollowed)

	return s
}

func (s *FollowServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &follow.Follow{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *FollowServer) mustGetPost(r *http.Request) *follow.Follow {
	var id = r.URL.Query().Get("id")
	var p, err = cache_follow.Get(id)
	web.AssertNil(err)
	return p
}

func (s *FollowServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *FollowServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &follow.Follow{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *FollowServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
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
	res, ok := follow.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = follow.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *FollowServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(follow.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *FollowServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
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
	res, ok := follow.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = follow.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *FollowServer) HanldeCountFollowNumber(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var query = map[string]interface{}{
		"dtime":    0,
		"personid": sess.UserID,
	}
	var count, err = follow.Count(query)
	web.AssertNil(err)
	s.SendData(w, count)
}

func (s *FollowServer) HanldeCheckFollowed(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	personId := r.URL.Query().Get("personid")
	query := map[string]interface{}{
		"userid":   sess.UserID,
		"personid": personId,
	}

	var res, err = follow.GetFollow(query)
	web.AssertNil(err)
	s.SendData(w, res)
}
