package comment

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_comment "huntsub/huntsub-map-server/cache/org/comment"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/comment"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type CommentServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewCommentServer() *CommentServer {
	var s = &CommentServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/getcommentsbypost", s.HanldeGetCommentsByPost)

	return s
}

func (s *CommentServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &comment.Comment{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *CommentServer) mustGetPost(r *http.Request) *comment.Comment {
	var id = r.URL.Query().Get("id")
	var p, err = cache_comment.Get(id)
	web.AssertNil(err)
	return p
}

func (s *CommentServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *CommentServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newNewCmt = &comment.Comment{}
	s.MustDecodeBody(r, &newNewCmt)
	var p = s.mustGetPost(r)
	res, err := p.Update(newNewCmt)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *CommentServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := comment.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = comment.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *CommentServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(comment.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *CommentServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	// var value = r.URL.Query().Get("value")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := comment.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = comment.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *CommentServer) HanldeGetCommentsByPost(w http.ResponseWriter, r *http.Request) {
	var postID = r.URL.Query().Get("id")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime":         0,
		"postid":        postID,
		"commentrootid": bson.M{"$eq": ""},
	}

	var data, ok = NewCommentForm(postID, _type, skip, limit)
	web.AssertNil(ok)
	var count, _ok = comment.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  data,
	})
}
