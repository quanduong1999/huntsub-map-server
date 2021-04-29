package report

import "db/mgo"

type Report struct {
	mgo.BaseModel `bson:",inline"`
	PostID        string `json:"postid" bson:"postid"`
	Label         string `json:"label" bson:"label"`
	// AboutPost     string `json:"aboutpost" bson:"aboutpost"`
	UserID  string `json:"userid" bson:"userid"`
	OwnerID string `json:"ownerid" bson:"ownerid"`
}

func (s *Report) GetLabel() string {
	return s.Label
}
