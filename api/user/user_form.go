package user

import (
	"http/web"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/chatroom"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/o/org/user"

	"gopkg.in/mgo.v2/bson"
)

func NewUserForm(arr []*user.User, userid string) []*chatroom.ChatRoomForm {
	var s = []*chatroom.ChatRoomForm{}
	for _, c := range arr {
		x := &chatroom.ChatRoomForm{}
		x.Person.Name = c.Name
		x.Person.Avatar = c.Avatar
		x.Person.UserID = c.ID
		x.Person.TimeOut = c.StatusActive.TimeOut
		// x.RoomID =
		var query = bson.M{
			"dtime":  0,
			"userid": c.ID,
		}
		cn, err := channel.GetChannel(query)
		web.AssertNil(err)
		x.Person.Job = cn.Job

		ran, err := rank.GetRank(query)
		web.AssertNil(err)
		x.Person.RankName = ran.RankName

		whfb, err := whfeedback.GetWHFeedBack(query)
		web.AssertNil(err)
		x.Person.FeedbackNumber = whfb.NumberFeedbackTime
		x.Person.StarNumber = whfb.PointAVG

		where := bson.M{
			"dtime": 0,
			"$and": []interface{}{
				bson.M{
					"users": bson.M{"$elemMatch": bson.M{"userid": c.ID}},
				},
				bson.M{
					"users": bson.M{"$elemMatch": bson.M{"userid": userid}},
				},
			},
		}
		room, err := chatroom.GetOne(where)
		if err != nil {

		} else {
			x.RoomID = room.ID
		}
		s = append(s, x)
	}
	return s
}

type InformationUser struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Home           user.Location `json:"home" bson:"home"`
	Avatar         string        `json:"avatar"`
	BackGround     string        `json:"background"`
	Job            string        `json:"job"`
	RankName       rank.RankName `json:"rankname" bson:"rankname"`
	LevelUpPercent int           `json:"level_up_percent"`
	StarNumber     float64       `json:"starnumber"`
	FeedbackNumber int           `json:"fb_number"`
	Phone          string        `json:"phone"`
	Email          string        `json:"email"`
	Introduction   string        `json:"introduction"`
}

func NewUserInfo(u *user.User, cn *channel.Channel, ran *rank.Rank, wfhb *whfeedback.WHFeedback) *InformationUser {
	var s = &InformationUser{}
	s.ID = u.ID
	s.Name = u.Name
	s.Home = u.Home
	s.Avatar = u.Avatar
	s.BackGround = u.BackGround
	s.Phone = u.Phone
	s.Email = u.Email

	s.Job = cn.Job
	s.Introduction = cn.Introduction

	s.RankName = ran.RankName
	s.LevelUpPercent = ran.LevelUpPercent

	s.StarNumber = wfhb.PointAVG
	s.FeedbackNumber = wfhb.NumberFeedbackTime

	return s
}
