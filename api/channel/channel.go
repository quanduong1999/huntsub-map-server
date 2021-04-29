package channel

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/action/follow"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/chatroom"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/rank"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type ChannelServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewChannelServer() *ChannelServer {
	var s = &ChannelServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByUserID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/search_by_job", s.HanldeSearchByJob)
	s.HandleFunc("/get_all_job_list", s.HandleGetAllJobList)

	return s
}

func (s *ChannelServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var p = &channel.Channel{}
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *ChannelServer) mustGetChannel(r *http.Request) *channel.Channel {
	var id = r.URL.Query().Get("id")
	var p, err = channel.GetChannel(map[string]interface{}{
		"dtime":  0,
		"userid": id,
	})
	web.AssertNil(err)
	return p
}

func (s *ChannelServer) mustGetChannelByID(r *http.Request) *channel.Channel {
	var id = r.URL.Query().Get("id")
	var p, err = channel.GetByID(id)
	web.AssertNil(err)
	return p
}

func (s *ChannelServer) HandleGetByUserID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetChannel(r)
	s.SendData(w, p)
}

func (s *ChannelServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newChannel = &channel.Channel{}
	s.MustDecodeBody(r, &newChannel)
	var p = s.mustGetChannelByID(r)
	res, err := p.Update(newChannel)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *ChannelServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := channel.GetPaginationTop(query, skip, limit)
	web.AssertNil(ok)
	var count, _ok = channel.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *ChannelServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetChannel(r)
	web.AssertNil(channel.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *ChannelServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	// var value = r.URL.Query().Get("value")
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := channel.GetPaginationTop(query, skip, limit)
	web.AssertNil(ok)
	var count, _ok = channel.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *ChannelServer) HanldeSearchByJob(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	web.AssertNil(err)
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	web.AssertNil(err)

	var job = r.URL.Query().Get("job")
	var query = bson.M{
		"dtime": 0,
		"$or": []interface{}{
			bson.M{
				"job": bson.M{"$regex": bson.RegEx{Pattern: job}},
			},
			bson.M{
				"describejob": bson.M{"$elemMatch": bson.M{"$regex": bson.RegEx{Pattern: job}}},
			},
		},
		"location.coordinates": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]interface{}{
						lng,
						lat,
					}, 1000,
				},
			},
		},
		"userid": bson.M{"$ne": sess.UserID},
	}

	cns, err := channel.GetPaginationTop(query, skip, limit)
	web.AssertNil(err)
	s.SendData(w, cns)

	data := []*InformationUser{}
	for _, c := range cns {
		//Get User
		u, err := cache_user.Get(c.UserID)
		web.AssertNil(err)
		var where = bson.M{
			"dtime":  0,
			"userid": c.UserID,
		}
		//Get Rank
		ran, err := rank.GetRank(where)
		web.AssertNil(err)
		whfb, err := whfeedback.GetWHFeedBack(where)
		web.AssertNil(err)
		fl, err := follow.GetFollow(map[string]interface{}{
			"dtime":    0,
			"userid":   sess.UserID,
			"personid": c.UserID,
		})
		followed := false
		followID := ""
		if err != nil {

		} else {
			followID = fl.ID
			followed = fl.IsFollow
		}
		//Checking has exits a room chat
		roomID := ""
		queyChatRoom := bson.M{
			"dtime": 0,
			"$and": []interface{}{
				bson.M{
					"users": bson.M{"$elemMatch": bson.M{"userid": c.UserID}},
				},
				bson.M{
					"users": bson.M{"$elemMatch": bson.M{"userid": sess.UserID}},
				},
			},
		}
		room, err := chatroom.GetOne(queyChatRoom)
		if err != nil {

		} else {
			roomID = room.ID
		}
		data = append(data, NewUserInfo(u, c, ran, whfb, followed, followID, roomID))
	}

	s.SendData(w, data)

}

func (s *ChannelServer) HandleGetAllJobList(w http.ResponseWriter, r *http.Request) {
	lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	web.AssertNil(err)
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	web.AssertNil(err)
	var sess = session.MustAuthScope(r)
	// var u = cache_user.Get(sess.UserID)

	var query = bson.M{
		"dtime": 0,
		// "$or": []interface{}{
		// 	bson.M{
		// 		"job": bson.M{"$regex": bson.RegEx{Pattern: job}},
		// 	},
		// 	bson.M{
		// 		"describejob": bson.M{"$elemMatch": bson.M{"$regex": bson.RegEx{Pattern: job}}},
		// 	},
		// },
		"location.coordinates": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]interface{}{
						lng,
						lat,
					}, 1000,
				},
			},
		},
		"userid": bson.M{"$ne": sess.UserID},
	}
	var group = bson.M{}
	group["_id"] = bson.M{
		"job": "$job",
	}
	group["count"] = bson.M{"$sum": 1}
	group["job"] = bson.M{"$first": "$job"}

	var pipeline = []bson.M{
		{"$match": query},
		{"$group": group},
	}

	var res = []*JobList{}
	err = channel.TableChannel.C().Pipe(pipeline).All(&res)
	web.AssertNil(err)

	s.SendData(w, res)
}
