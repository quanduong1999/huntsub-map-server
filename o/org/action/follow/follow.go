package follow

import "db/mgo"

type Follow struct {
	mgo.BaseModel `bson:",inline"`
	IsFollow      bool   `json:"isfollow" bson:"isfollow"`
	UserID        string `json:"userid" bson:"userid"`
	PersonID      string `json:"personid" bson:"personid"`
}

func (s *Follow) GetIsFollow() bool {
	return s.IsFollow
}
