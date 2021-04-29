package post

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_post "huntsub/huntsub-map-server/cache/org/post"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/post"
	userReport "huntsub/huntsub-map-server/o/report/user"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource
var acceptionPercent = 0.05

type PostServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewPostServer() *PostServer {
	var s = &PostServer{
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

func (s *PostServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	var p = &post.Post{}
	p.UserID = sess.UserID
	p.Location = u.Location
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *PostServer) mustGetPost(r *http.Request) *post.Post {
	var id = r.URL.Query().Get("id")
	var p, err = cache_post.Get(id)
	web.AssertNil(err)
	return p
}

func (s *PostServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetPost(r)
	s.SendData(w, p)
}

func (s *PostServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newPost = &post.Post{}
	s.MustDecodeBody(r, &newPost)
	var p = s.mustGetPost(r)
	res, err := p.Update(newPost)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *PostServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	sess := session.MustAuthScope(r)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(err)
	_type := r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
		// "$or": []interface{}{
		// 	bson.M{
		// 		"location": bson.M{
		// 			"$near": bson.M{
		// 				"$geometry": bson.M{
		// 					"type":        "Point",
		// 					"coordinates": u.Location.Coordinates,
		// 				},
		// 				"$maxDistance": 1000000,
		// 			},
		// 		},
		// 	},
		// 	bson.M{

		// 	},
		// },
		"location.coordinates": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]interface{}{
						u.Location.Coordinates[0],
						u.Location.Coordinates[1],
					}, 10000,
				},
			},
		},
		"type": _type,
	}
	res, ok := post.GetPostNearMe(query, skip, limit)
	web.AssertNil(ok)
	data := []*PostForm{}

	//Get UserReport Record
	urp, err := userReport.GetUserReport(map[string]interface{}{
		"dtime":  0,
		"userid": u.GetID(),
	})
	for _, c := range res {
		//check post has on user's interest
		index := common.CheckExist(c.Category, urp)
		if index == -1 {
			data = append(data, NewPostForm(c, sess.UserID))
		} else {
			if urp.AnalysisFields[index].Percent < acceptionPercent {

			} else {
				data = append(data, NewPostForm(c, sess.UserID))
			}
		}
	}
	s.SendData(w, data)
}

func (s *PostServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetPost(r)
	web.AssertNil(post.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *PostServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	userid := r.URL.Query().Get("userid")
	skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var query = map[string]interface{}{
		"dtime":  0,
		"userid": userid,
	}
	res, ok := post.GetPostNearMe(query, skip, limit)
	web.AssertNil(ok)
	data := []*PostForm{}
	for _, c := range res {
		data = append(data, NewPostForm(c, userid))
	}
	s.SendData(w, data)
}
