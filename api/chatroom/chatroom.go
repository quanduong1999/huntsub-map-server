package chatroom

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache "huntsub/huntsub-map-server/cache/org/chatroom"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/chatroom"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type ChatRoomServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewChatRoomServer() *ChatRoomServer {
	var s = &ChatRoomServer{
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

func (s *ChatRoomServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var p = &chatroom.ChatRoom{}
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *ChatRoomServer) mustGetPost(r *http.Request) *chatroom.ChatRoom {
	var id = r.URL.Query().Get("id")
	var p, err = cache.Get(id)
	web.AssertNil(err)
	return p
}

func (s *ChatRoomServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *ChatRoomServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &chatroom.ChatRoom{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *ChatRoomServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
		"users": bson.M{"$elemMatch": bson.M{"userid": sess.UserID}},
	}
	res, ok := chatroom.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	data := NewChatRoomForm(res, sess.UserID)
	s.SendData(w, data)
}

func (s *ChatRoomServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(chatroom.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *ChatRoomServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
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
	res, ok := chatroom.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = chatroom.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
