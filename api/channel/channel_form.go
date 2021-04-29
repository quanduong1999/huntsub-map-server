package channel

import (
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/o/org/user"
)

type InformationUser struct {
	UserID         string        `json:"userid"`
	Name           string        `json:"name"`
	Avatar         string        `json:"avatar"`
	BackGround     string        `json:"background"`
	Job            string        `json:"job"`
	RankName       rank.RankName `json:"rankname" bson:"rankname"`
	Level          int           `json:"level"`
	LevelUpPercent int           `json:"level_up_percent"`
	StarNumber     float64       `json:"starnumber"`
	FeedbackNumber int           `json:"fb_number"`
	Phone          string        `json:"phone"`
	Email          string        `json:"email"`
	Location       user.Location `json:"location" bson:"location"`
	Followed       bool          `json:"followed"`
	FollowID       string        `json:"followid"`
	RoomID         string        `json:"roomid"`
}

type JobList struct {
	Count int    `json:"count"`
	Job   string `json:"job"`
}

func NewUserInfo(u *user.User, cn *channel.Channel, ran *rank.Rank, wfhb *whfeedback.WHFeedback, followed bool, followID, roomID string) *InformationUser {
	var s = &InformationUser{}
	s.UserID = u.ID
	s.Name = u.Name
	s.Avatar = u.Avatar
	s.BackGround = u.BackGround
	s.Phone = u.Phone
	s.Email = u.Email
	s.Location = u.Location

	s.Job = cn.Job

	s.RankName = ran.RankName
	s.LevelUpPercent = ran.LevelUpPercent

	s.StarNumber = wfhb.PointAVG
	s.FeedbackNumber = wfhb.NumberFeedbackTime

	s.Followed = followed
	s.FollowID = followID
	s.RoomID = roomID
	return s
}
