package like

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/action/like"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type LikeServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewLikeServer() *LikeServer {
	var s = &LikeServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/check", s.HanldeCheckLike)

	return s
}

func (s *LikeServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &like.Like{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *LikeServer) mustGetPost(r *http.Request) *like.Like {
	var id = r.URL.Query().Get("id")
	var userid = r.URL.Query().Get("userid")
	var p, err = like.GetLike(map[string]interface{}{
		"dtime":  0,
		"postid": id,
		"userid": userid,
	})
	web.AssertNil(err)
	return p
}

func (s *LikeServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *LikeServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &like.Like{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *LikeServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := like.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = like.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *LikeServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(like.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *LikeServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
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
	res, ok := like.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = like.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *LikeServer) HanldeCheckLike(w http.ResponseWriter, r *http.Request) {
	postid := r.URL.Query().Get("postid")
	userid := r.URL.Query().Get("userid")

	query := map[string]interface{}{
		"dtime":  0,
		"postid": postid,
		"userid": userid,
	}

	_, err := like.GetLike(query)
	if err != nil {
		s.SendData(w, false)
	} else {
		s.SendData(w, true)
	}
}
