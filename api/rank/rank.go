package rank

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_rank "huntsub/huntsub-map-server/cache/org/rank"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/rank"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type RankServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewRankServer() *RankServer {
	var s = &RankServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/get_by_userid", s.HandleGetByUserID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/top", s.HandleGetTop)

	return s
}

func (s *RankServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &rank.Rank{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *RankServer) mustGetPost(r *http.Request) *rank.Rank {
	var id = r.URL.Query().Get("id")
	var p, err = cache_rank.Get(id)
	web.AssertNil(err)
	return p
}

func (s *RankServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *RankServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &rank.Rank{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *RankServer) HandleGetTop(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var query = map[string]interface{}{
		"dtime": 0,
		"location.coordinates": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]interface{}{
						u.Location.Coordinates[0],
						u.Location.Coordinates[1],
					}, 1000,
				},
			},
		},
	}
	res, ok := rank.GetPaginationTop(query, skip, limit)
	web.AssertNil(ok)
	data := []*RankForm{}
	for _, c := range res {
		data = append(data, NewRankForm(c))
	}
	var count, _ok = rank.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  data,
	})
}

func (s *RankServer) HandleGetByUserID(w http.ResponseWriter, r *http.Request) {
	var userID = r.URL.Query().Get("id")
	var where = map[string]interface{}{
		"dtime":  0,
		"userid": userID,
	}

	res, err := rank.GetRank(where)
	web.AssertNil(err)
	s.SendData(w, res)
}
