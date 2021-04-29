package rank

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/org/user"
)

type Rank struct {
	mgo.BaseModel    `bson:",inline"`
	UserID           string        `json:"userid" bson:"userid"`
	RankName         RankName      `json:"rankname" bson:"rankname"`
	Radius           int           `json:"radius" bson:"radius"`
	PostNumber       int           `json:"postnumber" bson:"postnumber"`
	ExperienceNumber int           `json:"experiencenumber" bson:"experiencenumber"`
	LevelUpPercent   int           `json:"level_up_percent" bson:"level_up_percent"`
	Location         user.Location `json:"location" bson:"location"`
	LikeNumber       int           `json:"likenumber" bson:"likenumber"`
	CommentNumber    int           `json:"commentnumber" bson:"commentnumber"`
	ShareNumber      int           `json:"sharenumber" bson:"sharenumber"`
	Job              string        `json:"job" bson:"job"`
}

func (s *Rank) GetRankName() RankName {
	return s.RankName
}

func (s *Rank) GetRadius() int {
	return s.Radius
}

func (s *Rank) GetPostNumber() int {
	return s.PostNumber
}

func (s *Rank) GetExperienceNumber() int {
	return s.ExperienceNumber
}

func (s *Rank) GetLevelUpPercent() int {
	return s.LevelUpPercent
}

func (s *Rank) GetLocation() user.Location {
	return s.Location
}

func (s *Rank) GetLikeNumber() int {
	return s.LikeNumber
}

func (s *Rank) GetCommentNumber() int {
	return s.CommentNumber
}

func (s *Rank) GetShareNumber() int {
	return s.ShareNumber
}
