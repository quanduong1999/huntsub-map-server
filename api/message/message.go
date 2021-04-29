package message

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"

	// cache_post "huntsub/huntsub-map-server/cache/org/message"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/message"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type MessageServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewMessageServer() *MessageServer {
	var s = &MessageServer{
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

func (s *MessageServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &message.Message{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *MessageServer) mustGetPost(r *http.Request) *message.Message {
	var id = r.URL.Query().Get("id")
	var p, err = message.GetByID(id)
	web.AssertNil(err)
	return p
}

func (s *MessageServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *MessageServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &message.Message{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *MessageServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var roomid = r.URL.Query().Get("roomid")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime":  0,
		"roomid": roomid,
	}
	res, ok := message.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = message.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *MessageServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(message.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *MessageServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
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
	res, ok := message.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = message.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
