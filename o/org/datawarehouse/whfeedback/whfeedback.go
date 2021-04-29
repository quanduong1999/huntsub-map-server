package whfeedback

import "db/mgo"

type WHFeedback struct {
	mgo.BaseModel      `bson:",inline"`
	UserID             string  `json:"userid" bson:"userid"`
	PointAVG           float64 `json:"pointavg" bson:"pointavg"`
	NumberFeedbackTime int     `json:"numberfeedbacktime" bson:"numberfeedbacktime"`
	FiveStar           int     `json:"fivestar" bson:"fivestar"`
	FourStar           int     `json:"fourstar" bson:"fourstar"`
	ThreeStar          int     `json:"threestar" bson:"threestar"`
	TwoStar            int     `json:"twostar" bson:"twostar"`
	OneStar            int     `json:"onestar" bson:"onestar"`
}

func (s *WHFeedback) GetPointAVG() float64 {
	return s.PointAVG
}
