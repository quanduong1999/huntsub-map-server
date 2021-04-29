package user

import "db/mgo"

type UserData struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string `json:"userid" bson:"userid"`
}

// func (s *UserData) GetPointAVG() float64 {
// 	return s.PointAVG
// }
