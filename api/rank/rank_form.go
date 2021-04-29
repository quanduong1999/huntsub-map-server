package rank

import (
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/rank"
)

type RankForm struct {
	Rank   *rank.Rank `json:"rank"`
	Avatar string     `json:"avatar"`
	UserID string     `json:"userid"`
	Name   string     `json:"name"`
	Job    string     `json:"job"`
}

func NewRankForm(ran *rank.Rank) *RankForm {
	var s = &RankForm{}

	s.Rank = ran
	usr, err := cache_user.Get(ran.UserID)
	web.AssertNil(err)
	s.Avatar = usr.Avatar
	s.UserID = usr.ID
	//Get Channel
	ch, err := channel.GetChannel(map[string]interface{}{
		"dtime":  0,
		"userid": ran.UserID,
	})
	web.AssertNil(err)
	s.Job = ch.Job
	s.Name = usr.Name

	return s
}
